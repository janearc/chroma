package main

import (
	"fmt"
	"sort"

	"github.com/janearc/libtheme-css/primitives/functions"
	"github.com/janearc/libtheme-css/primitives/swatch"
	"github.com/janearc/libtheme-css/spaces/ok"
	"github.com/janearc/libtheme-css/spaces/srgb"
)

// gradientWidth is how many columns the gradients take from the right
// of the pane: two of ground between them and the text, then two of the
// colourway's own colours, two of the wheel mixed in oklab, and two of
// the same wheel mixed in srgb.
const gradientWidth = 8

// wheel is the hue sweep the gradient draws, top to bottom: the screen's
// three lamps and the three colours two of them make together, and back
// to red, so the bottom meets the top.
var wheel = []string{"#ff0000", "#ffff00", "#00ff00", "#00ffff", "#0000ff",
	"#ff00ff", "#ff0000"}

// sweep is the wheel as a ramp mixed in one space.
func sweep(in functions.Mixer) functions.Ramp {
	swatches := make([]swatch.Swatch, len(wheel))
	for position, hex := range wheel {
		swatches[position] = srgb.MustHex(hex).Swatch()
	}
	return functions.Even(in, swatches...)
}

// lampCode is the escape that sets a swatch as the foreground (layer 38)
// or the ground (layer 48). a swatch the screen can't make is pulled in
// by chroma, keeping its lightness and hue, the way libtheme fits any
// colour; oklab's straight line between two lamps can bow outside them.
func lampCode(layer int, sample swatch.Swatch) string {
	if !srgb.In(sample) {
		fitted, _ := ok.Fit(ok.FromSwatch(sample).Polar(), srgb.In)
		sample = fitted.Rect().Swatch()
	}
	levels, _ := srgb.FromSwatch(sample)
	r, g, b := levels.Bytes()
	return fmt.Sprintf("\x1b[%d;2;%d;%d;%dm", layer, r, g, b)
}

// gradientRows are the strips down the right of the pane, a string of six
// cells for each row, top to bottom: the colourway's own colours, a
// colour to a cell, since a click takes the cell it lands in and those
// colours are each distinct.
//
// Then the wheel in oklab, then in srgb beside it, in half blocks so each
// row shows two steps of a smooth sweep.
func gradientRows(height int, theme functions.Ramp) []string {
	steps := float64(max(2*height-1, 1))
	rows := make([]string, height)
	for row := range rows {
		top, bottom := float64(2*row)/steps, float64(2*row+1)/steps
		rows[row] = lampCode(48, theme.At(top)) + "  "
		ramps := []functions.Ramp{sweep(ok.Mix), sweep(srgb.Mix)}
		for _, ramp := range ramps {
			rows[row] += lampCode(38, ramp.At(top)) +
				lampCode(48, ramp.At(bottom)) + "▀▀"
		}
		rows[row] += "\x1b[0m"
	}
	return rows
}

// palette is the colourway's own colours, each it sets once: the
// coloured ones by their hue round the wheel, so they sit beside the
// wheel's own, then the greys by lightness, light first. Each colour's
// oklch is worked out once, not in every comparison of the sort.
func palette(set *tuning) []colour {
	type entry struct {
		value colour
		polar ok.OKLCH
	}
	seen := map[string]bool{}
	coloured, greys := []entry{}, []entry{}
	for _, name := range set.roles.order {
		value, _ := set.roles.get(name)
		if seen[value.hex()] {
			continue
		}
		seen[value.hex()] = true
		each := entry{value, polarOf(value)}
		if each.polar.C < ok.Eye {
			greys = append(greys, each)
		} else {
			coloured = append(coloured, each)
		}
	}
	sort.SliceStable(coloured, func(first, second int) bool {
		return coloured[first].polar.H < coloured[second].polar.H
	})
	sort.SliceStable(greys, func(first, second int) bool {
		return greys[first].polar.L > greys[second].polar.L
	})
	colours := []colour{}
	for _, each := range append(coloured, greys...) {
		colours = append(colours, each.value)
	}
	return colours
}

// polarOf is a colour in oklch.
func polarOf(value colour) ok.OKLCH {
	return ok.FromSwatch(swatchOf(value)).Polar()
}

// themeRamp is the palette as a strip, mixed in oklab, top to bottom.
func themeRamp(set *tuning) functions.Ramp {
	swatches := []swatch.Swatch{}
	for _, value := range palette(set) {
		swatches = append(swatches, swatchOf(value))
	}
	switch len(swatches) {
	case 0:
		swatches = []swatch.Swatch{swatchOf(set.shown("ground"))}
		fallthrough
	case 1:
		swatches = append(swatches, swatches[0])
	}
	return functions.Even(ok.Mix, swatches...)
}

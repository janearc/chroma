package main

import (
	"fmt"
	"image"
	"os"
	"sort"

	"github.com/janearc/libtheme-css/spaces/ok"
	"github.com/janearc/libtheme-css/spaces/srgb"
)

// tile is a square of a picture averaged in linear light, with where it is
// and its colour in oklch.
type tile struct {
	x, y  int
	hex   string
	polar ok.OKLCH
}

// tiles averages a picture in squares of a side, skipping any square that
// is partly space, and prints the most colourful of them whose hue lies
// between two angles, so a colour that large boxes average away can be
// found where it is.
func tiles(path string, side int, from, to float64, count int) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	picture, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	at, bounds := levelsOf(picture), picture.Bounds()
	found := []tile{}
	keepAll := func(int, int) bool { return true }
	for y := bounds.Min.Y; y+side <= bounds.Max.Y; y += side {
		for x := bounds.Min.X; x+side <= bounds.Max.X; x += side {
			cell := box{"", x, y, x + side, y + side}
			square := averaged(at, cell, keepAll)
			if square.count < side*side {
				continue
			}
			value := srgb.FromLight(square.r, square.g, square.b)
			polar := ok.FromSwatch(value).Polar()
			if within(polar.H, from, to) {
				chip := tile{x, y, hex(value), polar}
				found = append(found, chip)
			}
		}
	}
	sort.Slice(found, func(i, k int) bool {
		return found[i].polar.C > found[k].polar.C
	})
	for _, each := range found[:min(count, len(found))] {
		fmt.Printf("x %4d y %4d  %s  L %.2f C %.3f H %3.0f\n",
			each.x, each.y, each.hex,
			each.polar.L, each.polar.C, each.polar.H)
	}
	fmt.Printf("%d squares in that hue\n", len(found))
	return nil
}

// within is whether a hue lies between two angles, going round past 360
// when the first is the larger.
func within(hue, from, to float64) bool {
	if from <= to {
		return hue >= from && hue <= to
	}
	return hue >= from || hue <= to
}

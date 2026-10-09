package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/janearc/libtheme-css/spaces/ok"
)

// the picker's rows: a readout, the colourway's own colours, a grid of
// hue across and lightness down, and a row of chroma.
const (
	paletteRows = 2
	gridRows    = 6
	pickerRows  = 1 + paletteRows + gridRows + 1
)

// pickerColour is the colour under a cell of the picker, counted from the
// picker's first row and its left edge in cells of two columns, for the
// colour steered now; false where a cell holds none.
func (screen *model) pickerColour(row, cell, cells int) (colour, bool) {
	now := polarOf(screen.set.shown(screen.stops[screen.steering].name))
	switch {
	case row >= 1 && row <= paletteRows:
		colours := palette(screen.set)
		index := (row-1)*cells + cell
		if index >= len(colours) {
			return colour{}, false
		}
		return colours[index], true
	case row > paletteRows && row <= paletteRows+gridRows:
		step := float64(row - paletteRows - 1)
		lightness := 0.92 - step*0.72/float64(gridRows-1)
		hue := 360 * float64(cell) / float64(cells)
		return fittedPolar(ok.OKLCH{
			L: lightness, C: math.Max(now.C, 0.06), H: hue}), true
	case row == pickerRows-1:
		chroma := 0.37 * float64(cell) / float64(max(cells-1, 1))
		return fittedPolar(
			ok.OKLCH{L: now.L, C: chroma, H: now.H}), true
	}
	return colour{}, false
}

// fittedPolar is an oklch colour pulled into what the screen can make,
// by chroma, as the strips are.
func fittedPolar(polar ok.OKLCH) colour {
	return fitted(polar.Rect().Swatch())
}

// picker draws the picker for the colour steered, a row each, every row
// width wide: the readout, then cells of two columns, a cell that would
// not read against the field's ground marked with a dot.
func (screen *model) picker(width int, tint string) []string {
	set, here := screen.set, screen.stops[screen.steering]
	value, against := set.shown(here.name), set.against(here)
	readout := fmt.Sprintf("  pick %s: %s  %.1f:1   · will not read "+
		"here   o or escape closes",
		here.label, value.hex(), contrast(value, against))
	short := max(width-len([]rune(readout)), 0)
	rows := []string{tint + readout + strings.Repeat(" ", short)}
	cells := width / 2
	for row := 1; row < pickerRows; row++ {
		var drawn strings.Builder
		for cell := range cells {
			value, has := screen.pickerColour(row, cell, cells)
			if !has {
				drawn.WriteString(tint + "  ")
				continue
			}
			mark := "  "
			if contrast(value, against) < reading.Contrast.ChipMin {
				dotInk := inkCode(controlInk(against, value))
				mark = dotInk + "· "
			}
			drawn.WriteString(groundCode(value) + mark + "\x1b[49m")
		}
		rows = append(rows,
			drawn.String()+strings.Repeat(" ", width-2*cells))
	}
	return rows
}

// pickAt is a click on the picker: the colour under it becomes the colour
// steered, and the source's own; a click on the readout or an empty cell
// does nothing. It says whether the click was on the picker.
func (screen *model) pickAt(x, y int) bool {
	top := screen.height - pickerRows
	if y < top {
		return false
	}
	cells := (screen.width - 6) / 2
	value, has := screen.pickerColour(y-top, x/2, cells)
	if has && x/2 < cells {
		screen.pick(value)
	}
	return true
}

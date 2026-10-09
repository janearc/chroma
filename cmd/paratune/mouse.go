package main

import (
	tea "charm.land/bubbletea/v2"
	"fmt"
	"github.com/charmbracelet/x/ansi"
	"github.com/janearc/libtheme-css/primitives/functions"
	"github.com/janearc/libtheme-css/primitives/swatch"
	"github.com/janearc/libtheme-css/spaces/ok"
	"github.com/janearc/libtheme-css/spaces/srgb"
	"slices"
	"strings"
	"unicode/utf8"
)

// mouse is what a click or the wheel does. A line of the sample steers a
// colour it shows, the next one each time it is clicked again; the wheel
// moves the colour steered, lighter and darker, and with shift, or sideways,
// warmer and cooler; the gradient gives the colour steered the hue clicked.
//
// While the menu or the keys are open the mouse does nothing.
func (screen *model) mouse(message tea.Msg) tea.Cmd {
	if screen.menu != nil || screen.keys {
		return nil
	}
	screen.note = ""
	switch message := message.(type) {
	case tea.MouseWheelMsg:
		shift := message.Mod&tea.ModShift != 0
		switch {
		case message.Button == tea.MouseWheelUp && shift,
			message.Button == tea.MouseWheelRight:
			screen.nudge(warmer)
		case message.Button == tea.MouseWheelDown && shift,
			message.Button == tea.MouseWheelLeft:
			screen.nudge(cooler)
		case message.Button == tea.MouseWheelUp:
			screen.nudge(lighter)
		case message.Button == tea.MouseWheelDown:
			screen.nudge(darker)
		}
	case tea.MouseClickMsg:
		if message.Button == tea.MouseLeft {
			screen.click(message.X, message.Y)
		}
	}
	return nil
}

// click is a left click at a cell of the pane.
func (screen *model) click(x, y int) {
	switch {
	case x >= screen.width-4:
		screen.pick(gradientColour(screen.height, y,
			x >= screen.width-2))
	case x >= screen.width-6:
		screen.pick(stripColour(themeRamp(screen.set),
			screen.height, y))
	case screen.picking && screen.pickAt(x, y):
	case y < screen.placed.rows && x >= screen.placed.left &&
		x < screen.placed.left+screen.placed.width:
		screen.keys = y == screen.placed.help
		screen.picking = screen.picking || y == screen.placed.field
	case y < len(screen.drawn):
		screen.steerAt(screen.drawn[y], x-2)
	}
}

// steerAt steers the colour drawn at a cell of a line, counted from the
// line's first cell.
//
// The role, of those the line shows and the page's ink and ground, drawn
// in the colour under the pointer, its letter's colour, or its ground's
// on a space or behind a letter cut out in the page's own ground, as a
// search match is.
//
// When several roles are drawn in that colour, a click on it again steers
// the next of them. Where no role has that colour, the click cycles
// through the line's colours as before.
func (screen *model) steerAt(clicked line, cell int) {
	set := screen.set
	ink := set.shown(set.inkName(screen.page))
	ground := set.shown(set.groundName(screen.page))
	fg, bg, space, found := cellColours(clicked.text, cell, ink, ground)
	if !found {
		// past the end of the line, the row shows its fill, or
		// the page's ground: a click there is on the ground.
		codes := strings.TrimPrefix(clicked.fill, "\x1b[")
		_, bg, _ = applyCodes(strings.TrimSuffix(codes, "m"),
			ink, ground, false, ink, ground)
		fg, space = ink, true
	}
	wanted := []colour{fg, bg}
	switch {
	case space:
		wanted = []colour{bg}
	case fg.hex() == ground.hex() && bg.hex() != ground.hex():
		wanted = []colour{bg, fg}
	}
	names := append(append([]string{}, clicked.shows...),
		set.inkName(screen.page), set.groundName(screen.page))
	for _, want := range wanted {
		matching := []int{}
		for _, name := range names {
			stop := screen.stopOf(name)
			if stop >= 0 && set.shown(name).hex() == want.hex() &&
				!slices.Contains(matching, stop) {
				matching = append(matching, stop)
			}
		}
		if len(matching) > 0 {
			next := 0
			at := slices.Index(matching, screen.steering)
			if at >= 0 {
				next = (at + 1) % len(matching)
			}
			screen.steering = matching[next]
			return
		}
	}
	screen.steerLine(clicked)
}

// cellColours is the letter and ground colour of one cell of a drawn line,
// following its colour codes from the start: the ink and ground until a
// code says otherwise, swapped under reverse video. It says whether the
// cell holds a space, and whether the line reaches the cell at all.
func cellColours(text string, cell int, ink, ground colour) (
	fg, bg colour, space, found bool) {
	fg, bg = ink, ground
	reverse, at := false, 0
	for len(text) > 0 {
		if strings.HasPrefix(text, "\x1b[") {
			end := strings.IndexByte(text, 'm')
			if end < 0 {
				break
			}
			fg, bg, reverse = applyCodes(text[2:end],
				fg, bg, reverse, ink, ground)
			text = text[end+1:]
			continue
		}
		letter, size := utf8.DecodeRuneInString(text)
		wide := max(ansi.StringWidth(string(letter)), 1)
		if cell >= at && cell < at+wide {
			if reverse {
				fg, bg = bg, fg
			}
			return fg, bg, letter == ' ', true
		}
		at += wide
		text = text[size:]
	}
	return fg, bg, false, false
}

// applyCodes follows one set of colour codes: a reset to the ink and
// ground, true colour for letter and ground, the default of each, and
// reverse video on and off.
func applyCodes(codes string, fg, bg colour, reverse bool,
	ink, ground colour) (colour, colour, bool) {
	parts := strings.Split(codes, ";")
	for index := 0; index < len(parts); index++ {
		switch parts[index] {
		case "", "0":
			fg, bg, reverse = ink, ground, false
		case "7":
			reverse = true
		case "27":
			reverse = false
		case "39":
			fg = ink
		case "49":
			bg = ground
		case "38", "48":
			if index+4 < len(parts) && parts[index+1] == "2" {
				var value colour
				fmt.Sscan(parts[index+2], &value.r)
				fmt.Sscan(parts[index+3], &value.g)
				fmt.Sscan(parts[index+4], &value.b)
				if parts[index] == "38" {
					fg = value
				} else {
					bg = value
				}
				index += 4
			}
		}
	}
	return fg, bg, reverse
}

// steerLine steers a colour a line shows, other than the page's own ink
// and ground: the first, or the next after the one already steered.
func (screen *model) steerLine(clicked line) {
	everywhere := map[string]bool{screen.set.groundName(screen.page): true,
		screen.set.inkName(screen.page): true}
	names := []string{}
	for _, name := range clicked.shows {
		if !everywhere[name] && screen.stopOf(name) >= 0 {
			names = append(names, name)
		}
	}
	if len(names) == 0 {
		return
	}
	next := names[0]
	for position, name := range names {
		if name == screen.stops[screen.steering].name {
			next = names[(position+1)%len(names)]
		}
	}
	screen.steering = screen.stopOf(next)
}

// stopOf is the index of a role's stop on the page showing, or -1.
func (screen *model) stopOf(name string) int {
	for index, here := range screen.stops {
		if here.page == screen.page && here.name == name {
			return index
		}
	}
	return -1
}

// pick sets the colour steered to one chosen whole, as the gradient
// gives it, and makes it the source's own.
func (screen *model) pick(value colour) {
	name := screen.stops[screen.steering].name
	screen.set.roles.put(name, value)
	screen.set.dirty = true
	screen.repaint()
}

// stripColour is the colour a strip draws at a row, at the top half of
// the cell, fitted as gradientColour fits it.
func stripColour(ramp functions.Ramp, height, row int) colour {
	steps := float64(max(2*height-1, 1))
	return fitted(ramp.At(float64(2*max(row, 0)) / steps))
}

// fitted is a swatch as paratune's colour, pulled into what the screen
// can make by chroma, as lampCode pulls it.
func fitted(sample swatch.Swatch) colour {
	if !srgb.In(sample) {
		polar, _ := ok.Fit(ok.FromSwatch(sample).Polar(), srgb.In)
		sample = polar.Rect().Swatch()
	}
	return colourOf(sample)
}

// gradientColour is the colour the gradient draws at a row: oklab's
// side or srgb's, at the top half of the cell, where a click lands
// nearer, fitted into what the screen can make as lampCode fits it.
func gradientColour(height, row int, srgbSide bool) colour {
	mixer := ok.Mix
	if srgbSide {
		mixer = srgb.Mix
	}
	return stripColour(sweep(mixer), height, row)
}

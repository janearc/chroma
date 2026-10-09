package main

import (
	"strings"

	libcolour "github.com/janearc/libreadme/colour"
	"github.com/janearc/libreadme/profile"
)

// reading is libreadme's current reader profile, compiled into libreadme:
// what a small labelled colour needs to be read, and the fill it is drawn
// on when the ground cannot give it that.
var reading = currentProfile()

// currentProfile is libreadme's current profile, without which paratune
// does not run.
func currentProfile() profile.Profile {
	current, err := profile.Current()
	if err != nil {
		panic("paratune: libreadme has no current profile: " +
			err.Error())
	}
	return current
}

// field is the role steered and its colour, written in that colour: on
// the float's ground when the reader can read it there, and when it
// cannot, on libreadme's fill for it, the colour's own hue drained of
// saturation, its lightness moved until the label reads.
func field(label string, value, ground colour, tint string) string {
	text := inkCode(value) + label + ": " + value.hex()
	if contrast(value, ground) >= reading.Contrast.ChipMin {
		return text + tint
	}
	fill, _ := libcolour.FillFor(value.reader(), ground.reader(),
		reading.Contrast.ChipMin, reading.Hue.ChipFillSaturation,
		reading.Hue.ChipFillLightnessStep)
	return groundCode(colour{fill.R, fill.G, fill.B}) + " " + text + " " +
		groundCode(ground) + tint
}

// controlInk is the ink paratune's own marks are drawn in, the float, the
// menu and the markers: the page's ink when it reads against the ground
// with room to spare, and otherwise moved there by libreadme, its
// lightness alone.
//
// The marks must read whatever the page's ink has been tuned to, so when
// no lightness of the ink's hue gets there, they are black or white,
// whichever reads more.
func controlInk(ink, ground colour) colour {
	band := libcolour.Band{Low: reading.Contrast.ChipMin + 1.5,
		High:   reading.Contrast.High,
		Centre: reading.Contrast.ChipMin + 3}
	if band.Contains(contrast(ink, ground)) {
		return ink
	}
	against := []libcolour.RGB{ground.reader()}
	if fitted, reached := band.Into(ink.reader(), against); reached {
		return colour{fitted.R, fitted.G, fitted.B}
	}
	black, white := colour{}, colour{255, 255, 255}
	if contrast(black, ground) > contrast(white, ground) {
		return black
	}
	return white
}

// inkCode is the escape that draws text in a colour.
func inkCode(value colour) string {
	return "\x1b[38;2;" + value.levels() + "m"
}

// fgName and bgName are a stop's field: the role drawn, and the role it
// is drawn on, when the field has one of its own; empty when it has
// none, as text on the page's own ground has no ground to copy.
func (set *tuning) fgName(here stop) string {
	switch {
	case !here.ground:
		return here.name
	case here.partner != "":
		return here.partner
	}
	return set.inkName(here.page)
}

// bgName is the role a stop's field is drawn on, empty when it has none.
func (set *tuning) bgName(here stop) string {
	switch {
	case here.ground:
		return here.name
	case grounds[here.partner]:
		return here.partner
	}
	return ""
}

// registers hold a copied foreground and a copied ground, for pasting
// into another field: c and v for the foreground, d and f for the ground,
// y and p for both.
type registers struct {
	fg, bg         colour
	haveFg, haveBg bool
}

// use carries out a register key on the stop, and says what happened.
func (held *registers) use(key string, here stop, set *tuning) string {
	fg, bg := set.fgName(here), set.bgName(here)
	switch key {
	case "c":
		return held.copy(&held.fg, &held.haveFg, fg, "foreground", set)
	case "v":
		return held.paste(held.fg, held.haveFg, fg, "foreground", set)
	case "d":
		return held.copy(&held.bg, &held.haveBg, bg, "ground", set)
	case "f":
		return held.paste(held.bg, held.haveBg, bg, "ground", set)
	case "y":
		return held.copy(&held.fg, &held.haveFg, fg,
			"foreground", set) + ", " +
			held.copy(&held.bg, &held.haveBg, bg, "ground", set)
	case "p":
		return held.paste(held.fg, held.haveFg, fg,
			"foreground", set) + ", " +
			held.paste(held.bg, held.haveBg, bg, "ground", set)
	}
	return ""
}

// copy puts a role's colour in a register.
func (held *registers) copy(into *colour, have *bool, name, what string,
	set *tuning) string {
	if name == "" {
		return "no " + what + " here"
	}
	*into, *have = set.shown(name), true
	return "copied " + what + " " + into.hex()
}

// paste sets a role from a register, making it the source's own.
func (held *registers) paste(value colour, have bool, name, what string,
	set *tuning) string {
	switch {
	case !have:
		return "no " + what + " copied"
	case name == "":
		return "no " + what + " here"
	}
	set.roles.put(name, value)
	set.dirty = true
	return "pasted " + what + " " + value.hex()
}

// copied is what the registers hold, as the float says it.
func (held *registers) copied() string {
	parts := []string{}
	if held.haveFg {
		parts = append(parts, "fg "+held.fg.hex())
	}
	if held.haveBg {
		parts = append(parts, "bg "+held.bg.hex())
	}
	return strings.Join(parts, "  ")
}

// keysText is the panel ? opens in place of the page: every key, and what
// the mouse does.
var keysText = []string{
	"keys: ? or escape closes this",
	"",
	"  arrows          lighter, darker, cooler, warmer",
	"  tab, shift-tab  the next colour on the page, and back",
	"  n, shift-n      the next colour in the colourway's own palette, " +
		"and back",
	"  r               every text colour here into your band: " +
		"dim up, glare down",
	"  o               the picker: your colours, a grid, chroma; " +
		"click to take",
	"  g, then g       mark this colour, steer another, g joins them; " +
		"g ungroups",
	"  - and =         night: the ground dimmer, contrast and chroma up; " +
		"= back",
	"  + and _         every colour's chroma, up and down",
	"  > and <         every text colour lighter and darker, grounds left",
	"  c v             copy and paste a field's foreground",
	"  d f             the same for its ground; y p for both",
	"  [               the menu: a colourway, then a page",
	"  backslash       ours, then theirs: charm's colorprofile at 256, " +
		"at 16",
	"  u               undo a step: each key or click that moved a colour",
	"  enter           keep the line, s saves, i installs",
	"  escape          leaves paratune; here, it closes this",
	"",
	"mouse",
	"",
	"  click a line    steer a colour it shows, the next each click",
	"  wheel           lighter and darker; with shift, warmer and cooler",
	"  the strips      on the right: click one to take its colour",
	"  the float       click the field for the picker, ? help for this",
}

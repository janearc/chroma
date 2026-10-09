package main

import (
	"fmt"
	"github.com/janearc/libtheme-css/colourway"
	"github.com/janearc/libtheme-css/spaces/ok"
	"slices"
	"strings"
)

// whole moves every colour the colourway sets at once, in oklch: its
// lightness by one factor and its chroma by another, hue held, each
// pulled back inside what the screen can show. = and - lift and dim,
// + and _ add and take away chroma. Roles the colourway leaves to their
// fallbacks follow the ones they fall back to.
func (screen *model) whole(lightness, chroma float64) {
	set := screen.set
	for _, name := range set.roles.order {
		value, _ := set.roles.get(name)
		polar := ok.FromSwatch(swatchOf(value)).Polar()
		polar.L = min(max(polar.L*lightness, 0), 1)
		polar.C = max(polar.C*chroma, 0)
		set.roles.put(name, fitted(polar.Rect().Swatch()))
	}
	set.dirty = true
	screen.repaint()
}

// night is - and, with step's sign turned, =: the grounds dim, every colour
// gains chroma, and each text colour moves away from the ground, darker under
// dark text or lighter over a dark ground, until its contrast with the dimmed
// ground is a little above what it was, or as far as the screen goes.
//
// One key, for when the room has gone dark.
func (screen *model) night(step float64) {
	set := screen.set
	oldGround := set.shown("ground")
	for _, name := range set.roles.order {
		if grounds[name] && name != "black" {
			screen.shift(name, 1-0.035*step, 1+0.05*step)
		}
	}
	newGround := set.shown("ground")
	for _, name := range set.roles.order {
		if grounds[name] && name != "black" {
			continue
		}
		value, _ := set.roles.get(name)
		target := contrast(value, oldGround) * (1 + 0.03*step)
		away := 0.98
		if value.luminance() > newGround.luminance() {
			away = 1.02
		}
		if step < 0 {
			away = 1 / away
		}
		screen.shift(name, 1, 1+0.05*step)
		for range 60 {
			now, _ := set.roles.get(name)
			reached := contrast(now, newGround) >= target
			if step < 0 {
				reached = contrast(now, newGround) <= target
			}
			if reached {
				break
			}
			screen.shift(name, away, 1)
		}
	}
	set.dirty = true
	screen.repaint()
}

// shift moves one role's lightness and chroma by factors, in oklch, hue
// held, pulled back inside what the screen can show.
func (screen *model) shift(name string, lightness, chroma float64) {
	value, _ := screen.set.roles.get(name)
	polar := ok.FromSwatch(swatchOf(value)).Polar()
	polar.L = min(max(polar.L*lightness, 0), 1)
	polar.C = max(polar.C*chroma, 0)
	screen.set.roles.put(name, fitted(polar.Rect().Swatch()))
}

// walk is n walking the palette for one role: the palette as it was when
// the walk began, with the role's colour then in it, and where it is now.
// Moving the role overwrites its colour, which may leave the palette, so
// the walk keeps its own copy to come back to.
type walk struct {
	name    string
	colours []colour
	at      int
}

// paletteStep is n and shift-n: the colour steered becomes the next of the
// colourway's own colours, or the one before, and the source's own. A walk
// goes on while the same role is steered and nothing else has moved it;
// otherwise a new one begins where the colour is now.
func (screen *model) paletteStep(step int) {
	name := screen.stops[screen.steering].name
	now := screen.set.shown(name)
	current := screen.walking
	if current == nil || current.name != name ||
		current.colours[current.at].hex() != now.hex() {
		current = walkFrom(name, now, palette(screen.set))
		screen.walking = current
	}
	count := len(current.colours)
	current.at = (current.at + step + count) % count
	screen.set.roles.put(name, current.colours[current.at])
	screen.set.dirty = true
	screen.note = fmt.Sprintf("palette %d of %d", current.at+1, count)
	screen.repaint()
}

// walkFrom is a walk beginning at a colour: the palette, with the colour
// put first when it is not already in it, so one step back, or all the
// way round, returns to it.
func walkFrom(name string, now colour, colours []colour) *walk {
	for index, value := range colours {
		if value.hex() == now.hex() {
			return &walk{name, colours, index}
		}
	}
	return &walk{name, append([]colour{now}, colours...), 0}
}

// isGround is whether a role is a ground, something drawn behind text:
// the page's grounds, a program's -ground roles, and the general grounds
// under a program's own name.
func isGround(name string) bool {
	return grounds[name] || strings.HasSuffix(name, "-ground") ||
		slices.Contains(colourway.Grounds,
			strings.TrimPrefix(name, "nvim-"))
}

// textLightness is > and <: every colour the colourway sets that is not a
// ground, lighter or darker by the same step of oklab lightness, hue and
// chroma held, each fitted into what the screen can make.
func (screen *model) textLightness(step float64) {
	set, moved := screen.set, 0
	for _, name := range set.roles.order {
		if isGround(name) {
			continue
		}
		value, _ := set.roles.get(name)
		polar := polarOf(value)
		polar.L = min(max(polar.L+step, 0), 1)
		set.roles.put(name, fittedPolar(polar))
		moved++
	}
	set.dirty = true
	shade := map[bool]string{true: "lighter", false: "darker"}[step > 0]
	screen.note = fmt.Sprintf("%d text colours %s", moved, shade)
	screen.repaint()
}

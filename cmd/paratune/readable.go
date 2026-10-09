package main

import (
	"fmt"

	libcolour "github.com/janearc/libreadme/colour"
)

// proseBand is the reader's band for prose, from libreadme's current
// profile: the floor, the ceiling, and the value to aim for.
func proseBand() libcolour.Band {
	return libcolour.Band{Low: reading.Contrast.Low,
		High: reading.Contrast.High, Centre: reading.Contrast.Centre}
}

// tokenBand is the band for a colour scanned rather than read, a heading
// or a keyword: the profile's floor for chips up to the prose ceiling,
// aiming just over the floor, so a token moves only as far as it must and
// stays a colour of its own rather than turning the ink's.
func tokenBand() libcolour.Band {
	return libcolour.Band{Low: reading.Contrast.ChipMin,
		High:   reading.Contrast.High,
		Centre: reading.Contrast.ChipMin + 0.5}
}

// readable is r: every text colour on the page showing put where the
// reader can read it against its ground, by libreadme's band, its
// lightness alone moved, hue and saturation kept.
//
// The page's ink is prose and goes to the prose band; the rest to the
// token band. A colour already inside its band is left as it is; a colour
// no lightness can bring in is left too, and counted.
func (screen *model) readable() {
	set, ink := screen.set, screen.set.inkName(screen.page)
	moved, unreachable := 0, 0
	for _, here := range screen.stops {
		if here.page != screen.page || here.ground {
			continue
		}
		value, ground := set.shown(here.name), set.against(here)
		band := tokenBand()
		if here.name == ink {
			band = proseBand()
		}
		if band.Contains(contrast(value, ground)) {
			continue
		}
		against := []libcolour.RGB{ground.reader()}
		fitted, reached := band.Into(value.reader(), against)
		if !reached {
			unreachable++
			continue
		}
		set.roles.put(here.name, colour{fitted.R, fitted.G, fitted.B})
		moved++
	}
	if moved > 0 {
		set.dirty = true
		screen.repaint()
	}
	screen.note = fmt.Sprintf("readable: %d moved", moved)
	if unreachable > 0 {
		screen.note += fmt.Sprintf(", %d no lightness can bring in",
			unreachable)
	}
}

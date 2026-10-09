package main

import (
	"math"

	"github.com/janearc/libtheme-css/colourway"
	"github.com/janearc/libtheme-css/css"
	"github.com/janearc/libtheme-css/primitives/swatch"
	"github.com/janearc/libtheme-css/spaces/ok"
	"github.com/janearc/libtheme-css/spaces/radiation"
	"github.com/janearc/libtheme-css/spaces/srgb"
)

// contrasts are what each kind of colour is lit to against the ground, from
// her reference dark profile: prose aims at 11 inside 10 to 12.5, and a
// colour scanned rather than read needs at least 4.5.
const (
	inkContrast       = 11
	colourContrast    = 6
	brightContrast    = 8.5
	whiteContrast     = 9
	brightWhite       = 12
	dimContrast       = 4.6
	bandContrast      = 1.25
	selectionContrast = 1.6
)

// screens are the screen's own colours, whose wavelengths the terminal's
// colours are taken at.
var screens = map[string]string{
	"red": "#ff0000", "yellow": "#ffff00", "green": "#00ff00",
	"cyan": "#00ffff", "blue": "#0000ff",
}

// visible carries visible light onto itself, so its Of is a colour's
// dominant wavelength: the one whose own colour has the same hue from white.
var visible = radiation.Profile{Name: "visible", Source: radiation.Visible,
	Render: radiation.Visible}

// fermiLong is a colourway from a long burst seen through the fermi
// profile: the deep indigo at its violet end for the ground, the burst's
// own colour for the ink, and each of the terminal's colours the
// wavelength the screen's own colour points at, lit to a contrast.
func fermiLong() colourway.Source {
	profile, shape := radiation.Fermi, bursts[0]
	ground := indigo(shape.strip(profile))
	burst := shape.colour(profile)
	roles := css.New()
	set := func(name string, value swatch.Swatch) {
		roles.Set(name, drawn(value))
	}
	set("ground", ground)
	set("ink", lit(burst, inkContrast, ground))
	set("cursor", lit(burst, inkContrast, ground))
	set("cursor-ink", ground)
	set("surface-1", lit(ground, selectionContrast, ground))
	set("selection-ink", lit(burst, inkContrast, ground))
	set("black", lit(ground, bandContrast, ground))
	hues := []string{"red", "green", "yellow", "blue", "cyan"}
	for _, bright := range []bool{false, true} {
		ratio, prefix := float64(colourContrast), ""
		if bright {
			ratio, prefix = brightContrast, "bright-"
		}
		for _, name := range hues {
			wave := swatch.Monochrome(dominant(name))
			set(prefix+name, lit(wave, ratio, ground))
		}
		set(prefix+"magenta", purple(ratio, ground))
	}
	set("white", lit(burst, whiteContrast, ground))
	set("bright-white", lit(burst, brightWhite, ground))
	set("bright-black", lit(ground, dimContrast, ground))
	return colourway.Source{Name: "fermi-long", About: about, Roles: roles,
		Origin: "sources/fermi-long.css"}
}

const about = "" +
	`fermi-long -- a long gamma-ray burst seen through libtheme's fermi
profile, 10 kev to 300 gev carried onto visible. the ground is the deep
indigo at its violet end; the ink is the whole burst as one colour, lit to
11 to 1. each of the terminal's colours is the wavelength the screen's own
colour points at, lit to 6 to 1, or 8.5 for the brights. magenta is in no
spectrum, so it is the screen's red mixed with violet. made by spectra,
github.com/janearc/chroma/cmd/spectra, 2026-10-05.`

// indigo is the strip's deep end: of the colours an eye sees between 400
// and 455 nm, the one with the most chroma.
func indigo(strip []column) swatch.Swatch {
	best, most := swatch.Black, -1.0
	for _, each := range strip {
		if each.Nm < 400 || each.Nm > 455 {
			continue
		}
		value := srgb.MustHex(each.Seen).Swatch()
		if chroma := ok.FromSwatch(value).Polar().C; chroma > most {
			best, most = value, chroma
		}
	}
	return best
}

// dominant is the wavelength the screen's own colour of that name points
// at, from white: the strip's colour most cleanly made of the screen's
// red, green and blue.
func dominant(name string) int {
	nm, _ := visible.Of(srgb.MustHex(screens[name]).Swatch())
	return int(math.Round(nm))
}

// purple is magenta, which no single wavelength makes: the screen's red
// wavelength mixed with violet, in the share that lands nearest the
// screen's own magenta hue.
func purple(ratio float64, ground swatch.Swatch) swatch.Swatch {
	warm := lit(swatch.Monochrome(dominant("red")), ratio, ground)
	violet := swatch.Monochrome(420)
	target := ok.FromSwatch(srgb.MustHex("#ff00ff").Swatch()).Polar().H
	best, nearest := warm, math.Inf(1)
	for share := 0.01; share < 1; share += 0.01 {
		wx, wy, wz := warm.XYZ()
		mix := share / (1 - share) * wy / max(lumOf(violet), 1e-9)
		vx, vy, vz := scaled(violet, mix).XYZ()
		value := lit(swatch.FromXYZ(wx+vx, wy+vy, wz+vz), ratio, ground)
		hue := ok.FromSwatch(value).Polar().H
		if gap := math.Abs(hue - target); gap < nearest {
			best, nearest = value, gap
		}
	}
	return best
}

// lit is a colour brightened or dimmed, its chromaticity held, until it
// stands at a contrast against the ground, as the screen draws it: above
// a dark ground, below a light one; or as far as the screen goes, when
// that is not enough.
func lit(
	sample swatch.Swatch, ratio float64, ground swatch.Swatch,
) swatch.Swatch {
	lo, hi := ok.FromSwatch(ground).Polar().L, 0.995
	if lo > 0.6 {
		hi = 0
	}
	for step := 0; step < 48; step++ {
		middle := (lo + hi) / 2
		tried := drawn(atLightness(sample, middle))
		if contrast(tried, ground) < ratio {
			lo = middle
		} else {
			hi = middle
		}
	}
	return drawn(atLightness(sample, hi))
}

// drawn is a swatch as the screen draws it: fitted into srgb, to the byte.
func drawn(sample swatch.Swatch) swatch.Swatch {
	return srgb.MustHex(hex(sample)).Swatch()
}

// lumOf is a swatch's luminance, Y, which is one for white.
func lumOf(sample swatch.Swatch) float64 {
	_, y, _ := sample.XYZ()
	return y
}

// contrast is the WCAG ratio between two swatches as the screen draws them.
func contrast(a, b swatch.Swatch) float64 {
	first, second := lumOf(drawn(a)), lumOf(drawn(b))
	return (max(first, second) + 0.05) / (min(first, second) + 0.05)
}

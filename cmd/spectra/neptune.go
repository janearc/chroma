package main

import (
	"github.com/janearc/libtheme-css/colourway"
	"github.com/janearc/libtheme-css/css"
	"github.com/janearc/libtheme-css/primitives/swatch"
	"github.com/janearc/libtheme-css/spaces/ok"
	"github.com/janearc/libtheme-css/spaces/srgb"
)

// groundLightness is how light a dark ground is drawn, and groundSaturation
// the most saturated it may be when its hue is near the ink's: her hue rule
// asks hues 40 degrees apart, or a ground under 0.35 saturation.
const (
	groundLightness  = 0.2
	groundSaturation = 0.35
)

// regions is a sheet of sampled regions: each region's name and colour.
func regions(path string) (map[string]swatch.Swatch, error) {
	raw, err := readData(path)
	if err != nil {
		return nil, err
	}
	roles := css.Read(string(raw)).Roles
	found := map[string]swatch.Swatch{}
	for _, name := range roles.Names() {
		found[name], _ = roles.Get(name)
	}
	return found, nil
}

// neptuneVoyager is a colourway from voyager's neptune: the limb in shadow
// for the ground, the white clouds for the ink, the disc's blue for the
// selection and the blues, the dark spot for dim text.
//
// Neptune has no red, green or yellow, so those are the screen's own
// wavelengths, as in fermi-long.
func neptuneVoyager(path string) (colourway.Source, error) {
	found, err := regions(path)
	if err != nil {
		return colourway.Source{}, err
	}
	ground := shaded(found["limb"])
	roles := css.New()
	set := func(name string, value swatch.Swatch) {
		roles.Set(name, drawn(value))
	}
	set("ground", ground)
	for _, name := range []string{"ink", "cursor", "selection-ink"} {
		set(name, lit(found["clouds"], inkContrast, ground))
	}
	set("cursor-ink", ground)
	set("surface-1", lit(found["disc"], selectionContrast, ground))
	set("black", lit(ground, bandContrast, ground))
	set("bright-black", lit(found["dark-spot"], dimContrast, ground))
	setSixteen(set, ground, map[string]swatch.Swatch{
		"blue": found["disc"], "cyan": found["clouds"]})
	set("white", lit(found["clouds"], whiteContrast, ground))
	set("bright-white", lit(found["clouds"], brightWhite, ground))
	return colourway.Source{Name: "neptune-voyager", About: voyagerAbout,
		Roles: roles, Origin: "sources/neptune-voyager.css"}, nil
}

// setSixteen sets the terminal's colours and their brights: each from the
// planet where it has one, else the wavelength the screen's own colour
// points at, lit to 6 to 1, or 8.5 for the brights.
func setSixteen(
	set func(string, swatch.Swatch), ground swatch.Swatch,
	own map[string]swatch.Swatch,
) {
	hues := []string{"red", "green", "yellow", "blue", "cyan"}
	for _, bright := range []bool{false, true} {
		ratio, prefix := float64(colourContrast), ""
		if bright {
			ratio, prefix = brightContrast, "bright-"
		}
		for _, name := range hues {
			value, has := own[name]
			if !has {
				value = swatch.Monochrome(dominant(name))
			}
			set(prefix+name, lit(value, ratio, ground))
		}
		set(prefix+"magenta", purple(ratio, ground))
	}
}

// shaded is a colour as a dark ground: dimmed to the ground's lightness,
// then made greyer until it is no more saturated than the hue rule allows.
func shaded(sample swatch.Swatch) swatch.Swatch {
	polar := ok.FromSwatch(atLightness(sample, groundLightness)).Polar()
	for step := 0; step < 200; step++ {
		if saturation(polar) <= groundSaturation {
			break
		}
		polar.C *= 0.97
	}
	return drawn(polar.Rect().Swatch())
}

// saturation is a colour's saturation as hsl gives it, as drawn.
func saturation(polar ok.OKLCH) float64 {
	value, _ := srgb.FromSwatch(drawn(polar.Rect().Swatch()))
	return value.HSL().S
}

const voyagerAbout = "" +
	`neptune-voyager -- neptune as voyager 2 showed it in august 1989, from
the green and orange filters of its narrow angle camera (nasa/jpl,
pia01492). the ground is the limb in shadow, greyed to her hue rule; the
ink is the white clouds, lit to 11 to 1; the selection and the blues are
the disc; dim text is the great dark spot. neptune has no red, green or
yellow, so those are the screen's own wavelengths, as in fermi-long.
samples in cmd/spectra/data/neptune-voyager.css. made by spectra,
github.com/janearc/chroma/cmd/spectra, 2026-10-05.`

// pluto is a light colourway from new horizons' pluto of 2015, the
// picture everyone knows: the heart for the ground, the whale for the
// ink, darkened to 10 to 1, the pole for the selection, the midlands for
// claude's band, the disc for dim text.
//
// Pluto has no red, green or blue to speak of, so the sixteen are the
// screen's own wavelengths, darkened to stand on the heart.
func pluto(path string) (colourway.Source, error) {
	found, err := regions(path)
	if err != nil {
		return colourway.Source{}, err
	}
	ground := found["heart"]
	roles := css.New()
	set := func(name string, value swatch.Swatch) {
		roles.Set(name, drawn(value))
	}
	set("ground", ground)
	for _, name := range []string{"ink", "cursor", "selection-ink"} {
		set(name, lit(found["whale"], 10, ground))
	}
	set("cursor-ink", ground)
	set("surface-1", found["pole"])
	set("black", found["midlands"])
	set("bright-black", lit(found["disc"], dimContrast, ground))
	setSixteen(set, ground, map[string]swatch.Swatch{})
	set("white", lit(found["disc"], whiteContrast, ground))
	set("bright-white", lit(found["whale"], brightWhite, ground))
	return colourway.Source{Name: "pluto", About: plutoAbout, Roles: roles,
		Origin: "sources/pluto.css"}, nil
}

const plutoAbout = "" +
	`pluto -- a light colourway, pluto as new horizons first showed it on
13 july 2015, the day before closest approach (nasa/jhuapl/swri,
pia19708). the ground is the heart; the ink is the whale, darkened to 10
to 1; the selection is the north polar cap, claude's band the midlands,
dim text the disc. pluto has no red, green or blue to speak of, so the
sixteen are the screen's own wavelengths, darkened to stand on the heart.
samples in cmd/spectra/data/pluto-2015.css. made by spectra,
github.com/janearc/chroma/cmd/spectra, 2026-10-05.`

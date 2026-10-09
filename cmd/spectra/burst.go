package main

import (
	"math"

	"github.com/janearc/libtheme-css/primitives/bands"
	"github.com/janearc/libtheme-css/primitives/swatch"
	"github.com/janearc/libtheme-css/spaces/radiation"
)

// burst is a gamma-ray burst's spectrum as the Band function gives it: a
// slope below the peak, a slope above it, and the energy in keV where the
// power per logarithm of energy is greatest.
type burst struct {
	Name  string  `json:"name"`
	Alpha float64 `json:"alpha"`
	Beta  float64 `json:"beta"`
	Peak  float64 `json:"peak"`
}

// power is the burst's power per logarithm of energy at an energy in keV:
// the energy squared times the Band function, up to a constant. Below the
// joint it is a power law cut off by an exponential; above, a power law.
func (shape burst) power(kev float64) float64 {
	fold := shape.Peak / (2 + shape.Alpha)
	joint := (shape.Alpha - shape.Beta) * fold
	if kev < joint {
		return kev * kev * math.Pow(kev/100, shape.Alpha) *
			math.Exp(-kev/fold)
	}
	return kev * kev * math.Pow(joint/100, shape.Alpha-shape.Beta) *
		math.Exp(shape.Beta-shape.Alpha) * math.Pow(kev/100, shape.Beta)
}

// light is the burst as light in bands across a profile's source range,
// one line at the middle of each band. The bands are even in the logarithm
// of wavelength, so power per logarithm is power per band.
func (shape burst) light(profile radiation.Profile) bands.Light {
	return lightOf(profile, shape.inBand)
}

// inBand is the burst's power in one of a profile's bands, read at the
// band's middle; a burst has light everywhere.
func (shape burst) inBand(lo, hi float64) (float64, bool) {
	middle := radiation.Energy(math.Sqrt(lo*hi)) / radiation.KeV
	return shape.power(middle), true
}

// colour is the burst's light, every band of it, as one colour once the
// profile has carried it into visible.
func (shape burst) colour(profile radiation.Profile) swatch.Swatch {
	return profile.Light(shape.light(profile))
}

// column is one band of a light as a prism would spread it once the
// profile has carried it into visible.
type column struct {
	// Nm is where it lands in visible.
	Nm float64 `json:"nm"`
	// Kev is the energy it stands for.
	Kev float64 `json:"kev"`
	// Power is relative to the strongest band.
	Power float64 `json:"power"`
	// Seen is as the eye sees it, all bands exposed alike.
	Seen string `json:"seen"`
	// Hue is the wavelength's own colour, at one lightness.
	Hue string `json:"hue"`
}

// strip is the burst spread out by wavelength: a column for every band of
// the profile's source, shortest wavelength, highest energy, first.
func (shape burst) strip(profile radiation.Profile) []column {
	return spread(profile, shape.inBand)
}

// lightOf is a light in bands across a profile's source range, one line at
// the middle of each band, its power from power; a band with no light in
// it is left out.
func lightOf(
	profile radiation.Profile, power func(lo, hi float64) (float64, bool),
) bands.Light {
	lines := map[float64]float64{}
	first, last := bands.Of(profile.Source.Lo), bands.Of(profile.Source.Hi)
	for band := first; band <= last; band++ {
		lo, hi := band.Edges()
		if value, has := power(lo, hi); has {
			lines[math.Sqrt(lo*hi)] = value
		}
	}
	return bands.Lines(lines)
}

// spread is a light spread out by wavelength once a profile has carried it
// into visible: a column for every band of the profile's source, shortest
// wavelength first, its power from power. A band with no light in it has
// no colour, and is drawn empty.
func spread(
	profile radiation.Profile, power func(lo, hi float64) (float64, bool),
) []column {
	columns, seen, lit := []column{}, []swatch.Swatch{}, []bool{}
	strongest := 0.0
	first, last := bands.Of(profile.Source.Lo), bands.Of(profile.Source.Hi)
	for band := first; band <= last; band++ {
		lo, hi := band.Edges()
		middle := math.Sqrt(lo * hi)
		if middle < profile.Source.Lo || middle >= profile.Source.Hi {
			continue
		}
		value, has := power(lo, hi)
		strongest = max(strongest, value)
		columns = append(columns, column{
			Nm:    onto(profile, middle),
			Kev:   radiation.Energy(middle) / radiation.KeV,
			Power: value,
			Hue:   hex(atLightness(profile.Swatch(middle), 0.7)),
		})
		seen = append(seen, scaled(profile.Swatch(middle), value))
		lit = append(lit, has)
	}
	exposure := exposed(seen, 0.4)
	for index := range columns {
		columns[index].Power /= strongest
		if lit[index] {
			columns[index].Seen = hex(scaled(seen[index], exposure))
		}
	}
	return columns
}

// onto is where a source wavelength lands in a profile's render range,
// evenly in the logarithm of wavelength, as the profile itself maps it.
func onto(profile radiation.Profile, nm float64) float64 {
	source, render := profile.Source, profile.Render
	t := math.Log(nm/source.Lo) / math.Log(source.Hi/source.Lo)
	return render.Lo * math.Pow(render.Hi/render.Lo, t)
}

// softening is a burst's colour as its peak falls, the way a burst's
// spectrum softens over its seconds: from the burst's own peak down to a
// tenth of it, each at one lightness so only hue and chroma change.
func (shape burst) softening(profile radiation.Profile, steps int) []string {
	ramp := []string{}
	for step := 0; step < steps; step++ {
		later := shape
		fall := float64(step) / float64(steps-1)
		later.Peak = shape.Peak * math.Pow(0.1, fall)
		chip := atLightness(later.colour(profile), 0.72)
		ramp = append(ramp, hex(chip))
	}
	return ramp
}

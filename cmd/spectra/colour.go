package main

import (
	"math"

	"github.com/janearc/libtheme-css/primitives/swatch"
	"github.com/janearc/libtheme-css/spaces/ok"
	"github.com/janearc/libtheme-css/spaces/srgb"
)

// hex is a swatch as a screen draws it, pulled into srgb by chroma when it
// falls outside, holding lightness and hue, as paratune's fitted does.
func hex(sample swatch.Swatch) string {
	if !srgb.In(sample) {
		polar, _ := ok.Fit(ok.FromSwatch(sample).Polar(), srgb.In)
		sample = polar.Rect().Swatch()
	}
	colour, _ := srgb.FromSwatch(sample)
	return colour.Hex()
}

// scaled is a swatch made k times as strong: the same colour, more light.
func scaled(sample swatch.Swatch, k float64) swatch.Swatch {
	x, y, z := sample.XYZ()
	return swatch.FromXYZ(x*k, y*k, z*k)
}

// atLightness is a swatch brightened or dimmed to an oklab lightness. Oklab
// takes a cube root of the light, so scaling the light by the cube of the
// ratio moves lightness exactly and leaves the colour's chromaticity alone.
func atLightness(sample swatch.Swatch, lightness float64) swatch.Swatch {
	current := ok.FromSwatch(sample).Polar().L
	if current <= 0 {
		return sample
	}
	return scaled(sample, math.Pow(lightness/current, 3))
}

// exposed is the one factor that brings the brightest of the swatches to a
// luminance, as a camera's exposure does for a whole frame.
func exposed(samples []swatch.Swatch, luminance float64) float64 {
	brightest := 0.0
	for _, sample := range samples {
		_, y, _ := sample.XYZ()
		brightest = max(brightest, y)
	}
	if brightest == 0 {
		return 0
	}
	return luminance / brightest
}

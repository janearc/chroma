package main

import (
	"testing"

	"github.com/janearc/libtheme-css/spaces/ok"
)

// = and - move every colour's lightness, + and _ every colour's chroma,
// and each marks the colourway unsaved.
func TestWhole(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	polar := func(name string) ok.OKLCH {
		return ok.FromSwatch(swatchOf(screen.set.shown(name))).Polar()
	}
	ground, contrastBefore := polar("ground").L, contrast(screen.set.shown("ink"), screen.set.shown("ground"))
	red := polar("red").C
	screen.press("-")
	if polar("ground").L >= ground {
		t.Error("- did not dim the ground")
	}
	if got := contrast(screen.set.shown("ink"), screen.set.shown("ground")); got <= contrastBefore {
		t.Errorf("- took the ink's contrast from %.2f to %.2f", contrastBefore, got)
	}
	if polar("red").C <= red {
		t.Error("- did not raise chroma")
	}
	if !screen.set.dirty {
		t.Error("- did not mark the colourway unsaved")
	}
	red = polar("red").C
	screen.press("+")
	if polar("red").C <= red {
		t.Errorf("+ left red's chroma at %.3f from %.3f", polar("red").C, red)
	}
	screen.press("_")
	screen.press("_")
	if polar("red").C >= red {
		t.Errorf("_ twice left red's chroma at %.3f", polar("red").C)
	}
}

// > lightens every text colour the colourway sets by one step and leaves
// every ground; < takes them back.
func TestTextLightness(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	ground, heading := screen.set.shown("ground"), polarOf(screen.set.shown("md-heading"))
	screen.press(">")
	if screen.set.shown("ground") != ground {
		t.Error("> moved the ground")
	}
	if moved := polarOf(screen.set.shown("md-heading")).L - heading.L; moved < 0.015 || moved > 0.025 {
		t.Errorf("> moved md-heading's lightness %.3f", moved)
	}
	screen.press("<")
	if back := polarOf(screen.set.shown("md-heading")).L - heading.L; back > 0.005 || back < -0.005 {
		t.Errorf("> then < left md-heading %.3f from where it was", back)
	}
}

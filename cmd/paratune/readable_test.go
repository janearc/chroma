package main

import (
	"math"
	"strings"
	"testing"
)

// r puts the page's ink in the prose band and a dim colour over the floor
// for tokens, each keeping its hue, and leaves a colour that reads as it
// is; the grounds are not touched.
func TestReadable(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	steer(t, screen, shellPage, "ink")
	set := screen.set
	ground := colour{16, 16, 16}
	set.roles.put("ground", ground)
	set.roles.put("ink", colour{64, 64, 64})
	set.roles.put("red", colour{64, 16, 16})
	green := colour{64, 192, 64}
	set.roles.put("green", green)
	screen.press("r")
	if ratio := contrast(set.shown("ink"), ground); !proseBand().Contains(ratio) {
		t.Errorf("the ink is %.1f:1 after r", ratio)
	}
	red := set.shown("red")
	if ratio := contrast(red, ground); ratio < reading.Contrast.ChipMin {
		t.Errorf("red is %.1f:1 after r", ratio)
	}
	if hue := polarOf(red).H; math.Abs(hue-polarOf(colour{64, 16, 16}).H) > 15 {
		t.Errorf("red's hue went to %.0f", hue)
	}
	if set.shown("green") != green || set.shown("ground") != ground {
		t.Error("r moved a colour that reads, or a ground")
	}
	if !strings.HasPrefix(screen.note, "readable: ") || !set.dirty {
		t.Errorf("r said %q, unsaved %v", screen.note, set.dirty)
	}
}

// a page where every colour already reads is left alone, and r says so.
func TestReadableLeavesAReadablePage(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.press("r")
	screen.set.dirty = false
	before := screen.set.roles.sheet().String()
	screen.press("r")
	if screen.set.roles.sheet().String() != before || screen.set.dirty {
		t.Error("a second r moved something")
	}
	if !strings.HasPrefix(screen.note, "readable: 0 moved") {
		t.Errorf("the second r said %q", screen.note)
	}
}

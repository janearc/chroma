package main

import (
	"fmt"
	"strings"
	"testing"
)

// c copies a field's foreground and v pastes it into another's, which
// makes it the source's own and the colourway unsaved.
func TestCopyPasteForeground(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	heading := screen.set.shown("md-heading")
	steer(t, screen, markdownPage, "md-heading")
	screen.press("c")
	steer(t, screen, markdownPage, "md-link")
	screen.press("v")
	if got, own := screen.set.roles.get("md-link"); !own || got != heading {
		t.Errorf("md-link is %s, set %v; want %s", got.hex(), own, heading.hex())
	}
	if !screen.set.dirty {
		t.Error("a paste did not mark the colourway unsaved")
	}
}

// d copies a field's ground and f pastes it: a heading's ground into
// code's ground, since each is the other's field.
func TestCopyPasteGround(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.set.roles.put("md-heading-ground", colour{11, 22, 33})
	steer(t, screen, markdownPage, "md-heading")
	screen.press("d")
	steer(t, screen, markdownPage, "md-code")
	screen.press("f")
	if got, _ := screen.set.roles.get("md-code-ground"); got != (colour{11, 22, 33}) {
		t.Errorf("md-code-ground is %s", got.hex())
	}
}

// text with no ground of its own has none to copy or paste into, and a
// register with nothing in it pastes nothing; each says so.
func TestNothingToCopyOrPaste(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	steer(t, screen, markdownPage, "md-quote")
	screen.press("d")
	if screen.note != "no ground here" {
		t.Errorf("d said %q", screen.note)
	}
	screen.press("f")
	if screen.note != "no ground copied" {
		t.Errorf("f said %q", screen.note)
	}
	screen.press("v")
	if screen.note != "no foreground copied" {
		t.Errorf("v said %q", screen.note)
	}
	if screen.set.dirty {
		t.Error("nothing pasted, and the colourway is unsaved")
	}
}

// on vim's own ground, the foreground register goes to vim's ink: a
// copied ink can never become a ground.
func TestVimGroundTakesNoInk(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.set.roles.put("nvim-ground", colour{200, 210, 190})
	screen.stops = stopsOf(screen.set.roles)
	steer(t, screen, shellPage, "ink")
	screen.press("c")
	steer(t, screen, vimPage, "nvim-ground")
	screen.press("v")
	if got, _ := screen.set.roles.get("nvim-ground"); got != (colour{200, 210, 190}) {
		t.Errorf("vim's ground became %s", got.hex())
	}
	if !strings.HasPrefix(screen.note, "pasted foreground") {
		t.Errorf("v said %q", screen.note)
	}
}

// y and p copy and paste both, in one key each.
func TestBoth(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.set.roles.put("md-heading-ground", colour{11, 22, 33})
	heading := screen.set.shown("md-heading")
	steer(t, screen, markdownPage, "md-heading")
	screen.press("y")
	steer(t, screen, markdownPage, "md-small-heading")
	screen.press("p")
	if got := screen.set.shown("md-small-heading"); got != heading {
		t.Errorf("the small heading is %s", got.hex())
	}
	if got := screen.set.shown("md-small-heading-ground"); got != (colour{11, 22, 33}) {
		t.Errorf("its ground is %s", got.hex())
	}
}

// the field steered is written in its own colour: on the ground when the
// reader can read it there, and on libreadme's chip fill when it cannot,
// with the ground put back after.
func TestFieldInItsColour(t *testing.T) {
	ground := colour{128, 138, 99}
	legible := colour{20, 20, 20}
	got := field("ink", legible, ground, "")
	if !strings.HasPrefix(got, inkCode(legible)+"ink: #141414") || strings.Contains(got, "\x1b[48;2;") {
		t.Errorf("a readable field is drawn %q", got)
	}
	faint := colour{120, 130, 95}
	got = field("dim", faint, ground, "")
	codes := escapes.FindAllString(got, -1)
	var fill colour
	fmt.Sscanf(codes[0], "\x1b[48;2;%f;%f;%fm", &fill.r, &fill.g, &fill.b)
	if ratio := contrast(faint, fill); ratio < reading.Contrast.ChipMin {
		t.Errorf("the label is %.1f:1 on its fill %s", ratio, fill.hex())
	}
	if !strings.HasPrefix(got, "\x1b[48;2;") || strings.HasPrefix(got, groundCode(ground)) ||
		!strings.Contains(got, inkCode(faint)+"dim: "+faint.hex()) ||
		!strings.HasSuffix(got, groundCode(ground)) {
		t.Errorf("an unreadable field is drawn %q", got)
	}
}

// n makes the colour steered the next of the colourway's own colours, and
// shift-n the one before; going all the way round comes back to where it
// started, and each step is the source's own.
func TestPaletteKey(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	steer(t, screen, markdownPage, "md-heading")
	colours := palette(screen.set)
	start := screen.set.shown("md-heading").hex()
	screen.press("n")
	first := screen.set.shown("md-heading").hex()
	if first == start || !screen.set.dirty {
		t.Fatalf("n left md-heading at %s", first)
	}
	screen.press("N")
	if got := screen.set.shown("md-heading").hex(); got != start {
		t.Errorf("shift-n went to %s, not back to %s", got, start)
	}
	for range colours {
		screen.press("n")
	}
	if got := screen.set.shown("md-heading").hex(); got != start {
		t.Errorf("all the way round ended at %s, not %s", got, start)
	}
	if !strings.Contains(screen.note, "of "+fmt.Sprint(len(colours))) {
		t.Errorf("n said %q", screen.note)
	}
}

// a colour only the steered role holds is not lost to n: shift-n brings
// it back, and so does going all the way round.
func TestPaletteKeepsAUniqueColour(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	steer(t, screen, markdownPage, "md-link")
	unique := colour{0x12, 0x34, 0x56}
	screen.set.roles.put("md-link", unique)
	screen.press("n")
	if screen.set.shown("md-link").hex() == unique.hex() {
		t.Fatal("n did not move md-link")
	}
	screen.press("N")
	if got := screen.set.shown("md-link").hex(); got != unique.hex() {
		t.Errorf("shift-n went to %s, not back to %s", got, unique.hex())
	}
	for range len(screen.walking.colours) {
		screen.press("n")
	}
	if got := screen.set.shown("md-link").hex(); got != unique.hex() {
		t.Errorf("all the way round ended at %s, not %s", got, unique.hex())
	}
}

// paratune's own marks read on any ground, whatever the page's ink has
// been tuned to: an ink the ground swallows, on a dark ground and on a
// mid-tone one; an ink inside the band is kept, and one past the ceiling,
// which glares, is brought under it.
func TestControlInkReads(t *testing.T) {
	floor := reading.Contrast.ChipMin
	for _, pair := range [][2]colour{
		{{9, 6, 2}, {18, 8, 4}},
		{{130, 135, 120}, {128, 138, 99}},
		{{250, 250, 250}, {255, 255, 255}},
	} {
		if ratio := contrast(controlInk(pair[0], pair[1]), pair[1]); ratio < floor {
			t.Errorf("on %s the marks are %.1f:1", pair[1].hex(), ratio)
		}
	}
	reads, dark := colour{170, 160, 140}, colour{16, 16, 16}
	if got := controlInk(reads, dark); got != reads {
		t.Errorf("an ink that reads became %s", got.hex())
	}
	if ratio := contrast(controlInk(colour{250, 250, 250}, dark), dark); ratio > reading.Contrast.High {
		t.Errorf("an ink past the ceiling stayed at %.1f:1", ratio)
	}
}

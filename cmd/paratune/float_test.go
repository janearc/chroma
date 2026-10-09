package main

import (
	"strings"
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
)

// plainRows is the float's rows without their colours.
func plainRows(screen *model) []string {
	rows := []string{}
	for _, text := range screen.float("", colour{128, 138, 99}) {
		rows = append(rows, escapes.ReplaceAllString(text, ""))
	}
	return rows
}

// at rest the float names the colourway and the page, then the field in
// its own colour with its ratio, and ends with ? help.
func TestFloatAtRest(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	for which, name := range pageNames {
		screen.page = page(which)
		screen.steering = screen.first(screen.page)
		rows := plainRows(screen)
		here := screen.stops[screen.steering]
		steered := false
		for _, row := range rows[2:] {
			steered = steered || strings.Contains(row, here.label+": ") && strings.Contains(row, ":1")
		}
		if rows[0] != "test-dusk" || rows[1] != name || !steered || rows[len(rows)-1] != "? help" {
			t.Errorf("on %s the float is %q", name, rows)
		}
	}
}

// a key makes the float the field's row alone; a minute without one
// brings back all of it, and a look before the minute is up does not.
func TestFloatWhilePressing(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	if _, command := screen.Update(tea.KeyPressMsg{Code: tea.KeyUp}); command == nil {
		t.Fatal("a key scheduled no look at the quiet")
	}
	if rows := plainRows(screen); len(rows) != 1 {
		t.Errorf("while pressing the float is %q", rows)
	}
	screen.Update(quietMsg{})
	if !screen.busy {
		t.Error("a look a moment after the key brought the float back")
	}
	screen.lastStirred = time.Now().Add(-2 * quietAfter)
	screen.Update(quietMsg{})
	if rows := plainRows(screen); screen.busy || rows[len(rows)-1] != "? help" {
		t.Errorf("after a quiet minute the float is %q", rows)
	}
}

// the float sits at the right of the top rows, every row still the
// pane's width; a click on ? help opens the keys, and a click elsewhere
// on it steers nothing beneath it.
func TestFloatOnThePage(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	content := screen.View().Content
	rows := strings.Split(content, "\n")
	if !strings.Contains(escapes.ReplaceAllString(rows[0], ""), "test-dusk") {
		t.Errorf("the top row is %q", escapes.ReplaceAllString(rows[0], ""))
	}
	placed, steering := screen.placed, screen.steering
	if placed.help < 0 || placed.left+placed.width >= screen.width-gradientWidth+2 {
		t.Fatalf("the float is at %+v", placed)
	}
	screen.Update(tea.MouseClickMsg{X: placed.left + 1, Y: 0, Button: tea.MouseLeft})
	if screen.keys || screen.steering != steering {
		t.Error("a click on the float's name did something")
	}
	screen.Update(tea.MouseClickMsg{X: placed.left + 1, Y: placed.help, Button: tea.MouseLeft})
	if !screen.keys {
		t.Error("a click on ? help did not open the keys")
	}
}

// a row with a wide glyph cut at the float's edge keeps the row's width.
func TestOverlayKeepsWidth(t *testing.T) {
	ground, page := colour{40, 40, 60}, colour{20, 20, 30}
	for shift := 0; shift < 3; shift++ {
		rows := []string{strings.Repeat("x", 30+shift) + strings.Repeat("漢", 20)}
		placed := overlay(rows, []string{"name", "? help"}, 70, ground, page)
		if width := ansi.StringWidth(rows[0]); width != 70 {
			t.Errorf("shifted %d, the float at %d: the row is %d wide", shift,
				placed.left, width)
		}
	}
}

// the float says when the colourway sets few of the page's colours, so a
// page drawn mostly in derived colours is not taken for its own.
func TestFloatSaysFew(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "plain", home)
	screen.page = vimPage
	screen.steering = screen.first(screen.page)
	if rows := plainRows(screen); !strings.HasPrefix(rows[2], "sets ") {
		t.Errorf("the float is %q", rows)
	}
}

// a walk through the palette ends when another colourway is opened.
func TestWalkEndsWithItsColourway(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.press("n")
	if screen.walking == nil {
		t.Fatal("n began no walk")
	}
	screen.open("plain")
	if screen.walking != nil {
		t.Error("the walk went on into another colourway")
	}
}

// the float's marks read on the float's own ground on every page ground,
// the mid-tones included, where the marks can only just clear the floor
// against the page and the float must not take that away.
func TestFloatMarksRead(t *testing.T) {
	for level := 0; level <= 255; level += 5 {
		for _, ground := range []colour{{float64(level), float64(level), float64(level)},
			{164, 106, 30}, {float64(level), float64(level) * 0.65, float64(level) * 0.2}} {
			for _, ink := range []colour{{9, 6, 2}, {130, 135, 120}, {250, 250, 250}} {
				marks := controlInk(ink, ground)
				if ratio := contrast(marks, floatGround(ground, marks)); ratio < reading.Contrast.ChipMin {
					t.Fatalf("on %s, ink %s: the float's marks are %.2f:1", ground.hex(), ink.hex(), ratio)
				}
			}
		}
	}
}

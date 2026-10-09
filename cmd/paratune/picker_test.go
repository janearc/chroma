package main

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
)

// o opens the picker and o or escape closes it, escape without leaving;
// with it open the frame is still the pane, every row its width.
func TestPickerOpensAndCloses(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	for _, step := range []struct {
		key  string
		open bool
	}{{"o", true}, {"esc", false}, {"o", true}, {"o", false}} {
		if command := screen.press(step.key); command != nil {
			t.Fatalf("%s gave a command: paratune would have left", step.key)
		}
		if screen.picking != step.open {
			t.Errorf("after %s the picker is open %v", step.key, screen.picking)
		}
	}
	screen.press("o")
	for size := range 2 {
		screen.width, screen.height = 88+10*size, 26+8*size
		rows := strings.Split(screen.View().Content, "\n")
		if len(rows) != screen.height {
			t.Fatalf("%d rows with the picker open", len(rows))
		}
		for number, row := range rows {
			if width := ansi.StringWidth(row); width != screen.width {
				t.Fatalf("row %d is %d wide with the picker open", number, width)
			}
		}
	}
}

// a click on a cell of the picker takes the colour drawn there: a colour
// of the colourway's own, then one from the grid.
func TestPickerClick(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.press("o")
	screen.View()
	name, top, cells := screen.stops[screen.steering].name, screen.height-pickerRows, (screen.width-6)/2
	want, _ := screen.pickerColour(1, 2, cells)
	screen.Update(tea.MouseClickMsg{X: 5, Y: top + 1, Button: tea.MouseLeft})
	if got := screen.set.shown(name); got != want || !screen.set.dirty {
		t.Errorf("a palette cell gave %s, not %s", got.hex(), want.hex())
	}
	want, _ = screen.pickerColour(4, 10, cells)
	screen.Update(tea.MouseClickMsg{X: 20, Y: top + 4, Button: tea.MouseLeft})
	if got := screen.set.shown(name); got.hex() != want.hex() {
		t.Errorf("a grid cell gave %s, not %s", got.hex(), want.hex())
	}
}

// a cell is marked when its colour would not read against the field's
// ground, and only then.
func TestPickerMarks(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	here := screen.stops[screen.steering]
	against := screen.set.against(here)
	rows := screen.picker(82, "")
	cells := 82 / 2
	for row := 1; row < pickerRows; row++ {
		text := escapes.ReplaceAllString(rows[row], "")
		for cell := range cells {
			value, has := screen.pickerColour(row, cell, cells)
			if !has {
				continue
			}
			marked := []rune(text)[2*cell] == '·'
			if marked != (contrast(value, against) < reading.Contrast.ChipMin) {
				t.Fatalf("row %d cell %d: marked %v at %.1f:1", row, cell, marked,
					contrast(value, against))
			}
		}
	}
}

// a click on the field in the float opens the picker.
func TestFloatFieldOpensThePicker(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.View()
	if screen.placed.field < 0 {
		t.Fatal("the float has no field row")
	}
	screen.Update(tea.MouseClickMsg{X: screen.placed.left + 2, Y: screen.placed.field,
		Button: tea.MouseLeft})
	if !screen.picking {
		t.Error("a click on the float's field did not open the picker")
	}
}

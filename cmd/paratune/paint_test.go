package main

import (
	"slices"
	"testing"

	tea "charm.land/bubbletea/v2"
)

// a held key does not paint on every press: each press moves the colour
// at once, one paint is scheduled, and the pane is painted when it comes,
// once, however many presses came before it.
func TestPaintingKeepsUp(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	paints := 0
	screen.painter = func(*tuning) { paints++ }
	before := screen.set.shown("ink")
	scheduled := 0
	for range 10 {
		if _, command := screen.Update(tea.KeyPressMsg{Code: tea.KeyUp}); command != nil {
			scheduled++
		}
	}
	if screen.set.shown("ink") == before {
		t.Fatal("ten presses did not move the ink")
	}
	if paints != 0 || scheduled != 1 {
		t.Errorf("ten presses painted %d times and scheduled %d paints", paints, scheduled)
	}
	screen.Update(paintMsg{})
	screen.Update(paintMsg{})
	if paints != 1 {
		t.Errorf("the paints that came painted %d times, not once", paints)
	}
}

// a paint is one call to tmux: every option set in it, joined by tmux's
// separator, the pane named in each.
func TestPaintIsOneCall(t *testing.T) {
	dir, home := fixture(t)
	set := opened(t, dir, "test-dusk", home).set
	args := paintArgs("%9", set)
	commands := [][]string{{}}
	for _, arg := range args {
		if arg == ";" {
			commands = append(commands, []string{})
			continue
		}
		commands[len(commands)-1] = append(commands[len(commands)-1], arg)
	}
	if want := 2 + len(paintedNumbers()); len(commands) != want {
		t.Fatalf("%d commands in the call, want %d", len(commands), want)
	}
	for _, command := range commands {
		if len(command) != 6 || !slices.Equal(command[:4], []string{"set", "-p", "-t", "%9"}) {
			t.Errorf("a command is %q", command)
		}
	}
}

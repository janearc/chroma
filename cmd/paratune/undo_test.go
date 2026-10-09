package main

import (
	"maps"
	"testing"

	tea "charm.land/bubbletea/v2"
)

// TestUndoComesBackFromWhite lifts every text colour until the ink is
// white, where its hue is gone, then undoes every step, and finds every
// colour as it opened.
func TestUndoComesBackFromWhite(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	start := maps.Clone(screen.set.roles.values)
	for range 80 {
		screen.Update(key(">", '>'))
	}
	if got := screen.set.shown("ink").hex(); got != "#ffffff" {
		t.Fatalf("eighty lifts left the ink at %s, not white", got)
	}
	if len(screen.set.history) == 0 {
		t.Fatal("eighty lifts kept no steps to undo")
	}
	for range 200 {
		if len(screen.set.history) == 0 {
			break
		}
		screen.Update(key("u", 'u'))
	}
	if count := len(screen.set.history); count != 0 {
		t.Fatalf("two hundred undos left %d steps; an undo is being kept", count)
	}
	if !maps.Equal(screen.set.roles.values, start) {
		t.Error("undoing every step did not bring every colour back")
	}
	if screen.note != "undone: back to how it opened" {
		t.Errorf("the last undo said %q", screen.note)
	}
	screen.Update(key("u", 'u'))
	if screen.note != "nothing to undo: this is how it opened" {
		t.Errorf("an undo with nothing left said %q", screen.note)
	}
}

// TestUndoKeepsOnlySteps moves around the page without changing a
// colour, which keeps nothing, and one change, which keeps one step.
func TestUndoKeepsOnlySteps(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	screen.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	if count := len(screen.set.history); count != 0 {
		t.Errorf("moving between colours kept %d steps", count)
	}
	screen.Update(tea.KeyPressMsg{Code: tea.KeyUp})
	if count := len(screen.set.history); count != 1 {
		t.Errorf("one change kept %d steps", count)
	}
	screen.Update(key("u", 'u'))
	if count := len(screen.set.history); count != 0 {
		t.Errorf("an undo left %d steps, and kept itself as one", count)
	}
}

// TestUndoToTheSavedColoursIsSaved finds a colourway undone back to what
// was last saved counted as saved, and one undone short of it not.
func TestUndoToTheSavedColoursIsSaved(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.Update(tea.KeyPressMsg{Code: tea.KeyUp})
	screen.set.save()
	screen.Update(tea.KeyPressMsg{Code: tea.KeyUp})
	screen.Update(tea.KeyPressMsg{Code: tea.KeyUp})
	screen.Update(key("u", 'u'))
	if !screen.set.dirty {
		t.Error("undone one step of two past the save, and counted saved")
	}
	screen.Update(key("u", 'u'))
	if screen.set.dirty {
		t.Error("undone back to the save, and counted unsaved")
	}
}

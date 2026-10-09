package main

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/colorprofile"
)

// TestOtherDisplayGoesRound steps ours, theirs at 256, theirs at 16, and
// back to ours, saying each step in the note.
func TestOtherDisplayGoesRound(t *testing.T) {
	screen := &model{}
	want := []colorprofile.Profile{colorprofile.ANSI256, colorprofile.ANSI, colorprofile.TrueColor}
	for _, profile := range want {
		if screen.otherDisplay() == nil {
			t.Fatal("no command to change the display")
		}
		if displays[screen.display].profile != profile {
			t.Errorf("display %d is %v, want %v", screen.display, displays[screen.display].profile, profile)
		}
		if screen.note != displays[screen.display].note {
			t.Errorf("note %q", screen.note)
		}
	}
}

// TestDrawWithSetsTheProfile is the message bubbletea reads to change the
// renderer's colour profile.
func TestDrawWithSetsTheProfile(t *testing.T) {
	msg, ok := drawWith(colorprofile.ANSI256)().(tea.ColorProfileMsg)
	if !ok || msg.Profile != colorprofile.ANSI256 {
		t.Errorf("got %#v", msg)
	}
}

// TestBackslashChangesTheDisplay is the key reaching otherDisplay.
func TestBackslashChangesTheDisplay(t *testing.T) {
	screen := &model{}
	if cmd := screen.press(`\`); cmd == nil || screen.display != 1 {
		t.Errorf("display %d after backslash", screen.display)
	}
}

// TestFloatSaysTheirs names charm's display beside the page while it is
// on, and nothing there on ours.
func TestFloatSaysTheirs(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	for which, display := range displays {
		screen.display = which
		want := pageNames[screen.page]
		if display.name != "" {
			want += ", " + display.name
		}
		if rows := plainRows(screen); rows[1] != want {
			t.Errorf("display %d: the float says %q, want %q", which, rows[1], want)
		}
	}
}

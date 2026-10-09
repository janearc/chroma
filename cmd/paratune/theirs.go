package main

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/colorprofile"
)

// displays are what backslash steps through: paratune's own true colour,
// then the colourway as charm's colorprofile draws it for a 256 colour
// terminal and for a 16 colour one.
//
// Theirs is bubbletea's renderer told the terminal is smaller, so every
// colour on the page goes through colorprofile's own conversion; the
// pane's palette and ground are the colourway's still, as a terminal with
// the colourway installed would be.
var displays = []struct {
	profile colorprofile.Profile
	name    string // for the float while it is not ours
	note    string // said once, on the step
}{
	{colorprofile.TrueColor, "", "ours: true colour"},
	{colorprofile.ANSI256, "theirs 256",
		"theirs: 256 colours, through charm's colorprofile"},
	{colorprofile.ANSI, "theirs 16",
		"theirs: 16 colours, through charm's colorprofile"},
}

// otherDisplay steps to the next display and has bubbletea draw with it.
// The screen is cleared after, because the renderer redraws only cells
// that change and a cell left alone would keep the last display's colour.
func (screen *model) otherDisplay() tea.Cmd {
	screen.display = (screen.display + 1) % len(displays)
	next := displays[screen.display]
	screen.note = next.note
	return tea.Sequence(drawWith(next.profile), tea.ClearScreen)
}

// drawWith is the message that sets the renderer's colour profile:
// bubbletea applies it on the way to Update.
func drawWith(profile colorprofile.Profile) tea.Cmd {
	return func() tea.Msg { return tea.ColorProfileMsg{Profile: profile} }
}

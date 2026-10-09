package main

import (
	"github.com/janearc/libtheme-css/dialects/ghostty"
	"strings"
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
)

// the arrows move the steered role, make it the source's own, mark the
// colourway unsaved, and never leave a lamp outside nought to 255.
func TestArrows(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	steer(t, screen, shellPage, "ink")
	before := screen.set.shown("ink")
	screen.press("up")
	after := screen.set.shown("ink")
	if after.r <= before.r || !screen.set.dirty {
		t.Errorf("up took %s to %s, dirty %v", before.hex(), after.hex(), screen.set.dirty)
	}
	for range 200 {
		screen.press("up")
	}
	if top := screen.set.shown("ink"); top.r > 255 || top.g > 255 || top.b > 255 {
		t.Errorf("lighter went past white: %v", top)
	}
	for range 400 {
		screen.press("down")
	}
	if bottom := screen.set.shown("ink"); bottom.r < 0 || bottom.g < 0 || bottom.b < 0 {
		t.Errorf("darker went past black: %v", bottom)
	}
	screen.press("right")
	screen.press("left")
}

// a move on an unset role makes it the source's own.
func TestMovingAnUnsetRoleSetsIt(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	steer(t, screen, markdownPage, "md-heading")
	if _, own := screen.set.roles.get("md-heading"); own {
		t.Fatal("md-heading was set to begin with")
	}
	screen.press("up")
	if _, own := screen.set.roles.get("md-heading"); !own {
		t.Error("moving md-heading did not set it")
	}
}

// tab walks the page's stops and nothing else, round to the start.
func TestTabStaysOnThePage(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.page, screen.steering = claudePage, screen.first(claudePage)
	start, visited := screen.steering, 0
	for {
		screen.press("tab")
		visited++
		if screen.stops[screen.steering].page != claudePage {
			t.Fatal("tab left the page")
		}
		if screen.steering == start || visited > len(screen.stops) {
			break
		}
	}
	screen.press("shift+tab")
	screen.press("tab")
	if screen.steering != start {
		t.Error("shift-tab and tab are not opposites")
	}
}

// a page opens on its ink.
func TestFirstIsTheInk(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	for which, ink := range map[page]string{shellPage: "ink",
		claudePage: "claude-text", markdownPage: "md-body", vimPage: "ink"} {
		if got := screen.stops[screen.first(which)].name; got != ink {
			t.Errorf("page %d opens on %s, not %s", which, got, ink)
		}
	}
}

// key is a key press as bubbletea sends it.
func key(text string, code rune) tea.KeyPressMsg {
	return tea.KeyPressMsg{Text: text, Code: code}
}

// [ opens the menu and ] closes it unchanged, as does escape.
func TestBracketsOpenAndClose(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	for _, closer := range []tea.KeyPressMsg{key("]", ']'), {Code: tea.KeyEscape}} {
		screen.Update(key("[", '['))
		if screen.menu == nil {
			t.Fatal("[ did not open the menu")
		}
		screen.Update(closer)
		if screen.menu != nil {
			t.Errorf("%s did not close the menu", closer.String())
		}
	}
}

// the menu goes to the colourway and page chosen, steering its ink.
func TestMenuGoesToThePage(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.openMenu()
	screen.choice = choice{Name: "plain", Page: markdownPage}
	screen.menu.State = huh.StateCompleted
	screen.steerMenu(key("x", 'x'))
	if screen.set.source.Name != "plain" || screen.page != markdownPage ||
		screen.stops[screen.steering].name != "md-body" {
		t.Errorf("went to %s page %d steering %s", screen.set.source.Name,
			screen.page, screen.stops[screen.steering].name)
	}
}

// a colourway with unsaved moves can be left, and has them again, still
// unsaved, when it is turned back to.
func TestUnsavedIsKept(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.press("up")
	moved := screen.set.shown("ink")
	screen.open("plain")
	if screen.set.source.Name != "plain" || screen.set.dirty {
		t.Fatalf("on %s, dirty %v", screen.set.source.Name, screen.set.dirty)
	}
	screen.open("test-dusk")
	if !screen.set.dirty || screen.set.shown("ink") != moved {
		t.Errorf("test-dusk came back dirty %v with ink %s, not %s",
			screen.set.dirty, screen.set.shown("ink").hex(), moved.hex())
	}
}

// drive hands paratune a message as bubbletea would, then every message
// the commands it returns give back, until there are none. A command
// that takes longer than a moment, a cursor's blink, is let go.
func drive(screen *model, message tea.Msg) {
	queue := []tea.Msg{message}
	for steps := 0; len(queue) > 0 && steps < 200; steps++ {
		next := queue[0]
		queue = queue[1:]
		_, command := screen.Update(next)
		queue = append(queue, run(command)...)
	}
}

// run is what a command gives back, a batch opened out.
func run(command tea.Cmd) []tea.Msg {
	if command == nil {
		return nil
	}
	given := make(chan tea.Msg, 1)
	go func() { given <- command() }()
	select {
	case message := <-given:
		if batch, ok := message.(tea.BatchMsg); ok {
			messages := []tea.Msg{}
			for _, each := range batch {
				messages = append(messages, run(each)...)
			}
			return messages
		}
		if message == nil {
			return nil
		}
		return []tea.Msg{message}
	case <-time.After(100 * time.Millisecond):
		return nil
	}
}

// her keys, exactly: [ opens the menu; up to another colourway and enter;
// down to a page and enter. paratune is then on that colourway, drawn in
// it, on that page, with unsaved moves on the one it left or without.
func TestMenuByKeys(t *testing.T) {
	for _, unsaved := range []bool{false, true} {
		dir, home := fixture(t)
		screen := opened(t, dir, "test-dusk", home)
		if unsaved {
			screen.press("up")
		}
		for _, message := range []tea.Msg{key("[", '['), tea.KeyPressMsg{Code: tea.KeyUp},
			tea.KeyPressMsg{Code: tea.KeyEnter}, tea.KeyPressMsg{Code: tea.KeyDown},
			tea.KeyPressMsg{Code: tea.KeyEnter}} {
			drive(screen, message)
		}
		if screen.menu != nil {
			t.Fatalf("unsaved %v: the menu is still open", unsaved)
		}
		if screen.set.source.Name != "plain" || screen.page != claudePage {
			t.Errorf("unsaved %v: on %s page %d, want plain page %d", unsaved,
				screen.set.source.Name, screen.page, claudePage)
		}
		view := screen.View().Content
		if !strings.Contains(view, groundCode(screen.set.shown("ground"))) {
			t.Errorf("unsaved %v: the page is not drawn on plain's ground", unsaved)
		}
	}
}

// ctrl+c leaves from anywhere, the menu open or not.
func TestControlCLeaves(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.openMenu()
	_, command := screen.Update(tea.KeyPressMsg{Code: 'c', Mod: tea.ModCtrl})
	if command == nil {
		t.Fatal("ctrl+c did nothing with the menu open")
	}
	quit := false
	for _, message := range run(command) {
		_, isQuit := message.(tea.QuitMsg)
		quit = quit || isQuit
	}
	if !quit {
		t.Error("ctrl+c did not quit")
	}
}

// a page drawn mostly in fallbacks says so, and one the colourway sets
// does not.
func TestUnsetNote(t *testing.T) {
	dir, home := fixture(t)
	plain := opened(t, dir, "plain", home)
	plain.page = vimPage
	if note := plain.unsetNote(); !strings.Contains(note, "sets 2 of") {
		t.Errorf("plain's vim page says %q", note)
	}
	full := opened(t, dir, "test-dusk", home)
	full.page = vimPage
	if note := full.unsetNote(); note != "" {
		t.Errorf("test-dusk's vim page says %q", note)
	}
}

// once the page is drawn, tab steers only colours it draws: on the shell
// page that is all sixteen, since the colour test names each; on the
// markdown page it is never the block's text colour, which the sample
// draws nothing in.
func TestTabSteersWhatIsDrawn(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	for _, which := range []page{shellPage, markdownPage, claudePage} {
		screen.page, screen.steering = which, screen.first(which)
		screen.View()
		drawn, visited := screen.drawnRoles(), map[string]bool{}
		for range screen.stops {
			screen.press("tab")
			name := screen.stops[screen.steering].name
			if !drawn[name] {
				t.Errorf("on page %d tab steered %s, which nothing draws", which, name)
			}
			visited[name] = true
		}
		if which == shellPage {
			for _, name := range ghostty.Ansi {
				if !visited[name] {
					t.Errorf("tab never reached %s on the shell page", name)
				}
			}
		}
		if visited["md-block-text"] {
			t.Error("tab steered md-block-text on the markdown page")
		}
		if !visited[screen.set.groundName(which)] {
			t.Errorf("tab never reached the ground on page %d", which)
		}
	}
}

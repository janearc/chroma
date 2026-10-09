package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/janearc/libtheme-css/colourway"
	"github.com/janearc/libtheme-css/css"
	"github.com/janearc/libtheme-css/dialects/ghostty"
	"github.com/janearc/libtheme-css/dialects/nvim"
	"github.com/janearc/libtheme-css/spaces/srgb"
)

// terminalRoles are what a terminal theme needs: the ground, the ink and
// the sixteen.
func terminalRoles() *css.Sheet {
	roles := css.New()
	roles.Set("ground", srgb.MustHex("#808a63").Swatch())
	roles.Set("ink", srgb.MustHex("#38241e").Swatch())
	for number, name := range ghostty.Ansi {
		roles.Set(name, srgb.RGB8(uint8(number*15), 60, uint8(200-number*10)).Swatch())
	}
	return roles
}

// fullRoles are a terminal's, an editor's and claude code's text: every
// program renders.
func fullRoles() *css.Sheet {
	roles := terminalRoles()
	for position, entry := range nvim.Roles {
		if _, found := roles.Get(entry.Role); !found {
			roles.Set(entry.Role, srgb.RGB8(uint8(40+position*9), 30, 50).Swatch())
		}
	}
	roles.Set("claude-text", srgb.MustHex("#c5e699").Swatch())
	return roles
}

// fixture is a directory of colourways, written as a save writes them,
// and an empty home to install into: test-dusk, which sets every
// program's roles, and plain, which sets only a terminal's, the rest
// derived by libtheme.
func fixture(t *testing.T) (dir, home string) {
	t.Helper()
	t.Setenv("XDG_CONFIG_HOME", "")
	dir, home = t.TempDir(), t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "sources"), 0o755); err != nil {
		t.Fatal(err)
	}
	for name, roles := range map[string]*css.Sheet{"test-dusk": fullRoles(),
		"plain": terminalRoles()} {
		source := colourway.Source{Name: name, About: name + ", for a test",
			Roles: roles, Origin: "sources/" + name + ".css"}
		if err := source.Write(sourcePath(dir, name)); err != nil {
			t.Fatal(err)
		}
	}
	return dir, home
}

// bare writes a colourway with only a ground and an ink, which no
// program can render: no sixteen for the terminal or the editor, and
// nothing but the ink for markdown.
func bare(t *testing.T, dir string) {
	t.Helper()
	roles := css.New()
	roles.Set("ground", srgb.MustHex("#808a63").Swatch())
	roles.Set("ink", srgb.MustHex("#38241e").Swatch())
	source := colourway.Source{Name: "bare", About: "bare, for a test", Roles: roles,
		Origin: "sources/bare.css"}
	if err := source.Write(sourcePath(dir, "bare")); err != nil {
		t.Fatal(err)
	}
}

// sourcePath is where a colourway's source sits.
func sourcePath(dir, name string) string {
	return filepath.Join(dir, "sources", name+".css")
}

// opened is paratune with a colourway open, sized as her pane is.
func opened(t *testing.T, dir, name, home string) *model {
	t.Helper()
	set, err := load(dir, name)
	if err != nil {
		t.Fatal(err)
	}
	screen := &model{home: home, text: "a line of text\n\nanother, with more words in it",
		set: set, stops: stopsOf(set.roles), width: 88, height: 34}
	screen.steering = screen.first(shellPage)
	return screen
}

// steer turns to a stop by its name on a page, failing if there is none.
func steer(t *testing.T, screen *model, which page, name string) stop {
	t.Helper()
	for index, here := range screen.stops {
		if here.page == which && here.name == name {
			screen.page, screen.steering = which, index
			return here
		}
	}
	t.Fatalf("no %s on page %d", name, which)
	return stop{}
}

// read is a file's text, failing the test if it cannot be read.
func read(t *testing.T, path string) string {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(raw)
}

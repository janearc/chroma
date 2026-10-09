package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/janearc/libtheme-css/colourway"
	"github.com/janearc/libtheme-css/css"
	"github.com/janearc/libtheme-css/spaces/srgb"
)

// held is a source exactly as a save writes it.
func held() string {
	roles := css.New()
	roles.Set("ground", srgb.MustHex("#808a63").Swatch())
	roles.Set("ink", srgb.MustHex("#38241e").Swatch())
	return colourway.Commented("a colourway\nin two lines") + roles.String()
}

// a source as a save writes it is held. Anything a save would not write
// back the same is refused, by its line: whatever it is, not only the
// kinds listed here, since the check is the round trip itself.
func TestUnheld(t *testing.T) {
	if lost := unheld(held()); lost != "" {
		t.Fatalf("a written source would lose %s", lost)
	}
	ink := "  --ink: #38241e; /* oklch(28% 0.033 37.4) */"
	for _, case_ := range []struct{ name, from, to string }{
		{"an eight-digit hex", ink, "  --ink: #38241e80;"},
		{"capitals in a hex", ink, strings.ToUpper(ink[:16]) + ink[16:]},
		{"a comment of its own", ink, ink + "\n  /* the brown */"},
		{"an oklch note edited", ink, "  --ink: #38241e; /* my brown */"},
		{"a role named twice", ink, ink + "\n  --ink: #000000; /* oklch(0% 0 none) */"},
		{"a second selector", "}\n", "}\n.dark { --ink: #000000; }\n"},
		{"a colour not in hex", ink, "  --ink: oklch(28% 0.033 37.4);"},
		{"a blank line inside", ink, "\n" + ink},
		{"no last newline", "}\n", "}"},
		{"a property that is not a colour", ink, ink + "\n  --width: 4px;"},
	} {
		text := strings.Replace(held(), case_.from, case_.to, 1)
		if text == held() {
			t.Fatalf("%s: the case changed nothing", case_.name)
		}
		if unheld(text) == "" {
			t.Errorf("%s was taken as held:\n%s", case_.name, text)
		}
	}
}

// the line unheld names is the first that a save would change.
func TestUnheldNamesTheLine(t *testing.T) {
	text := strings.Replace(held(), "  --ink: #38241e;", "  --ink: #38241E;", 1)
	lost := unheld(text)
	if !strings.HasPrefix(lost, "line 7:") || !strings.Contains(lost, "#38241E") {
		t.Errorf("named %q", lost)
	}
}

// the roles put in are the sheet that comes out, in the order they were
// first put, at the byte shown.
func TestRolesSheet(t *testing.T) {
	all := rolesOf(css.New())
	all.put("ink", colour{56.4, 36, 30})
	all.put("ground", colour{128, 138, 99})
	all.put("ink", colour{57, 36, 30})
	if names := strings.Join(all.sheet().Names(), " "); names != "ink ground" {
		t.Errorf("order %q", names)
	}
	for _, rule := range all.sheet().Rules() {
		if rule.Name == "ink" && rule.Value != "#39241e" {
			t.Errorf("ink saved as %s", rule.Value)
		}
	}
}

// a move saved is in the source, every other role as it was, and the
// colourway is clean again.
func TestSaveWritesTheMove(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	before := read(t, sourcePath(dir, "test-dusk"))
	steer(t, screen, shellPage, "ground")
	screen.press("up")
	if !screen.set.dirty {
		t.Fatal("a move did not mark the colourway unsaved")
	}
	note := screen.set.save()
	if note != "saved and rendered" {
		t.Errorf("saving said %q", note)
	}
	after := read(t, sourcePath(dir, "test-dusk"))
	changed := 0
	beforeLines, afterLines := strings.Split(before, "\n"), strings.Split(after, "\n")
	if len(beforeLines) != len(afterLines) {
		t.Fatalf("%d lines became %d", len(beforeLines), len(afterLines))
	}
	for number := range beforeLines {
		if beforeLines[number] != afterLines[number] {
			changed++
			if !strings.Contains(afterLines[number], "--ground:") {
				t.Errorf("a line that is not the ground changed: %s", afterLines[number])
			}
		}
	}
	if changed != 1 || screen.set.dirty {
		t.Errorf("%d lines changed, dirty %v", changed, screen.set.dirty)
	}
	theme := read(t, filepath.Join(dir, "ghostty", "test-dusk"))
	if !strings.Contains(theme, "background = "+screen.set.shown("ground").hex()) {
		t.Error("the theme was not rendered with the moved ground")
	}
}

// a source holding something a save would lose is not written, not by a
// byte, and saving says why.
func TestSaveRefusesALossySource(t *testing.T) {
	dir, home := fixture(t)
	path := sourcePath(dir, "test-dusk")
	lossy := strings.Replace(read(t, path), ":root {", ":root {\n  /* mine */", 1)
	if err := os.WriteFile(path, []byte(lossy), 0o644); err != nil {
		t.Fatal(err)
	}
	screen := opened(t, dir, "test-dusk", home)
	screen.press("up")
	note := screen.set.save()
	if !strings.HasPrefix(note, "not saved") || !strings.Contains(note, "/* mine */") {
		t.Errorf("saving said %q", note)
	}
	if read(t, path) != lossy {
		t.Error("the lossy source was written")
	}
	if !screen.set.dirty {
		t.Error("the refused move was taken as saved")
	}
}

// saving says which programs refused to render, and why: a colourway
// with only a ground and an ink gives none of them enough.
func TestSaveSaysWhatRefused(t *testing.T) {
	dir, home := fixture(t)
	bare(t, dir)
	screen := opened(t, dir, "bare", home)
	screen.press("up")
	note := screen.set.save()
	for _, want := range []string{"not rendered", "ghostty (", "nvim (", "glamour ("} {
		if !strings.Contains(note, want) {
			t.Errorf("saving said %q", note)
		}
	}
}

// installing with unsaved moves writes nothing.
func TestInstallRefusesUnsaved(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	screen.press("up")
	if note := screen.set.install(home); !strings.HasPrefix(note, "unsaved") {
		t.Errorf("install said %q", note)
	}
	if entries, _ := os.ReadDir(home); len(entries) != 0 {
		t.Error("install wrote into the home while unsaved")
	}
}

// install renders the saved source afresh and copies exactly that: the
// files it copies are the files the render wrote, the moved colour in
// them, the colourway and its twin for each program.
func TestInstallCopiesAFreshRender(t *testing.T) {
	dir, home := fixture(t)
	screen := opened(t, dir, "test-dusk", home)
	steer(t, screen, shellPage, "ground")
	screen.press("up")
	screen.set.save()
	note := screen.set.install(home)
	if note != "installed 8 files where each program reads them" {
		t.Errorf("install said %q", note)
	}
	for rendered, installed := range map[string]string{
		"ghostty/test-dusk":                  ".config/ghostty/themes/test-dusk",
		"ghostty/test-dusk-inverted":         ".config/ghostty/themes/test-dusk-inverted",
		"nvim/colors/test-dusk.lua":          ".config/nvim/colors/test-dusk.lua",
		"nvim/colors/test-dusk-inverted.lua": ".config/nvim/colors/test-dusk-inverted.lua",
		"vim/colors/test-dusk.vim":           ".vim/colors/test-dusk.vim",
		"vim/colors/test-dusk-inverted.vim":  ".vim/colors/test-dusk-inverted.vim",
		"glamour/test-dusk.json":             ".config/glamour/test-dusk.json",
		"glamour/test-dusk-inverted.json":    ".config/glamour/test-dusk-inverted.json",
	} {
		if read(t, filepath.Join(dir, rendered)) != read(t, filepath.Join(home, installed)) {
			t.Errorf("%s is not what was rendered", installed)
		}
	}
	theme := read(t, filepath.Join(home, ".config/ghostty/themes/test-dusk"))
	if !strings.Contains(theme, "background = "+screen.set.shown("ground").hex()) {
		t.Error("the installed theme does not have the moved ground")
	}
}

// a file left in the colourways by an older render, for a program the
// source no longer renders, is never installed.
func TestInstallNeverCopiesStale(t *testing.T) {
	dir, home := fixture(t)
	bare(t, dir)
	stale := filepath.Join(dir, "nvim", "colors", "bare.lua")
	if err := os.MkdirAll(filepath.Dir(stale), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(stale, []byte("-- stale\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	screen := opened(t, dir, "bare", home)
	note := screen.set.install(home)
	if _, err := os.Stat(filepath.Join(home, ".config/nvim/colors/bare.lua")); err == nil {
		t.Error("the stale scheme was installed")
	}
	if !strings.HasPrefix(note, "installed 0 files") || !strings.Contains(note, "nvim (") {
		t.Errorf("install said %q", note)
	}
}

// TestInstallSaysWhatItReplaced installs over a different theme of the
// same name, as i is pressed to do, and says so.
func TestInstallSaysWhatItReplaced(t *testing.T) {
	dir, home := fixture(t)
	theirs := filepath.Join(home, ".config/ghostty/themes/test-dusk")
	if err := os.MkdirAll(filepath.Dir(theirs), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(theirs, []byte("background = #000000\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	screen := opened(t, dir, "test-dusk", home)
	note := screen.set.install(home)
	if !strings.Contains(note, "1 took the place of a different file") {
		t.Errorf("install over a different theme said %q", note)
	}
}

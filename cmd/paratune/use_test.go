package main

import (
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

// themeLine is ghostty's pointer's pattern, the one the tests repoint.
var themeLine = pointers[0].pattern

// repoint plans and writes one config's change, as use does for each of
// its three once all three are planned.
func repoint(path string, pattern *regexp.Regexp, line string) error {
	each, err := plan(path, pattern, line)
	if err != nil {
		return err
	}
	return each.write()
}

// written makes a file with some text in a fresh directory.
func written(t *testing.T, name, text string, permissions os.FileMode) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(text), permissions); err != nil {
		t.Fatal(err)
	}
	return path
}

// TestRepointOneLine replaces the one line that names a theme, keeps the
// file as it was beside it, and leaves the other lines alone.
func TestRepointOneLine(t *testing.T) {
	before := "font-size = 15\ntheme = fermi-midnight\nfaint-opacity = 0.5\n"
	path := written(t, "config", before, 0o600)
	if err := repoint(path, themeLine, "theme = pluto"); err != nil {
		t.Fatal(err)
	}
	after, _ := os.ReadFile(path)
	if string(after) != "font-size = 15\ntheme = pluto\nfaint-opacity = 0.5\n" {
		t.Errorf("after: %q", after)
	}
	kept, _ := os.ReadFile(path + ".bak-paratune")
	if string(kept) != before {
		t.Errorf("backup: %q", kept)
	}
	if info, _ := os.Stat(path); info.Mode().Perm() != 0o600 {
		t.Errorf("mode %v, want 0600", info.Mode().Perm())
	}
	leftover, _ := filepath.Glob(filepath.Join(filepath.Dir(path), ".config.paratune-*"))
	if len(leftover) > 0 {
		t.Errorf("temporary files left: %v", leftover)
	}
}

// TestRepointKeepsFirstBackup keeps the file from before paratune ever
// changed it, however many times it switches.
func TestRepointKeepsFirstBackup(t *testing.T) {
	path := written(t, "config", "theme = first\n", 0o644)
	for _, name := range []string{"second", "third"} {
		if err := repoint(path, themeLine, "theme = "+name); err != nil {
			t.Fatal(err)
		}
	}
	kept, _ := os.ReadFile(path + ".bak-paratune")
	if string(kept) != "theme = first\n" {
		t.Errorf("backup: %q", kept)
	}
}

// TestRepointFollowsSymlink changes the file a link points to and leaves
// the link a link.
func TestRepointFollowsSymlink(t *testing.T) {
	target := written(t, "dotfiles/config", "theme = a\n", 0o644)
	link := filepath.Join(t.TempDir(), "config")
	if err := os.Symlink(target, link); err != nil {
		t.Fatal(err)
	}
	if err := repoint(link, themeLine, "theme = b"); err != nil {
		t.Fatal(err)
	}
	if info, _ := os.Lstat(link); info.Mode()&os.ModeSymlink == 0 {
		t.Error("the link was replaced by a file")
	}
	after, _ := os.ReadFile(target)
	if string(after) != "theme = b\n" {
		t.Errorf("target: %q", after)
	}
	if _, err := os.Stat(target + ".bak-paratune"); err != nil {
		t.Errorf("no backup beside the target: %v", err)
	}
}

// TestRepointKeepsLineEndings keeps a carriage return at the end of the
// line it replaces.
func TestRepointKeepsLineEndings(t *testing.T) {
	path := written(t, "config", "theme = a\r\nfont-size = 15\r\n", 0o644)
	if err := repoint(path, themeLine, "theme = b"); err != nil {
		t.Fatal(err)
	}
	after, _ := os.ReadFile(path)
	if string(after) != "theme = b\r\nfont-size = 15\r\n" {
		t.Errorf("after: %q", after)
	}
}

// TestRepointRefuses leaves a file alone when the line is missing or there
// is more than one of it.
func TestRepointRefuses(t *testing.T) {
	for _, text := range []string{"font-size = 15\n", "theme = a\ntheme = b\n", "# theme = a\n"} {
		path := written(t, "config", text, 0o644)
		if err := repoint(path, themeLine, "theme = pluto"); err == nil {
			t.Errorf("%q: no error", text)
		}
		after, _ := os.ReadFile(path)
		if string(after) != text {
			t.Errorf("%q changed to %q", text, after)
		}
		if _, err := os.Stat(path + ".bak-paratune"); err == nil {
			t.Errorf("%q: a backup was written for a refused change", text)
		}
	}
}

// TestUseChangesNothingOnRefusal reads every config before writing any: a
// neovim config without its colorscheme line leaves ghostty's and glow's
// as they were.
func TestUseChangesNothingOnRefusal(t *testing.T) {
	home := t.TempDir()
	texts := map[string]string{
		".config/ghostty/config":            "theme = fermi-midnight\n",
		".config/nvim/init.lua":             "vim.opt.number = true\n",
		"Library/Preferences/glow/glow.yml": "style: \"auto\"\n",
	}
	for file, text := range texts {
		path := filepath.Join(home, file)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(text), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if err := use(t.TempDir(), "pluto", home); err == nil {
		t.Fatal("no error for a config without its line")
	}
	for file, text := range texts {
		after, _ := os.ReadFile(filepath.Join(home, file))
		if string(after) != text {
			t.Errorf("%s changed to %q", file, after)
		}
	}
}

// TestUseRefusesANameOrNoHome keeps a name that could break a config line
// or a neovim command out of both, and will not switch without a home.
func TestUseRefusesANameOrNoHome(t *testing.T) {
	for _, name := range []string{`pluto")`, "Pluto", "a b", ""} {
		if err := use(t.TempDir(), name, t.TempDir()); err == nil {
			t.Errorf("%q passed", name)
		}
	}
	if err := use(t.TempDir(), "pluto", ""); err == nil {
		t.Error("an empty home passed")
	}
}

// TestPointersMatchTheirLines checks each real pattern against the line it
// is for, and that the line it writes is one it would match again.
func TestPointersMatchTheirLines(t *testing.T) {
	lines := []string{"theme = fermi-midnight", `vim.cmd.colorscheme("twilight-deep")`,
		`style: "~/.config/glamour/aqua.json"`}
	for index, each := range pointers {
		if !each.pattern.MatchString(lines[index]) {
			t.Errorf("%s: misses %q", each.file("~"), lines[index])
		}
		written := each.line("/home", "pluto")
		if !regexp.MustCompile(each.pattern.String()).MatchString(written) {
			t.Errorf("%s: writes %q, which it would not find again", each.file("~"), written)
		}
	}
}

// TestPointersFollowTheConfigFolder finds ghostty's and neovim's configs,
// and the style glow is pointed at, in the folder XDG_CONFIG_HOME names,
// where install puts the themes.
func TestPointersFollowTheConfigFolder(t *testing.T) {
	home, config := t.TempDir(), t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", config)
	if got := pointers[0].file(home); got != filepath.Join(config, "ghostty/config") {
		t.Errorf("ghostty's config is %s", got)
	}
	if got := pointers[1].file(home); got != filepath.Join(config, "nvim/init.lua") {
		t.Errorf("neovim's config is %s", got)
	}
	want := `style: "` + filepath.Join(config, "glamour/pluto.json") + `"`
	if got := pointers[2].line(home, "pluto"); got != want {
		t.Errorf("glow is pointed with %s", got)
	}
}

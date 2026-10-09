package main

import (
	"os"
	"path/filepath"
	"testing"
)

// TestInstallVerbInstallsTheRepository installs the repository's own
// colourways into an empty home, and finds one theme where ghostty reads
// it and one source where paratune looks.
func TestInstallVerbInstallsTheRepository(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_CONFIG_HOME", "")
	t.Setenv("LANG", "en_GB.UTF-8")
	if err := install([]string{filepath.Join("..", "..", "colourways")}); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{".config/ghostty/themes/vaporwave-dusk",
		".config/colourways/sources/vaporwave-dusk.css",
		".vim/colors/vaporwave-dusk.vim"} {
		if _, err := os.Stat(filepath.Join(home, path)); err != nil {
			t.Errorf("no %s", path)
		}
	}
}

// TestInstallVerbWantsFlagsFirst refuses a flag after the folder, which
// go's flags would otherwise pass over in silence.
func TestInstallVerbWantsFlagsFirst(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	if err := install([]string{"colourways", "--replace"}); err == nil {
		t.Error("a flag after the folder was taken without a word")
	}
}

package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/janearc/libtheme-css/colourway"
)

// TestShippedRendersAreTheirSources renders every colourway in the
// repository again and finds each file the same as the one shipped, so a
// render is never stale and never edited by hand. A program skipped is a
// failure here, though render allows it: the readme tells people to
// install each shipped colourway in every program.
func TestShippedRendersAreTheirSources(t *testing.T) {
	shipped := filepath.Join("..", "..", "colourways")
	sources, err := filepath.Glob(filepath.Join(shipped, "sources", "*.css"))
	if err != nil || len(sources) == 0 {
		t.Fatalf("no sources under %s: %v", shipped, err)
	}
	made := map[string]bool{}
	for _, path := range sources {
		source, err := colourway.Read(path)
		if err != nil {
			t.Fatal(err)
		}
		dir := t.TempDir()
		for _, written := range colourway.Render(source, dir) {
			if written.Err != nil {
				t.Errorf("%s: %v", source.Name, written.Err)
				continue
			}
			fresh, err := os.ReadFile(written.Path)
			if err != nil {
				t.Fatal(err)
			}
			relative := strings.TrimPrefix(written.Path, dir+string(filepath.Separator))
			made[filepath.ToSlash(relative)] = true
			old, err := os.ReadFile(filepath.Join(shipped, relative))
			if err != nil {
				t.Errorf("%s: %s is not shipped", source.Name, relative)
				continue
			}
			if !bytes.Equal(fresh, old) {
				t.Errorf("%s: %s is not what its source renders; run colourway render",
					source.Name, relative)
			}
		}
	}
	found, err := orphans(shipped, made)
	if err != nil {
		t.Fatal(err)
	}
	for _, orphan := range found {
		t.Errorf("%s is shipped, and no source renders it", orphan)
	}
}

// orphans are the shipped renders no source makes, as when a source is
// renamed and its old files stay behind.
func orphans(shipped string, made map[string]bool) ([]string, error) {
	found := []string{}
	for _, folder := range []string{"ghostty", "nvim/colors", "vim/colors", "glamour"} {
		entries, err := os.ReadDir(filepath.Join(shipped, folder))
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			relative := folder + "/" + entry.Name()
			hidden := strings.HasPrefix(entry.Name(), ".")
			if !entry.IsDir() && !hidden && !made[relative] {
				found = append(found, relative)
			}
		}
	}
	return found, nil
}

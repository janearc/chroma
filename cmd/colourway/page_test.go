package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestPageIsTheReadmeInOneFile builds the page from the repository's
// colourways and finds it english, every picture inside it, every
// colourway a name at the top, and every place for words shown.
func TestPageIsTheReadmeInOneFile(t *testing.T) {
	folder := filepath.Join("..", "..", "colourways")
	out := filepath.Join(t.TempDir(), "colourways.html")
	if err := page(folder, out); err != nil {
		t.Fatal(err)
	}
	raw, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	text := string(raw)
	sources, _ := filepath.Glob(filepath.Join(folder, "sources", "*.css"))
	readme, _ := os.ReadFile(filepath.Join(folder, "README.md"))
	for want, count := range map[string]int{
		`<html lang="en">`:            1,
		`src="data:image/png;base64,`: strings.Count(string(readme), "<img "),
		`<span id="in-`:               len(sources),
		`:root:has(#in-`:              len(sources),
		`(words: `:                    strings.Count(string(readme), "<!-- words: "),
	} {
		if got := strings.Count(text, want); got != count {
			t.Errorf("%q is in the page %d times, not %d", want, got, count)
		}
	}
	if strings.Contains(text, `src="pictures/`) {
		t.Error("a picture is named by its path, not carried inside the page")
	}
}

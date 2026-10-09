package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	gmhtml "github.com/yuin/goldmark/renderer/html"
)

// pageDefault is the colourway a page opens in: a dark one turned down,
// so nobody is given a light page unless they choose one.
const pageDefault = "vaporwave-dusk"

// cleared are the roles the page reads, cleared at the top of each
// colourway's block, so a role a colourway does not set falls back to its
// own ink and not to the default's.
var cleared = []string{"--md-body", "--md-heading", "--md-code",
	"--md-code-ground", "--md-rule", "--md-quote", "--heading", "--code",
	"--panel", "--dim", "--border"}

// wordsMark is a place left for the author's words, and pictureSrc a
// picture the readme names by its path in the folder.
var (
	wordsMark  = regexp.MustCompile(`(?m)^<!-- (words: .*) -->$`)
	pictureSrc = regexp.MustCompile(`src="(pictures/[a-z0-9-]+\.png)"`)
)

// pictures draws every colourway in a colourways folder into its
// pictures folder, a png named for each source.
func pictures(folder string) error {
	sources, err := filepath.Glob(filepath.Join(folder, "sources", "*.css"))
	if err != nil {
		return err
	}
	if len(sources) == 0 {
		return fmt.Errorf("%s has no sources/*.css to draw", folder)
	}
	drawn := filepath.Join(folder, "pictures")
	if err := os.MkdirAll(drawn, 0o755); err != nil {
		return err
	}
	for _, source := range sources {
		name := strings.TrimSuffix(filepath.Base(source), ".css")
		out := filepath.Join(drawn, name+".png")
		if err := picture(source, out); err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
		fmt.Println("drew", name)
	}
	return nil
}

// page writes a colourways folder's readme as one page, with every
// picture inside it, drawn in the default colourway, every colourway a
// click away by name at its top, and the places for words shown.
func page(folder, out string) error {
	readme, err := os.ReadFile(filepath.Join(folder, "README.md"))
	if err != nil {
		return err
	}
	sheets, names, err := themeSheets(folder)
	if err != nil {
		return err
	}
	style, err := os.ReadFile(filepath.Join(folder, "preview.css"))
	if err != nil {
		return err
	}
	text := wordsMark.ReplaceAll(readme, []byte("> *($1)*"))
	var body bytes.Buffer
	markdown := goldmark.New(goldmark.WithExtensions(extension.GFM),
		goldmark.WithRendererOptions(gmhtml.WithUnsafe()))
	if err := markdown.Convert(text, &body); err != nil {
		return err
	}
	inlined, err := inlinePictures(folder, body.String())
	if err != nil {
		return err
	}
	var doc strings.Builder
	doc.WriteString("<!doctype html>\n<html lang=\"en\">\n<head>\n" +
		"<meta charset=\"utf-8\">\n" +
		"<meta name=\"viewport\" content=\"width=device-width\">\n" +
		"<title>colourways</title>\n<style>\n")
	doc.WriteString(sheets + string(style))
	doc.WriteString("</style>\n</head>\n<body>\n")
	doc.WriteString(namesRow(names) + inlined + "</body>\n</html>\n")
	if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(out, []byte(doc.String()), 0o644); err != nil {
		return err
	}
	fmt.Printf("wrote %s, in %s; the names at its top change it\n",
		out, pageDefault)
	return nil
}

// themeSheets is every colourway's sheet for the page: the default's as
// the page's own, then each behind its name, its roles cleared first; and
// the colourways' names, in order.
func themeSheets(folder string) (string, []string, error) {
	sources, err := filepath.Glob(filepath.Join(folder, "sources", "*.css"))
	if err != nil {
		return "", nil, err
	}
	first := filepath.Join(folder, "sources", pageDefault+".css")
	base, err := os.ReadFile(first)
	if err != nil {
		return "", nil, err
	}
	reset := strings.Join(cleared, ": initial; ") + ": initial;"
	sheets, names := []string{string(base)}, []string{}
	for _, source := range sources {
		raw, err := os.ReadFile(source)
		if err != nil {
			return "", nil, err
		}
		name := strings.TrimSuffix(filepath.Base(source), ".css")
		opener := fmt.Sprintf(":root:has(#in-%s:target) { %s",
			name, reset)
		sheet := strings.Replace(string(raw), ":root {", opener, 1)
		sheets = append(sheets, sheet)
		names = append(names, name)
	}
	return strings.Join(sheets, "\n"), names, nil
}

// namesRow is the row of names at the top: each a link that turns the
// page into that colourway, beside an anchor the link targets.
func namesRow(names []string) string {
	var row strings.Builder
	row.WriteString("<p class=\"themes\">read this page in:\n")
	for _, name := range names {
		escaped := html.EscapeString(name)
		fmt.Fprintf(&row, "<span id=\"in-%s\"></span>", escaped)
		fmt.Fprintf(&row, "<a href=\"#in-%s\">%s</a>\n",
			escaped, escaped)
	}
	row.WriteString("</p>\n")
	return row.String()
}

// inlinePictures puts each picture the page names inside it, so the page
// is one file that can go anywhere.
func inlinePictures(folder, body string) (string, error) {
	var missing error
	carry := func(src string) string {
		path := pictureSrc.FindStringSubmatch(src)[1]
		raw, err := os.ReadFile(filepath.Join(folder, path))
		if err != nil {
			missing = err
			return src
		}
		return `src="data:image/png;base64,` +
			base64.StdEncoding.EncodeToString(raw) + `"`
	}
	return pictureSrc.ReplaceAllStringFunc(body, carry), missing
}

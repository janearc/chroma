package main

import (
	"github.com/alecthomas/chroma/v2/styles"
	"slices"
	"strings"

	"charm.land/glamour/v2"
	gstyle "github.com/janearc/libtheme-css/dialects/glamour"
)

// markdownCode is the markdown page's code block: short, so the page
// fits, with one of each kind of token chroma colours. its indents are
// spaces, since a tab drawn into a frame has no tab stops to go to.
const markdownCode = `// sweep is the wheel as a ramp mixed in one space.
func sweep(in functions.Mixer) functions.Ramp {
    const space = "oklab"
    swatches := make([]swatch.Swatch, 7)
    for position, hex := range wheel {
        swatches[position] = srgb.MustHex(hex).Swatch()
    }
    return functions.Even(in, swatches...)
}`

// sampleMarkdown is what paratune hands glamour: a short readme with one
// of everything her style colours, and a little go as its code block.
const sampleMarkdown = "# ramps\n\n" +
	"A ramp turns a number between 0 and 1 into a colour, from a few\n" +
	"**stops** and a way of *mixing* them. " +
	"[oklab](https://bottosson.github.io/posts/oklab/)\n" +
	"is where a straight line looks straight; " +
	"`srgb.Mix` goes through mud.\n\n" +
	"## using one\n\n" +
	"- build it with `functions.Even` and the stops in order\n" +
	"- ask it for a colour with `At`, or a strip with `Samples`\n\n" +
	"> a ramp answers what colour t is, " +
	"and never asks which cell it is in.\n\n" +
	"```go\n" + markdownCode + "\n```\n\n" +
	"### where t comes from\n\n" +
	"That is the drawing's business, not the ramp's.\n"

// glamourChroma is the name glamour registers its code colours under.
const glamourChroma = "charm"

// markdownSample is the sample as glamour draws it in the colourway's
// style, said by libtheme's glamour dialect, wrapped to width, with
// chroma colouring the code in true colour. If either refuses, the
// reason is shown in its place, so a broken style is seen and not hidden.
func markdownSample(set *tuning, width int) []string {
	style, err := gstyle.Of(set.resolution().Roles)
	if err != nil {
		return []string{"glamour: " + err.Error()}
	}
	// glamour registers its code colours with chroma once, under one name,
	// and never again, so every later render drew code in the first
	// colours it was given. taking the name away makes it register them
	// fresh from this style.
	delete(styles.Registry, glamourChroma)
	renderer, err := glamour.NewTermRenderer(
		glamour.WithStylesFromJSONBytes(style),
		glamour.WithWordWrap(width),
		glamour.WithChromaFormatter("terminal16m"))
	if err != nil {
		return []string{"glamour: " + err.Error()}
	}
	drawn, err := renderer.Render(sampleMarkdown)
	if err != nil {
		return []string{"glamour: " + err.Error()}
	}
	return strings.Split(strings.TrimRight(drawn, "\n"), "\n")
}

// markdownWords are words of the sample that show a role inside a line
// of body text.
var markdownWords = []struct {
	word string
	role []string
}{
	{"few stops", []string{"md-strong"}},
	{"mixing", []string{"md-emphasis"}},
	{"oklab", []string{"md-link", "md-link-text"}},
	{"srgb.Mix", []string{"md-code", "md-code-ground"}},
	{"functions.Even", []string{"md-code", "md-code-ground"}},
	{" At ", []string{"md-code", "md-code-ground"}},
	{"Samples", []string{"md-code", "md-code-ground"}},
}

// markdownLines are glamour's drawing of the sample, each line with the
// roles it shows: headings by their marks, the quote by its bar, bullets
// by theirs, code by chroma's reading of the same line, and body text by
// the words in it.
//
// A code line shows the block's text colour only where a character chroma
// leaves unstyled is drawn in it, not its spaces.
func markdownLines(set *tuning, width int) []line {
	code := map[string][]string{}
	for _, tokens := range tokenLines(markdownCode) {
		text, shows := "", []string{}
		for _, token := range tokens {
			text += token.Value
			name := chromaRole(token.Type)
			if name == "" && strings.TrimSpace(token.Value) != "" {
				name = "md-block-text"
			}
			if name != "" && !slices.Contains(shows, name) {
				shows = append(shows, name)
			}
		}
		code[strings.TrimSpace(text)] = shows
	}
	drawn := markdownSample(set, width)
	lines := make([]line, len(drawn))
	for position, text := range drawn {
		bare := strings.TrimSpace(escapes.ReplaceAllString(text, ""))
		lines[position] = line{text: text,
			shows: markdownShows(bare, code)}
	}
	return lines
}

// markdownShows is the roles one line of glamour's drawing shows.
func markdownShows(plain string, code map[string][]string) []string {
	if shows, ok := code[plain]; ok && plain != "" {
		return shows
	}
	switch {
	case plain == "":
		return nil
	case strings.HasPrefix(plain, "### "):
		return []string{"md-small-heading", "md-small-heading-ground"}
	case strings.HasPrefix(plain, "#"):
		return []string{"md-heading", "md-heading-ground"}
	case strings.HasPrefix(plain, "│"):
		return []string{"md-quote"}
	}
	shows := []string{"md-body"}
	if strings.HasPrefix(plain, "•") {
		shows = append(shows, "md-bullet")
	}
	for _, entry := range markdownWords {
		if strings.Contains(" "+plain+" ", entry.word) {
			shows = append(shows, entry.role...)
		}
	}
	return shows
}

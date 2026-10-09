package main

import (
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
)

// sampleCode is the go both code samples show: glamour's code block on
// the markdown page, and the file open in vim on the vim page.
const sampleCode = `package main

import (
	"fmt"

	"github.com/janearc/libtheme-css/spaces/ok"
)

// gradientWidth is how many columns the gradient takes.
const gradientWidth = 6

// sweep is the wheel as a ramp mixed in one space.
func sweep(in functions.Mixer) functions.Ramp {
	swatches := make([]swatch.Swatch, len(wheel))
	for position, hex := range wheel {
		swatches[position] = srgb.MustHex(hex).Swatch()
	}
	return functions.Even(in, swatches...)
}

func unused() {
	ramp := sweep(ok.Mix)
	fmt.Println("a ramp of", 7, "stops")
}`

// tokenLines are go source as chroma's lexer reads it, a line at a time:
// each token's kind and its text, with tabs opened to four spaces.
func tokenLines(source string) [][]chroma.Token {
	lexer := lexers.Get("go")
	iterator, err := lexer.Tokenise(nil, source)
	if err != nil {
		return [][]chroma.Token{{{Type: chroma.Text, Value: source}}}
	}
	lines := [][]chroma.Token{{}}
	for _, token := range iterator.Tokens() {
		tabbed := strings.ReplaceAll(token.Value, "\t", "    ")
		pieces := strings.Split(tabbed, "\n")
		for position, piece := range pieces {
			if position > 0 {
				lines = append(lines, []chroma.Token{})
			}
			if piece != "" {
				last := len(lines) - 1
				lines[last] = append(lines[last], chroma.Token{
					Type: token.Type, Value: piece})
			}
		}
	}
	return lines
}

// chromaRole is the markdown role a kind of token is coloured by in her
// glamour style, or empty for plain text.
func chromaRole(kind chroma.TokenType) string {
	switch {
	case kind.InCategory(chroma.Comment):
		return "md-comment"
	case kind == chroma.KeywordType, kind == chroma.NameBuiltin:
		return "md-type"
	case kind.InCategory(chroma.Keyword):
		return "md-keyword"
	case kind == chroma.NameFunction, kind == chroma.NameClass,
		kind == chroma.NameAttribute:
		return "md-function"
	case kind == chroma.NameConstant,
		kind.InSubCategory(chroma.LiteralNumber):
		return "md-number"
	case kind.InCategory(chroma.Literal):
		return "md-string"
	case kind.InCategory(chroma.Operator),
		kind.InCategory(chroma.Punctuation):
		return "md-operator"
	case kind == chroma.Name, kind == chroma.NameOther:
		return "md-name"
	}
	return ""
}

// vimGroup is the name in her vim scheme's table that a kind of token is
// drawn in, following the groups the scheme sets: Comment in dim; String,
// Number and Constant in string; Keyword and Statement in heading; Function
// and Special in link; Type in code; Operator in dim; the rest in ink.
func vimGroup(kind chroma.TokenType) string {
	switch {
	case kind.InCategory(chroma.Comment):
		return "dim"
	case kind == chroma.KeywordType:
		return "code"
	case kind == chroma.KeywordConstant, kind.InCategory(chroma.Literal):
		return "string"
	case kind.InCategory(chroma.Keyword):
		return "heading"
	case kind == chroma.NameFunction, kind == chroma.NameBuiltin:
		return "link"
	case kind.InCategory(chroma.Operator):
		return "dim"
	}
	return "ink"
}

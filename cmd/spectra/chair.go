package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/janearc/libtheme-css/spaces/ok"
	"github.com/janearc/libtheme-css/spaces/srgb"
)

// paintPane colours one tmux pane as a ghostty theme would, with every
// colour's chroma times k and its lightness times dim, hue held: the
// optometrist's chair, one candidate at a time, nothing written to any file.
//
// The pane's own text and ground come from the theme's foreground and
// background, the palette numbers from its palette lines, as paratune
// paints its own pane.
func paintPane(pane, theme string, k, dim float64) error {
	file, err := os.Open(theme)
	if err != nil {
		return err
	}
	defer file.Close()
	stronger := func(value string) string {
		polar := ok.FromSwatch(srgb.MustHex(value).Swatch()).Polar()
		polar.C *= k
		polar.L = min(polar.L*dim, 1)
		return hex(polar.Rect().Swatch())
	}
	style, palette := map[string]string{}, []string{}
	lines := bufio.NewScanner(file)
	for lines.Scan() {
		key, value, found := strings.Cut(lines.Text(), "=")
		if !found {
			continue
		}
		key, value = strings.TrimSpace(key), strings.TrimSpace(value)
		switch key {
		case "background", "foreground":
			style[key] = stronger(value)
		case "palette":
			number, colour, _ := strings.Cut(value, "=")
			name := fmt.Sprintf("pane-colours[%s]", number)
			palette = append(palette, ";", "set", "-p", "-t", pane,
				name, stronger(colour))
		}
	}
	look := "fg=" + style["foreground"] + ",bg=" + style["background"]
	args := []string{"set", "-p", "-t", pane, "window-style", look,
		";", "set", "-p", "-t", pane, "window-active-style", look}
	args = append(args, palette...)
	return exec.Command("tmux", args...).Run()
}

// unpaintPane puts a pane back on the terminal's own colours.
func unpaintPane(pane string) error {
	return exec.Command("tmux",
		"set", "-pu", "-t", pane, "window-style", ";",
		"set", "-pu", "-t", pane, "window-active-style", ";",
		"set", "-pu", "-t", pane, "pane-colours").Run()
}

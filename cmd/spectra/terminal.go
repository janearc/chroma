package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/janearc/libtheme-css/primitives/swatch"
	"github.com/janearc/libtheme-css/spaces/radiation"
	"github.com/janearc/libtheme-css/spaces/srgb"
)

// labelWidth is the room each line gives its name, and chipWidth the room
// for the light as one colour at the end of the line.
const (
	labelWidth = 15
	chipWidth  = 6
)

// row is one line of the pane: its name, its colours across the strip,
// empty where there is nothing to show, and the light as one colour.
type row struct {
	label   string
	colours []string
	chip    string
}

// show draws a set of rows in the pane it runs in, a line each: the
// bursts and the spectra, or neptune. It redraws when the pane is resized
// and stays until it is closed.
func show(rows func() []row) {
	draw(rows)
	resized := make(chan os.Signal, 1)
	signal.Notify(resized, syscall.SIGWINCH)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	for {
		select {
		case <-resized:
			draw(rows)
		case <-stop:
			fmt.Print("\x1b[?25h")
			return
		}
	}
}

// everything is the bursts and then the spectra.
func everything() []row {
	return append(burstRows(), planetRows()...)
}

// draw clears the pane and draws every row to its width.
func draw(rows func() []row) {
	width := paneWidth()
	stripWidth := max(width-labelWidth-chipWidth, 8)
	lines := []string{}
	for _, each := range rows() {
		line := padded(each.label) + cells(each.colours, stripWidth)
		if each.chip != "" {
			line += "  " + lamp(each.chip, chipWidth-2)
		}
		lines = append(lines, line)
	}
	fmt.Print("\x1b[?25l\x1b[H\x1b[2J" + strings.Join(lines, "\r\n"))
}

// burstRows are the bursts: through the whole fermi profile, where a
// burst's colour is its high-energy slope, so a steeper one is redder; then
// through the burst monitor, where the peak shows.
func burstRows() []row {
	chip := func(shape burst, profile radiation.Profile) string {
		return hex(atLightness(shape.colour(profile), 0.72))
	}
	steep := burst{
		Name: "a steep burst", Alpha: -1.0, Beta: -3.0, Peak: 200,
	}
	rows := []row{
		{"fermi", seenOf(bursts[0].strip(radiation.Fermi)),
			chip(bursts[0], radiation.Fermi)},
		{"fermi steep", seenOf(steep.strip(radiation.Fermi)),
			chip(steep, radiation.Fermi)},
	}
	for _, shape := range bursts {
		rows = append(rows, row{"gbm " + strings.Fields(shape.Name)[1],
			seenOf(shape.strip(gbm)), chip(shape, gbm)})
	}
	return rows
}

// planetRows are the spectra in data/: neptune's disc in daylight, the
// light neptune's clouds add and its dark spot takes, and
// pluto as new horizons' natural colour picture showed it, by region.
func planetRows() []row {
	day, rows := daylight(), []row{}
	add := func(label, file string, turn func(spectrum) spectrum) {
		curve, err := readSpectrum(dataPath(file))
		if err != nil || len(curve.nm) == 0 {
			rows = append(rows, row{label: label + " ?"})
			return
		}
		curve = turn(curve)
		chip := ""
		if curve.reaches(400, 700) {
			chip = curve.colour()
		}
		rows = append(rows, row{label, curve.strip(600), chip})
	}
	add("neptune", "neptune-disc.csv",
		func(c spectrum) spectrum { return c.times(day) })
	add("clouds add", "neptune-bright-clouds-sbs-difference.csv",
		func(c spectrum) spectrum { return c.scaledBy(1).times(day) })
	return append(rows, samplesRow("pluto 2015", "pluto-2015.css",
		"whale", "midlands", "heart", "pole"))
}

// samplesRow is regions sampled from a picture, which are colours and not
// a spectrum, side by side, so it has no chip.
func samplesRow(label, file string, names ...string) row {
	found, err := regions(dataPath(file))
	if err != nil {
		return row{label: label + " ?"}
	}
	colours := []string{}
	for _, name := range names {
		for range 24 {
			colours = append(colours, hex(found[name]))
		}
		colours = append(colours, "", "")
	}
	return row{label: label, colours: colours[:len(colours)-2]}
}

// paneWidth is the width of the tmux pane this runs in, or 80 outside one.
func paneWidth() int {
	output, err := exec.Command("tmux", "display", "-p", "-t",
		os.Getenv("TMUX_PANE"), "#{pane_width}").Output()
	if err != nil {
		return 80
	}
	width, err := strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil {
		return 80
	}
	return width
}

// padded is a label at the width every line gives it.
func padded(label string) string {
	gap := max(labelWidth-1-len([]rune(label)), 1)
	return " " + label + strings.Repeat(" ", gap)
}

// seenOf is a strip's colours as the eye sees them.
func seenOf(strip []column) []string {
	colours := make([]string, len(strip))
	for index, each := range strip {
		colours[index] = each.Seen
	}
	return colours
}

// cells draws colours across a number of cells, each cell the average, in
// light, of the colours that fall in it, or the nearest when there are
// fewer colours than cells; a cell with no colour in it is left blank.
func cells(colours []string, count int) string {
	var drawn strings.Builder
	for cell := 0; cell < count; cell++ {
		first := cell * len(colours) / count
		last := max((cell+1)*len(colours)/count, first+1)
		x, y, z, share := 0.0, 0.0, 0.0, 0.0
		from, to := min(first, len(colours)), min(last, len(colours))
		for _, each := range colours[from:to] {
			if each == "" {
				continue
			}
			cx, cy, cz := srgb.MustHex(each).Swatch().XYZ()
			x, y, z, share = x+cx, y+cy, z+cz, share+1
		}
		if share == 0 {
			drawn.WriteString(" ")
			continue
		}
		mix := swatch.FromXYZ(x/share, y/share, z/share)
		drawn.WriteString(lamp(hex(mix), 1))
	}
	return drawn.String()
}

// lamp is a run of cells lit in one colour, and the pane's own colours
// after it.
func lamp(colour string, count int) string {
	value := srgb.MustHex(colour)
	r, g, b := value.Bytes()
	cell := strings.Repeat(" ", count)
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm%s\x1b[0m", r, g, b, cell)
}

// webb carries webb's near infrared into visible the way the fermi profile
// carries gamma rays, evenly in the logarithm of wavelength: 1.4 microns
// lands at violet and 5.3 at the deep red.
var webb = radiation.Profile{Name: "webb",
	Source: radiation.Range{Lo: 1400, Hi: 5300}, Render: radiation.Visible}

// neptuneRows are neptune in daylight, then as webb saw it, carried into
// visible: nirspec's spectra of the disc, its high clouds and its clear
// half; and nircam's four numbers for the disc and the clouds.
func neptuneRows() []row {
	rows := []row{}
	curve, err := readSpectrum(dataPath("neptune-disc.csv"))
	if err != nil || len(curve.nm) == 0 {
		rows = append(rows, row{label: "neptune ?"})
	} else {
		curve = curve.times(daylight())
		rows = append(rows,
			row{"neptune", curve.strip(600), curve.colour()})
	}
	files := [][2]string{
		{"nirspec disc", "neptune-nirspec-disc.csv"},
		{"nirspec cloud", "neptune-nirspec-clouds.csv"},
		{"nirspec clear", "neptune-nirspec-clear.csv"},
		{"nircam disc", "neptune-nircam-disc.csv"},
		{"nircam cloud", "neptune-nircam-clouds.csv"},
	}
	for _, each := range files {
		rows = append(rows, webbRow(each[0], each[1]))
	}
	return rows
}

// webbRow is a spectrum in data carried into visible by the webb profile,
// with its light as one colour.
func webbRow(label, file string) row {
	curve, err := readSpectrum(dataPath(file))
	if err != nil || len(curve.nm) == 0 {
		return row{label: label + " ?"}
	}
	curve = curve.perLogarithm()
	chip := webb.Light(lightOf(webb, curve.inBand))
	return row{label, seenOf(spread(webb, curve.inBand)),
		hex(atLightness(chip, 0.72))}
}

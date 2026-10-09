// spectra turns light that libtheme can carry into colours a page can draw:
// gamma-ray bursts through a false colour profile, and the reflectance
// spectra in data/ lit by daylight. It writes page/data.js, which
// page/index.html draws.
//
// It finds cmd/spectra/page from anywhere inside chroma.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/janearc/chroma/internal/age"
	"github.com/janearc/libtheme-css/colourway"
	"github.com/janearc/libtheme-css/css"
	"github.com/janearc/libtheme-css/spaces/radiation"
	"github.com/janearc/libtheme-css/spaces/srgb"
)

// bursts are round values typical of the BATSE and Fermi catalogues: a long
// burst peaks near 200 keV, and a short one is harder, peaking higher with
// a shallower slope below the peak.
var bursts = []burst{
	{Name: "a long burst", Alpha: -1.0, Beta: -2.3, Peak: 200},
	{Name: "a short burst", Alpha: -0.5, Beta: -2.3, Peak: 600},
}

// gbm is Fermi's burst monitor alone, 8 keV to 40 MeV, the instrument that
// sees the peak of nearly every burst. The whole mission's profile reaches
// 300 GeV, where a burst has only its tail.
var gbm = radiation.Profile{
	Name: "fermi gbm",
	Source: radiation.Range{Lo: radiation.FromEnergy(40 * radiation.MeV),
		Hi: radiation.FromEnergy(8 * radiation.KeV)},
	Render: radiation.Visible,
}

// build and built are stamped by game build: the commit, and the
// commit's time. --age prints them.
var build, built = "dev", ""

// usage is what spectra prints for help, and for a verb it does not know
// or one given too few arguments.
const usage = `usage: spectra [VERB ARGS]

with no verb, spectra writes cmd/spectra/page/data.js, which
cmd/spectra/page/index.html draws: gamma-ray bursts through fermi's
profiles, and the reflectance spectra in data/ lit by daylight.

  show [neptune]                  draw the bursts and spectra, or neptune,
                                  in this pane until you close it
  paint PANE THEME [K [DIM]]      paint a tmux pane with THEME, its chroma
                                  times K and its lightness times DIM
  unpaint PANE                    put a pane back on the terminal's colours
  webb FOLDER                     neptune's spectra and colours from webb's
                                  nirspec cubes and nircam pictures
  voyager PICTURE OUT             neptune's colours sampled from PICTURE
  neptune-voyager PICTURE OUT     the neptune-voyager colourway
  pluto PICTURE OUT               pluto's colours sampled from PICTURE
  pluto-2015 PICTURE OUT          the same, with the boxes for 2015
  pluto-colourway PICTURE OUT     the pluto colourway
  fermi-long OUT                  the fermi-long colourway
  from-kept NAME OUT              a colourway from a line paratune kept,
                                  read from standard input
  tiles PICTURE [SIDE]            the most colourful squares of PICTURE
  --age                           which commit this binary is, and how old
`

// main runs the verb it is given, or, with none, writes page/data.js.
func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--age", "version":
			fmt.Println(age.Of("spectra", build, built, time.Now()))
			return
		case "help", "-h", "--help":
			fmt.Print(usage)
			return
		}
	}
	if len(os.Args) > 1 && os.Args[1] == "show" {
		if len(os.Args) > 2 && os.Args[2] == "neptune" {
			show(neptuneRows)
			return
		}
		show(everything)
		return
	}
	if len(os.Args) > 3 && os.Args[1] == "paint" {
		strength, dim := 1.0, 1.0
		if len(os.Args) > 4 {
			fmt.Sscan(os.Args[4], &strength)
		}
		if len(os.Args) > 5 {
			fmt.Sscan(os.Args[5], &dim)
		}
		err := paintPane(os.Args[2], os.Args[3], strength, dim)
		if err != nil {
			fmt.Fprintln(os.Stderr, "spectra:", err)
			os.Exit(1)
		}
		return
	}
	if len(os.Args) > 2 && os.Args[1] == "unpaint" {
		if err := unpaintPane(os.Args[2]); err != nil {
			fmt.Fprintln(os.Stderr, "spectra:", err)
			os.Exit(1)
		}
		return
	}
	if len(os.Args) > 2 && os.Args[1] == "webb" {
		err := fromNirspec(os.Args[2])
		if err == nil {
			err = fromNircam(os.Args[2])
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "spectra:", err)
			os.Exit(1)
		}
		return
	}
	if len(os.Args) > 3 && os.Args[1] == "voyager" {
		samples, err := sample(os.Args[2], [4]int{912, 851, 1246, 1064})
		if err == nil {
			err = sheetOf("neptune-voyager", voyagerSource,
				samples).Write(os.Args[3])
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "spectra:", err)
			os.Exit(1)
		}
		return
	}
	if len(os.Args) > 3 && os.Args[1] == "from-kept" {
		if err := fromKept(os.Args[2], os.Args[3]); err != nil {
			fmt.Fprintln(os.Stderr, "spectra:", err)
			os.Exit(1)
		}
		return
	}
	if len(os.Args) > 2 && os.Args[1] == "tiles" {
		side := 100
		if len(os.Args) > 3 {
			fmt.Sscan(os.Args[3], &side)
		}
		if err := tiles(os.Args[2], side, 0, 360, 12); err != nil {
			fmt.Fprintln(os.Stderr, "spectra:", err)
			os.Exit(1)
		}
		return
	}
	if len(os.Args) > 3 && os.Args[1] == "pluto" {
		samples, err := boxes(os.Args[2], plutoBoxes)
		if err == nil {
			err = sheetOf("pluto", plutoSource, samples).
				Write(os.Args[3])
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "spectra:", err)
			os.Exit(1)
		}
		return
	}
	if len(os.Args) > 3 && os.Args[1] == "pluto-2015" {
		samples, err := boxes(os.Args[2], pluto2015Boxes)
		if err == nil {
			err = sheetOf("pluto-2015", pluto2015Source, samples).
				Write(os.Args[3])
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "spectra:", err)
			os.Exit(1)
		}
		return
	}
	if len(os.Args) > 3 && os.Args[1] == "pluto-colourway" {
		source, err := pluto(os.Args[2])
		if err == nil {
			err = source.Write(os.Args[3])
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "spectra:", err)
			os.Exit(1)
		}
		return
	}
	if len(os.Args) > 3 && os.Args[1] == "neptune-voyager" {
		source, err := neptuneVoyager(os.Args[2])
		if err == nil {
			err = source.Write(os.Args[3])
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "spectra:", err)
			os.Exit(1)
		}
		return
	}
	if len(os.Args) > 2 && os.Args[1] == "fermi-long" {
		if err := fermiLong().Write(os.Args[2]); err != nil {
			fmt.Fprintln(os.Stderr, "spectra:", err)
			os.Exit(1)
		}
		return
	}
	if len(os.Args) > 1 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}
	drawn := []map[string]any{}
	for _, profile := range []radiation.Profile{radiation.Fermi, gbm} {
		drawn = append(drawn, through(profile))
	}
	data := map[string]any{"profiles": drawn}
	dir, err := pageDir()
	if err == nil {
		err = write(filepath.Join(dir, "data.js"), data)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "spectra:", err)
		os.Exit(1)
	}
	fmt.Println("spectra: wrote", filepath.Join(dir, "data.js"))
}

// through is every burst as one profile carries it into visible.
func through(profile radiation.Profile) map[string]any {
	drawn := []map[string]any{}
	for _, shape := range bursts {
		chip := hex(atLightness(shape.colour(profile), 0.72))
		drawn = append(drawn, map[string]any{
			"burst":     shape,
			"colour":    chip,
			"strip":     shape.strip(profile),
			"softening": shape.softening(profile, 32),
		})
	}
	return map[string]any{
		"name":   profile.Name,
		"kevLo":  radiation.Energy(profile.Source.Hi) / radiation.KeV,
		"kevHi":  radiation.Energy(profile.Source.Lo) / radiation.KeV,
		"nmLo":   profile.Render.Lo,
		"nmHi":   profile.Render.Hi,
		"bursts": drawn,
	}
}

// write puts the page's data where index.html reads it, as a script.
func write(path string, page map[string]any) error {
	text, err := json.Marshal(page)
	if err != nil {
		return err
	}
	script := append([]byte("window.SPECTRA = "), text...)
	return os.WriteFile(path, append(script, ";\n"...), 0o644)
}

// voyagerSource is where the voyager samples came from, written above them.
const voyagerSource = `
neptune-voyager -- neptune as voyager 2 showed it, sampled from one
picture and the picture not kept.
source: NASA/JPL, PIA01492 "Neptune Full Disk View", Voyager 2 narrow
angle camera, August 1989, 4.4 million miles out, 4 days 20 hours before
closest approach. https://science.nasa.gov/photojournal/neptune-full-disk-view
file: PIA01492.jpg, 2188 x 2185, 258997 bytes, fetched 2026-10-05.
colour: made from images through the green and orange filters only, so
these are two bands of light shown as three lamps, not a spectrum. each
region is its pixels averaged in linear light, written to the byte; the
disc is every pixel brighter than 0.01 luminance.`

// plutoBoxes are the regions sampled from new horizons' natural colour
// pluto, in pixels of the 8000 by 8000 picture: the heart's western lobe,
// the north polar cap, the pale midlands between, and the whale.
var plutoBoxes = []box{
	{"heart", 4300, 4000, 5300, 5400},
	{"pole", 3400, 900, 4800, 1500},
	{"midlands", 2000, 2400, 3400, 3400},
	{"whale", 1200, 5100, 2800, 6000},
}

// plutoSource is where the pluto samples came from, written above them.
const plutoSource = `
pluto -- pluto as new horizons showed it, sampled from one picture and
the picture not kept.
source: NASA/JHUAPL/SwRI/Alex Parker, "True Colors of Pluto", New Horizons
MVIC, 14 July 2015, 35,445 km out; recalibrated to natural colour, released
23 July 2018. https://science.nasa.gov/resource/true-colors-of-pluto/
file: BIG_P_COLOR_2_TRUE_COLOR1.png, 8000 x 8000, 57098975 bytes, fetched
2026-10-05.
colour: a single MVIC colour scan, processed to approximate what an eye
would see; not a spectrum. each region is its pixels averaged in linear
light, written to the byte; space, under 0.01 luminance, is left out.`

// pluto2015Boxes are the same regions in the 1024 by 1024 picture of
// pluto's big heart.
var pluto2015Boxes = []box{
	{"heart", 427, 512, 569, 668},
	{"pole", 427, 228, 626, 313},
	{"midlands", 284, 370, 427, 484},
	{"whale", 256, 626, 370, 711},
}

// pluto2015Source is where the 2015 pluto samples came from.
const pluto2015Source = `
pluto-2015 -- pluto as new horizons first showed it, the day before
closest approach: the picture everyone knows, sampled and not kept.
source: NASA/JHUAPL/SwRI, PIA19708 "Pluto's Big Heart in Color", LORRI,
13 July 2015, 768,000 km out, coloured with lower resolution colour from
Ralph taken earlier that day.
https://science.nasa.gov/photojournal/plutos-big-heart-in-color
file: PIA19708.tif, 1024 x 1024, 3146976 bytes, fetched 2026-10-05.
colour: black and white detail coloured from another camera's coarser
colour, processed for release before the 2018 recalibration; not a
spectrum. each region is its pixels averaged in linear light, written to
the byte; space, under 0.01 luminance, is left out.`

// fromKept is a colourway made from a line paratune kept, read from the
// standard input: its roles, under a new name, written as a save writes
// a source, so paratune opens it as its own.
func fromKept(name, path string) error {
	line, err := io.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	fields := strings.Fields(string(line))
	if len(fields) < 4 {
		return fmt.Errorf("a kept line is a date, a time, " +
			"a colourway and its roles")
	}
	roles := css.New()
	for _, pair := range fields[3:] {
		role, hex, found := strings.Cut(pair, "=")
		if !found {
			return fmt.Errorf("%q is not a role and a colour", pair)
		}
		value, err := srgb.FromHex(hex)
		if err != nil {
			return err
		}
		roles.Set(role, value.Swatch())
	}
	about := name + " -- made in paratune from " + fields[2] + "."
	return colourway.Source{Name: name, About: about, Roles: roles,
		Origin: "sources/" + name + ".css"}.Write(path)
}

// pageDir is the folder that holds page/index.html: page/ here when
// spectra runs in cmd/spectra, or cmd/spectra/page in this folder or the
// nearest one above it.
func pageDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	if holdsPage(filepath.Join(dir, "page")) {
		return filepath.Join(dir, "page"), nil
	}
	for {
		page := filepath.Join(dir, "cmd", "spectra", "page")
		if holdsPage(page) {
			return page, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("no cmd/spectra/page here or " +
				"above; run spectra from inside chroma")
		}
		dir = parent
	}
}

// holdsPage is whether a folder has the page's index.html in it.
func holdsPage(dir string) bool {
	_, err := os.Stat(filepath.Join(dir, "index.html"))
	return err == nil
}

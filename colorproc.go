package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	convert := flags.Bool("convert", false, "Convert a #RGB value to HSL")

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if *convert {
		rgb := flags.Arg(0)
		if len(rgb) == 0 {
			return errors.New("missing #RGB argument")
		}

		if rgb[0] == '#' {
			rgb = rgb[1:]
		}

		if len(rgb) != 6 {
			return errors.New("invalid #RGB argument: " + rgb)
		}

		r, err := strconv.ParseUint(rgb[0:2], 16, 8)
		if err != nil {
			return err
		}
		g, err := strconv.ParseUint(rgb[2:4], 16, 8)
		if err != nil {
			return err
		}
		b, err := strconv.ParseUint(rgb[4:], 16, 8)
		if err != nil {
			return err
		}

		h, s, l := RgbToHsl(int(r), int(g), int(b))

		fmt.Fprintf(stdout, "#%s\n", rgb)
		fmt.Fprintf(stdout, "RGB: (%v, %v, %v)\n", r, g, b)
		fmt.Fprintf(stdout, "HSL: (%v, %v%%, %v%%)\n", h, s, l)

		return nil
	}

	return errors.New("missing parameters")
}

func toDeg(hue float64) int {
	r := hue * 60.0
	if r < 0 {
		r += 360.0
	}
	return int(math.Round(r))
}

func toPercent(v float64) int {
	return int(math.Round(v * 100.0))
}

func RgbToHsl(r, g, b int) (int, int, int) {
	vmax := max(r, g, b)
	vmin := min(r, g, b)

	// normalize
	nr := float64(r) / 255.0
	ng := float64(g) / 255.0
	nb := float64(b) / 255.0
	nvmax := float64(vmax) / 255.0
	nvmin := float64(vmin) / 255.0

	h := 0.0
	s := 0.0
	l := (nvmax + nvmin) / 2.0

	if vmax != vmin {
		c := nvmax - nvmin
		s = c / (1.0 - math.Abs(2.0*l-1.0))

		switch vmax {
		case r:
			h = (ng - nb) / c
			h = math.Mod(h, 6.0)
		case g:
			h = (nb-nr)/c + 2.0
		case b:
			h = (nr-ng)/c + 4.0
		}
	}

	return toDeg(h), toPercent(s), toPercent(l)
}

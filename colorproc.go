package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	if err := Run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func Run(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)

	convert := flags.Bool("convert", false, "Convert a #RGB value to HSL")

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if *convert {
		hex := Hex(flags.Arg(0))
		if len(hex) == 0 {
			return errors.New("missing #RGB argument")
		}

		rgb, err := hex.ToRgb()
		if err != nil {
			return err
		}

		hsl := rgb.ToHsl()

		fmt.Println(hex)
		fmt.Println(rgb)
		fmt.Println(hsl)

		rgb2 := hsl.ToRgb()
		hex2 := rgb2.ToHex()

		fmt.Println(rgb2)
		fmt.Println(hex2)

		hsl2 := rgb2.ToHsl()
		hex3 := hsl2.ToHex()
		fmt.Println(hsl2)
		fmt.Println(hex3)

		return nil
	}

	return errors.New("missing parameters")
}

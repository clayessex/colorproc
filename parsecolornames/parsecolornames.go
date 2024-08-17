package main

import (
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

const (
	ListFilename   = "colornames.csv.gz"
	OutputPath     = "../tmp/colornames/"
	OutputFilename = "colornames.go"
)

type Color struct {
	name string
	rgb  string
}

func main() {
	f, err := os.Open(ListFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	z, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer z.Close()

	colorNames := make([]Color, 0, 30000)

	scanner := bufio.NewScanner(z)
	first := true
	for scanner.Scan() {
		if first {
			first = false
			continue
		}

		fields := strings.Split(scanner.Text(), ",")
		if len(fields) < 2 {
			log.Printf("Line is not valid CSV: [%s]", scanner.Text())
			continue
		}

		name := fields[0]
		rgb := fields[1]

		colorNames = append(colorNames, Color{name, rgb})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	writeColors(OutputPath, OutputFilename, colorNames)
}

// Create a Go code file containing the colors slice
func writeColors(filepath string, filename string, colors []Color) {
	if _, err := os.Stat(filepath); errors.Is(err, fs.ErrNotExist) {
		os.Mkdir(filepath, 0775)
	}

	fout, err := os.Create(filepath + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	fout.WriteString(`
package colornames

var List = []struct{
    Name string
    Rgb string
}{
`)

	for _, v := range colors {
		fmt.Fprintf(fout, `    {Name: "%s", Rgb: "%v"},`, v.name, v.rgb)
		fmt.Fprintf(fout, "\n")
	}

	fout.WriteString("}\n")
}

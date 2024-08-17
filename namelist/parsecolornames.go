package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ListFilename   = "colornames.csv.gz"
	OutputFilename = "colornames.go.out"
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

	writeColors(OutputFilename, colorNames)
}

// Create a Go code file containing the colors slice
func writeColors(filename string, colors []Color) {
	fout, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	fout.WriteString(`
package main

var ColorNames = []struct{
    name string
    rgb string
}{
`)

	for _, v := range colors {
		fmt.Fprintf(fout, `    {name: "%s", rgb: "%v"},`, v.name, v.rgb)
		fmt.Fprintf(fout, "\n")
	}

	fout.WriteString("}\n")
}

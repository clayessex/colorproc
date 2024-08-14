package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/clayessex/colorproc/colors"
)

const (
	ListFilename   = "colornames.csv.gz"
	OutputFilename = "colornames.go.out"
)

func main() {
	fout, err := os.Create(OutputFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	fout.WriteString(`
const ColorNames = []struct{
    hsl Hsl
    hex Hex
    name string
}{
`)

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

	scanner := bufio.NewScanner(z)
	maxLines := 20
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
		hex := colors.Hex(fields[1])
		hsl, err := hex.ToHsl()
		if err != nil {
			log.Printf("Line color is not parsable: [%s]", scanner.Text())
		}
		hslstring := hsl.F3String()

		// TODO: create printers for go Hsl and Hex types

		fmt.Fprintf(fout, "\thsl: %v, hex: %v, name: %s\n", hslstring, hex, name)

		if maxLines--; maxLines == 0 {
			log.Print("reached max lines")
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fout.WriteString("}\n")
}

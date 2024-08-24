package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/clayessex/colorproc/build/colornamesorted"
)

const (
	OutputFilename = "color-display.html"
	OutputPath     = "../build/display/"
)

func main() {
	WriteColorList(OutputPath, OutputFilename)
}

func WriteColorList(filepath, filename string) {
	if _, err := os.Stat(filepath); errors.Is(err, fs.ErrNotExist) {
		os.Mkdir(filepath, 0775)
	}

	f, err := os.Create(filepath + filename)
	if err != nil {
		log.Fatal("unable to create: ", filepath+filename)
	}
	defer f.Close()

	f.WriteString(Head)
	for _, v := range colornamesorted.List {
		fmt.Fprintf(f, Swatch, v.Rgb, v.Name, v.Rgb)
	}
	f.WriteString(Tail)
}

const Swatch = `
	<div class="swatch-block">
		<div class="swatch" style="background-color: %s">&nbsp; </div>
		<div class="color-name">%s</div>
		<div>%s</div>
	</div>`

const Head = `<html>
<body>
<head>
<style>
body {
	background-color: #282828;
}
main {
  display: flex;
  flex-wrap: wrap;
  color: #b0b0b0;
  font-size: small;
}
.swatch {
	border-radius: 0.5rem;
	padding: 1rem;
	margin: 1rem;
	height: 48px;
	width: 48px;
}
.color-name {
	font-size: x-small;
}
.swatch-block {
	text-align: center;
}
</style>
</head>

<main>
`

const Tail = `
</main>

</body>
</html>
`

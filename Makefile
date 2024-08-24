

SOURCE_URL = https://github.com/meodai/color-names/raw/master/src/colornames.csv


all: build/display/color-display.html


## Download the list of color names
build/csv/colornames.csv.gz:
	mkdir build/csv
	wget -O build/csv/colornames.csv ${SOURCE_URL}
	gzip build/csv/colornames.csv

## Parse the list and generate a go source file from the color names
build/colornames/colornames.go: build/csv/colornames.csv.gz
	cd parsecolornames && go run .

## Sort the color list according to a simple Hue clustering algorithm
build/colornamesorted/colornamesorted.go: build/colornames/colornames.go
	cd colorsort && go run .

build/display/color-display.html: build/colornamesorted/colornamesorted.go
	cd display && go run .

clean:
	# rm -rf build/csv
	rm -rf build/colornames
	rm -rf build/colornamesorted


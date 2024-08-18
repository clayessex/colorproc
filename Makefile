

all: build/colornames/colornames.go


SOURCE_URL = https://github.com/meodai/color-names/raw/master/src/colornames.csv


build/csv/colornames.csv.gz:
	mkdir build/csv
	wget -O build/csv/colornames.csv ${SOURCE_URL}
	gzip build/csv/colornames.csv

build/colornames/colornames.go: build/csv/colornames.csv.gz
	cd parsecolornames && go run .


clean:
	# rm -rf build/csv
	rm -rf build/colornames
	rm -rf build/colornamesorted


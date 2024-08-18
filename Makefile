

all: build/csv/colornames.csv.gz


SOURCE_URL = https://github.com/meodai/color-names/raw/master/src/colornames.csv


build/csv/colornames.csv.gz:
	mkdir build/csv
	wget -O build/csv/colornames.csv ${SOURCE_URL}
	gzip build/csv/colornames.csv

colornames:

clean:
	rm -rf build/csv

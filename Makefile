all: build
.PHONY: all

build:
	go build -o _output/bin/sync-vendor github.com/deads2k/sync-vendor
.PHONY: build

clean:
	rm -rf _output
.PHONY: build

.PHONY: build test clean

VERSION := $(shell git describe --tags --always --dirty)
BUILDTIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
GITCOMMIT := $(shell git rev-parse HEAD)
LDFLAGS := -X "qr_generator/cmd.Version=$(VERSION)" \
           -X "qr_generator/cmd.BuildTime=$(BUILDTIME)" \
           -X "qr_generator/cmd.GitCommit=$(GITCOMMIT)"

build:
	go build -ldflags "$(LDFLAGS)" -o bin/qr

test:
	go test -v ./...

clean:
	rm -rf bin/ 

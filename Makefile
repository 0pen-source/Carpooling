#! /usr/bin/make
#
# Targets:
# - "depend" retrieves the Go packages needed to run the linter and tests
# - "format" runs goimports on all go files
# - "lint" runs the linters and checks the code format using gometalinter
# - "test" runs the tests
# - "build" build the executable file
#
# Meta targets:
# - "all" is the default target, it runs all the targets in the order above.
#
DIRS=$(shell go list -f {{.Dir}} ./... | grep -v /vendor/)
DEPEND=\
			 golang.org/x/tools/cmd/cover \
			 github.com/alecthomas/gometalinter \
			 github.com/tools/godep

all: depend format lint test build

# install/update all dependency
depend:
	@go get -v $(DEPEND)
	@gometalinter -i

format:
	@for d in $(DIRS) ; do \
		if [ "`goimports -w -srcdir $$d $$d/*.go | tee /dev/stderr`" ]; then \
			echo "^ - Failed to format go files in directory $$d" && echo && exit 1; \
		fi \
	done

lint:
	@for d in $(DIRS) ; do \
		if [ "`goimports -l $$d/*.go | tee /dev/stderr`" ]; then \
			echo "^ - Repo contains improperly formatted go files" && echo && exit 1; \
		fi \
	done
	@gometalinter --vendor \
		--deadline=60s \
		--enable=golint \
		--enable=vetshadow \
	    --enable=gocyclo \
        --enable=gosimple \
        --enable=staticcheck \
		--enable=deadcode \
		--enable=goconst \
		--enable=dupl \
		--linter='dupl:dupl -plumbing -threshold {duplthreshold} ./*.go | grep -v "_test.go"::(?P<path>[^\s][^:]+?\.go):(?P<line>\d+)-\d+:\s*(?P<message>.*)'  \
		./...

#test:
#	@bash test.sh

build:
	@go build -v

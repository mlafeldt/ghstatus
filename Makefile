GO = go

all: build

get:
	$(GO) get -d -x ./...

build: get
	$(GO) build -x ./...

install: get
	$(GO) install -x ./...

test_dep: get
	$(GO) get -x github.com/stretchr/testify
	$(GO) test -i -x ./...

test: test_dep
	$(GO) test -v -x ./...

clean:
	$(GO) clean -x ./...

.PHONY: all get build install test_dep test clean

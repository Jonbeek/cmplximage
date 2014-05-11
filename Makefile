gopath = GOPATH=$(CURDIR)
packages = cmplximage examples

all: build

build: fmt
	$(gopath) go build $(packages)

fmt:
	$(gopath) go fmt $(packages)

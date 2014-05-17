gopath = GOPATH=$(CURDIR)
packages = cmplximage examples/...

all: install

install: fmt
	$(gopath) go install $(packages)

fmt:
	$(gopath) go fmt $(packages)

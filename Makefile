REPO := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all:
	GOPATH=$(REPO) go build -o i2proxy

clean:
	GOPATH=$(REPO) go clean
	rm -f inproxy

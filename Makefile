GOPATH=$(PWD)

all:
	go build inproxy
	go build configure

clean:
	go clean
	rm -f inproxy
	rm -f configure

.PHONY: test clean deps lint

PROG=gophorg

all: test ${PROG} deps # lint  <- Add back when Issue 15 is resolved.

deps:
	go get .

${PROG}:
	go build

test:
	go test

lint:
	golint -set_exit_status .
	staticcheck .

clean:
	rm ${PROG}

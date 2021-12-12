.PHONY: test clean deps

PROG=gophorg

all: test ${PROG} deps

deps:
	go get .

${PROG}:
	go build

test:
	go test

clean:
	rm ${PROG}

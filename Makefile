.PHONY: test clean deps lint

PROG=gophorg

all: test ${PROG} deps lint

deps:
	go get .

${PROG}:
	go build

test:
	go test

lint:
	golint -set_exit_status .

clean:
	rm ${PROG}

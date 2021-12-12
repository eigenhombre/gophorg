.PHONY: all test clean

PROG=gophorg

all:
	make test
	make ${PROG}

${PROG}:
	go build

test:
	go test

clean:
	rm ${PROG}

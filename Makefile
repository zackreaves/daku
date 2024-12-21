BINARY_NAME=daku

.PHONY: all
all: main

deps:
	go get github.com/mattn/go-sqlite3 

main:
	go build -o ${BINARY_NAME} ./

clean:
	go clean

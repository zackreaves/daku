BINARY_NAME=daku

.PHONY: all
all: main deps

deps:
	go get github.com/mattn/go-sqlite3 
	go get github.com/lib/pq

main:
	go build -o ${BINARY_NAME} ./

clean:
	go clean

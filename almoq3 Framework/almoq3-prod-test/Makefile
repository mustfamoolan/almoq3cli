.PHONY: build run clean

BINARY_NAME=almoq3-prod-test

build:
	go build -o bin/$(BINARY_NAME) main.go

run:
	go run main.go

clean:
	rm -rf bin/
	go clean

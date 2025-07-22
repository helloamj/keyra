BINARY_NAME = keyra

.PHONY: all build run clean

all: build

build:
	go build -o $(BINARY_NAME) ./cmd/...

run: build
	./$(BINARY_NAME)

clean:
	go clean
	rm -f $(BINARY_NAME)

NAME := lastpost

all: test lint

build: dependencies
	go build -o bin/$(NAME)

test: dependencies
	go test ./...

lint: dependencies
	@which -s gometalinter || (go get github.com/alecthomas/gometalinter && gometalinter --install)
	gometalinter --deadline 1m ./...

dependencies:
	go get -t

.PHONY: build dependencies lint test

NAME := lastpost

GOX_OSARCH := darwin/386 darwin/amd64 linux/amd64 windows/386 windows/amd64

default: test lint

test: dependencies
	go test ./...

lint: dependencies
	@which -s gometalinter || (go get github.com/alecthomas/gometalinter && gometalinter --install)
	gometalinter --deadline 1m ./...

dependencies:
	go get -t

release: dependencies clean
	go get github.com/mitchellh/gox
	gox \
	  -osarch="$(GOX_OSARCH)" \
	  -output="release/$(NAME)_{{.OS}}_{{.Arch}}" \
	  -osarch="darwin/amd64" \
	  ./...
	cd release/; for f in *; do mv -v $$f $(NAME); tar -zcf $$f.tar.gz $(NAME); rm $(NAME); done

clean:
	rm -rf release/

.PHONY: default clean dependencies lint release test

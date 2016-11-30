.PHONY: generate
generate:
	go generate ./...

.PHONY: build
build:
	go build ./...
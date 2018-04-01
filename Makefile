BINARY ?= go-hello-server

.PHONY: setup
setup:
	glide install --strip-vendor

.PHONY: build_linux
build_linux: setup
	env GOOS=linux GOARCH=amd64 go build -o $(BINARY) ./cmd/server/hello_server.go

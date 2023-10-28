buf = go run github.com/bufbuild/buf/cmd/buf@latest
gobin = $(shell go env GOPATH)/bin

all: yaml proto golang

$(gobin)/buf:
	go install github.com/bufbuild/buf/cmd/buf@latest

$(gobin)/gofumpt:
	go install mvdan.cc/gofumpt@latest

$(gobin)/golangci-lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

$(gobin)/protoc-gen-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

$(gobin)/yamlfmt:
	go install github.com/google/yamlfmt/cmd/yamlfmt@latest

proto: $(gobin)/buf $(gobin)/protoc-gen-go
	rm -rf pkg/apis
	buf format -w
	buf generate --template build/buf.gen.yaml

golang: $(gobin)/gofumpt $(gobin)/golangci-lint
	gofumpt -w .
	go test ./...
	go mod tidy
	golangci-lint run

yaml: $(gobin)/yamlfmt
	yamlfmt .

buf = go run github.com/bufbuild/buf/cmd/buf@latest
gobin = $(shell go env GOPATH)/bin

all: yaml golang

$(gobin)/buf:
	go install github.com/bufbuild/buf/cmd/buf@latest

$(gobin)/gofumpt:
	go install mvdan.cc/gofumpt@latest

$(gobin)/golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

$(gobin)/protoc-gen-connect-go:
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

$(gobin)/protoc-gen-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

$(gobin)/yamlfmt:
	go install github.com/google/yamlfmt/cmd/yamlfmt@latest

proto: $(gobin)/buf $(gobin)/protoc-gen-connect-go $(gobin)/protoc-gen-go
	rm -rf pkg/apis
	buf format -w
	buf generate --template build/buf.gen.yaml
	go install ./cmd/protoc-gen-tryout
	buf generate --template build/buf.gen.tryout.yaml

golang: $(gobin)/gofumpt $(gobin)/golangci-lint proto
	gofumpt -w .
	go test ./...
	go mod tidy
	golangci-lint run

yaml: $(gobin)/yamlfmt
	yamlfmt .

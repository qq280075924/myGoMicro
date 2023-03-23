GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto@v1.30.0
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@v1.0.0

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/myGoMicro.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o myGoMicro *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t myGoMicro:latest .
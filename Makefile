
PWD = $(shell pwd)
proto_dir=./protos
pb_dir=./proto

OUTPUT_DIR := $(PWD)/bin
TARGET := validator

COMMIT=$(shell git rev-parse --short HEAD)
CURVER=$(shell git rev-parse --abbrev-ref HEAD | awk -F'/' '{print $$2}')
VERSION=$(CURVER)_$(COMMIT)

BUILD_TIME=$(shell date "+%Y-%m-%d %H:%M:%S")
GIT_COMMIT_ID=$(shell git rev-parse HEAD)



proto: clean
	@echo "=============> proto"
	protoc --proto_path=${proto_dir} --go_out=paths=source_relative:${pb_dir} ${proto_dir}/*.proto
	protoc-go-inject-tag -input="./${pb_dir}/*.pb.go"

clean:
	@echo "=============> clean"
	@rm -vf ${pb_dir}/*.pb.go

run:
	go run ./main.go

.PHONY:build
build:
	@echo "VERSION: $(VERSION), BUILD_TIME: $(BUILD_TIME), GIT_COMMIT_ID: $(GIT_COMMIT_ID)"
  		GO111MODULE=on GOOS=linux go build -o validator main.go \
  		-a -ldflags "-X 'main.Version=$(VERSION)'-X 'main.BuildTime=$(BUILD_TIME)' -X 'main.GitCommitID=$(GIT_COMMIT_ID)'" \
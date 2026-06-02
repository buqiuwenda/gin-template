.PHONY: build run wire proto tidy

APP_NAME := gin-template
CONFIG   ?= configs/config.example.yaml

build:
	go build -o bin/$(APP_NAME) ./cmd

run: build
	./bin/$(APP_NAME) web -c $(CONFIG)

wire:
	@bash scripts/gen_wire.sh

proto:
	@bash scripts/gen_proto.sh

tidy:
	go mod tidy

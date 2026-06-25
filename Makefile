.PHONY: build run wire proto tidy gen-config gen-doc gen-db

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


gen-config:
	protoc --go_out=. --go_opt=module=github.com/buqiuwenda/gin-template --go-grpc_out=. --go-grpc_opt=module=github.com/buqiuwenda/gin-template internal/conf/conf.proto

# 生成数据库模型
gen-db:
	@reverse -f reverse.yaml

# 生成 swagger 文档（依赖 third_party/googleapis，由 make proto 拉取）
GOOGLEAPIS := third_party/googleapis

gen-doc:
	@test -f $(GOOGLEAPIS)/google/api/annotations.proto || $(MAKE) proto
	protoc --proto_path=$(GOOGLEAPIS) \
		--proto_path=. \
		--openapiv2_out=./docs \
		--openapiv2_opt=output_format=json \
		api/v1/**/*.proto
  

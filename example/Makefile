NAME?=example-server
VERSION?=0.0.1
LICENSE?=false
PACKAGE=server

# 检查操作系统
ifeq ($(OS),Windows_NT)
    PACKAGE := $(PACKAGE).exe
endif

clean:
	@if [ -d "./$(NAME)" ]; then rm -rf "./$(NAME)"; fi
	@rm -f *.tar.gz

wire:
	@rm -f "./cmd/server/wire_gen.go"
	@wire gen "./cmd/server"
	@echo "Wire Gen Successful!"
pb:
	@echo "start generate protobuf..."
	@rm -f "./app/protobuf/pb/*"
	@cd "app/protobuf" && protoc -I=./proto/ --go_out=. proto/*.proto
	@protoc-go-inject-tag -input="./app/protobuf/pb/*.pb.go"
	@cd "app/protobuf" && protoc -I=./proto --go-grpc_out=. proto/*.proto
	@git add "./app/protobuf/pb/*"
	@echo "protobuf generate successful!"

build:clean
	@echo "Start build ${PACKAGE}..."
	@mkdir -p "$(NAME)/resource"
	@cp -r "./resource" "$(NAME)"
	@go build -ldflags "-X 'github.com/lrayt/small-sparrow/core.AppNAME=$(NAME)' -X 'github.com/lrayt/small-sparrow/core.VERSION=$(VERSION)' -X 'github.com/lrayt/small-sparrow/core.VerifyLICENSE=$(LICENSE)'" -a -v -o "$(NAME)/${PACKAGE}" "./cmd/server"
	@tar zcvf "$(NAME)-v$(VERSION).tar.gz" "$(NAME)"
	@rm -rf "./$(NAME)"
	@echo "Packaging successful!"
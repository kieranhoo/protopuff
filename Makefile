GENERATED_DIR=pkg/gen
PROTO_DIR=internal/proto

b:
	go build -ldflags="-s -w" -o ./bin/exe ./cmd

gen:
	@if [ -d "$(GENERATED_DIR)" ]; then \
		rm -rf $(GENERATED_DIR); \
	fi
	mkdir -p $(GENERATED_DIR)

	protoc --proto_path=$(PROTO_DIR) --go_out=$(GENERATED_DIR) \
    --go-grpc_out=$(GENERATED_DIR) \
	--grpc-gateway_out $(GENERATED_DIR) \
    $(PROTO_DIR)/*.proto

setup:
	go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

s:
	go run cmd/main.go s

w:
	go run cmd/main.go w

m:
	go run cmd/main.go m

client:
	go run cmd/client/client.go

sqlc-setup:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
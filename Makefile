GENERATED_DIR=internal/gen
PROTO_DIR=proto/service

gen:
	@if [ -d "$(GENERATED_DIR)" ]; then \
		rm -rf $(GENERATED_DIR); \
	fi
	mkdir -p $(GENERATED_DIR)

	protoc --proto_path=proto --go_out=$(GENERATED_DIR) \
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

client:
	go run cmd/client/client.go
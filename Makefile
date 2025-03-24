PROTO_DIR = api
OUT_DIR = internal/pb

# PROTOC_GEN_GO = protoc-gen-go
# PROTOC_GEN_GO_GRPC = protoc-gen-go-grpc
# PROTOC_GEN_GRPC_GATEWAY = protoc-gen-grpc-gateway

# all: clean generate

# Xóa file cũ trước khi generate
# clean:
# 	rm -rf $(OUT_DIR)/*.pb.go $(OUT_DIR)/*.pb.gw.go $(OUT_DIR)/*.pb_grpc.go

# Generate pb.go, pb_grpc.go, pb.gw.go
generate:
	protoc -I=. \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		--grpc-gateway_out=$(OUT_DIR) \
		--grpc-gateway_opt generate_unbound_methods=true \
		--openapiv2_out . \
		--openapiv2_opt generate_unbound_methods=true \
		$(PROTO_DIR)/*.proto




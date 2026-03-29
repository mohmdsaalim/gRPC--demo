# ══════════════════════════════════════════════════════════════
# Makefile — run these commands to build and run the project
# ══════════════════════════════════════════════════════════════

# Generate Go code from .proto files
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/user.proto proto/order.proto

# Run User Service (Terminal 1)
user:
	go run user-service/main.go

# Run Order Service (Terminal 2)
order:
	go run order-service/main.go

# Run Client (Terminal 3)
client:
	go run client/main.go

# Install protoc plugins (run once)
install-tools:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: proto user order client install-tools

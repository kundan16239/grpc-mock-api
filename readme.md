# Go gRPC User Service

This project implements a simple gRPC-based User Service in Go with basic CRUD operations.  
It includes `proto` definitions, server implementation, and repository + service layers.

---

## 🚀 Prerequisites

Before running, ensure you have:

- Go (>= 1.20) installed → [Download](https://go.dev/dl/)
- Protocol Buffers Compiler (`protoc`) → [Install Guide](https://grpc.io/docs/protoc-installation/)
- Go plugins for `protoc`:
  ```sh
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  export PATH="$PATH:$(go env GOPATH)/bin"


# Run this after writing proto file to generate
protoc \                        
  --proto_path=api/grpc/proto \
  --go_out=proto_gen --go_opt=paths=source_relative \
  --go-grpc_out=proto_gen --go-grpc_opt=paths=source_relative \
  api/grpc/proto/user.proto


# Run Grpc Server
go run ./cmd/grpc_server/main.go

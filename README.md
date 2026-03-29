# gRPC Demo — Go Microservices (User + Order)

A minimal gRPC project in Go with **2 services** to understand how gRPC works.

---

## 📁 Project Structure

```
grpc-demo/
├── proto/                  ← .proto files (the CONTRACT)
│   ├── user.proto
│   ├── order.proto
│   ├── userpb/             ← Auto-generated Go code (after protoc)
│   │   ├── user.pb.go
│   │   └── user_grpc.pb.go
│   └── orderpb/
│       ├── order.pb.go
│       └── order_grpc.pb.go
├── user-service/           ← gRPC Server #1 (port 50051)
│   └── main.go
├── order-service/          ← gRPC Server #2 (port 50052)
│   └── main.go             ← Also a gRPC CLIENT to User Service!
├── client/                 ← Test client that calls both services
│   └── main.go
├── go.mod
└── Makefile
```

---

## 🚀 Setup & Run (Step by Step)

### Prerequisites

```bash
# 1. Install Go (1.21+)
#    https://go.dev/dl/

# 2. Install Protocol Buffer compiler
#    macOS:   brew install protobuf
#    Ubuntu:  sudo apt install protobuf-compiler
#    Windows: Download from https://github.com/protocolbuffers/protobuf/releases

# 3. Install Go plugins for protoc
make install-tools
# OR manually:
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 4. Make sure $GOPATH/bin is in your PATH
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Generate Code & Run

```bash
# Step 1: Generate Go code from .proto files
make proto

# Step 2: Download dependencies
go mod tidy

# Step 3: Open 3 terminals and run:

# Terminal 1 — Start User Service
make user
# Output: 🚀 User Service running on :50051

# Terminal 2 — Start Order Service
make order
# Output: 🔗 Connected to User Service on :50051
# Output: 🚀 Order Service running on :50052

# Terminal 3 — Run the client
make client
```

---

## 🧠 gRPC Core Concepts

### What is gRPC?

gRPC = **g**oogle **R**emote **P**rocedure **C**all

Instead of sending JSON over HTTP like REST, gRPC uses:
- **Protocol Buffers (protobuf)** for data serialization (binary, smaller, faster)
- **HTTP/2** for transport (multiplexing, streaming, header compression)
- **Strongly typed contracts** (.proto files) shared between client & server

### The 5 Main Components

```
┌─────────────────────────────────────────────────────┐
│  1. PROTO FILE (.proto)                             │
│     → Defines messages (data) and services (RPCs)   │
│     → The "contract" between client and server      │
└────────────────────┬────────────────────────────────┘
                     │ protoc (compiler)
                     ▼
┌─────────────────────────────────────────────────────┐
│  2. GENERATED CODE (pb.go + grpc.pb.go)             │
│     → pb.go: message structs + serialization        │
│     → grpc.pb.go: client stub + server interface    │
└────────┬──────────────────────────────┬─────────────┘
         ▼                              ▼
┌──────────────────┐          ┌──────────────────────┐
│  3. SERVER       │          │  4. CLIENT           │
│  Implements the  │◄─────────│  Uses generated      │
│  service inter-  │  gRPC    │  client stub to      │
│  face from proto │  call    │  call server methods  │
└──────────────────┘          └──────────────────────┘
         │
         ▼
┌──────────────────────────────────────────────────────┐
│  5. gRPC SERVER RUNTIME                              │
│  → grpc.NewServer() creates the server               │
│  → RegisterXxxServer() binds your implementation     │
│  → server.Serve(listener) starts accepting calls     │
└──────────────────────────────────────────────────────┘
```

### 4 Types of RPC

```
TYPE                    DESCRIPTION                         USED IN THIS PROJECT?
─────────────────────────────────────────────────────────────────────────────────
Unary                   1 request → 1 response              ✅ GetUser, CreateOrder
Server Streaming        1 request → stream of responses     ✅ ListUsers, GetOrdersByUser
Client Streaming        stream of requests → 1 response     ❌ (not in this demo)
Bidirectional Streaming stream ↔ stream                     ❌ (not in this demo)
```

### gRPC vs REST

```
FEATURE          gRPC                        REST
──────────────────────────────────────────────────────
Protocol         HTTP/2                      HTTP/1.1
Data Format      Protobuf (binary)           JSON (text)
Contract         .proto file (strict)        OpenAPI/none (loose)
Streaming        Built-in (4 types)          Not native
Code Gen         Automatic                   Manual/tools
Speed            Very fast                   Slower
Browser Support  Needs gRPC-Web proxy        Native
Best For         Microservices, internal     Public APIs, web
```

### Error Handling

gRPC has its OWN error codes (not HTTP status codes):

```go
import "google.golang.org/grpc/codes"
import "google.golang.org/grpc/status"

// Return a gRPC error
return nil, status.Errorf(codes.NotFound, "user %d not found", id)

// Common codes:
// codes.OK              = success
// codes.NotFound        = resource doesn't exist
// codes.InvalidArgument = bad input
// codes.Internal        = server error
// codes.Unauthenticated = not logged in
// codes.PermissionDenied = not authorized
```

### The protoc Command Explained

```bash
protoc \
  --go_out=.                          # Generate message structs → *.pb.go
  --go_opt=paths=source_relative      # Put generated files next to .proto
  --go-grpc_out=.                     # Generate service stubs → *_grpc.pb.go
  --go-grpc_opt=paths=source_relative # Same path rule
  proto/user.proto                    # Input .proto file
```

---

## 🔗 How the Services Connect

```
                    ┌──────────────┐
                    │    CLIENT    │
                    │  (Terminal 3)│
                    └───┬─────┬───┘
                        │     │
            GetUser     │     │  CreateOrder
            ListUsers   │     │  GetOrder
                        │     │  GetOrdersByUser
                        ▼     ▼
┌──────────────────┐    ┌──────────────────┐
│  USER SERVICE    │◄───│  ORDER SERVICE   │
│  :50051          │    │  :50052          │
│                  │    │                  │
│  - GetUser()     │    │  - CreateOrder() │
│  - ListUsers()   │    │  - GetOrder()    │
│                  │    │  - GetOrdersByUser│
└──────────────────┘    └──────────────────┘

Order Service CALLS User Service to verify
the user exists before creating an order.
This is service-to-service gRPC communication!
```

---

## 📝 Key Takeaways

1. **Proto file is the contract** — define once, generate code for any language
2. **protoc generates two files**: `*.pb.go` (messages) and `*_grpc.pb.go` (service stubs)
3. **Server implements the interface**, client uses the generated stub
4. **Streaming** lets you send multiple messages over a single connection
5. **Service-to-service calls** work the same way — one service is both server AND client
6. **Error codes** are gRPC-specific, not HTTP status codes
# gRPC--demo

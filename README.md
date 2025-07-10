## Implementation steps

#### Clone repository
```bash
git clone git@github.com:xreyc/grpc-auth.git
cd grpc-auth
```

#### Initialize module
```bash
go mod init github.com/xreyc/grpc-auth
```

#### Add grpc-contract as submodule
```bash
git submodule add https://github.com/xreyc/grpc-contract.git contract
git submodule update --init --recursive
```

These will generate

```
grpc-auth/contract/auth/v1/user.proto
```

#### Install require tools
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

#### Generate code for proto
```bash
mkdir -p internal/gen/go
```

Then run

```bash
protoc \
  --go_out=internal/gen/go \
  --go-grpc_out=internal/gen/go \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  -I contract \
  contract/auth/v1/user.proto
```

This will generate
```
internal/gen/go/auth/v1/user.pb.go
internal/gen/go/auth/v1/user_grpc.pb.go
```

#### Install dependencies
```bash
go get google.golang.org/grpc
go get github.com/gin-gonic/gin
```

#### Project structure
```
grpc-auth/
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── contract/                    # Submodule (grpc-contract)
├── internal/
│   ├── gen/go/                  # Generated .pb.go files
│   ├── handler/grpc/            # gRPC implementation
│   └── server/                  # gRPC server startup logic
├── go.mod
├── go.sum
└── README.md
```

#### Implement grpc handler
`internal/handler/grpc/user_handler.go`
```go
package grpc

import (
    "context"

    authv1 "github.com/xreyc/grpc-auth/internal/gen/go/auth/v1"
)

type UserHandler struct {
    authv1.UnimplementedUserServiceServer
}

func NewUserHandler() *UserHandler {
    return &UserHandler{}
}

func (h *UserHandler) GetUserDetails(ctx context.Context, req *authv1.GetUserRequest) (*authv1.GetUserResponse, error) {
    // Hardcoded response
    return &authv1.GetUserResponse{
        Username:  req.GetUsername(),
        Email:     "xreyc@example.com",
        FullName:  "Reyco Seguma",
    }, nil
}
```

#### Implement grpc startup code
`internal/server/grpc.go`
```go
package server

import (
    "log"
    "net"

    authv1 "github.com/xreyc/grpc-auth/internal/gen/go/auth/v1"
    userHandler "github.com/xreyc/grpc-auth/internal/handler/grpc"
    "google.golang.org/grpc"
)

func StartGRPCServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    authv1.RegisterUserServiceServer(grpcServer, userHandler.NewUserHandler())

    log.Println("gRPC server listening on :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
```

#### Implement entry point
`cmd/server/main.go`
```go
package main

import (
    "github.com/xreyc/grpc-auth/internal/server"
)

func main() {
    server.StartGRPCServer()
}
```

#### Run server
```bash
go run cmd/server/main.go
```

#### Create a Makefile
```makefile
APP_NAME := grpc-auth
MAIN := cmd/server/main.go

.PHONY: run build clean proto tidy

# Run the gRPC server
run:
	go run $(MAIN)

# Build the binary
build:
	go build -o bin/$(APP_NAME) $(MAIN)

# Clean build artifacts
clean:
	rm -rf bin

# Generate proto files
proto:
	protoc \
		--go_out=internal/gen/go \
		--go-grpc_out=internal/gen/go \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		-I contract \
		contract/auth/v1/user.proto

# Tidy up dependencies
tidy:
	go mod tidy
```

Usage
```bash
make run       # Run the gRPC server
make build     # Build binary to ./bin/grpc-auth
make proto     # Regenerate .pb.go files
make tidy      # Clean up go.mod/go.sum
make clean     # Remove ./bin folder
```
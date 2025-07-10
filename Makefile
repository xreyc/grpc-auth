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


# Variables
BINARY_NAME=api
OUTPUT_DIR=bin/


run:
	@go run cmd/api/main.go

test:
	@go test ./..

# Build the application
build:
	@go build -o $(OUTPUT_DIR)$(BINARY_NAME) cmd/api/main.go

# Compile and cross-compile the application
compile:
	@echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)$(BINARY_NAME)-linux-amd64 cmd/api/main.go
	GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)$(BINARY_NAME)-windows-amd64 cmd/api/main.go
	GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)$(BINARY_NAME)-darwin-amd64 cmd/api/main.go

# Clean up
clean:
	@echo "Cleaning up..."
	@rm -rf $(OUTPUT_DIR)

# Ensure dependencies are kept in sync
deps:
	@go mod tidy
	@go mod verify
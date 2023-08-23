
# Variables
BINARY_NAME=api
OUTPUT_DIR=bin/


run:
	@go run cmd/api/main.go

test:
	@go test ./... -v

test-coverage:
	@go test -v ./... -coverprofile=coverage.out
	@go tool cover -func=coverage.out

# Build the application
build:
	@go mod tidy
	@go mod verify
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)$(BINARY_NAME) cmd/api/main.go

docker-build:
	@echo "Cleaning up..."
	@rm -rf $(OUTPUT_DIR)

	@go mod tidy
	@go mod verify
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)$(BINARY_NAME) cmd/api/main.go
	@docker compose -f docker-compose.build.yml build

docker-push:
	@docker compose -f docker-compose.build.yml push

docker-run:
	@docker compose -f docker-compose.yml up -d

docker-stop:
	@docker compose -f docker-compose.yml down

# Compile and cross-compile the application
compile:
	@echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)$(BINARY_NAME)-linux-amd64 cmd/api/main.go
	#GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)$(BINARY_NAME)-windows-amd64 cmd/api/main.go
	#GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)$(BINARY_NAME)-darwin-amd64 cmd/api/main.go

# Clean up
clean:
	@echo "Cleaning up..."
	@rm -rf $(OUTPUT_DIR)

# Ensure dependencies are kept in sync
deps:
	@go mod tidy
	@go mod verify
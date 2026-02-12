.PHONY: build clean install test run

# Binary name
BINARY_NAME=smartcommit

# Build the application
build:
	@echo "Building..."
	go build -o bin/$(BINARY_NAME) main.go
	@echo "Build complete: bin/$(BINARY_NAME)"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -rf dist/
	go clean
	@echo "Clean complete"

# Install to GOPATH/bin
install:
	@echo "Installing..."
	go install
	@echo "Install complete"

# Run tests
test:
	@echo "Running tests..."
	go test ./... -v
	@echo "Tests complete"

# Run the application
run:
	go run main.go

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)-linux-amd64 main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY_NAME)-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/$(BINARY_NAME)-darwin-arm64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_NAME)-windows-amd64.exe main.go
	@echo "Multi-platform build complete"
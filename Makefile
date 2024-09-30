# Makefile

# Variables
GO_BIN=go
NEXT_BIN=npx next
BUILD_DIR=build
MODULE_NAME=barr.io
FRONTEND_DIR=frontend

# Default target
.PHONY: all
all: build

# Build Go application
.PHONY: build-go
build-go:
	@echo "Building Go application..."
	cd backend/api && $(GO_BIN) build -o $(BUILD_DIR) $(GO_APP)

# Build Next.js application
.PHONY: build-next
build-next:
	@echo "Building Next.js application..."
	cd $(FRONTEND_DIR) && $(NEXT_BIN) build

# Combined build target
build: build-go build-next

# Run the Go application
.PHONY: run-go
run-go:
	@echo "Running Go application..."
	cd backend/api && $(GO_BIN) run main.go

# Serve Next.js application
serve-next:
	@echo "Starting Next.js application..."
	cd $(FRONTEND_DIR) && $(NEXT_BIN) dev

# Initialize gqlgen
.PHONY: gqlgen-init
gqlgen-init:
	@echo "Initializing gqlgen..."
	cd backend $(GO_BIN) run github.com/99designs/gqlgen init


# Clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)/* $(FRONTEND_DIR)/.next

# Makefile

# Go build command
GO_BUILD = go build

# Output directory
OUTPUT_DIR = bin

# Binary name
BINARY_NAME = main

# Default target
default: build

# Build the Go project
build:
	$(GO_BUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME) ./cmd/main

# Run the built binary
run:
	./$(OUTPUT_DIR)/$(BINARY_NAME)

# Clean the built files
clean:
	rm -rf $(OUTPUT_DIR)

.PHONY: default build run clean

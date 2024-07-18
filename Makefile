# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=server
BUILD_DIR=bin

all: test build

build:
	@echo "Building the application..."
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) app/cmd/server/main.go

test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

clean:
	@echo "Cleaning up..."
	$(GOCLEAN)
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

run: build
	@echo "Running the application..."
	./$(BUILD_DIR)/$(BINARY_NAME)


.PHONY: all build clean test run deps

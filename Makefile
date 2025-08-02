GO_CMD=go
GO_BUILD=$(GO_CMD) build

MAIN_PATH=cmd/trackme/main.go
BINARY_NAME=trackme
BINARY_PATH=bin/$(BINARY_NAME)

build:
	@echo "Building $(BINARY_NAME) ..."
	@$(GO_BUILD) -o $(BINARY_PATH) $(MAIN_PATH)
	@echo "Binary created at $(BINARY_PATH)"

run:
	@echo "Running $(BINARY_NAME) ..."
	@./$(BINARY_PATH)

deploy-railway:
	@./scripts/deploy-railway.sh

.PHONY: build run deploy
BIN := bin/$(shell basename $(CURDIR))
MAIN := .

.PHONY: build
build:
	@echo "Building..."
	@go build -o $(BIN) $(MAIN)

.PHONY: swag
swag:

MAIN := cmd/api/main.go
BIN := main

.PHONY: all build run test clean watch

all: build

build:
	@echo "Building..."
	@templ generate
	@go build -o $(BIN) $(MAIN)

run:
	@echo "Running..."
	@templ generate
	@go run $(MAIN)

test:
	@echo "Testing..."
	@go test ./... -v

clean:
	@echo "Cleaning..."
	@rm -f $(BIN)

watch:
	@mkdir -p tmp
	@./scripts/watch.sh

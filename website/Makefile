MAIN := cmd/api/main.go
BIN := website

.PHONY: all build run test clean watch shutdown

all: build

build:
	@echo "Building..."
	@templ generate
	@go build -o $(BIN) $(MAIN)

run:
	@templ generate
	@go run $(MAIN)

test:
	@echo "Testing..."
	@go test ./... -v

clean:
	@echo "Cleaning..."
	@rm -f $(BIN)

# shutting down first in case there is an instance already running
watch:
	@./scripts/shutdown.sh
	@./scripts/watch.sh

shutdown:
	@./scripts/shutdown.sh

MAIN := cmd/api/main.go
BIN := website

.PHONY: all build run test clean watch

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

watch:
	@mkdir -p tmp
	@./scripts/watch.sh

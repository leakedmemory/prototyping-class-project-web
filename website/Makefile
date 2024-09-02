MAIN := cmd/api/main.go
BIN := website

.PHONY: all build run test clean watch shutdown restart

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
	@bash -c "./scripts/watch.sh > tmp/air.log 2>&1 &"
	@echo "Watching..."

shutdown:
	@./scripts/shutdown.sh

restart:
	@./scripts/shutdown.sh
	@bash -c "./scripts/watch.sh > tmp/air.log 2>&1 &"
	@echo "Watching..."

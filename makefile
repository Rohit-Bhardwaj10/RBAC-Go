APP_NAME=rbac-api
MAIN_PATH=./cmd/api/main.go
BINARY_NAME=bin/$(APP_NAME)

.PHONY: help build run dev clean test fmt tidy docker-up docker-down logs

help:
	@echo "Available commands:"
	@echo "  make build         - Build the Go binary"
	@echo "  make run           - Run the app"
	@echo "  make dev           - Run with live reload (if using air)"
	@echo "  make clean         - Remove built binaries"
	@echo "  make test          - Run tests"
	@echo "  make fmt           - Format code"
	@echo "  make tidy          - Clean go.mod"
	@echo "  make docker-up     - Start postgres via docker-compose"
	@echo "  make docker-down   - Stop docker containers"
	@echo "  make logs          - View docker logs"

build:
	@echo "Building..."
	@mkdir -p bin
	go build -o $(BINARY_NAME) $(MAIN_PATH)

run:
	@echo "Running..."
	go run $(MAIN_PATH)

dev:
	@echo "Starting dev mode..."
	air

clean:
	@echo "Cleaning..."
	rm -rf bin

test:
	go test ./...

fmt:
	go fmt ./...

tidy:
	go mod tidy

docker-up:
	docker compose up -d

docker-down:
	docker compose down

logs:
	docker compose logs -f

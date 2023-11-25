# Simple Makefile for a Go project

postgres: 
	docker run --name postgres-alpine -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -d postgres:alpine

createdb:
	docker exec -it postgres-alpine createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres-alpine drop simple_bank

migrateup:
	migrate -path internal/db/migration -database "postgres://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migration -database "postgres://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc-install:
	docker run --rm -v C:/Users/subro/Desktop/x/Go/simplebank:/src -w /src kjconroy/sqlc generate
	
sqlc-generate:
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate

sqlc-compile:
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc compile
# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	# @go test ./tests -v
	@go test -v -cover ./...

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if [ -x "$(GOPATH)/bin/air" ]; then \
	    "$(GOPATH)/bin/air"; \
		@echo "Watching...";\
	else \
	    read -p "air is not installed. Do you want to install it now? (y/n) " choice; \
	    if [ "$$choice" = "y" ]; then \
			go install github.com/cosmtrek/air@latest; \
	        "$(GOPATH)/bin/air"; \
				@echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

# .PHONY: all build run test clean
.PHONY: postgres createdb 

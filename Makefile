APP_NAME=crud-golang
MAIN_FILE=cmd/main.go

.PHONY: build run clean deps

build:
	go build -o $(APP_NAME) $(MAIN_FILE)

run:
	go run $(MAIN_FILE)

deps:
	go mod tidy
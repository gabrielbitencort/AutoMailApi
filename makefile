.PHONY: default run build test docs clean

# Variables
APP_NAME=AutoMailApi

# Tasks
default: run-with-docs

run:
	@go run main.go
run-with-docs:
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@del $(APP_NAME).exe 2>NUL
	@rmdir /S /Q docs
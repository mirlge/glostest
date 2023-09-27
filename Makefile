BINARY_NAME=glostest

build:
	go build -o out/${BINARY_NAME} cmd/glostest/main.go

dev:
	go run cmd/glostest/main.go example.json

tidy:
	go mod tidy


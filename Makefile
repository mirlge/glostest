BINARY_NAME=glostest

build:
	go build -o out/${BINARY_NAME} cmd/glostest/main.go

run:
	go run cmd/glostest/main.go

tidy:
	go mod tidy


BINARY_NAME=glossary_test

build:
	go build -o ${BINARY_NAME} cmd/glossary_test/main.go

run:
	go run cmd/glossary_test/main.go

tidy:
	go mod tidy


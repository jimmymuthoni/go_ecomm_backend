build:
	@go build -o bin/go_eccom_backend

run: build
	@./bin/go_eccom_backend

test:
	@go test -v ./...
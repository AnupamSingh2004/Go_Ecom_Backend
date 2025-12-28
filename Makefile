build:
	@go build -o bin/Ecom_Backend cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/Ecom_Backend
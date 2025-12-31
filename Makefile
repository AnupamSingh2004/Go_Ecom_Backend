build:
	@go build -o bin/Ecom_Backend cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/Ecom_Backend

migration:
	@~/go/bin/migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
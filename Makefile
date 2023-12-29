all: test gen-doc run

gen-doc:
	swag init -g cmd/app.go

run:
	go run main.go

test:
	go test ./...

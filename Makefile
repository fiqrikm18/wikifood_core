all: test gen-doc run

gen-doc:
	swag init -g cmd/app.go

run:
	go run main.go

test:
	go test ./...

test-cover:
	go test ./... --cover

migrate-up:
	migrate -path ./database/migrations/ -database "postgres://postgres:postgres@127.0.0.1:5432/wikifood?sslmode=disable" -verbose up

migrate-down:
	migrate -path ./database/migrations/ -database "postgres://postgres:postgres@127.0.0.1:5432/wikifood?sslmode=disable" -verbose down

create-scheme:
	migrate create -ext sql -dir database/migrations/ -seq $(name)
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
	migrate -source database/migrations/ -database postgres://postgres:postgres@127.0.0.1:5432/wikifood -verbose up

migrate-down:
	migrate -source database/migrations/ -database postgres://postgres:postgres@127.0.0.1:5432/wikifood -verbose down

create-scheme:
	migrate create -ext sql -dir database/migrations/ -seq $(seq)
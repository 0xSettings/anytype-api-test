APP_NAME=flow-anytype-crud

run:
	go run ./cmd/main.go

build:
	go build -o bin/$(APP_NAME) ./cmd/main.go

tidy:
	go mod tidy

test:
	go test ./...

clean:
	rm -rf bin

env:
	cp .env.example .env

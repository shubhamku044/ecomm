build:
	@go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go

test:
	@go test -v ./...

run: 
	@go run cmd/$(APP_NAME)/main.go

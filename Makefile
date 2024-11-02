build:
	@go build -0 bin/backend-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/backend-api
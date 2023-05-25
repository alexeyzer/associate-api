run:
	@go run cmd/main.go

generate:
	@buf generate
lint:
	@golangci-lint run

build-image:
	docker build --no-cache -t associate-api .

up:
	goose up ""
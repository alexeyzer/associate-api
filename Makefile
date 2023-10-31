run:
	@go run cmd/main.go

generate:
	@buf generate
lint:
	@golangci-lint run

build-image:
	docker build --no-cache -t associate-api .

up:
	goose postgres "host=database user=test password=test dbname=test sslmode=disable" status
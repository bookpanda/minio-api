watch:
	air

server:
	go run cmd/main.go

swagger:
	swag init -d ./cmd -o ./docs -md ./docs/markdown
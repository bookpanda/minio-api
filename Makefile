watch:
	air

server:
	go run cmd/main.go

mock-gen:
	mockgen -source ./internal/repository/file/file.repository.go -destination ./mocks/repository/file.mock.go
	mockgen -source ./internal/service/file/file.service.go -destination ./mocks/service/file.mock.go

swagger:
	swag init -d ./internal/file -g ../../cmd/main.go -o ./docs -md ./docs/markdown --parseDependency --parseInternal
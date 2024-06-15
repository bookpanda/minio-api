watch:
	air

server:
	go run cmd/main.go

mock-gen:
	mockgen -source ./internal/repository/file/file.repository.go -destination ./mocks/repository/file.repository.go
	mockgen -source ./internal/service/file/file.service.go -destination ./mocks/service/file.service.go
	mockgen -source ./internal/client/minio.client.go -destination ./mocks/client/minio.client.go

swagger:
	swag init -d ./internal/file -g ../../cmd/main.go -o ./docs -md ./docs/markdown --parseDependency --parseInternal
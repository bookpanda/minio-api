watch:
	air

server:
	go run cmd/main.go

stage:
	docker-compose -f docker-compose.stage.yml up

prod:
	docker-compose -f docker-compose.prod.yml up

mock-gen:
	mockgen -source ./internal/repository/file/file.repository.go -destination ./mocks/repository/file.repository.go
	mockgen -source ./internal/service/file/file.service.go -destination ./mocks/service/file.service.go
	mockgen -source ./internal/client/http/http.client.go -destination ./mocks/client/http/http.client.go
	mockgen -source ./internal/client/store/store.client.go -destination ./mocks/client/store/store.client.go
	mockgen -source ./internal/router/context.go -destination ./mocks/router/context.go
	mockgen -source ./internal/validator/validator.go -destination ./mocks/validator/validator.go
	mockgen -source ./metrics/requests.metrics.go -destination ./mocks/metrics/requests.metrics.go

test:
	go vet ./...
	go test  -v -coverpkg ./internal/... -coverprofile coverage.out -covermode count ./internal/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

swagger:
	swag init -d ./internal/file -g ../../cmd/main.go -o ./docs -md ./docs/markdown --parseDependency --parseInternal
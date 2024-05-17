
.PHONY: migrate-up
migrate-up:
	@go run main.go migrate up

.PHONY: migrate-down
migrate-down:
	@go run main.go migrate down

.PHONY: migrate-fresh
migrate-fresh:
	@go run main.go migrate fresh

run : 
	go run main.go rest

docs : 
	swag init -g internal/delivery/http/rest.go --parseDependency true --parseInternal

mock-repository:
	mockgen -source=./internal/repository/author.go -destination=./shared/mock/repository/author_mock.go -package repository
	mockgen -source=./internal/repository/user.go -destination=./shared/mock/repository/user_mock.go -package repository
	mockgen -source=./internal/repository/book.go -destination=./shared/mock/repository/book_mock.go -package repository
	mockgen -source=./internal/repository/transaction.go -destination=./shared/mock/repository/transaction_mock.go -package repository

mock-pkg:
	mockgen -source=./pkg/elasticsearch/elasticsearch.go -destination=./shared/mock/pkg/elasticsearch_mock.go -package pkg

test:
	go test -v -cover -count=1 -failfast ./... -coverprofile="coverage.out"
include .env

LOCAL_MIGRATION_DIR=$(MIGRATIONS_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=55010 dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"
MIGRATION_NAME=""

### Project tools
install:
	make install-linter
	go install mvdan.cc/gofumpt@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

generate-proto:
	protoc --proto_path api/proto \
    	--go_out=pkg/generated/wisdom --go_opt=paths=source_relative \
    	--go-grpc_out=pkg/generated/wisdom --go-grpc_opt=paths=source_relative \
		api/proto/wisdom.proto

install-linter:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

### Continues Integration code quality check
prepare:
	make tidy
	make fmt
	make lint
	make test

tidy:
	go mod tidy
	go mod vendor

fmt:
	gofumpt -w .

lint:
	golangci-lint run ./pkg/... ./internal/...

test:
	go test -v -race ./pkg/... ./internal/...

### Migrations
migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

migration-create:
	goose -dir ${LOCAL_MIGRATION_DIR} create ${MIGRATION_NAME} sql

### Continues Deployment
docker-up:
	docker compose up --build --detach

docker-stop:
	docker compose stop

docker-down:
	docker compose down --remove-orphans --volumes
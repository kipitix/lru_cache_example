export GOOSE_DRIVER := postgres
export GOOSE_DBSTRING := host=localhost port=5432 user=postgres password=postgres dbname=lrucache_psql sslmode=disable
export APP_PORT := 8080
export APP_DSN := $(GOOSE_DBSTRING)

goose-install:
	@go install github.com/pressly/goose/v3/cmd/goose@latest

bench-install:
	@go install github.com/cmpxchg16/gobench@latest

migrate:
	@goose -dir=./migrations up

unmigrate:
	@goose -dir=./migrations down

run:
	@go run ./cmd

compose-up:
	@docker-compose up -d --build

compose-down:
	@docker-compose down --remove-orphans

stress:
	@gobench -u http://localhost:8080/user/email@example.com -k=true -c 500 -t 10

curl-user:
	@curl -s localhost:8080/user/email@example.com
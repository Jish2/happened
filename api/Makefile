

.PHONY: watch migrate


deps:
	@go install "github.com/air-verse/air@v1.61.1"
	
fmt:
	@go fmt ./...

test:
	@go test ./...

api:
	@go build -o ./bin/api ./cmd/api

gen:
	@buf generate

watch: db
	@air

# Applies the specified schema file to the database, similar to Terraform.
migrate:
	@atlas schema apply --url "postgres://admin:admin@localhost:5433/happened_db?sslmode=disable" --to "file://sql/schema.sql" --dev-url "docker://postgres/15"

# Connects to the local Postgres database
connect:
	@psql "postgres://admin:admin@localhost:5433/happened_db"

# Startup the development database
db:
	@docker compose up -d

# Introspects database and returns current schema
inspect:
	@atlas schema inspect -u "postgres://admin:admin@localhost:5433/happened_db?sslmode=disable" --format '{{ sql . }}' > schema.sql


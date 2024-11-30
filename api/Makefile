

default: api

.PHONY: watch migrate

.PHONY: openapi
openapi: # Generate OpenAPI specification
deps:
	@go install "github.com/air-verse/air@v1.61.1"
	@brew tap golangci/tap
	@brew install golangci/tap/golangci-lint

check:
	@go vet ./...
	@golangci-lint run ./...

openapi:
	@./bin/api openapi > openapi.yaml


.PHONY: mocks
mocks: # Generate mocks
	@go generate ./...

clean:
	@rm -rf ./bin

deps:
	@go install "github.com/air-verse/air@v1.61.1"
	@go install "github.com/vektra/mockery/v2@v2.49.1"
	
fmt:
	@go fmt ./...

test:
	@go test -cover ./...

api:
	@go build -o ./bin/api ./cmd/api


IMAGE_NAME := anmho/happened
GIT_COMMIT_HASH = $(shell git rev-parse --short HEAD)
TAG ?= $(GIT_COMMIT_HASH)
image:
	@docker build -t $(IMAGE_NAME):$(TAG) .
	@echo docker.io/$(IMAGE_NAME):$(TAG) | tee image.txt

push:
	@docker push $(IMAGE_NAME):$(TAG)
	@echo docker.io/$(IMAGE_NAME):$(TAG) | tee image.txt

watch:
	@air

# Database commands
# Startup the development database
db:
	@docker compose up -d

# Applies the specified schema file to the database, similar to Terraform.
migrate:
	@atlas schema apply --url "postgres://admin:admin@localhost:5433/happened_db?sslmode=disable" --to "file://sql/schema.sql" --dev-url "docker://postgres/15"

# Connects to the local Postgres database
connect:
	@psql "postgres://admin:admin@localhost:5433/happened_db"

# Introspects database and returns current schema
inspect:
	@atlas schema inspect -u "postgres://admin:admin@localhost:5433/happened_db?sslmode=disable" --format '{{ sql . }}' > schema.sql


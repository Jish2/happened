

.PHONY: deps
deps:
	@go install github.com/air-verse/air@latest


.PHONY: build
build:
	@cd api && go build -o ./bin/api ./cmd/api
.PHONY: watch
watch:
	@air --build.cmd "make build" --build.bin "./api/bin/api"


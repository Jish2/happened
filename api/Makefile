




.PHONY: api gen watch openapi

api:
	@go build -o ./bin/api ./cmd/api

.PHONY: gen
gen:
	@buf generate

.PHONY: watch
watch:
	@air --build.cmd "make api" --build.bin "./bin/api"

# openapi:
# 	@curl http://localhost:8080/openapi.yaml >
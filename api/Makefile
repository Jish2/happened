

.PHONY: api
api:
	@go build -o ./bin/api ./cmd/api

.PHONY: gen
gen:
	@buf generate

.PHONY: watch
watch:
	@air --build.cmd "make api" --build.bin "./bin/api"
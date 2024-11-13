

.PHONY: api
api:
	@go build -o ./bin/api ./cmd/api

.PHONY: gen
gen:
	@buf generate

.PHONY: run
run:
	@./bin/api -stage development

.PHONY: watch
watch:
	@air --build.cmd "make api" --build.bin "make run"

.PHONY: image
image:
	@docker build --platform linux/amd64 -t anmho/happened -f Dockerfile  .

.PHONY: push
push: image
	@docker push anmho/happened
	
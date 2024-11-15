# .PHONY: deps
# deps:
# 	@go install "github.com/air-verse/air@v1.61.1"
# 	@brew tap hashicorp/tap
# 	@brew install hashicorp/tap/terraform

deps:
	@npm install -g orval

.PHONY: openapi gen

watch:
	$(MAKE) -C api watch

# Server must be running to generate the latest spec.
openapi:
	$(MAKE) -C api openapi

# Generates the client SDK from the server's current OpenAPI spec.
gen: openapi
	@orval --input ./api/openapi.yaml --output ./client/gen/openapi.ts



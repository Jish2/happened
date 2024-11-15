
deps:
	@npm install -g orval

watch:
	$(MAKE) -C api watch

# Generates the client SDK from the server's current OpenAPI spec.
gen:
	@orval --input ./api/openapi.yaml --output ./client/gen/openapi.ts



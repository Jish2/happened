

deps:
	@npm install -g orval

.PHONY: openapi gen

watch:
	$(MAKE) -C api watch

# Server must be running to generate the latest spec.
# openapi:
# 	$(MAKE) -C api openapi

# Generates the client SDK from the server's current OpenAPI spec.
gen:
	@orval --input ./api/openapi.yaml --output ./client/gen/openapi.ts



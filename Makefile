
default:
	@echo "Starting API and Client SDK builds..."
	@(make -C api) & pid1=$$!; \
	 (make gen) & pid2=$$!; \
	 wait $$pid1; status1=$$?; \
	 wait $$pid2; status2=$$?; \
	 if [ $$status1 -eq 0 ] && [ $$status2 -eq 0 ]; then \
	     echo "API and Client SDK built successfully!"; \
	 else \
	     echo "One or more processes failed."; \
	     [ $$status1 -ne 0 ] && echo "API build failed with status $$status1"; \
	     [ $$status2 -ne 0 ] && echo "Client SDK build failed with status $$status2"; \
	     exit 1; \
	 fi

test:
	$(MAKE) -C api test

.PHONY: gen
deps:
	@npm install -g orval

clean:
	@rm -rf ./client/gen

watch:
	$(MAKE) -C api watch

.PHONY: gen
# Generates the client SDK from the server's current OpenAPI spec.
gen:
	@$(MAKE) -C api openapi
	@orval --input ./api/openapi.yaml --output ./client-v2/gen/openapi.ts



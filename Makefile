.PHONY: deps
deps:
	@brew install bufbuild/buf/buf
	@go install "google.golang.org/protobuf/cmd/protoc-gen-go@latest"
	@go install "connectrpc.com/connect/cmd/protoc-gen-connect-go@latest"
	@go install "github.com/air-verse/air@v1.61.1"
	@brew tap hashicorp/tap
	@brew install hashicorp/tap/terraform


.PHONY: api
api:
	$(MAKE) -C api

.PHONY: watch
watch:
	$(MAKE) -C api watch

.PHONY: gen
gen:
	$(MAKE) -C api gen
	@cd client; yarn gen

.PHONY: init-tf
init-tf:
	@terraform -chdir=./terraform init

.PHONY: tf
tf:
	@terraform -chdir=./terraform apply

.PHONY: destroy
destroy:
	@terraform -chdir=./terraform destroy



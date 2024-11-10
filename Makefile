

.PHONY: deps
deps:
	@go install "github.com/bufbuild/buf/cmd/buf@latest"
	@go install "github.com/fullstorydev/grpcurl/cmd/grpcurl@latest"
	@go install "google.golang.org/protobuf/cmd/protoc-gen-go@latest"
	@go install "connectrpc.com/connect/cmd/protoc-gen-connect-go@latest"
	@go install "github.com/air-verse/air@latest"


.PHONY: build
build:
	@cd api && go build -o ./bin/api ./cmd/api

.PHONY: watch
watch:
	@air --build.cmd "make build" --build.bin "./api/bin/api" --build.include_dir api

.PHONY: gen
gen:
	@buf generate

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



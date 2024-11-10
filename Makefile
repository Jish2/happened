

.PHONY: deps
deps:
	@go install "github.com/bufbuild/buf/cmd/buf@latest"
	@go install "github.com/fullstorydev/grpcurl/cmd/grpcurl@latest"
	@go install "google.golang.org/protobuf/cmd/protoc-gen-go@latest"
	@go install "connectrpc.com/connect/cmd/protoc-gen-connect-go@latest"
	@go install "github.com/air-verse/air@latest"
	@npm install -g @connectrpc/protoc-gen-connect-es @bufbuild/protoc-gen-es


.PHONY: build
build:
	@cd api && go build -o ./bin/api ./cmd/api

.PHONY: watch
watch:
	@air --build.cmd "make build" --build.bin "./api/bin/api"

.PHONY: gen
gen:
	@buf generate
	

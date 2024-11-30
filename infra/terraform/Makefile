
.PHONY: init-tf tf destroy

deps:
	@brew tap hashicorp/tap
	@brew install hashicorp/tap/terraform

init-tf:
	@terraform init

apply:
	@terraform apply

destroy:
	@terraform destroy

fmt:
	@terraform fmt
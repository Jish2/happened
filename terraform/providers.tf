

provider "aws" {
  region = var.aws_region
}

provider "google" {
  project = var.gcp_project_id
}
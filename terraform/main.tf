
module "s3" {
  source = "./modules/aws/s3"
  region = "us-west-2"
}

module "cloud_run" {
  source              = "./modules/gcp/cloud_run"
  project_id          = var.gcp_project_id
  image_url           = var.api_image_url
  location            = var.gcp_region
  deletion_protection = var.cloud_run_deletion_protection
  name                = var.api_name
  port                = var.api_port
}


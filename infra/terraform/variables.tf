

variable "aws_region" {
  description = "The AWS Region to deploy"
  type        = string
  default     = "us-west-2"
}

variable "gcp_project_id" {
  description = "The GCP Project ID"
  type        = string
  default     = "happened-441602"
}

variable "gcp_region" {
  description = "The GCP Region"
  type        = string
  default     = "us-west1"
}


variable "api_image_url" {
  description = "Image URL to run on Cloud Run"
  type        = string
  #   default     = "docker.io/anmho/happened:latest"
  default = "docker.io/hashicorp/http-echo"
}

variable "api_name" {
  description = "Name of the Cloud Run service to deploy"
  type        = string
  default     = "happened-api"
}

variable "api_port" {
  description = "Port API is listening on"
  type        = string
  default     = "5678"
}

variable "cloud_run_deletion_protection" {
  description = "Enable deletion protection"
  type        = bool
  default     = false
}
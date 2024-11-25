variable "project_id" {
  description = "The project ID to deploy to"
  type = string
}

variable "image_url" {
  description = "The url of the image to deploy to cloud run from a container registry"
  type = string
}
variable "location" {
  description = "location to deploy on Cloud Run"
  type = string
}

variable "deletion_protection" {
  description = "Enable deletion protection or not"
  type = bool
}

variable "name" {
  description = "Service name"
  type = string
}

variable "port" {
  description = "Application container port"
  type = string
}
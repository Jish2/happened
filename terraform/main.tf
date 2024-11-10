provider "aws" {
  region = var.region
}

resource "aws_s3_bucket" "happened-images" {
  bucket = "happened-images"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}


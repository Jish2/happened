terraform {
  backend "s3" {
    bucket = "happened-terraform"
    key    = "./state"
    region = "us-west-2"
  }
}
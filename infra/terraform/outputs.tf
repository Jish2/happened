output "s3_iam_admin_username" {
  value = module.s3.iam_admin_username
}

output "s3_iam_admin_id" {
  value = module.s3.iam_admin_id
}

output "s3_iam_admin_secret" {
  value     = module.s3.iam_admin_secret
  sensitive = true
}

output "s3_aws_region" {
  value = module.s3.aws_region
}

output "s3_bucket_name" {
  value = module.s3.bucket_name
}
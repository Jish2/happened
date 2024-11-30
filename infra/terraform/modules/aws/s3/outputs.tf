# Output the bucket name and ARN
output "bucket_name" {
  value = aws_s3_bucket.happened-bucket.id
}

output "bucket_arn" {
  value = aws_s3_bucket.happened-bucket.arn
}

output "iam_admin_username" {
  value = aws_iam_access_key.admin.user
}

output "iam_admin_id" {
  value = aws_iam_access_key.admin.id
}

output "iam_admin_secret" {
  value = aws_iam_access_key.admin.secret
  sensitive = true
}

output "aws_region" {
  value = var.region
}


# Output the bucket name and ARN
output "bucket_name" {
  value = aws_s3_bucket.happened-bucket.id
}

output "bucket_arn" {
  value = aws_s3_bucket.happened-bucket.arn
}

# Outputs for credentials
output "access_key_id" {
  value     = aws_iam_access_key.admin.id
  sensitive = true
}

output "secret_access_key" {
  value     = aws_iam_access_key.admin.secret
  sensitive = true
}


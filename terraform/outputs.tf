
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

# Go to the aws console to get access key id and secret access key


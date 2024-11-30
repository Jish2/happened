# Create IAM user
resource "aws_iam_user" "admin" {
  name = "happened-admin"  # Choose your username
  path = "/"
}

# Create access key
resource "aws_iam_access_key" "admin" {
  user = aws_iam_user.admin.name
}

# Create S3 bucket
resource "aws_s3_bucket" "happened-bucket" {
  bucket = "happened-bucket"  # Replace with your desired bucket name
}

# Block all public access
resource "aws_s3_bucket_public_access_block" "happened-bucket" {
  bucket = aws_s3_bucket.happened-bucket.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# Create bucket policy
data "aws_iam_policy_document" "admin_bucket_policy" {
  statement {
    effect = "Allow"
    actions = [
      "s3:*"
    ]
    resources = [
      aws_s3_bucket.happened-bucket.arn,
      "${aws_s3_bucket.happened-bucket.arn}/*"
    ]
  }
}


resource "aws_iam_policy" "admin_bucket_policy" {
  name = "happened-s3-admin-access"
  description = "Policy granting admin access to specific S3 buckets"
  policy      = data.aws_iam_policy_document.admin_bucket_policy.json
}

resource "aws_iam_user_policy_attachment" "s3_admin" {
  user = "happened-admin"
  policy_arn = aws_iam_policy.admin_bucket_policy.arn
}
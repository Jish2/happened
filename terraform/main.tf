provider "aws" {
  region = var.region
}

# Get current user identity
data "aws_caller_identity" "current" {}

resource "aws_s3_bucket" "happened-images" {
  bucket = "happened-images"
}


# Disable block public access settings
resource "aws_s3_bucket_public_access_block" "example" {
  bucket = aws_s3_bucket.happened-images.id

  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}


# S3 bucket policy document
data "aws_iam_policy_document" "s3_limited_access" {
  # GetObject-only access for all users
  statement {
    effect = "Allow"
    actions = [
      "s3:GetObject"
    ]
    resources = [
      "${aws_s3_bucket.happened-images.arn}/*"
    ]
    principals {
      type = "AWS"
      identifiers = ["*"]
    }
  }

  # Full access for current user
  statement {
    effect = "Allow"
    actions = [
      "s3:PutObject",
      "s3:GetObject",
      "s3:DeleteObject",
      "s3:ListBucket",
      "s3:GetBucketLocation"
    ]
    resources = [
      aws_s3_bucket.happened-images.arn,
      "${aws_s3_bucket.happened-images.arn}/*"
    ]
    principals {
      type = "AWS"
      identifiers = [data.aws_caller_identity.current.arn]
    }
  }
}

# Attach policy to S3 bucket
resource "aws_s3_bucket_policy" "bucket_policy" {
  bucket = aws_s3_bucket.happened-images.id
  policy = data.aws_iam_policy_document.s3_limited_access.json
}

# IAM user policy for additional permissions
data "aws_iam_policy_document" "user_s3_access" {
  statement {
    effect = "Allow"
    actions = [
      "s3:ListBucket",
      "s3:GetObject",
      "s3:PutObject",
      "s3:DeleteObject",
      "s3:GetBucketLocation"
    ]
    resources = [
      aws_s3_bucket.happened-images.arn,
      "${aws_s3_bucket.happened-images.arn}/*"
    ]
  }
}

# Create IAM policy
resource "aws_iam_policy" "user_s3_policy" {
  name        = "user-s3-limited-access"
  description = "Policy allowing full access to current user and GetObject for others"
  policy      = data.aws_iam_policy_document.user_s3_access.json
}

# Attach policy to current user
resource "aws_iam_user_policy_attachment" "attach_s3_access" {
  user       = data.aws_caller_identity.current.user_id
  policy_arn = aws_iam_policy.user_s3_policy.arn
}
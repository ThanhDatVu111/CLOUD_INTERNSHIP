terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# Configure the AWS provider
provider "aws" {
  region = var.region
}

# S3 Bucket Resource with Description
resource "aws_s3_bucket" "backend_bucket" {
  bucket        = var.backend_bucket
  acl           = var.bucket_acl
  force_destroy = true # This setting allows the bucket to be destroyed even if it contains objects

  tags = {
    Name        = "backend-bucket"
    Environment = var.environment
  }

  versioning {
    enabled = true # Enable versioning for the S3 bucket
  }

  lifecycle {
    prevent_destroy = false # Allow the bucket to be destroyed
  }
}

# S3 Bucket Policy
resource "aws_s3_bucket_policy" "backend_bucket_policy" {
  bucket = aws_s3_bucket.backend_bucket.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid = "AllowPutObject"
        Effect = "Allow"
        Principal = "*"
        Action = "s3:PutObject"
        Resource = "arn:aws:s3:::${aws_s3_bucket.backend_bucket.id}/*"
      }
    ]
  })
}

# S3 Bucket CORS Configuration
resource "aws_s3_bucket_cors_configuration" "backend_bucket_cors" {
  bucket = aws_s3_bucket.backend_bucket.id

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["GET", "PUT", "POST", "DELETE"]
    allowed_origins = ["http://localhost:3000"]
    expose_headers  = []
    max_age_seconds = 3000
  }
}

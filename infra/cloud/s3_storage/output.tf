output "bucket_name" {
  description = "The name of the S3 bucket"
  value       = var.backend_bucket
}

output "bucket_acl" {
  description = "The ACL of the S3 bucket"
  value       = var.bucket_acl
}

output "region" {
  description = "The AWS region used"
  value       = var.region
}

output "environment" {
  description = "The environment of the S3 bucket"
  value       = var.environment
}
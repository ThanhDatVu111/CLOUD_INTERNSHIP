# Variables Block
variable "backend_bucket" {
  type    = string
  default = "chakratech-bucket-2024"
}

variable "environment" {
  type    = string
  default = "development"
}

variable "region" {
  type    = string
  default = "us-west-2"
}

variable "bucket_acl" {
  description = "The ACL to apply to the S3 bucket"
  type        = string
  default     = "private" # Changed to private as a safer default
}

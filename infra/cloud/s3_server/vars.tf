variable "region" {
  type    = string
  default = "us-west-2"
}

variable "env_name" {
  description = "Name of the environment (dev)"
  type        = string
  default     = "dev"  
}

# Variable for the service account role name
variable "service_account_role_name" {
  description = "Name for the IAM role used by the service account"
  type        = string
  default     = "s3-read-only-role"
}


variable "s3_access_policy_name" {
  description = "IAM policy granting S3 access (GET/PUT/POST)"
  type        = string
  default     = "s3-read-only" 
}

variable "service"{
    description = "Type of service that will be used"
    type = string
    default = "s3.amazonaws.com"
}

variable "backend_bucket"{
  description = "The s3 resource that is allowed access"
  type = string
  default = "${var.env}-backend-platform-2024"
}
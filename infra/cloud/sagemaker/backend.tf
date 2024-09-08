# Setup S3 bucket for the SageMaker backend
terraform {
  backend "s3" {
    bucket         = "infra-state-2024"
    key            = ".terraform/terraform.tfstate"
    region         = "us-west-2"
    dynamodb_table = "infra_lock"
    encrypt        = true
  }
}
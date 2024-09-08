variable "region" {
  description = "AWS region"
  type        = string
  default     = "us-west-2"
}

variable "vpc_id" {
  description = "The ID of the existing VPC"
  type        = string
  default     = ""
}

variable "subnet_ids" {
  description = "A list of subnet IDs within the VPC"
  type        = list(string)
  default     = [""]
}

variable "sagemaker_model_artifact_path" {
  description = "The S3 path to the model artifacts"
  type        = string
  default     = ""
}

variable "sagemaker_image_uri" {
  description = "The URI of the Docker image to use for the SageMaker model"
  type        = string
  default     = ""
}
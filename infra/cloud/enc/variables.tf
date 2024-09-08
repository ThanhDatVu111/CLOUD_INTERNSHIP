variable "region" {
  type    = string
  default = "us-west-2"
}

variable "ami" {
  description = "ID of the Amazon Machine Image (AMI) to use"
  type        = string
  default     = "ami-12345678"
}

variable "instance_type" {
  description = "EC2 the instance type"
  type        = string
  default     = "t2.micro"
}

variable "encrypted_credentials" {
  description = "The path to the encrypted credentials file"
  type        = string
  default     = "credentials/login.enc.json"
}


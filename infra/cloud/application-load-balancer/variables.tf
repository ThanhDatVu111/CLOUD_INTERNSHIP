variable "name" {
  description = "The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen"
  type        = string
  default     = "backend"
}

variable "region" {
  type    = string
  default = "us-west-2"
}

variable "azs" {
  description = "Availability Zones"
  type        = list(string)
  default     = ["us-west-2a", "us-west-2b"]
}

variable "security_groups" {
  description = "A list of security group IDs to assign to the LB"
  type        = list(string)
  default     = []
}

variable "vpc_id" {
  description = "Identifier of the VPC where the security group will be created"
  type        = string
  default     = "vpc-0ae1622eadb6ea9ab"
}

variable "subnets" {
  description = "A list of subnet IDs to attach to the LB. Subnets cannot be updated for Load Balancers of type `network`. Changing this value for load balancers of type `network` will force a recreation of the resource"
  type        = list(string)
  default     = ["04066e8968a9da8ef"]
}

variable "subnet_mapping" {
  description = "A list of subnet mapping blocks describing subnets to attach to load balancer"
  type        = list(map(string))
  default     = []
}

variable "client_cidr_block" {
  description = "The IPv4 CIDR block to assign client IP addresses from (minimum size /22)"
  type        = string
  default     = "10.0.0.0/23"
}

variable "public_subnet_cidrs" {
  description = "Public Subnet CIDR values"
  type        = list(string)
  default     = ["10.0.1.0/24"]
}

variable "ami" {
  description = "Amazon Linux AMI"
  type        = string
  default     = "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2"
}

variable "instances_per_subnet" {
  description = "Number of EC2 instances in each private subnet"
  type        = number
  default     = 1
}

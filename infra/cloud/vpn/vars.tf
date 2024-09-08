variable "region" {
  type    = string
  default = "us-west-2"
}

#todo to get the vpc id from another module, that be greate
variable "vpc_id" {
  type        = string
  default     = "vpc-0ff1aa60e31ce30e4"
  description = "vpc id for the us-west-2 region"
}

variable "client_cidr_block" {
  type        = string
  default     = "20.0.0.0/22"
  description = "The IPv4 CIDR block to assign client IP addresses from (minimum size /22)"
}


variable "root_certificate_chain_arn" {
  description = "ARN of the client certificate used to authenticate clients connecting to the Client VPN endpoint"
  type        = string
  default     = "arn:aws:acm:us-west-2:431095965741:certificate/81395312-83fe-4a79-b568-c8c036e59453"
}

variable "dns_servers" {
  type        = list(string)
  default     = [] #Empty list to use default VPC DNS servers
  description = "a list of custom DNS servers for the Client VPN endpoint"
}


variable "env" {
  description = "CloudWatch log environment"
  type        = string
  default     = "dev"
}

variable "cloudwatch_log_group" {
  description = "The name of the CloudWatch log group"
  type        = string
  default     = "/aws/vpn/clientvpn"
}

variable "domain_name" {
  type        = string
  description = "This a string of domain names"
  default     = "chakratechwork.com"
}



variable "subnet_id" {
  description = "The subnet_id the vpn will be routed to"
  type        = string
  default     = "subnet-07b11f3108a5430de"
}

variable "subnet_cidr"{
  description = "subnet cidr"
  type        = string
  default     = "10.0.3.0/24"
}

variable "vpc_cidr"{
  description = "vpc_cidr"
  type        = string
  default     = "10.0.0.0/16"
}
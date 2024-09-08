variable "aws_region" {
  description = "AWS region where resources will be provisioned"
  type        = string
  default     = "us-west-2"
}

variable "db_engine" {
  description = "Database engine type"
  type        = string
  default     = "postgres"
}

variable "db_engine_version" {
  description = "Database engine version"
  type        = string
  default     = "16.3" # Latest stable version as of June 2024
}

variable "db_instance_class" {
  description = "Instance class for the RDS instance"
  type        = string
  default     = "db.t3.micro"
}

variable "db_allocated_storage" {
  description = "Allocated storage size (in GB) for the RDS instance"
  type        = number
  default     = 20
}

variable "db_name" {
  description = "Name of the database to create on the RDS instance"
  type        = string
  default     = "platformdb"
}

variable "db_username" {
  description = "Username for database access"
  type        = string
  default     = "postgres"
}

variable "db_password" {
  description = "Password for database access"
  type        = string
  default     = "Chakratech.12345.78910"
}

variable "webserver_from_port" {
  description = "Port number for webserver ingress"
  type        = number
  default     = 443
}

variable "database_security_groups" {
  description = "List of security group IDs allowed to access the database"
  type        = list(string)
  default     = []
}
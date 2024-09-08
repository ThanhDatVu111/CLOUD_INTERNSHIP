variable "region" {
  type          = string
  default       = "us-west-2"
}

variable "domain_name"{
    type        = string
    description = "This a string of domain names"
    default     = "chakratechwork.com"
}


variable "zone_id"{
  type          = string
  description   = "AZ IDs provides a consistent way of identifying the location"
  default       = "/hostedzone/Z053448522CB81YE8FVY5"
}
variable "alternative_names"{
  type          = list
  description   = "list of alternative domain names"
  default       = []
} 

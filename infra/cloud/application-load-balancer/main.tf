provider "aws" {
  region = var.region
}

resource "aws_lb" "application" {
  name               = "app-${var.name}"
  load_balancer_type = "application"
  subnets            = var.subnets
  security_groups    = var.security_groups
}

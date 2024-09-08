terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# Configure the AWS provider
provider "aws" {
  region = var.region
}

resource "aws_acm_certificate" "vpn_server" {
  domain_name       = "*.${var.domain_name}"
  validation_method = "DNS"
  tags = {
    Name = "VPN"
  }
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate_validation" "vpn_server" {
  certificate_arn = aws_acm_certificate.vpn_server.arn
  timeouts {
    create = "5m"
  }
}

# Client VPN Endpoint resource
resource "aws_ec2_client_vpn_endpoint" "client_vpn" {
  description            = "Client VPN Endpoint"
  client_cidr_block      = var.client_cidr_block
  server_certificate_arn = aws_acm_certificate_validation.vpn_server.certificate_arn
  security_group_ids     = [aws_security_group.vpn_access.id]
  vpc_id                 = var.vpc_id
  split_tunnel           = true

  authentication_options {
    type                       = "certificate-authentication"
    root_certificate_chain_arn = var.root_certificate_chain_arn
  }

  connection_log_options {
    enabled               = false
    cloudwatch_log_group  = "${var.env}${var.cloudwatch_log_group}"
    cloudwatch_log_stream = "client_vpn-{aws_ec2_client_vpn_endpoint.client_vpn.id}"
  }

  dns_servers        = var.dns_servers
  transport_protocol = "tcp"
}


# Security Group for Client VPN (replace with your rules)
resource "aws_security_group" "vpn_access" {
  name        = "client_vpn_sg"
  description = "Security Group for Client VPN access"
  vpc_id      = var.vpc_id

  # Allow inbound traffic from anywhere on port 443
  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  # Allow all outbound traffic
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_ec2_client_vpn_network_association" "vpn_association" {
  client_vpn_endpoint_id = aws_ec2_client_vpn_endpoint.client_vpn.id
  subnet_id              = var.subnet_id
}

# VPN Authorization Rule
resource "aws_ec2_client_vpn_authorization_rule" "vpn_auth_rule" {
  client_vpn_endpoint_id = aws_ec2_client_vpn_endpoint.client_vpn.id
  target_network_cidr    = var.vpc_cidr
  authorize_all_groups   = true
}

# VPN Route
resource "aws_ec2_client_vpn_route" "vpn_route" {
  client_vpn_endpoint_id = aws_ec2_client_vpn_endpoint.client_vpn.id
  destination_cidr_block = var.vpc_cidr
  target_vpc_subnet_id   = var.subnet_id
}

# Security Group for VPN Clients
resource "aws_security_group" "vpn_sg" {
  name   = "vpn-sg"
  vpc_id = var.vpc_id

  ingress {
    from_port   = 0
    to_port     = 65535
    protocol    = "tcp"
    cidr_blocks = ["${var.client_cidr_block}"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "vpn-sg"
  }
}
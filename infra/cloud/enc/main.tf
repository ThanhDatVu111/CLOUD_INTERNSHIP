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

data "external" "decrypt_credentials" {
  program = ["bash", "-c", "sops -d ${var.encrypted_credentials}"]
}

locals {
  decrypted_credentials = jsondecode(jsonencode(data.external.decrypt_credentials.result))
}

resource "aws_instance" "example" {
  ami           = var.ami
  instance_type = var.instance_type

  user_data = templatefile("templates/userdata.tpl", {
    username = local.decrypted_credentials.username
    password = local.decrypted_credentials.password
  })
}


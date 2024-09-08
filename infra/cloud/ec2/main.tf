provider "aws" {
  region  = var.region
}

data "template_file" "user_data" {
  template = file("../scripts/user-data.yaml")
}

# Each instance provisioned by the resource block with count will have a different incrementing value for count.index - starting with zero.
# This configuration uses count.index and modulo division to assign each instance to a private subnet.
# Because the default value of instances_per_subnet is 1, Terraform will provision 1 EC2 instance per subnet.
resource "aws_instance" "go_server" {
  ami           = var.ami
  instance_type = "t2.nano"
  count = var.instances_per_subnet * length(var.subnets)
  subnet_id     = var.subnets[count.index % length(var.subnets)]

  tags = {
    Name = "backend-${count.index}"
  }

  user_data = <<EOF
  # Install Golang
  #!/bin/bash
  sudo apt update
  sudo apt install -y golang
  
  # Verify installation
  go version

  # Clone repository 
  ${data.template_file.user_data.rendered}

  cd data-platform

  # Build and run application
  make

  EOF
}
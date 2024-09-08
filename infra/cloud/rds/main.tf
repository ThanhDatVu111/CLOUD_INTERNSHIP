# configured aws provider with proper credentials
provider "aws" {
  region = var.aws_region
}

# create default vpc if one does not exist
data "aws_vpc" "default" {
 default = true
}

# create security group for the web server
resource "aws_security_group" "webserver_security_group" {
  name        = "backend"
  description = "enable http access on port 80"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    description = "http access"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "backend"
  }
}


# Create subnets
resource "aws_subnet" "subnet_a" {
  vpc_id            = data.aws_vpc.default.id
  cidr_block        = "172.31.3.0/24"  
  availability_zone = "us-west-2a"   
  map_public_ip_on_launch = true     

  tags = {
    Name = "subnet-a"
  }
}

resource "aws_subnet" "subnet_b" {
  vpc_id            = data.aws_vpc.default.id
  cidr_block        = "172.31.4.0/24"  
  availability_zone = "us-west-2b"   
  map_public_ip_on_launch = true     
  tags = {
    Name = "subnet-b"
  }
}

resource "aws_db_subnet_group" "db_subnet_group" {
  name        = "my-db-subnet-group"
  subnet_ids  = [aws_subnet.subnet_a.id, aws_subnet.subnet_b.id]
  description = "Subnet group for my RDS instance"

  tags = {
    Name = "My DB Subnet Group"
  }
}


# create security group for the database
resource "aws_security_group" "database_security_group" {
  name        = "database security group"
  description = "enable portgres access on port 5432"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    description     = "portgres access"
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.webserver_security_group.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "database_security_group"
  }
}

# create the rds instance
resource "aws_db_instance" "db_instance" {
  engine               = var.db_engine
  engine_version       = var.db_engine_version
  multi_az             = false
  identifier           = "backend-instance"
  username             = var.db_username
  password             = var.db_password
  instance_class       = var.db_instance_class
  allocated_storage    = var.db_allocated_storage
  db_name             = var.db_name
  db_subnet_group_name  = aws_db_subnet_group.db_subnet_group.name
  vpc_security_group_ids = [aws_security_group.database_security_group.id]

  backup_retention_period = 7 # Number of days to retain automated backups
  backup_window = "03:00-04:00" # Preferred UTC backup window (hh24:mi-hh24:mi format)
  maintenance_window = "mon:04:00-mon:04:30" # Preferred UTC maintenance window

  # Enable automated backups
  skip_final_snapshot = false
  }
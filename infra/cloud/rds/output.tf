output "rds_instance_endpoint" {
  value = aws_db_instance.db_instance.endpoint
}

output "rds_instance_username" {
  value = aws_db_instance.db_instance.username
}

output "rds_instance_password" {
  value = var.db_password 
}
output "id" {
  description = "Map of id(s) for the instance(s) we created"
  value       = { for k, v in aws_instance.go_server : k => v.id }
}

output "arn" {
  description = "Map of ARN(s) of the instance(s)"
  value       = { for k, v in aws_instance.go_server : k => v.arn }
}

output "ec2_instance_state" {
  description = "The state of the instance. One of: `pending`, `running`, `shutting-down`, `terminated`, `stopping`, `stopped`"
  value       = { for k, v in aws_instance.go_server : k => v.instance_state }
}

output "public_ip" {
  description = "The Public IP address of the instance."
  value       = { for k, v in aws_instance.go_server : k => v.public_ip }
}

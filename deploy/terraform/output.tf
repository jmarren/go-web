output "app_server_ip" {
  description = "ip address of app server instance"
  value       = aws_instance.app_server.public_ip
}

output "db_server_ip" {
  description = "ip address of db server instance"
  value       = aws_instance.db_server.public_ip
}


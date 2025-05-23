
output "playful_1_instance_ip" {
  description = "ip address of playful 1 instance"
  value       = aws_instance.playful_1_instance.public_ip
}

# output "ssh_key_content" {
#   value = data.local_file.playful_pub_key.content
# }
#

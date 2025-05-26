
# security group for all ec2 instances
resource "aws_security_group" "all_ec2" {
  vpc_id = aws_vpc.vpc.id
  name   = "${var.app_name}_app_server_security_group"

  lifecycle {
    create_before_destroy = true
  }
}

# security group for db server
resource "aws_security_group" "db" {
  vpc_id = aws_vpc.vpc.id
  name   = "${var.app_name}_db_server_security_group"
}

# allow ssh from my IP
resource "aws_vpc_security_group_ingress_rule" "allow_ssh_from_my_ip" {
  security_group_id = aws_security_group.all_ec2.id
  cidr_ipv4         = "${chomp(data.http.my_ip.body)}/32"
  from_port         = 22
  ip_protocol       = "tcp"
  to_port           = 22
}

# allow connections on port 5432 from my IP
resource "aws_vpc_security_group_ingress_rule" "allow_5432" {
  security_group_id = aws_security_group.db.id
  cidr_ipv4         = "${chomp(data.http.my_ip.body)}/32"
  from_port         = 5432
  ip_protocol       = "tcp"
  to_port           = 5432
}

# allow egress on all ports 
resource "aws_vpc_security_group_egress_rule" "allow_all" {
  security_group_id = aws_security_group.all_ec2.id
  from_port         = 0
  to_port           = 0
  ip_protocol       = -1
  cidr_ipv4         = "0.0.0.0/0"
}

# resource "aws_vpc_security_group_ingress_rule" "allow_connection_on_5432" {
#   security_group_id = aws_security_group.ec2_security_group.id
#   cidr_ipv4         = "${chomp(data.http.my_ip.body)}/32"
#   from_port         = 22
#   ip_protocol       = "tcp"
#   to_port           = 22
# }


# ingress {
#   from_port = 22
#   to_port   = 22
#   protocol  = "tcp"
#   # cidr_blocks = ["0.0.0.0/0"]
#   # security_groups = [aws_security_group.loop_jump_box_security_group.id]
#   cidr_blocks = ["${chomp(data.http.my_ip.body)}/32"]
# }
#
# egress {
#   from_port   = 0
#   to_port     = 0
#   protocol    = -1
#   cidr_blocks = ["0.0.0.0/0"]
# }



# resource "aws_security_group" "ec2_security_group" {
#   vpc_id = aws_vpc.vpc.id
#   name   = "${var.app_name}_app_server_ec2_security_group"
#
#   ingress {
#     from_port = 22
#     to_port   = 22
#     protocol  = "tcp"
#     # cidr_blocks = ["0.0.0.0/0"]
#     # security_groups = [aws_security_group.loop_jump_box_security_group.id]
#     cidr_blocks = ["${chomp(data.http.my_ip.body)}/32"]
#   }
#
#   egress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = -1
#     cidr_blocks = ["0.0.0.0/0"]
#   }
# }



# egress {
#   from_port   = 0
#   to_port     = 0
#   protocol    = -1
#   cidr_blocks = ["0.0.0.0/0"]
# }


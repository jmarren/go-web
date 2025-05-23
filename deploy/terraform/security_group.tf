resource "aws_security_group" "playful_ec2_security_group" {
  vpc_id = aws_vpc.playful_vpc.id
  name   = "playful_ec2_security_group"

  ingress {
    from_port = 22
    to_port   = 22
    protocol  = "tcp"
    # cidr_blocks = ["0.0.0.0/0"]
    # security_groups = [aws_security_group.loop_jump_box_security_group.id]
    cidr_blocks = ["${chomp(data.http.my_ip.body)}/32"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = -1
    cidr_blocks = ["0.0.0.0/0"]
  }
}

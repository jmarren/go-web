resource "aws_instance" "playful_1_instance" {
  ami           = "ami-084568db4383264d4"
  instance_type = "t2.micro"
  subnet_id     = aws_subnet.playful_public_east_1a.id
  key_name      = aws_key_pair.playful_1_key_pair.key_name


  vpc_security_group_ids = [aws_security_group.playful_ec2_security_group.id]

  tags = {
    Name = "playful_instance"
  }
}




resource "aws_key_pair" "playful_1_key_pair" {
  key_name   = "playful_1_key_pair"
  public_key = data.local_file.playful_pub_key.content
}




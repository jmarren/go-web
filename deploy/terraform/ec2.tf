resource "aws_instance" "app_server" {
  ami           = "ami-084568db4383264d4"
  instance_type = "t2.micro"
  subnet_id     = aws_subnet.subnet_public_east_1a.id
  key_name      = aws_key_pair.main.key_name


  vpc_security_group_ids = [aws_security_group.all_ec2.id]

  tags = {
    Name = "${var.app_name}_app_server"
  }
}


resource "aws_instance" "db_server" {
  ami           = "ami-084568db4383264d4"
  instance_type = "t2.micro"
  subnet_id     = aws_subnet.subnet_public_east_1a.id
  key_name      = aws_key_pair.main.key_name


  vpc_security_group_ids = [aws_security_group.all_ec2.id, aws_security_group.db.id]

  tags = {
    Name = "${var.app_name}_db_server"
  }
}



resource "aws_key_pair" "main" {
  key_name   = "${var.app_name}_public_key"
  public_key = data.local_file.public_key.content
}




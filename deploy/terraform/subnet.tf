
resource "aws_subnet" "subnet_public_east_1a" {
  vpc_id            = aws_vpc.vpc.id
  cidr_block        = "10.0.4.0/24"
  availability_zone = "us-east-1a"

  map_public_ip_on_launch = true

  tags = {
    Name = "${var.app_name}_subnet_public_east_1a"
  }
}

resource "aws_route_table" "public_subnet_route_table" {
  vpc_id = aws_vpc.vpc.id

  # route to Internet Gateway
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }


  tags = {
    Name = "${var.app_name}_public_subnet_route_table"
  }
}


resource "aws_route_table_association" "public_east_1a_route_table_assoc" {
  subnet_id      = aws_subnet.subnet_public_east_1a.id
  route_table_id = aws_route_table.public_subnet_route_table.id
}


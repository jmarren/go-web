
resource "aws_internet_gateway" "playful_igw" {
  vpc_id = aws_vpc.playful_vpc.id

  tags = {
    Name = "playful_igw"
  }
}

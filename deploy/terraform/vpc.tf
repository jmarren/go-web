resource "aws_vpc" "playful_vpc" {
  cidr_block       = var.vpc_cidr_block
  instance_tenancy = "default"

  tags = {
    Name = "playful_vpc"
  }
}


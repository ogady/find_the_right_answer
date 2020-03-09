resource "aws_vpc" "vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
  tags = {
    Name = "FTRA_VPC"
  }
}

resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.vpc.id
  tags = {
    Name = "FTRA_IGW"
  }
}


resource "aws_eip" "ngw_eip_0" {
  vpc        = true
  depends_on = [aws_internet_gateway.igw]
  tags = {
    Name = "FTRA_EIP_0"
  }
}

resource "aws_eip" "ngw_eip_1" {
  vpc        = true
  depends_on = [aws_internet_gateway.igw]
  tags = {
    Name = "FTRA_EIP_1"
  }
}

resource "aws_nat_gateway" "ngw_0" {
  allocation_id = aws_eip.ngw_eip_0.id
  subnet_id     = aws_subnet.pub_subnet_0.id
  depends_on    = [aws_internet_gateway.igw]
  tags = {
    Name = "FTRA_NGW_0"
  }
}

resource "aws_nat_gateway" "ngw_1" {
  allocation_id = aws_eip.ngw_eip_1.id
  subnet_id     = aws_subnet.pub_subnet_1.id
  depends_on    = [aws_internet_gateway.igw]
  tags = {
    Name = "FTRA_NGW_1"
  }
}


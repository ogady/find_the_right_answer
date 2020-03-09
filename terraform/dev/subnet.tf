resource "aws_subnet" "pub_subnet_0" {
  vpc_id                  = aws_vpc.vpc.id
  cidr_block              = "10.0.1.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "ap-northeast-1a"
  tags = {
    Name = "FTRA_PUB_SUB_0"
  }
}

resource "aws_subnet" "pub_subnet_1" {
  vpc_id                  = aws_vpc.vpc.id
  cidr_block              = "10.0.2.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "ap-northeast-1c"
  tags = {
    Name = "FTRA_PUB_SUB_1"
  }
}

resource "aws_route_table" "pub_rtbl" {
  vpc_id = aws_vpc.vpc.id
  tags = {
    Name = "FTRA_PUB_RT"
  }
}


resource "aws_route_table_association" "pub_sub_0_route" {
  subnet_id      = aws_subnet.pub_subnet_0.id
  route_table_id = aws_route_table.pub_rtbl.id

}

resource "aws_route_table_association" "pub_sub_1_route" {
  subnet_id      = aws_subnet.pub_subnet_1.id
  route_table_id = aws_route_table.pub_rtbl.id

}

resource "aws_subnet" "pri_subnet_0" {
  vpc_id                  = aws_vpc.vpc.id
  cidr_block              = "10.0.65.0/24"
  map_public_ip_on_launch = false
  availability_zone       = "ap-northeast-1a"
  tags = {
    Name = "FTRA_PRI_SUB_0"
  }
}

resource "aws_subnet" "pri_subnet_1" {
  vpc_id                  = aws_vpc.vpc.id
  cidr_block              = "10.0.66.0/24"
  map_public_ip_on_launch = false
  availability_zone       = "ap-northeast-1a"
  tags = {
    Name = "FTRA_PRI_SUB_1"
  }
}

resource "aws_route" "pub_route" {
  route_table_id         = aws_route_table.pub_rtbl.id
  gateway_id             = aws_internet_gateway.igw.id
  destination_cidr_block = "0.0.0.0/0"
}

resource "aws_route_table" "pri_rtbl_0" {
  vpc_id = aws_vpc.vpc.id
  tags = {
    Name = "FTRA_PRI_RT_0"
  }
}

resource "aws_route_table" "pri_rtbl_1" {
  vpc_id = aws_vpc.vpc.id
  tags = {
    Name = "FTRA_PRI_RT_1"
  }
}

resource "aws_route" "pri_route_0" {
  route_table_id         = aws_route_table.pri_rtbl_0.id
  gateway_id             = aws_nat_gateway.ngw_0.id
  destination_cidr_block = "0.0.0.0/0"

}

resource "aws_route" "pri_route_1" {
  route_table_id         = aws_route_table.pri_rtbl_1.id
  gateway_id             = aws_nat_gateway.ngw_1.id
  destination_cidr_block = "0.0.0.0/0"

}

resource "aws_route_table_association" "pri_sub_0_route" {
  subnet_id      = aws_subnet.pri_subnet_0.id
  route_table_id = aws_route_table.pri_rtbl_0.id

}


resource "aws_route_table_association" "pri_sub_1_route" {
  subnet_id      = aws_subnet.pri_subnet_1.id
  route_table_id = aws_route_table.pri_rtbl_1.id

}

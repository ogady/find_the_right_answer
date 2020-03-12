variable "env" {}
variable "vpc_id" {}
variable "vpc_cidr" {}

resource "aws_internet_gateway" "igw" {
  vpc_id = var.vpc_id
  tags = {
    Name = "FTRA_${var.env}_IGW"
  }
}

resource "aws_eip" "ngw_eip_0" {
  vpc        = true
  depends_on = [aws_internet_gateway.igw]
  tags = {
    Name = "FTRA_${var.env}_EIP_0"
  }
}

resource "aws_eip" "ngw_eip_1" {
  vpc        = true
  depends_on = [aws_internet_gateway.igw]
  tags = {
    Name = "FTRA_${var.env}_EIP_1"
  }
}


# private subnet
resource "aws_subnet" "pri_subnet_0" {
  vpc_id                  = var.vpc_id
  cidr_block              = cidrsubnet(var.vpc_cidr, 8, 10)
  map_public_ip_on_launch = false
  availability_zone       = "ap-northeast-1a"
  tags = {
    Name = "FTRA_${var.env}_PRI_SUB_0"
  }
}

resource "aws_subnet" "pri_subnet_1" {
  vpc_id                  = var.vpc_id
  cidr_block              = cidrsubnet(var.vpc_cidr, 8, 11)
  map_public_ip_on_launch = false
  availability_zone       = "ap-northeast-1c"
  tags = {
    Name = "FTRA_${var.env}_PRI_SUB_1"
  }
}

resource "aws_route_table" "pri_rtbl_0" {
  vpc_id = var.vpc_id
  tags = {
    Name = "FTRA_${var.env}_PRI_RT_0"
  }
}

resource "aws_route_table" "pri_rtbl_1" {
  vpc_id = var.vpc_id
  tags = {
    Name = "FTRA_${var.env}_PRI_RT_1"
  }
}

resource "aws_route" "pri_route_0" {
  route_table_id         = aws_route_table.pri_rtbl_0.id
  nat_gateway_id         = aws_nat_gateway.ngw_0.id
  destination_cidr_block = "0.0.0.0/0"

}

resource "aws_route" "pri_route_1" {
  route_table_id         = aws_route_table.pri_rtbl_1.id
  nat_gateway_id         = aws_nat_gateway.ngw_1.id
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

# nat gateway
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

# public subnet
resource "aws_subnet" "pub_subnet_0" {
  vpc_id                  = var.vpc_id
  cidr_block              = cidrsubnet(var.vpc_cidr, 8, 1)
  map_public_ip_on_launch = true
  availability_zone       = "ap-northeast-1a"
  tags = {
    Name = "FTRA_${var.env}_PUB_SUB_0"
  }
}

resource "aws_subnet" "pub_subnet_1" {
  vpc_id                  = var.vpc_id
  cidr_block              = cidrsubnet(var.vpc_cidr, 8, 2)
  map_public_ip_on_launch = true
  availability_zone       = "ap-northeast-1c"
  tags = {
    Name = "FTRA_${var.env}_PUB_SUB_1"
  }
}

resource "aws_route_table" "pub_rtbl" {
  vpc_id = var.vpc_id
  tags = {
    Name = "FTRA_${var.env}_PUB_RT"
  }
}

resource "aws_route" "pub_route" {
  route_table_id         = aws_route_table.pub_rtbl.id
  gateway_id             = aws_internet_gateway.igw.id
  destination_cidr_block = "0.0.0.0/0"
}

resource "aws_route_table_association" "pub_sub_0_route" {
  subnet_id      = aws_subnet.pub_subnet_0.id
  route_table_id = aws_route_table.pub_rtbl.id

}

resource "aws_route_table_association" "pub_sub_1_route" {
  subnet_id      = aws_subnet.pub_subnet_1.id
  route_table_id = aws_route_table.pub_rtbl.id

}


output "pub_sub_0_id" {
  value = aws_subnet.pub_subnet_0.id
}

output "pub_sub_1_id" {
  value = aws_subnet.pub_subnet_1.id
}

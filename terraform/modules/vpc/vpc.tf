variable "env" {}

resource "aws_vpc" "vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
  tags = {
    Name = "FTRA_${var.env}_VPC"
  }
}

resource "aws_vpc_endpoint" "dynamodb" {
  vpc_id       = aws_vpc.vpc.id
  service_name = "com.amazonaws.ap-northeast-1.dynamodb"

  tags = {
    Name = "FTRA_${var.env}_VPC_EP"
  }
}

output "vpc_id" {
  value = aws_vpc.vpc.id
}

output "vpc_cidr" {
  value = aws_vpc.vpc.cidr_block
}

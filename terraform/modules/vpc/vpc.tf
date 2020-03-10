variable "env" {}

resource "aws_vpc" "vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true
  tags = {
    Name = "FTRA_${var.env}_VPC"
  }
}


output "vpc_id" {
  value = aws_vpc.vpc.id
}

output "vpc_cidr" {
  value = aws_vpc.vpc.cidr_block
}

provider "aws" {
  region = "ap-northeast-1"
}

terraform {
  backend "s3" {
    bucket = "find-the-right-answer-prod-tfstate"
    key    = "prod/terraform.tfstate"
    region = "ap-northeast-1"
  }
}

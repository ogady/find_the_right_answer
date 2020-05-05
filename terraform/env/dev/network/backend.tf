provider "aws" {
  region = "ap-northeast-1"
}

terraform {
  backend "s3" {
    bucket = "find-the-right-answer-dev-tfstate"
    key    = "dev/network/terraform.tfstate"
    region = "ap-northeast-1"
  }
  # backend "local" {
  #   path = "terraform.tfstate"
  # }
}

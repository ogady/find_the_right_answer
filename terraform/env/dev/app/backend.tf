provider "aws" {
  region = "ap-northeast-1"
}

terraform {
  backend "s3" {
    bucket = "ftra-dev-tfstates"
    key    = "dev/app/terraform.tfstate"
    region = "ap-northeast-1"
  }

  # backend "local" {
  #   path = "terraform.tfstate"
  # }
}

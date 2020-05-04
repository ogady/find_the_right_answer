provider "aws" {
  region = "ap-northeast-1"
}

terraform {
  backend "s3" {
    bucket = "find-the-right-answer-ui-dev-tfstate"
    key    = "dev/ui/terraform.tfstate"
    region = "ap-northeast-1"
  }
}

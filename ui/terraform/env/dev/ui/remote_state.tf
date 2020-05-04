data "terraform_remote_state" "network" {
  backend = "s3"

  config = {
    bucket = "find-the-right-answer-dev-tfstate"
    key    = "dev/network/terraform.tfstate"
    region = "ap-northeast-1"
  }
}

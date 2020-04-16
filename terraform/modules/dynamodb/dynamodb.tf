variable "env" {}

resource "aws_dynamodb_table" "dynamodb_topic" {
  name           = "topic"
  billing_mode   = "PAY_PER_REQUEST"
  read_capacity  = 4
  write_capacity = 4
  hash_key       = "StartChar"
  range_key      = "TopicPiece"

  attribute {
    name = "StartChar"
    type = "S"
  }

  attribute {
    name = "TopicPiece"
    type = "S"
  }


  tags = {
    Name = "FTRA_${var.env}_Topic"
  }
}

resource "aws_dynamodb_table" "dynamodb_topicPiece" {
  name           = "topic_piece"
  billing_mode   = "PROVISIONED"
  read_capacity  = 4
  write_capacity = 4
  hash_key       = "TopicPiece"

  attribute {
    name = "TopicPiece"
    type = "S"
  }


  tags = {
    Name = "FTRA_${var.env}_TopicPiece"
  }
}

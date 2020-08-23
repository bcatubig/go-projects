resource "aws_dynamodb_table" "url" {
  hash_key = "ShortID"
  name     = "url-shortener"

  attribute {
    name = "ShortID"
    type = "S"
  }

  attribute {
    name = "Url"
    type = "S"
  }

  tags = merge(local.tags, {
    Name = "dynamodb-url-shortener"
  })
}

terraform {
  required_version = "~> 0.13.0"
}

provider "aws" {
  version = "~> 3.3.0"
}

locals {
  tags = {
    Owner       = "brandon@notbsd.com"
    Environment = "development"
  }
}

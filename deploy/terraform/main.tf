terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
  # alias  = "northern-california"
}

data "http" "my_ip" {
  url = "http://checkip.amazonaws.com/"
}


data "local_file" "playful_pub_key" {
  filename = "${path.cwd}/deploy/keys/key.pub"
}



data "aws_caller_identity" "current" {}



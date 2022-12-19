# terraform {
#   backend "s3" {
#     bucket         = "nassau-state"
#     dynamodb_table = "nassau-state"
#     key            = "services/upload-api.tfstate"
#     region         = "us-east-1"
#     encrypt        = true
#   }
# }

provider "aws" {
  region              = local.region
  allowed_account_ids = ["895216607862"]

  default_tags {
    tags = local.tags
  }
}

locals {
  region = "us-east-1"
  name   = "upload-api"
  env    = terraform.workspace

  tags = {
    TF-Managed = true
    Env        = local.env
  }
}

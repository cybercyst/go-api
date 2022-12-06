terraform {
  backend "s3" {
    bucket         = "terraform-state-ddp-infra"
    key            = "services/go-api.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "terraform-state-ddp-infra-lock"
  }
}

provider "aws" {
  region              = local.region
  allowed_account_ids = ["882892008441"]

  default_tags {
    tags = local.tags
  }
}

locals {
  region = "us-east-1"
  name   = "go-api"
  env    = terraform.workspace

  tags = {
    TF-Managed = true
    Env        = local.env
  }
}

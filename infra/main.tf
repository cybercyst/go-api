terraform {
  required_providers {

    aws = {
      source  = "hashicorp/aws"
      version = "4.48.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "3.4.3"
    }
  }
}

provider "aws" {
  region              = local.region
  allowed_account_ids = ["895216607862"]

  default_tags {
    tags = local.tags
  }
}

locals {
  region = "us-east-1"
  env    = terraform.workspace
  name   = "upload-api-${local.env}"

  tags = {
    TF-Managed = true
    Env        = local.env
  }
}

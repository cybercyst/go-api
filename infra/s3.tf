resource "aws_s3_bucket" "_" {
  bucket = "${local.name}-${random_string.random.result}"
  # This will allow us to delete a bucket with files in it
  force_destroy = true
}

resource "aws_s3_bucket_acl" "_" {
  bucket = aws_s3_bucket._.id
  acl    = "private"
}

resource "random_string" "random" {
  length  = 4
  lower   = true
  numeric = false
  special = false
  upper   = false
}

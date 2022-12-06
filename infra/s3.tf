resource "aws_s3_bucket" "_" {
  bucket = "go-api-images-${local.env}"
}

resource "aws_s3_bucket_acl" "_" {
  bucket = aws_s3_bucket._.id
  acl    = "private"
}

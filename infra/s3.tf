resource "aws_s3_bucket" "_" {
  bucket = local.name
  # This will allow us to delete a bucket with files in it
  force_destroy = true
}

resource "aws_s3_bucket_acl" "_" {
  bucket = aws_s3_bucket._.id
  acl    = "private"
}

terraform {
  required_version = ">= v1.5.5"
}

module "bucket" {
  source  = "terraform-aws-modules/s3-bucket/aws"
  version = "4.11.0"

  bucket = var.name

  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false

  control_object_ownership = true
  object_ownership         = "ObjectWriter"

  versioning = {
    status = "Disabled"
  }

  website = {
    index_document = "index.html"
  }

  attach_policy = true
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Sid    = "PublicReadForGetBucketObjects",
        Effect = "Allow",
        Principal = {
          AWS = "*"
        },
        Action   = "s3:GetObject",
        Resource = "arn:aws:s3:::${var.name}/*"
      }
    ]
  })

  grant = var.grant
  owner = var.owner
}

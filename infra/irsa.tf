data "aws_iam_policy_document" "write_policy" {
  statement {
    actions = [
      "s3:PutObject",
      "s3:GetObject"
    ]

    resources = [
      "${aws_s3_bucket._.arn}/*"
    ]
  }
}

resource "aws_iam_policy" "write_policy" {
  name   = "${local.name}-write-policy"
  policy = data.aws_iam_policy_document.write_policy.json
}

module "terraform_irsa_role" {
  source = "terraform-aws-modules/iam/aws//modules/iam-role-for-service-accounts-eks"

  role_name = "${local.name}-irsa"

  oidc_providers = {
    main = {
      provider_arn               = var.eks_oidc_provider_arn
      namespace_service_accounts = var.namespace_service_accounts
    }
  }

  role_policy_arns = {
    WriteBucket = aws_iam_policy.write_policy.arn
  }
}

variable "eks_oidc_provider_arn" {
  type        = string
  description = "The ARN for EKS OIDC provider"
}

variable "namespace_service_accounts" {
    type = list(string)
    description = "The namespaced service accounts to attach the IRSA to"
}
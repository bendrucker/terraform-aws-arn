variable "partition" {
  type        = string
  description = "The partition in which the resource is located. A partition is a group of AWS Regions."
}

variable "service" {
  type        = string
  description = "The service namespace that identifies the AWS product. For example, s3 for Amazon S3 resources."
}

variable "region" {
  type        = string
  description = "The Region. For example, us-east-2 for US East (Ohio)."
}

variable "account_id" {
  type        = string
  description = "The ID of the AWS account that owns the resource, without the hyphens. For example, 123456789012."
}

variable "resource_id" {
  type        = string
  description = "The resource identifier. This part of the ARN can be the name or ID of the resource or a resource path. For example, user/Bob for an IAM user or instance/i-1234567890abcdef0 for an EC2 instance."
}

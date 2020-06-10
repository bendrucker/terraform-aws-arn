# terraform-aws-arn [![tests workflow status](https://github.com/bendrucker/terraform-aws-arn/workflows/tests/badge.svg?branch=master)](https://github.com/bendrucker/terraform-aws-arn/actions?query=workflow%3Atests) [![terraform module](https://img.shields.io/badge/terraform-module-623CE4)](https://registry.terraform.io/modules/bendrucker/arn/aws)

> Terraform module that accepts the components of an AWS ARN and outputs the string form

This module provides the inverse to [`data.aws_arn`](https://www.terraform.io/docs/providers/aws/d/arn.html). Where possible, prefer refererring to an ARN directly (as an attribute of a resource or data source) rather than constructing it. 

Use this module to construct:

* ARNs for resources that do not yet exist but will be created in the future
* ARNs with wildcard segments (for use in IAM policies)

## Usage

Only the `service` and `resource_id` variables are required, since all ARNs include them. Variables `partition` (default: "aws"), `account_id` (default: ""), and `region` (default: "") are optional. 

```tf
module "object" {
  source = "bendrucker/arn/aws"

  service     = "s3"
  resource_id = "my-bucket/my-object"
}

output "object_arn" {
  value = module.object.arn
}
```

## Testing

```sh
go test -v ./...
```

## License

MIT © [Ben Drucker](http://bendrucker.me)

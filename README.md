# Website Bucket Module
## Pipeline Status

[![.github/workflows/pipeline.yml](https://github.com/armakuni/terraform-aws-static-website-bucket/actions/workflows/pipeline.yml/badge.svg?branch=main)](https://github.com/armakuni/terraform-aws-static-website-bucket/actions/workflows/pipeline.yml)

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= v1.5.5 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_bucket"></a> [bucket](#module\_bucket) | terraform-aws-modules/s3-bucket/aws | 4.0.1 |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_grant"></a> [grant](#input\_grant) | An ACL policy grant. Conflicts with `acl` | `any` | `[]` | no |
| <a name="input_name"></a> [name](#input\_name) | The name for the bucket | `string` | n/a | yes |
| <a name="input_owner"></a> [owner](#input\_owner) | Bucket owner's display name and ID. Conflicts with `acl` | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_s3_bucket_id"></a> [s3\_bucket\_id](#output\_s3\_bucket\_id) | The bucket ID |
<!-- END_TF_DOCS -->

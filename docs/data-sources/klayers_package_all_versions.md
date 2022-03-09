# klayers_package_latest_version Data Source

Use this data source to get the access to the last AWS arn of your (Klayers lambda layer)[https://github.com/keithrozario/Klayers].

## Example Usage
```hcl
data "klayers_package_all_versions" "pandas" {
  name           = "requests"
  region         = "eu-west-1"
  python_version = "3.8"
}

output "klayers_package_all_versions_arns" {
  value = data.klayers_package_all_versions.pandas.arns
}
```

## Argument Reference
* `name` - (Required) The name of the layer.
* `region` - (Required) The AWS region the layer resides in.
* `python version` - (Optional) The python version. Default : 3.9

## Attribute Reference
* `arns` - List of Amazon Resource Name (ARNs) identifying the layers.
* `api` - The api hostname.
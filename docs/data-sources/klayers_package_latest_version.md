# klayers_package_latest_version Data Source

Use this data source to get the access to the last AWS arn of your [Klayers lambda layer](https://github.com/keithrozario/Klayers).

## Example Usage
```hcl
data "klayers_package_latest_version" "pandas" {
  name           = "pandas"
  region         = "eu-west-1"
  python_version = "3.8"
}

output "klayers_package_latest_version_arn" {
  value = data.klayers_package_latest_version.pandas.arn
}
```

## Argument Reference
* `name` - (Required) The name of the layer.
* `region` - (Required) The AWS region the layer resides in.
* `python version` - (Optional) The python version. Default : 3.9

## Attribute Reference
* `arn` - Amazon Resource Name (ARN) identifying the layer.
* `api` - The api hostname.

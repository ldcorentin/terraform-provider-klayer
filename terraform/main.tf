data "klayers_package_latest_version" "pandas" {
  name   = "pandas"
  region = "eu-west-1"
}

data "klayers_package_latest_version" "requests" {
  name           = "requests"
  region         = "eu-west-1"
  python_version = "3.8"
}

data "klayers_package_all_versions" "requests" {
  name           = "requests"
  region         = "eu-west-1"
  python_version = "3.8"
}

output "klayers_pandas_latest_package_arn" {
  value = data.klayers_package_latest_version.pandas.arn
}

output "klayers_requests_latest_package_arn" {
  value = data.klayers_package_latest_version.requests.arn
}

output "klayers_requests_versions_package_arn" {
  value = data.klayers_package_all_versions.requests.arns
}

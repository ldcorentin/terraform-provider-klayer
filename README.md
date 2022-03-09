# [terraform-provider-klayers](https://registry.terraform.io/providers/ldcorentin/klayers/latest)

A terraform provider to manage your klayers (https://github.com/keithrozario/Klayers).

## Example
```terraform
terraform {
  required_providers {
    klayer = {
      version = "~> 1.0.0"
      source  = "ldcorentin/klayer"
    }
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

data "klayers_package_latest_version" "pandas" {
  name   = "pandas"
  region = "us-east-1"
}

resource "aws_lambda_function" "test_lambda" {
  filename      = "index.zip"
  function_name = "klayers-test"
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "index.handler"
  runtime       = "python3.9"
  layers = [
    data.klayers_package_latest_version.pandas
  ]
}
```

You will find two klayers providers. The `klayers` one is deprecated because it is not synchronize anymore with the repo (I have no answer from the support to resynchromize it).

The right provider to use is : `klayer`!
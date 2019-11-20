# Terraform Provider - Active Directory

[![GolangCI](https://golangci.com/badges/github.com/golangci/golangci-lint.svg)](https://golangci.com)
![CircleCI](https://img.shields.io/circleci/build/github/adlerrobert/terraform-provider-activedirectory?style=flat-square&logo=circleci&label=CircleCI)
[![codecov](https://codecov.io/gh/adlerrobert/terraform-provider-activedirectory/branch/master/graph/badge.svg)](https://codecov.io/gh/adlerrobert/terraform-provider-activedirectory)
[![GitHub license](https://img.shields.io/github/license/adlerrobert/terraform-provider-activedirectory.svg?style=flat-square&cacheSeconds=3600)](https://github.com/adlerrobert/terraform-provider-activedirectory/blob/master/LICENSE)

[![GitHub release](https://img.shields.io/github/release/adlerrobert/terraform-provider-activedirectory.svg?style=flat-square&cacheSeconds=3600)](https://GitHub.com/adlerrobert/terraform-provider-activedirectory/releases/)
[![GitHub tag](https://img.shields.io/github/tag/adlerrobert/terraform-provider-activedirectory.svg?style=flat-square&cacheSeconds=3600)](https://github.com/adlerrobert/terraform-provider-activedirectory/tags/)

This is a Terraform  Provider to work with Active Directory.

This provider currently supports only computer objects, but more active directory resources are planned. Please feel free to contribute.

For general information about Terraform, visit the [official website][3] and the [GitHub project page][4].

[3]: https://terraform.io/
[4]: https://github.com/hashicorp/terraform

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12+
- [Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)

## Developing the Provider
TODO

## Using the Provider
The provider is useful for adding and managing computer objects in Active Directory.
### Example
```hcl
# Configure the AD Provider
provider "activedirectory" {
  ad_host     = "ad.example.org"
  ad_port     = 389
  use_tls       = true
  bind_user     = "cn=admin,dc=example,dc=org"
  bind_password = "admin"
}

# Add computer to Active Directory
resource "activedirectory_computer" "foo" {
  name           = "TestComputerTF"                       # update will force destroy and new
  ou             = "CN=Computers,DC=example,DC=org"       # can be updated
  description    = "terraform sample server"              # can be updated
}
```

### Updating Dependencies
```console
$ go get github.com/adlerrobert/terraform-provider-activedirectory
$ go mod tidy
$ go mod vendor
```

## Testing the Provider
TODO

## Contributing
TODO

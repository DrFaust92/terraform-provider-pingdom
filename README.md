# terraform-provider-pingdom

This project is a [terraform](http://www.terraform.io/) provider for [pingdom](https://www.pingdom.com/).

This currently only supports working with basic HTTP and ping checks.

This supports Pingdom API v3.1: [API reference docs](https://docs.pingdom.com/api/)

## Requirements

- Terraform 0.12.x
- Go 1.14 (to build the provider plugin)

## Usage

### Use provider

```hcl
terraform {
  required_providers {
    pingdom = {
      source  = "drfaust92/pingdom"
      version = "~> 1.2"
    }
  }
}

variable "pingdom_api_token" {}
variable "solarwinds_user" {}
variable "solarwinds_passwd" {}

provider "pingdom" {
    api_token         = "${var.pingdom_api_token}"
    solarwinds_user   = "${var.solarwinds_user}"
    solarwinds_passwd = "${var.solarwinds_passwd}"
}
```

Alternatively, you can use environment variables to avoid sensitive values in your code:

```sh
export PINGDOM_API_TOKEN=YOUR_API_TOKEN
export SOLARWINDS_PASSWD=YOUR_SOLARWINDS_PASSWD
```

You don't need to declare those values in your code:

```hcl
terraform {
  required_providers {
    pingdom = {
      source  = "drfaust92/pingdom"
      version = "~> 1.2"
    }
  }
}

variable "solarwinds_user" {}

provider "pingdom" {
    solarwinds_user   = "${var.solarwinds_user}"
}
```

### Basic Check

```hcl
resource "pingdom_check" "example" {
    type        = "http"
    name        = "my http check"
    host        = "example.com"
    resolution  = 5
}

resource "pingdom_check" "example_with_alert" {
  type                     = "http"
  name                     = "my http check"
  host                     = "example.com"
  resolution               = 5
  sendnotificationwhendown = 2 # alert after 5 mins, with resolution 5*(2-1)
  integrationids = [
    12345678,
    23456789
  ]
  userids = [
    24680,
    13579
  ]
}

resource "pingdom_check" "ping_example" {
  type       = "ping"
  name       = "my ping check"
  host       = "example.com"
  resolution = 1
  userids = [
    24680
  ]
}
```

Apply with:

```sh
 terraform apply \
    -var 'pingdom_api_token=YOUR_API_TOKEN'
    -var 'solarwinds_user=YOUR_SOLARWINDS_USER'
    -var 'solarwinds_passwd=YOUR_SOLARWINDS_PASSWD'
```

### Using attributes from other resources

```hcl
variable "heroku_email" {}
variable "heroku_api_key" {}

variable "pingdom_api_token" {}
variable "solarwinds_user" {}
variable "solarwinds_passwd" {}



provider "heroku" {
  email   = var.heroku_email
  api_key = var.heroku_api_key
}

provider "pingdom" {
  api_token         = var.pingdom_api_token
  solarwinds_user   = var.solarwinds_user
  solarwinds_passwd = var.solarwinds_passwd
}

resource "heroku_app" "example" {
  name   = "my-app"
  region = "us"
}

resource "pingdom_check" "example" {
  name       = "my check"
  host       = heroku_app.example.heroku_hostname
  resolution = 5
}
```

### Teams

```hcl
resource "pingdom_team" "test" {
  name = "The Test team"
  member_ids = [
    pingdom_contact.first_contact.id,
  ]
}
```

### Contacts

Note that all contacts _must_ have both a high and low severity notification

```hcl
resource "pingdom_contact" "first_contact" {
  name = "johndoe"

  sms_notification {
    number   = "5555555555"
    severity = "HIGH"
  }

  sms_notification {
    number       = "3333333333"
    country_code = "91"
    severity     = "LOW"
    provider     = "esendex"
  }

  email_notification {
    address  = "test@test.com"
    severity = "LOW"
  }
}

resource "pingdom_contact" "second_contact" {
  name   = "janedoe"
  paused = true

  email_notification {
    address  = "test@test.com"
    severity = "LOW"
  }

  email_notification {
    address  = "test@test.com"
    severity = "HIGH"
  }
}
```

### Maintenance

The maintenance resource is used to define a one time or repetitive maintenance window and bond the maintenance window with one or more uptime and tms checks.

```hcl
resource "pingdom_check" "test" {
  name = "test-check"
  host = "www.example.com"
  type = "http"
}

resource "pingdom_maintenance" "test" {
  description    = "test-maintenance"
  from           = 2717878693
  to             = 2718878693
  effectiveto    = 2718978693
  recurrencetype = "week"
  repeatevery    = 4
  uptimeids      = [pingdom_check.test.id]
}
```

### Occurrence

An occurrence resource usually represents a group of maintenance occurrences, as determined by the triple
(maintenanceid, effective_from, effective_to). This triple is effectively a query against all existing
maintenance occurrences. Please note that 'effective_from' and 'effective_to' are different from the attributes
pair 'from' and 'to' of maintenance/occurrence. The latter are used to specify the start/end within a maintenance
cycle, while the former are purely query conditions used to retrieve occurrence objects.

It is not possible to import occurrences as they are queries, which does not exist on Pingdom.

'from' and 'to' attributes can be updated, which results in all occurrences matched by the query being
updated. In other words, all occurrences matched by a single resource will share the same values for `from` and `to`.

```hcl
resource "pingdom_occurrence" "test" {
  maintenance_id = pingdom_maintenance.test.id
  effective_from = pingdom_maintenance.test.from
  effective_to   = pingdom_maintenance.test.effectiveto
  from           = pingdom_maintenance.test.from
  to             = "2021-04-10T22:00:00+08:00"
}
```

## Documentation

Full documentation for every resource and data source is on the [Terraform Registry](https://registry.terraform.io/providers/DrFaust92/pingdom/latest/docs), generated from the [`docs/`](docs/) directory.

## Develop The Provider

### Dependencies for building from source

If you need to build from source, you should have a working Go environment setup. If not check out the Go [getting started](http://golang.org/doc/install) guide.

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) for dependency management. To fetch all dependencies run `go get` inside this repository.

### Build

```sh
make build
```

The binary will then be available at `_build/terraform-provider-pingdom_VERSION`.

### Install

```sh
make install
```

This will place the binary under `$HOME/.terraform.d/plugins/OS_ARCH/terraform-provider-pingdom_VERSION`. After installing you will need to run `terraform init` in any project using the plugin.

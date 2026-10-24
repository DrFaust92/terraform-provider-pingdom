---
layout: "pingdom"
page_title: "Pingdom: pingdom_integration"
sidebar_current: "docs-pingdom-data-integration"
description: |-
  Provides a data for a Pingdom integration
---

# pingdom\_integration

Provides a way to fetch data on a integration.

~> **Note:** This data source uses the SolarWinds API and requires the `solarwinds_user` and `solarwinds_passwd` provider arguments (or the `SOLARWINDS_USER` and `SOLARWINDS_PASSWD` environment variables). The Pingdom API token alone is not enough.

## Example Usage

```hcl
data "pingdom_integration" "example" {
  name = pingdom_integration.example.name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The integration name to look up.

## Attributes Reference

* `provider_name` - The integration provider name.
* `active` - Whether the integration is active.
* `url` - The integration url.

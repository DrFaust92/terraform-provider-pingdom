---
layout: "pingdom"
page_title: "Pingdom: pingdom_integrations"
sidebar_current: "docs-pingdom-data-integrations"
description: |-
  Provides a data for a Pingdom integrations
---

# pingdom\_integrations

Provides a way to fetch data on a integrations.

~> **Note:** This data source uses the SolarWinds API and requires the `solarwinds_user` and `solarwinds_passwd` provider arguments (or the `SOLARWINDS_USER` and `SOLARWINDS_PASSWD` environment variables). The Pingdom API token alone is not enough.

## Example Usage

```hcl
data "pingdom_integrations" "example" {}
```

## Argument Reference

The following arguments are supported:

No arguments can be passed.

## Attributes Reference

* `names` - The integrations names.
* `ids` - The integrations ids.

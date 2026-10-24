---
layout: "pingdom"
page_title: "Pingdom: pingdom_integration"
sidebar_current: "docs-pingdom-resource-integration"
description: |-
  Provides a Pingdom integration.
---

# pingdom\_integration

Provides a Pingdom integration resource, for example a webhook that is triggered by check alerts.

~> **Note:** This resource uses the SolarWinds API and requires the `solarwinds_user` and `solarwinds_passwd` provider arguments (or the `SOLARWINDS_USER` and `SOLARWINDS_PASSWD` environment variables). The Pingdom API token alone is not enough.

## Example Usage

```hcl
resource "pingdom_integration" "example" {
  provider_name = "webhook"
  active        = true
  name          = "slack"
  url           = "https://hooks.slack.com/services/XXX"
}
```

## Argument Reference

The following arguments are supported:

* `provider_name` - (Required) The integration provider. One of `webhook` or `librato` (`librato` is not currently operational).
* `active` - (Required) Whether the integration is active.
* `name` - (Required) The name of the integration.
* `url` - (Optional) The integration URL. Required when `provider_name` is `webhook`.

## Attributes Reference

* `id` - The integration ID. Use it in `pingdom_check.integrationids`.

## Import

Integrations can be imported using their ID, e.g.

```sh
terraform import pingdom_integration.example 1234
```

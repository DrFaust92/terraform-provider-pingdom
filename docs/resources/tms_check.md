---
layout: "pingdom"
page_title: "Pingdom: pingdom_tms_check"
sidebar_current: "docs-pingdom-resource-tms-check"
description: |-
  Provides a Pingdom transaction (TMS) check.
---

# pingdom\_tms\_check

Provides a Pingdom transaction monitoring (TMS) check resource.

## Example Usage

```hcl
resource "pingdom_tms_check" "example" {
  name     = "homepage"
  interval = 10

  steps {
    fn = "go_to"
    args = {
      url = "https://www.example.com"
    }
  }

  steps {
    fn = "url"
    args = {
      url = "https://www.example.com"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the check.
* `steps` - (Required) The steps of the transaction, at least one. See [Steps](#steps) below.
* `active` - (Optional) Whether the check is enabled.
* `contact_ids` - (Optional) Set of contact IDs to notify.
* `team_ids` - (Optional) Set of team IDs to notify.
* `integration_ids` - (Optional) Set of integration IDs to trigger.
* `custom_message` - (Optional) Message to include in alerts.
* `interval` - (Optional) How often the check runs, in minutes. One of `5`, `10`, `20`, `60`, `720` or `1440`. Default is `10`.
* `region` - (Optional) The region the check runs from. Default is `us-east`.
* `send_notification_when_down` - (Optional) Number of consecutive failures before the check is considered down. Default is `1`.
* `security_level` - (Optional) How important alerts are when the check fails. One of `low` or `high`. Default is `high`.
* `tags` - (Optional) Tags for the check. Tag names may contain `A-Z`, `a-z`, `0-9`, `_` and `-`, up to 64 characters.
* `metadata` - (Optional) Metadata for recorded transactions. See [Metadata](#metadata) below.

### Steps

* `fn` - (Required) The function the step performs.
* `args` - (Optional) Map of arguments for the function.

### Metadata

* `authentication` - (Optional) Map of authentication settings.
* `disable_websecurity` - (Optional) Whether to disable web security.
* `height` - (Optional) Browser viewport height.
* `width` - (Optional) Browser viewport width.

## Attributes Reference

* `id` - The TMS check ID.

## Import

TMS checks can be imported using their ID, e.g.

```sh
terraform import pingdom_tms_check.example 1234
```

---
layout: "pingdom"
page_title: "Pingdom: pingdom_maintenance"
sidebar_current: "docs-pingdom-resource-maintenance"
description: |-
  Provides a Pingdom maintenance window.
---

# pingdom\_maintenance

Provides a Pingdom maintenance window resource. A maintenance window can be one-off or recurring, and applies to one or more uptime and transaction (TMS) checks.

## Example Usage

```hcl
resource "pingdom_check" "example" {
  name = "example-check"
  host = "www.example.com"
  type = "http"
}

resource "pingdom_maintenance" "example" {
  description    = "weekly maintenance"
  from           = "2066-01-02T22:00:00+08:00"
  to             = "2066-01-02T23:00:00+08:00"
  effectiveto    = "2067-01-02T23:00:00+08:00"
  recurrencetype = "week"
  repeatevery    = 4
  uptimeids      = [pingdom_check.example.id]
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Required) The description of the maintenance window.
* `from` - (Required) The start of the initial maintenance window, in RFC3339 format, e.g. `2066-01-02T22:00:00+08:00`.
* `to` - (Required) The end of the initial maintenance window, in RFC3339 format.
* `effectiveto` - (Optional) The end of the recurrence, in RFC3339 format. Defaults to `to`.
* `recurrencetype` - (Optional) The type of recurrence. One of `none`, `day`, `week`, `month` or `nth-day-of-week-in-month`. Default is `none`.
* `repeatevery` - (Optional) Repeat every n-th day, week or month. Default is `0`.
* `uptimeids` - (Optional) Set of uptime check IDs to assign to the maintenance window.
* `tmsids` - (Optional) Set of transaction (TMS) check IDs to assign to the maintenance window.

## Attributes Reference

* `id` - The maintenance window ID.

## Import

Maintenance windows can be imported using their ID, e.g.

```sh
terraform import pingdom_maintenance.example 1234
```

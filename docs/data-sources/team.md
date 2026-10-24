---
layout: "pingdom"
page_title: "Pingdom: pingdom_team"
sidebar_current: "docs-pingdom-data-team"
description: |-
  Provides a data for a Pingdom team
---

# pingdom\_team

Provides a way to fetch data on a team.

## Example Usage

```hcl
data "pingdom_team" "example" {
  name = "gob"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The team name to look up.

## Attributes Reference

* `id` - The team ID.
* `member_ids` - A list of memebers ids.

---
layout: "pingdom"
page_title: "Pingdom: pingdom_teams"
sidebar_current: "docs-pingdom-data-teams"
description: |-
  Provides a data for a Pingdom teams
---

# pingdom\_team

Provides a way to fetch data on teams.

## Example Usage

```hcl
data "pingdom_teams" "example" {}
```

## Argument Reference

No arguments can be passed.

## Attributes Reference

* `names` - A list of team names.
* `ids` - A list of team ids.

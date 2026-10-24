---
layout: "pingdom"
page_title: "Pingdom: pingdom_team"
sidebar_current: "docs-pingdom-resource-team"
description: |-
  Provides a Pingdom alerting team.
---

# pingdom\_team

Provides a Pingdom alerting team resource.

## Example Usage

```hcl
resource "pingdom_team" "example" {
  name = "The Test team"
  member_ids = [
    pingdom_contact.example.id,
  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the team.
* `member_ids` - (Optional) Set of contact IDs that are members of the team.

## Attributes Reference

* `id` - The team ID.

## Import

Teams can be imported using their ID, e.g.

```sh
terraform import pingdom_team.example 1234
```

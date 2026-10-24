---
layout: "pingdom"
page_title: "Pingdom: pingdom_contacts"
sidebar_current: "docs-pingdom-data-contacts"
description: |-
  Provides a data for a Pingdom contacts
---

# pingdom\_contacts

Provides a way to fetch data on a contacts.

## Example Usage

```hcl
data "pingdom_contacts" "example" {}
```

## Argument Reference

The following arguments are supported:

No arguments can be passed.

## Attributes Reference

* `names` - The contacts names.
* `ids` - The contacts ids.
* `types` - Contact types.

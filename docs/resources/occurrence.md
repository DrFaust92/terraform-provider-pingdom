---
layout: "pingdom"
page_title: "Pingdom: pingdom_occurrence"
sidebar_current: "docs-pingdom-resource-occurrence"
description: |-
  Provides a group of Pingdom maintenance occurrences.
---

# pingdom\_occurrence

Provides a resource that manages a group of maintenance occurrences, selected by the query `(maintenance_id, effective_from, effective_to)`.

`effective_from` and `effective_to` are only query conditions used to find occurrences. They are different from `from` and `to`, which set the start and end of the matched occurrences within a maintenance cycle. Updating `from` or `to` updates every occurrence matched by the query, so all occurrences matched by one resource share the same `from` and `to`.

## Example Usage

```hcl
resource "pingdom_occurrence" "example" {
  maintenance_id = pingdom_maintenance.example.id
  effective_from = pingdom_maintenance.example.from
  effective_to   = pingdom_maintenance.example.effectiveto
  from           = pingdom_maintenance.example.from
  to             = "2066-01-02T22:30:00+08:00"
}
```

## Argument Reference

The following arguments are supported:

* `maintenance_id` - (Required) The ID of the maintenance window the occurrences belong to. Changing this forces a new resource.
* `effective_from` - (Required) The start of the occurrence query, in RFC3339 format. Use the current time or later to only match future occurrences.
* `effective_to` - (Required) The end of the occurrence query, in RFC3339 format.
* `from` - (Optional) The start time to set on the matched occurrences, in RFC3339 format.
* `to` - (Optional) The end time to set on the matched occurrences, in RFC3339 format.

## Attributes Reference

* `id` - The resource ID.
* `size` - The number of occurrences matched by the query.

## Import

Occurrences cannot be imported, because the resource is a query rather than an object that exists in Pingdom.

---
layout: "pingdom"
page_title: "Pingdom: pingdom_contact"
sidebar_current: "docs-pingdom-data-contact"
description: |-
  Provides a data for a Pingdom contact
---

# pingdom\_contact

Provides a way to fetch data on a contact.

## Example Usage

```hcl
data "pingdom_contact" "example" {
  name = pingdom_contact.example.name
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The contact name to look up.

## Attributes Reference

* `id` - The contact ID.
* `paused` - Whether contact is paused.
* `teams` - Teams connected to contact.
* `sms_notification` - SMS notifcation details.
* `email_notification` - Email notifcation details.

### SMS Notification

* `country_code` - Country code
* `number` - Phone number
* `provider` - Provider
* `severity` - Contact target's severity level

### Email Notification

* `address` - Email address
* `severity` - Contact target's severity level

### Teams

* `id` - Team id.
* `name` - Team name.

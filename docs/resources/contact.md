---
layout: "pingdom"
page_title: "Pingdom: pingdom_contact"
sidebar_current: "docs-pingdom-resource-contact"
description: |-
  Provides a Pingdom contact.
---

# pingdom\_contact

Provides a Pingdom alerting contact resource.

~> **Note:** Every contact must have at least one `HIGH` and one `LOW` severity notification.

## Example Usage

```hcl
resource "pingdom_contact" "example" {
  name = "johndoe"

  sms_notification {
    number   = "5555555555"
    severity = "HIGH"
  }

  sms_notification {
    number       = "3333333333"
    country_code = "91"
    severity     = "LOW"
    provider     = "esendex"
  }

  email_notification {
    address  = "test@example.com"
    severity = "LOW"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the contact.
* `paused` - (Optional) Whether alerts for this contact are disabled. Default is `false`.
* `sms_notification` - (Optional) SMS notification targets. See [SMS Notification](#sms-notification) below.
* `email_notification` - (Optional) Email notification targets. See [Email Notification](#email-notification) below.

### SMS Notification

* `number` - (Required) The phone number.
* `severity` - (Required) The severity of this notification. One of `HIGH` or `LOW`.
* `country_code` - (Optional) The country code. Default is `1`.
* `provider` - (Optional) The SMS provider. One of `nexmo`, `bulksms`, `esendex` or `cellsynt` (`bulksms` is not currently operational). Default is `nexmo`.

### Email Notification

* `address` - (Required) The email address to notify.
* `severity` - (Required) The severity of this notification. One of `HIGH` or `LOW`.

## Attributes Reference

* `id` - The contact ID.

## Import

Contacts can be imported using their ID, e.g.

```sh
terraform import pingdom_contact.example 1234
```

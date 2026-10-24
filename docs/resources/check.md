---
layout: "pingdom"
page_title: "Pingdom: pingdom_check"
sidebar_current: "docs-pingdom-resource-check"
description: |-
  Provides a Pingdom Check
---

# pingdom\_check

Provides a Pingdom check resource.

This allows you to manage your pingdom checks.

## Example Usage

```hcl
resource "pingdom_check" "example" {
  type                     = "http"
  name                     = "my http check"
  host                     = "example.com"
  resolution               = 5
  sendnotificationwhendown = 2
  integrationids           = [pingdom_integration.example.id]
  teamids                  = [pingdom_team.example.id]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Check name.
* `host` - (Required) Target host.
* `type` - (Required) Type of check. Valid values are `http`, `tcp`, `ping` and `dns`.
* `custom_message` - (Optional) Custom message that will be added to email and webhook alerts.
* `encryption` - (Optional) Connection encryption.
* `expectedip` - (Optional) Expected IP. Valid only for `dns` check type.
* `integrationids` - (Optional) Integration identifiers.
* `nameserver` - (Optional) DNS server to use. Valid only for `dns` check type.
* `notifyagainevery` - (Optional) Notify again every n result. 0 means that no extra notifications will be sent.
* `notifywhenbackup` - (Optional) Notify when back up again.
* `password` - (Optional) Password to auth with. Stored in state, since the API returns it.
* `password_wo` - (Optional) Write-only alternative to `password`, never stored in Terraform state or plan output. Requires Terraform 1.11 or later. Conflicts with `password`. See [write-only arguments](https://developer.hashicorp.com/terraform/language/manage-sensitive-data/write-only).
* `password_wo_version` - (Optional) Required with `password_wo`. Terraform cannot detect changes to write-only values, change this value (e.g. `"2"`) to send an updated `password_wo`.
* `paused` - (Optional) Whether the check is paused.
* `port` - (Optional) Target port.
* `postdata` - (Optional) Data that should be posted to the web page, for example submission data for a sign-up or login form. The data needs to be formatted in the same way as a web browser would send it to the web server.
* `probefilters` - (Optional) Filters used for probe selections. Overwrites previous filters for check. To remove all filters from a check, use an empty value. Comma separated key:value pairs. Currently only region is supported. Possible values are `EU`, `NA`, `APAC`, and `LATAM`. For example, "region: NA".
* `requestheaders` - (Optional) Custom HTTP header. The entry value should contain a one-element string array. The element should contain headerName and headerValue colon-separated. To add more than one header send other parameters named requestheaders{number}.
* `resolution` - (Optional) How often should the check be tested? (minutes). Valid values are `1`, `5`, `15`, `30`, `60`. Default value is `5`.
* `responsetime_threshold` - (Optional) Triggers a down alert if the response time exceeds threshold specified in ms (Not available for Starter and Free plans.)
* `sendnotificationwhendown` - (Optional) Send notification when down X times. Default value is `2`.
* `shouldcontain` - (Optional) Target site should contain this string. Note! This parameter cannot be used together with the parameter `shouldnotcontain`, use only one of them in your request. Valid for `http` check.
* `shouldnotcontain` - (Optional) Target site should NOT contain this string. Note! This parameter cannot be used together with the parameter `shouldcontain`, use only one of them in your request. Valid for `http` check.
* `ssl_down_days_before` - (Optional) Treat the target site as down if a certificate expires within the given number of days. This parameter will be ignored if `verify_certificate` is set to `false`. It will appear provided `verify_certificate` is `true` and `ssl_down_days_before` value is greater than or equals `1`. Default value is `0`.
* `tags` - (Optional) List of tags for check. The maximum length of a tag is 64 characters.
* `stringtoexpect` - (Optional) String to expect in response. Valid for `tcp` check.
* `stringtosend` - (Optional) String to send. Valid for `tcp` check.
* `teamids` - (Optional) Teams to alert.
* `url` - (Optional) Path to target on server.
* `userids` - (Optional) Users to alert.
* `username` - (Optional) Username to auth with.
* `verify_certificate` - (Optional) Treat target site as down if an invalid/unverifiable certificate is found. Default value is `true`.
* `ipv6` - (Optional) Wheter to enable IPV6 on check. Default value is `false`.

## Attributes Reference

* `id` - The check ID.

## Import

Checks can be imported using their id, e.g.

```sh
terraform import pingdom_check.example check_id
```

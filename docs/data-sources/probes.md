---
layout: "pingdom"
page_title: "Pingdom: pingdom_probes"
sidebar_current: "docs-pingdom-data-probes"
description: |-
  Provides the list of Pingdom probe servers.
---

# pingdom\_probes

Provides the list of Pingdom probe servers, for example to allow Pingdom's IP addresses through a firewall.

## Example Usage

```hcl
data "pingdom_probes" "na" {
  only_active = true
  region      = "NA"
}

resource "aws_security_group_rule" "pingdom" {
  type              = "ingress"
  from_port         = 443
  to_port           = 443
  protocol          = "tcp"
  cidr_blocks       = [for ip in data.pingdom_probes.na.ipv4_addresses : "${ip}/32"]
  security_group_id = aws_security_group.example.id
}
```

## Argument Reference

The following arguments are supported:

* `only_active` - (Optional) Only return active probes. Default is `false`.
* `region` - (Optional) Only return probes in this region. One of `NA`, `EU`, `APAC` or `LATAM`.

## Attributes Reference

* `ipv4_addresses` - List of the probes' IPv4 addresses.
* `ipv6_addresses` - List of the probes' IPv6 addresses. Probes without an IPv6 address are skipped.
* `probes` - List of probes. Each probe has the following attributes:
  * `id` - The probe ID.
  * `name` - The probe name.
  * `active` - Whether the probe is active.
  * `hostname` - The probe hostname.
  * `ip` - The probe IPv4 address.
  * `ipv6` - The probe IPv6 address.
  * `city` - The city the probe is in.
  * `country` - The country the probe is in.
  * `country_iso` - The ISO code of the country the probe is in.
  * `region` - The region the probe is in.

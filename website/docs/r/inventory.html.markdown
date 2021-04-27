---
layout: "dcnm"
page_title: "DCNM: dcnm_inventory"
sidebar_current: "docs-dcnm-resource-inventory"
description: |-
  Manages DCNM inventory modules
---

# dcnm_inventory #
Manages DCNM inventory modules

## Example Usage ##

```hcl

resource "dcnm_inventory" "first" {
  fabric_name   = "fab2"
  username      = "username for DCNM switches"
  password      = "password for DCNM switches"
  max_hops      = 0
  preserve_config = "false"
  auth_protocol = 0
  config_timeout = 10
  switch_config {
    ip   = "switch IP"
    role = "leaf"
  }
  switch_config {
    ip   = "switch IP"
    role = "spine"
  }
  switch_config {
    ip   = "switch IP"
    role = "leaf"
  }
}

```


## Argument Reference ##

* `fabric_name` - (Required) fabric name under which inventory should be created.
* `username` - (Required) username for the the switch.
* `password` - (Required) password for the the switch.
* `max_hops` - (Optional) maximum number hops for switch. Ranging from 0 to 10, default value is 0.
* `auth_protocol` - (Optional) authentication protocol for switch. Mapping is as `0 : "MD5", 1: "SHA", 2 : "MD5_DES", 3 : "MD5_AES", 4 : "SHA_DES", 5 : "SHA_AES"`
* `preserve_config` - (Optional) flag to preserve the configuration of switch. Default value is "false".
* `platform` - (Optional) platform name for the switch.
* `second_timeout` - (Optional) second timeout value for switch.
* `config_timeout` - (Optional) configuration timeout value in minutes. Default value is "5".

* `switch_config` - (Required) switch configuration block for inventory resource. It consists of the information regarding switches.
* `switch_config.ip` - (Required) ip Address of switch.
* `switch_config.role` - (Optional) role of the switch. Allowed values are "leaf", "spine", "border", "border_spine", "border_gateway", "border_gateway_spine", "super_spine", "border_super_spine", "border_gateway_super_spine".

* `deploy` - (Optional) deploy flag for the switch. Default value is "true".

## Attribute Reference

* `id` - Dn for the switch inventory.

* `switch_config` - Switch configuration block for inventory.
* `switch_config.switch_name` - Name of the switch.
* `switch_config.switch_db_id` - DB ID for the switch.
* `switch_config.serial_number` - Serial number of the switch.
* `switch_config.model` - Model name of the switch.
* `switch_config.mode` - Mode of the switch.

## Importing ##

An existing switch inventory can be [imported][docs-import] into this resource via its fabric and name, using the following command:
[docs-import]: https://www.terraform.io/docs/import/index.html


```
terraform import dcnm_inventory.example <fabric_name>:<switch_name>
```
---
layout: "sdwan"
page_title: "SD-WAN: sdwan_device_uuids"
sidebar_current: "docs-sdwan-data-source-device_uuids"
description: |-
  Data source for SD-WAN Device UUID List
---
# sdwan_device_uuids Data Source #
Data source for SD-WAN Device UUID List

## Example Usage ##

```hcl

data "sdwan_device_uuids" "get_devices" {
}

```

## Common Attribute Reference ##
* `device` - List of Devices on vManage server (see [below for nested schema](#nestedblock--device))

<a id="nestedblock--device"></a>

## Nested Schema for device
* `uuid` - UUID of the Device
* `device_model` - Model used by the Device
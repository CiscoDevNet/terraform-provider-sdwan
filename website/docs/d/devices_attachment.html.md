---
layout: "sdwan"
page_title: "SD-WAN: sdwan_devices_attachment"
sidebar_current: "docs-sdwan-data-source-devices_attachment"
description: |-
  Data Source for SD-WAN Devices Attachment 
---
# sdwan_devices_attachment #
Data Source for SD-WAN Devices Attachment

## Example Usage ##

```hcl

data "sdwan_devices_attachment" "name" {
  device_template_id = "a137ff27-6370-4214-b060-17d29452a297"
}

```
## Argument Reference ##
* `device_template_id` - (Required) Unique Device Template ID

## Common Attribute Reference ##
* `template_type` - Configuration type for  Device Template
* `devices_list` - List of Devices attached to the Device Template (see [below for nested schema](#nestedblock--devices_list))

<a id="nestedblock--devices_list"></a>
## Nested Schema for devices_list
* `chassis_number` - Chassis number of the Device attached to the Device Template

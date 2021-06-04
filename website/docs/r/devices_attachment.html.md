---
layout: "sdwan"
page_title: "SD-WAN: sdwan_devices_attachment"
sidebar_current: "docs-sdwan-resource-devices_attachment"
description: |-
  Manages SD-WAN Devices Attachment
---

# sdwan_devices_attachment #
Manages SD-WAN Device Attachment

## Example Usage ##

```hcl
# example for Devices to Device Template attachment

resource "sdwan_devices_attachment" "name" {
  timeout = 600
  device_template_id = "a137ff27-6370-4214-b060-17d29452a297"
  file = "featureInput.csv"
  devices_list {
    chassis_number = "CSR-F3EF66F0-A1E8-4C11-9C12-3C95ED91A2CC"
  }
  devices_list {
    chassis_number = "CSR-AE05722E-FD8F-41A1-A799-A9398646DE6F"
  }
  devices_list {
    chassis_number = "CSR-DC54554A-5BE7-49BD-936B-702ECD2FBE4B"
  }
}

```
## Common Argument Reference ##
* `device_template_id` - (Required) Unique Device Template ID
* `file` - (Required) Path of the file with Device variables to be passed
* `devices_list` - (Required) List of Devices to be attached to the Device Template (see [below for nested schema](#nestedblock--devices_list))
* `timeout` - (Optional) Maximum Time period (in seconds) for which the operation should take place, Default: 600

<a id="nestedblock--devices_list"></a>
## Nested Schema for devices_list ##
* `chassis_number` - (Required) Chassis number of the Device attached to the Device Template
* `attach` - (Optional) Boolean flag to indicate whether the Device is attached or not, Default: true

## Attribute Reference for devices_list ##
* `status` - Status of Device Attachment
* `activity` - List of Activities at the time of Device attachment

---
layout: "sdwan"
page_title: "SD-WAN: sdwan_devices_attachment"
sidebar_current: "docs-sdwan-resource-devices_attachment"
description: |-
  Manages SD-WAN Devices Attachment
---

# sdwan_devices_attachment Resource #
Manages SD-WAN Device Attachment

## Example Usage ##

```hcl
# example for Devices to Device Template attachment
resource "sdwan_devices_attachment" "example_devices_attachment" {
  device_template_id = sdwan_device_template.cisco_device_template.template_id
  file = "Template.csv"

  devices_list{
    chassis_number = "CSR-53D3BC42-6F6C-4529-8D03-4B2903D1E688"
    }
  devices_list{
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

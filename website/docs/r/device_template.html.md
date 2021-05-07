---
layout: "sdwan"
page_title: "SD-WAN: sdwan_device_template"
sidebar_current: "docs-sdwan-resource-device_template"
description: |-
  Manages SD-WAN Device Templates
---

# sdwan_device_template #
Manages SD-WAN Device Templates

## Example Usage ##

```hcl
# example for Device Template created from CLI
resource "sdwan_device_template" "example_feature_template"{
  template_name = "example1"
  template_description = "For testing purpose"
  device_type = "vedge-cloud"
  device_role = "sdwan-edge"
  config_type = "file"
  factory_default = false
  template_configuration = "`"
}

# example for Device Template created from Feature Templates
resource "sdwan_device_template" "example_cli" {
  template_name = "example2"
  template_description = "For testing purpose"
  device_type = "vedge-cloud"
  device_role = "sdwan-edge"
  config_type = "template"
  factory_default = false
  connection_preference = false
  connection_preference_required = false

  general_templates{
    template_id = "c8746871-cb8e-4ab7-a5f8-948c624f19ac"
    template_type = "system-vedge"
  }
  general_templates{
    template_id = "5083367c-0db3-4dd7-bc05-ad1afbbceaf4"
    template_type = "vpn-vedge"
  }

  general_templates{
    template_id = "79437c57-9b0f-4364-b060-55e827ac7e45"
    template_type = "vpn-vedge"
  }

  policy_id = "8715a21d-9367-47ea-9bc6-e25163ed9513"
  
  feature_template_uid_range = ["string", 123]

}

```
## Common Argument Reference ##
* `template_name` - (Required) Unique Device Template Name (length (1 - 128), must be alphanumeric and should not include [, ^, <, >, !, &, \, ", ], *, $)
* `template_description` - (Required) Long Description of Device Template
* `device_type` - (Required) Type of device for which Device Template should be created, allowed values: "vedge-100", "vedge-100-B", "vedge-100-M", "vedge-100-WM", "vedge-1000", "vedge-2000", "vedge-5000", "vedge-cloud", "vedge-ISR1100-4G", "vedge-ISR1100-4GLTE", "vedge-ISR1100-6G", "vedge-ISR1100X-4G", "vedge-ISR1100X-6G", "vmanage", "vsmart"
* `device_role` - (Required) Device role of the device, allowed values: "sdwan-edge", "service-node"
* `factory_default` - (Required) Flag indicating whether Device Template is factory default or not, allowed values: `true` or `false`
* `config_type` - (Required) Configuration type for  Device Template, allowed values: "template" (for Device Template created from Feature Template), "file" (for Device Template created from CLI) 

## Argument Reference for Device Template created from Feature Templates ##
* `policy_id` - (Required) policyId for  Device Template
* `feature_template_uid_range` - (Required) featureTemplateUidRange for  Device Template
* `general_templates` - (Required) List of Feature Templates and thier Sub Templates thourgh which Device Template is created (should not be empty ,see [below for nested schema])
* `connection_preference_required` - (Optional) connectionPreferenceRequired flag for Device Template, allowed values: `true` or `false`
* `connection_preference` - (Optional) connectionPreference flag for Device Template, allowed values: `true` or `false`

## Nested Schema for general_templates ##
* `template_id` - (Required) Template Id of Feature Template
* `template_type` - (Required) Template Type of Feature Template, allowed values: "system-vedge", "ntp", "vpn-vedge", "vpn-vedge-interface"
* `sub_templates` - (Optional) List of Sub Templates associated with feature Template (see [below for nested schema])

## Nested Schema for sub_templates ##
* `template_id` - (Required) Template Id of Feature Template
* `template_type` - (Required) Template Type of Feature Template, allowed values: "system-vedge", "ntp", "vpn-vedge", "vpn-vedge-interface"

## Argument Reference for Device Template created from CLI ##
* `template_configuration` - (Required) Template Configuration for  Device Template

## Common Attributes ##
* `template_id` - Device Template id for  Device Template
* `template_class` - Template Class type for  Device Template

## Attributes for Device Template created from CLI ##
* `last_updated_by` - User who updated  Device Template last
* `last_updated_on` - Time when Device Template was last updated
* `rid` - Request ID for Device Template
* `created_on` - Time when  Device Template was created
* `created_by` - User who created Device Template
* `feature` - Feature for Device Template
* `template_attached` - Number of templates attached to Device Template





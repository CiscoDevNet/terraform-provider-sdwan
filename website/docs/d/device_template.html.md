---
layout: "sdwan"
page_title: "SD-WAN: sdwan_device_template"
sidebar_current: "docs-sdwan-data-source-device_template"
description: |-
  Data source for SD-WAN Device Templates 
---
# sdwan_device_template #
Data source for SD-WAN Device Templates

## Example Usage ##

```hcl

data "sdwan_device_template" "name" {
    template_name = "example_template"
}

```
## Argument Reference ##
* `template_name` - (Required) Unique Device Template Name

## Common Attribute Reference ##
* `template_description` - Long Description of Device Template
* `device_type` - Type of device supported by  Device Template
* `device_role` - Device role for the Device Template
* `config_type` - Configuration type for  Device Template
* `factory_default` - Flag indicating whether Device Template is factory default or not
* `template_class` - Template Class type for  Device Template
* `template_id` - Device Template id for  Device Template

## Attribute Reference for Device Template created from Feature Templates ##
* `policy_id` - policyId for  Device Template
* `feature_template_uid_range` - feature_template_uid_range for  Device Template
* `connection_preference_required` - connectionPreferenceRequired flag for Device Template
* `connection_preference` - connectionPreference flag for Device Template
* `general_templates` - List of Feature Templates and thier Sub Templates thourgh which Device Template is created (see [below for nested schema](#nestedblock--general_templates))


<a id="nestedblock--general_templates"></a>

## Nested Schema for general_templates
* `template_id` - Template Id of Feature Template
* `template_type` - Template Type of Feature Template
* `sub_templates` - List of Sub Templates associated with feature Template (see [below for nested schema](#nestedblock--sub_templates))

<a id="nestedblock--sub_templates"></a>

## Nested Schema for sub_templates
* `template_id` - Template Id of Feature Template
* `template_type` - Template Type of Feature Template

## Attribute Reference for Device Template created from CLI ##
* `last_updated_by` - User who updated  Device Template last
* `template_configuration` - Template Configuration for  Device Template
* `created_on` - Time when  Device Template was created
* `rid` - Request ID for Device Template
* `feature` - Feature for Device Template
* `created_by` - User who created Device Template
* `last_updated_on` - Time when Device Template was last updated
* `template_attached` - Number of templates attached to Device Template


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
    template_type = "aaa"
  }

  general_templates{
    template_id = "5083367c-0db3-4dd7-bc05-ad1afbbceaf4"
    template_type = "bfd-vedge"
  }

  general_templates{
    template_id = "79437c57-9b0f-4364-b060-55e827ac7e45"
    template_type = "omp-vedge"
  }

  general_templates{
    template_id = "486d419f-4e6c-44a5-a6fb-7b5ccf94ff90"
    template_type = "security-vedge"
  }
  
  general_templates{
    template_id = "45ea940a-45d2-4fd9-8da2-570a1a6d6874"
    template_type = "system-vedge"
    sub_templates{
      template_id = "171e9bd4-7a7b-460d-b692-83f0d5ce0124"
      template_type = "logging"
    }
    sub_templates{
      template_id = "edf3d309-91d4-45be-98d9-cfd57a05a479"
      template_type = "ntp"
    }
  }

  general_templates{
    template_id = "79437c57-9b0f-4364-b060-55e827ac7e45"
    template_type = "vpn-vedge"
    sub_templates{
      template_id = " e1b5b6e9-3b54-4279-a532-a2aaaef3e6a1"
      template_type = "vpn-vedge-interface"
    }
  }

  policy_id = "8715a21d-9367-47ea-9bc6-e25163ed9513"
  
  feature_template_uid_range = ["string", 123]

}

```
## Common Argument Reference ##
* `template_name` - (Required) Unique Device Template Name, Must not include <, >, !, &, ", or white space; maximum 128 characters
* `template_description` - (Required) Long Description of Device Template
* `device_type` - (Required) Type of device for which Device Template should be created, allowed values: 
"vedge-100", "vedge-100-B", "vedge-100-M", "vedge-100-WM", "vedge-1000", "vedge-2000", "vedge-5000",   "vedge-cloud", "vedge-ISR1100-4G", "vedge-ISR1100-4GLTE", "vedge-ISR1100-6G", "vedge-ISR1100X-4G", "vedge-ISR1100X-6G", "vedge-ASR-1001-HX", "vedge-ASR-1001-X", "vedge-ASR-1002-HX", "vedge-ASR-1002-X", "vedge-C1101-4P", "vedge-C1101-4PLTEP", "vedge-C1101-4PLTEPW", "vedge-C1109-2PLTEGB", "vedge-C1109-2PLTEUS", "vedge-C1109-2PLTEVZ", "vedge-C1109-4PLTE2PW", "vedge-C1109-4PLTE2P", "vedge-C1111-4P", "vedge-C1111-4PW", "vedge-C1111-4PLTELA", "vedge-C1111-8PLTEEA", "vedge-C1111-8PW", "vedge-C1111-8P", "vedge-C1111-8PLTELAW", "vedge-C1111-8PLTEEAW", "vedge-C1111-8PLTELA", "vedge-C1111X-8P", "vedge-C1113-8PLTEEAW", "vedge-C1113-8PLTELA",  "vedge-C1113-8PMLTEEA", "vedge-C1116-4P", "vedge-C1116-4PLTEEA", "vedge-C1117-4P", "vedge-C1117-4PLTEEA", "vedge-C1117-4PLTELA", "vedge-C1117-4PM", "vedge-C1117-4PMLTEEA", "vedge-C1121-4P", "vedge-C1121-4PLTEP", "vedge-C1121-8P", "vedge-C1121-8PLTEP", "vedge-C1121-8PLTEPW", "vedge-C1121X-8P", "vedge-C1121X-8PLTEP", "vedge-C1121X-8PLTEPW", "vedge-C1126-8PLTEP", "vedge-C1126X-8PLTEP", "vedge-C1127-8PLTEP", "vedge-C1127-8PMLTEP", "vedge-C1127X-8PLTEP", "vedge-C1127X-8PMLTEP", "vedge-C1128-8PLTEP", "vedge-C1161-8P", "vedge-C1161X-8P", "vedge-C1161X-8PLTEP", "vedge-ISR-4221", "vedge-ISR-4221X", "vedge-ISR-4321", "vedge-ISR-4331", "vedge-ISR-4351", "vedge-ISR-4431", "vedge-ISR-4451-X", "vedge-ISR-4461", "vedge-ISRv", "vedge-CSR-1000v"
* `device_role` - (Required) Device role of the device, allowed values: "sdwan-edge", "service-node"
* `factory_default` - (Required) Flag indicating whether Device Template is factory default or not, allowed values: `true` or `false`
* `config_type` - (Required) Configuration type for  Device Template, allowed values: "template" (for Device Template created from Feature Template), "file" (for Device Template created from CLI) 

## Argument Reference for Device Template created from Feature Templates ##
* `policy_id` - (Required) policyId for  Device Template
* `feature_template_uid_range` - (Required) featureTemplateUidRange for  Device Template
* `general_templates` - (Required) List of Feature Templates and thier Sub Templates thourgh which Device Template is created (should not be empty , (see [below for nested schema](#nestedblock--general_templates)))
* `connection_preference_required` - (Optional) connectionPreferenceRequired flag for Device Template, allowed values: `true` or `false`
* `connection_preference` - (Optional) connectionPreference flag for Device Template, allowed values: `true` or `false`


<a id="nestedblock--general_templates"></a>

## Nested Schema for general_templates ##
* `template_id` - (Required) Template Id of Feature Template
* `template_type` - (Required) Template Type of Feature Template, allowed values: "aaa", "bfd-vedge", "omp-vedge", "security-vedge", "system-vedge",  "vpn-vedge", "cedge_aaa", "cisco_bfd", "cisco_omp", "cisco_security", "cisco_system",  "cisco_vpn", "cedge_global" <br>
<strong>NOTE</strong><br>
`aaa`, `bfd-vedge`, `omp-vedge`, `security-vedge`, `system-vedge`, `vpn-vedge` are required for vEdge Devices.<br>
`cedge_aaa`, `cisco_bfd`, `cisco_omp`, `cisco_security`, `cisco_system`, `cisco_vpn`, `cedge_global` are required for cEdge Devices
* `sub_templates` - (Optional) List of Sub Templates associated with feature Template (see [below for nested schema](#nestedblock--sub_templates))


<a id="nestedblock--sub_templates"></a>

## Nested Schema for sub_templates ##
* `sub_template_id` - (Required) Template Id of Feature Template
* `sub_template_type` - (Required) Template Type of Feature Template, allowed values: "logging" (this can be allowed only with `system-vedge` type `general_templates`), "ntp" (this can be allowed only with `system-vedge` type `general_templates`), "vpn-vedge-interface" (this can be allowed only with`vpn-vedge` type `general_templates`), "cisco_logging" (this can be allowed only with `cisco_system` type `general_templates`), "cisco_ntp"(this can be allowed only with `"cisco_system"` type `general_templates`), "cisco_vpn_interface" (this can be allowed only with `cisco_vpn` type `general_templates`)

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

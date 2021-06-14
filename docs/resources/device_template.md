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
resource "sdwan_device_template" "cisco_device_template" {
  template_name = "Demo-Device-Template0"
  template_description = "For demo purpose"
  device_type = "vedge-CSR-1000v"
  device_role = "sdwan-edge"
  config_type = "template"
  factory_default = false
  connection_preference = false
  connection_preference_required = false

  # Required factory default feature templates
  general_templates{
    template_id = "f4422bdc-0933-4138-b4c4-8caa7ea169f6"
    template_type = "cedge_aaa"
  }

  general_templates{
    template_id = "f9e7dee5-0db5-492d-80d6-224f893caefc"
    template_type = "cisco_bfd"
  }

  general_templates{
    template_id = "39740dbf-b641-487f-b1d9-07f6cd9ce3e5"
    template_type = "cisco_omp"
  }

  general_templates{
    template_id = "b495a21b-eca8-4d5f-97e7-19508edeba8f"
    template_type = "cisco_security"
  }

  general_templates{
    template_id = "b3625ec5-512e-4a64-b780-b16ebceb60b5"
    template_type = "cedge_global"
  }


  general_templates{
    template_id = sdwan_system_feature_template.example_cedge_system.template_id
    template_type = "cisco_system"
    sub_templates{
      sub_template_id = "c9e48a21-87c3-4211-8c67-d22f64e851c2"
      sub_template_type = "cisco_logging"
    }
    # Optional Cisco NTP feature template
    sub_templates{
      sub_template_id = sdwan_ntp_feature_template.example_cedge_ntp.template_id
      sub_template_type = "cisco_ntp"
    }
  }

  general_templates{
    template_id = sdwan_vpn_feature_template.example_cedge_vpn0.template_id
    template_type = "cisco_vpn"
    #Optional Cisco VPN Interface feature template
    sub_templates{
      sub_template_id = sdwan_vpn_interface_feature_template.example_cedge_vpn_interface.template_id
      sub_template_type = "cisco_vpn_interface"
    }
  }

  general_templates{
    template_id = sdwan_vpn_feature_template.example_cedge_vpn512.template_id
    template_type = "cisco_vpn"
    #Optional Cisco VPN Interface feature template
    sub_templates{
      sub_template_id = sdwan_vpn_interface_feature_template.example_cedge_vpn_interface.template_id
      sub_template_type = "cisco_vpn_interface"
    }
  }

  policy_id = ""

  feature_template_uid_range = []
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

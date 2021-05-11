---
layout: "sdwan"
page_title: "SD-WAN: sdwan_ntp_feature_template"
sidebar_current: "docs-sdwan-data-source-ntp_feature_template"
description: |-
  Data source for SD-WAN Ntp Feature Template
---

# sdwan_ntp_feature_template #
Data source for SD-WAN Ntp Feature Templates

## Example Usage ##

```hcl

data "sdwan_ntp_feature_template" "name" {
    template_name = "example_template"
}

```


## Argument Reference ##

* `template_name` - (Required) Unique  Ntp Type Feature Template Name

## Attribute Reference ##

* `template_description` - Long Description of  Ntp type Feature Template
* `device_type` - Type of device which supports  Feature Template
* `template_min_version` - Minimum Version for the  Feature template
* `template_definition` - Configuration for  Ntp type Feature Template, (see [below for nested schema])
* `factory_default` - Boolean flag indicating whether  Ntp type Feature Template is factory default or not
* `attached_masters_count` - Number of Device Templates attached to  Ntp type Feature Template
* `devices_attached` - Number of Devices attached to  Ntp type Feature Template
* `last_updated_by` - User who updated the  Ntp type Feature Template latest
* `last_updated_on` - Time for the latest Update of  Ntp type Feature Template
* `g_template_class` - Template Class type for  Ntp type Feature Template
* `template_id` - Unique ID for  Ntp type Feature Template
* `config_type` - Type of configuration for  Ntp type Feature Template
* `created_on` - Time of creation of  Ntp type Feature Template
* `rid` - rid for the  Ntp type Feature Template
* `feature` - Feature for the  Ntp type Feature Template

## Nested Schema for template_definition
* `server` - Server configuration for the  Ntp type Feature Template, (see [below for nested schema])
* `authentication` - Authentication configuration for the  Ntp type Feature Template, (see [below for nested schema])
* `trusted` - Trusted key for the  Ntp type Feature Template

## Nested Schema for server
* `hostname` - IP address of an NTP server, or a DNS server that knows how to reach the NTP server.
* `key` - Key Id for configuration of  Ntp type Feature Template
* `vpn` - Vpn ID for configuration of  Ntp type Feature Template
* `version` - Version number of the NTP protocol software for configuration of  Ntp type Feature Template
* `source_interface` - Name of a specific interface to use for outgoing NTP packets
* `prefer` - Boolean flag, Click On if multiple NTP servers are at the same stratum level and you want one to be preferred.

## Nested Schema for authentication
* `id` - MD5 authentication key ID for configuration of Ntp type Feature Template
* `value` - Key value for configuration of  Ntp type Feature Template
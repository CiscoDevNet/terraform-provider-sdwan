---
layout: "sdwan"
page_title: "SD-WAN: sdwan_ntp_feature_template"
sidebar_current: "docs-sdwan-resource-ntp_feature_template"
description: |-
  Manages SD-WAN Ntp Type Feature Templates
---

# sdwan_ntp_feature_template #
Manages SD-WAN Ntp Type Feature Templates

## Example Usage ##

```hcl
# example for Ntp Feature Template

resource "sdwan_ntp_feature_template" "example_ntp_feature_template" {  
  template_name = "example1"
  template_description = "For testing purposes"
  device_type = ["vedge-1000"]
  template_min_version = "15.0.0"  
  factory_default = false
}


# example for Ntp Feature Template

resource "sdwan_ntp_feature_template" "example_ntp_feature_template" {
  template_name = "example2"
  template_description = "For testing purposes"
  device_type = ["vedge-1000"]
  template_min_version = "15.0.0"
  template_definition {
    server {
      hostname = "time1.google."  
      key = 1    
      vpn = 0
      version = 4
      source_interface = "abc"
      prefer = false
    }
    
    server {
      hostname = "198.00.200.100"     
      key = 2
      vpn = 1
      version = 3 
    }

    authentication {
      id = 1
      value = "12345"
    }

    authentication {
      id = 2
      value = "xyz"
    }

    trusted = [ "2"]
  }
  factory_default = false
}
```


## Argument Reference ##

* `template_name` - (Required) Unique  Ntp Type Feature Template Name, Must not include <, >, !, &, ", or white space; maximum 128 characters
* `template_description` - (Required) Long Description of  Ntp type Feature Template, Should not be empty
* `device_type` - (Required) Type of device which supports  Feature Template, Allowed values: "vbond","vedge-100","vedge-100-B","vedge-100-M","vedge-100-WM","vedge-1000","vedge-2000","vedge-5000","vedge-cloud","vedge-ISR1100-4G","vedge-ISR1100-4GLTE","vedge-ISR1100-6G","vedge-ISR1100X-4G","vedge-ISR1100X-6G","vedge-nfvis-CSP-5444","vedge-nfvis-CSP-5456","vedge-nfvis-CSP2100","vedge-nfvis-CSP2100-X1","vedge-nfvis-CSP2100-X2","vedge-nfvis-UCSC-E","vedge-nfvis-UCSC-M5"
* `template_min_version` - (Required) Minimum Version for the  Feature template
* `factory_default` - (Required) Boolean flag indicating whether  Ntp type Feature Template is factory default or not
* `template_definition` - (Optional) Configuration for  Ntp type Feature Template, (see [below for nested schema])

## Attribute Reference ##

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
* `server` - (Optional) Server configuration for the  Ntp type Feature Template, see [below for nested schema])
* `authentication` - (Optional) Authentication configuration for the  Ntp type Feature Template, see [below for nested schema])
* `trusted` - (Optional) Enter the MD5 authentication key id to designate the key as trustworthy.

## Nested Schema for server
* `hostname` - (Required) IP address of an NTP server, or a DNS server that knows how to reach the NTP server.
* `key` - (Required) MD5 key associated with the NTP server, to enable MD5 authentication, Range: [1, 65535]
* `vpn` - (Required) Vpn ID for configuration of  Ntp type Feature Template, Range: [1, 65530]
* `version` - (Required) Version number of the NTP protocol software, Range: [1, 4]
* `source_interface` - (Optional) Name of a specific interface to use for outgoing NTP packets.
* `prefer` - (Optional) Boolean flag, Click On if multiple NTP servers are at the same stratum level and you want one to be preferred.

## Nested Schema for authentication
* `id` - (Required) MD5 authentication key ID for configuration of Ntp type Feature Template, Range: [1, 65535]
* `value` - (Required) Enter either a cleartext key or an AES-encrypted key for authentication of  Ntp server
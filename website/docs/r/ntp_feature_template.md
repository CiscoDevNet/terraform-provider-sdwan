---
layout: "sdwan"
page_title: "SD-WAN: sdwan_ntp_feature_template"
sidebar_current: "docs-sdwan-resource-ntp_feature_template"
description: |-
  Manages SD-WAN NTP Type Feature Templates
---

# sdwan_ntp_feature_template #
Manages SD-WAN NTP Type Feature Templates

## Example Usage ##

```hcl
# example for NTP Feature Template

resource "sdwan_ntp_feature_template" "example_ntp_feature_template" {  
  template_name = "example1"
  template_description = "For testing purposes"
  device_type = [	
    "vedge-1000",
		"vedge-2000",
		"vedge-cloud",
		"vedge-5000",
		]    
  template_type = "ntp"
  template_min_version = "15.0.0"
   template_definition {   
   
    server {
      hostname = "198.00.200.100"     
      key = 1
      vpn = 0
      version = 4 
    }

    authentication {
      id = 1
      value = "12345"
    }
   
    trusted = [ "1"]
  }
  factory_default = false
}


# example for Cisco NTP Feature Template

resource "sdwan_ntp_feature_template" "example_cisco_ntp_feature_template" {
  template_name = "example2"
  template_description = "For testing purposes"
  device_type = [
    "vedge-C1161-8P",
		"vedge-C1121X-8P",
		"vedge-C8200-1N-4T",
		"vedge-ISR-4331",
		"vedge-C1127X-8PMLTEP",
		"vedge-C1117-4PMLTEEAWE",
		"vedge-ISR-4451-X",
		"vedge-C1113-8PLTEEA",
		"vedge-IR-1821",
		"vedge-ASR-1001-X"]
  template_type = "cisco_ntp"
  template_min_version = "15.0.0"
  template_definition {      
    server {
      hostname = "time1.google."  
      key = 2    
      vpn = 5
      version = 3
      source_interface = "axyz"
      prefer = false
    }

    server {
      hostname = "198.00.200.100"     
      key = 1
      vpn = 0
      version = 4 
    }

    master {
      enable = true
      stratum = 5
      source = "abc"
    }

    authentication {
      id = 1
      value = "12345"
    }

    authentication {
      id = 2
      value = "xyzdf"
    }

    trusted = [ "2"]
  }
  factory_default = false
}
```


## Argument Reference ##

* `template_name` - (Required) Unique NTP Type Feature Template Name, Must not include <, >, !, &, ", or white space; maximum 128 characters
* `template_description` - (Required) Long Description of NTP type Feature Template, Should not be empty
* `device_type` - (Required) Type of device which supports NTP Feature Template, Allowed values for NTP: "vedge-1000", "vedge-2000", "vedge-cloud",
"vedge-5000", "vedge-ISR1100-6G", "vedge-100-B", "vedge-ISR1100-4G",
"vedge-100", "vsmart", "vedge-ISR1100-4GLTE", "vedge-100-WM", "vmanage",
"vbond", "vedge-100-M", "vedge-ISR1100X-6G", "vedge-ISR1100X-4G"
Allowed values for Cisco NTP: "vedge-C8500-12X4QC", "vedge-C1111-4PLTEEA",
"vedge-C1161-8P", "vedge-C1117-4PLTEEAW", "vedge-C1121X-8P",
"vedge-C8200-1N-4T", "vedge-ISR-4331", "vedge-C1127X-8PMLTEP","vedge-C1117-4PMLTEEAWE", "vedge-ISR-4451-X", "vedge-C1113-8PLTEEA",
"vedge-IR-1821", "vedge-ASR-1001-X", "vedge-ISR-4321","vedge-C1116-4PLTEEAWE", "vedge-C1109-4PLTE2P", "vedge-C1121-8P", "vedge-ASR-1002-HX", "vedge-C1111-8PLTEEAW", "vedge-C1112-8PWE", "vedge-C1101-4PLTEP", "vedge-ISR1100-4GLTENA-XE", "vedge-C1111-8PLTELA",
"vedge-IR-1835", "vedge-C1121X-8PLTEP", "vedge-IR-1833", "vedge-C8300-1N1S-4T2X", "vedge-C1121-4P", "vedge-ISR-4351",
"vedge-C1117-4PLTELA", "vedge-C1116-4PWE", "vedge-nfvis-C8200-UCPE",
"vedge-C1113-8PM", "vedge-IR-1831", "vedge-C1127-8PLTEP","vedge-C1121-8PLTEPW", "vedge-C1113-8PW", "vedge-ASR-1001-HX",
"vedge-C1128-8PLTEP", "vedge-C1113-8PLTEEAW", "vedge-C1117-4PW",
"vedge-C1116-4P", "vedge-C1113-8PMLTEEA","vedge-C1112-8P",
"vedge-ISR-4461", "vedge-C1116-4PLTEEA", "vedge-ISR-4221",
"vedge-C1117-4PM", "vedge-C1113-8PLTELAWZ", "vedge-C1117-4PMWE",
"vedge-C1109-2PLTEVZ", "vedge-C1113-8P", "vedge-C1117-4P",
"vedge-nfvis-ENCS5400", "vedge-C8300-2N2S-6T", "vedge-C1127-8PMLTEP",
"vedge-ESR-6300", "vedge-ISR-4221X", "vedge-ISR1100-4GLTEGB-XE",
"vedge-C8500-12X", "vedge-C1109-2PLTEGB", "vedge-CSR-1000v",
"vedge-C1113-8PLTEW", "vedge-C1121X-8PLTEPW", "vedge-ISR1100-6G-XE",
"vedge-C1121-4PLTEP", "vedge-C1111-8PLTEEA", "vedge-C1117-4PLTEEA",
"vedge-C1127X-8PLTEP", "vedge-C1109-2PLTEUS", "vedge-C1112-8PLTEEAWE",
"vedge-C1161X-8P", "vedge-C8500L-8S4X", "vedge-C1111-8PW",
"vedge-C1161X-8PLTEP", "vedge-C1101-4PLTEPW", "vedge-ISR1100X-4G-XE",
"vedge-IR-1101", "vedge-C1111-4P", "vedge-C1111-4PW",   "vedge-C1111-8P",
"vedge-C1117-4PMLTEEA", "vedge-C1113-8PLTELA", "vedge-C1111-8PLTELAW",
"vedge-C1161-8PLTEP", "vedge-ISR1100X-6G-XE", "vedge-ISR-4431",
"vedge-C1101-4P", "vedge-C1109-4PLTE2PW", "vedge-C1113-8PMWE",
"vedge-C1118-8P", "vedge-C1126-8PLTEP", "vedge-C8300-1N1S-6T",
"vedge-C1121-8PLTEP", "vedge-C8300-2N2S-4T2X", "vedge-C1112-8PLTEEA",
"vedge-C1111-4PLTELA", "vedge-ASR-1002-X", "vedge-C1111X-8P",
"vedge-C1126X-8PLTEP", "vedge-ASR-1006-X", "vedge-C8000V",
"vedge-ISR1100-4G-XE", "vedge-C1117-4PLTELAWZ", "vedge-ISRv"
* `template_type` - (Required) Type of Feature Template, Allowed values: "ntp", "cisco_ntp"
* `template_min_version` - (Required) Minimum Version for the Feature template
* `factory_default` - (Required) Boolean flag indicating whether NTP type Feature Template is factory default or not, If we set it true we can not update or delete resource.
* `template_definition` - (Required) Configuration for NTP type Feature Template, (see [below for nested schema])

## Attribute Reference ##

* `attached_masters_count` - Number of Device Templates attached to NTP type Feature Template
* `devices_attached` - Number of Devices attached to NTP type Feature Template
* `last_updated_by` - User who updated the NTP type Feature Template latest
* `last_updated_on` - Time for the latest Update of NTP type Feature Template
* `g_template_class` - Template Class type for NTP type Feature Template
* `template_id` - Unique ID for NTP type Feature Template
* `config_type` - Type of configuration for NTP type Feature Template
* `created_on` - Time of creation of NTP type Feature Template
* `rid` - rid for the NTP type Feature Template
* `feature` - Feature for the NTP type Feature Template

## Nested Schema for template_definition
* `server` - (Optional) Server configuration for the NTP type Feature Template, see [below for nested schema])
* `master` - (Optional) Master configuration for the Cisco NTP type Feature Template, see [below for nested schema])
* `authentication` - (Optional) Authentication configuration for the NTP type Feature Template, see [below for nested schema])
* `trusted` - (Optional) Enter the MD5 authentication key id to designate the key as trustworthy

## Nested Schema for server
* `hostname` - (Required) IP address of an NTP server, or a DNS server that knows how to reach the NTP server
* `key` - (Required) MD5 key associated with the NTP server, to enable MD5 authentication, Range: [1, 65535]
* `vpn` - (Required) Vpn ID for configuration of NTP type Feature Template, Range: [1, 65530]
* `version` - (Required) Version number of the NTP protocol software, Range: [1, 4], Default: 4
* `source_interface` - (Optional) Name of a specific interface to use for outgoing NTP packets
* `prefer` - (Optional) It is a boolean flag, set true if multiple NTP servers are at the same stratum level and you want one to be preferred

## Nested Schema for master
* `enable` - (Optional) Boolean flag indicating whether master is enabled for configuration of Cisco NTP type Feature Template
* `stratum` - (Optional) Stratum for configuration of Cisco NTP type Feature Template, Range: [1, 15] 
* `source` - (Optional) Source for configuration of Cisco NTP type Feature Template

## Nested Schema for authentication
* `id` - (Required) MD5 authentication key ID for configuration of NTP type Feature Template, Range: [1, 65535]
* `value` - (Required) It is either a cleartext key or an AES-encrypted key for authentication of NTP server
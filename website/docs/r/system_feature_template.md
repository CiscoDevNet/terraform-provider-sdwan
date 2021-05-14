---
layout: "sdwan"
page_title: "SD-WAN: sdwan_system_feature_template"
sidebar_current: "docs-sdwan-resource-system_feature_template"
description: |-
  Manages SD-WAN System Type Feature Templates
---

# sdwan_system_feature_template #
Manages SD-WAN System Type Feature Templates

## Example Usage ##

```hcl
# example for System Feature Template

resource "sdwan_system_feature_template" "example_system_feature_template" {  
  template_name = "example1"
  template_description = "For testing purposes"
  device_type = ["vedge-1000"]
  template_min_version = "15.0.0"
  template_definition {
    system_basic {  
      site_id = 5
      system_ip = "198.12.23.178"
      timezone = "UTC"
    }
  }
  factory_default = false
}


# example for System Feature Template

resource "sdwan_system_feature_template" "example_system_feature_template" {
  template_name = "example2"
  template_description = "For testing purposes"
  device_type = ["vedge-1000"]
  template_min_version = "15.0.0"
  template_definition {
    system_basic {  
      site_id = 5
      system_ip = "198.12.23.178"
      overlay_id = 2
      timezone = "Asia/Dubai"
      hostname = "hostname"
      location = "Place"
      device_groups = [ "abc","def" ]
      controller_groups = [ "12","34"]
      description = "some description"
      console_baud_rate = "19200"
      max_omp_sessions = 3
      tcp_optimization_enabled = true
    }

  system_gps {
    longitude = 23
    latitude = 19
  }

  system_trackers {
    tracker_name = "abcd"
    tracker_type = "ip-address"
    tracker_endpoint_value = "198.23.34.67"
  }
  
  system_trackers {
    tracker_name = "tracker2"
    tracker_type = "dns-name"
    tracker_endpoint_value = "dnsName"
  }
  system_trackers {
    tracker_name = "abc"
    tracker_type = "api-url"
    tracker_endpoint_value = "cisco.com/docs"
  }
  system_advanced {
    control_session_pps = 10
    system_tunnel_mtu = 560
    port_hop = false
    port_offset = 5
    dns_cache_timeout = 6
    track_transport = false
    vbond_local = true
    vbond_remote = "abcd"
    track_interface_tag = 433
    multicast_buffer_percent = 56
    usb_controller = true
    track_default_gateway = false
    host_policer_pps = 5689
    icmp_error_pps = 67
    allow_same_site_tunnels = true
    route_consistency_check = true
    admin_tech_on_failure = false
    idle_timeout = 67
    eco_friendly_mode = true
    on_demand_enable = true
    on_demand_idle_timeout = 78
  }
  }
  factory_default = false
}
```


## Argument Reference ##

* `template_name` - (Required) Unique  System Type Feature Template Name, Must not include <, >, !, &, ", or white space; maximum 128 characters
* `template_description` - (Required) Long Description of  System type Feature Template, Should not be empty
* `device_type` - (Required) Type of device which supports  Feature Template, Allowed values: "vbond","vedge-100","vedge-100-B","vedge-100-M","vedge-100-WM","vedge-1000","vedge-2000","vedge-5000","vedge-cloud","vedge-ISR1100-4G","vedge-ISR1100-4GLTE","vedge-ISR1100-6G","vedge-ISR1100X-4G","vedge-ISR1100X-6G","vedge-nfvis-CSP-5444","vedge-nfvis-CSP-5456","vedge-nfvis-CSP2100","vedge-nfvis-CSP2100-X1","vedge-nfvis-CSP2100-X2","vedge-nfvis-UCSC-E","vedge-nfvis-UCSC-M5"
* `template_min_version` - (Required) Minimum Version for the  Feature template
* `template_definition` - (Required) Configuration for  System type Feature Template, (see [below for nested schema])
* `factory_default` - (Required) Boolean flag indicating whether  System type Feature Template is factory default or not

## Attribute Reference ##

* `attached_masters_count` - Number of Device Templates attached to  System type Feature Template
* `devices_attached` - Number of Devices attached to  System type Feature Template
* `last_updated_by` - User who updated the  System type Feature Template latest
* `last_updated_on` - Time for the latest Update of  System type Feature Template
* `g_template_class` - Template Class type for  System type Feature Template
* `template_id` - Unique ID for  System type Feature Template
* `config_type` - Type of configuration for  System type Feature Template
* `created_on` - Time of creation of  System type Feature Template
* `rid` - rid for the  System type Feature Template
* `feature` - Feature for the  System type Feature Template

## Nested Schema for template_definition
* `system_basic` - (Required) Basic configuration for the  System type Feature Template, see [below for nested schema])
* `system_gps` - (Optional) GPS configuration for the  System type Feature Template, see [below for nested schema])
* `system_trackers` - (Optional) Trackers for the  System type Feature Template, see [below for nested schema])
* `system_advanced` - (Optional) Advanced configuration for the  System type Feature Template, see [below for nested schema])

## Nested Schema for system_basic
* `site_id` - (Required) Site ID for configuration of System type Feature Template, Range: [1, 4294967295]
* `system_ip` - (Required) System IP for configuration of  System type Feature Template, allowed values: IPv4 address
* `timezone` - (Required) Timezone for configuration of  System type Feature Template
* `overlay_id` - (Optional) Overlay ID for configuration of  System type Feature Template, Deafult: 1
* `hostname` - (Optional)  Hostname for configuration of  System type Feature Template, String length: [1, 32] 
* `location` - (Optional) Location for configuration of  System type Feature Template, String length: [1, 128] 
* `device_groups` - (Optional) Device Groups for configuration of  System type Feature Template
* `controller_groups` - (Optional) Controller Group List for configuration of  System type Feature Template
* `description` - (Optional) Description for configuration of  System type Feature Template
* `console_baud_rate` - (Optional) Console Baud Rate for configuration of  System type Feature Template, Allowed values: "1200","2400","4800","9600","19200","38400","57600","115200"
* `max_omp_sessions` - (Optional) Maximum OMP Sessions for configuration of  System type Feature Template, Range: [0, 100], Deafult: 2
* `tcp_optimization_enabled` - (Optional) Boolean flag indicating whether  TCP Optimization is enabled for configuration of  System type Feature Template, Default: false

## Nested Schema for system_gps
* `latitude` - (Optional) Latitude for configuration of System type Feature Template
* `longitude` - (Optional) Longitude for configuration of  System type Feature Template

## Nested Schema for system_trackers
* `tracker_name` - (Required) Name of Tracker for configuration of System type Feature Template, Must not include <, >, !, &, ", or white space; maximum 128 characters
* `tracker_type` - (Required) Type of Tracker for configuration of System type Feature Template, Allowed values: "ip-address", "dns-name", "api-url"
* `tracker_endpoint_value` - (Required) Value of Tracker Endpoint for configuration of System type Feature Template
* `tracker_threshold` - (Optional) Threshold of Tracker for configuration of System type Feature Template, Range: [100, 1000], Default: 300
* `tracker_interval` - (Optional) Interval of Tracker for configuration of System type Feature Template, Range: [10, 600], Default: 60
* `tracker_multiplier` - (Optional) Multiplier of Tracker for configuration of System type Feature Template, Range: [1, 10], Default: 3

## Nested Schema for system_advanced
* `control_session_pps` - (Optional) Policer Rate(pps) for Control Sessions for configuration of System type Feature Template, Range: [1, 65535], Default: 300
* `system_tunnel_mtu` - (Optional) MTU of System's internal DTLS Tunnel for configuration of System type Feature Template, Range: [500, 2000], Default: 1024
* `port_hop` - (Optional) Boolean flag to indicate whether Port Hopping is enabled for configuration of System type Feature Template, Default: true
* `port_offset` - (Optional) TLOC Port Offset for configuration of System type Feature Template, Range: [0, 19]
* `dns_cache_timeout` - (Optional) DNS Cache Timeout(minutes) for configuration of System type Feature Template, Range: [1, 30], Default: 2
* `track_transport` - (Optional) Boolean flag to indicate whether tracking of transport is enabled for configuration of System type Feature Template, Default: true
* `vbond_local` - (Optional) Boolean flag to indicate whether the local device is set as vBond for configuration of System type Feature Template, Default: false
* `vbond_remote` - (Optional) vBond IP address for configuration of System type Feature Template
* `track_interface_tag` - (Optional) Interface tracking for configuration of System type Feature Template, Range: [1, 4294967295]
* `multicast_buffer_percent` - (Optional) Percentage of buffer multicast packets can consume for configuration of System type Feature Template, Range: [5, 100]
* `usb_controller` - (Optional) Boolean flag to indicate whether external USB Controller on the device is enabled for configuration of System type Feature Template, Default: false
* `track_default-gateway` - (Optional) Boolean flag to indicate whether default Gateway Tracking on the device is enabled for configuration of System type Feature Template, Default: true
* `host_policer_pps` - (Optional) Rate(1000-25000) at which to police packets bound to the control plane for configuration of System type Feature Template, Range: [1000, 25000], Default: 20000
* `icmp_error_pps` - (Optional) Rate at which to police ICMP error messages for configuration of System type Feature Template, Range: [1, 200], Default: 100
* `allow_same_site_tunnels` - (Optional) Boolean flag to indicate whether tunnels are allowed to be formed between vEdges in the same site for configuration of System type Feature Template, Default: false
* `route_consistency_check` - (Optional) Boolean flag to indicate whether route consistency check is enabled for configuration of System type Feature Template, Default: false
* `admin_tech_on_failure` - (Optional) Boolean flag to indicate whether the collection of admin-tech before reboot due to daemon failure is enabled for configuration of System type Feature Template, Default: true
* `idle_timeout` - (Optional) Idle CLI timeout(minutes) for configuration of System type Feature Template, Range: [0, 300]
* `eco_friendly_mode` - (Optional) Boolean flag to indicate whether vEdge Cloud router can run in the eco-friendly mode for configuration of System type Feature Template, Default: false
* `on_demand_enable` - (Optional) Boolean flag to indicate whether On-Demand Tunnel is enabled for configuration of System type Feature Template, Default: false
* `on_demand_idle_timeout` - (Optional) Idle timeout for On-demand Tunnels for configuration of System type Feature Template, Default: 10



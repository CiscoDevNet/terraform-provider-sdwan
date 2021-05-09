---
layout: "sdwan"
page_title: "SD-WAN: sdwan_system_feature_template"
sidebar_current: "docs-sdwan-data-source-system_feature_template"
description: |-
  Data source for SD-WAN System Feature Template
---

# sdwan_system_feature_template #
Data source for SD-WAN System Feature Templates

## Example Usage ##

```hcl

data "sdwan_system_feature_template" "name" {
    template_name = "example_template"
}

```


## Argument Reference ##

* `template_name` - (Required) Unique  System Type Feature Template Name

## Attribute Reference ##

* `template_description` - Long Description of  System type Feature Template
* `device_type` - Type of device which supports  Feature Template
* `template_min_version` - Minimum Version for the  Feature template
* `template_definition` - Configuration for  System type Feature Template, (see [below for nested schema])
* `factory_default` - Boolean flag indicating whether  System type Feature Template is factory default or not
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
* `system_basic` - Basic configuration for the  System type Feature Template, (see [below for nested schema])
* `system_gps` - GPS configuration for the  System type Feature Template, (see [below for nested schema])
* `system_trackers` - Trackers for the  System type Feature Template, (see [below for nested schema])
* `system_advanced` - Advanced configuration for the  System type Feature Template, (see [below for nested schema])

## Nested Schema for system_basic
* `site_id` - Site ID for configuration of System type Feature Template
* `system_ip` - System IP for configuration of  System type Feature Template
* `overlay_id` - Overlay ID for configuration of  System type Feature Template
* `timezone` - Timezone for configuration of  System type Feature Template
* `hostname` - Hostname for configuration of  System type Feature Template
* `location` - Location for configuration of  System type Feature Template
* `device_groups` - Device Groups for configuration of  System type Feature Template
* `controller_groups` - Controller Group List for configuration of  System type Feature Template
* `description` - Description for configuration of  System type Feature Template
* `console_baud_rate` - Console Baud Rate for configuration of  System type Feature Template
* `max_omp_sessions` - Maximum OMP Sessions for configuration of  System type Feature Template
* `tcp_optimization_enabled` - Boolean flag indicating whether  TCP Optimization is enabled for configuration of  System type Feature Template

## Nested Schema for system_gps
* `latitude` - Latitude for configuration of System type Feature Template
* `longitude` - Longitude for configuration of  System type Feature Template

## Nested Schema for system_trackers
* `tacker_name` - Name of Tracker for configuration of System type Feature Template
* `tacker_type` - Type of Tracker for configuration of System type Feature Template
* `tacker_endpoint_value` - Value of Tracker Endpoint for configuration of System type Feature Template
* `tacker_threshold` - Threshold of Tracker for configuration of System type Feature Template
* `tacker_interval` - Interval of Tracker for configuration of System type Feature Template
* `tacker_multiplier` - Multiplier of Tracker for configuration of System type Feature Template

## Nested Schema for system_advanced
* `control_session_pps` - Policer Rate(pps) for Control Sessions for configuration of System type Feature Template
* `system_tunnel_mtu` - MTU of System's internal DTLS Tunnel for configuration of System type Feature Template
* `port_hop` - Boolean flag to indicate whether Port Hopping is enabled for configuration of System type Feature Template
* `port_offset` - TLOC Port Offset for configuration of System type Feature Template
* `dns_cache_timeout` - DNS Cache Timeout(minutes) for configuration of System type Feature Template
* `track_transport` - Boolean flag to indicate whether tracking of transport is enabled for configuration of System type Feature Template
* `vbond_local` - Boolean flag to indicate whether the local device is set as vBond for configuration of System type Feature Template
* `vbond_remote` - vBond IP address for configuration of System type Feature Template
* `track_interface_tag` - Interface tracking for configuration of System type Feature Template
* `multicast_buffer_percent` - Percentage(5-100) of buffer multicast packets can consume for configuration of System type Feature Template
* `usb_controller` - Boolean flag to indicate whether external USB Controller on the device is enabled for configuration of System type Feature Template
* `track_default-gateway` - Boolean flag to indicate whether default Gateway Tracking on the device is enabled for configuration of System type Feature Template
* `host_policer_pps` - Rate(1000-25000) at which to police packets bound to the control plane for configuration of System type Feature Template
* `icmp_error_pps` - Rate at which to police ICMP error messages for configuration of System type Feature Template
* `allow_same_site_tunnels` - Boolean flag to indicate whether tunnels are allowed to be formed between vEdges in the same site for configuration of System type Feature Template
* `route_consistency_check` - Boolean flag to indicate whether route consistency check is enabled for configuration of System type Feature Template
* `admin_tech_on_failure` - Boolean flag to indicate whether the collection of admin-tech before reboot due to daemon failure is enabled for configuration of System type Feature Template
* `idle_timeout` - Idle CLI timeout(minutes) for configuration of System type Feature Template
* `eco_friendly_mode` - Boolean flag to indicate whether vEdge Cloud router can run in the eco-friendly mode for configuration of System type Feature Template
* `on_demand_enable` - Boolean flag to indicate whether On-Demand Tunnel is enabled for configuration of System type Feature Template
* `on_demand_idle_timeout` - Idle timeout for On-demand Tunnels for configuration of System type Feature Template


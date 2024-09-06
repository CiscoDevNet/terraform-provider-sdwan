---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_system_omp_feature Resource - terraform-provider-sdwan"
subcategory: "Feature"
description: |-
  This resource can manage a System OMP Feature.
  Minimum SD-WAN Manager version: 20.12.0
---

# sdwan_system_omp_feature (Resource)

This resource can manage a System OMP Feature.
  - Minimum SD-WAN Manager version: `20.12.0`

## Example Usage

```terraform
resource "sdwan_system_omp_feature" "example" {
  name                        = "Example"
  description                 = "My Example"
  feature_profile_id          = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  graceful_restart            = true
  overlay_as                  = 10
  paths_advertised_per_prefix = 4
  ecmp_limit                  = 4
  shutdown                    = false
  omp_admin_distance_ipv4     = 10
  omp_admin_distance_ipv6     = 20
  advertisement_interval      = 1
  graceful_restart_timer      = 43200
  eor_timer                   = 300
  holdtime                    = 60
  advertise_ipv4_bgp          = false
  advertise_ipv4_ospf         = false
  advertise_ipv4_ospf_v3      = false
  advertise_ipv4_connected    = false
  advertise_ipv4_static       = false
  advertise_ipv4_eigrp        = false
  advertise_ipv4_lisp         = false
  advertise_ipv4_isis         = false
  advertise_ipv6_bgp          = true
  advertise_ipv6_ospf         = true
  advertise_ipv6_connected    = true
  advertise_ipv6_static       = true
  advertise_ipv6_eigrp        = true
  advertise_ipv6_lisp         = true
  advertise_ipv6_isis         = true
  ignore_region_path_length   = false
  transport_gateway           = "prefer"
  site_types                  = ["type-1"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `name` (String) The name of the Feature

### Optional

- `advertise_ipv4_bgp` (Boolean) BGP
  - Default value: `false`
- `advertise_ipv4_bgp_variable` (String) Variable name
- `advertise_ipv4_connected` (Boolean) Connected
  - Default value: `true`
- `advertise_ipv4_connected_variable` (String) Variable name
- `advertise_ipv4_eigrp` (Boolean) EIGRP
  - Default value: `false`
- `advertise_ipv4_eigrp_variable` (String) Variable name
- `advertise_ipv4_isis` (Boolean) ISIS
  - Default value: `false`
- `advertise_ipv4_isis_variable` (String) Variable name
- `advertise_ipv4_lisp` (Boolean) LISP
  - Default value: `false`
- `advertise_ipv4_lisp_variable` (String) Variable name
- `advertise_ipv4_ospf` (Boolean) OSPF
  - Default value: `false`
- `advertise_ipv4_ospf_v3` (Boolean) OSPFV3
  - Default value: `false`
- `advertise_ipv4_ospf_v3_variable` (String) Variable name
- `advertise_ipv4_ospf_variable` (String) Variable name
- `advertise_ipv4_static` (Boolean) Static
  - Default value: `true`
- `advertise_ipv4_static_variable` (String) Variable name
- `advertise_ipv6_bgp` (Boolean) BGP
  - Default value: `false`
- `advertise_ipv6_bgp_variable` (String) Variable name
- `advertise_ipv6_connected` (Boolean) Connected
  - Default value: `false`
- `advertise_ipv6_connected_variable` (String) Variable name
- `advertise_ipv6_eigrp` (Boolean) EIGRP
  - Default value: `false`
- `advertise_ipv6_eigrp_variable` (String) Variable name
- `advertise_ipv6_isis` (Boolean) ISIS
  - Default value: `false`
- `advertise_ipv6_isis_variable` (String) Variable name
- `advertise_ipv6_lisp` (Boolean) LISP
  - Default value: `false`
- `advertise_ipv6_lisp_variable` (String) Variable name
- `advertise_ipv6_ospf` (Boolean) OSPF
  - Default value: `false`
- `advertise_ipv6_ospf_variable` (String) Variable name
- `advertise_ipv6_static` (Boolean) Static
  - Default value: `false`
- `advertise_ipv6_static_variable` (String) Variable name
- `advertisement_interval` (Number) Advertisement Interval (seconds)
  - Range: `0`-`65535`
  - Default value: `1`
- `advertisement_interval_variable` (String) Variable name
- `description` (String) The description of the Feature
- `ecmp_limit` (Number) Set maximum number of OMP paths to install in cEdge route table
  - Range: `1`-`0`
  - Default value: `4`
- `ecmp_limit_variable` (String) Variable name
- `eor_timer` (Number) EOR Timer
  - Range: `1`-`3600`
  - Default value: `300`
- `eor_timer_variable` (String) Variable name
- `graceful_restart` (Boolean) Graceful Restart for OMP
  - Default value: `true`
- `graceful_restart_timer` (Number) Graceful Restart Timer (seconds)
  - Range: `1`-`604800`
  - Default value: `43200`
- `graceful_restart_timer_variable` (String) Variable name
- `graceful_restart_variable` (String) Variable name
- `holdtime` (Number) Hold Time (seconds)
  - Default value: `60`
- `holdtime_variable` (String) Variable name
- `ignore_region_path_length` (Boolean) Treat hierarchical and direct (secondary region) paths equally
  - Default value: `false`
- `ignore_region_path_length_variable` (String) Variable name
- `omp_admin_distance_ipv4` (Number) OMP Admin Distance IPv4
  - Range: `1`-`255`
  - Default value: `251`
- `omp_admin_distance_ipv4_variable` (String) Variable name
- `omp_admin_distance_ipv6` (Number) OMP Admin Distance IPv6
  - Range: `1`-`255`
  - Default value: `251`
- `omp_admin_distance_ipv6_variable` (String) Variable name
- `overlay_as` (Number) Overlay AS Number
  - Range: `1`-`4294967295`
- `overlay_as_variable` (String) Variable name
- `paths_advertised_per_prefix` (Number) Number of Paths Advertised per Prefix
  - Range: `1`-`16`
  - Default value: `4`
- `paths_advertised_per_prefix_variable` (String) Variable name
- `shutdown` (Boolean) Shutdown
  - Default value: `false`
- `shutdown_variable` (String) Variable name
- `site_types` (Set of String) Site Types
- `site_types_variable` (String) Variable name
- `transport_gateway` (String) Transport Gateway Path Behavior
  - Choices: `prefer`, `ecmp-with-direct-path`
- `transport_gateway_variable` (String) Variable name

### Read-Only

- `id` (String) The id of the Feature
- `version` (Number) The version of the Feature

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_system_omp_feature.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```
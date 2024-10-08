---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_switchport_feature_template Data Source - terraform-provider-sdwan"
subcategory: "(Classic) Feature Templates"
description: |-
  This data source can read the Switchport feature template.
---

# sdwan_switchport_feature_template (Data Source)

This data source can read the Switchport feature template.

## Example Usage

```terraform
data "sdwan_switchport_feature_template" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) The id of the feature template
- `name` (String) The name of the feature template

### Read-Only

- `age_out_time` (Number) Set when a MAC table entry ages out (0 to disable, 10-1000000 otherwise)
- `age_out_time_variable` (String) Variable name
- `description` (String) The description of the feature template
- `device_types` (Set of String) List of supported device types
- `interfaces` (Attributes List) Interface name: GigabitEthernet0/<>/<> when present (see [below for nested schema](#nestedatt--interfaces))
- `module_type` (String) Module type
- `slot` (Number) Number of Slots
- `static_mac_addresses` (Attributes List) Add static MAC address entries for interface (see [below for nested schema](#nestedatt--static_mac_addresses))
- `sub_slot` (Number) Number of Sub-Slots
- `template_type` (String) The template type
- `version` (Number) The version of the feature template

<a id="nestedatt--interfaces"></a>
### Nested Schema for `interfaces`

Read-Only:

- `dot1x_authentication_order` (Set of String) Specify authentication methods in the order of preference
- `dot1x_authentication_order_variable` (String) Variable name
- `dot1x_control_direction` (String) Set uni or bi directional authorization mode
- `dot1x_control_direction_variable` (String) Variable name
- `dot1x_critical_vlan` (Number) Set Critical VLAN
- `dot1x_critical_vlan_variable` (String) Variable name
- `dot1x_enable` (Boolean) Set 802.1x on off
- `dot1x_enable_criticial_voice_vlan` (Boolean) Enable Critical Voice VLAN
- `dot1x_enable_criticial_voice_vlan_variable` (String) Variable name
- `dot1x_enable_periodic_reauth` (Boolean) Enable Periodic Reauthentication
- `dot1x_enable_periodic_reauth_variable` (String) Variable name
- `dot1x_enable_variable` (String) Variable name
- `dot1x_guest_vlan` (Number) Set vlan to drop non-802.1x enabled clients into if client is not in MAB list
- `dot1x_guest_vlan_variable` (String) Variable name
- `dot1x_host_mode` (String) Set host mode
- `dot1x_host_mode_variable` (String) Variable name
- `dot1x_mac_authentication_bypass` (Boolean) MAC Authentication Bypass
- `dot1x_mac_authentication_bypass_variable` (String) Variable name
- `dot1x_pae_enable` (Boolean) Set 802.1x Interface Pae Type
- `dot1x_pae_enable_variable` (String) Variable name
- `dot1x_periodic_reauth_inactivity_timeout` (Number) Periodic Reauthentication Inactivity Timeout (in seconds)
- `dot1x_periodic_reauth_inactivity_timeout_variable` (String) Variable name
- `dot1x_periodic_reauth_interval` (Number) Periodic Reauthentication Interval (in seconds)
- `dot1x_periodic_reauth_interval_variable` (String) Variable name
- `dot1x_port_control` (String) Set Port-Control Mode
- `dot1x_port_control_variable` (String) Variable name
- `dot1x_restricted_vlan` (Number) Set Restricted VLAN ID
- `dot1x_restricted_vlan_variable` (String) Variable name
- `duplex` (String) Duplex mode
- `duplex_variable` (String) Variable name
- `name` (String) Set Interface name
- `name_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `shutdown` (Boolean) Administrative state
- `shutdown_variable` (String) Variable name
- `speed` (String) Set interface speed
- `speed_variable` (String) Variable name
- `switchport_access_vlan` (Number) Set VLAN identifier associated with bridging domain
- `switchport_access_vlan_variable` (String) Variable name
- `switchport_mode` (String) Set type of switch port: access/trunk
- `switchport_trunk_allowed_vlans` (String) Configure VLAN IDs used with the trunk
- `switchport_trunk_allowed_vlans_variable` (String) Variable name
- `switchport_trunk_native_vlan` (Number) Configure VLAN ID used for native VLAN
- `switchport_trunk_native_vlan_variable` (String) Variable name
- `voice_vlan` (Number) Configure Voice Vlan
- `voice_vlan_variable` (String) Variable name


<a id="nestedatt--static_mac_addresses"></a>
### Nested Schema for `static_mac_addresses`

Read-Only:

- `if_name` (String) Interface name: GigabitEthernet0/<>/<>
- `if_name_variable` (String) Variable name
- `mac_address` (String) Set MAC address in xxxx.xxxx.xxxx format
- `mac_address_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `vlan` (Number) Configure VLAN ID used with the mac and interface
- `vlan_variable` (String) Variable name

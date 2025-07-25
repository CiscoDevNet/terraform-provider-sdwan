---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_traffic_data_policy_definition Data Source - terraform-provider-sdwan"
subcategory: "(Classic) Centralized Policies"
description: |-
  This data source can read the Traffic Data Policy Definition .
---

# sdwan_traffic_data_policy_definition (Data Source)

This data source can read the Traffic Data Policy Definition .

## Example Usage

```terraform
data "sdwan_traffic_data_policy_definition" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object

### Read-Only

- `default_action` (String) Default action, either `accept` or `drop`
- `description` (String) The description of the policy definition
- `name` (String) The name of the policy definition
- `sequences` (Attributes List) List of sequences (see [below for nested schema](#nestedatt--sequences))
- `type` (String) Type
- `version` (Number) The version of the object

<a id="nestedatt--sequences"></a>
### Nested Schema for `sequences`

Read-Only:

- `action_entries` (Attributes List) List of action entries (see [below for nested schema](#nestedatt--sequences--action_entries))
- `base_action` (String) Base action, either `accept` or `drop`
- `id` (Number) Sequence ID
- `ip_type` (String) Sequence IP type, either `ipv4`, `ipv6` or `all`
- `match_entries` (Attributes List) List of match entries (see [below for nested schema](#nestedatt--sequences--match_entries))
- `name` (String) Sequence name
- `type` (String) Sequence type

<a id="nestedatt--sequences--action_entries"></a>
### Nested Schema for `sequences.action_entries`

Read-Only:

- `cflowd` (Boolean) Enable cflowd
- `counter` (String) Counter name
- `dre_optimization` (Boolean) Enable DRE optimization
- `fallback_to_routing` (Boolean) Enable fallback to routing
- `log` (Boolean) Enable logging
- `loss_correction` (String) Loss correction
- `loss_correction_fec` (String) Loss correction FEC
- `loss_correction_fec_threshold` (String) Loss correction FEC threshold
- `loss_correction_packet_duplication` (String) Loss correction packet duplication
- `nat_parameters` (Attributes List) List of NAT parameters (see [below for nested schema](#nestedatt--sequences--action_entries--nat_parameters))
- `nat_pool` (String) NAT pool
- `nat_pool_id` (Number) NAT pool ID
- `redirect_dns` (String) Redirect DNS
- `redirect_dns_address` (String) Redirect DNS IP address
- `redirect_dns_type` (String) Redirect DNS type
- `secure_internet_gateway` (Boolean) Enable secure internet gateway
- `service_node_group` (String) Service node group
- `set_parameters` (Attributes List) List of set parameters (see [below for nested schema](#nestedatt--sequences--action_entries--set_parameters))
- `tcp_optimization` (Boolean) Enable TCP optimization
- `type` (String) Type of action entry

<a id="nestedatt--sequences--action_entries--nat_parameters"></a>
### Nested Schema for `sequences.action_entries.nat_parameters`

Read-Only:

- `fallback` (Boolean) Fallback
- `type` (String) Type of NAT parameter
- `vpn_id` (Number) DSCP


<a id="nestedatt--sequences--action_entries--set_parameters"></a>
### Nested Schema for `sequences.action_entries.set_parameters`

Read-Only:

- `dscp` (Number) DSCP
- `forwarding_class` (String) Forwarding class
- `local_tloc_list_color` (String) Local TLOC list color. Space separated list of colors.
- `local_tloc_list_encap` (String) Local TLOC list encapsulation.
- `local_tloc_list_restrict` (Boolean) Local TLOC list restrict
- `next_hop` (String) Next hop IP
- `next_hop_loose` (Boolean) Use routing table entry to forward the packet in case Next-hop is not available
- `policer_list_id` (String) Policer list ID
- `policer_list_version` (Number) Policer list version
- `preferred_color_group_list` (String) Preferred color group list ID
- `preferred_color_group_list_version` (Number) Preferred color group list version
- `service_tloc_color` (String) Service TLOC color
- `service_tloc_encapsulation` (String) Service TLOC encapsulation
- `service_tloc_ip` (String) Service TLOC IP address
- `service_tloc_list_id` (String) Service TLOC list ID
- `service_tloc_list_version` (Number) Service TLOC list version
- `service_tloc_local` (Boolean) Service TLOC Local
- `service_tloc_restrict` (Boolean) Service TLOC Restrict
- `service_type` (String) Service type
- `service_vpn_id` (Number) Service VPN ID
- `tloc_color` (String) TLOC color
- `tloc_encapsulation` (String) TLOC encapsulation
- `tloc_ip` (String) TLOC IP address
- `tloc_list_id` (String) TLOC list ID
- `tloc_list_version` (Number) TLOC list version
- `type` (String) Type of set parameter
- `vpn_id` (Number) DSCP



<a id="nestedatt--sequences--match_entries"></a>
### Nested Schema for `sequences.match_entries`

Read-Only:

- `application_list_id` (String) Application list ID
- `application_list_version` (Number) Application list version
- `destination_data_prefix_list_id` (String) Destination Data Prefix list ID
- `destination_data_prefix_list_version` (Number) Destination Data Prefix list version
- `destination_ip` (String) Destination IP
- `destination_port` (String) Destination port, 0-65535 (Single value, range or multiple values separated by spaces)
- `destination_region` (String) Destination region
- `dns` (String) DNS request or response
- `dns_application_list_id` (String) DNS Application list ID
- `dns_application_list_version` (Number) DNS Application list version
- `dscp` (String) DSCP value
- `icmp_message` (String) ICMP Message
- `packet_length` (Number) Packet length
- `plp` (String) PLP
- `protocol` (String) IP Protocol, 0-255 (Single value or multiple values separated by spaces)
- `source_data_prefix_list_id` (String) Source Data Prefix list ID
- `source_data_prefix_list_version` (Number) Source Data Prefix list version
- `source_ip` (String) Source IP
- `source_port` (String) Source port, 0-65535 (Single value, range or multiple values separated by spaces)
- `tcp` (String) TCP flags
- `traffic_to` (String) Traffic to
- `type` (String) Type of match entry

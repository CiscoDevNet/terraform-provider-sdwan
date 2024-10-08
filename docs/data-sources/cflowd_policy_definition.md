---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_cflowd_policy_definition Data Source - terraform-provider-sdwan"
subcategory: "(Classic) Centralized Policies"
description: |-
  This data source can read the Cflowd Policy Definition .
---

# sdwan_cflowd_policy_definition (Data Source)

This data source can read the Cflowd Policy Definition .

## Example Usage

```terraform
data "sdwan_cflowd_policy_definition" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object

### Read-Only

- `active_flow_timeout` (Number) Active flow timeout in seconds
- `collectors` (Attributes List) List of collectors (see [below for nested schema](#nestedatt--collectors))
- `description` (String) The description of the policy definition
- `flow_refresh` (Number) Flow refresh in seconds
- `inactive_flow_timeout` (Number) Inactive flow timeout in seconds
- `name` (String) The name of the policy definition
- `protocol` (String) Protocol, either `ipv4`, `ipv6` or `all`
- `remarked_dscp` (Boolean) Collect remarked DSCP
- `sampling_interval` (Number) Flow sampling interval
- `tos` (Boolean) Collect TOS record field
- `type` (String) Type
- `version` (Number) The version of the object

<a id="nestedatt--collectors"></a>
### Nested Schema for `collectors`

Read-Only:

- `bfd_metrics_exporting` (Boolean) BFD metrics exporting
- `export_spreading` (String) Export spreading
- `exporting_interval` (Number) Exporting interval
- `ip_address` (String) IP address
- `port` (Number) Port
- `source_interface` (String) Source interface
- `transport` (String) Transport protocol
- `vpn_id` (Number) VPN ID

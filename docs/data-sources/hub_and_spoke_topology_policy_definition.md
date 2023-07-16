---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_hub_and_spoke_topology_policy_definition Data Source - terraform-provider-sdwan"
subcategory: "Centralized Policies"
description: |-
  This data source can read the Hub and Spoke Topology policy definition.
---

# sdwan_hub_and_spoke_topology_policy_definition (Data Source)

This data source can read the Hub and Spoke Topology policy definition.

## Example Usage

```terraform
data "sdwan_hub_and_spoke_topology_policy_definition" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the policy definition

### Read-Only

- `description` (String) The description of the policy definition
- `name` (String) The name of the policy definition
- `topologies` (Attributes List) List of topologies (see [below for nested schema](#nestedatt--topologies))
- `type` (String) The policy definition type
- `version` (Number) The version of the policy definition
- `vpn_list_id` (String) VPN list ID
- `vpn_list_version` (Number) VPN list version

<a id="nestedatt--topologies"></a>
### Nested Schema for `topologies`

Read-Only:

- `name` (String) Topology name
- `spokes` (Attributes List) List of spokes (see [below for nested schema](#nestedatt--topologies--spokes))

<a id="nestedatt--topologies--spokes"></a>
### Nested Schema for `topologies.spokes`

Read-Only:

- `hubs` (Attributes List) List of hubs (see [below for nested schema](#nestedatt--topologies--spokes--hubs))
- `site_list_id` (String) Site list ID
- `site_list_version` (Number) Site list version

<a id="nestedatt--topologies--spokes--hubs"></a>
### Nested Schema for `topologies.spokes.hubs`

Read-Only:

- `site_list_id` (String) Site list ID
- `site_list_version` (Number) Site list version
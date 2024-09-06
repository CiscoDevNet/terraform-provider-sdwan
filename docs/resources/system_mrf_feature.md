---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_system_mrf_feature Resource - terraform-provider-sdwan"
subcategory: "Feature"
description: |-
  This resource can manage a System MRF Feature.
  Minimum SD-WAN Manager version: 20.12.0
---

# sdwan_system_mrf_feature (Resource)

This resource can manage a System MRF Feature.
  - Minimum SD-WAN Manager version: `20.12.0`

## Example Usage

```terraform
resource "sdwan_system_mrf_feature" "example" {
  name                    = "Example"
  description             = "My Example"
  feature_profile_id      = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  region_id               = 1
  secondary_region_id     = 2
  role                    = "edge-router"
  enable_migration_to_mrf = "enabled"
  migration_bgp_community = 100
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the Feature

### Optional

- `description` (String) The description of the Feature
- `enable_migration_to_mrf` (String) Enable migration mode to Multi-Region Fabric
  - Choices: `enabled`, `enabled-from-bgp-core`
- `feature_profile_id` (String) Feature Profile ID
- `migration_bgp_community` (Number) Set BGP community during migration from BGP-core based network
  - Range: `1`-`4294967295`
- `region_id` (Number) Set region ID
  - Range: `1`-`63`
- `role` (String) Set the role for router
  - Choices: `edge-router`, `border-router`
- `role_variable` (String) Variable name
- `secondary_region_id` (Number) Set secondary region ID
  - Range: `1`-`63`
- `secondary_region_id_variable` (String) Variable name

### Read-Only

- `id` (String) The id of the Feature
- `version` (Number) The version of the Feature

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_system_mrf_feature.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```
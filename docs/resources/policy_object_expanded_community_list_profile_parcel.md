---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_policy_object_expanded_community_list_profile_parcel Resource - terraform-provider-sdwan"
subcategory: "Profile Parcels"
description: |-
  This resource can manage a Policy Object Expanded Community List profile parcel.
  Minimum SD-WAN Manager version: 20.12.0
---

# sdwan_policy_object_expanded_community_list_profile_parcel (Resource)

This resource can manage a Policy Object Expanded Community List profile parcel.
  - Minimum SD-WAN Manager version: `20.12.0`

## Example Usage

```terraform
resource "sdwan_policy_object_expanded_community_list_profile_parcel" "example" {
  name                     = "Example"
  description              = "My Example"
  feature_profile_id       = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  expanded_community_lists = ["abcd"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `expanded_community_lists` (Set of String) Expanded Community List
- `feature_profile_id` (String) Feature Profile ID
- `name` (String) The name of the profile parcel

### Optional

- `description` (String) The description of the profile parcel
- `expanded_community_lists_variable` (String) Variable name

### Read-Only

- `id` (String) The id of the profile parcel
- `version` (Number) The version of the profile parcel

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_policy_object_expanded_community_list_profile_parcel.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```
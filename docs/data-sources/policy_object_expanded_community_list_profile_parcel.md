---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_policy_object_expanded_community_list_profile_parcel Data Source - terraform-provider-sdwan"
subcategory: "Profile Parcels"
description: |-
  This data source can read the Policy Object Expanded Community List profile parcel.
---

# sdwan_policy_object_expanded_community_list_profile_parcel (Data Source)

This data source can read the Policy Object Expanded Community List profile parcel.

## Example Usage

```terraform
data "sdwan_policy_object_expanded_community_list_profile_parcel" "example" {
  id                 = "f6b2c44c-693c-4763-b010-895aa3d236bd"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `id` (String) The id of the profile parcel

### Read-Only

- `description` (String) The description of the profile parcel
- `expanded_community_lists` (Set of String) Expanded Community List
- `expanded_community_lists_variable` (String) Variable name
- `name` (String) The name of the profile parcel
- `version` (Number) The version of the profile parcel
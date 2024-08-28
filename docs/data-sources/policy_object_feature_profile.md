---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_policy_object_feature_profile Data Source - terraform-provider-sdwan"
subcategory: "Feature Profiles"
description: |-
  This data source can read the Policy Object Feature Profile .
---

# sdwan_policy_object_feature_profile (Data Source)

This data source can read the Policy Object Feature Profile .

## Example Usage

```terraform
data "sdwan_policy_object_feature_profile" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object

### Read-Only

- `description` (String) Description
- `name` (String) The name of the policy object feature profile
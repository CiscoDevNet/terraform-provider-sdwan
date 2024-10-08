---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_extended_community_list_policy_object Data Source - terraform-provider-sdwan"
subcategory: "(Classic) Policy Objects"
description: |-
  This data source can read the Extended Community List Policy Object .
---

# sdwan_extended_community_list_policy_object (Data Source)

This data source can read the Extended Community List Policy Object .

## Example Usage

```terraform
data "sdwan_extended_community_list_policy_object" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object

### Read-Only

- `entries` (Attributes List) List of entries (see [below for nested schema](#nestedatt--entries))
- `name` (String) The name of the policy object
- `version` (Number) The version of the object

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Read-Only:

- `community` (String) Extended community value, e.g. `community soo 10.0.0.1:30` or `community rt 100:10`

---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_ipv6_prefix_list_policy_object Resource - terraform-provider-sdwan"
subcategory: "Policy Objects"
description: |-
  This resource can manage a IPv6 Prefix List policy object.
---

# sdwan_ipv6_prefix_list_policy_object (Resource)

This resource can manage a IPv6 Prefix List policy object.

## Example Usage

```terraform
resource "sdwan_ipv6_prefix_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      prefix = "2001:1:1:2::/64"
      le     = 80
      ge     = 128
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `entries` (Attributes List) List of entries (see [below for nested schema](#nestedatt--entries))
- `name` (String) The name of the policy object

### Read-Only

- `id` (String) The id of the policy object
- `version` (Number) The version of the feature template

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Required:

- `prefix` (String) IP prefix list entry, e.g. `2001:1:1:2::/64`

Optional:

- `ge` (Number) Greater equal
- `le` (Number) Lesser equal

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_ipv6_prefix_list_policy_object.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```
---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_allow_url_list_policy_object Resource - terraform-provider-sdwan"
subcategory: "(Classic) Policy Objects"
description: |-
  This resource can manage a Allow URL List Policy Object .
---

# sdwan_allow_url_list_policy_object (Resource)

This resource can manage a Allow URL List Policy Object .

## Example Usage

```terraform
resource "sdwan_allow_url_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      url = "cisco.com"
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

- `id` (String) The id of the object
- `version` (Number) The version of the object

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Optional:

- `url` (String) URL

## Import

Import is supported using the following syntax:

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
terraform import sdwan_allow_url_list_policy_object.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```

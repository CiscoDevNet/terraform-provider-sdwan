---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_configuration_group Data Source - terraform-provider-sdwan"
subcategory: "Configuration Groups"
description: |-
  This data source can read the Configuration Group .
---

# sdwan_configuration_group (Data Source)

This data source can read the Configuration Group .

## Example Usage

```terraform
data "sdwan_configuration_group" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The id of the object

### Read-Only

- `description` (String) Description
- `devices` (Attributes List) List of devices (see [below for nested schema](#nestedatt--devices))
- `feature_profile_ids` (Set of String) List of feature profile IDs
- `feature_versions` (List of String) List of all associated feature versions
- `name` (String) The name of the configuration group
- `solution` (String) Type of solution
- `topology_devices` (Attributes List) List of topology device types (see [below for nested schema](#nestedatt--topology_devices))
- `topology_site_devices` (Number) Number of devices per site

<a id="nestedatt--devices"></a>
### Nested Schema for `devices`

Read-Only:

- `deploy` (Boolean) Deploy to device if enabled.
- `id` (String) Device ID
- `variables` (Attributes Set) List of variables (see [below for nested schema](#nestedatt--devices--variables))

<a id="nestedatt--devices--variables"></a>
### Nested Schema for `devices.variables`

Read-Only:

- `list_value` (List of String) Use this instead of `value` in case value is of type `List`.
- `name` (String) Variable name
- `value` (String) Variable value



<a id="nestedatt--topology_devices"></a>
### Nested Schema for `topology_devices`

Read-Only:

- `criteria_attribute` (String) Criteria attribute
- `criteria_value` (String) Criteria value
- `unsupported_features` (Attributes List) List of unsupported features (see [below for nested schema](#nestedatt--topology_devices--unsupported_features))

<a id="nestedatt--topology_devices--unsupported_features"></a>
### Nested Schema for `topology_devices.unsupported_features`

Read-Only:

- `parcel_id` (String) Parcel ID
- `parcel_type` (String) Parcel type

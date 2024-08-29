---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_service_object_tracker_profile_parcel Data Source - terraform-provider-sdwan"
subcategory: "Profile Parcels"
description: |-
  This data source can read the Service Object Tracker profile parcel.
---

# sdwan_service_object_tracker_profile_parcel (Data Source)

This data source can read the Service Object Tracker profile parcel.

## Example Usage

```terraform
data "sdwan_service_object_tracker_profile_parcel" "example" {
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
- `interface` (String) interface name
- `interface_variable` (String) Variable name
- `name` (String) The name of the profile parcel
- `object_tracker_id` (Number) Object tracker ID
- `object_tracker_id_variable` (String) Variable name
- `object_tracker_type` (String) objectTrackerType:Interface SIG Route
- `route_ip` (String) IP address
- `route_ip_variable` (String) Variable name
- `route_mask` (String) IP mask
- `route_mask_variable` (String) Variable name
- `version` (Number) The version of the profile parcel
- `vpn` (Number) VPN
- `vpn_variable` (String) Variable name
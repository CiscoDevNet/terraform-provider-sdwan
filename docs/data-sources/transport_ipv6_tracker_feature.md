---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_transport_ipv6_tracker_feature Data Source - terraform-provider-sdwan"
subcategory: "Feature"
description: |-
  This data source can read the Transport IPv6 Tracker Feature.
---

# sdwan_transport_ipv6_tracker_feature (Data Source)

This data source can read the Transport IPv6 Tracker Feature.

## Example Usage

```terraform
data "sdwan_transport_ipv6_tracker_feature" "example" {
  id                 = "f6b2c44c-693c-4763-b010-895aa3d236bd"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `feature_profile_id` (String) Feature Profile ID
- `id` (String) The id of the Feature

### Read-Only

- `description` (String) The description of the Feature
- `endpoint_api_url` (String) API url of endpoint
- `endpoint_api_url_variable` (String) Variable name
- `endpoint_dns_name` (String) Endpoint DNS Name
- `endpoint_dns_name_variable` (String) Variable name
- `endpoint_ip` (String) Endpoint IP
- `endpoint_ip_variable` (String) Variable name
- `endpoint_tracker_type` (String) Endpoint Tracker Type
- `endpoint_tracker_type_variable` (String) Variable name
- `interval` (Number) Interval
- `interval_variable` (String) Variable name
- `multiplier` (Number) Multiplier
- `multiplier_variable` (String) Variable name
- `name` (String) The name of the Feature
- `threshold` (Number) Threshold
- `threshold_variable` (String) Variable name
- `tracker_name` (String) Tracker Name
- `tracker_name_variable` (String) Variable name
- `tracker_type` (String) Tracker Type
- `tracker_type_variable` (String) Variable name
- `version` (Number) The version of the Feature
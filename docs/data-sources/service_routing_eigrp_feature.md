---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_service_routing_eigrp_feature Data Source - terraform-provider-sdwan"
subcategory: "Features - Service"
description: |-
  This data source can read the Service Routing EIGRP Feature.
---

# sdwan_service_routing_eigrp_feature (Data Source)

This data source can read the Service Routing EIGRP Feature.

## Example Usage

```terraform
data "sdwan_service_routing_eigrp_feature" "example" {
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

- `authentication_type` (String) Set EIGRP router authentication type
- `authentication_type_variable` (String) Variable name
- `autonomous_system_id` (Number) Set autonomous system ID <1..65535>
- `autonomous_system_id_variable` (String) Variable name
- `description` (String) The description of the Feature
- `filter` (Boolean) Selective route download
- `filter_variable` (String) Variable name
- `hello_interval` (Number) Set EIGRP hello interval
- `hello_interval_variable` (String) Variable name
- `hmac_authentication_key` (String) Set hmac-sha-256 authentication key
- `hmac_authentication_key_variable` (String) Variable name
- `hold_time` (Number) Set EIGRP hold time
- `hold_time_variable` (String) Variable name
- `interfaces` (Attributes List) Configure IPv4 Static Routes (see [below for nested schema](#nestedatt--interfaces))
- `md5_keys` (Attributes List) Set keychain details (see [below for nested schema](#nestedatt--md5_keys))
- `name` (String) The name of the Feature
- `networks` (Attributes List) Configure the networks for EIGRP to advertise (see [below for nested schema](#nestedatt--networks))
- `redistributes` (Attributes List) Redistribute routes into EIGRP (see [below for nested schema](#nestedatt--redistributes))
- `route_policy_id` (String)
- `version` (Number) The version of the Feature

<a id="nestedatt--interfaces"></a>
### Nested Schema for `interfaces`

Read-Only:

- `name` (String) Set interface name
- `name_variable` (String) Variable name
- `shutdown` (Boolean) Enable/disable EIGRP
- `shutdown_variable` (String) Variable name
- `summary_addresses` (Attributes List) Set summary addresses (see [below for nested schema](#nestedatt--interfaces--summary_addresses))

<a id="nestedatt--interfaces--summary_addresses"></a>
### Nested Schema for `interfaces.summary_addresses`

Read-Only:

- `address` (String)
- `address_variable` (String) Variable name
- `mask` (String)
- `mask_variable` (String) Variable name



<a id="nestedatt--md5_keys"></a>
### Nested Schema for `md5_keys`

Read-Only:

- `key_id` (Number) Set MD5 key ID
- `key_id_variable` (String) Variable name
- `key_string` (String) Set MD5 key
- `key_string_variable` (String) Variable name


<a id="nestedatt--networks"></a>
### Nested Schema for `networks`

Read-Only:

- `ip_address` (String)
- `ip_address_variable` (String) Variable name
- `mask` (String)
- `mask_variable` (String) Variable name


<a id="nestedatt--redistributes"></a>
### Nested Schema for `redistributes`

Read-Only:

- `protocol` (String) Set the protocol to redistribute routes from
- `protocol_variable` (String) Variable name
- `route_policy_id` (String)

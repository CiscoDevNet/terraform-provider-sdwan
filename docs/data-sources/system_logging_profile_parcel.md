---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_system_logging_profile_parcel Data Source - terraform-provider-sdwan"
subcategory: "Profile Parcels"
description: |-
  This data source can read the System Logging profile parcel.
---

# sdwan_system_logging_profile_parcel (Data Source)

This data source can read the System Logging profile parcel.

## Example Usage

```terraform
data "sdwan_system_logging_profile_parcel" "example" {
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
- `disk_enable` (Boolean) Enable logging to local disk
- `disk_enable_variable` (String) Variable name
- `disk_file_rotate` (Number) Set number of syslog files to create before discarding oldest files
- `disk_file_rotate_variable` (String) Variable name
- `disk_file_size` (Number) Set maximum size of file before it is rotated
- `disk_file_size_variable` (String) Variable name
- `ipv4_servers` (Attributes List) Enable logging to remote server (see [below for nested schema](#nestedatt--ipv4_servers))
- `ipv6_servers` (Attributes List) Enable logging to remote ipv6 server (see [below for nested schema](#nestedatt--ipv6_servers))
- `name` (String) The name of the profile parcel
- `tls_profiles` (Attributes List) Configure a TLS profile (see [below for nested schema](#nestedatt--tls_profiles))
- `version` (Number) The version of the profile parcel

<a id="nestedatt--ipv4_servers"></a>
### Nested Schema for `ipv4_servers`

Read-Only:

- `hostname_ip` (String) Set hostname or IPv4 address of server
- `hostname_ip_variable` (String) Variable name
- `priority` (String) Set logging level for messages logged to server
- `priority_variable` (String) Variable name
- `source_interface` (String) Set interface to use to reach syslog server
- `source_interface_variable` (String) Variable name
- `tls_enable` (Boolean) Enable TLS Profile
- `tls_enable_variable` (String) Variable name
- `tls_properties_custom_profile` (Boolean) Define custom profile
- `tls_properties_custom_profile_variable` (String) Variable name
- `tls_properties_profile` (String) Configure a TLS profile
- `tls_properties_profile_variable` (String) Variable name
- `vpn` (Number) Set hostname or IPv4 address of server
- `vpn_variable` (String) Variable name


<a id="nestedatt--ipv6_servers"></a>
### Nested Schema for `ipv6_servers`

Read-Only:

- `hostname_ip` (String) Set IPv6 hostname or IPv6 address of server
- `hostname_ip_variable` (String) Variable name
- `priority` (String) Set logging level for messages logged to server
- `priority_variable` (String) Variable name
- `source_interface` (String) Set interface to use to reach syslog server
- `source_interface_variable` (String) Variable name
- `tls_enable` (Boolean) Enable TLS Profile
- `tls_enable_variable` (String) Variable name
- `tls_properties_custom_profile` (Boolean) Define custom profile
- `tls_properties_custom_profile_variable` (String) Variable name
- `tls_properties_profile` (String) Configure a TLS profile
- `tls_properties_profile_variable` (String) Variable name
- `vpn` (Number) Set hostname or IPv4 address of server
- `vpn_variable` (String) Variable name


<a id="nestedatt--tls_profiles"></a>
### Nested Schema for `tls_profiles`

Read-Only:

- `cipher_suites` (Set of String) Syslog secure server ciphersuites
- `cipher_suites_variable` (String) Variable name
- `profile` (String) Specify the name of the TLS profile
- `profile_variable` (String) Variable name
- `tls_version` (String) TLS Version
- `tls_version_variable` (String) Variable name
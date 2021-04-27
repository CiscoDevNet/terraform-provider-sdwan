---
layout: "dcnm"
page_title: "DCNM: dcnm_vrf"
sidebar_current: "docs-dcnm-resource-vrf"
description: |-
  Manages DCNM VRF
---

# dcnm_vrf #
Manages DCNM VRF

## Example Usage ##

```hcl

resource "dcnm_vrf" "first" {
  fabric_name             = "fab2"
  name                    = "check"
  vlan_id                 = 2002
  segment_id              = "50016"
  vlan_name               = "check"
  description             = "vrf creation"
  intf_description        = "vrf"
  tag                     = "1250"
  max_bgp_path            = 2
  max_ibgp_path           = 4
  trm_enable              = false
  rp_external_flag        = true
  rp_address              = "1.2.3.4"
  loopback_id             = 15
  mutlicast_address       = "10.0.0.2"
  mutlicast_group         = "224.0.0.1/4"
  ipv6_link_local_flag    = "true"
  trm_bgw_msite_flag      = true
  advertise_host_route    = true
  advertise_default_route = "true"
  static_default_route    = false
  deploy                  = true
  attachments {
    serial_number = "9EQ00OGQYV6"
    vlan_id       = 2300
    attach        = false
    loopback_id   = 70
    loopback_ipv4 = "1.2.3.4"
  }
}

```


## Argument Reference ##

* `name` - (Required) name of Object VRF.
* `fabric_name` - (Required) fabric name under which VRF should be created.
* `segment_id` - (Optional) VRF-Segment id. This field is auto-calculated if not provided. However while creating multiple VRFs in the same plan use this field to reserve the VRF id to avoid any conflicts due to concurrent execution. 

<strong>Note: </strong> For auto-generation of segment-id while creating multiple VRFs in the same plan, Use the depends on functionality of terraform to avoid any segment-id conflicts.

* `vlan` - (Optional) vlan Id for the VRF.
* `vlan_name` - (Optional) vlan name for the VRF.
* `description` - (Optional) description for the VRF.
* `intf_description` - (Optional) intf desscription for the VRF.
* `tag` - (Optional) tag for the VRF. Ranging from 0 to 4294967295.
* `max_bgp_path` - (Optional) maximum BGP path value for the VRF. Ranging from 1 to 64.
* `max_ibgp_path` - (Optional) maximum iBGP path value for the VRF. Ranging from 1 to 64.
* `trm_enable` - (Optional) trm enable flag for the VRF. Allowed values are "true" and "false".
* `rp_external_flag` - (Optional) rp external flag for the VRF. Allowed values are "true" and "false".
* `rp_address` - (Optional) rp address for the VRF.
* `loopback_id` - (Optional) loopback ip address for the VRF. Ranging from 0 to 1023.
* `mutlicast_group` - (Optional) multicast group address for the VRF. Ranging from 224.0.0.0/4 to 239.255.255.255/4.
* `mutlicast_address` - (Optional) multicast address for the VRF.
* `ipv6_link_local_flag` - (Optional) ipv6 link local enable flag for the VRF. Allowed values are "true" and "false".
* `trm_bgw_msite_flag` - (Optional) trm bgw multisite enable flag for the VRF. Allowed values are "true" and "false".
* `advertise_host_route` - (Optional) advertise host route enable flag for the VRF. Allowed values are "true" and "false".
* `advertise_default_route` - (Optional) advertise default route enable flag for the VRF. Allowed values are "true" and "false".
* `static_default_route` - (Optional) configure static default route enable flag for the VRF. Allowed values are "true" and "false".
* `template` - (Optional) template name for the VRF. Values allowed "Default_VRF_Universal". Default is "Default_VRF_Universal".
* `mtu` - (Optional) mtu value for the VRF. Ranging from 68 to 9216.
* `extension_template` - (Optional) extension Template name for the VRF. Values allowed are "Default_VRF_Extension_Universal". Default is "Default_VRF_Extension_Universal".
* `service_template` - (Optional) service template name for the VRF.
* `source` - (Optional) source for the VRF.

* `deploy` - (Optional) deploy flag, used to deploy the VRF. Default value is "true".
* `deploy_timeout` - (Optional) deployment timeout, used as the limiter for the deployment status check for VRF resource. It is in the unit of seconds and default value is "300".

* `attachments` - (Optional) attachment Block, have information regarding the switches which should be attached or detached to/from VRF. If `deploy` is "true", then atleast one attachment must be configured.
* `attachments.serial_number` - (Required) serial number of the switch.
* `attachments.vlan_id` - (Optional) vlan ID for the switch associated with VRF. If not mentioned then VRF's default vlan id will be used for attachment.
* `attachments.attach` - (Optional) attach flag for switch. Default value is "true".
* `attachments.free_form_config` - (Optional) free form configuration for the switch attachment.
* `attachments.extension_values` - (Optional) extension values for switch attachment.
* `attachments.loopback_id` - (Optional) loopback id for the switch attachment.
* `attachments.loopback_ipv4` - (Optional) loopback ipv4 address for the switch attachment.
* `attachments.loopback_ipv6` - (Optional) loopback ipv6 address for the switch attachment. 


## Attribute Reference

The only attribute that this resource exports is the `id`, which is set to the
Dn of the VRF.

## Importing ##

An existing VRF can be [imported][docs-import] into this resource via its fabric and name, using the following command:
[docs-import]: https://www.terraform.io/docs/import/index.html


```
terraform import dcnm_vrf.example <fabric_name>:<vrf_name>
```
---
layout: "dcnm"
page_title: "DCNM: dcnm_network"
sidebar_current: "docs-dcnm-resource-network"
description: |-
  Manages DCNM Network
---

# dcnm_network #
Manages DCNM Network

## Example Usage ##

```hcl

resource "dcnm_network" "first" {
  fabric_name     = "fab2"
  name            = "first"
  network_id      = "1234"
  display_name    = "net1"
  description     = "first network from terraform"
  vrf_name        = "VRF1012"
  vlan_id         = 2300
  vlan_name       = "vlan1"
  ipv4_gateway    = "192.0.3.1/24"
  ipv6_gateway    = "2001:db8::1/64"
  mtu             = 1500
  secondary_gw_1  = "192.0.3.1/24"
  secondary_gw_2  = "192.0.3.1/24"
  arp_supp_flag   = true
  ir_enable_flag  = false
  mcast_group     = "239.1.2.2"
  dhcp_1          = "1.2.3.4"
  dhcp_2          = "1.2.3.5"
  dhcp_vrf        = "VRF1012"
  loopback_id     = 100
  tag             = "1400"
  rt_both_flag    = true
  trm_enable_flag = true
  l3_gateway_flag = true

  deploy = true
  attachments {
    serial_number = "9EQ00OGQYV6"
    vlan_id       = 2400
    attach        = true
    switch_ports = ["Ethernet1/4", "Ethernet1/3"]
  }
  attachments {
    serial_number = "9AYOFL6LTML"
    vlan_id       = 2500
    attach        = true
    switch_ports = ["Ethernet1/4"]
  }
}

```


## Argument Reference ##

* `name` - (Required) name of network object.
* `fabric_name` - (Required) fabric name under which network should be created.
* `network_id` - (Optional) Network id to be associated with the network to be created. Pass this value while creating multiple networks in single plan to avoid conflict of autogenerating ids. Id will be computed by DCNM if not provided.

<strong>Note: </strong> For auto-generation of network-id while creating multiple networks in the same plan, use the depends on functionality of terraform to avoid any network-id conflicts.  

* `display_name` - (Optional) display name for the network object. If not mentioned, then `name` will be considered as `display_name`.
* `description` - (Optional) description for the network.
* `vrf_name` - (Optional) name of the vrf which should be associated with the network. If not given then will be configured as "NA" with `l2_only_flag` as "true".
* `vlan_id` - (Optional) vlan number for the network.
* `vlan_name` - (Optional) vlan name for the network.
* `ipv4_gateway` - (Optional) ipv4 address of gateway for the network.
* `ipv6_gateway` - (Optional) ipv6 address of gateway for the network.
* `mtu` - (Optional) mtu value for the network. Ranging from 68 to 9216.
* `tag` - (Optional) tag for the Network. Ranging from 0 to 4294967295.
* `secondary_gw_1` - (Optional) ipv4 secondary gateway 1 for the network.
* `secondary_gw_2` - (Optional) ipv4 secondary gateway 2 for the network.
* `arp_supp_flag` - (Optional) arp suppression flag for the network.
* `ir_enable_flag` - (Optional) ingress replication flag for the network.
* `mcast_group` - (Optional) multicast group address for the network.
* `dhcp_1` - (Optional) ipv4 address of DHCP server 1 for the network.
* `dhcp_2` - (Optional) ipv4 address of DHCP server 2 for the network.
* `dhcp_vrf` - (Optional) vrf name of DHCP server for the network.
* `loopback_id` - (Optional) loopback id for the network. Ranging from 0 to 1023.
* `rt_both_flag` - (Optional) l2 VNI route-target both enable flag for the network.
* `trm_enable_flag` - (Optional) TRM enable flag for the network.
* `l3_gateway_flag` - (Optional) enable L3 gateway on border flag for the network. 
* `template` - (Optional) template name for the network. Values allowed "Default_VRF_Universal" and "Service_Network_Universal". Default is "Default_VRF_Universal".
* `extension_template` - (Optional) extension Template name for the network. Values allowed are "Default_Network_Extension_Universal". Default is "Default_Network_Extension_Universal".
* `service_template` - (Optional) service template name for the network.
* `source` - (Optional) source for the network.

* `deploy` - (Optional) deploy flag, used to deploy the network. Default value is "true".
* `deploy_timeout` - (Optional) deployment timeout, used as the limiter for the deployment status check for network resource. It is in the unit of seconds and default value is "300".

* `attachments` - (Optional) attachment block, have information regarding the switches which should be attached or detached to/from network. If `deploy` is "true", then atleast one attachment must be configured.
* `attachments.serial_number` - (Required) serial number of the switch.
* `attachments.vlan_id` - (Optional) vlan ID for the switch associated with network. If not mentioned then network's default vlan id will be used for attachment.
* `attachments.attach` - (Optional) attach flag for switch. Default value is "true".
* `attachments.dot1_qvlan` - (Optional) dot1 qvlan for switch attachment.
* `attachments.switch_ports` - (Optional) list of port name(i.e. interface names) for switch attachment.
* `attachments.untagged` - (Optional) untagged flag for switch attachment.
* `attachments.free_form_config` - (Optional) free form configuration for the switch attachment.
* `attachments.extension_values` - (Optional) extension values for switch attachment.
* `attachments.instance_values` - (Optional) instance values for switch attachment.

## Attribute Reference

* `id` - Dn for the network.
* `l2_only_flag` - Layer 2 only flag. If VRF is not set then `l2_only_flag` will be set to true.

## Importing ##

An existing network can be [imported][docs-import] into this resource via its fabric and name, using the following command:
[docs-import]: https://www.terraform.io/docs/import/index.html


```
terraform import dcnm_network.example <fabric_name>:<network_name>
```

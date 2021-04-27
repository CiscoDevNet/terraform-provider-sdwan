---
layout: "dcnm"
page_title: "DCNM: dcnm_interface"
sidebar_current: "docs-dcnm-resource-interface"
description: |-
  Manages DCNM interface modules
---

# dcnm_interface #
Manages DCNM interface modules

## Example Usage ##

Port Channel interface:

```hcl
resource "dcnm_interface" "second" {
  policy        = "int_port_channel_access_host_11_1"
  type          = "port-channel"
  name          = "port-channel502"
  fabric_name   = "fab2"
  switch_name_1 = "leaf2"

  mode            = "active"
  bpdu_gaurd_flag = "true"
  mtu             = "jumbo"
  allowed_vlans   = "none"
  access_vlans    = "10"
  pc_interface    = ["e1/6", "eth1/9"]
}
```

Loopback interface:

```hcl
resource "dcnm_interface" "second" {
  fabric_name = "fab2"
  name        = "loopback5"
  type        = "loopback"
  policy      = "int_loopback_11_1"

  switch_name_1             = "leaf1"
  ipv4                      = "1.2.3.4"
  loopback_tag              = "1234"
  vrf                       = "MyVRF"
  loopback_ls_routing       = "ospf"
  loopback_routing_tag      = "1234"
  loopback_router_id        = "10"
  loopback_replication_mode = "Multicast"
  description               = "creation from terraform"
  ipv6                      = "2001::0"
}
```

vPC interface:

```hcl
resource "dcnm_interface" "second" {
  policy        = "int_vpc_trunk_host_11_1"
  type          = "vpc"
  name          = "vPC1"
  fabric_name   = "fab2"
  switch_name_1 = "leaf1"

  switch_name_2           = "leaf2"
  vpc_peer1_id            = "501"
  vpc_peer2_id            = "502"
  mode                    = "active"
  bpdu_gaurd_flag         = "true"
  mtu                     = "jumbo"
  vpc_peer1_allowed_vlans = "none"
  vpc_peer2_allowed_vlans = "none"
  vpc_peer1_access_vlans  = "10"
  vpc_peer2_access_vlans  = "20"
  vpc_peer1_interface     = ["e1/5", "eth1/7"]
  vpc_peer2_interface     = ["e1/5", "eth1/7"]
}
```

Sub-interface:

```hcl
resource "dcnm_interface" "second" {
  policy        = "int_subif_11_1"
  type          = "sub-interface"
  name          = "Ethernet1/41.8"
  fabric_name   = "fab2"
  switch_name_1 = "leaf1"

  vrf               = "MyVRF"
  subinterface_vlan = "8"
  ipv4              = "1.2.3.4"
  ipv6              = "2001::0"
  ipv4_prefix       = "24"
  ipv6_prefix       = "65"
  subinterface_mtu  = "9216"
  description       = "creation from terraform"

  deploy = false
}
```

## Common Argument Reference ##

* `fabric_name` - (Required) fabric name under which interface should be created.
* `name` - (Required) name of the interface. It must be in proper format for example, for loopback: "loopback5", for port-channel "port-channel5", for virtual port channel "vPC17", for sub-interface "Ethernet1/41.8" and for ethernet "Ethernet1/4".
* `type` - (Required) type of the interface. Allowed values are "loopback", "port-channel", "vpc", "sub-interface", "ethernet".
* `policy` - (Required) policy name for the interface.
* `switch_name_1` - (Required) name of the switch which should be associated to the interface.
* `admin_state` - (Optional) administrative state for the interface. Allowed values are "true" and "false". Default value is "true".
* `deploy` - (Optional) deploy flag for the deployment of interface. Allowed values are "true" and "false". Default value is "true".

## Argument Reference for loopback Interface ##

* `vrf` - (Optional) vrf name for the loopback interface.
* `ipv4` - (Optional) ipv4 address for the loopback interface.
* `ipv6` - (Optional) ipv6 address for the loopback interface.
* `loopback_tag` - (Optional) tag for the loopback interface.
* `loopback_routing_tag` - (Optional) routing tag for the loopback interface.
* `loopback_ls_routing` - (Optional) link state routing protocol for the loopback interface.
* `loopback_router_id` - (Optional) router id for the loopback interface.
* `loopback_replication_mode` - (Optional) replication mode for the loopback interface.
* `configuration` - (Optional) configuration for the loopback interface.
* `description` - (Optional) description for the loopback interface.

## Argument Reference for port-channel Interface ##

* `pc_interface` - (Optional) list of port channel member interface for port-channel interface.
* `access_vlans` - (Optional) access vlans for the port-channel interface.
* `mode` - (Optional) mode for the port-channel interface. Allowed values are "on", "active" and "passive".
* `bpdu_gaurd_flag` - (Optional) BPDU flag for the port-channel interface. Allowed values are "true", "false" and "no".
* `port_fast_flag` - (Optional) port type fast flag for the port-channel interface.
* `mtu` - (Optional) mtu for the port-channel interface. Allowed values are "jumbo" and "default". 
* `allowed_vlans` - (Optional) allowed vlans for the port-channel interface. Allowed values are "none", "all" or vlan ranges(1-200,500-2000,3000) 
* `configuration` - (Optional) configuration for the port-channel interface.
* `description` - (Optional) description for the port-channel interface.

## Argument Reference for vPC Interface ##

* `switch_name_2` - (Optional) name of the second switch with which interface should be associated. 
* `vpc_peer1_id` - (Optional) peer1 port-channel id for the vPC interface.
* `vpc_peer2_id` - (Optional) peer2 port-channel id for the vPC interface.
* `vpc_peer1_interface` - (Optional) list of peer1 member interface for the vPC interface.
* `vpc_peer2_interface` - (Optional) list of peer2 member interface for the vPC interface.
* `mode` - (Optional)  mode for the vPC interface. Allowed values are "on", "active" and "passive".
* `bpdu_gaurd_flag` - (Optional) BPDU flag for the vPC interface. Allowed values are "true", "false" and "no".
* `port_fast_flag` - (Optional) port type fast flag for the vPC interface.
* `mtu` - (Optional) mtu for the vPC interface. Allowed values are "jumbo" and "default".
* `vpc_peer1_allowed_vlans` - (Optional) peer1 allowed vlans for the vPC interface. Allowed values are "none", "all" or vlan ranges(1-200,500-2000,3000) 
* `vpc_peer2_allowed_vlans` - (Optional) peer2 allowed vlans for the vPC interface. Allowed values are "none", "all" or vlan ranges(1-200,500-2000,3000) 
* `vpc_peer1_access_vlans` - (Optional) peer1 access vlans for the vPC interface.
* `vpc_peer2_access_vlans` - (Optional) peer2 access vlans for the vPC interface.
* `vpc_peer1_desc` - (Optional) peer1 description for the vPC interface.
* `vpc_peer2_desc` - (Optional) peer2 description for the vPC interface.
* `vpc_peer1_conf` - (Optional) peer1 configuration for the vPC interface.
* `vpc_peer2_conf` - (Optional) peer2 configuration for the vPC interface.

## Argument Reference for sub-interface Interface ##

* `subinterface_vlan` - (Optional) vlan for the sub-interface.
* `vrf` - (Optional) vrf for the sub-interface.
* `ipv4` - (Optional) ipv4 address for the sub-interface.
* `ipv6` - (Optional) ipv6 address for the sub-interface.
* `ipv6_prefix` - (Optional) ipv6 prefic for the sub-interface.
* `ipv4_prefix` - (Optional) ipv4 prefix for the sub-interface.
* `subinterface_mtu` - (Optional) mtu for the sub-interface.
* `configuration` - (Optional) configuration for the sub-interface.
* `description` - (Optional) description for the sub-interface.

## Argument Reference for ethernet Interface ##

* `bpdu_gaurd_flag` - (Optional) BPDU flag for the ethernet interface. Allowed values are "true", "false" and "no".
* `port_fast_flag` - (Optional) port type fast flag for the ethernet interface.
* `mtu` - (Optional) mtu for the ethernet interface. Allowed values are "jumbo" and "default". If `policy` is configured as "epl_routed_intf" or "int_routed_host_11_1", then allowed value range is from 576 to 9216.
* `ethernet_speed` - (Optional) speed of the ethernet. Allowed values are "Auto", "100Mb", "1Gb", "10Gb", "25Gb",	"40Gb" and "100Gb".
* `allowed_vlans` - (Optional) allowed vlans for the ethernet interface. Allowed values are "none", "all" or vlan ranges(1-200,500-2000,3000)
* `configuration` - (Optional) configuration for the ethernet.
* `description` - (Optional) description for the ethernet.
* `ipv4` - (Optional) ipv4 address for the ethernet.
* `ipv6` - (Optional) ipv6 address for the ethernet.
* `ipv6_prefix` - (Optional) ipv6 prefix for the ethernet.
* `ipv4_prefix` - (Optional) ipv4 prefix for the ethernet.
* `access_vlans` - (Optional) access vlans for the ethernet interface.


## Attribute Reference

* `serial_number` - Dn for the interface module.

## Importing ##

An existing interface can be [imported][docs-import] into this resource via its serial number, type and name, using the following command:
[docs-import]: https://www.terraform.io/docs/import/index.html


```
terraform import dcnm_interface.example <type>:<serial_number>:<name>:<fabric_name>
```
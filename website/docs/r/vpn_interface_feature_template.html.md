---
layout: "sdwan"
page_title: "SD-WAN: sdwan_vpn_interface_feature_template"
sidebar_current: "docs-sdwan-resource-vpn_interface_feature_template"
description: |-
  Resource for SD-WAN VPN interface Feature Template
---

# sdwan_vpn_interface_feature_template #
Resource for SD-WAN VPN Interface Feature Templates

## Example Usage ##

```hcl

# example for Cisco VPN Interface Feature Template

resource "sdwan_vpn_interface_feature_template" "name1" {
  template_name = "Sample_VPN_Interface_1"
  template_description = "For testing purposes"
  device_type = ["vedge-C1161-8P"]
  template_type = "cisco_vpn_interface"
  template_min_version = "15.0.0"
  factory_default = false

  template_definition {  
     vpn_interface_basic {
        shutdown= true
       description = "for testing purpose"
     }
  }
}

# example for VPN Vedge Interface Feature Template

resource "sdwan_vpn_interface_feature_template" "name2" {
  template_name = "Sample_VPN_Interface_2"
  template_description = "For testing purposes"
  device_type = ["vedge-100"]
  template_type = "vpn-vedge-interface"
  template_min_version = "15.0.0"
  factory_default = false
  template_definition {     
     vpn_interface_basic {
       shutdown= true
       description = "for testing purpose"
       ipv4 {
         primary_address = "192.168.10.2/16"
         secondary_address {
           address = "192.168.11.2/16"           
         }
         secondary_address {
           address = "192.168.10.3/16"           
         }
         secondary_address {
           address = "192.168.10.4/16"           
         }
         secondary_address {
           address = "192.168.10.5/16"           
         }
         dhcp_helper = ["192.168.10.5"]
         dhcp_distance = 1
       }

       ipv6 {
          primary_address = "2001:0db8:85a3:0000:0000:8a2e:0370:7334/12"
       
          secondary_address {
           address = "2001:0db8:85a3:0000:0000:8a2e:0370:7335/12"           
          }
          secondary_address {
           address = "2001:0db8:85a3:0000:0000:8a2e:0370:7336/12"           
          }
          dhcp_distance = 1
          dhcp_rapid_commit = true
          dhcp_helper {
            address = "2001:0db8:85a3:0000:0000:8a2e:0370:7336"
            vpn = 1
          }
          dhcp_helper {
            address = "2001:0db8:85a3:0000:0000:8a2e:0370:7337"
            vpn = 1
          }
          dhcp_helper {
            address = "2001:0db8:85a3:0000:0000:8a2e:0370:7338"
            vpn = 1
          }
          dhcp_helper {
            address = "2001:0db8:85a3:0000:0000:8a2e:0370:7339"
            vpn = 2
          }
          dhcp_helper {
            address = "2001:0db8:85a3:0000:0000:8a2e:0370:7341"
            vpn = 1
          }
          dhcp_helper {
            address = "2001:0db8:85a3:0000:0000:8a2e:0370:7342"
            vpn = 1
          }
          dhcp_helper {
            address = "2001:0db8:85a3:0000:0000:8a2e:0370:7343"
            vpn = 1
          }
          dhcp_helper {
            address = "2001:0db8:85a3:0000:0000:8a2e:0370:7344"
            vpn = 1
          }
       }       

       block_non_source_ip = true
       bandwidth_upstream = 123
       bandwidth_downstream = 1333
     }

     vpn_interface_tunnel {
       per_tunnel_qos = true
       per_tunnel_qos_aggregator = true
       tunnels_bandwidth_percent = 90
       color = "red"
       restrict = true
       groups = [1,2333,42949]
       border = true
       control_connection = true
       maximum_control_connections = 5
       vbond_as_stun_server = true
       exclude_controller_group_list = [1,2,3]
       vmanage_connection_preference = 5
       port_hop = true
       low_bandwidth_link = true
       allow_service {
         all = true
         bgp = true
         dhcp = true
         dns = true
         icmp = true
         netconf = true
         ntp = true
         ospf = true
         ssh = true
         stun = true
         https = true
         snmp = true
       }
       encapsulation {
         gre = true
         gre_preference = 101
         gre_weight = 5
         ipsec  = true
         ipsec_preference = 100
         ipsec_weight = 6
         carrier = "carrier3"
         bind_loopback_tunnel = "test"
         last_resort_circuit = true
         hold_time = 101
         nat_refresh_interval = 51
         hello_interval = 200
         hello_tolerance = 20
         gre_tunnel_destination_ip = "198.162.10.11"
       }
     }

    vpn_interface_acl_qos {
      adapt_period = 100
      shaping_rate_upstream_min = 111
      shaping_rate_upstream_max = 222
      shaping_rate_upstream_default = 150
      shaping_rate_downstream_min = 333
      shaping_rate_downstream_max = 444
      shaping_rate_downstream_default = 350
      shaping_rate = 123
      qos_map = "test"
      rewrite_rule = "124abc"
      ipv4_ingress_access_list = "abc1"
      ipv4_egress_access_list = "abc2"
      ipv6_ingress_access_list = "xyz1"
      ipv6_egress_access_list = "xyz2"
      ingress_policer_name = "abc4"
      egress_policer_name = "abc5"
    }
     vpn_interface_nat {
       ipv4 {
         nat_type = "loopback"
         refresh_mode = "outbound"
         log_nat_flow_creations_or_deletions = false
         udp_timeout = 50
         tcp_timeout = 100
         block_icmp = false
         respond_to_ping = true
         nat_pool_range_start = "192.162.15.10"
         nat_pool_range_end = "192.162.15.20"
         nat_pool_prefix_length = 100
         overload = true
         nat_inside_source_loopback_interface = "asvcd"

         port_forward {
           port_start_range = 30
           port_end_range = 35
           protocol = "tcp"
           vpn = 0
           private_ip = "192.162.15.20"
         }
          port_forward {
           port_start_range = 40
           port_end_range = 45
           protocol = "udp"
           vpn = 1
           private_ip = "192.162.15.21"
         }

         static_nat {
           source_ip = "192.168.20.10"
           translated_source_ip = "192.168.20.11"
           source_vpn_id = 1
           static_nat_direction = "inside"
           protocol = "tcp"
           source_port = 20
           translate_port = 25
         }
         static_nat {
           source_ip = "192.168.20.11"
           translated_source_ip = "192.168.20.12"
           source_vpn_id = 2
           static_nat_direction = "outside"
           protocol = "udp"
           source_port = 30
           translate_port = 35
         }
       }
       ipv6 {
        nat64 = true
       }
     }

     vpn_interface_vrrp {
        ipv4 {
         group_id = 1
         priority = 102
         timer = 9
         track_omp = false
         track_prefix_list = "test"
         ip_address = "198.165.10.20"
        }
        ipv6 {
         group_id = 11
         priority = 105
         timer = 900
         track_omp = false
         track_prefix_list = "test"
         link_local_ipv6_address = "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
         global_ipv6_prefix = "2001:0db8:85a3:0000:0000:8a2e:0370:7334/12"
       }
     }

    vpn_interface_arp {
      ip_address = "198.168.10.30"
      mac_address = "00:1B:44:11:3A:B7"
    }  
    vpn_interface_arp {
      ip_address = "198.168.10.35"
      mac_address = "00:1B:44:11:3A:B8"
    }  

    vpn_interface_advanced {
      duplex = "half"
      mac_address = "00:1B:44:11:3A:B7"
      ip_mtu = 666
      pmtu_discovery = true
      flow_control = "both"
      tcp_mss = 555
      speed = "1000"
      clear_dont_fragment = true
      static_ingess_qos = 5
      arp_timeout = 1111
      autonegotiation = true
      tloc_extension = "test"
      power_over_ethernet = false
      load_interval = 60
      tracker = ["abc"]
      icmp_redirect_disable = true
      gre_tunnel_source_ip = "192.168.11.11"
      xconnect = "22"
      ip_directed_broadcast = true
    }

    vpn_interface_trustsec {
      enable_sgt_propagation = true
      propagate = true
      security_group_tag = 20
      trusted = true
    }

    vpn_interface_8021x {
      radius_server = ["198", "test"]
      account_interim_interval = 12
      nas_identifier = 2
      nas_ip = "192.168.2.10"
      wake_on_lan = true
      control_direction = "in-only"
      reauthentication = 2
      inactivity = 23
      host_mode = "multi-host"
      advanced_options {
        vlan {
          authentication_fail_vlan = 30
          guest_vlan = 40
          authentication_reject_vlan = 45
          default_vlan = 50
        }
        dynamic_authentication_server {
          das_port = 55
          client = "192.168.15.2"
          secret_key = "key"
          time_window = 60
          required_timestamp = true
          vpn = 0
        }
        mac_authentication_bypass {
          server = true
          mac_authentication_bypass_entries = ["00:0a:95:9d:68:16"]
        }

        request_attributes {
          authentication {
            id = 1
            syntax_choice = "string"
            value = "test"           
          }
          authentication {
            id = 2
            syntax_choice = "integer"
            value = "10"
            
          }
          accounting {
            id = 3
            syntax_choice = "octet"
            value = "101"            
          }
        }
      }
    }    
  }
}


```

## Argument Reference ##

* `template_name` - (Required) Unique VPN Interface Type Feature Template Name, Must not include <, >, !, &, ", or white space; maximum 128 characters
* `template_description` - (Required) Long Description of VPN Interface type Feature Template, Should not be empty
* `device_type` - (Required) Type of device which supports Feature Template, Allowed values for Vedge VPN Interface type Feature Template : "vedge-100", "vedge-100-WM", "vedge-ISR1100-4GLTE", "vedge-100-B", "vedge-1000", "vedge-ISR1100-6G", "vedge-100-M", "vedge-2000",	"vedge-ISR1100X-4G", "vedge-ISR1100-4G", "vedge-5000","vedge-cloud", "vedge-ISR1100X-6G"
Allowed values for Cisco VPN Interface type Feature Template : "vedge-C8500-12X4QC", "vedge-ISRv", "vedge-C1111-4PLTEEA", "vedge-C1161-8P", "vedge-C1117-4PLTEEAW",
"vedge-C1121X-8P", "vedge-C8200-1N-4T", "vedge-ISR-4331", "vedge-C1127X-8PMLTEP", "vedge-C1117-4PMLTEEAWE", "vedge-ISR-4451-X", "vedge-C1113-8PLTEEA", "vedge-IR-1821",
"vedge-ASR-1001-X", "vedge-ISR-4321", "vedge-C1116-4PLTEEAWE", "vedge-C1109-4PLTE2P", "vedge-C1121-8P", "vedge-ASR-1002-HX", "vedge-C1111-8PLTEEAW", "vedge-C1112-8PWE",
"vedge-C1101-4PLTEP", "vedge-ISR1100-4GLTENA-XE", "vedge-C1111-8PLTELA", "vedge-IR-1835", "vedge-C1121X-8PLTEP", "vedge-IR-1833", "vedge-C8300-1N1S-4T2X", "vedge-C1121-4P", "vedge-ISR-4351", "vedge-C1117-4PLTELA", "vedge-C1116-4PWE", "vedge-nfvis-C8200-UCPE", "vedge-C1113-8PM", "vedge-IR-1831", "vedge-C1127-8PLTEP",
"vedge-C1121-8PLTEPW", "vedge-C1113-8PW", "vedge-ASR-1001-HX", "vedge-C1128-8PLTEP", "vedge-C1113-8PLTEEAW", "vedge-C1117-4PW", "vedge-C1116-4P", "vedge-C1113-8PMLTEEA",
"vedge-C1112-8P", "vedge-ISR-4461", "vedge-C1116-4PLTEEA", "vedge-ISR-4221", "vedge-C1117-4PM", "vedge-C1113-8PLTELAWZ", "vedge-C1117-4PMWE", "vedge-C1109-2PLTEVZ",
"vedge-C1113-8P", "vedge-C1117-4P", "vedge-nfvis-ENCS5400", "vedge-C8300-2N2S-6T", "vedge-C1127-8PMLTEP", "vedge-ESR-6300", "vedge-ISR-4221X", vedge-ISR1100-4GLTEGB-XE",
"vedge-C8500-12X", "vedge-C1109-2PLTEGB", "vedge-CSR-1000v", "vedge-C1113-8PLTEW", "vedge-C1121X-8PLTEPW", "vedge-ISR1100-6G-XE", "vedge-C1121-4PLTEP", "vedge-C1111-8PLTEEA", "vedge-C1117-4PLTEEA", "vedge-C1127X-8PLTEP", "vedge-C1109-2PLTEUS",  vedge-C1112-8PLTEEAWE", "vedge-C1161X-8P", "vedge-C8500L-8S4X", "vedge-C1111-8PW", "vedge-C1161X-8PLTEP", "vedge-C1101-4PLTEPW", "vedge-ISR1100X-4G-XE", "vedge-IR-1101", "vedge-C1111-4P", "vedge-C1111-4PW", "vedge-C1111-8P",
"vedge-C1117-4PMLTEEA", "vedge-C1113-8PLTELA", "vedge-C1111-8PLTELAW", "vedge-C1161-8PLTEP", "vedge-ISR1100X-6G-XE", "vedge-ISR-4431", "vedge-C1101-4P", "vedge-C1109-4PLTE2PW", "vedge-C1113-8PMWE", "vedge-C1118-8P", "vedge-C1126-8PLTEP", "vedge-C8300-1N1S-6T", "vedge-C1121-8PLTEP", "vedge-C8300-2N2S-4T2X", "vedge-C1112-8PLTEEA", "vedge-C1111-4PLTELA", "vedge-ASR-1002-X", "vedge-C1111X-8P", "vedge-C1126X-8PLTEP", "vedge-ASR-1006-X", "vedge-C8000V", "vedge-ISR1100-4G-XE",
"vedge-C1117-4PLTELAWZ"
* `template_type` - (Required) Template type for the VPN Interface type Feature Template, Allowed Value: "vpn-vedge-interface","cisco_vpn_interface"
* `template_min_version` - (Required) Minimum Version for the Feature template
* `factory_default` - (Required) Boolean flag indicating whether VPN Interface type Feature Template is factory default or not. If we set it true we can not update or delete resource.
* `template_definition` - (Required) Configuration for VPN Interface type Feature Template, (see [below for nested schema](#nestedblock--template_definition))


## Attribute Reference ##

* `attached_masters_count` - Number of Device Templates attached to VPN interface type Feature Template
* `devices_attached` - Number of Devices attached to VPN interface type Feature Template
* `last_updated_by` - User who updated the VPN interface type Feature Template latest
* `last_updated_on` - Time for the latest Update of VPN interface type Feature Template
* `g_template_class` - Template Class type for VPN interface type Feature Template
* `template_id` - Unique ID for VPN interface type Feature Template
* `config_type` - Type of configuration for VPN interface type Feature Template
* `created_on` - Time of creation of VPN interface type Feature Template
* `rid` - rid for the VPN interface type Feature Template
* `feature` - Feature for the VPN interface type Feature Template

<a id="nestedblock--template_definition"></a>

## Nested Schema for template_definition
* `vpn_interface_basic` - (Required) Basic configuration for the  VPN Interface type Feature Template, Max Items: 1, Min Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_basic))
* `vpn_interface_tunnel` - (Optional) Tunnel configuration for the  VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_tunnel))
* `vpn_interface_nat` - (Optional) NAT(Network Address Translation) configuration for the  VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_nat))
* `vpn_interface_vrrp` - (Optional) VRRP(Virtual Router Redundancy Protocol) configuration for the  VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_vrrp))
* `vpn_interface_acl_qos` - (Optional) ACL(Apply Access Lists) and QoS  configuration for the  VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_acl_qos))
* `vpn_interface_arp` - (Optional) ARP(Address Resolution Protocol) configuration for the  VPN Interface type Feature Template, (see [below for nested schema](#nestedblock--vpn_interface_arp))
* `vpn_interface_8021x` - (Optional) IEEE 802.1X Authentication configuration for the Vedge VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_8021x))
* `vpn_interface_trustsec` - (Optional) TrustSec configuration for the Cisco VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_trustsec)) 
* `vpn_interface_advanced` - (Optional) Advanced configuration for the VPN Interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--vpn_interface_advanced))


<a id="nestedblock--vpn_interface_basic"></a>

## Nested Schema for vpn_interface_basic
* `shutdown` - (Optional) Shutdown flag for VPN interface type Feature Template, Default: true
* `description` - (Optional) Interface Description for VPN interface type Feature Template
* `ipv4` - (Optional) IPv4 configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--ipv4))
* `ipv6` - (Optional) IPv6 configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--ipv6))
* `block_non_source_ip` - (Optional) Flag indicating whether forwarded traffic matches with interface's IP prefix range for VPN interface type Feature Template, Default: false
* `bandwidth_upstream` - (Optional) Value of bandwidth above which to generate notifications in case of transmitted traffic for VPN interface type Feature Template, Range: [1, 2147483647]
* `bandwidth_downstream` - (Optional) Value of bandwidth above which to generate notifications in case of received traffic for VPN interface type Feature Template, Range: [1, 2147483647]


<a id="nestedblock--vpn_interface_tunnel"></a>

## Nested Schema for vpn_interface_tunnel
* `per_tunnel_qos` - (Optional) Flag indicating whether per_tunnel_qos is enabled or not for VPN interface type Feature Template, Default: false
* `per_tunnel_qos_aggregator` - (Optional) Flag indicating whether per_tunnel_qos_aggregator is enabled or not for VPN interface type Feature Template, Default: false
* `tunnels_bandwidth_percent` - (Optional) Value of tunnel bandwith percentage for VPN interface type Feature Template, Range: [1, 99]
* `color` - (Optional) TLOC color value for VPN interface type Feature Template, Allowed values: "3g", "biz-internet", "blue", "bronze", "custom1",
"custom2", "custom3", "default", "gold", "green", "Ite", "metro-ethernet",
"mpls", "private1", "private2", "private3", "private4", "private5", "private6", "public-internet", "red", "silver"
* `restrict` - (Optional) Restrict flag for VPN interface type Feature Template, Default: false
* `groups` - (Optional) List of Groups in tunnle for VPN interface type Feature Template, Range: [1, 4294967295]
* `border` - (Optional) Flag indicating whether Border is enabled or not for VPN interface type Feature Template, Default: false
* `control_connection` - (Optional) Flag indicating TLOC connection is enabled or not for VPN interface type Feature Template, Allowed value: true, false
* `maximum_control_connections` - (Optional) Maximum number of vSamart controllers that WAN tunnel interface can connect for VPN interface type Feature Template, Range: [0, 100]
* `vbond_as_stun_server` - (Optional) Flag indicating whether Session Traversal Utilities for NAT(STUN) is enabled or not for VPN interface type Feature Template, Default: false
* `exclude_controller_group_list` - (Optional) List of vSmart controllers that the tunnel interface is not allowed to connect for VPN interface type Feature Template, Range: [0, 100]
* `vmanage_connection_preference` - (Optional) Flag indicating preference for using a tunnel interface to exchange control traffic with the vManage NMS for VPN interface type Feature Template, Range: [0, 8], Default: 5
* `port_hop` - (Optional) Flag indicating whether Port Hop is enabled or not for VPN interface type Feature Template, Default: true
* `low_bandwidth_link` - (Optional) Flag indicating whether tunnle is set as a low-bandwidth link for VPN interface type Feature Template, Default: false
* `allow_service` - (Optional) Nested block for allow or disallow various services for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--allow_service))
* `encapsulation` - (Optional) Encapsulation configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--encapsulation))


<a id="nestedblock--vpn_interface_nat"></a>

## Nested Schema for vpn_interface_nat
* `ipv4` - (Optional) NAT IPv4 configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--natv4))
* `ipv6` - (Optional) NAT IPv6 configuration for VPN interface type Feature Template, Max Items: 1 (see [below for nested schema](#nestedblock--natv6))



<a id="nestedblock--vpn_interface_vrrp"></a>

## Nested Schema for vpn_interface_vrrp
* `ipv4` - (Optional) VRRP IPv4 configuration for VPN interface type Feature Template, Max Items for Cisco VPN Interface type Feature Template: 1, Max Items for Vedge VPN Interface type Feature Template: 5,(see [below for nested schema](#nestedblock--vrrpv4))
* `ipv6` - (Optional) VRRP IPv6 configuration for VPN interface type Feature Template, Max Items for VPN Interface type Feature Template: 1,(see [below for nested schema](#nestedblock--vrrpv6))


<a id="nestedblock--vpn_interface_acl_qos"></a>

## Nested Schema for vpn_interface_acl_qos
* `adapt_period` - (Optional) Value of Adapt Period for Cisco VPN interface type Feature Template, Range: [1, 720]
* `shaping_rate_upstream_min` - (Optional) Minimum value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template, Range: [8, 100000000]
* `shaping_rate_upstream_max` - (Optional) Maximum value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template, Range: [8, 100000000]
* `shaping_rate_upstream_default` - (Optional) Default value of shaping rate in case of transmitted traffic for Cisco VPN interface type Feature Template, Range: [8, 100000000]
* `shaping_rate_downstream_min` - (Optional) Minimum value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template, Range: [8, 100000000]
* `shaping_rate_downstream_max` - (Optional) Maximum value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template, Range: [8, 100000000]
* `shaping_rate_downstream_default` - (Optional) Default value of shaping rate in case of received traffic for Cisco VPN interface type Feature Template, Range: [8, 100000000]
* `shaping_rate` - (Optional) Value of aggregate traffic transmission rate on the VPN interface for VPN interface type Feature Template, Range: [8, 100000000]
* `qos_map` - (Optional) Name of the QoS map to apply to packets being transmitted out the interface for VPN interface type Feature Template
* `rewrite_rule` - (Optional) Name of the rewrite rule to apply on the interface for VPN interface type Feature Template
* `ipv4_ingress_access_list` - (Optional) Name of the access list to apply to IPv4 packets being received on the interface for VPN interface type Feature Template
* `ipv4_egress_access_list` - (Optional) Name of the access list to apply to IPv4 packets being transmitted on the interface for VPN interface type Feature Template
* `ipv6_ingress_access_list` - (Optional) Name of the access list to apply to IPv6 packets being received on the interface for VPN interface type Feature Template
* `ipv6_egress_access_list` - (Optional) Name of the access list to apply to IPv6 packets being transmitted on the interface for VPN interface type Feature Template
* `ingress_policer_name` - (Optional) Name of the policer to apply to packets received on the interface for VPN interface type Feature Template
* `egress_policer_name` - (Optional) Name of the policer to apply to packets being transmitted on the interface for Vedge VPN interface type Feature Template


<a id="nestedblock--vpn_interface_arp"></a>

## Nested Schema for vpn_interface_arp
* `ip_address` - (Required) IP address for the ARP entry in dotted decimal notation or as a fully qualified host name for VPN interface type Feature Template, Allowed values: IPv4 address
* `mac_address` - (Required) MAC address in colon-separated hexadecimal notation for VPN interface type Feature Template, Allowed values: MAC address



<a id="nestedblock--vpn_interface_8021x"></a>

## Nested Schema for vpn_interface_8021x
* `radius_server` - (Optional) List of tags of the RADIUS server to use for 802.1X authentication for Vedge VPN interface type Feature Template
* `account_interim_interval` - (Optional) Value of how often to send 802.1X interim accounting updates to the RADIUS server for Vedge VPN interface type Feature Template, Range: [1, 1440]
* `nas_identifier` - (Optional) The NAS identifier of the local router
* `nas_ip` - (Optional) The NAS IP address of the local router for Vedge VPN interface type Feature Template, Allowed values: IPv4 address
* `wake_on_lan` - (Optional) The flag that enable client to be powered up when the router receives an Ethernet magic packet frame for Vedge VPN interface type Feature Template, Allowed value: true, false  Default: false
* `control_direction` - (Optional) Direction type of packets for Vedge VPN interface type Feature Template, Allowed values: "in-and-out",					"in-only", Default: "in-and-out"
* `reauthentication` - (Optional) Value indicating how often to reauthenticate 802.1X clients for Vedge VPN interface type Feature Template, Range: [0, 1440], Default: 0
* `inactivity` - (Optional) Value indicating how long to wait before revoking an 802.1X client's network access for Vedge VPN interface type Feature Template, Range: [0, 1440], Default: 60
* `host_mode` - (Optional) Value indicating whether an 802.1X interface grants access to a single client or to multiple clients for Vedge VPN interface type Feature Template, Allowed values: "multi-auth", "multi-host",
"single-host", Default: "single-host"
* `advanced_options` - (Optional) Advance configuration of vpn_interface_8021x for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--advanced_options))


<a id="nestedblock--vpn_interface_trustsec"></a>

## Nested Schema for vpn_interface_trustsec
* `enable_sgt_propagation` - (Optional) Flag indicating whether SGT propagation is enabled or not for Cisco VPN interface type Feature Template, Allowed value: true, false  Default: false
* `propagate` - (Optional) Flag indicating whether propagate is enabled or not for Cisco VPN interface type Feature Template, Allowed value: true, false
* `security_group_tag` - (Optional) Value of security group tag for Cisco VPN interface type Feature Template
* `trusted` - (Optional) Flag indicating whether trusted is enable or not for Cisco VPN interface type Feature Template, Allowed value: true, false  Default: false



<a id="nestedblock--vpn_interface_advanced"></a>

## Nested Schema for vpn_interface_advanced
* `duplex` - (Optional) Value of duplex to run interface in auto, full, or half mode for VPN interface type Feature Template, Allowed values for Cisco VPN Interface type feature template: "auto", "full", "half", Allowed values for Vedge  VPN Interface type feature template: "full", "half"
* `mac_address` - (Optional) Value of MAC address to associate with the interface, in colon-separated hexadecimal notation, Allowed values: MAC address
* `ip_mtu` - (Optional) Value of ip mtu to specify the maximum MTU size of packets on the interface, Range: [576, 2000], Default: 1500
* `pmtu_discovery` - (Optional) Flag indicating whether path MTU is enable or not for Vedge VPN interface type Feature Template, Allowed value: true, false  Default: false
* `flow_control` - (Optional) Value of bidirectional flow control settings for Vedge VPN interface type Feature Template, Allowed values: "autoneg",
"both", "egress", "ingress", "none", Default:  "autoneg"
* `tcp_mss` - (Optional) Maximum segment size (MSS) of TPC SYN packets passing through the router, Range for Cisco VPN Interface : [500, 1460], Range for Vedge VPN Interface : [552, 1460]
* `speed` - (Optional) Value of speed to use when the remote end of the connection does not support autonegotiation, Allowed values for Cisco VPN Interface type feature template: "10",	"100", "1000", "10000", "2500",  Allowed values for Vedge  VPN Interface type feature template: "10",	"100", "1000", "10000" 
* `clear_dont_fragment` - (Optional) Flag indicating Don't Fragment (DF) bit in the IPv4 packet header for packets being transmitted out the interface is clar or not, Allowed value: true, false  Default: false
* `static_ingess_qos` - (Optional) Queue number to use for incoming traffic for Vedge VPN interface type Feature Template, Range: [0, 7]
* `arp_timeout` - (Optional) Value indicating how long it takes for a dynamically learned ARP entry to time out, Range for Cisco VPN Interface : [0, 2147483], Range for Vedge VPN Interface : [0, 2678400] Default: 1500
* `autonegotiation` - (Optional) Flag indicating autonegotiation mode, Allowed value: true, false  Default: true
* `tloc_extension` - (Optional) Name of a physical interface on the same router that connects to the WAN transport
* `power_over_ethernet` - (Optional) Flag indicating whether PoE(Power over Ethernet) is enabled or not for Vedge VPN interface type Feature Template, Allowed value: true, false  Default: false
* `load_interval` - (Optional) Value of load interval for Cisco VPN interface type Feature Template, Range: [30, 600], Allowed values: value can only be multiple of 30
* `tracker` - (Optional) List of a tracker name to track the status of transport interfaces that connect to the internet for VPN interface type Feature Template
* `icmp_redirect_disable` - (Optional) Flag indicating whether ICMP redirect disable or not on interface for VPN interface type Feature Template, Allowed value: true, false  Default: false
* `gre_tunnel_source_ip` - (Optional) IP address of the extended WAN interface, Allowed values: IPv4 address
* `xconnect` - (Optional) The name of a physical interface on the same router that connects to the WAN transport for VPN interface type Feature Template
* `ip_directed_broadcast` - (Optional) Flag indicate whether directed IP braodcast is enabled or not for VPN interface type Feature Template, Allowed value: true, false  Default: false


<a id="nestedblock--ipv4"></a>

## Nested Schema for Basic ipv4
* `primary_address` - (Optional) Value of IPv4 prefix for VPN interface type Feature Template
* `secondary_address` - (Optional) Value of IPv4 prefix for VPN interface type Feature Template, Max Items: 4(see [below for nested schema](#nestedblock--secondary_addressv4))
* `dhcp_distance` - (Optional) Value of IPv4 DHCP distance for VPN interface type Feature Template, Range: [1, 65536], either primary_address or dhcp_distance is configured not both
* `dhcp_helper` - (Optional) List of IPv4 or IPv6 DHCP Helper for VPN interface type Feature Template, Allowed values: IPv4 or IPv6 address


<a id="nestedblock--ipv6"></a>

## Nested Schema for Basic ipv6
* `primary_address` - (Optional) Value of IPv6 prefix for VPN interface type Feature Template
* `secondary_address` - (Optional) Value of IPv6 prefix for VPN interface type Feature Template, Max Items: 2(see [below for nested schema](#nestedblock--secondary_addressv6))
* `dhcp_distance` - (Optional) Value of IPv6 DHCP distance for Vedge VPN interface type Feature Template, Range: [1, 65536], either primary_address or dhcp_distance is configured not both
* `dhcp_rapid_commit` - (Optional) Value of IPv6 DHCP rapid commit enabled or not for Vedge VPN interface type Feature Template
* `dhcp_helper` - (Optional) DHCP Helper value to designate the interface as a DHCP helper on a router for VPN interface type Feature Template, Max Items: 8(see [below for nested schema](#nestedblock--dhcp_helperV6))


<a id="nestedblock--allow_service"></a>

## Nested Schema for Tunnel allow_service
* `all` - (Optional) Flag indicating whether All serivce is enabled or not for VPN interface type Feature Template, Default: false
* `bgp` - (Optional) Flag indicating whether BGP is enabled or not for VPN interface type Feature Template, Default: false
* `dhcp` - (Optional) Flag indicating whether DHCP is enabled or not for VPN interface type Feature Template, Default: true
* `dns` - (Optional) Flag indicating whether DNS is enabled or not for VPN interface type Feature Template, Default: true
* `icmp` - (Optional) Flag indicating whether ICMP is enabled or not for VPN interface type Feature Template, Default: true
* `netconf` - (Optional) Flag indicating whether NETCONF is enabled or not for VPN interface type Feature Template, Default: false
* `ntp` - (Optional) Flag indicating whether NTP is enabled or not for VPN interface type Feature Template, Default: false
* `ospf` - (Optional) Flag indicating whether OSPF is enabled or not for VPN interface type Feature Template, Default: false
* `ssh` - (Optional) Flag indicating whether SSH is enabled or not for VPN interface  type Feature Template, Default: false
* `stun` -(Optional) Flag indicating whether STUN is enabled or not for VPN interface type Feature Template, Default: false
* `https` - (Optional) Flag indicating whether HTTPS is enabled or not for VPN interface type Feature Template, Default: true
* `snmp` - (Optional) Flag indicating whether SNMP is enabled or not for Cisco VPN interface type Feature Template, Default: true


<a id="nestedblock--encapsulation"></a>

## Nested Schema for Tunnel encapsulation
* `gre` - (Optional) Flag indicating whether gre is enabled or not for VPN interface type Feature Template, Default: false
* `gre_preference` - (Optional) Value of GRE preference for VPN interface type Feature Template, Range: [0, 4294967295]
* `gre_weight` - (Optional) Value of GRE weight for VPN interface type Feature Template, Range: [1, 255], Default: 1
* `ipsec` - (Optional) Flag indicating whether ipsec is enabled or not for VPN interface type Feature Template, Default: true
* `ipsec_preference` - (Optional) Value of IPsec preference for VPN interface type Feature Template, Range: [0, 4294967295]
* `ipsec_weight` - (Optional) Value of IPsec weight for VPN interface type Feature Template, Range: [1, 255], Default: 1
* `carrier` - (Optional) Carrier name or private network identifier to associate with the tunnel for VPN interface type Feature Template, Allowed values: "carrier1", "carrier2", "carrier3", "carrier4", "carrier5", "carrier6",
"carrier7", "carrier8", "default", Default: "default"
* `bind_loopback_tunnel` - (Optional) Name of a physical interface to bind to a loopback interface for VPN interface type Feature Template
* `last_resort_circuit` - (Optional) Flag indicating whether tunnel interface is used as the circuit of last resort for VPN interface type Feature Template, Default: false
* `hold_time` - (Optional) Value of hold time for Vedge VPN interface type Feature Template, Range: [100, 10000], Default: 7000
* `nat_refresh_interval` - (Optional) Value of interval between NAT refresh packets sent on a DTLS or TLS WAN transport connection for VPN interface type Feature Template, Range: [1, 60], Default: 5
* `hello_interval` - (Optional) Value of interval between Hello packets sent on a DTLS or TLS WAN transport connection for VPN interface type Feature Template, Range: [10, 600000], Default: 1000
* `hello_tolerance` - (Optional) Value of the time to wait for a Hello packet on a DTLS or TLS WAN transport connection before declaring that transport tunnel to be down for VPN interface type Feature Template, 
Range: [12, 6000], Default: 12
* `gre_tunnel_destination_ip` - (Optional) Value of GRE tunnel destination IP for VPN interface type Feature Template, Allowed value: IPv4 Address


<a id="nestedblock--natv4"></a>

## Nested Schema for NAT ipv4
* `nat_type` - (Optional) NAT type for Cisco VPN interface type Feature Template, Allowed value: "interface", "pool", "loopback" Default: "interface"
* `refresh_mode` - (Optional) Refresh mode for VPN interface type Feature Template, Allowed value: "bi-directional", "outbound", Default: "outbound"
* `log_nat_flow_creations_or_deletions` - (Optional) Flag indicating whether log nat flow creations or deletions are enabled or not for VPN interface type Feature Template, Default: false
* `udp_timeout` - (Optional) The time when NAT translations over UDP sessions time out for VPN interface type Feature Template, Range: [1, 65536] minutes, Default: 1
* `tcp_timeout` - (Optional) The time when NAT translations over TCP sessions time out for VPN interface type Feature Template, Range: [1, 65536] minutes, Default: 60
* `block_icmp` - (Optional) Flag indicating whether to block icmp or not for VPN interface type Feature Template, Default: true
* `respond_to_ping` - (Optional) Flag indicating whether to respond to ping or not for VPN interface type Feature Template, Default: false
* `nat_pool_range_start` - (Optional) Value of starting IP address for the NAT pool for VPN interface type Feature Template, Allowed value: IPv4 Address
* `nat_pool_range_end` - (Optional) Value of closing IP address for the NAT pool for VPN interface type Feature Template, Allowed value: IPv4 Address
* `nat_pool_prefix_length` - (Optional) Value of NAT pool prefix length for Cisco VPN interface type Feature Template, Range: [1, 65536], Default: 60
* `overload` - (Optional) Flag indicating whether to enable per-port translation or not for Cisco VPN interface type Feature Template, Default: true
* `nat_inside_source_loopback_interface` - (Optional) Value of NAT inside source loopback interface for Cisco VPN interface type Feature Template, String length Range: [1, 32]
* `port_forward` - (Optional) Configure port forwarding rule for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--port_forward))
* `static_nat` - (Optional) Configure static nat for VPN interface type Feature Template(see [below for nested schema](#nestedblock--static_nat))


<a id="nestedblock--natv6"></a>

## Nested Schema for NAT ipv6
* `nat64` - (Optional) Flag indicating whether to enable nat64 or not for VPN interface type Feature Template, Default: false

<a id="nestedblock--vrrpv4"></a>

## Nested Schema for VRRP ipv4
* `group_id` - (Required) The virtual router ID, which is a numeric identifier of the virtual router for VPN interface type Feature Template, 
Range: [1, 255]
* `track_prefix_list` - (Required) Value of remote prefixes, which is defined in a prefix list configured on the local router for VPN interface type Feature Template
* `priority` - (Optional) The priority level of the router. The router with the highest priority is elected as primary VRRP router. If two routers have the same priority, the one with the higher IP address is elected as primary VRRP router for VPN interface type Feature Template, Range: [1, 254], Default: 100
* `timer` - (Optional) Timer value specify how often the primary VRRP router sends VRRP advertisement messages. If subordinate routers miss three consecutive VRRP advertisements, they elect a new primary VRRP routers for VPN interface type Feature Template, Range for Vedge VPN interface: [1, 3600] Seconds, Range for Cisco VPN interface: [1, 40950] Milliseconds, Default: 1 Second
* `track_omp` - (Optional) Flag indicating whether to enable for VRRP to track the Overlay Management Protocol (OMP) session running on the WAN connection or not for VPN interface type Feature Template, Default: false
* `ip_address` - (Optional) Value of IPv4 address of the virtual router for VPN interface type Feature Template, Allowed value: IPv4 Address

<a id="nestedblock--vrrpv6"></a>

## Nested Schema for VRRP ipv6
* `group_id` - (Required) The virtual router ID, which is a numeric identifier of the virtual router for VPN interface type Feature Template, 
Range: [1, 255]
* `track_prefix_list` - (Required) Value of remote prefixes, which is defined in a prefix list configured on the local router for VPN interface type Feature Template
* `link_local_ipv6_address` - (Required) Value of IPv6 address of the virtual router for VPN interface type Feature Template, Allowed value: IPv6 Address
* `priority` - (Optional) The priority level of the router. The router with the highest priority is elected as primary VRRP router. If two routers have the same priority, the one with the higher IP address is elected as primary VRRP router for VPN interface type Feature Template, Range: [1, 254], Default: 100
* `timer` - (Optional) Timer value specify how often the primary VRRP router sends VRRP advertisement messages. If subordinate routers miss three consecutive VRRP advertisements, they elect a new primary VRRP routers for VPN interface type Feature Template, Range for Vedge VPN interface: [1, 3600] Seconds, Range for Cisco VPN interface: [1, 40950] Milliseconds, Default: 1 Second
* `track_omp` - (Optional) Flag indicating whether to enable for VRRP to track the Overlay Management Protocol (OMP) session running on the WAN connection or not for VPN interface type Feature Template, Default: false
* `global_ipv6_prefix` - (Optional) Value of global IPv6 prefix of the virtual router for VPN interface type Feature Template

<a id="nestedblock--advanced_options"></a>

## Nested Schema for 8021.X advanced_options
* `vlan` - (Optional) Configure VLAN for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--vlan))
* `dynamic_authentication_server` - (Optional) Configure Dynamic Authentication Server for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--dynamic_authentication_server))
* `mac_authentication_bypass` - (Optional) Configure MAC Authentication Bypass for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--mac_authentication_bypass))
* `request_attributes` - (Optional) Configure Request Attributes for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--request_attributes))

<a id="nestedblock--secondary_addressv4"></a>

## Nested Schema for Basic IPv4 secondary_address
* `address` - (Required) Value of IPv4 prefix for VPN interface type Feature Template

<a id="nestedblock--secondary_addressv6"></a>

## Nested Schema for Basic IPv6 secondary_address
* `address` - (Required) Value of IPv6 prefix for VPN interface type Feature Template


<a id="nestedblock--dhcp_helperV6"></a>

## Nested Schema for Basic IPv6 dhcp_helper
* `address` - (Required) Value of IPv6 address for VPN interface type Feature Template, Allowed value: IPv6 Address
* `vpn` - (Optional) Value of VPN id for VPN interface type Feature Template, Range: [1, 65536]

<a id="nestedblock--port_forward"></a>

## Nested Schema for NAT ipv4 port_forward
* `port_start_range` - (Required) Value of port starting range for Vedge VPN interface type Feature Template, Range: [0, 65535]
* `port_end_range` - (Required) Value of port ending range for Vedge VPN interface type Feature Template, Range: [0, 65535]
* `protocol` - (Required) Protocol for Vedge VPN interface type Feature Template, Allowed value: "tcp", "udp"
* `vpn` - (Required) VPN id for Vedge VPN interface type Feature Template, Range: [0, 65535]
* `private_ip` - (Required) Private IP address for Vedge VPN interface type Feature Template, Allowed value: IPv4 Address


<a id="nestedblock--static_nat"></a>

## Nested Schema for NAT ipv4 static_nat
* `source_ip` - (Required) Value of the inside local address as source IP address for VPN interface type Feature Template, Allowed value: IPv4 Address
* `translated_source_ip` - (Required) Value of the inside global address as the translated source IP address for VPN interface type Feature Template, Allowed value: IPv4 Address
* `source_vpn_id` - (Optional) Configures Source VPN ID, which is the Service VPN in which the host resides for VPN interface type Feature Template, Range: [0, 65530], Default: 0
* `static_nat_direction` - (Optional) Set the direction in which to perform network address translation for VPN interface type Feature Template, Allowed value for Vedge VPN interface type Feature Template : "inside", "outside" Default: "inside" Allowed value for Cisco VPN interface type Feature Template : "inside"
* `protocol` - (Optional) Protocol for Vedge VPN interface type Feature Template, Allowed value: "tcp", "udp"
* `source_port` - (Optional) Value of source port for Vedge VPN interface type Feature Template, Range: [0, 65535]
* `translate_port` - (Optional) Value of translate port for Vedge VPN interface type Feature Template, Range: [0, 65535]


<a id="nestedblock--vlan"></a>

## Nested Schema for vlan
* `authentication_fail_vlan` - (Optional) Value of authentication fail VLAN for Vedge VPN interface type Feature Template, Range: [1, 65536]
* `guest_vlan` - (Optional) Value of guest VLAN for Vedge VPN interface type Feature Template, Range: [1, 65536]
* `authentication_reject_vlan` - (Optional) Value of authentication reject VLAN for Vedge VPN interface type Feature Template, Range: [1, 65536]
* `default_vlan` - (Optional) Value of default VLAN for Vedge VPN interface type Feature Template, Range: [1, 65536]

<a id="nestedblock--dynamic_authentication_server"></a>

## Nested Schema for dynamic_authentication_server
* `das_port` - (Optional) Value of Dynamic Authentication Server port for Vedge VPN interface type Feature Template, Range: [1, 65535], 
Default: 3799,
* `client` - (Optional) Value of Client IP address for Vedge VPN interface type Feature Template, Allowed value: IPv4 Address
* `secret_key` - (Optional) Secret key Value for Vedge VPN interface type Feature Template
* `time_window` - (Optional) Time window value for Vedge VPN interface type Feature Template, Range: [10, 1000] Seconds, Default: 300,
* `required_timestamp` - (Optional) Flag indicating whether to enable required timestamp or not for Vedge VPN interface type Feature Template, Default: false
* `vpn` - (Optional) Value of VPN ID for Vedge VPN interface type Feature Template, Range: [0, 65535], Default: 0


<a id="nestedblock--mac_authentication_bypass"></a>

## Nested Schema for mac_authentication_bypass
* `server` - (Optional) Flag indicating whether to enable server or not for Vedge VPN interface type Feature Template, Default: false
* `mac_authentication_bypass_entries` - (Optional) List of MAC Authenticaion bypass entries for Vedge VPN interface type Feature Template, Max Items: 8


<a id="nestedblock--request_attributes"></a>

## Nested Schema for request_attributes
* `authentication` - (Optional) Configure Authentication for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--authentication))
* `accounting` - (Optional) Configure Accounting for Vedge VPN interface type Feature Template(see [below for nested schema](#nestedblock--accounting))

<a id="nestedblock--authentication"></a>

## Nested Schema for request_attributes authentication
* `id` - (Required) ID for Vedge VPN interface type Feature Template, Range: [1, 64]
* `syntax_choice` - (Required) Syntax Choice for Vedge VPN interface type Feature Template, Allowed value: "string", "integer", "octet"
* `value` - (Required) Value for Vedge VPN interface type Feature Template


<a id="nestedblock--accounting"></a>

## Nested Schema for request_attributes accounting
* `id` - (Required) ID for Vedge VPN interface type Feature Template, Range: [1, 64]
* `syntax_choice` - (Required) Syntax Choice for Vedge VPN interface type Feature Template, Allowed value: "string", "integer", "octet"
* `value` - (Required) Value for Vedge VPN interface type Feature Template










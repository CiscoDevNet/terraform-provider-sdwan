terraform {
  required_providers {
    sdwan = {
      version = "0.1"
      source  = "sdwan.com/labs/sdwan"
    }
  }
}

provider "sdwan" {  
   # cisco-sdwan user name
  username = "user name for vManage"
  # cisco-sdwan password
  password = "password for vManage"
  # cisco-sdwan url
  url      = "https://my-cisco-vmanage.com"
  insecure = true
}

data "sdwan_vpn_interface_feature_template" "name1"{
    template_name="Sample_Vedge_VPN_Interface"
}

output "output1" {
  value = data.sdwan_vpn_interface_feature_template.name1
}

data "sdwan_vpn_interface_feature_template" "name2"{
    template_name="Sample_Cisco_VPN_Interface"
}

output "output2" {
  value = data.sdwan_vpn_interface_feature_template.name2
}

resource "sdwan_vpn_interface_feature_template" "name3" {
  template_name = "Sample_Cisco_VPN_Interface_1"
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

resource "sdwan_vpn_interface_feature_template" "name4" {
  template_name = "Sample_Vedge_VPN_Interface_1"
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
         nat_inside_source_loopback_interface = "abc"

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


resource "sdwan_vpn_interface_feature_template" "example_cedge_vpn_interface" {
  template_name = "Cisco-VPN-Interface-demo"
  template_description = "Minimal VPN Interface Type cEdge Feature Template"
  template_type = "cisco_vpn_interface"
  device_type = ["vedge-CSR-1000v"]
  template_min_version = "20.4.1"
  factory_default = false
  template_definition {
    vpn_interface_basic {
        shutdown= true 
        description = "for demo"
     }
  }
}
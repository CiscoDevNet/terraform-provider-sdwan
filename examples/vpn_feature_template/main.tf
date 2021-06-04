terraform {
  required_providers {
    sdwan = {
      version = "0.1"
      source  = "sdwan.com/labs/sdwan"
    }
  }
}

# configure provider with your cisco sdwan credentials.
provider "sdwan" {
  # cisco-sdwan user name
  username = "user name for vManage"
  # cisco-sdwan password
  password = "password for vManage"
  # cisco-sdwan url
  url      = "https://my-cisco-vmanage.com"
  insecure = true
}

data "sdwan_vpn_feature_template" "example_datasource" {
  template_name = "example_template"
}

output "example_output" {
  value = data.sdwan_vpn_feature_template.example_datasource
}


resource "sdwan_vpn_feature_template" "exmple_resource1" {
  template_name = "example_cisco_vpn"
  template_type = "cisco_vpn"
  factory_default = false
  device_type = [ "vedge-C1111-8P" ]
  template_description = "For testing purpose"
  template_min_version = "15.0.0"
  template_definition {
    vpn_basic {
        vpn_id = 0
    }
  }
}

resource "sdwan_vpn_feature_template" "example_resource2" {
  template_name = "example_vpn_vedg"
  factory_default = false
  template_type = "vpn-vedge"
  device_type = [ "vedge-cloud" ]
  template_description = "For testing purpose"
  template_min_version = "15.0.0"
  template_definition {
    vpn_basic {
      vpn_id = 0 
      name = "sample_name"
      ecmp_hash_key = "true"
    }
     vpn_dns {
      primary_dns_ipv4 = "1.1.1.1"
      secondary_dns_ipv4 = "1.1.1.2"
      primary_dns_ipv6 = "1::"
      secondary_dns_ipv6 = "2::"
      vpn_host {
        hostname = "h1"
        ip = ["1.1.1.1","2.2.2.2"]
      }
      vpn_host {
        hostname = "h2"
        ip = ["1::"]
      }
     }
    ipv4_route {
      prefix = "1.1.1.1/1"
      gateway_type = "dhcp"
    }
    ipv4_route {
      prefix = "1.1.1.1/2"
      gateway_type = "null0"
      null0_distance = "6"
    }
    ipv4_route {
      prefix = "1.1.1.1/3"
      gateway_type = "null0"
    }
    ipv4_route {
      prefix = "1.1.1.1/4"
      gateway_type = "next-hop"
      next_hop {
        next_hop_address = "1.1.1.1"
      }
      next_hop {
        next_hop_address="2.2.2.2"
        next_hop_distance = "2"
      }
    }
    service_route {
      prefix = "1.1.1.1/4"
      service_type= "FW"
    }
    nat64 {
      pool_name = "smaple_name"
      start_address = "1.1.1.1"
      end_address = "2.2.2.2"
      overload = "true"
    }
    ipv6_route {
      prefix = "1::/1"
      gateway_type = "null0"
      null0_distance = "10"
    }
     ipv6_route {
      prefix = "1::/2"
      gateway_type = "null0"
    }
    ipv6_route {
      prefix = "2::/1"
      gateway_type = "next-hop"
      next_hop {
        next_hop_address = "1::"
        next_hop_distance = "12"
      }
      next_hop {
        next_hop_address = "2::"
      }
    }
    te_service_enabled = true
  }
}
terraform {
  required_providers {
    sdwan = {
      version = "0.1"
      source  = "sdwan.com/labs/sdwan"
    }
  }
}

#configure provider with your cisco sdwan credentials.
provider "sdwan" {
  # cisco-sdwan user name
  username = "user name for vManage"
  # cisco-sdwan password
  password = "password for vManage"
  # cisco-sdwan url
  url      = "https://my-cisco-vmanage.com"
  insecure = true
}

data "sdwan_system_feature_template" "name1"{
    template_name="Example"
}

output "output1" {
  value = data.sdwan_system_feature_template.name1
}

resource "sdwan_system_feature_template" "name2" {
  template_name = "Sample1"
  template_description = "For testing purposes"
  device_type = ["vedge-1000"]
  template_min_version = "15.0.0"
  template_definition {
    system_basic {  
      site_id = 5
      system_ip = "198.12.23.178"
      overlay_id = 2
      timezone = "Asia/Dubai"
      hostname = "hostname"
      location = "Place"
      device_groups = [ "abc","def" ]
      controller_groups = [ "12","34"]
      description = "some description"
      console_baud_rate = "19200"
      max_omp_sessions = 3
      tcp_optimization_enabled = true
    }

    system_gps {
      longitude = 23
      latitude = 19
    }

    system_trackers {
      tracker_name = "abcd"
      tracker_type = "ip-address"
      tracker_endpoint_value = "198.23.34.67"
    }
    
    system_trackers {
      tracker_name = "tracker2"
      tracker_type = "dns-name"
      tracker_endpoint_value = "dnsName"
    }
    system_trackers {
      tracker_name = "abc"
      tracker_type = "api-url"
      tracker_endpoint_value = "cisco.com/docs"
    }
    system_advanced {
      control_session_pps = 10
      system_tunnel_mtu = 560
      port_hop = false
      port_offset = 5
      dns_cache_timeout = 6
      track_transport = false
      vbond_local = true
      vbond_remote = "abcd"
      track_interface_tag = 433
      multicast_buffer_percent = 56
      usb_controller = true
      track_default_gateway = false
      host_policer_pps = 5689
      icmp_error_pps = 67
      allow_same_site_tunnels = true
      route_consistency_check = true
      admin_tech_on_failure = false
      idle_timeout = 67
      eco_friendly_mode = true
      on_demand_enable = true
      on_demand_idle_timeout = 78
    }
  }
  factory_default = false
}

resource "sdwan_system_feature_template" "name3" {
  
  template_name = "Sample2"
  template_description = "Description changed"
  device_type = ["vedge-1000"]
  template_min_version = "20.4.1"
  template_definition {
    system_basic {  
      site_id = 5
      system_ip = "198.12.23.178"
      timezone = "UTC"
    }
  }
  factory_default = false
}

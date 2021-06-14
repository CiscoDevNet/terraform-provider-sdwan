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

data "sdwan_system_feature_template" "name2"{
    template_name="vEdge-Example1"
}

output "output1" {
  value = data.sdwan_system_feature_template.name2
}

resource "sdwan_system_feature_template" "name3" {
  template_name = "vEdge-Example1"
  template_description = "For testing purposes"
  template_type = "System"
  device_type = ["vedge-1000", "vedge-cloud"]
  template_min_version = "15.0.0"
  template_definition {
    system_basic {  
      overlay_id = 2
      timezone = "Asia/Dubai"
      location = "Place"
      device_groups = [ "abc","def" ]
      controller_groups = [ "12","34"]
      description = "some description"
      console_baud_rate = "19200"
      max_omp_sessions = 3
    }

    system_gps {
      longitude = 23
      latitude = 19
    }

    system_trackers {
      tracker_name = "abcd"
      tracker_endpoint_type = "ip-address"
      tracker_endpoint_value = "198.23.34.67"
    }
    
    system_trackers {
      tracker_name = "tracker2"
      tracker_endpoint_type = "dns-name"
      tracker_endpoint_value = "dnsName"
    }
    system_trackers {
      tracker_name = "abc"
      tracker_endpoint_type = "api-url"
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

resource "sdwan_system_feature_template" "name4" {
  
  template_name = "vEdge-Example2"
  template_description = "Description changed"
  template_type = "System"
  device_type = ["vedge-1000", "vedge-cloud"]
  template_min_version = "20.4.1"
  template_definition {
    system_basic {
    }

    system_advanced {      
    }
  }
  factory_default = false
}

resource "sdwan_system_feature_template" "name5" {
  template_name = "CSR-Example1"
  template_description = "For testing purposes"
  device_type = ["vedge-CSR-1000v"]
  template_type = "Cisco System"
  template_min_version = "15.0.0"
  template_definition {
    system_basic {  
      overlay_id = 2
      timezone = "Asia/Dubai"
      location = "Place"
      device_groups = [ "abc","def" ]
      controller_groups = [ "12","34"]
      description = "some description"
      console_baud_rate = "19200"
      max_omp_sessions = 3
    }

    system_gps {
      longitude = 23
      latitude = 19
    }

    system_trackers {
      tracker_name = "abcd"
      tracker_type = "static-route"
      tracker_endpoint_type = "ip-address"
      tracker_endpoint_value = "198.23.34.67"
    }
    
    system_trackers {
      tracker_name = "tracker2"
      tracker_endpoint_type = "dns-name"
      tracker_endpoint_value = "dnsName"
    }
    system_trackers {
      tracker_name = "abc"
      tracker_type = "interface"
      tracker_endpoint_type = "api-url"
      tracker_endpoint_value = "cisco.com/docs"
    }
    system_advanced {
      control_session_pps = 10
      port_hop = false
      port_offset = 5
      track_transport = false
      track_interface_tag = 433
      track_default_gateway = false
      admin_tech_on_failure = false
      idle_timeout = 67
      on_demand_enable = true
      on_demand_idle_timeout = 78
    }
  }
  factory_default = false
}

resource "sdwan_system_feature_template" "example_cedge_system" {
  
  template_name = "CSR-Example2"
  template_description = "Description changed"
  device_type = ["vedge-CSR-1000v"]
  template_type = "Cisco System"
  template_min_version = "20.4.1"
  factory_default = false
  template_definition {    
    system_basic {  
    }

    system_advanced {
    }
  }
}
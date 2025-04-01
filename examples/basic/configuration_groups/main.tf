resource "sdwan_system_feature_profile" "system_01" {
  name        = "system_01"
  description = "My system feature profile"
}

resource "sdwan_transport_feature_profile" "transport_01" {
  name        = "transport_01"
  description = "My transport feature profile"
}

resource "sdwan_system_basic_feature" "system_01_basic" {
  name               = "system_01_basic"
  feature_profile_id = sdwan_system_feature_profile.system_01.id
}

resource "sdwan_system_aaa_feature" "system_01_aaa" {
  name               = "system_01_aaa"
  feature_profile_id = sdwan_system_feature_profile.system_01.id
  server_auth_order  = ["local"]
  users = [{
    name     = "admin"
    password = "admin"
  }]
}

resource "sdwan_system_bfd_feature" "system_01_bfd" {
  name               = "system_01_bfd"
  feature_profile_id = sdwan_system_feature_profile.system_01.id
}

resource "sdwan_system_global_feature" "system_01_global" {
  name               = "system_01_global"
  feature_profile_id = sdwan_system_feature_profile.system_01.id
}

resource "sdwan_system_logging_feature" "system_01_logging" {
  name               = "system_01_logging"
  feature_profile_id = sdwan_system_feature_profile.system_01.id
}

resource "sdwan_system_omp_feature" "system_01_omp" {
  name               = "system_01_omp"
  feature_profile_id = sdwan_system_feature_profile.system_01.id
}

resource "sdwan_transport_wan_vpn_feature" "transport_01_wan_vpn" {
  name               = "transport_01_wan_vpn"
  feature_profile_id = sdwan_transport_feature_profile.transport_01.id
  vpn                = 0
}

resource "sdwan_transport_wan_vpn_interface_ethernet_feature" "transport_01_wan_vpn_interface_ethernet" {
  name                         = "transport_01_wan_vpn_interface_ethernet"
  feature_profile_id           = sdwan_transport_feature_profile.transport_01.id
  transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.transport_01_wan_vpn.id
  interface_name               = "GigabitEthernet1"
  shutdown                     = false
  ipv4_dhcp_distance           = 1
  tunnel_interface             = true
  tunnel_interface_encapsulations = [
    {
      encapsulation = "ipsec"
    }
  ]
}

resource "sdwan_configuration_group" "config_group_01" {
  name                = "config_group_01"
  description         = "My config group"
  solution            = "sdwan"
  feature_profile_ids = [sdwan_system_feature_profile.system_01.id, sdwan_transport_feature_profile.transport_01.id]
  devices = [{
    id     = "C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B"
    deploy = true
    variables = [
      {
        name  = "host_name"
        value = "edge1"
      },
      {
        name  = "pseudo_commit_timer"
        value = 0
      },
      {
        name  = "site_id"
        value = 1
      },
      {
        name  = "system_ip"
        value = "10.1.1.1"
      },
      {
        name  = "ipv6_strict_control"
        value = "false"
      }
    ]
  }]
  feature_versions = [
    sdwan_system_basic_feature.system_01_basic.version,
    sdwan_system_aaa_feature.system_01_aaa.version,
    sdwan_system_bfd_feature.system_01_bfd.version,
    sdwan_system_global_feature.system_01_global.version,
    sdwan_system_logging_feature.system_01_logging.version,
    sdwan_system_omp_feature.system_01_omp.version,
    sdwan_transport_wan_vpn_feature.transport_01_wan_vpn.version,
    sdwan_transport_wan_vpn_interface_ethernet_feature.transport_01_wan_vpn_interface_ethernet.version,
  ]
}

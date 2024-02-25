# Attach device template to a device and provide values 
# for associated variables
resource "sdwan_attach_feature_device_template" "attach_Edge" {
  id = "a5e2ba39-87a6-4d38-9053-7ca4c781857d"
  devices = [{
    id = "C8K-E2D91A30-814B-F22D-4E7B-7D39BE74EE32"
    variables = {
      system_site_id                           = "102"
      system_system_ip                         = "10.10.1.11"
      system_longitude                         = "-121.932"
      system_latitude                          = "37.411"
      system_host_name                         = "terraform_dc-cedge01"
      vpn_0_if_ipv4_address                    = "10.10.23.6/30"
      vpn_0_if_description                     = "GigabitEthernet2.wan-rtr01"
      vpn_0_if_name                            = "GigabitEthernet2"
      public_internet_vpn_0_if_ipv4_address    = "10.10.23.38/30"
      public_internet_vpn_0_if_description     = "internet-link"
      public_internet_vpn_0_if_name            = "GigabitEthernet4"
      public_internet_0_vpn_next_hop_ip_addres = "10.10.23.37"
      vpn_0_next_hop_ip_address                = "10.10.23.5"
      vpn_512_if_ipv4_address                  = "10.10.20.172/24"
      vpn_512_if_description                   = "port.sbx-mgmt"
      vpn_512_if_name                          = "GigabitEthernet1"
      vpn_512_next_hop_ip_address              = "10.10.20.254"
      vpn_1_if_ipv4_address                    = "10.10.20.182/24"
      vpn_1_if_description                     = "port.dc-phys"
      vpn_1_if_name                            = "GigabitEthernet3"
      vpn_1_next_hop_ip_address                = "10.10.20.254"
    }
  }]
}

resource "sdwan_transport_wan_vpn_interface_ipsec_feature" "example" {
  name                                = "Example"
  description                         = "My Example"
  feature_profile_id                  = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  transport_wan_vpn_profile_parcel_id = "140331f6-5418-4755-a059-13c77eb96037"
  interface_name                      = "ipsec987"
  shutdown                            = true
  interface_description               = "ipsec987"
  ipv4_address                        = "9.7.5.4"
  ipv4_subnet_mask                    = "255.255.255.0"
  tunnel_source_ipv4_address          = "1.3.5.88"
  tunnel_source_ipv4_subnet_mask      = "255.255.255.0"
  tunnel_source_interface             = "GigabitEthernet8"
  tunnel_destination_ipv4_address     = "2.55.67.99"
  tunnel_destination_ipv4_subnet_mask = "255.255.255.0"
  application_tunnel_type             = "none"
  tcp_mss                             = 1460
  clear_dont_fragment                 = false
  ip_mtu                              = 1500
  dpd_interval                        = 10
  dpd_retries                         = 3
  ike_preshared_key                   = "123"
  ike_version                         = 1
  ike_integrity_protocol              = "main"
  ike_rekey_interval                  = 14400
  ike_ciphersuite                     = "aes256-cbc-sha1"
  ike_diffie_hellman_group            = "16"
  ike_id_local_end_point              = "xxx"
  ike_id_remote_end_point             = "xxx"
  ipsec_rekey_interval                = 3600
  ipsec_replay_window                 = 512
  ipsec_ciphersuite                   = "aes256-gcm"
  perfect_forward_secrecy             = "group-16"
  tunnel_route_via                    = "2222"
}

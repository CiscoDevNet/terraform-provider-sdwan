resource "sdwan_system_remote_access_feature" "example" {
  name                                = "Example"
  description                         = "My Example"
  feature_profile_id                  = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  connection_type_ssl                 = false
  any_connect_eap_authentication_type = "user"
  ipv4_pool_size                      = 50
  ipv6_pool_size                      = 1024
  enable_certificate_list_check       = false
  psk_authentication_type             = "aaa"
  psk_authentication_pre_shared_key   = "Cisco123"
  radius_group_name                   = "radius-1"
  aaa_derive_name_from_peer_identity  = "MyPassword"
  aaa_enable_accounting               = false
  ikev2_local_ike_identity_type       = "EMAIL"
  ikev2_local_ike_identity_value      = "abc@xyz.com"
  ikev2_security_association_lifetime = 86400
  ikev2_anti_dos_threshold            = 99
  ipsec_enable_anti_replay            = false
  ipsec_anti_replay_window_size       = 64
  ipsec_security_association_lifetime = 3600
  ipsec_enable_perfect_foward_secrecy = false
}

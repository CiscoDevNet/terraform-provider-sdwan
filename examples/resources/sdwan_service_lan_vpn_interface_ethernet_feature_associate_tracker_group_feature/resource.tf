resource "sdwan_service_lan_vpn_interface_ethernet_feature_associate_tracker_group_feature" "example" {
  feature_profile_id                            = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  service_lan_vpn_feature_id                    = "140331f6-5418-4755-a059-13c77eb96037"
  service_lan_vpn_interface_ethernet_feature_id = "140331f6-5418-4755-a059-13c77eb96037"
  service_tracker_group_feature_id              = "140331f6-5418-4755-a059-13c77eb96037"
}

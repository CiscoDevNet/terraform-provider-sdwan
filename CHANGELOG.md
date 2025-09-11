## 0.7.2 (unreleased)

- BREAKING CHANGE: Fix type of `port` attribute of `sdwan_port_list_policy_object` to support multiple values, [Link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/460)
- BREAKING CHANGE: Rename `target_interface` attribute of `sdwan_application_priority_qos_policy` resource to `target_interfaces`
- BREAKING CHANGE: Adds `ipv4_configuration_type` and `ipv6_configuration_type` to `sdwan_service_lan_vpn_interface_ethernet_feature` resource
- Add `community` support to the `sdwan_route_policy_definition` resource and data source

## 0.7.1

- Fix issue, where concurrent update of `localized_policy` resources fails with `Failed to update variables` error
- Fix issue in `sdwan_attach_feature_device_template` resource where manual detachment of template in the UI is not recognized by Terraform, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/340)
- Fix issue with embedded objects in the `transport_acl` and `service_acl` resources and data sources, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/381)
- Adds `translate_rib_metric` support to the `sdwan_service_routing_bgp_feature` resource and data source
- Adds `translate_rib_metric` support to the `sdwan_service_routing_ospf_feature` resource and data source
- Adds `translate_rib_metric` support to the `sdwan_service_routing_ospfv3_ipv4_feature` resource and data source
- Adds `translate_rib_metric` support to the `sdwan_service_routing_ospfv3_ipv6_feature` resource and data source
- Use `Set` type for `match_entries` and `actions` in `sdwan_application_aware_routing_policy_definition`, `sdwan_custom_control_topology_policy_definition`, `sdwan_ipv4_acl_policy_definition`, `sdwan_ipv4_device_acl_policy_definition`, `sdwan_ipv6_acl_policy_definition`, `sdwan_ipv6_device_acl_policy_definition`, `sdwan_route_policy_definition`, `sdwan_traffic_data_policy_definition`, `sdwan_zone_based_firewall_policy_definition`

## 0.7.0

- BREAKING CHANGE: Rename `tls_ssl_profile_version` attribute of `sdwan_tls_ssl_decryption_policy_definition` resource to `tls_ssl_profile_policy_version`
- BREAKING CHANGE: Rename `sla_class_list` attribute of `sdwan_application_aware_routing_policy_definition` resource to `sla_class_list_id`
- BREAKING CHANGE: Rename `preferred_color_group_list` attribute of `sdwan_application_aware_routing_policy_definition` resource to `preferred_color_group_list_id`
- BREAKING CHANGE: Rename `preferred_color_group_list` attribute of `sdwan_traffic_data_policy_definition` resource to `preferred_color_group_list_id`
- Fix issue, where NAT IPv4 and IPv6 is always configured on `sdwan_transport_wan_vpn_interface_ethernet_feature`, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/485)
- Add missing options under `unsupported_features` attribute of `sdwan_configuration_group`, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/478)
- Add `sdwan_policy_group` resource and data source
- Adds `enhanced_app_aware_routing` support to the `sdwan_cisco_system_feature_template` resource and data source
- Add `gateway` attribute for `ipv4_static_routes` and `ipv6_static_routes` in `sdwan_service_lan_vpn_feature` resource
- BREAKING CHANGE: Rename `interface` attribute of `sdwan_cisco_vpn_feature_template` resource to `interfaces`
- Fix type of `interfaces` attribute of `sdwan_cisco_vpn_feature_template` under `ipv4_static_ipsec_routes` & `ipv4_static_gre_routes`, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/474)
- Fix type of `as_path_prepend` attribute of `sdwan_service_route_policy_feature` and `sdwan_transport_route_policy_feature`, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/474)
- Fix issue causing the `sdwan_policy_object_preferred_color_group` resource to fail, [Link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/495)
- BREAKING CHANGE: Refactor `max_number_of_prefixes`, `restart_interval` and `threshold` of `sdwan_service_routing_bgp_feature` resource, [Link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/475)

## 0.6.2

- Fix issue causing changes to applied feature templates to fail, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/417)
- Fix issue with service lan vpn feature when route-policy is not defined, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/450)
- Prevent updates to `sdwan_activate_centralized_policy` resources being reverted, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/452)
- BREAKING CHANGE: Fix type of `dscp` attribute of `sdwan_application_aware_routing_policy_definition` and `sdwan_traffic_data_policy_definition`, [Link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/460)
- Fix issue where required attributes were marked as optional in the documentation
- Fix issue causing `sdwan_transport_route_policy_feature` to fail to apply when `default_action` is set to `reject`, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/439)
- BREAKING CHANGE: Refactor `max_number_of_prefixes`, `restart_interval` and `threshold` of `sdwan_transport_routing_bgp_feature` resource, [Link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/413)

## 0.6.1

- BREAKING CHANGE: Replaces `no_authentication` attribute of `sdwan_transport_cellular_profile_feature` with `requires_authentication`
- Correctly apply default values with `sdwan_transport_route_policy_feature` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/410)
- Change `server_auth_order` attribute of `sdwan_system_aaa_feature` resource to a set, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/430)
- Support changes in `sdwan_system_mrf_feature` resource in version 20.13+, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/427)
- Applying default action with `ipv4_device_access_policy` or other resources can cause idempotency issues, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/412)
- Optimize and fix attachment of device template to multiple devices, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/417)

## 0.6.0

- Delete configuration group when an error is encountered during creation, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/390)
- Refresh devices of a configuration group before deletion, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/391)
- BREAKING CHANGE: Convert `feature_profiles` attribute of `sdwan_configuration_group` to set of strings and rename to `feature_profile_ids`
- Add `sdwan_tag` resource and data source
- Adds `sdwan_dns_security_policy` resource and data source

## 0.5.3

- Fix template rendering issue with `sdwan_cisco_vpn_interface_cellular_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/378)
- Fix reading of lists in various features, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/380), [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/381)
- Add `sdwan_transport_wan_vpn_interface_cellular_feature_associate_tracker_feature` resource and data source
- Add `sdwan_transport_wan_vpn_interface_cellular_feature_associate_tracker_group_feature` resource and data source
- Add `sdwan_transport_wan_vpn_interface_gre_feature_associate_tracker_feature` resource and data source
- Add `sdwan_transport_wan_vpn_interface_ipsec_feature_associate_tracker_feature` resource and data source
- Add `sdwan_service_lan_vpn_interface_ethernet_feature_associate_tracker_feature` resource and data source
- Add `sdwan_service_lan_vpn_interface_ethernet_feature_associate_tracker_group_feature` resource and data source

## 0.5.2

- Update documentation categories
- Add `sdwan_service_dhcp_server_feature` resource and data source
- Add `sdwan_service_lan_vpn_interface_ethernet_feature_associate_dhcp_server_feature` resource and data source
- Add `sdwan_service_lan_vpn_interface_ipsec_feature_associate_dhcp_server_feature` resource and data source
- Add `sdwan_service_lan_vpn_interface_svi_feature_associate_dhcp_server_feature` resource and data source
- Add `hub` option to `tunnel_qos_mode` attribute of `sdwan_cisco_vpn_interface_cellular_feature` resource
- Add `cellular` template types to `sdwan_feature_device_template` resource

## 0.5.1

- Enhance update logic of policy objects, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/pull/360)
- Fix issue with `sdwan_other_ucse_feature` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/320)
- BREAKING CHANGE: Adds `gateway` attribute to `sdwan_transport_wan_vpn_feature` resource
- Allow port ranges in ACL feature resources, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/364)
- Add `sdwan_transport_cellular_controller_feature` resource and data source
- Add `sdwan_transport_cellular_profile_feature` resource and data source
- Add `sdwan_transport_wan_vpn_feature_associate_routing_bgp_feature` resource and data source
- Add `sdwan_transport_wan_vpn_feature_associate_routing_ospf_feature` resource and data source
- Add `sdwan_transport_wan_vpn_feature_associate_routing_ospfv3_ipv4_feature` resource and data source
- Add `sdwan_transport_wan_vpn_feature_associate_routing_ospfv3_ipv6_feature` resource and data source
- Add `sdwan_service_lan_vpn_feature_associate_routing_bgp_feature` resource and data source
- Add `sdwan_service_lan_vpn_feature_associate_routing_eigrp_feature` resource and data source
- Add `sdwan_service_lan_vpn_feature_associate_multicast_feature` resource and data source
- Add `sdwan_service_lan_vpn_feature_associate_routing_ospf_feature` resource and data source
- Add `sdwan_service_lan_vpn_feature_associate_routing_ospfv3_ipv4_feature` resource and data source
- Add `sdwan_service_lan_vpn_feature_associate_routing_ospfv3_ipv6_feature` resource and data source
- Add `sdwan_transport_wan_vpn_interface_ethernet_feature_associate_tracker_feature` resource and data source
- Add `sdwan_transport_wan_vpn_interface_ethernet_feature_associate_tracker_group_feature` resource and data source
- Add `sdwan_transport_wan_vpn_interface_ethernet_feature_associate_ipv6_tracker_feature` resource and data source
- Add `sdwan_transport_wan_vpn_interface_ethernet_feature_associate_ipv6_tracker_group_feature` resource and data source

## 0.5.0

- Add `sdwan_policy_object_application_list` resource and data source
- Add `sdwan_policy_object_sla_class_list` resource and data source
- Add `sdwan_policy_object_app_probe_class` resource and data source
- Add `sdwan_policy_object_as_path_list` resource and data source
- Add `sdwan_policy_object_vpn_group` resource and data source
- Add `sdwan_policy_object_security_data_ipv4_prefix_list` resource and data source
- Add `sdwan_policy_object_security_fqdn_list` resource and data source
- Add `sdwan_policy_object_security_geolocation_list` resource and data source
- Add `sdwan_policy_object_security_ips_signature` resource and data source
- Add `sdwan_policy_object_security_local_domain_list` resource and data source
- Add `sdwan_policy_object_security_port_list` resource and data source
- Add `sdwan_policy_object_security_scalable_group_tag_list` resource and data source
- Add `sdwan_policy_object_security_url_allow_list` resource and data source
- Add `sdwan_policy_object_security_url_block_list` resource and data source
- Add `sdwan_policy_object_preferred_color_group` resource and data source
- Add `sdwan_policy_object_security_identity_list` resource and data source
- Add `sdwan_policy_object_security_local_application_list` resource and data source
- Add `sdwan_policy_object_standard_community_list` resource and data source
- Add `sdwan_dns_security_feature_profile` resource and data source
- Add `sdwan_sig_security_feature_profile` resource and data source
- Add `sdwan_embedded_security_feature_profile` resource and data source
- Enhance `sdwan_configuration_group` resource and data source to support device associations, variables and deployment
- Add `sdwan_service_multicast_feature` resource and data source
- Add `sdwan_service_routing_eigrp_feature` resource and data source
- Add `sdwan_service_wireless_lan_feature` resource and data source
- Add `sdwan_transport_t1_e1_controller_feature` resource and data source
- Add import to `sdwan_attach_feature_device_template` and `sdwan_attach_feature_device_template` resources
- Add `sdwan_application_priority_traffic_policy_policy` resource and data source
- Add `sdwan_transport_ipv4_acl_feature` resource and data source
- Add `sdwan_service_ipv4_acl_feature` resource and data source
- Add `sdwan_transport_ipv6_acl_feature` resource and data source
- Add `sdwan_service_ipv6_acl_feature` resource and data source
- Add workaround for issue when deleting security policy objects where the reference count is not updated
- BREAKING CHANGE: Adds `ipv4_configuration_type` and `ipv6_configuration_type` to `sdwan_transport_management_vpn_interface_ethernet_feature` resource
- BREAKING CHANGE: Adds `ipv4_configuration_type` and `ipv6_configuration_type` to `sdwan_transport_wan_vpn_interface_ethernet_feature` resource
- BREAKING CHANGE: Adds `gateway` to `sdwan_transport_management_vpn_feature` resource

## 0.4.1

- Update documentation categories

## 0.4.0

- Fix issue when reading deleted `sdwan_cli_config_profile_parcel` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/291)
- Add `sdwan_transport_tracker_group_feature` resource and data source
- Add `sdwan_transport_tracker_feature` resource and data source
- Add `sdwan_service_tracker_group_feature` resource and data source
- Add `sdwan_transport_ipv6_tracker_group_feature` resource and data source
- Add `sdwan_service_object_tracker_feature` resource and data source
- Add `sdwan_service_object_tracker_group_feature` resource and data source
- BREAKING CHANGE: Rename `enable_crl_check` attribute of `sdwan_system_remote_access_profile_parcel` resource to `enable_certificate_list_check`
- BREAKING CHANGE: Rename `psk_selection` attribute of `sdwan_system_remote_access_profile_parcel` resource to `psk_authentication_type`
- BREAKING CHANGE: Rename `aaa_derive_name_identity` attribute of `sdwan_system_remote_access_profile_parcel` resource to `aaa_derive_name_from_peer_identity`
- BREAKING CHANGE: Rename `aaa_derive_name_domain` attribute of `sdwan_system_remote_access_profile_parcel` resource to `aaa_derive_name_from_peer_domain`
- BREAKING CHANGE: Removes `any_connect_eap_profile_download_status` and `any_connect_eap_profile_file_name` attributes of `sdwan_system_remote_access_profile_parcel`
- Add `sdwan_policy_object_feature_profile` resource and data source
- Add `sdwan_policy_object_class_map` resource and data source
- Add `sdwan_policy_object_color_list` resource and data source
- Add `sdwan_policy_object_data_ipv6_prefix_list` resource and data source
- Add `sdwan_policy_object_data_ipv4_prefix_list` resource and data source
- Add `sdwan_policy_object_tloc_list` resource and data source
- Add `sdwan_policy_object_expanded_community_list` resource and data source
- Add `sdwan_policy_object_extended_community_list` resource and data source
- Add `sdwan_policy_object_mirror` resource and data source
- Add `sdwan_policy_object_policer` resource and data source
- Add `sdwan_policy_object_ipv4_prefix_list` resource and data source
- Add `sdwan_policy_object_ipv6_prefix_list` resource and data source
- Add `sdwan_other_feature_profile` resource and data source
- Add `sdwan_other_ucse_feature` resource and data source
- Add `sdwan_other_thousandeyes_feature` resource and data source
- Add `sdwan_transport_gps_feature` resource and data source
- Add `sdwan_transport_routing_ospf_feature` resource and data source
- Add `sdwan_transport_routing_ospfv3_ipv4_feature` resource and data source
- Add `sdwan_transport_routing_ospfv3_ipv6_feature` resource and data source
- Add `sdwan_service_routing_ospf_feature` resource and data source
- Add `sdwan_service_routing_ospfv3_ipv4_feature` resource and data source
- Add `sdwan_service_routing_ospfv3_ipv6_feature` resource and data source
- Add `sdwan_service_routing_bgp_feature` resource and data source
- Add `sdwan_service_switchport_feature` resource and data source
- Add `sdwan_transport_route_policy_feature` resource and data source
- Add `sdwan_service_route_policy_feature` resource and data source
- BREAKING CHANGE: Rename all profile parcel resources and data sources being used in configuration groups to "Features" with a `_feature` suffix
- BREAKING CHANGE: Rename `sdwan_application_priority_qos_policy_profile_parcel` resource and data source to `sdwan_application_priority_qos_policy`
- Fix issue when reading deleted `sdwan_configuration_group` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/297)
- BREAKING CHANGE: Rename `service_lan_vpn_profile_parcel_id` attribute of `sdwan_service_lan_vpn_interface_ethernet_feature`, `sdwan_service_lan_vpn_interface_gre_feature`, `sdwan_service_lan_vpn_interface_ipsec_feature` and `sdwan_service_lan_vpn_interface_svi_feature` resource to `service_lan_vpn_feature_id`
- BREAKING CHANGE: Rename `transport_management_vpn_profile_parcel_id` attribute of `sdwan_transport_management_vpn_interface_ethernet_feature` resource to `transport_management_vpn_feature_id`
- BREAKING CHANGE: Rename `transport_wan_vpn_profile_parcel_id` attribute of `sdwan_transport_wan_vpn_interface_cellular_feature`, `sdwan_transport_wan_vpn_interface_ethernet_feature`, `sdwan_transport_wan_vpn_interface_gre_feature`, `sdwan_transport_wan_vpn_interface_ipsec_feature` and `sdwan_transport_wan_vpn_interface_t1_e1_serial_feature` resource to `transport_wan_vpn_feature_id`

## 0.3.13

- Add workaround for issue when concurrently deleting security policy objects where the reference count is not updated
- BREAKING CHANGE: Remove default value of `failure_mode` attribute of `sdwan_security_policy` resource

## 0.3.12

- Add `protocol_type` attribute to `sdwan_zone_based_firewall_policy_definition` resource and data source
- Fix payload issue with `sdwan_cisco_sig_credentials_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/247)

## 0.3.11

- Add workaround for issue when concurrently deleting localized policy objects where the reference count is not updated

## 0.3.10

- Add `sdwan_transport_management_vpn_interface_ethernet_profile_parcel` resource and data source
- Add `sdwan_service_lan_vpn_interface_ethernet_profile_parcel` resource and data source
- Fix type of `controller_group_list` attribute of `sdwan_cisco_system_feature_template` resource and data source, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/243)
- Add `sdwan_application_priority_feature_profile` resource and data source
- Add `sdwan_application_priority_qos_policy_profile_parcel` resource and data source
- Add `sdwan_transport_wan_vpn_interface_ethernet_profile_parcel` resource and data source
- Add support for `bfd_metrics_exporting` and `exporting_interval` in `sdwan_cflowd_policy_definition` resource and data source, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/122)
- Add `sdwan_zone_based_firewall_policy_definition` resource and data source, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/246)
- BREAKING CHANGE: Rename `basic_configuration_description` attribute of `sdwan_transport_management_vpn_interface_ethernet_profile_parcel` resource to `interface_description`, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/270)
- BREAKING CHANGE: Rename `basic_configuration_description` attribute of `sdwan_transport_management_vpn_profile_parcel` resource to `vpn_description`, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/270)
- BREAKING CHANGE: Rename `ipv4_route_gateway_next_ho` attribute of `sdwan_transport_management_vpn_profile_parcel` resource to `next_hops`, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/266)
- Add workaround for issue when deleting localized policy objects where the reference count is not updated
- Add `sdwan_service_lan_vpn_interface_ipsec_profile_parcel` resource and data source
- Add `sdwan_service_lan_vpn_interface_gre_profile_parcel` resource and data source
- Add `sdwan_service_lan_vpn_interface_svi_profile_parcel` resource and data source
- Add `sdwan_transport_wan_vpn_interface_gre_profile_parcel` resource and data source
- Add `sdwan_transport_wan_vpn_interface_ipsec_profile_parcel` resource and data source
- Add `sdwan_transport_wan_vpn_interface_t1_e1_serial_profile_parcel` resource and data source
- Add `sdwan_transport_wan_vpn_interface_cellular_profile_parcel` resource and data source
- Add `sdwan_system_ipv4_device_access_profile_parcel` resource and data source
- Add `sdwan_system_ipv6_device_access_profile_parcel` resource and data source

## 0.3.9

- Fix issue with rendering feature template payloads
- Add `sdwan_transport_routing_bgp_profile_parcel` resource and data source
- Add `sdwan_service_lan_vpn_profile_parcel` resource and data source
- Add `sdwan_system_ntp_profile_parcel` resource and data source
- Add `sdwan_system_performance_monitoring_profile_parcel` resource and data source
- Handle 429 responses with retry logic
- Add `sdwan_system_remote_access_profile_parcel` resource and data source
- Add `sdwan_transport_ipv6_tracker_profile_parcel` resource and data source
- Add `sdwan_service_tracker_profile_parcel` resource and data source
- Use existing variable values if none are provided with `sdwan_attach_feature_device_template` resource
- Add `sdwan_cli_config_profile_parcel` resource and data source
- Add `sdwan_system_logging_profile_parcel` resource and data source
- Add `sdwan_transport_management_vpn_profile_parcel` resource and data source
- Fix issue when updating certain policy definitions, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/230)

## 0.3.8

- Add `sdwan_system_flexible_port_speed_profile_parcel` resource and data source
- Add `sdwan_system_mrf_profile_parcel` resource and data source
- Add `sdwan_system_snmp_profile_parcel` resource and data source
- Add `sdwan_system_basic_profile_parcel` resource and data source
- Fix issue with deleted feature profile resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/204)
- Add `sdwan_transport_wan_vpn_profile_parcel` resource and data source
- Add `advertise_hub_tlocs`, `all_hubs_are_equal`, `tloc_list_id`, `ipv4_prefix_list_ids`, `ipv6_prefix_list_ids` and `preference` to `sdwan_hub_and_spoke_topology_policy_definition` resource and data source
- Fix issue with deleted profile parcel resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/204)
- Fix issue with incorrect feature template payloads, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/213)
- Add warning and continue if there are no changes to deploy with `sdwan_attach_feature_device_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/220)
- Add `sdwan_system_security_profile_parcel` resource and data source
- Fix issue with `dot1x_enable` attribute of `sdwan_switchport_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/219)
- Add possible types to `sdwan_localized_policy` and `sdwan_feature_device_template` resource
- Add `type` read-only attribute to localized policy definitions
- Add `type` read-only attribute to centralized policy definitions
- Add `sdwan_configuration_group` resource and data source

## 0.3.7

- Add `tracker`, `interval`, `interval_variable` and `zscaler_surrogate_refresh_time` to `sdwan_cisco_secure_internet_gateway_feature_template` resource and data source
- Add `sdwan_cisco_trustsec_feature_template` resource and data source
- Add `sdwan_cisco_vpn_interface_gre_feature_template` resource and data source
- Add `sdwan_cellular_profile_feature_template` resource and data source
- Add `sdwan_cedge_igmp_feature_template` resource and data source
- Add `sdwan_cedge_multicast_feature_template` resource and data source
- Add `sdwan_cisco_wireless_lan_feature_template` resource and data source
- Add `sdwan_gps_feature_template` resource and data source
- Add `sdwan_eigrp_feature_template` resource and data source
- Add `sdwan_cedge_pim_feature_template` resource and data source
- Add `sdwan_cisco_ospfv3_feature_template` resource and data source
- Add `sdwan_security_app_hosting_feature_template` resource and data source
- Add `sdwan_cellular_cedge_profile_feature_template` resource and data source
- Add `sdwan_vpn_interface_cellular_feature_template` resource and data source
- Add `sdwan_vpn_interface_dsl_pppoa_feature_template` resource and data source
- Add `sdwan_vpn_interface_dsl_pppoe_feature_template` resource and data source
- Add `sdwan_vpn_interface_ethernet_pppoe_feature_template` resource and data source
- Add `sdwan_vpn_interface_dsl_ipoe_feature_template` resource and data source
- Add `sdwan_vpn_interface_multilink_feature_template` resource and data source
- Add `sdwan_vpn_interface_t1_e1_serial_feature_template` resource and data source
- Add `sdwan_cellular_controller_feature_template` resource and data source
- BREAKING CHANGE: Increases the minimum version for all profile parcel resource and data sources to 20.12.0
- Fix issue when changing multiple already deployed templates including device templates
- Fix error when refreshing state of `sdwan_attach_feature_device_template` resource and device template does no longer exist
- Fix refreshing state of empty lists in responses
- Fix issue with empty `action_entries` list in `sdwan_route_policy_definition` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/186)

## 0.3.6

- Fix issue with address family configuration of `sdwan_cisco_bgp_feature_template` resource
- Fix issue with area configuration of `sdwan_cisco_ospf_feature_template` resource
- BREAKING CHANGE: Rename `send_lifetime` attribute of `sdwan_cisco_security_feature_template` resource to `send_lifetime_local`
- BREAKING CHANGE: Rename `accept_lifetime` attribute of `sdwan_cisco_security_feature_template` resource to `accept_lifetime_local`
- Improve error handling when detaching device templates
- Fix templating issue with `sdwan_cisco_vpn_interface_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/165)
- Fix templating issue with `sdwan_cisco_sig_credentials_feature_template` resouce, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/162)
- Enable `global` flag for `sdwan_cisco_sig_credentials_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/163)
- Fix templating issue with `sdwan_cisco_secure_internet_gateway_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/164)
- Fix templating issue with `sdwan_cisco_secure_internet_gateway_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/160)
- BREAKING CHANGE: Rename `aup_enabled` attribute of `sdwan_cisco_secure_internet_gateway_feature_template` resource and data source to `zscaler_aup_enabled`
- BREAKING CHANGE: Rename `aup_block_internet_until_accepted` attribute of `sdwan_cisco_secure_internet_gateway_feature_template` resource and data source to `zscaler_aup_block_internet_until_accepted`
- BREAKING CHANGE: Rename `aup_force_ssl_inspection` attribute of `sdwan_cisco_secure_internet_gateway_feature_template` resource and data source to `zscaler_aup_force_ssl_inspection`
- BREAKING CHANGE: Rename `aup_timeout` attribute of `sdwan_cisco_secure_internet_gateway_feature_template` resource and data source to `zscaler_aup_timeout`
- Fix templating of certain boolean values, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/133), [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/134)
- BREAKING CHANGE: Rename `dot1x_voice_vlan` attribute of `sdwan_switchport_feature_template` resource and data source to `voice_vlan`
- Use `set` type for list attributes with primitive values
- Fix templating issue with `sdwan_cisco_security_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/157)
- Fix templating of endpoint and object trackers with `sdwan_cisco_system_feature_template` resource, [link](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues/159)

## 0.3.5

- Fix import operation of profile parcel resources
- BREAKING CHANGE: Convert `protocol`, `source_port`, `destination_port` attribute of `sdwan_application_aware_routing_definition` to strings to support multiple values
- Add `sdwan_system_omp_profile_parcel` resource and data source
- Add `sdwan_security_policy` resource and data source
- BREAKING CHANGE: Rename `von_configurations` attribute of `sdwan_cedge_aaa_feature_template` resource and data source to `vpn_configurations`
- Add `strict` and `fallbackToBestPath` options to SLA class action parameters of `sdwan_application_aware_routing_policy_definition` resource and data source
- BREAKING CHANGE: Rename `radius_dynamic_author_cts_authorization_list` attribute of `sdwan_cedge_aaa_feature_template` resource and data source to `radius_trustsec_cts_authorization_list`
- Add option to configure `ipv6-unicast` address families to `sdwan_cisco_bgp_feature_template` resource and data source
- Add `stub` and `nssa` options to `sdwan_cisco_ospf_feature_template` resource and data source

## 0.3.4

- Add `sdwan_switchport_feature_template` resource and data source
- Add `sdwan_vpn_interface_svi_feature_template` resource and data source

## 0.3.3

- Fix next hop action and match entry of `sdwan_route_policy_definition` resource and data source

## 0.3.2

- BREAKING CHANGE: Rename `source_port` attribute to `source_ports` of `sdwan_ipv4_acl_policy_definition` resource and data source
- BREAKING CHANGE: Rename `destination_port` attribute to `destination_ports` of `sdwan_ipv4_acl_policy_definition` resource and data source
- BREAKING CHANGE: Rename `source_port` attribute to `source_ports` of `sdwan_ipv6_acl_policy_definition` resource and data source
- BREAKING CHANGE: Rename `destination_port` attribute to `destination_ports` of `sdwan_ipv6_acl_policy_definition` resource and data source
- BREAKING CHANGE: Rename `source_port` attribute to `source_ports` of `sdwan_ipv4_device_acl_policy_definition` resource and data source
- BREAKING CHANGE: Rename `source_port` attribute to `source_ports` of `sdwan_ipv6_device_acl_policy_definition` resource and data source
- Add `traffic_class` attribute to `sdwan_ipv6_acl_policy_definition` resource and data source
- Fix next hop action entry of `sdwan_route_policy_definition` resource and data source

## 0.3.1

- Fix idempotency issue with DNS redirect configuration of `sdwan_traffic_data_policy_definition` resource
- Add option to select both `gre` and `ipsec` encaps for local TLOC set action of `sdwan_traffic_data_policy_definition` resource
- Fix issue with `exportTo` action entry of `sdwan_custom_control_topology_policy_definition` resource
- Fix issue with service TLOC lists of `sdwan_traffic_data_policy_definition` resource

## 0.3.0

- Add `sdwan_system_feature_profile` resource and data source
- Add `sdwan_service_feature_profile` resource and data source
- Add `sdwan_transport_feature_profile` resource and data source
- Add `sdwan_cli_feature_profile` resource and data source
- Add `sdwan_system_aaa_profile_parcel` resource and data source
- Add `sdwan_system_banner_profile_parcel` resource and data source
- Add `sdwan_system_bfd_profile_parcel` resource and data source
- Fix issue with `sdwan_centralized_policy` not accepting all valid parameters
- Fix issue with `sdwan_custom_control_topology_policy_definition` not accepting all possible values
- BREAKING CHANGE: Convert `protocol`, `source_port`, `destination_port` attributes of `sdwan_traffic_data_policy_definition` to strings to support multiple values
- Fix issue with `sdwan_traffic_data_policy_definition` not supporting local and restrict options
- Add `sdwan_dns_security_policy_definition` resource and data source
- Add `sdwan_system_global_profile_parcel` resource and data source
- BREAKING CHANGE: Rename `sdwan_acl_policy_definition` resource and data source to `sdwan_ipv4_acl_policy_definition` and update attributes
- BREAKING CHANGE: Rename `sdwan_device_acl_policy_definition` resource and data source to `sdwan_ipv4_device_acl_policy_definition` and update attributes
- Add `sdwan_ipv6_acl_policy_definition` resource and data source
- Add `sdwan_ipv6_device_acl_policy_definition` resource and data source

## 0.2.11

- Fix idempotency issue with various resource and boolean values
- Remove redundant `per_tunnel_qos` and `per_tunnel_qos_aggregator` attributes from `sdwan_cisco_vpn_interface_feature_template` resource and data source
- Add `base_action` attribute to `sdwan_custom_control_topology_policy_definition` resource and data source
- Add `base_action` attribute to `sdwan_traffic_data_policy_definition` resource and data source
- Allow full path variables in `sdwan_attach_feature_device_template` resource `variables` section
- Always include empty lists in payloads to fix various GUI compatibility issues
- Skip `sdwan_attach_feature_device_template` template detachment when device template is no longer attached

## 0.2.10

- Fix issue with `sdwan_cisco_ospf_feature_template` resource when not configuring route policies
- Fix issue with `sdwan_cisco_snmp_feature_template` resource when configuring SNMPv2
- Add `sdwan_device` data source
- Add `sdwan_vedge_inventory` data source
- Fix idempotency issue with `sdwan_cisco_ospf_feature_template` resource when enabling `default_information_originate`

## 0.2.9

- Add `sdwan_advanced_malware_protection_policy_definition` resource and data source
- BREAKING CHANGE: Rename `group` attribute of `sdwan_cedge_aaa_feature_template` resource and data source to `groups` and fix type
- Use type `Set` for `device_types` attributes of feature template resources and data sources
- Add `sdwan_tls_ssl_decryption_policy_definition` resource and data source
- Add `sdwan_advanced_inspection_profile_policy_definition` resource and data source

## 0.2.8

- Correctly populate system-level variables when attaching existing (factory default) feature templates
- Add `sdwan_intrusion_prevention_policy_definition` resource and data source
- Add `sdwan_tls_ssl_profile_policy_definition` resource and data source
- BREAKING CHANGE: Rename `timers_spf_initial_holf` attribute of `sdwan_cisco_ospf_feature_template` resource and data source to `timers_spf_initial_hold`
- Fix issue with setting an `omp_tag` match condition using a `sdwan_route_policy_definition` resource

## 0.2.7

- Add option to feature template data sources to identify templates by its name

## 0.2.6

- Fix deployment of attached device templates without feature template changes

## 0.2.5

- Add `sdwan_local_application_list_policy_object` resource and data source
- Add `sdwan_domain_list_policy_object` resource and data source
- Add `sdwan_ips_signature_list_policy_object` resource and data source
- Add `sdwan_allow_url_list_policy_object` resource and data source
- Add `sdwan_block_url_list_policy_object` resource and data source
- Add `sdwan_zone_list_policy_object` resource and data source
- Add `sdwan_port_list_policy_object` resource and data source
- Add `sdwan_protocol_list_policy_object` resource and data source
- Add `sdwan_geo_location_list_policy_object` resource and data source
- Add `sdwan_data_fqdn_prefix_list_policy_object` resource and data source
- Add `sdwan_object_group_policy_definition` resource and data source
- Add `sdwan_rule_set_policy_definition` resource and data source

## 0.2.4

- BREAKING CHANGE: Rename `ipv6_adresses` attribute of `sdwan_cisco_vpn_interface_feature_template` resource and data source to `ipv6_addresses`

## 0.2.3

- Add `sdwan_centralized_policy` resource and data source
- Add `sdwan_activate_centralized_policy` resource

## 0.2.2

- Add `sdwan_application_aware_routing_policy_definition` resource and data source
- Add `sdwan_cflowd_policy_definition` resource and data source
- Add `sdwan_traffic_data_policy_definition` resource and data source
- Add `set_parameters` attribute to `sdwan_acl_policy_definition` resource and data source
- BREAKING CHANGE: Remove `dscp` and `next_hop` attributes of `sdwan_acl_policy_definition` resource and data source
- BREAKING CHANGE: Rename `layer2cos` attribute of `sdwan_rewrite_rule_policy_definition` resource and data source to `layer2_cos`
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_class_map_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_app_probe_class_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_mirror_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_policer_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_preferred_color_group_policy_object` resource and data source
- BREAKING CHANGE: Remove `entries` attribute of `sdwan_sla_class_policy_object` resource and data source
- Add `app_probe_class_version` attribute to `sdwan_sla_class_policy_object` resource and data source

## 0.2.1

- Add `sdwan_cisco_ospf_feature_template` resource and data source
- Add `sdwan_cisco_vpn_interface_ipsec_feature_template` resource and data source
- Add `sdwan_cisco_secure_internet_gateway_feature_template` resource and data source
- Add `sdwan_hub_and_spoke_topology_policy_definition` resource and data source
- Add `sdwan_mesh_topology_policy_definition` resource and data source
- Add `sdwan_custom_control_topology_policy_definition` resource and data source
- Add `sdwan_vpn_membership_policy_definition` resource and data source

## 0.2.0

- BREAKING CHANGE: Completely revamped the provider based on `github.com/netascode/terraform-provider-sdwan` codebase, replacing all existing resources and data sources

## 0.1.0 (July 23, 2021)

- Initial Release

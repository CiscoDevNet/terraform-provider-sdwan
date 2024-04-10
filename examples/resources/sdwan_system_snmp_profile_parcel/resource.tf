resource "sdwan_system_snmp_profile_parcel" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  shutdown           = false
  contact_person     = "wixie.cisco"
  location_of_device = "SHANGHAI"
  views = [
    {
      name = "VIEW1"
      oid = [
        {
          id      = "1.3.6.1.4.1.9.9.394"
          exclude = false
        }
      ]
    }
  ]
  communities = [
    {
      name          = "example"
      user_label    = "COMMUNITY1"
      view          = "VIEW1"
      authorization = "read-only"
    }
  ]
  groups = [
    {
      name           = "GROUP1"
      security_level = "auth-priv"
      view           = "VIEW1"
    }
  ]
  users = [
    {
      name                    = "USER1"
      authentication_protocol = "sha"
      authentication_password = "$CRYPT_CLUSTER$su56l1Z0Tk4Qc9N7+T/uOg==$sD6b0HLqEdI+RNwsEOoLcQ=="
      privacy_protocol        = "aes-256-cfb-128"
      privacy_password        = "$CRYPT_CLUSTER$su56l1Z0Tk4Qc9N7+T/uOg==$sD6b0HLqEdI+RNwsEOoLcQ=="
      group                   = "GROUP1"
    }
  ]
  trap_target_servers = [
    {
      vpn_to_reach_trap_target_server         = 1
      ip_address_of_snmp_server               = "10.75.221.156"
      udp_port_number_of_snmp_server          = 161
      user_label                              = "TARGET1"
      user                                    = "USER1"
      source_interface_for_outgoing_snmp_trap = "GigabitEthernet1"
    }
  ]
}

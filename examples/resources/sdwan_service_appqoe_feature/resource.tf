resource "sdwan_service_appqoe_feature" "example" {
  name               = "Example"
  description        = "My Example"
  feature_profile_id = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  appqoe_device_role = "forwarder"
  forwarder_controller_groups = [
    {
      appnav_controllers = [
        {
          address = "192.168.2.1"
          vpn     = 1
        }
      ]
    }
  ]
  forwarder_service_node_groups = [
    {
      service_nodes = [
        {
          address = "192.168.2.2"
        }
      ]
    }
  ]
  forwarder_service_contexts = [
    {
      appnav_controller_group = "ACG-APPQOE"
      service_node_group      = "SNG-APPQOE"
      enable                  = true
      vpn                     = 0
    }
  ]
}

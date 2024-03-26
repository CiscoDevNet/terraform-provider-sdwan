resource "sdwan_cedge_igmp_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  interfaces = [
    {
      name = "Ethernet0"
      join_groups = [
        {
          group_address = "235.1.1.1"
          source        = "1.2.3.4"
        }
      ]
    }
  ]
}

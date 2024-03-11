resource "sdwan_cedge_igmp_feature_template" "example" {
  name         = "Example"
  description  = "My Example"
  device_types = ["vedge-C8000V"]
  interface = [
    {
      name = "example"
      join_group = [
        {
          group_address = "1.2.3.4"
          source        = "1.2.3.4"
        }
      ]
    }
  ]
}

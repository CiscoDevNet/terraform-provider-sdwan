resource "sdwan_cedge_multicast_feature_template" "example" {
  name             = "Example"
  description      = "My Example"
  device_types     = ["vedge-C8000V"]
  spt_only         = true
  local_replicator = true
  threshold        = 200
}

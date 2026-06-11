resource "sdwan_network_hierarchy_node" "example" {
  parent_id   = "9cdc05d1-5306-41ef-8487-85829a4cfbe6"
  name        = "EMEA-Group"
  description = "EMEA group"
  type        = "group"
}

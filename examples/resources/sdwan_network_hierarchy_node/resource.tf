resource "sdwan_network_hierarchy_node" "example" {
  parent_group = "Global"
  name         = "EMEA-Group"
  description  = "EMEA group"
  type         = "group"
}

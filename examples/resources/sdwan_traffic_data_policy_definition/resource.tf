resource "sdwan_traffic_data_policy_definition" "example" {
  name           = "Example"
  description    = "My description"
  default_action = "drop"
  sequences = [
    {
      id      = 1
      name    = "Seq1"
      type    = "applicationFirewall"
      ip_type = "ipv4"
      match_entries = [
        {
          type                = "appList"
          application_list_id = "e3aad846-abb9-425f-aaa8-9ed17b9c8d7c"
        }
      ]
      action_entries = [
        {
          type = "log"
          log  = true
        }
      ]
    }
  ]
}

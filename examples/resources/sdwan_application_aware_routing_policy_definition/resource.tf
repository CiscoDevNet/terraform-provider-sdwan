resource "sdwan_application_aware_routing_policy_definition" "example" {
  name        = "Example"
  description = "My description"
  sequences = [
    {
      id      = 1
      name    = "Region1"
      ip_type = "ipv4"
      match_entries = [
        {
          type                = "appList"
          application_list_id = "e3aad846-abb9-425f-aaa8-9ed17b9c8d7c"
        }
      ]
      action_entries = [
        {
          type                       = "backupSlaPreferredColor"
          backup_sla_preferred_color = "bronze"
        }
      ]
    }
  ]
  default_action = [
    {
      sla_class_list_id = "ef743b4c-ca10-4cac-8f11-8ce902a47f57"
    }
  ]
}

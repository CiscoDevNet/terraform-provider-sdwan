resource "sdwan_local_application_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      application = "cisco-collab-video"
    }
  ]
}

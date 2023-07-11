resource "sdwan_application_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      application = "netflix"
    }
  ]
}

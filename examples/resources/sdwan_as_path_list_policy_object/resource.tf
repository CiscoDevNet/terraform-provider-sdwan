resource "sdwan_as_path_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      as_path = "^1239_[0-9]*$"
    }
  ]
}

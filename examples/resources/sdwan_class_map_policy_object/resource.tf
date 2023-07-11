resource "sdwan_class_map_policy_object" "example" {
  name = "Example"
  entries = [
    {
      queue = 2
    }
  ]
}

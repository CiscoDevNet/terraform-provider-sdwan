resource "sdwan_color_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      color = "blue"
    }
  ]
}

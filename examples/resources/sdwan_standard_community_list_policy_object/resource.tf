resource "sdwan_standard_community_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      community = "100:1000"
    }
  ]
}

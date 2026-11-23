resource "sdwan_extended_community_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      community = "rt 100:10"
    }
  ]
}

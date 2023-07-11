resource "sdwan_preferred_color_group_policy_object" "example" {
  name = "Example"
  entries = [
    {
      primary_color_preference   = "blue bronze"
      primary_path_preference    = "direct-path"
      secondary_color_preference = "3g"
      secondary_path_preference  = "multi-hop-path"
      tertiary_color_preference  = "custom1"
      tertiary_path_preference   = "all-paths"
    }
  ]
}

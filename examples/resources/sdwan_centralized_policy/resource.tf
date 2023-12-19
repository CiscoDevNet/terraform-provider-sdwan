resource "sdwan_centralized_policy" "example" {
  name        = "Example"
  description = "My description"
  definitions = [
    {
      id   = "2081c2f4-3f9f-4fee-8078-dcc8904e368d"
      type = "data"
      entries = [
        {
          site_list_ids = ["2081c2f4-3f9f-4fee-8078-dcc8904e368d"]
          vpn_list_ids  = ["7d0c2444-8743-4414-add0-866945ea9f70"]
          direction     = "in"
        }
      ]
    }
  ]
}

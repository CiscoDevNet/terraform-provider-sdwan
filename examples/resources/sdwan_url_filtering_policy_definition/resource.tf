resource "sdwan_url_filtering_policy_definition" "example" {
  name                  = "Example"
  description           = "My description"
  mode                  = "security"
  alerts                = ["blacklist"]
  web_categories        = ["alcohol-and-tobacco"]
  web_categories_action = "allow"
  web_reputation        = "moderate-risk"
  target_vpns           = ["1"]
  block_page_action     = "text"
  block_page_contents   = "Access to the requested page has been denied. Please contact your Network Administrator"
}

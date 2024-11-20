resource "sdwan_policy_object_unified_url_filtering" "example" {
  name                  = "Example"
  description           = "My Example"
  feature_profile_id    = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  web_categories_action = "block"
  web_categories        = ["confirmed-spam-sources"]
  web_reputation        = "suspicious"
  url_allow_list_id     = "2ad58d78-59ee-46d3-86dd-7b6b7ca09f38"
  url_block_list_id     = "2ad58d78-59ee-46d3-86dd-7b6b7ca09f38"
  block_page_action     = "text"
  block_page_contents   = "Access to the requested page has been denied. Please contact your Network Administrator"
  enable_alerts         = true
  alerts                = ["blacklist"]
}

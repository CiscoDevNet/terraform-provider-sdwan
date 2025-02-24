resource "sdwan_tag" "example" {
  name        = "TAG_1"
  description = "My tag"
  devices     = ["C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B"]
}

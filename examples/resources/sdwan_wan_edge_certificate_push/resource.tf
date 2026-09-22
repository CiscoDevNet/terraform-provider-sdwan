resource "sdwan_wan_edge_certificate_push" "example" {
  triggers = {
    # change this value (e.g. to a timestamp) to force a new push to the controllers
    revision = "1"
  }
}

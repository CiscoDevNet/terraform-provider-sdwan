resource "sdwan_mirror_policy_object" "example" {
  name = "Example"
  entries = [
    {
      remote_destination_ip = "10.1.1.1"
      source_ip             = "10.2.1.1"
    }
  ]
}

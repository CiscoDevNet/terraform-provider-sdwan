resource "sdwan_ips_signature_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      generator_id = 1111
      signature_id = 2222
    }
  ]
}

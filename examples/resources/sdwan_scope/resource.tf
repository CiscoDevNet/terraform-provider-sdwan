resource "sdwan_scope" "example" {
  name        = "west1"
  description = "West Coast Domain1"
  objects = [
    {
      object_type = "feature-profile"
      object_ids  = ["003c4e0b-a2da-486d-8ee7-0e6dd4551bae"]
    }
  ]
}

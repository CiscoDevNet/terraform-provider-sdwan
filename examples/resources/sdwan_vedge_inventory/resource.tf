resource "sdwan_vedge_inventory" "example" {
  name = "va-001"
  devices = [
    {
      chassis_number = "2081c2f4-3f9f-4fee-8078-dcc8904e368d"
      site_id        = "400"
      serial_number  = "2AJC9DJ"
      hostname       = "vEdge-123"
      validity       = "valid"
      device_type    = "vedge"
    }
  ]
}

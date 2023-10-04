resource "sdwan_device" "example" {
  serial_number = "2AJC9DJ"
  name          = "va-001"
  devices = [
    {
      device_id     = "2081c2f4-3f9f-4fee-8078-dcc8904e368d"
      uuid          = "2081c2f4-3f9f-4fee-8078-dcc8904e368d"
      site_id       = "400"
      serial_number = "2AJC9DJ"
      hostname      = "vEdge-123"
      reachability  = "reachable"
      status        = "normal"
      state         = "green"
    }
  ]
}

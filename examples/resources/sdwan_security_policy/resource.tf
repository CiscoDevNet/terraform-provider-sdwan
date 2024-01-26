resource "sdwan_security_policy" "example" {
  name        = "Example"
  description = "Example"
  mode        = "security"
  use_case    = "custom"
  definitions = [
    {
      id   = "7d299c34-981c-4fb3-9167-6be44ab1691f"
      type = "urlFiltering"
    }
  ]
  failure_mode = "close"
  logging = [
    {
      external_syslog_server_ip  = "10.0.0.1"
      external_syslog_server_vpn = "123"
    }
  ]
}

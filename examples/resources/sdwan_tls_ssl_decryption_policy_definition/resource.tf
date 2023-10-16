resource "sdwan_tls_ssl_decryption_policy_definition" "example" {
  name           = "Example"
  description    = "My description"
  mode           = "security"
  default_action = "noIntent"
  network_rules = [
    {
      base_action = "doNotDecrypt"
      rule_id     = 4
      rule_name   = "Example"
      rule_type   = "sslDecryption"
      source_and_destination_configuration = [
        {
          option = "destinationIp"
          value  = "10.0.0.0/12"
        }
      ]
    }
  ]
  ssl_decryption_enabled        = "true"
  expired_certificate           = "drop"
  untrusted_certificate         = "drop"
  certificate_revocation_status = "none"
  unknown_revocation_status     = "drop"
  unsupported_protocol_versions = "drop"
  unsupported_cipher_suites     = "drop"
  failure_mode                  = "close"
  rsa_key_pair_modulus          = "2048"
  ec_key_type                   = "P384"
  certificate_lifetime_in_days  = 1
  minimal_tls_version           = "TLSv1.2"
  use_default_ca_cert_bundle    = true
}

resource "sdwan_policy_object_unified_tls_ssl_decryption" "example" {
  name                          = "Example"
  description                   = "My Example"
  feature_profile_id            = "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac"
  expired_certificate           = "drop"
  untrusted_certificate         = "drop"
  certificate_revocation_status = "none"
  unsupported_protocol_versions = "drop"
  unsupported_cipher_suites     = "drop"
  failure_mode                  = "close"
  default_ca_certificate_bundle = true
  rsa_keypair_modules           = "2048"
  ec_key_type                   = "P256"
  certificate_lifetime          = "1"
  minimal_tls_ver               = "TLSv1"
}

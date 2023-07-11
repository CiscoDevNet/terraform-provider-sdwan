resource "sdwan_feature_device_template" "test" {
  name        = "device_group_1"
  description = "My device template"
  device_type = "vedge-ISR-4331"
  general_templates = [{
    id   = sdwan_cisco_system_feature_template.system.id
    type = sdwan_cisco_system_feature_template.system.template_type
  }]
}

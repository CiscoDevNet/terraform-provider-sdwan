resource "sdwan_cisco_system_feature_template" "system" {
  name               = "TF_SYSTEM_1"
  description        = "My cisco system feature template"
  device_types       = ["vedge-C8000V"]
  hostname_variable  = "var_hostname"
  system_ip_variable = "var_system_ip"
  site_id_variable   = "var_site_id"
  console_baud_rate  = "115200"
}

resource "sdwan_feature_device_template" "device_template_1" {
  name        = "TF_DT1_ALL"
  description = "My device template 1"
  device_type = "vedge-C8000V"
  general_templates = [{
    id      = sdwan_cisco_system_feature_template.system.id
    version = sdwan_cisco_system_feature_template.system.version
    type    = sdwan_cisco_system_feature_template.system.template_type
  }]
}

resource "sdwan_attach_feature_device_template" "test" {
  id      = sdwan_feature_device_template.device_template_1.id
  version = sdwan_feature_device_template.device_template_1.version
  devices = [{
    id = "C8K-CC678D1C-8EDF-3966-4F51-ABFAB64F5ABE"
    variables = {
      var_site_id   = "1001"
      var_system_ip = "1.1.1.1"
      var_hostname  = "router1"
    }
  }]
}

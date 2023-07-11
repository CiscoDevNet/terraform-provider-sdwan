---
subcategory: "Guides"
page_title: "Updating Templates and Policies"
description: |-
    Updating Templates and Policies
---

# Updating Templates and Policies

Similary to the UI workflow, once device templates are attached to a device, updates to templates or policies require the device template to be re-attached to the device.

Each template and policy resource in Terraform exposes a version attribute which can be used to link resources and trigger a re-attachment of the device template if required. Due to brevity the following configuration is not complete and only shows a single feature template being used:

```terraform
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
```

This provider supports updates to a single feature template, device template and policy per device within a single `terraform apply` operation. This is due to vManage locking the templates and policies once a change has been made. If for example two feature templates used in the same device template are being updated, a `terraform apply` operation will show the following warning:

```
│ Warning: Client Warning
│ 
│   with sdwan_cisco_system_feature_template.system,
│   on main.tf line q, in resource "sdwan_cisco_system_feature_template" "system":
│  1: resource "sdwan_cisco_system_feature_template" "system" {
│ 
│ Failed to modify template due to template being locked by another change. Template changes will not be applied. Re-run 'terraform apply' to try again.
```

A subsequent `terraform plan` will show that the feature template has not been updated. Another `terraform apply` will eventually also apply the second updated feature template.

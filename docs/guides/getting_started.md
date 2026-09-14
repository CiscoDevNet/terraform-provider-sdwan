---
subcategory: "Guides"
page_title: "Getting Started - Classic Templates"
description: |-
    Getting Started - Classic Templates
---

# Getting Started - Classic Templates

This guide covers the **classic** device/feature template model (pre-20.12). If you are starting a new deployment on SD-WAN Manager 20.15 or 20.18, use [Getting Started - Configuration and Policy Groups](https://registry.terraform.io/providers/CiscoDevNet/sdwan/latest/docs/guides/configuration_groups) instead - classic templates remain supported for existing deployments but are not the recommended model for new ones.

This example demonstrates how the provider can be used to attach an existing device template to a device. The configuration for the first part can be found [here](https://github.com/CiscoDevNet/terraform-provider-sdwan/tree/main/examples/basic/getting_started). The configuration for the second part can be found [here](https://github.com/CiscoDevNet/terraform-provider-sdwan/tree/main/examples/basic/getting_started_part_2).

## Environment

This example is built specifically for the [Cisco SD-WAN 20.10 Sandbox](https://devnetsandbox.cisco.com/DevNet/catalog/sd-wan). 
For use with other environments, it is required to adapt the code accordingly since it refers to specific template/device IDs and names available in the mentioned sandbox environment.

Navigate to the [Cisco SD-WAN 20.10 page](https://devnetsandbox.cisco.com/DevNet/catalog/sd-wan) and click **Launch**. Select the preferred schedule and press the **Launch** button. The setup typically takes around 15 minutes. During the setup process and after successful completion, you are informed via email.

Note the IP address, port, username and password for HTTPS provided in the **Credentials** tab of the sandbox page for a later step.

Connect to the sandbox VPN based on the instructions and credentials provided in the success email.

### Prepare the Sandbox Device 

This sample code was tested with the device **dc-cedge01** in the sandbox. The mentioned device is attached to a device template by default. Before we can attach a template via Terraform to the device, it is required to detach the device. 

Therefore, open Cisco Catalyst SD-WAN Manager and follow the below steps:   

- Open the **Menu**
- Click the **Configuration** menu element
- Click the **Templates** menu element
- Click the **Device Templates** tab
- Click the three dots in the row of the **DC_cEdge_Template**
- Select **Detach devices**
- Move the device **dc-cedge01** to the **Selected devices'** list
- Press **Detach** and wait for the task status to change to **Success**.

## Provider Configuration

First of all we need to add the necessary provider configuration to the Terraform configuration file:

```hcl
terraform {
  required_providers {
    sdwan = {
      source = "CiscoDevNet/sdwan"
    }
  }
}

provider "sdwan" {
  username = "admin"
  password = "password"
  url      = "https://10.1.1.1"
}
```

## Attach Existing Device Template

In order to attach a device template, we first need to either create it or refer to an existing device template by its ID.

```hcl
resource "sdwan_attach_feature_device_template" "attach_Edge" {
  id = "a5e2ba39-87a6-4d38-9053-7ca4c781857d"
  devices = [{
    id = "C8K-E2D91A30-814B-F22D-4E7B-7D39BE74EE32"
    variables = {
      system_site_id                           = "102"
      system_system_ip                         = "10.10.1.11"
      system_longitude                         = "-121.932"
      system_latitude                          = "37.411"
      system_host_name                         = "terraform_dc-cedge01"
      vpn_0_if_ipv4_address                    = "10.10.23.6/30"
      vpn_0_if_description                     = "GigabitEthernet2.wan-rtr01"
      vpn_0_if_name                            = "GigabitEthernet2"
      public_internet_vpn_0_if_ipv4_address    = "10.10.23.38/30"
      public_internet_vpn_0_if_description     = "internet-link"
      public_internet_vpn_0_if_name            = "GigabitEthernet4"
      public_internet_0_vpn_next_hop_ip_addres = "10.10.23.37"
      vpn_0_next_hop_ip_address                = "10.10.23.5"
      vpn_512_if_ipv4_address                  = "10.10.20.172/24"
      vpn_512_if_description                   = "port.sbx-mgmt"
      vpn_512_if_name                          = "GigabitEthernet1"
      vpn_512_next_hop_ip_address              = "10.10.20.254"
      vpn_1_if_ipv4_address                    = "10.10.20.182/24"
      vpn_1_if_description                     = "port.dc-phys"
      vpn_1_if_name                            = "GigabitEthernet3"
      vpn_1_next_hop_ip_address                = "10.10.20.254"
    }
  }]
}
```

## Create New Feature and Device Templates

Before we continue we need to detach the previously attached device template by running `terraform destroy`. After detaching the device template we can replace the existing Terraform config with following parts. At first we create a new system feature template:

```hcl
resource "sdwan_cisco_system_feature_template" "TF_system_template" {
  name               = "TF_system_template"
  description        = "Test bootstrap system template"
  device_types       = ["vedge-C8000V"]
  system_ip_variable = "var_system_ip"
  hostname_variable  = "var_hostname"
  site_id_variable   = "var_site_id"
  longitude          = -130
  latitude           = 40
  console_baud_rate  = "9600"
}
```

Some of the parameter are device specific and we therefore define variables (indicated by `_variable` attribute suffix) instead of hardcoding the values.

For the remaining feature templates we rely on factory default templates where we first need to resolve the ID of templates by using data sources.

```hcl
data "sdwan_cisco_logging_feature_template" "Factory_Default_Cisco_Logging_Template" {
  name = "Factory_Default_Cisco_Logging_Template"
}

data "sdwan_cisco_omp_feature_template" "Factory_Default_Cisco_OMP_ipv46_Template" {
  name = "Factory_Default_Cisco_OMP_ipv46_Template"
}

data "sdwan_cedge_aaa_feature_template" "Factory_Default_AAA_CISCO_Template" {
  name = "Factory_Default_AAA_CISCO_Template"
}

data "sdwan_cisco_bfd_feature_template" "Factory_Default_Cisco_BFD_Template" {
  name = "Factory_Default_Cisco_BFD_Template"
}

data "sdwan_cisco_security_feature_template" "Factory_Default_Cisco_Security_Template" {
  name = "Factory_Default_Cisco_Security_Template"
}

data "sdwan_cisco_vpn_feature_template" "DC_VPN_0_cEdge_Template_v1" {
  name = "2ee186d6-6686-48ef-97f6-bb456250fb2f"
}

data "sdwan_cisco_vpn_interface_feature_template" "DC_VPN_0_MPLS_Template_v1" {
  name = "DC_VPN_0_cEdge_Template_v1"
}

data "sdwan_cisco_vpn_interface_feature_template" "DC_VPN_512_cEdge_Template_v1" {
  name = "DC_VPN_512_cEdge_Template_v1"
}

data "sdwan_cedge_global_feature_template" "Factory_Default_Global_CISCO_Template" {
  name = "Factory_Default_Global_CISCO_Template"
}
```

Now we are ready to define the resource managing the device template itself.

```hcl
resource "sdwan_feature_device_template" "TF_device_template" {
  name        = "TF_DT"
  description = "My device template via terraform. Using new and available feature templates."
  device_type = "vedge-C8000V"
  device_role = "sdwan-edge"
  general_templates = [{
    id      = sdwan_cisco_system_feature_template.TF_system_template.id
    type    = sdwan_cisco_system_feature_template.TF_system_template.template_type
    version = sdwan_cisco_system_feature_template.TF_system_template.version
    },
    { id   = data.sdwan_cisco_logging_feature_template.Factory_Default_Cisco_Logging_Template.id
      type = data.sdwan_cisco_logging_feature_template.Factory_Default_Cisco_Logging_Template.template_type
    },
    { id   = data.sdwan_cedge_aaa_feature_template.Factory_Default_AAA_CISCO_Template.id
      type = data.sdwan_cedge_aaa_feature_template.Factory_Default_AAA_CISCO_Template.template_type
    },
    { id   = data.sdwan_cisco_omp_feature_template.Factory_Default_Cisco_OMP_ipv46_Template.id
      type = data.sdwan_cisco_omp_feature_template.Factory_Default_Cisco_OMP_ipv46_Template.template_type
    },
    { id   = data.sdwan_cisco_bfd_feature_template.Factory_Default_Cisco_BFD_Template.id
      type = data.sdwan_cisco_bfd_feature_template.Factory_Default_Cisco_BFD_Template.template_type
    },
    { id   = data.sdwan_cisco_security_feature_template.Factory_Default_Cisco_Security_Template.id
      type = data.sdwan_cisco_security_feature_template.Factory_Default_Cisco_Security_Template.template_type
    },
    { id   = data.sdwan_cisco_vpn_feature_template.DC_VPN_0_cEdge_Template_v1.id
      type = data.sdwan_cisco_vpn_feature_template.DC_VPN_0_cEdge_Template_v1.template_type
      sub_templates = [{
        id   = data.sdwan_cisco_vpn_interface_feature_template.DC_VPN_0_MPLS_Template_v1.id
        type = data.sdwan_cisco_vpn_interface_feature_template.DC_VPN_0_MPLS_Template_v1.template_type
      }]
    },
    { id   = data.sdwan_cisco_vpn_interface_feature_template.DC_VPN_512_cEdge_Template_v1.id
      type = data.sdwan_cisco_vpn_interface_feature_template.DC_VPN_512_cEdge_Template_v1.template_type
    },
    { id   = data.sdwan_cedge_global_feature_template.Factory_Default_Global_CISCO_Template.id
      type = data.sdwan_cedge_global_feature_template.Factory_Default_Global_CISCO_Template.template_type
    }
  ]
}
```

Take note of the version attribute of the referenced system feature template, which is important to track changes of already deployed feature and device templates. More information is available [below](#updating-templates-and-policies).

Once the device template is defined we can attach it to one or more devices.

```hcl
resource "sdwan_attach_feature_device_template" "custom_attach" {
  id      = sdwan_feature_device_template.TF_device_template.id
  version = sdwan_feature_device_template.TF_device_template.version
  devices = [{
    id = "C8K-E2D91A30-814B-F22D-4E7B-7D39BE74EE32"
    variables = {
      var_system_ip                            = "10.10.1.11"
      var_site_id                              = "105"
      var_hostname                             = "terraform_dc-cedge01_2"
      vpn_512_next_hop_ip_address              = "10.10.20.254"
      vpn_0_next_hop_ip_address                = "10.10.23.5"
      public_internet_0_vpn_next_hop_ip_addres = "10.10.23.37"
      vpn_0_if_description                     = "GigabitEthernet2.wan-rtr01"
      vpn_0_if_name                            = "GigabitEthernet2"
      vpn_0_if_ipv4_address                    = "10.10.23.6/30"
    }
  }]
}
```

Running `terraform apply` will now create the templates and attach the device template to the device.

## Update Deployed Feature Templates

After deploying templates we can update them and a subsequent `terraform apply` will then update the templates and redeploy them accordingly. We can update the `console_baud_rate` of the previously created and deployed system feature template and run `terraform apply` again to see how this not only updates the feature template but also triggers an update of the device template and its attachment.

## Updating Templates and Policies

Configuration Groups and Policy Groups (UX 2.0) use a different mechanism for tracking changes - see [Tracking Changes](https://registry.terraform.io/providers/CiscoDevNet/sdwan/latest/docs/guides/configuration_groups#tracking-changes) in the Configuration and Policy Groups guide instead.

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

This provider supports updates to a single feature template, device template and policy per device within a single `terraform apply` operation. This is due to SD-WAN Manager locking the templates and policies once a change has been made. If for example two feature templates used in the same device template are being updated, a `terraform apply` operation will show the following warning:

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

# Define and create new system FEATURE TEMPLATE. Set up multiple variables.
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

# Read existing FEATURE TEMPLATES.
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

# Define and create new DEVICE TEMPLATE referencing the above feature templates
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

# Attach DEVICE TEMPLATE to device and provide values for multiple variables
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

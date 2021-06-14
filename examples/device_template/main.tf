terraform {
  required_providers {
    sdwan = {
      version = "0.1"
      source  = "sdwan.com/labs/sdwan"
    }
  }
}

#configure provider with your cisco sdwan credentials.
provider "sdwan" {
  # cisco-sdwan user name
  username = "user name for vManage"
  # cisco-sdwan password
  password = "password for vManage"
  # cisco-sdwan url
  url      = "https://my-cisco-vmanage.com"
  insecure = true
}

data "sdwan_device_template" "example1"{
    template_name="Smaple_for_data_Source"
}

data "sdwan_device_template" "example2"{
  template_name = "test1"
}

resource "sdwan_device_template" "example_resource1"{
  template_name = "resource_testing_1"
  template_description="For testing purpose"
  device_type="vedge-cloud"
  device_role="sdwan-edge"
  config_type="file"
  factory_default=false
  template_configuration="`"
}

resource "sdwan_device_template" "cisco_device_template" {
  template_name = "Demo-Device-Template0"
  template_description = "For demo purpose"
  device_type = "vedge-CSR-1000v"
  device_role = "sdwan-edge"
  config_type = "template"
  factory_default = false
  connection_preference = false
  connection_preference_required = false

  # Required factory default feature templates
  general_templates{
    template_id = "f4422bdc-0933-4138-b4c4-8caa7ea169f6"
    template_type = "cedge_aaa"
  }

  general_templates{
    template_id = "f9e7dee5-0db5-492d-80d6-224f893caefc"
    template_type = "cisco_bfd"
  }

  general_templates{
    template_id = "39740dbf-b641-487f-b1d9-07f6cd9ce3e5"
    template_type = "cisco_omp"
  }

  general_templates{
    template_id = "b495a21b-eca8-4d5f-97e7-19508edeba8f"
    template_type = "cisco_security"
  }

  general_templates{
    template_id = "b3625ec5-512e-4a64-b780-b16ebceb60b5"
    template_type = "cedge_global"
  }


  general_templates{
    template_id = sdwan_system_feature_template.example_cedge_system.template_id
    template_type = "cisco_system"
    sub_templates{
      sub_template_id = "c9e48a21-87c3-4211-8c67-d22f64e851c2"
      sub_template_type = "cisco_logging"
    }
    # Optional Cisco NTP feature template
    sub_templates{
      sub_template_id = sdwan_ntp_feature_template.example_cedge_ntp.template_id
      sub_template_type = "cisco_ntp"
    }
  }

  general_templates{
    template_id = sdwan_vpn_feature_template.example_cedge_vpn0.template_id
    template_type = "cisco_vpn"
    #Optional Cisco VPN Interface feature template
    sub_templates{
      sub_template_id = sdwan_vpn_interface_feature_template.example_cedge_vpn_interface.template_id
      sub_template_type = "cisco_vpn_interface"
    }
  }

  general_templates{
    template_id = sdwan_vpn_feature_template.example_cedge_vpn512.template_id
    template_type = "cisco_vpn"
    #Optional Cisco VPN Interface feature template
    sub_templates{
      sub_template_id = sdwan_vpn_interface_feature_template.example_cedge_vpn_interface.template_id
      sub_template_type = "cisco_vpn_interface"
    }
  }

  policy_id = ""

  feature_template_uid_range = []
}





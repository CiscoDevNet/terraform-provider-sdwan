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

resource "sdwan_device_template" "example_resource2" {
  template_name = "example2"
  template_description = "For testing purpose"
  device_type = "vedge-cloud"
  device_role = "sdwan-edge"
  config_type = "template"
  factory_default = false
  connection_preference = false
  connection_preference_required = false

  general_templates{
    template_id = "c8746871-cb8e-4ab7-a5f8-948c624f19ac"
    template_type = "aaa"
  }

  general_templates{
    template_id = "5083367c-0db3-4dd7-bc05-ad1afbbceaf4"
    template_type = "bfd-vedge"
  }

  general_templates{
    template_id = "79437c57-9b0f-4364-b060-55e827ac7e45"
    template_type = "omp-vedge"
  }

  general_templates{
    template_id = "486d419f-4e6c-44a5-a6fb-7b5ccf94ff90"
    template_type = "security-vedge"
  }
  
  general_templates{
    template_id = "45ea940a-45d2-4fd9-8da2-570a1a6d6874"
    template_type = "system-vedge"
    sub_templates{
      template_id = "171e9bd4-7a7b-460d-b692-83f0d5ce0124"
      template_type = "logging"
    }
    sub_templates{
      template_id = "edf3d309-91d4-45be-98d9-cfd57a05a479"
      template_type = "ntp"
    }
  }

  general_templates{
    template_id = "79437c57-9b0f-4364-b060-55e827ac7e45"
    template_type = "vpn-vedge"
    sub_templates{
      template_id = " e1b5b6e9-3b54-4279-a532-a2aaaef3e6a1"
      template_type = "vpn-vedge-interface"
    }
  }

  policy_id = "8715a21d-9367-47ea-9bc6-e25163ed9513"
  
  feature_template_uid_range = ["string", 123]

}




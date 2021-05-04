terraform {
  required_providers {
    sdwan = {
      version = "0.1"
      source  = "sdwan.com/labs/sdwan"
    }
  }
}

provider "sdwan" {
  
}

# data "sdwan_device_template" "example1"{
#     template_name="Smaple_for_data_Source"
# }

# data "sdwan_device_template" "example2"{
#   template_name = "test1"
# }

resource "sdwan_device_template" "example_resource1"{
  template_name = "resource_testing_1"
  template_description="For testing purpose"
  device_type="vedge-cloud"
  device_role="sdwan-edge"
  config_type="file"
  factory_default=false
  template_configuration="harsh"
}

resource "sdwan_device_template" "example_resource2" {
  template_name = "resource_testing_2"
  template_description="For testing purpose1"
  device_type="vedge-cloud"
  device_role="sdwan-edge"
  config_type="template"
  factory_default=false
  connection_preference = true
  connection_preference_required = true

  general_templates{
    template_id="c8746871-cb8e-4ab7-a5f8-948c624f19ac"
    template_type="system-vedge"
  }
  general_templates{
    template_id="5083367c-0db3-4dd7-bc05-ad1afbbceaf4"
    template_type="vpn-vedge"
  }

  general_templates{
    template_id="79437c57-9b0f-4364-b060-55e827ac7e45"
    template_type="vpn-vedge"
  }

  policy_id = ""
  
  feature_template_uid_range=[]

}




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

data "sdwan_device_template" "example1"{
    template_name="Smaple_for_data_Source"
}

data "sdwan_device_template" "example2"{
  template_name = "test1"
}

output "output1" {
  value = data.sdwan_device_template.example1.device_type
}

output "output2" {
  value = data.sdwan_device_template.example2
}
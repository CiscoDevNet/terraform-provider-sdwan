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

data "sdwan_device_template" "name"{
    template_name="vSmart"
}

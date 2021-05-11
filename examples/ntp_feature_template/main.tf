terraform {
  required_providers {
    sdwan = {
      version = "0.1"
      source  = "sdwan.com/labs/sdwan"
    }
  }
}

provider "sdwan" {  
  url = var.VMANAGE_URL
  username = var.VMANAGE_USERNAME
  password = var.VMANAGE_PASSWORD
}

data "sdwan_ntp_feature_template" "name"{
    template_name="Sample_NTP"
}

output "output1" {
  value = data.sdwan_ntp_feature_template.name
}

resource "sdwan_ntp_feature_template" "name1" {
  template_name = "Sample_NTP_1"
  template_description = "For testing purposes 1"
  device_type = ["vedge-1000"]
  template_min_version = "20.4.1"
  template_definition {      
    server {
      hostname = "time1.google."  
      key = 2    
      vpn = 5
      version = 3
      source_interface = "axyz"
      prefer = false
    }

    server {
      hostname = "198.00.200.100"     
      key = 1
      vpn = 0
      version = 4 
    }

    authentication {
      id = 1
      value = "12345"
    }

    authentication {
      id = 2
      value = "xyzdf"
    }

    trusted = [ "2"]
  }
  factory_default = false
}

resource "sdwan_ntp_feature_template" "name2" {
  template_name = "Sample_NTP_2"
  template_description = "Sample_NTP_2 Description changed"
  device_type = ["vbond"]
  template_min_version = "20.4.1"
  factory_default = false
}

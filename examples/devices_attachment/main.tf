terraform {
  required_providers {
    sdwan = {
      version = "0.1"
      source  = "sdwan.com/labs/sdwan"
    }
  }
}

provider "sdwan" {
  url = ""
  username = ""
  password = ""
}

    data "sdwan_devices_attachment" "name"{
      device_template_id = "a137ff27-6370-4214-b060-17d29452a297"
    }

    output "output1" {
      value = data.sdwan_devices_attachment.name
    }


    resource "sdwan_devices_attachment" "name" {
      timeout = 600
      device_template_id = "a137ff27-6370-4214-b060-17d29452a297"
      template_type = "feature"
      file = "featureInput.csv"
      devices_list {
        chassis_number = "CSR-F3EF66F0-A1E8-4C11-9C12-3C95ED91A2CC"
      }
      devices_list {
        chassis_number = "CSR-AE05722E-FD8F-41A1-A799-A9398646DE6F"
      }
      devices_list {
        chassis_number = "CSR-DC54554A-5BE7-49BD-936B-702ECD2FBE4B"
      }
    }

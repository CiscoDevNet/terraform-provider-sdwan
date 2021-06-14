terraform {
  required_providers {
    sdwan = {
      version = "0.1"
      source  = "sdwan.com/labs/sdwan"
    }
  }
}

provider "sdwan" {
  url = "https://my-cisco-vmanage.com"
  username = "user name for vManage"
  password = "password for vManage"
}

data "sdwan_devices_attachment" "name"{
  device_template_id = "a137ff27-6370-4214-b060-17d29452a297"
}

output "output1" {
  value = data.sdwan_devices_attachment.name
}

resource "sdwan_devices_attachment" "example_devices_attachment" {
  device_template_id = sdwan_device_template.cisco_device_template.template_id
  file = "Template.csv"

  devices_list{
    chassis_number = "CSR-53D3BC42-6F6C-4529-8D03-4B2903D1E688"
    }
  devices_list{
    chassis_number = "CSR-DC54554A-5BE7-49BD-936B-702ECD2FBE4B"
    }  
}







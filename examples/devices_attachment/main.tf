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

resource "sdwan_devices_attachment" "device_attachment1" {
  timeout = 100
  device_template_id = "5477b564-db5e-4195-8d76-00b6d471a8b6"
  template_type = "feature"  
  file = "Template1.csv"
  devices_list {
    chassis_number = "CSR-330E65C0-8D52-41B2-940A-398119726CCE"
  }
   devices_list {
    chassis_number = "CSR-2AD10A1D-325E-48CC-9D80-A9B5B7BBF7CB"
  }  
}


resource "sdwan_devices_attachment" "device_attachment2" {
  timeout = 100
  device_template_id = "b884a99a-e399-4e8b-8676-fb1a3ca21f7c"
  template_type = "feature"  
  file = "Template2.csv"
  devices_list {
    chassis_number = "1590f587-31aa-431c-80ac-8dc15f23c7b7"
  }
   devices_list {
    chassis_number = "f05c0fb9-56c5-4d17-bd91-067253dedeff"
  }  
}


resource "sdwan_devices_attachment" "device_attachment3" {
  timeout = 50
  device_template_id = "186f0ff1-2ea5-4a2c-b97f-d2576217b953"
  template_type = "cli"  
  file = "Template3.csv"
  devices_list {
    chassis_number = "CSR-DF5C9569-CA62-421A-97AB-F0765C1854D6"
  }
   devices_list {
    chassis_number = "CSR-F3EF66F0-A1E8-4C11-9C12-3C95ED91A2CC"
  }  
   devices_list {
    chassis_number = "CSR-AE05722E-FD8F-41A1-A799-A9398646DE6F"
  }  
}

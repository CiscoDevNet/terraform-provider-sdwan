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


resource "sdwan_devices_attachment" "name1" {
  timeout = 700
  device_template_id = "a3c15599-637f-4ee1-89c6-9c58b011c4bd"
  template_type = "cli"
  file = "Template.csv"
  devices_list {
    chassis_number = "ebdc8bd9-17e5-4eb3-a5e0-f438403a83de"
  }
  devices_list {
    chassis_number = "f21dbb35-30b3-47f4-93bb-d2b2fe092d35"
  }
}

resource "sdwan_devices_attachment" "name2" {
  timeout = 600
  device_template_id = "f202c9fe-1ef8-483d-b2a0-d8f107a8e715"
  template_type = "feature"
  file = "fi2.csv"
  devices_list {
    chassis_number = "52c7911f-c5b0-45df-b826-3155809a2a1a"
  }
}

resource "sdwan_devices_attachment" "name3" {
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

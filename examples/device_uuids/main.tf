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

data "sdwan_device_uuids" "get_devices" {
}

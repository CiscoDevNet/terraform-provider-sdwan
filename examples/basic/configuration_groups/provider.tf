terraform {
  required_providers {
    sdwan = {
      source = "CiscoDevNet/sdwan"
    }
  }
}

provider "sdwan" {
  // By default uses env $SDWAN_URL $SDWAN_USERNAME $SDWAN_PASSWORD
}

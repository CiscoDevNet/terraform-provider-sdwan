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

data "sdwan_ntp_feature_template" "test1"{
    template_name="Sample_NTP"
}

data "sdwan_ntp_feature_template" "test2"{   
    template_name="Sample_Cisco_NTP"
}
output "output1" {
  value = data.sdwan_ntp_feature_template.test1
}
output "output2" {
  value = data.sdwan_ntp_feature_template.test2
}
resource "sdwan_ntp_feature_template" "name1" {
  template_name = "Sample_Cisco_NTP"
  template_description = "For testing purposes 1"
  device_type = ["vedge-C1161-8P","vedge-C1117-4PLTEEAW",
		"vedge-C1121X-8P",
		"vedge-C8200-1N-4T",
		"vedge-ISR-4331",
		"vedge-C1127X-8PMLTEP",
		"vedge-C1117-4PMLTEEAWE",
		"vedge-ISR-4451-X",
		"vedge-C1113-8PLTEEA",
		"vedge-IR-1821",
		"vedge-ASR-1001-X",]
  template_type = "cisco_ntp"
  template_min_version = "15.0.0"
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

    master {
      enable = true
      stratum = 5
      source = "abc"
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
  template_name = "Sample_NTP"
  template_description = "For testing purposes 2"
  device_type = [	
    "vedge-1000",
		"vedge-2000",
		"vedge-cloud",
		"vedge-5000",
		]    
  template_type = "ntp"
  template_min_version = "20.4.1"
   template_definition {   
   
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
   
    trusted = [ "1"]
  }
  factory_default = false
}

resource "sdwan_ntp_feature_template" "example_cedge_ntp" {
  template_name = "Cisco-NTP-demo"
  template_description = "Minimal NTP Type cEdge Feature Template"
  template_type = "cisco_ntp"
  device_type = ["vedge-CSR-1000v"]
  template_min_version = "20.4.1"
  factory_default = false
  template_definition {
    server {
      hostname = "198.00.200.100"     
      key = 1
      vpn = 0
      version = 4 
    }
    
  }
}
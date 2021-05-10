---
layout: "sdwan"
page_title: "Provider: SDWAN"
sidebar_current: "docs-sdwan-index"
description: |-
  The Cisco SDWAN provider is used to interact with the resources provided by Cisco vManage. The provider needs to be configured with the proper credentials before it can be used.
---
  

Overview
--------------------------------------------------
Cisco SD-WAN is a cloud-delivered overlay WAN architecture connecting branches to data centers and multi-cloud environments through a single fabric. Cisco SD-WAN helps ensure predictable user experience for applications, optimizes Software-as-a-Service(SaaS), Infrastructure-as-a-Service(IaaS), and Platform-as-a-Service(PaaS) connections; and offers integrated security either on-premises or in the cloud. Analytics capabilities deliver the visibility and insights necessary to isolate and resolve issues promptly and deliver intelligent data analysis for planning and what-if scenarios.

Cisco SDWAN Provider
------------
The Cisco SDWAN provider is used to interact with the resources provided by Cisco vManage. The provider needs to be configured with the proper credentials before it can be used.

Authentication
-------------- 

Authentication with user-id and password.  
 example:  

 ```hcl

    provider "sdwan" {
      # cisco-sdwan user name
      username = "admin"
      # cisco-sdwan password
      password = "password"
      # cisco-sdwan url
      url      = "https://my-cisco-sdwan.com"
      insecure = true
    }
 
 ```

Example Usage
------------
```hcl

#configure provider with your cisco sdwan credentials.
provider "sdwan" {
  # cisco-sdwan user name
  username = "admin"
  # cisco-sdwan password
  password = "password"
  # cisco-sdwan url
  url      = "https://my-cisco-sdwan.com"
  insecure = true
}

resource "sdwan_device_template" "test" {
  template_name = "example1"
  template_description = "For testing purpose"
  device_type = "vedge-cloud"
  device_role = "sdwan-edge"
  config_type = "file"
  factory_default = false
  template_configuration = "`"
}

```

Argument Reference
------------------
Following arguments are supported with Cisco SDWAN terraform provider.

 * `username` - (Required) This is the Cisco SDWAN username, which is required to authenticate with CISCO SDWAN.
 * `password` - (Required) Password of the user mentioned in username argument. It is required when you want to use token-based authentication.
 * `url` - (Required) URL for CISCO SDWAN.
 * `insecure` - (Optional) This determines whether to use insecure HTTP connection or not. Default value is `true`.
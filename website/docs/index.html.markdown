---
layout: "sdwan"
page_title: "Provider: SDWAN"
sidebar_current: "docs-sdwan-index"
description: |-
  <Add Description for SDWAN Provider>
---
  

Overview
--------------------------------------------------
Add Overview for SDWAN Provider>

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

resource "sdwan_feature_template" "test" {
#  fabric_name = "fab1"
#  name = "MyVRF"
#  description = "This feature template is created by terraform"
}

```

Argument Reference
------------------
Following arguments are supported with Cisco SDWAN terraform provider.

 * `username` - (Required) This is the Cisco SDWAN username, which is required to authenticate with CISCO SDWAN.
 * `password` - (Required) Password of the user mentioned in username argument. It is required when you want to use token-based authentication.
 * `url` - (Required) URL for CISCO SDWAN.
 * `insecure` - (Optional) This determines whether to use insecure HTTP connection or not. Default value is `true`.
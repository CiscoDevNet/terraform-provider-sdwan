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
Example:  

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

Authentication with through proxy URL.  
Example 1 (with non-authenticated proxy server):  

 ```hcl

    provider "sdwan" {
      # cisco-sdwan user name
      username = "admin"
      # cisco-sdwan password
      password = "password"
      # cisco-sdwan url
      url      = "https://my-cisco-sdwan.com"
      insecure = true
      proxy_url = "my-proxy-url"
    }
 
 ```

 Example 2 (with authenticated proxy server):  

 ```hcl

    provider "sdwan" {
      # cisco-sdwan user name
      username = "admin"
      # cisco-sdwan password
      password = "password"
      # cisco-sdwan url
      url      = "https://my-cisco-sdwan.com"
      insecure = true
      proxy_url = "my-proxy-url"
      proxy_creds = "proxy credentials in form of username:password"
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
 * `proxy_url` - (Optional) URL for the proxy server.
 * `proxy_creds` - (Optional) credentials for the proxy server, required in case of authenticated proxy server.
 * `rate_limit` - (Optional) Rate limit of API requests per second. Allowed range is from `0` to `500`, default value is `500`

 <strong>NOTE:</strong> username, password, url, proxy_url and proxy_creds can be set through environment variables as well using environment variables SDWAN_USERNAME, SDWAN_PASSWORD, SDWAN_URL, SDWAN_PROXY_URL and SDWAN_PROXY_CREDS respectively.
---
layout: "dcnm"
page_title: "DCNM: dcnm_rest"
sidebar_current: "docs-dcnm-resource-rest"
description: |-
  Manages DCNM rest modules
---

# dcnm_rest #
Manages DCNM rest modules

## Example Usage ##

```hcl

resource "dcnm_rest" "first" {
  path    = "/rest/top-down/fabrics/fab2/networks/import"
  payload = <<EOF
  {
    "displayName": "check_rest",
    "fabric": "fab2",
    "networkExtensionTemplate": "Default_Network_Extension_Universal",
    "networkId": "30006",
    "networkName": "import",
    "networkTemplate": "Default_Network_Universal",
    "networkTemplateConfig": "{\"networkName\":\"import\",\"segmentId\":\"30006\",\"vlanId\":2300,\"mtu\":1500,\"gatewayIpAddress\":\"192.0.3.1/24\",\"gatewayIpV6Address\":\"2001:db8::1/64\",\"vlanName\":\"vlan2\",\"intfDescription\":\"second network from terraform\",\"secondaryGW1\":\"192.0.3.1/24\",\"secondaryGW2\":\"192.0.3.1/24\",\"suppressArp\":true,\"mcastGroup\":\"239.1.2.2\",\"dhcpServerAddr1\":\"1.2.3.4\",\"dhcpServerAddr2\":\"1.2.3.4\",\"vrfDhcp\":\"VRF1012\",\"loopbackId\":100,\"tag\":\"1400\",\"trmEnabled\":true,\"rtBothAuto\":true,\"enableL3OnBorder\":true}",
    "vrf": "MyVRF"
  }
 EOF 
}

```


## Argument Reference ##

* `path` - (Required) DCNM REST endpoint, where the data is being sent.
* `method` - (Optional) HTTP method. Allowed values are "GET", "PUT", "POST", "DELETE".
* `payload` - (Required) JSON encoded payload data.

NOTE: This resource will not work well in the case of Terraform destroy if there is a change in the terraform configuration required to destroy the object from the DCNM, as Destroy only has the access to the data in the state file. To destroy the objects created via dcnm_rest in such cases modify the payload and method and use the Terraform apply instead.

## Attribute Reference

No attributes are exported.

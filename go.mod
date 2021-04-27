module github.com/CiscoDevNet/terraform-provider-sdwan

go 1.16

replace github.com/CiscoDevNet/sdwan-go-client => ../sdwan-go-client

require (
	github.com/CiscoDevNet/sdwan-go-client v0.0.1
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1
)

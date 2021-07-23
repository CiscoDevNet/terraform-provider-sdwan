module github.com/CiscoDevNet/terraform-provider-sdwan

go 1.16

// replace github.com/CiscoDevNet/sdwan-go-client => ../sdwan-go-client

require (
	github.com/CiscoDevNet/sdwan-go-client v0.1.0
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1
)

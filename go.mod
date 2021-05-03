module github.com/CiscoDevNet/terraform-provider-sdwan

go 1.16

replace github.com/CiscoDevNet/sdwan-go-client => ../sdwan-go-client

require (
	github.com/CiscoDevNet/sdwan-go-client v0.0.1
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320 // indirect
	github.com/hashicorp/go-multierror v1.0.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.3.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.6.1
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/go-testing-interface v1.0.4 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
)

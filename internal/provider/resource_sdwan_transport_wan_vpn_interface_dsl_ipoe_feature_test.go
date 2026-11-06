// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/tidwall/gjson"
)

// End of section. //template:end imports

func TestAccSdwanTransportWANVPNInterfaceDSLIPoEVariableMode(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("set SDWAN_2015 or SDWAN_2018 for the matching Manager lab")
	}
	config := strings.ReplaceAll(testAccSdwanTransportWANVPNInterfaceDSLIPoEProfileParcelConfig_all(),
		`tunnel_qos_mode = "spoke"`, `tunnel_qos_mode_variable = "{{tunnel_qos_mode}}"`)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: testAccSdwanTransportWANVPNInterfaceDSLIPoEPrerequisitesProfileParcelConfig + config,
			Check: resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test",
				"tunnel_qos_mode_variable", "{{tunnel_qos_mode}}"),
		}},
	})
}

func TestAccSdwanTransportWANVPNInterfaceDSLIPoEHubMode(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("set SDWAN_2015 or SDWAN_2018 for the matching Manager lab")
	}
	config := strings.ReplaceAll(testAccSdwanTransportWANVPNInterfaceDSLIPoEProfileParcelConfig_all(),
		`tunnel_qos_mode = "spoke"`, `tunnel_qos_mode = "hub"`)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: testAccSdwanTransportWANVPNInterfaceDSLIPoEPrerequisitesProfileParcelConfig + config,
			Check: resource.ComposeTestCheckFunc(
				resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "tunnel_qos_mode", "hub"),
				resource.TestCheckNoResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "adaptive_qos"),
				resource.TestCheckNoResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "adaptive_qos_period"),
				resource.TestCheckNoResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "shaping_rate_upstream"),
				resource.TestCheckNoResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "shaping_rate_downstream"),
			),
		}},
	})
}

func TestAccSdwanTransportWANVPNInterfaceDSLIPoEPerTunnelQosVariable(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("set SDWAN_2015 or SDWAN_2018 for the matching Manager lab")
	}
	config := strings.ReplaceAll(testAccSdwanTransportWANVPNInterfaceDSLIPoEProfileParcelConfig_all(),
		`tunnel_per_tunnel_qos = true`, `tunnel_per_tunnel_qos_variable = "{{per_tunnel_qos}}"`)
	config = strings.ReplaceAll(config, `tunnel_qos_mode = "spoke"`, "")
	globalConfig := strings.TrimSuffix(config, "}\n") + "bandwidth_downstream = 5000\n}\n"
	variableConfig := strings.TrimSuffix(config, "}\n") + "bandwidth_downstream_variable = \"{{downstream}}\"\n}\n"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccSdwanTransportWANVPNInterfaceDSLIPoEPrerequisitesProfileParcelConfig + config,
				ExpectError: regexp.MustCompile("Missing Downstream Bandwidth"),
			},
			{
				Config: testAccSdwanTransportWANVPNInterfaceDSLIPoEPrerequisitesProfileParcelConfig + globalConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "bandwidth_downstream", "5000"),
					resource.TestCheckNoResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "tunnel_qos_mode"),
					resource.TestCheckNoResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "tunnel_qos_mode_variable"),
				),
			},
			{
				Config: testAccSdwanTransportWANVPNInterfaceDSLIPoEPrerequisitesProfileParcelConfig + variableConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "bandwidth_downstream_variable", "{{downstream}}"),
					resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "tunnel_per_tunnel_qos_variable", "{{per_tunnel_qos}}"),
					resource.TestCheckNoResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "tunnel_qos_mode"),
					resource.TestCheckNoResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "tunnel_qos_mode_variable"),
				),
			},
		},
	})
}

func TestTransportWANVPNInterfaceDSLIPoEVariableQosPayload(t *testing.T) {
	for _, managerVersion := range []string{"20.15.1", "20.18.1"} {
		for _, mode := range []string{"", "spoke", "hub", "variable"} {
			t.Run(managerVersion+"/"+mode, func(t *testing.T) {
				data := TransportWANVPNInterfaceDSLIPoE{
					TunnelPerTunnelQosVariable:  types.StringValue("{{qos}}"),
					BandwidthDownstreamVariable: types.StringValue("{{downstream}}"),
					TunnelBandwidthPercent:      types.Int64Value(60),
				}
				if mode == "variable" {
					data.TunnelQosModeVariable = types.StringValue("{{mode}}")
				} else if mode != "" {
					data.TunnelQosMode = types.StringValue(mode)
				}
				body := gjson.Parse(data.toBody(context.Background(), version.Must(version.NewVersion(managerVersion))))
				if body.Get("data.tunnel.mode").Exists() || body.Get("data.tunnel.bandwidthPercent").Exists() {
					t.Error("variable per-tunnel QoS must omit mode and bandwidth percentage")
				}
				if body.Get("data.tunnel.perTunnelQos.optionType").String() != "variable" || body.Get("data.bandwidthDownstream.value").String() != "{{downstream}}" {
					t.Error("per-tunnel QoS or downstream bandwidth variable was not preserved")
				}
			})
		}
	}
}

func TestTransportWANVPNInterfaceDSLIPoEAdaptiveQosMode(t *testing.T) {
	fields := []string{
		"adaptiveQoS", "adaptPeriod", "shapingRateUpstream", "shapingRateDownstream",
		"shapingRateUpstreamConfig.minShapingRateUpstream",
		"shapingRateUpstreamConfig.maxShapingRateUpstream",
		"shapingRateUpstreamConfig.defaultShapingRateUpstream",
		"shapingRateDownstreamConfig.minShapingRateDownstream",
		"shapingRateDownstreamConfig.maxShapingRateDownstream",
		"shapingRateDownstreamConfig.defaultShapingRateDownstream",
	}
	for _, managerVersion := range []string{"20.15.1", "20.18.1"} {
		for _, mode := range []string{"", "spoke", "hub", "variable"} {
			for _, option := range []string{"default", "global", "variable"} {
				t.Run(managerVersion+"/"+mode+"/"+option, func(t *testing.T) {
					data := TransportWANVPNInterfaceDSLIPoE{TunnelPerTunnelQos: types.BoolValue(true)}
					if mode == "variable" {
						data.TunnelQosModeVariable = types.StringValue("{{mode}}")
					} else if mode != "" {
						data.TunnelQosMode = types.StringValue(mode)
					}
					if option != "default" {
						data.AdaptiveQos = types.BoolValue(true)
						data.ShapingRateUpstream = types.BoolValue(true)
						data.ShapingRateDownstream = types.BoolValue(true)
					}
					if option == "global" {
						data.AdaptiveQosPeriod = types.Int64Value(20)
						data.MinShapingRateUpstream = types.Int64Value(100)
						data.MaxShapingRateUpstream = types.Int64Value(1000)
						data.DefaultShapingRateUpstream = types.Int64Value(500)
						data.MinShapingRateDownstream = types.Int64Value(200)
						data.MaxShapingRateDownstream = types.Int64Value(2000)
						data.DefaultShapingRateDownstream = types.Int64Value(1000)
					} else if option == "variable" {
						data.AdaptiveQosPeriodVariable = types.StringValue("{{period}}")
						data.MinShapingRateUpstreamVariable = types.StringValue("{{min_up}}")
						data.MaxShapingRateUpstreamVariable = types.StringValue("{{max_up}}")
						data.DefaultShapingRateUpstreamVariable = types.StringValue("{{default_up}}")
						data.MinShapingRateDownstreamVariable = types.StringValue("{{min_down}}")
						data.MaxShapingRateDownstreamVariable = types.StringValue("{{max_down}}")
						data.DefaultShapingRateDownstreamVariable = types.StringValue("{{default_down}}")
					}
					body := gjson.Parse(data.toBody(context.Background(), version.Must(version.NewVersion(managerVersion))))
					for index, field := range fields {
						value := body.Get("data.aclQos." + field)
						expected := mode != "hub" && (option != "default" || index < 4)
						if value.Exists() != expected {
							t.Errorf("%s exists = %t, want %t; value: %s", field, value.Exists(), expected, value.Raw)
						}
					}
					if mode == "hub" && (body.Get("data.aclQos.shapingRateUpstreamConfig").Exists() || body.Get("data.aclQos.shapingRateDownstreamConfig").Exists()) {
						t.Error("hub payload must not contain adaptive bandwidth configuration objects")
					}
				})
			}
		}
	}
}

func TestTransportWANVPNInterfaceDSLIPoEQosBandwidth(t *testing.T) {
	for _, managerVersion := range []string{"20.15.1", "20.18.1"} {
		for _, mode := range []string{"", "spoke", "hub", "variable"} {
			t.Run(managerVersion+"/"+mode, func(t *testing.T) {
				data := TransportWANVPNInterfaceDSLIPoE{TunnelPerTunnelQos: types.BoolValue(true)}
				if mode == "variable" {
					data.TunnelQosModeVariable = types.StringValue("{{tunnel_qos_mode}}")
				} else if mode != "" {
					data.TunnelQosMode = types.StringValue(mode)
				}
				body := gjson.Parse(data.toBody(context.Background(), version.Must(version.NewVersion(managerVersion))))
				bandwidth := body.Get("data.tunnel.bandwidthPercent")
				if mode == "hub" {
					if bandwidth.Get("optionType").String() != "default" || bandwidth.Get("value").Int() != 50 {
						t.Errorf("hub bandwidth = %s, want default 50", bandwidth.Raw)
					}
				} else if bandwidth.Exists() {
					t.Errorf("bandwidth must not be sent for mode %q: %s", mode, bandwidth.Raw)
				}
				serializedMode := body.Get("data.tunnel.mode")
				if mode == "" {
					if serializedMode.Exists() {
						t.Error("an omitted mode must remain omitted; no provider mode default")
					}
				} else if mode == "variable" {
					if serializedMode.Get("optionType").String() != "variable" || serializedMode.Get("value").String() != "{{tunnel_qos_mode}}" {
						t.Errorf("variable mode was not preserved: %s", serializedMode.Raw)
					}
					for _, option := range []string{"global", "variable"} {
						if option == "global" {
							data.TunnelBandwidthPercent = types.Int64Value(60)
						} else {
							data.TunnelBandwidthPercent = types.Int64Null()
							data.TunnelBandwidthPercentVariable = types.StringValue("{{bandwidth}}")
						}
						configuredBody := gjson.Parse(data.toBody(context.Background(), version.Must(version.NewVersion(managerVersion))))
						configuredBandwidth := configuredBody.Get("data.tunnel.bandwidthPercent")
						if configuredBandwidth.Get("optionType").String() != option {
							t.Errorf("explicit bandwidth option = %s, want %s", configuredBandwidth.Raw, option)
						}
					}
				} else if serializedMode.Get("optionType").String() != "global" || serializedMode.Get("value").String() != mode {
					t.Errorf("explicit mode was not preserved: %s", serializedMode.Raw)
				}
			})
		}
	}
}

func TestTransportWANVPNInterfaceDSLIPoEVersionedPayload(t *testing.T) {
	for _, managerVersion := range []string{"20.15.1", "20.18.1"} {
		for _, option := range []string{"default", "global", "variable"} {
			t.Run(managerVersion+"/"+option, func(t *testing.T) {
				data := TransportWANVPNInterfaceDSLIPoE{
					Ipv4AddressType: types.StringValue("dynamic"),
				}
				if option == "global" {
					data.TunnelInterfaceColorDescription = types.StringValue("DSL WAN")
					data.TunnelInterfaceFullPortHop = types.BoolValue(true)
					data.AclIpv4EgressFeatureId = types.StringValue("11111111-1111-1111-1111-111111111111")
					data.AclIpv4IngressFeatureId = data.AclIpv4EgressFeatureId
					data.AclIpv6EgressFeatureId = data.AclIpv4EgressFeatureId
					data.AclIpv6IngressFeatureId = data.AclIpv4EgressFeatureId
				} else if option == "variable" {
					data.TunnelInterfaceColorDescriptionVariable = types.StringValue("{{description}}")
					data.TunnelInterfaceFullPortHopVariable = types.StringValue("{{port_hop}}")
					data.AclIpv4EgressFeatureIdVariable = types.StringValue("{{acl}}")
					data.AclIpv4IngressFeatureIdVariable = data.AclIpv4EgressFeatureIdVariable
					data.AclIpv6EgressFeatureIdVariable = data.AclIpv4EgressFeatureIdVariable
					data.AclIpv6IngressFeatureIdVariable = data.AclIpv4EgressFeatureIdVariable
				}
				body := gjson.Parse(data.toBody(context.Background(), version.Must(version.NewVersion(managerVersion))))
				for _, field := range []string{"colorDescription", "fullPortHop"} {
					value := body.Get("data.tunnel." + field)
					if managerVersion == "20.15.1" {
						if value.Exists() {
							t.Errorf("20.18-only field %s was sent to 20.15: %s", field, value.Raw)
						}
					} else if value.Get("optionType").String() != option {
						t.Errorf("%s optionType = %q, want %q", field, value.Get("optionType").String(), option)
					}
				}
				for _, field := range []string{"ipv4AclEgress", "ipv4AclIngress", "ipv6AclEgress", "ipv6AclIngress"} {
					value := body.Get("data.aclQos." + field)
					if option == "default" || (option == "variable" && managerVersion == "20.15.1") {
						if value.Exists() {
							t.Errorf("unsupported or unset ACL %s was serialized: %s", field, value.Raw)
						}
					} else if value.Get("refId.optionType").String() != option {
						t.Errorf("%s refId.optionType = %q, want %q", field, value.Get("refId.optionType").String(), option)
					}
				}
				distance := body.Get("data.ethernet.intfIpAddress.dynamic.dynamicDhcpDistance")
				if distance.Get("optionType").String() != "default" || distance.Get("value").Int() != 1 {
					t.Errorf("omitted DHCP distance did not serialize as default 1: %s", distance.Raw)
				}
			})
		}
	}
}

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanTransportWANVPNInterfaceDSLIPoEProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "shutdown", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "vdsl_slot", "0/0/0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "interface_name", "GigabitEthernet1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "ipv4_address", "192.0.2.10"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "ipv4_subnet_mask", "255.255.255.0"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "tunnel_interface_enabled", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "tunnel_per_tunnel_qos", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "tunnel_qos_mode", "spoke"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "tunnel_interface_color", "default"))
	if os.Getenv("SDWAN_2018") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "tunnel_interface_color_description", "DSL WAN"))
	}
	if os.Getenv("SDWAN_2018") != "" {
		checks = append(checks, resource.TestCheckResourceAttr("sdwan_transport_wan_vpn_interface_dsl_ipoe_feature.test", "tunnel_interface_full_port_hop", "false"))
	}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTransportWANVPNInterfaceDSLIPoEPrerequisitesProfileParcelConfig + testAccSdwanTransportWANVPNInterfaceDSLIPoEProfileParcelConfig_minimum(),
			},
			{
				Config: testAccSdwanTransportWANVPNInterfaceDSLIPoEPrerequisitesProfileParcelConfig + testAccSdwanTransportWANVPNInterfaceDSLIPoEProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanTransportWANVPNInterfaceDSLIPoEPrerequisitesProfileParcelConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_transport_wan_vpn_feature" "test" {
  name               = "TF_TEST_WAN"
  description        = "Terraform test"
  feature_profile_id = sdwan_transport_feature_profile.test.id
  vpn                = 0
}
`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum
func testAccSdwanTransportWANVPNInterfaceDSLIPoEProfileParcelConfig_minimum() string {
	config := `resource "sdwan_transport_wan_vpn_interface_dsl_ipoe_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MIN"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id` + "\n"
	config += `	vdsl_slot = "0/0/0"` + "\n"
	config += `	interface_name = "GigabitEthernet1"` + "\n"
	config += `	ipv4_address_type = "dynamic"` + "\n"
	config += `	tunnel_interface_enabled = true` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTransportWANVPNInterfaceDSLIPoEProfileParcelConfig_all() string {
	config := `resource "sdwan_transport_wan_vpn_interface_dsl_ipoe_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.test.id` + "\n"
	config += `	shutdown = true` + "\n"
	config += `	vdsl_slot = "0/0/0"` + "\n"
	config += `	interface_name = "GigabitEthernet1"` + "\n"
	config += `	ipv4_address_type = "static"` + "\n"
	config += `	ipv4_address = "192.0.2.10"` + "\n"
	config += `	ipv4_subnet_mask = "255.255.255.0"` + "\n"
	config += `	tunnel_interface_enabled = true` + "\n"
	config += `	tunnel_per_tunnel_qos = true` + "\n"
	config += `	tunnel_qos_mode = "spoke"` + "\n"
	config += `	tunnel_interface_color = "default"` + "\n"
	if os.Getenv("SDWAN_2018") != "" {
		config += `	tunnel_interface_color_description = "DSL WAN"` + "\n"
	}
	if os.Getenv("SDWAN_2018") != "" {
		config += `	tunnel_interface_full_port_hop = false` + "\n"
	}
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

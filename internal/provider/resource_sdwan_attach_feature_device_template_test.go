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

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanAttachFeatureDeviceTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanAttachFeatureDeviceTemplateConfig_all("1.1.1.1"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_attach_feature_device_template.test", "devices.0.variables.system_site_id", "1001"),
					resource.TestCheckResourceAttr("sdwan_attach_feature_device_template.test", "devices.0.variables.system_system_ip", "1.1.1.1"),
					resource.TestCheckResourceAttr("sdwan_attach_feature_device_template.test", "devices.0.variables.system_host_name", "router1"),
					resource.TestCheckResourceAttr("sdwan_attach_feature_device_template.test", "devices.0.variables.vpn_if_name_Default_vEdge_DHCP_Tunnel_Interface", "GigabitEthernet1"),
				),
			},
			{
				Config: testAccSdwanAttachFeatureDeviceTemplateConfig_all("1.1.1.2"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_attach_feature_device_template.test", "devices.0.variables.system_system_ip", "1.1.1.2"),
				),
			},
		},
	})
}

func testAccSdwanAttachFeatureDeviceTemplateConfig_all(ip string) string {
	return fmt.Sprintf(`
		data "sdwan_cisco_system_feature_template" "Factory_Default_Cisco_System_Template" {
			name = "Factory_Default_Cisco_System_Template"
		}
		
		data "sdwan_cisco_logging_feature_template" "Factory_Default_Cisco_Logging_Template" {
			name = "Factory_Default_Cisco_Logging_Template"
		}
		
		data "sdwan_cisco_omp_feature_template" "Factory_Default_Cisco_OMP_ipv46_Template" {
			name = "Factory_Default_Cisco_OMP_ipv46_Template"
		}
		
		data "sdwan_cisco_bfd_feature_template" "Factory_Default_Cisco_BFD_Template" {
			name = "Factory_Default_Cisco_BFD_Template"
		}
		
		data "sdwan_cisco_security_feature_template" "Factory_Default_Cisco_Security_Template" {
			name = "Factory_Default_Cisco_Security_Template"
		}
		
		data "sdwan_cisco_vpn_feature_template" "Factory_Default_Cisco_VPN_0_Template" {
			name = "Factory_Default_Cisco_VPN_0_Template"
		}
		
		data "sdwan_cisco_vpn_interface_feature_template" "Factory_Default_Cisco_DHCP_Tunnel_Interface" {
			name = "Factory_Default_Cisco_DHCP_Tunnel_Interface"
		}
		
		data "sdwan_cisco_vpn_feature_template" "Factory_Default_Cisco_VPN_512_Template" {
			name = "Factory_Default_Cisco_VPN_512_Template"
		}
		
		data "sdwan_cedge_global_feature_template" "Factory_Default_Global_CISCO_Template" {
			name = "Factory_Default_Global_CISCO_Template"
		}

		resource "sdwan_feature_device_template" "TF_device_template" {
			name        = "TF_TEST"
			description = "Terraform test."
			device_type = "vedge-C8000V"
			device_role = "sdwan-edge"
			general_templates = [
				{
					id   = data.sdwan_cisco_system_feature_template.Factory_Default_Cisco_System_Template.id
					type = data.sdwan_cisco_system_feature_template.Factory_Default_Cisco_System_Template.template_type
				},
				{
					id   = data.sdwan_cisco_logging_feature_template.Factory_Default_Cisco_Logging_Template.id
					type = data.sdwan_cisco_logging_feature_template.Factory_Default_Cisco_Logging_Template.template_type
				},
				{
					id   = data.sdwan_cisco_omp_feature_template.Factory_Default_Cisco_OMP_ipv46_Template.id
					type = data.sdwan_cisco_omp_feature_template.Factory_Default_Cisco_OMP_ipv46_Template.template_type
				},
				{
					id   = data.sdwan_cisco_bfd_feature_template.Factory_Default_Cisco_BFD_Template.id
					type = data.sdwan_cisco_bfd_feature_template.Factory_Default_Cisco_BFD_Template.template_type
				},
				{
					id   = data.sdwan_cisco_security_feature_template.Factory_Default_Cisco_Security_Template.id
					type = data.sdwan_cisco_security_feature_template.Factory_Default_Cisco_Security_Template.template_type
				},
				{
					id   = data.sdwan_cisco_vpn_feature_template.Factory_Default_Cisco_VPN_0_Template.id
					type = data.sdwan_cisco_vpn_feature_template.Factory_Default_Cisco_VPN_0_Template.template_type
					sub_templates = [{
						id   = data.sdwan_cisco_vpn_interface_feature_template.Factory_Default_Cisco_DHCP_Tunnel_Interface.id
						type = data.sdwan_cisco_vpn_interface_feature_template.Factory_Default_Cisco_DHCP_Tunnel_Interface.template_type
					}]
				},
				{
					id   = data.sdwan_cisco_vpn_feature_template.Factory_Default_Cisco_VPN_512_Template.id
					type = data.sdwan_cisco_vpn_feature_template.Factory_Default_Cisco_VPN_512_Template.template_type
				},
				{
					id   = data.sdwan_cedge_global_feature_template.Factory_Default_Global_CISCO_Template.id
					type = data.sdwan_cedge_global_feature_template.Factory_Default_Global_CISCO_Template.template_type
				}
			]
		}

		resource "sdwan_attach_feature_device_template" "test" {
			id      = sdwan_feature_device_template.TF_device_template.id
			devices = [{
				id = "C8K-CC678D1C-8EDF-3966-4F51-ABFAB64F5ABE"
				variables = {
					system_system_ip                                = "%s"
					system_site_id                                  = "1001"
					system_host_name                                = "router1"
					vpn_if_name_Default_vEdge_DHCP_Tunnel_Interface = "GigabitEthernet1"
				}
			}]
		}`, ip)
}

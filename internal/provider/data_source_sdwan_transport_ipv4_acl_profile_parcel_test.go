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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwanTransportIPv4ACLProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2012") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2012")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_ipv4_acl_profile_parcel.test", "default_action", "drop"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_ipv4_acl_profile_parcel.test", "sequences.0.sequence_id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_ipv4_acl_profile_parcel.test", "sequences.0.sequence_name", "AccessControlList1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_ipv4_acl_profile_parcel.test", "sequences.0.conditions.0.packet_length", "1500"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_ipv4_acl_profile_parcel.test", "sequences.0.conditions.0.source_ports.0.port", "8000"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_ipv4_acl_profile_parcel.test", "sequences.0.conditions.0.tcp_state", "syn"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_ipv4_acl_profile_parcel.test", "sequences.0.actions.0.accept_dscp", "60"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_ipv4_acl_profile_parcel.test", "sequences.0.actions.0.accept_counter_name", "COUNTER_1"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_ipv4_acl_profile_parcel.test", "sequences.0.actions.0.accept_log", "false"))
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_transport_ipv4_acl_profile_parcel.test", "sequences.0.actions.0.accept_next_hop", "1.2.3.4"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanTransportIPv4ACLPrerequisitesProfileParcelConfig + testAccDataSourceSdwanTransportIPv4ACLProfileParcelConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccDataSourceSdwanTransportIPv4ACLPrerequisitesProfileParcelConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwanTransportIPv4ACLProfileParcelConfig() string {
	config := `resource "sdwan_transport_ipv4_acl_profile_parcel" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	default_action = "drop"` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  sequence_id = 1` + "\n"
	config += `	  sequence_name = "AccessControlList1"` + "\n"
	config += `	  conditions = [{` + "\n"
	config += `		dscps = [16]` + "\n"
	config += `		packet_length = 1500` + "\n"
	config += `		protocols = [1]` + "\n"
	config += `      source_ports = [{` + "\n"
	config += `			port = 8000` + "\n"
	config += `		}]` + "\n"
	config += `		tcp_state = "syn"` + "\n"
	config += `	}]` + "\n"
	config += `	  actions = [{` + "\n"
	config += `		accept_dscp = 60` + "\n"
	config += `		accept_counter_name = "COUNTER_1"` + "\n"
	config += `		accept_log = false` + "\n"
	config += `		accept_next_hop = "1.2.3.4"` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "sdwan_transport_ipv4_acl_profile_parcel" "test" {
			id = sdwan_transport_ipv4_acl_profile_parcel.test.id
			feature_profile_id = sdwan_transport_feature_profile.test.id
		}
	`
	return config
}

// End of section. //template:end testAccDataSourceConfig

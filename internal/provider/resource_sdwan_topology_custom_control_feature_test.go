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

// Section below is generated&owned by "gen/generator.go". //template:begin testAcc
func TestAccSdwanTopologyCustomControlProfileParcel(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "default_action", "reject"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "target_level", "SITE"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.name", "Rule1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.base_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.type", "route"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.ip_type", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.match_entries.0.omp_tag", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.match_entries.0.origin", "connected"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.match_entries.0.originator", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.match_entries.0.tloc_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.match_entries.0.tloc_color", "bronze"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.match_entries.0.tloc_encapsulation", "ipsec"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.action_entries.0.set_parameters.0.preference", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.action_entries.0.set_parameters.0.omp_tag", "100"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{

			{
				Config: testAccSdwanTopologyCustomControlPrerequisitesProfileParcelConfig + testAccSdwanTopologyCustomControlProfileParcelConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanTopologyCustomControlPrerequisitesProfileParcelConfig = `
resource "sdwan_topology_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigMinimum

// End of section. //template:end testAccConfigMinimum

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTopologyCustomControlProfileParcelConfig_all() string {
	config := `resource "sdwan_topology_custom_control_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_ALL"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	config += `	feature_profile_id = sdwan_topology_feature_profile.test.id` + "\n"
	config += `	default_action = "reject"` + "\n"
	config += `	target_level = "SITE"` + "\n"
	config += `	target_inbound_sites = ["SITE_100"]` + "\n"
	config += `	target_outbound_sites = ["SITE_200"]` + "\n"
	config += `	sequences = [{` + "\n"
	config += `	  id = 1` + "\n"
	config += `	  name = "Rule1"` + "\n"
	config += `	  base_action = "accept"` + "\n"
	config += `	  type = "route"` + "\n"
	config += `	  ip_type = "ipv4"` + "\n"
	config += `	  match_entries = [{` + "\n"
	config += `		omp_tag = 100` + "\n"
	config += `		origin = "connected"` + "\n"
	config += `		originator = "1.2.3.4"` + "\n"
	config += `		tloc_ip = "1.2.3.4"` + "\n"
	config += `		tloc_color = "bronze"` + "\n"
	config += `		tloc_encapsulation = "ipsec"` + "\n"
	config += `	}]` + "\n"
	config += `	  action_entries = [{` + "\n"
	config += `      set_parameters = [{` + "\n"
	config += `			preference = 100` + "\n"
	config += `			omp_tag = 100` + "\n"
	config += `		}]` + "\n"
	config += `	}]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

func TestAccSdwanTopologyCustomControlProfileParcelMultiSequence(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "default_action", "reject"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "target_level", "SITE"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.#", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.id", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.name", "Rule1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.base_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.type", "route"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.ip_type", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.match_entries.0.omp_tag", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.match_entries.0.origin", "connected"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.match_entries.0.originator", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.match_entries.0.tloc_ip", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.match_entries.0.tloc_color", "bronze"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.match_entries.0.tloc_encapsulation", "ipsec"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.0.action_entries.0.set_parameters.0.preference", "100"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.1.id", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.1.name", "Rule2"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.1.base_action", "reject"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.1.type", "tloc"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.1.ip_type", "ipv4"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.1.match_entries.0.omp_tag", "200"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.2.id", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.2.name", "Rule3"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.2.base_action", "accept"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.2.type", "route"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.2.ip_type", "ipv6"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_topology_custom_control_feature.test", "sequences.2.action_entries.0.set_parameters.0.omp_tag", "300"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTopologyCustomControlPrerequisitesProfileParcelConfig + testAccSdwanTopologyCustomControlProfileParcelConfig_multiSequence(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

func testAccSdwanTopologyCustomControlProfileParcelConfig_multiSequence() string {
	config := `resource "sdwan_topology_custom_control_feature" "test" {` + "\n"
	config += ` name = "TF_TEST_MULTI"` + "\n"
	config += ` description = "Terraform integration test with multiple sequences"` + "\n"
	config += `	feature_profile_id = sdwan_topology_feature_profile.test.id` + "\n"
	config += `	default_action = "reject"` + "\n"
	config += `	target_level = "SITE"` + "\n"
	config += `	target_outbound_sites = ["SITE_100"]` + "\n"
	config += `	sequences = [` + "\n"
	config += `	  {` + "\n"
	config += `	    id = 1` + "\n"
	config += `	    name = "Rule1"` + "\n"
	config += `	    base_action = "accept"` + "\n"
	config += `	    type = "route"` + "\n"
	config += `	    ip_type = "ipv4"` + "\n"
	config += `	    match_entries = [{` + "\n"
	config += `		  omp_tag = 100` + "\n"
	config += `		  origin = "connected"` + "\n"
	config += `		  originator = "1.2.3.4"` + "\n"
	config += `		  tloc_ip = "1.2.3.4"` + "\n"
	config += `		  tloc_color = "bronze"` + "\n"
	config += `		  tloc_encapsulation = "ipsec"` + "\n"
	config += `	    }]` + "\n"
	config += `	    action_entries = [{` + "\n"
	config += `        set_parameters = [{` + "\n"
	config += `		    preference = 100` + "\n"
	config += `		  }]` + "\n"
	config += `	    }]` + "\n"
	config += `	  },` + "\n"
	config += `	  {` + "\n"
	config += `	    id = 2` + "\n"
	config += `	    name = "Rule2"` + "\n"
	config += `	    base_action = "reject"` + "\n"
	config += `	    type = "tloc"` + "\n"
	config += `	    ip_type = "ipv4"` + "\n"
	config += `	    match_entries = [{` + "\n"
	config += `		  omp_tag = 200` + "\n"
	config += `	    }]` + "\n"
	config += `	  },` + "\n"
	config += `	  {` + "\n"
	config += `	    id = 3` + "\n"
	config += `	    name = "Rule3"` + "\n"
	config += `	    base_action = "accept"` + "\n"
	config += `	    type = "route"` + "\n"
	config += `	    ip_type = "ipv6"` + "\n"
	config += `	    action_entries = [{` + "\n"
	config += `        set_parameters = [{` + "\n"
	config += `		    omp_tag = 300` + "\n"
	config += `		  }]` + "\n"
	config += `	    }]` + "\n"
	config += `	  }` + "\n"
	config += `	]` + "\n"
	config += `}` + "\n"
	return config
}

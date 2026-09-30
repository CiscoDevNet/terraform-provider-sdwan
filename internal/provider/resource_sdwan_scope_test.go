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
func TestAccSdwanScope(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_scope.test", "name", "west1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_scope.test", "description", "West Coast Domain1"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_scope.test", "objects.0.object_type", "feature-profile"))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanScopePrerequisitesConfig + testAccSdwanScopeConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanScopePrerequisitesConfig = `
resource "sdwan_policy_object_feature_profile" "test" {
  name = "POLICY_OBJECT_FP_1"
  description = "My policy object feature profile 1"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanScopeConfig_all() string {
	config := `resource "sdwan_scope" "test" {` + "\n"
	config += `	name = "west1"` + "\n"
	config += `	description = "West Coast Domain1"` + "\n"
	config += `	objects = [{` + "\n"
	config += `	  object_type = "feature-profile"` + "\n"
	config += `	  object_ids = [sdwan_policy_object_feature_profile.test.id]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

const testAccSdwanScopeMultiPrerequisitesConfig = `
resource "sdwan_cli_device_template" "test_multi_dt1" {
  name              = "TF_SCOPE_MULTI_CLI_DT_1"
  description       = "Scope multi-object test CLI device template 1"
  device_type       = "vedge-ISR-4331"
  cli_type          = "device"
  cli_configuration = " system\n host-name             R1-ISR4331-1200-1"
}

resource "sdwan_cli_device_template" "test_multi_dt2" {
  name              = "TF_SCOPE_MULTI_CLI_DT_2"
  description       = "Scope multi-object test CLI device template 2"
  device_type       = "vedge-ISR-4331"
  cli_type          = "device"
  cli_configuration = " system\n host-name             R2-ISR4331-1200-2"
}

resource "sdwan_system_feature_profile" "test_multi_sysfp" {
  name        = "TF_SCOPE_MULTI_SYSTEM_FP_1"
  description = "Scope multi-object test system feature profile 1"
}

`

func testAccSdwanScopeMultiConfig() string {
	config := `resource "sdwan_scope" "test_multi" {` + "\n"
	config += `	name = "tf-acc-scope-multi"` + "\n"
	config += `	description = "Scope with multiple object groups"` + "\n"
	config += `	objects = [{` + "\n"
	config += `	  object_type = "device-template"` + "\n"
	config += `	  object_ids = [sdwan_cli_device_template.test_multi_dt1.id, sdwan_cli_device_template.test_multi_dt2.id]` + "\n"
	config += `	}, {` + "\n"
	config += `	  object_type = "feature-profile"` + "\n"
	config += `	  object_ids = [sdwan_system_feature_profile.test_multi_sysfp.id]` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"
	return config
}

func TestAccSdwanScopeMultipleObjects(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}

	config := testAccSdwanScopeMultiPrerequisitesConfig + testAccSdwanScopeMultiConfig()

	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_scope.test_multi", "name", "tf-acc-scope-multi"))
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_scope.test_multi", "description", "Scope with multiple object groups"))
	// Exactly the two groups declared in config - no controller-injected extras.
	checks = append(checks, resource.TestCheckResourceAttr("sdwan_scope.test_multi", "objects.#", "2"))
	// `objects` is a SetNestedAttribute, so element indices are not stable;
	// assert set membership rather than positional objects.0.* / objects.1.* paths.
	checks = append(checks, resource.TestCheckTypeSetElemNestedAttrs("sdwan_scope.test_multi", "objects.*", map[string]string{
		"object_type":  "device-template",
		"object_ids.#": "2",
	}))
	checks = append(checks, resource.TestCheckTypeSetElemNestedAttrs("sdwan_scope.test_multi", "objects.*", map[string]string{
		"object_type":  "feature-profile",
		"object_ids.#": "1",
	}))

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
			// Drift guard. Re-running the identical config with PlanOnly forces a
			// refresh (a real GET against the Manager) followed by a re-plan. If the
			// read path adopted the controller-injected feature-template
			// associations into state, or dropped the device-template group because
			// its ids arrive under the "templateId" key, the refreshed state will no
			// longer match the config and the plan will be non-empty - failing here.
			{
				Config:             config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			// Import verification. ImportStateVerify is deliberately false: on import
			// there is no prior practitioner state to reconcile against, so the
			// provider intentionally adopts the full server-side view of the scope,
			// including the controller-injected associations. That imported state is
			// therefore expected NOT to equal the state produced from the config
			// above, and a strict attribute-by-attribute comparison would fail by
			// design.
			{
				ResourceName:      "sdwan_scope.test_multi",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

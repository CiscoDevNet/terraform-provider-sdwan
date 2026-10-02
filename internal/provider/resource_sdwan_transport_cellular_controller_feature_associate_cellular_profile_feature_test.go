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
func TestAccSdwanTransportCellularControllerFeatureAssociateCellularProfileFeature(t *testing.T) {
	if os.Getenv("SDWAN_2015") == "" && os.Getenv("SDWAN_2018") == "" {
		t.Skip("skipping test, set environment variable SDWAN_2015 or SDWAN_2018")
	}
	var checks []resource.TestCheckFunc
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanTransportCellularControllerFeatureAssociateCellularProfileFeaturePrerequisitesConfig + testAccSdwanTransportCellularControllerFeatureAssociateCellularProfileFeatureConfig_all(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

// End of section. //template:end testAcc

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
const testAccSdwanTransportCellularControllerFeatureAssociateCellularProfileFeaturePrerequisitesConfig = `
resource "sdwan_transport_feature_profile" "test" {
  name        = "TF_TEST"
  description = "Terraform test"
}

resource "sdwan_transport_cellular_controller_feature" "test" {
  name               = "TF_TEST_CELLULAR_CONTROLLER"
  description        = "Terraform test"
  feature_profile_id = sdwan_transport_feature_profile.test.id
  cellular_id        = "0/3/0"
}

resource "sdwan_transport_cellular_profile_feature" "test" {
  name               = "TF_TEST_CELLULAR_PROFILE"
  description        = "Terraform test"
  feature_profile_id = sdwan_transport_feature_profile.test.id
  profile_id         = 2
  access_point_name  = "apn1"
}

`

// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccConfigAll
func testAccSdwanTransportCellularControllerFeatureAssociateCellularProfileFeatureConfig_all() string {
	config := `resource "sdwan_transport_cellular_controller_feature_associate_cellular_profile_feature" "test" {` + "\n"
	config += `	feature_profile_id = sdwan_transport_feature_profile.test.id` + "\n"
	config += `	transport_cellular_controller_feature_id = sdwan_transport_cellular_controller_feature.test.id` + "\n"
	config += `	transport_cellular_profile_feature_id = sdwan_transport_cellular_profile_feature.test.id` + "\n"
	config += `}` + "\n"
	return config
}

// End of section. //template:end testAccConfigAll

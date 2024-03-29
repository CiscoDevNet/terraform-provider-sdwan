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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanSystemOMPProfileParcel(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanSystemOMPProfileParcelConfig_minimum(),
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: testAccSdwanSystemOMPProfileParcelConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "graceful_restart", "true"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "overlay_as", "10"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "paths_advertised_per_prefix", "4"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "ecmp_limit", "4"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "shutdown", "false"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "omp_admin_distance_ipv4", "10"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "omp_admin_distance_ipv6", "20"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertisement_interval", "1"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "graceful_restart_timer", "43200"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "eor_timer", "300"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "holdtime", "60"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv4_bgp", "false"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv4_ospf", "false"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv4_ospf_v3", "false"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv4_cpnnected", "false"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv4_static", "false"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv4_eigrp", "false"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv4_lisp", "false"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv4_isis", "false"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv6_bgp", "true"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv6_ospf", "true"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv6_connected", "true"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv6_static", "true"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv6_eigrp", "true"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv6_lisp", "true"),
					resource.TestCheckResourceAttr("sdwan_system_omp_profile_parcel.test", "advertise_ipv6_isis", "true"),
				),
			},
		},
	})
}

func testAccSdwanSystemOMPProfileParcelConfig_minimum() string {
	return `
	resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

	resource "sdwan_system_omp_profile_parcel" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		feature_profile_id = sdwan_system_feature_profile.test.id
	}
	`
}

func testAccSdwanSystemOMPProfileParcelConfig_all() string {
	return `
	resource "sdwan_system_feature_profile" "test" {
  name = "TF_TEST"
  description = "Terraform test"
}

	resource "sdwan_system_omp_profile_parcel" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		feature_profile_id = sdwan_system_feature_profile.test.id
		graceful_restart = true
		overlay_as = 10
		paths_advertised_per_prefix = 4
		ecmp_limit = 4
		shutdown = false
		omp_admin_distance_ipv4 = 10
		omp_admin_distance_ipv6 = 20
		advertisement_interval = 1
		graceful_restart_timer = 43200
		eor_timer = 300
		holdtime = 60
		advertise_ipv4_bgp = false
		advertise_ipv4_ospf = false
		advertise_ipv4_ospf_v3 = false
		advertise_ipv4_cpnnected = false
		advertise_ipv4_static = false
		advertise_ipv4_eigrp = false
		advertise_ipv4_lisp = false
		advertise_ipv4_isis = false
		advertise_ipv6_bgp = true
		advertise_ipv6_ospf = true
		advertise_ipv6_connected = true
		advertise_ipv6_static = true
		advertise_ipv6_eigrp = true
		advertise_ipv6_lisp = true
		advertise_ipv6_isis = true
	}
	`
}

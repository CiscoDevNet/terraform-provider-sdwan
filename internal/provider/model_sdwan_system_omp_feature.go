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
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type SystemOMP struct {
	Id                               types.String `tfsdk:"id"`
	Version                          types.Int64  `tfsdk:"version"`
	Name                             types.String `tfsdk:"name"`
	Description                      types.String `tfsdk:"description"`
	FeatureProfileId                 types.String `tfsdk:"feature_profile_id"`
	GracefulRestart                  types.Bool   `tfsdk:"graceful_restart"`
	GracefulRestartVariable          types.String `tfsdk:"graceful_restart_variable"`
	OverlayAs                        types.Int64  `tfsdk:"overlay_as"`
	OverlayAsVariable                types.String `tfsdk:"overlay_as_variable"`
	PathsAdvertisedPerPrefix         types.Int64  `tfsdk:"paths_advertised_per_prefix"`
	PathsAdvertisedPerPrefixVariable types.String `tfsdk:"paths_advertised_per_prefix_variable"`
	EcmpLimit                        types.Int64  `tfsdk:"ecmp_limit"`
	EcmpLimitVariable                types.String `tfsdk:"ecmp_limit_variable"`
	Shutdown                         types.Bool   `tfsdk:"shutdown"`
	ShutdownVariable                 types.String `tfsdk:"shutdown_variable"`
	OmpAdminDistanceIpv4             types.Int64  `tfsdk:"omp_admin_distance_ipv4"`
	OmpAdminDistanceIpv4Variable     types.String `tfsdk:"omp_admin_distance_ipv4_variable"`
	OmpAdminDistanceIpv6             types.Int64  `tfsdk:"omp_admin_distance_ipv6"`
	OmpAdminDistanceIpv6Variable     types.String `tfsdk:"omp_admin_distance_ipv6_variable"`
	AdvertisementInterval            types.Int64  `tfsdk:"advertisement_interval"`
	AdvertisementIntervalVariable    types.String `tfsdk:"advertisement_interval_variable"`
	GracefulRestartTimer             types.Int64  `tfsdk:"graceful_restart_timer"`
	GracefulRestartTimerVariable     types.String `tfsdk:"graceful_restart_timer_variable"`
	EorTimer                         types.Int64  `tfsdk:"eor_timer"`
	EorTimerVariable                 types.String `tfsdk:"eor_timer_variable"`
	Holdtime                         types.Int64  `tfsdk:"holdtime"`
	HoldtimeVariable                 types.String `tfsdk:"holdtime_variable"`
	AdvertiseIpv4Bgp                 types.Bool   `tfsdk:"advertise_ipv4_bgp"`
	AdvertiseIpv4BgpVariable         types.String `tfsdk:"advertise_ipv4_bgp_variable"`
	AdvertiseIpv4Ospf                types.Bool   `tfsdk:"advertise_ipv4_ospf"`
	AdvertiseIpv4OspfVariable        types.String `tfsdk:"advertise_ipv4_ospf_variable"`
	AdvertiseIpv4OspfV3              types.Bool   `tfsdk:"advertise_ipv4_ospf_v3"`
	AdvertiseIpv4OspfV3Variable      types.String `tfsdk:"advertise_ipv4_ospf_v3_variable"`
	AdvertiseIpv4Connected           types.Bool   `tfsdk:"advertise_ipv4_connected"`
	AdvertiseIpv4ConnectedVariable   types.String `tfsdk:"advertise_ipv4_connected_variable"`
	AdvertiseIpv4Static              types.Bool   `tfsdk:"advertise_ipv4_static"`
	AdvertiseIpv4StaticVariable      types.String `tfsdk:"advertise_ipv4_static_variable"`
	AdvertiseIpv4Eigrp               types.Bool   `tfsdk:"advertise_ipv4_eigrp"`
	AdvertiseIpv4EigrpVariable       types.String `tfsdk:"advertise_ipv4_eigrp_variable"`
	AdvertiseIpv4Lisp                types.Bool   `tfsdk:"advertise_ipv4_lisp"`
	AdvertiseIpv4LispVariable        types.String `tfsdk:"advertise_ipv4_lisp_variable"`
	AdvertiseIpv4Isis                types.Bool   `tfsdk:"advertise_ipv4_isis"`
	AdvertiseIpv4IsisVariable        types.String `tfsdk:"advertise_ipv4_isis_variable"`
	AdvertiseIpv6Bgp                 types.Bool   `tfsdk:"advertise_ipv6_bgp"`
	AdvertiseIpv6BgpVariable         types.String `tfsdk:"advertise_ipv6_bgp_variable"`
	AdvertiseIpv6Ospf                types.Bool   `tfsdk:"advertise_ipv6_ospf"`
	AdvertiseIpv6OspfVariable        types.String `tfsdk:"advertise_ipv6_ospf_variable"`
	AdvertiseIpv6Connected           types.Bool   `tfsdk:"advertise_ipv6_connected"`
	AdvertiseIpv6ConnectedVariable   types.String `tfsdk:"advertise_ipv6_connected_variable"`
	AdvertiseIpv6Static              types.Bool   `tfsdk:"advertise_ipv6_static"`
	AdvertiseIpv6StaticVariable      types.String `tfsdk:"advertise_ipv6_static_variable"`
	AdvertiseIpv6Eigrp               types.Bool   `tfsdk:"advertise_ipv6_eigrp"`
	AdvertiseIpv6EigrpVariable       types.String `tfsdk:"advertise_ipv6_eigrp_variable"`
	AdvertiseIpv6Lisp                types.Bool   `tfsdk:"advertise_ipv6_lisp"`
	AdvertiseIpv6LispVariable        types.String `tfsdk:"advertise_ipv6_lisp_variable"`
	AdvertiseIpv6Isis                types.Bool   `tfsdk:"advertise_ipv6_isis"`
	AdvertiseIpv6IsisVariable        types.String `tfsdk:"advertise_ipv6_isis_variable"`
	IgnoreRegionPathLength           types.Bool   `tfsdk:"ignore_region_path_length"`
	IgnoreRegionPathLengthVariable   types.String `tfsdk:"ignore_region_path_length_variable"`
	TransportGateway                 types.String `tfsdk:"transport_gateway"`
	TransportGatewayVariable         types.String `tfsdk:"transport_gateway_variable"`
	SiteTypes                        types.Set    `tfsdk:"site_types"`
	SiteTypesVariable                types.String `tfsdk:"site_types_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemOMP) getModel() string {
	return "system_omp"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemOMP) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/omp", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemOMP) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.GracefulRestartVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"gracefulRestart.optionType", "variable")
			body, _ = sjson.Set(body, path+"gracefulRestart.value", data.GracefulRestartVariable.ValueString())
		}
	} else if data.GracefulRestart.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"gracefulRestart.optionType", "default")
			body, _ = sjson.Set(body, path+"gracefulRestart.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"gracefulRestart.optionType", "global")
			body, _ = sjson.Set(body, path+"gracefulRestart.value", data.GracefulRestart.ValueBool())
		}
	}

	if !data.OverlayAsVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"overlayAs.optionType", "variable")
			body, _ = sjson.Set(body, path+"overlayAs.value", data.OverlayAsVariable.ValueString())
		}
	} else if data.OverlayAs.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"overlayAs.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"overlayAs.optionType", "global")
			body, _ = sjson.Set(body, path+"overlayAs.value", data.OverlayAs.ValueInt64())
		}
	}

	if !data.PathsAdvertisedPerPrefixVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sendPathLimit.optionType", "variable")
			body, _ = sjson.Set(body, path+"sendPathLimit.value", data.PathsAdvertisedPerPrefixVariable.ValueString())
		}
	} else if data.PathsAdvertisedPerPrefix.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"sendPathLimit.optionType", "default")
			body, _ = sjson.Set(body, path+"sendPathLimit.value", 4)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"sendPathLimit.optionType", "global")
			body, _ = sjson.Set(body, path+"sendPathLimit.value", data.PathsAdvertisedPerPrefix.ValueInt64())
		}
	}

	if !data.EcmpLimitVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ecmpLimit.optionType", "variable")
			body, _ = sjson.Set(body, path+"ecmpLimit.value", data.EcmpLimitVariable.ValueString())
		}
	} else if data.EcmpLimit.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ecmpLimit.optionType", "default")
			body, _ = sjson.Set(body, path+"ecmpLimit.value", 4)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ecmpLimit.optionType", "global")
			body, _ = sjson.Set(body, path+"ecmpLimit.value", data.EcmpLimit.ValueInt64())
		}
	}

	if !data.ShutdownVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"shutdown.optionType", "variable")
			body, _ = sjson.Set(body, path+"shutdown.value", data.ShutdownVariable.ValueString())
		}
	} else if data.Shutdown.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"shutdown.optionType", "default")
			body, _ = sjson.Set(body, path+"shutdown.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"shutdown.optionType", "global")
			body, _ = sjson.Set(body, path+"shutdown.value", data.Shutdown.ValueBool())
		}
	}

	if !data.OmpAdminDistanceIpv4Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv4.optionType", "variable")
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv4.value", data.OmpAdminDistanceIpv4Variable.ValueString())
		}
	} else if data.OmpAdminDistanceIpv4.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv4.optionType", "default")
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv4.value", 251)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv4.optionType", "global")
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv4.value", data.OmpAdminDistanceIpv4.ValueInt64())
		}
	}

	if !data.OmpAdminDistanceIpv6Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv6.optionType", "variable")
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv6.value", data.OmpAdminDistanceIpv6Variable.ValueString())
		}
	} else if data.OmpAdminDistanceIpv6.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv6.optionType", "default")
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv6.value", 251)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv6.optionType", "global")
			body, _ = sjson.Set(body, path+"ompAdminDistanceIpv6.value", data.OmpAdminDistanceIpv6.ValueInt64())
		}
	}

	if !data.AdvertisementIntervalVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertisementInterval.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertisementInterval.value", data.AdvertisementIntervalVariable.ValueString())
		}
	} else if data.AdvertisementInterval.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertisementInterval.optionType", "default")
			body, _ = sjson.Set(body, path+"advertisementInterval.value", 1)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertisementInterval.optionType", "global")
			body, _ = sjson.Set(body, path+"advertisementInterval.value", data.AdvertisementInterval.ValueInt64())
		}
	}

	if !data.GracefulRestartTimerVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"gracefulRestartTimer.optionType", "variable")
			body, _ = sjson.Set(body, path+"gracefulRestartTimer.value", data.GracefulRestartTimerVariable.ValueString())
		}
	} else if data.GracefulRestartTimer.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"gracefulRestartTimer.optionType", "default")
			body, _ = sjson.Set(body, path+"gracefulRestartTimer.value", 43200)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"gracefulRestartTimer.optionType", "global")
			body, _ = sjson.Set(body, path+"gracefulRestartTimer.value", data.GracefulRestartTimer.ValueInt64())
		}
	}

	if !data.EorTimerVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"eorTimer.optionType", "variable")
			body, _ = sjson.Set(body, path+"eorTimer.value", data.EorTimerVariable.ValueString())
		}
	} else if data.EorTimer.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"eorTimer.optionType", "default")
			body, _ = sjson.Set(body, path+"eorTimer.value", 300)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"eorTimer.optionType", "global")
			body, _ = sjson.Set(body, path+"eorTimer.value", data.EorTimer.ValueInt64())
		}
	}

	if !data.HoldtimeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"holdtime.optionType", "variable")
			body, _ = sjson.Set(body, path+"holdtime.value", data.HoldtimeVariable.ValueString())
		}
	} else if data.Holdtime.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"holdtime.optionType", "default")
			body, _ = sjson.Set(body, path+"holdtime.value", 60)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"holdtime.optionType", "global")
			body, _ = sjson.Set(body, path+"holdtime.value", data.Holdtime.ValueInt64())
		}
	}

	if !data.AdvertiseIpv4BgpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.bgp.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv4.bgp.value", data.AdvertiseIpv4BgpVariable.ValueString())
		}
	} else if data.AdvertiseIpv4Bgp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.bgp.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv4.bgp.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.bgp.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv4.bgp.value", data.AdvertiseIpv4Bgp.ValueBool())
		}
	}

	if !data.AdvertiseIpv4OspfVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.ospf.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv4.ospf.value", data.AdvertiseIpv4OspfVariable.ValueString())
		}
	} else if data.AdvertiseIpv4Ospf.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.ospf.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv4.ospf.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.ospf.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv4.ospf.value", data.AdvertiseIpv4Ospf.ValueBool())
		}
	}

	if !data.AdvertiseIpv4OspfV3Variable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.ospfv3.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv4.ospfv3.value", data.AdvertiseIpv4OspfV3Variable.ValueString())
		}
	} else if data.AdvertiseIpv4OspfV3.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.ospfv3.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv4.ospfv3.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.ospfv3.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv4.ospfv3.value", data.AdvertiseIpv4OspfV3.ValueBool())
		}
	}

	if !data.AdvertiseIpv4ConnectedVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.connected.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv4.connected.value", data.AdvertiseIpv4ConnectedVariable.ValueString())
		}
	} else if data.AdvertiseIpv4Connected.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.connected.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv4.connected.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.connected.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv4.connected.value", data.AdvertiseIpv4Connected.ValueBool())
		}
	}

	if !data.AdvertiseIpv4StaticVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.static.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv4.static.value", data.AdvertiseIpv4StaticVariable.ValueString())
		}
	} else if data.AdvertiseIpv4Static.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.static.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv4.static.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.static.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv4.static.value", data.AdvertiseIpv4Static.ValueBool())
		}
	}

	if !data.AdvertiseIpv4EigrpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.eigrp.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv4.eigrp.value", data.AdvertiseIpv4EigrpVariable.ValueString())
		}
	} else if data.AdvertiseIpv4Eigrp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.eigrp.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv4.eigrp.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.eigrp.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv4.eigrp.value", data.AdvertiseIpv4Eigrp.ValueBool())
		}
	}

	if !data.AdvertiseIpv4LispVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.lisp.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv4.lisp.value", data.AdvertiseIpv4LispVariable.ValueString())
		}
	} else if data.AdvertiseIpv4Lisp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.lisp.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv4.lisp.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.lisp.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv4.lisp.value", data.AdvertiseIpv4Lisp.ValueBool())
		}
	}

	if !data.AdvertiseIpv4IsisVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.isis.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv4.isis.value", data.AdvertiseIpv4IsisVariable.ValueString())
		}
	} else if data.AdvertiseIpv4Isis.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.isis.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv4.isis.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv4.isis.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv4.isis.value", data.AdvertiseIpv4Isis.ValueBool())
		}
	}

	if !data.AdvertiseIpv6BgpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.bgp.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv6.bgp.value", data.AdvertiseIpv6BgpVariable.ValueString())
		}
	} else if data.AdvertiseIpv6Bgp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.bgp.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv6.bgp.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.bgp.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv6.bgp.value", data.AdvertiseIpv6Bgp.ValueBool())
		}
	}

	if !data.AdvertiseIpv6OspfVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.ospf.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv6.ospf.value", data.AdvertiseIpv6OspfVariable.ValueString())
		}
	} else if data.AdvertiseIpv6Ospf.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.ospf.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv6.ospf.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.ospf.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv6.ospf.value", data.AdvertiseIpv6Ospf.ValueBool())
		}
	}

	if !data.AdvertiseIpv6ConnectedVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.connected.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv6.connected.value", data.AdvertiseIpv6ConnectedVariable.ValueString())
		}
	} else if data.AdvertiseIpv6Connected.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.connected.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv6.connected.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.connected.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv6.connected.value", data.AdvertiseIpv6Connected.ValueBool())
		}
	}

	if !data.AdvertiseIpv6StaticVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.static.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv6.static.value", data.AdvertiseIpv6StaticVariable.ValueString())
		}
	} else if data.AdvertiseIpv6Static.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.static.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv6.static.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.static.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv6.static.value", data.AdvertiseIpv6Static.ValueBool())
		}
	}

	if !data.AdvertiseIpv6EigrpVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.eigrp.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv6.eigrp.value", data.AdvertiseIpv6EigrpVariable.ValueString())
		}
	} else if data.AdvertiseIpv6Eigrp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.eigrp.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv6.eigrp.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.eigrp.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv6.eigrp.value", data.AdvertiseIpv6Eigrp.ValueBool())
		}
	}

	if !data.AdvertiseIpv6LispVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.lisp.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv6.lisp.value", data.AdvertiseIpv6LispVariable.ValueString())
		}
	} else if data.AdvertiseIpv6Lisp.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.lisp.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv6.lisp.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.lisp.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv6.lisp.value", data.AdvertiseIpv6Lisp.ValueBool())
		}
	}

	if !data.AdvertiseIpv6IsisVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.isis.optionType", "variable")
			body, _ = sjson.Set(body, path+"advertiseIpv6.isis.value", data.AdvertiseIpv6IsisVariable.ValueString())
		}
	} else if data.AdvertiseIpv6Isis.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.isis.optionType", "default")
			body, _ = sjson.Set(body, path+"advertiseIpv6.isis.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"advertiseIpv6.isis.optionType", "global")
			body, _ = sjson.Set(body, path+"advertiseIpv6.isis.value", data.AdvertiseIpv6Isis.ValueBool())
		}
	}

	if !data.IgnoreRegionPathLengthVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ignoreRegionPathLength.optionType", "variable")
			body, _ = sjson.Set(body, path+"ignoreRegionPathLength.value", data.IgnoreRegionPathLengthVariable.ValueString())
		}
	} else if data.IgnoreRegionPathLength.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"ignoreRegionPathLength.optionType", "default")
			body, _ = sjson.Set(body, path+"ignoreRegionPathLength.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"ignoreRegionPathLength.optionType", "global")
			body, _ = sjson.Set(body, path+"ignoreRegionPathLength.value", data.IgnoreRegionPathLength.ValueBool())
		}
	}

	if !data.TransportGatewayVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"transportGateway.optionType", "variable")
			body, _ = sjson.Set(body, path+"transportGateway.value", data.TransportGatewayVariable.ValueString())
		}
	} else if data.TransportGateway.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"transportGateway.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"transportGateway.optionType", "global")
			body, _ = sjson.Set(body, path+"transportGateway.value", data.TransportGateway.ValueString())
		}
	}

	if !data.SiteTypesVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"siteTypes.optionType", "variable")
			body, _ = sjson.Set(body, path+"siteTypes.value", data.SiteTypesVariable.ValueString())
		}
	} else if data.SiteTypes.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"siteTypes.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"siteTypes.optionType", "global")
			var values []string
			data.SiteTypes.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"siteTypes.value", values)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemOMP) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.GracefulRestart = types.BoolNull()
	data.GracefulRestartVariable = types.StringNull()
	if t := res.Get(path + "gracefulRestart.optionType"); t.Exists() {
		va := res.Get(path + "gracefulRestart.value")
		if t.String() == "variable" {
			data.GracefulRestartVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GracefulRestart = types.BoolValue(va.Bool())
		}
	}
	data.OverlayAs = types.Int64Null()
	data.OverlayAsVariable = types.StringNull()
	if t := res.Get(path + "overlayAs.optionType"); t.Exists() {
		va := res.Get(path + "overlayAs.value")
		if t.String() == "variable" {
			data.OverlayAsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OverlayAs = types.Int64Value(va.Int())
		}
	}
	data.PathsAdvertisedPerPrefix = types.Int64Null()
	data.PathsAdvertisedPerPrefixVariable = types.StringNull()
	if t := res.Get(path + "sendPathLimit.optionType"); t.Exists() {
		va := res.Get(path + "sendPathLimit.value")
		if t.String() == "variable" {
			data.PathsAdvertisedPerPrefixVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PathsAdvertisedPerPrefix = types.Int64Value(va.Int())
		}
	}
	data.EcmpLimit = types.Int64Null()
	data.EcmpLimitVariable = types.StringNull()
	if t := res.Get(path + "ecmpLimit.optionType"); t.Exists() {
		va := res.Get(path + "ecmpLimit.value")
		if t.String() == "variable" {
			data.EcmpLimitVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EcmpLimit = types.Int64Value(va.Int())
		}
	}
	data.Shutdown = types.BoolNull()
	data.ShutdownVariable = types.StringNull()
	if t := res.Get(path + "shutdown.optionType"); t.Exists() {
		va := res.Get(path + "shutdown.value")
		if t.String() == "variable" {
			data.ShutdownVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Shutdown = types.BoolValue(va.Bool())
		}
	}
	data.OmpAdminDistanceIpv4 = types.Int64Null()
	data.OmpAdminDistanceIpv4Variable = types.StringNull()
	if t := res.Get(path + "ompAdminDistanceIpv4.optionType"); t.Exists() {
		va := res.Get(path + "ompAdminDistanceIpv4.value")
		if t.String() == "variable" {
			data.OmpAdminDistanceIpv4Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OmpAdminDistanceIpv4 = types.Int64Value(va.Int())
		}
	}
	data.OmpAdminDistanceIpv6 = types.Int64Null()
	data.OmpAdminDistanceIpv6Variable = types.StringNull()
	if t := res.Get(path + "ompAdminDistanceIpv6.optionType"); t.Exists() {
		va := res.Get(path + "ompAdminDistanceIpv6.value")
		if t.String() == "variable" {
			data.OmpAdminDistanceIpv6Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OmpAdminDistanceIpv6 = types.Int64Value(va.Int())
		}
	}
	data.AdvertisementInterval = types.Int64Null()
	data.AdvertisementIntervalVariable = types.StringNull()
	if t := res.Get(path + "advertisementInterval.optionType"); t.Exists() {
		va := res.Get(path + "advertisementInterval.value")
		if t.String() == "variable" {
			data.AdvertisementIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertisementInterval = types.Int64Value(va.Int())
		}
	}
	data.GracefulRestartTimer = types.Int64Null()
	data.GracefulRestartTimerVariable = types.StringNull()
	if t := res.Get(path + "gracefulRestartTimer.optionType"); t.Exists() {
		va := res.Get(path + "gracefulRestartTimer.value")
		if t.String() == "variable" {
			data.GracefulRestartTimerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GracefulRestartTimer = types.Int64Value(va.Int())
		}
	}
	data.EorTimer = types.Int64Null()
	data.EorTimerVariable = types.StringNull()
	if t := res.Get(path + "eorTimer.optionType"); t.Exists() {
		va := res.Get(path + "eorTimer.value")
		if t.String() == "variable" {
			data.EorTimerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EorTimer = types.Int64Value(va.Int())
		}
	}
	data.Holdtime = types.Int64Null()
	data.HoldtimeVariable = types.StringNull()
	if t := res.Get(path + "holdtime.optionType"); t.Exists() {
		va := res.Get(path + "holdtime.value")
		if t.String() == "variable" {
			data.HoldtimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Holdtime = types.Int64Value(va.Int())
		}
	}
	data.AdvertiseIpv4Bgp = types.BoolNull()
	data.AdvertiseIpv4BgpVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.bgp.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.bgp.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4BgpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4Bgp = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv4Ospf = types.BoolNull()
	data.AdvertiseIpv4OspfVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.ospf.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.ospf.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4OspfVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4Ospf = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv4OspfV3 = types.BoolNull()
	data.AdvertiseIpv4OspfV3Variable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.ospfv3.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.ospfv3.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4OspfV3Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4OspfV3 = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv4Connected = types.BoolNull()
	data.AdvertiseIpv4ConnectedVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.connected.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.connected.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4ConnectedVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4Connected = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv4Static = types.BoolNull()
	data.AdvertiseIpv4StaticVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.static.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.static.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4StaticVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4Static = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv4Eigrp = types.BoolNull()
	data.AdvertiseIpv4EigrpVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.eigrp.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.eigrp.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4EigrpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4Eigrp = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv4Lisp = types.BoolNull()
	data.AdvertiseIpv4LispVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.lisp.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.lisp.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4LispVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4Lisp = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv4Isis = types.BoolNull()
	data.AdvertiseIpv4IsisVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.isis.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.isis.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4IsisVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4Isis = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv6Bgp = types.BoolNull()
	data.AdvertiseIpv6BgpVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv6.bgp.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv6.bgp.value")
		if t.String() == "variable" {
			data.AdvertiseIpv6BgpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv6Bgp = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv6Ospf = types.BoolNull()
	data.AdvertiseIpv6OspfVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv6.ospf.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv6.ospf.value")
		if t.String() == "variable" {
			data.AdvertiseIpv6OspfVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv6Ospf = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv6Connected = types.BoolNull()
	data.AdvertiseIpv6ConnectedVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv6.connected.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv6.connected.value")
		if t.String() == "variable" {
			data.AdvertiseIpv6ConnectedVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv6Connected = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv6Static = types.BoolNull()
	data.AdvertiseIpv6StaticVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv6.static.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv6.static.value")
		if t.String() == "variable" {
			data.AdvertiseIpv6StaticVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv6Static = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv6Eigrp = types.BoolNull()
	data.AdvertiseIpv6EigrpVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv6.eigrp.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv6.eigrp.value")
		if t.String() == "variable" {
			data.AdvertiseIpv6EigrpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv6Eigrp = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv6Lisp = types.BoolNull()
	data.AdvertiseIpv6LispVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv6.lisp.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv6.lisp.value")
		if t.String() == "variable" {
			data.AdvertiseIpv6LispVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv6Lisp = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv6Isis = types.BoolNull()
	data.AdvertiseIpv6IsisVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv6.isis.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv6.isis.value")
		if t.String() == "variable" {
			data.AdvertiseIpv6IsisVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv6Isis = types.BoolValue(va.Bool())
		}
	}
	data.IgnoreRegionPathLength = types.BoolNull()
	data.IgnoreRegionPathLengthVariable = types.StringNull()
	if t := res.Get(path + "ignoreRegionPathLength.optionType"); t.Exists() {
		va := res.Get(path + "ignoreRegionPathLength.value")
		if t.String() == "variable" {
			data.IgnoreRegionPathLengthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IgnoreRegionPathLength = types.BoolValue(va.Bool())
		}
	}
	data.TransportGateway = types.StringNull()
	data.TransportGatewayVariable = types.StringNull()
	if t := res.Get(path + "transportGateway.optionType"); t.Exists() {
		va := res.Get(path + "transportGateway.value")
		if t.String() == "variable" {
			data.TransportGatewayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TransportGateway = types.StringValue(va.String())
		}
	}
	data.SiteTypes = types.SetNull(types.StringType)
	data.SiteTypesVariable = types.StringNull()
	if t := res.Get(path + "siteTypes.optionType"); t.Exists() {
		va := res.Get(path + "siteTypes.value")
		if t.String() == "variable" {
			data.SiteTypesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SiteTypes = helpers.GetStringSet(va.Array())
		}
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemOMP) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.GracefulRestart = types.BoolNull()
	data.GracefulRestartVariable = types.StringNull()
	if t := res.Get(path + "gracefulRestart.optionType"); t.Exists() {
		va := res.Get(path + "gracefulRestart.value")
		if t.String() == "variable" {
			data.GracefulRestartVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GracefulRestart = types.BoolValue(va.Bool())
		}
	}
	data.OverlayAs = types.Int64Null()
	data.OverlayAsVariable = types.StringNull()
	if t := res.Get(path + "overlayAs.optionType"); t.Exists() {
		va := res.Get(path + "overlayAs.value")
		if t.String() == "variable" {
			data.OverlayAsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OverlayAs = types.Int64Value(va.Int())
		}
	}
	data.PathsAdvertisedPerPrefix = types.Int64Null()
	data.PathsAdvertisedPerPrefixVariable = types.StringNull()
	if t := res.Get(path + "sendPathLimit.optionType"); t.Exists() {
		va := res.Get(path + "sendPathLimit.value")
		if t.String() == "variable" {
			data.PathsAdvertisedPerPrefixVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PathsAdvertisedPerPrefix = types.Int64Value(va.Int())
		}
	}
	data.EcmpLimit = types.Int64Null()
	data.EcmpLimitVariable = types.StringNull()
	if t := res.Get(path + "ecmpLimit.optionType"); t.Exists() {
		va := res.Get(path + "ecmpLimit.value")
		if t.String() == "variable" {
			data.EcmpLimitVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EcmpLimit = types.Int64Value(va.Int())
		}
	}
	data.Shutdown = types.BoolNull()
	data.ShutdownVariable = types.StringNull()
	if t := res.Get(path + "shutdown.optionType"); t.Exists() {
		va := res.Get(path + "shutdown.value")
		if t.String() == "variable" {
			data.ShutdownVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Shutdown = types.BoolValue(va.Bool())
		}
	}
	data.OmpAdminDistanceIpv4 = types.Int64Null()
	data.OmpAdminDistanceIpv4Variable = types.StringNull()
	if t := res.Get(path + "ompAdminDistanceIpv4.optionType"); t.Exists() {
		va := res.Get(path + "ompAdminDistanceIpv4.value")
		if t.String() == "variable" {
			data.OmpAdminDistanceIpv4Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OmpAdminDistanceIpv4 = types.Int64Value(va.Int())
		}
	}
	data.OmpAdminDistanceIpv6 = types.Int64Null()
	data.OmpAdminDistanceIpv6Variable = types.StringNull()
	if t := res.Get(path + "ompAdminDistanceIpv6.optionType"); t.Exists() {
		va := res.Get(path + "ompAdminDistanceIpv6.value")
		if t.String() == "variable" {
			data.OmpAdminDistanceIpv6Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OmpAdminDistanceIpv6 = types.Int64Value(va.Int())
		}
	}
	data.AdvertisementInterval = types.Int64Null()
	data.AdvertisementIntervalVariable = types.StringNull()
	if t := res.Get(path + "advertisementInterval.optionType"); t.Exists() {
		va := res.Get(path + "advertisementInterval.value")
		if t.String() == "variable" {
			data.AdvertisementIntervalVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertisementInterval = types.Int64Value(va.Int())
		}
	}
	data.GracefulRestartTimer = types.Int64Null()
	data.GracefulRestartTimerVariable = types.StringNull()
	if t := res.Get(path + "gracefulRestartTimer.optionType"); t.Exists() {
		va := res.Get(path + "gracefulRestartTimer.value")
		if t.String() == "variable" {
			data.GracefulRestartTimerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GracefulRestartTimer = types.Int64Value(va.Int())
		}
	}
	data.EorTimer = types.Int64Null()
	data.EorTimerVariable = types.StringNull()
	if t := res.Get(path + "eorTimer.optionType"); t.Exists() {
		va := res.Get(path + "eorTimer.value")
		if t.String() == "variable" {
			data.EorTimerVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EorTimer = types.Int64Value(va.Int())
		}
	}
	data.Holdtime = types.Int64Null()
	data.HoldtimeVariable = types.StringNull()
	if t := res.Get(path + "holdtime.optionType"); t.Exists() {
		va := res.Get(path + "holdtime.value")
		if t.String() == "variable" {
			data.HoldtimeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Holdtime = types.Int64Value(va.Int())
		}
	}
	data.AdvertiseIpv4Bgp = types.BoolNull()
	data.AdvertiseIpv4BgpVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.bgp.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.bgp.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4BgpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4Bgp = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv4Ospf = types.BoolNull()
	data.AdvertiseIpv4OspfVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.ospf.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.ospf.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4OspfVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4Ospf = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv4OspfV3 = types.BoolNull()
	data.AdvertiseIpv4OspfV3Variable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.ospfv3.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.ospfv3.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4OspfV3Variable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4OspfV3 = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv4Connected = types.BoolNull()
	data.AdvertiseIpv4ConnectedVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.connected.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.connected.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4ConnectedVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4Connected = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv4Static = types.BoolNull()
	data.AdvertiseIpv4StaticVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.static.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.static.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4StaticVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4Static = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv4Eigrp = types.BoolNull()
	data.AdvertiseIpv4EigrpVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.eigrp.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.eigrp.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4EigrpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4Eigrp = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv4Lisp = types.BoolNull()
	data.AdvertiseIpv4LispVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.lisp.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.lisp.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4LispVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4Lisp = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv4Isis = types.BoolNull()
	data.AdvertiseIpv4IsisVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv4.isis.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv4.isis.value")
		if t.String() == "variable" {
			data.AdvertiseIpv4IsisVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv4Isis = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv6Bgp = types.BoolNull()
	data.AdvertiseIpv6BgpVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv6.bgp.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv6.bgp.value")
		if t.String() == "variable" {
			data.AdvertiseIpv6BgpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv6Bgp = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv6Ospf = types.BoolNull()
	data.AdvertiseIpv6OspfVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv6.ospf.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv6.ospf.value")
		if t.String() == "variable" {
			data.AdvertiseIpv6OspfVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv6Ospf = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv6Connected = types.BoolNull()
	data.AdvertiseIpv6ConnectedVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv6.connected.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv6.connected.value")
		if t.String() == "variable" {
			data.AdvertiseIpv6ConnectedVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv6Connected = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv6Static = types.BoolNull()
	data.AdvertiseIpv6StaticVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv6.static.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv6.static.value")
		if t.String() == "variable" {
			data.AdvertiseIpv6StaticVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv6Static = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv6Eigrp = types.BoolNull()
	data.AdvertiseIpv6EigrpVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv6.eigrp.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv6.eigrp.value")
		if t.String() == "variable" {
			data.AdvertiseIpv6EigrpVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv6Eigrp = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv6Lisp = types.BoolNull()
	data.AdvertiseIpv6LispVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv6.lisp.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv6.lisp.value")
		if t.String() == "variable" {
			data.AdvertiseIpv6LispVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv6Lisp = types.BoolValue(va.Bool())
		}
	}
	data.AdvertiseIpv6Isis = types.BoolNull()
	data.AdvertiseIpv6IsisVariable = types.StringNull()
	if t := res.Get(path + "advertiseIpv6.isis.optionType"); t.Exists() {
		va := res.Get(path + "advertiseIpv6.isis.value")
		if t.String() == "variable" {
			data.AdvertiseIpv6IsisVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdvertiseIpv6Isis = types.BoolValue(va.Bool())
		}
	}
	data.IgnoreRegionPathLength = types.BoolNull()
	data.IgnoreRegionPathLengthVariable = types.StringNull()
	if t := res.Get(path + "ignoreRegionPathLength.optionType"); t.Exists() {
		va := res.Get(path + "ignoreRegionPathLength.value")
		if t.String() == "variable" {
			data.IgnoreRegionPathLengthVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IgnoreRegionPathLength = types.BoolValue(va.Bool())
		}
	}
	data.TransportGateway = types.StringNull()
	data.TransportGatewayVariable = types.StringNull()
	if t := res.Get(path + "transportGateway.optionType"); t.Exists() {
		va := res.Get(path + "transportGateway.value")
		if t.String() == "variable" {
			data.TransportGatewayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TransportGateway = types.StringValue(va.String())
		}
	}
	data.SiteTypes = types.SetNull(types.StringType)
	data.SiteTypesVariable = types.StringNull()
	if t := res.Get(path + "siteTypes.optionType"); t.Exists() {
		va := res.Get(path + "siteTypes.value")
		if t.String() == "variable" {
			data.SiteTypesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SiteTypes = helpers.GetStringSet(va.Array())
		}
	}
}

// End of section. //template:end updateFromBody

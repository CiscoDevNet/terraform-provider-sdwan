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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
type SystemBasic struct {
	Id                               types.String                     `tfsdk:"id"`
	Version                          types.Int64                      `tfsdk:"version"`
	Name                             types.String                     `tfsdk:"name"`
	Description                      types.String                     `tfsdk:"description"`
	FeatureProfileId                 types.String                     `tfsdk:"feature_profile_id"`
	Timezone                         types.String                     `tfsdk:"timezone"`
	TimezoneVariable                 types.String                     `tfsdk:"timezone_variable"`
	ConfigDescription                types.String                     `tfsdk:"config_description"`
	ConfigDescriptionVariable        types.String                     `tfsdk:"config_description_variable"`
	Location                         types.String                     `tfsdk:"location"`
	LocationVariable                 types.String                     `tfsdk:"location_variable"`
	GpsLongitude                     types.Float64                    `tfsdk:"gps_longitude"`
	GpsLongitudeVariable             types.String                     `tfsdk:"gps_longitude_variable"`
	GpsLatitude                      types.Float64                    `tfsdk:"gps_latitude"`
	GpsLatitudeVariable              types.String                     `tfsdk:"gps_latitude_variable"`
	GpsGeoFencingEnable              types.Bool                       `tfsdk:"gps_geo_fencing_enable"`
	GpsGeoFencingRange               types.Int64                      `tfsdk:"gps_geo_fencing_range"`
	GpsGeoFencingRangeVariable       types.String                     `tfsdk:"gps_geo_fencing_range_variable"`
	GpsSmsEnable                     types.Bool                       `tfsdk:"gps_sms_enable"`
	GpsSmsMobileNumbers              []SystemBasicGpsSmsMobileNumbers `tfsdk:"gps_sms_mobile_numbers"`
	DeviceGroups                     types.Set                        `tfsdk:"device_groups"`
	DeviceGroupsVariable             types.String                     `tfsdk:"device_groups_variable"`
	ControllerGroups                 types.Set                        `tfsdk:"controller_groups"`
	ControllerGroupsVariable         types.String                     `tfsdk:"controller_groups_variable"`
	OverlayId                        types.Int64                      `tfsdk:"overlay_id"`
	OverlayIdVariable                types.String                     `tfsdk:"overlay_id_variable"`
	PortOffset                       types.Int64                      `tfsdk:"port_offset"`
	PortOffsetVariable               types.String                     `tfsdk:"port_offset_variable"`
	PortHopping                      types.Bool                       `tfsdk:"port_hopping"`
	PortHoppingVariable              types.String                     `tfsdk:"port_hopping_variable"`
	ControlSessionPps                types.Int64                      `tfsdk:"control_session_pps"`
	ControlSessionPpsVariable        types.String                     `tfsdk:"control_session_pps_variable"`
	TrackTransport                   types.Bool                       `tfsdk:"track_transport"`
	TrackTransportVariable           types.String                     `tfsdk:"track_transport_variable"`
	TrackInterfaceTag                types.Int64                      `tfsdk:"track_interface_tag"`
	TrackInterfaceTagVariable        types.String                     `tfsdk:"track_interface_tag_variable"`
	ConsoleBaudRate                  types.String                     `tfsdk:"console_baud_rate"`
	ConsoleBaudRateVariable          types.String                     `tfsdk:"console_baud_rate_variable"`
	MaxOmpSessions                   types.Int64                      `tfsdk:"max_omp_sessions"`
	MaxOmpSessionsVariable           types.String                     `tfsdk:"max_omp_sessions_variable"`
	MultiTenant                      types.Bool                       `tfsdk:"multi_tenant"`
	MultiTenantVariable              types.String                     `tfsdk:"multi_tenant_variable"`
	TrackDefaultGateway              types.Bool                       `tfsdk:"track_default_gateway"`
	TrackDefaultGatewayVariable      types.String                     `tfsdk:"track_default_gateway_variable"`
	AdminTechOnFailure               types.Bool                       `tfsdk:"admin_tech_on_failure"`
	AdminTechOnFailureVariable       types.String                     `tfsdk:"admin_tech_on_failure_variable"`
	IdleTimeout                      types.Int64                      `tfsdk:"idle_timeout"`
	IdleTimeoutVariable              types.String                     `tfsdk:"idle_timeout_variable"`
	OnDemandEnable                   types.Bool                       `tfsdk:"on_demand_enable"`
	OnDemandEnableVariable           types.String                     `tfsdk:"on_demand_enable_variable"`
	OnDemandIdleTimeout              types.Int64                      `tfsdk:"on_demand_idle_timeout"`
	OnDemandIdleTimeoutVariable      types.String                     `tfsdk:"on_demand_idle_timeout_variable"`
	TransportGateway                 types.Bool                       `tfsdk:"transport_gateway"`
	TransportGatewayVariable         types.String                     `tfsdk:"transport_gateway_variable"`
	EnhancedAppAwareRouting          types.String                     `tfsdk:"enhanced_app_aware_routing"`
	EnhancedAppAwareRoutingVariable  types.String                     `tfsdk:"enhanced_app_aware_routing_variable"`
	SiteTypes                        types.Set                        `tfsdk:"site_types"`
	SiteTypesVariable                types.String                     `tfsdk:"site_types_variable"`
	AffinityGroupNumber              types.Int64                      `tfsdk:"affinity_group_number"`
	AffinityGroupNumberVariable      types.String                     `tfsdk:"affinity_group_number_variable"`
	AffinityGroupPreferences         types.Set                        `tfsdk:"affinity_group_preferences"`
	AffinityGroupPreferencesVariable types.String                     `tfsdk:"affinity_group_preferences_variable"`
	AffinityPreferenceAuto           types.Bool                       `tfsdk:"affinity_preference_auto"`
	AffinityPreferenceAutoVariable   types.String                     `tfsdk:"affinity_preference_auto_variable"`
	AffinityPerVrfs                  []SystemBasicAffinityPerVrfs     `tfsdk:"affinity_per_vrfs"`
}

type SystemBasicGpsSmsMobileNumbers struct {
	Number         types.String `tfsdk:"number"`
	NumberVariable types.String `tfsdk:"number_variable"`
}

type SystemBasicAffinityPerVrfs struct {
	AffinityGroupNumber         types.Int64  `tfsdk:"affinity_group_number"`
	AffinityGroupNumberVariable types.String `tfsdk:"affinity_group_number_variable"`
	VrfRange                    types.String `tfsdk:"vrf_range"`
	VrfRangeVariable            types.String `tfsdk:"vrf_range_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data SystemBasic) getModel() string {
	return "system_basic"
}

// End of section. //template:end getModel

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data SystemBasic) getPath() string {
	return fmt.Sprintf("/v1/feature-profile/sdwan/system/%v/basic", url.QueryEscape(data.FeatureProfileId.ValueString()))
}

// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data SystemBasic) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	path := "data."

	if !data.TimezoneVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"clock.timezone.optionType", "variable")
			body, _ = sjson.Set(body, path+"clock.timezone.value", data.TimezoneVariable.ValueString())
		}
	} else if data.Timezone.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"clock.timezone.optionType", "default")
			body, _ = sjson.Set(body, path+"clock.timezone.value", "UTC")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"clock.timezone.optionType", "global")
			body, _ = sjson.Set(body, path+"clock.timezone.value", data.Timezone.ValueString())
		}
	}

	if !data.ConfigDescriptionVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"description.optionType", "variable")
			body, _ = sjson.Set(body, path+"description.value", data.ConfigDescriptionVariable.ValueString())
		}
	} else if data.ConfigDescription.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"description.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"description.optionType", "global")
			body, _ = sjson.Set(body, path+"description.value", data.ConfigDescription.ValueString())
		}
	}

	if !data.LocationVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"location.optionType", "variable")
			body, _ = sjson.Set(body, path+"location.value", data.LocationVariable.ValueString())
		}
	} else if data.Location.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"location.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"location.optionType", "global")
			body, _ = sjson.Set(body, path+"location.value", data.Location.ValueString())
		}
	}

	if !data.GpsLongitudeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"gpsLocation.longitude.optionType", "variable")
			body, _ = sjson.Set(body, path+"gpsLocation.longitude.value", data.GpsLongitudeVariable.ValueString())
		}
	} else if data.GpsLongitude.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"gpsLocation.longitude.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"gpsLocation.longitude.optionType", "global")
			body, _ = sjson.Set(body, path+"gpsLocation.longitude.value", data.GpsLongitude.ValueFloat64())
		}
	}

	if !data.GpsLatitudeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"gpsLocation.latitude.optionType", "variable")
			body, _ = sjson.Set(body, path+"gpsLocation.latitude.value", data.GpsLatitudeVariable.ValueString())
		}
	} else if data.GpsLatitude.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"gpsLocation.latitude.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"gpsLocation.latitude.optionType", "global")
			body, _ = sjson.Set(body, path+"gpsLocation.latitude.value", data.GpsLatitude.ValueFloat64())
		}
	}
	if data.GpsGeoFencingEnable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.enable.optionType", "default")
			body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.enable.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.enable.optionType", "global")
			body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.enable.value", data.GpsGeoFencingEnable.ValueBool())
		}
	}

	if !data.GpsGeoFencingRangeVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.range.optionType", "variable")
			body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.range.value", data.GpsGeoFencingRangeVariable.ValueString())
		}
	} else if data.GpsGeoFencingRange.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.range.optionType", "default")
			body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.range.value", 100)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.range.optionType", "global")
			body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.range.value", data.GpsGeoFencingRange.ValueInt64())
		}
	}
	if data.GpsSmsEnable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.sms.enable.optionType", "default")
			body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.sms.enable.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.sms.enable.optionType", "global")
			body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.sms.enable.value", data.GpsSmsEnable.ValueBool())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"gpsLocation.geoFencing.sms.mobileNumber", []interface{}{})
		for _, item := range data.GpsSmsMobileNumbers {
			itemBody := ""

			if !item.NumberVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "number.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "number.value", item.NumberVariable.ValueString())
				}
			} else if !item.Number.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "number.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "number.value", item.Number.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"gpsLocation.geoFencing.sms.mobileNumber.-1", itemBody)
		}
	}

	if !data.DeviceGroupsVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"deviceGroups.optionType", "variable")
			body, _ = sjson.Set(body, path+"deviceGroups.value", data.DeviceGroupsVariable.ValueString())
		}
	} else if data.DeviceGroups.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"deviceGroups.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"deviceGroups.optionType", "global")
			var values []string
			data.DeviceGroups.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"deviceGroups.value", values)
		}
	}

	if !data.ControllerGroupsVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"controllerGroupList.optionType", "variable")
			body, _ = sjson.Set(body, path+"controllerGroupList.value", data.ControllerGroupsVariable.ValueString())
		}
	} else if data.ControllerGroups.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"controllerGroupList.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"controllerGroupList.optionType", "global")
			var values []int64
			data.ControllerGroups.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"controllerGroupList.value", values)
		}
	}

	if !data.OverlayIdVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"overlayId.optionType", "variable")
			body, _ = sjson.Set(body, path+"overlayId.value", data.OverlayIdVariable.ValueString())
		}
	} else if data.OverlayId.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"overlayId.optionType", "default")
			body, _ = sjson.Set(body, path+"overlayId.value", 1)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"overlayId.optionType", "global")
			body, _ = sjson.Set(body, path+"overlayId.value", data.OverlayId.ValueInt64())
		}
	}

	if !data.PortOffsetVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"portOffset.optionType", "variable")
			body, _ = sjson.Set(body, path+"portOffset.value", data.PortOffsetVariable.ValueString())
		}
	} else if data.PortOffset.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"portOffset.optionType", "default")
			body, _ = sjson.Set(body, path+"portOffset.value", 0)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"portOffset.optionType", "global")
			body, _ = sjson.Set(body, path+"portOffset.value", data.PortOffset.ValueInt64())
		}
	}

	if !data.PortHoppingVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"portHop.optionType", "variable")
			body, _ = sjson.Set(body, path+"portHop.value", data.PortHoppingVariable.ValueString())
		}
	} else if data.PortHopping.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"portHop.optionType", "default")
			body, _ = sjson.Set(body, path+"portHop.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"portHop.optionType", "global")
			body, _ = sjson.Set(body, path+"portHop.value", data.PortHopping.ValueBool())
		}
	}

	if !data.ControlSessionPpsVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"controlSessionPps.optionType", "variable")
			body, _ = sjson.Set(body, path+"controlSessionPps.value", data.ControlSessionPpsVariable.ValueString())
		}
	} else if data.ControlSessionPps.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"controlSessionPps.optionType", "default")
			body, _ = sjson.Set(body, path+"controlSessionPps.value", 300)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"controlSessionPps.optionType", "global")
			body, _ = sjson.Set(body, path+"controlSessionPps.value", data.ControlSessionPps.ValueInt64())
		}
	}

	if !data.TrackTransportVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trackTransport.optionType", "variable")
			body, _ = sjson.Set(body, path+"trackTransport.value", data.TrackTransportVariable.ValueString())
		}
	} else if data.TrackTransport.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trackTransport.optionType", "default")
			body, _ = sjson.Set(body, path+"trackTransport.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"trackTransport.optionType", "global")
			body, _ = sjson.Set(body, path+"trackTransport.value", data.TrackTransport.ValueBool())
		}
	}

	if !data.TrackInterfaceTagVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trackInterfaceTag.optionType", "variable")
			body, _ = sjson.Set(body, path+"trackInterfaceTag.value", data.TrackInterfaceTagVariable.ValueString())
		}
	} else if data.TrackInterfaceTag.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trackInterfaceTag.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"trackInterfaceTag.optionType", "global")
			body, _ = sjson.Set(body, path+"trackInterfaceTag.value", data.TrackInterfaceTag.ValueInt64())
		}
	}

	if !data.ConsoleBaudRateVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"consoleBaudRate.optionType", "variable")
			body, _ = sjson.Set(body, path+"consoleBaudRate.value", data.ConsoleBaudRateVariable.ValueString())
		}
	} else if data.ConsoleBaudRate.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"consoleBaudRate.optionType", "default")
			body, _ = sjson.Set(body, path+"consoleBaudRate.value", "9600")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"consoleBaudRate.optionType", "global")
			body, _ = sjson.Set(body, path+"consoleBaudRate.value", data.ConsoleBaudRate.ValueString())
		}
	}

	if !data.MaxOmpSessionsVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"maxOmpSessions.optionType", "variable")
			body, _ = sjson.Set(body, path+"maxOmpSessions.value", data.MaxOmpSessionsVariable.ValueString())
		}
	} else if data.MaxOmpSessions.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"maxOmpSessions.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"maxOmpSessions.optionType", "global")
			body, _ = sjson.Set(body, path+"maxOmpSessions.value", data.MaxOmpSessions.ValueInt64())
		}
	}

	if !data.MultiTenantVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"multiTenant.optionType", "variable")
			body, _ = sjson.Set(body, path+"multiTenant.value", data.MultiTenantVariable.ValueString())
		}
	} else if data.MultiTenant.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"multiTenant.optionType", "default")
			body, _ = sjson.Set(body, path+"multiTenant.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"multiTenant.optionType", "global")
			body, _ = sjson.Set(body, path+"multiTenant.value", data.MultiTenant.ValueBool())
		}
	}

	if !data.TrackDefaultGatewayVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trackDefaultGateway.optionType", "variable")
			body, _ = sjson.Set(body, path+"trackDefaultGateway.value", data.TrackDefaultGatewayVariable.ValueString())
		}
	} else if data.TrackDefaultGateway.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"trackDefaultGateway.optionType", "default")
			body, _ = sjson.Set(body, path+"trackDefaultGateway.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"trackDefaultGateway.optionType", "global")
			body, _ = sjson.Set(body, path+"trackDefaultGateway.value", data.TrackDefaultGateway.ValueBool())
		}
	}

	if !data.AdminTechOnFailureVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"adminTechOnFailure.optionType", "variable")
			body, _ = sjson.Set(body, path+"adminTechOnFailure.value", data.AdminTechOnFailureVariable.ValueString())
		}
	} else if data.AdminTechOnFailure.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"adminTechOnFailure.optionType", "default")
			body, _ = sjson.Set(body, path+"adminTechOnFailure.value", true)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"adminTechOnFailure.optionType", "global")
			body, _ = sjson.Set(body, path+"adminTechOnFailure.value", data.AdminTechOnFailure.ValueBool())
		}
	}

	if !data.IdleTimeoutVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"idleTimeout.optionType", "variable")
			body, _ = sjson.Set(body, path+"idleTimeout.value", data.IdleTimeoutVariable.ValueString())
		}
	} else if data.IdleTimeout.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"idleTimeout.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"idleTimeout.optionType", "global")
			body, _ = sjson.Set(body, path+"idleTimeout.value", data.IdleTimeout.ValueInt64())
		}
	}

	if !data.OnDemandEnableVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"onDemand.onDemandEnable.optionType", "variable")
			body, _ = sjson.Set(body, path+"onDemand.onDemandEnable.value", data.OnDemandEnableVariable.ValueString())
		}
	} else if data.OnDemandEnable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"onDemand.onDemandEnable.optionType", "default")
			body, _ = sjson.Set(body, path+"onDemand.onDemandEnable.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"onDemand.onDemandEnable.optionType", "global")
			body, _ = sjson.Set(body, path+"onDemand.onDemandEnable.value", data.OnDemandEnable.ValueBool())
		}
	}

	if !data.OnDemandIdleTimeoutVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"onDemand.onDemandIdleTimeout.optionType", "variable")
			body, _ = sjson.Set(body, path+"onDemand.onDemandIdleTimeout.value", data.OnDemandIdleTimeoutVariable.ValueString())
		}
	} else if data.OnDemandIdleTimeout.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"onDemand.onDemandIdleTimeout.optionType", "default")
			body, _ = sjson.Set(body, path+"onDemand.onDemandIdleTimeout.value", 10)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"onDemand.onDemandIdleTimeout.optionType", "global")
			body, _ = sjson.Set(body, path+"onDemand.onDemandIdleTimeout.value", data.OnDemandIdleTimeout.ValueInt64())
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
			body, _ = sjson.Set(body, path+"transportGateway.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"transportGateway.optionType", "global")
			body, _ = sjson.Set(body, path+"transportGateway.value", data.TransportGateway.ValueBool())
		}
	}

	if !data.EnhancedAppAwareRoutingVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"epfr.optionType", "variable")
			body, _ = sjson.Set(body, path+"epfr.value", data.EnhancedAppAwareRoutingVariable.ValueString())
		}
	} else if data.EnhancedAppAwareRouting.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"epfr.optionType", "default")
			body, _ = sjson.Set(body, path+"epfr.value", "disabled")
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"epfr.optionType", "global")
			body, _ = sjson.Set(body, path+"epfr.value", data.EnhancedAppAwareRouting.ValueString())
		}
	}

	if !data.SiteTypesVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"siteType.optionType", "variable")
			body, _ = sjson.Set(body, path+"siteType.value", data.SiteTypesVariable.ValueString())
		}
	} else if data.SiteTypes.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"siteType.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"siteType.optionType", "global")
			var values []string
			data.SiteTypes.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"siteType.value", values)
		}
	}

	if !data.AffinityGroupNumberVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"affinityGroupNumber.optionType", "variable")
			body, _ = sjson.Set(body, path+"affinityGroupNumber.value", data.AffinityGroupNumberVariable.ValueString())
		}
	} else if data.AffinityGroupNumber.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"affinityGroupNumber.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"affinityGroupNumber.optionType", "global")
			body, _ = sjson.Set(body, path+"affinityGroupNumber.value", data.AffinityGroupNumber.ValueInt64())
		}
	}

	if !data.AffinityGroupPreferencesVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"affinityGroupPreference.optionType", "variable")
			body, _ = sjson.Set(body, path+"affinityGroupPreference.value", data.AffinityGroupPreferencesVariable.ValueString())
		}
	} else if data.AffinityGroupPreferences.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"affinityGroupPreference.optionType", "default")

		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"affinityGroupPreference.optionType", "global")
			var values []int64
			data.AffinityGroupPreferences.ElementsAs(ctx, &values, false)
			body, _ = sjson.Set(body, path+"affinityGroupPreference.value", values)
		}
	}

	if !data.AffinityPreferenceAutoVariable.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"affinityPreferenceAuto.optionType", "variable")
			body, _ = sjson.Set(body, path+"affinityPreferenceAuto.value", data.AffinityPreferenceAutoVariable.ValueString())
		}
	} else if data.AffinityPreferenceAuto.IsNull() {
		if true {
			body, _ = sjson.Set(body, path+"affinityPreferenceAuto.optionType", "default")
			body, _ = sjson.Set(body, path+"affinityPreferenceAuto.value", false)
		}
	} else {
		if true {
			body, _ = sjson.Set(body, path+"affinityPreferenceAuto.optionType", "global")
			body, _ = sjson.Set(body, path+"affinityPreferenceAuto.value", data.AffinityPreferenceAuto.ValueBool())
		}
	}
	if true {
		body, _ = sjson.Set(body, path+"affinityPerVrf", []interface{}{})
		for _, item := range data.AffinityPerVrfs {
			itemBody := ""

			if !item.AffinityGroupNumberVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "affinityGroupNumber.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "affinityGroupNumber.value", item.AffinityGroupNumberVariable.ValueString())
				}
			} else if item.AffinityGroupNumber.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "affinityGroupNumber.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "affinityGroupNumber.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "affinityGroupNumber.value", item.AffinityGroupNumber.ValueInt64())
				}
			}

			if !item.VrfRangeVariable.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vrfRange.optionType", "variable")
					itemBody, _ = sjson.Set(itemBody, "vrfRange.value", item.VrfRangeVariable.ValueString())
				}
			} else if item.VrfRange.IsNull() {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vrfRange.optionType", "default")

				}
			} else {
				if true {
					itemBody, _ = sjson.Set(itemBody, "vrfRange.optionType", "global")
					itemBody, _ = sjson.Set(itemBody, "vrfRange.value", item.VrfRange.ValueString())
				}
			}
			body, _ = sjson.SetRaw(body, path+"affinityPerVrf.-1", itemBody)
		}
	}
	return body
}

// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *SystemBasic) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Timezone = types.StringNull()
	data.TimezoneVariable = types.StringNull()
	if t := res.Get(path + "clock.timezone.optionType"); t.Exists() {
		va := res.Get(path + "clock.timezone.value")
		if t.String() == "variable" {
			data.TimezoneVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Timezone = types.StringValue(va.String())
		}
	}
	data.ConfigDescription = types.StringNull()
	data.ConfigDescriptionVariable = types.StringNull()
	if t := res.Get(path + "description.optionType"); t.Exists() {
		va := res.Get(path + "description.value")
		if t.String() == "variable" {
			data.ConfigDescriptionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ConfigDescription = types.StringValue(va.String())
		}
	}
	data.Location = types.StringNull()
	data.LocationVariable = types.StringNull()
	if t := res.Get(path + "location.optionType"); t.Exists() {
		va := res.Get(path + "location.value")
		if t.String() == "variable" {
			data.LocationVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Location = types.StringValue(va.String())
		}
	}
	data.GpsLongitude = types.Float64Null()
	data.GpsLongitudeVariable = types.StringNull()
	if t := res.Get(path + "gpsLocation.longitude.optionType"); t.Exists() {
		va := res.Get(path + "gpsLocation.longitude.value")
		if t.String() == "variable" {
			data.GpsLongitudeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GpsLongitude = types.Float64Value(va.Float())
		}
	}
	data.GpsLatitude = types.Float64Null()
	data.GpsLatitudeVariable = types.StringNull()
	if t := res.Get(path + "gpsLocation.latitude.optionType"); t.Exists() {
		va := res.Get(path + "gpsLocation.latitude.value")
		if t.String() == "variable" {
			data.GpsLatitudeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GpsLatitude = types.Float64Value(va.Float())
		}
	}
	data.GpsGeoFencingEnable = types.BoolNull()

	if t := res.Get(path + "gpsLocation.geoFencing.enable.optionType"); t.Exists() {
		va := res.Get(path + "gpsLocation.geoFencing.enable.value")
		if t.String() == "global" {
			data.GpsGeoFencingEnable = types.BoolValue(va.Bool())
		}
	}
	data.GpsGeoFencingRange = types.Int64Null()
	data.GpsGeoFencingRangeVariable = types.StringNull()
	if t := res.Get(path + "gpsLocation.geoFencing.range.optionType"); t.Exists() {
		va := res.Get(path + "gpsLocation.geoFencing.range.value")
		if t.String() == "variable" {
			data.GpsGeoFencingRangeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GpsGeoFencingRange = types.Int64Value(va.Int())
		}
	}
	data.GpsSmsEnable = types.BoolNull()

	if t := res.Get(path + "gpsLocation.geoFencing.sms.enable.optionType"); t.Exists() {
		va := res.Get(path + "gpsLocation.geoFencing.sms.enable.value")
		if t.String() == "global" {
			data.GpsSmsEnable = types.BoolValue(va.Bool())
		}
	}
	if value := res.Get(path + "gpsLocation.geoFencing.sms.mobileNumber"); value.Exists() {
		data.GpsSmsMobileNumbers = make([]SystemBasicGpsSmsMobileNumbers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemBasicGpsSmsMobileNumbers{}
			item.Number = types.StringNull()
			item.NumberVariable = types.StringNull()
			if t := v.Get("number.optionType"); t.Exists() {
				va := v.Get("number.value")
				if t.String() == "variable" {
					item.NumberVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.Number = types.StringValue(va.String())
				}
			}
			data.GpsSmsMobileNumbers = append(data.GpsSmsMobileNumbers, item)
			return true
		})
	}
	data.DeviceGroups = types.SetNull(types.StringType)
	data.DeviceGroupsVariable = types.StringNull()
	if t := res.Get(path + "deviceGroups.optionType"); t.Exists() {
		va := res.Get(path + "deviceGroups.value")
		if t.String() == "variable" {
			data.DeviceGroupsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DeviceGroups = helpers.GetStringSet(va.Array())
		}
	}
	data.ControllerGroups = types.SetNull(types.Int64Type)
	data.ControllerGroupsVariable = types.StringNull()
	if t := res.Get(path + "controllerGroupList.optionType"); t.Exists() {
		va := res.Get(path + "controllerGroupList.value")
		if t.String() == "variable" {
			data.ControllerGroupsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ControllerGroups = helpers.GetInt64Set(va.Array())
		}
	}
	data.OverlayId = types.Int64Null()
	data.OverlayIdVariable = types.StringNull()
	if t := res.Get(path + "overlayId.optionType"); t.Exists() {
		va := res.Get(path + "overlayId.value")
		if t.String() == "variable" {
			data.OverlayIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OverlayId = types.Int64Value(va.Int())
		}
	}
	data.PortOffset = types.Int64Null()
	data.PortOffsetVariable = types.StringNull()
	if t := res.Get(path + "portOffset.optionType"); t.Exists() {
		va := res.Get(path + "portOffset.value")
		if t.String() == "variable" {
			data.PortOffsetVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PortOffset = types.Int64Value(va.Int())
		}
	}
	data.PortHopping = types.BoolNull()
	data.PortHoppingVariable = types.StringNull()
	if t := res.Get(path + "portHop.optionType"); t.Exists() {
		va := res.Get(path + "portHop.value")
		if t.String() == "variable" {
			data.PortHoppingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PortHopping = types.BoolValue(va.Bool())
		}
	}
	data.ControlSessionPps = types.Int64Null()
	data.ControlSessionPpsVariable = types.StringNull()
	if t := res.Get(path + "controlSessionPps.optionType"); t.Exists() {
		va := res.Get(path + "controlSessionPps.value")
		if t.String() == "variable" {
			data.ControlSessionPpsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ControlSessionPps = types.Int64Value(va.Int())
		}
	}
	data.TrackTransport = types.BoolNull()
	data.TrackTransportVariable = types.StringNull()
	if t := res.Get(path + "trackTransport.optionType"); t.Exists() {
		va := res.Get(path + "trackTransport.value")
		if t.String() == "variable" {
			data.TrackTransportVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackTransport = types.BoolValue(va.Bool())
		}
	}
	data.TrackInterfaceTag = types.Int64Null()
	data.TrackInterfaceTagVariable = types.StringNull()
	if t := res.Get(path + "trackInterfaceTag.optionType"); t.Exists() {
		va := res.Get(path + "trackInterfaceTag.value")
		if t.String() == "variable" {
			data.TrackInterfaceTagVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackInterfaceTag = types.Int64Value(va.Int())
		}
	}
	data.ConsoleBaudRate = types.StringNull()
	data.ConsoleBaudRateVariable = types.StringNull()
	if t := res.Get(path + "consoleBaudRate.optionType"); t.Exists() {
		va := res.Get(path + "consoleBaudRate.value")
		if t.String() == "variable" {
			data.ConsoleBaudRateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ConsoleBaudRate = types.StringValue(va.String())
		}
	}
	data.MaxOmpSessions = types.Int64Null()
	data.MaxOmpSessionsVariable = types.StringNull()
	if t := res.Get(path + "maxOmpSessions.optionType"); t.Exists() {
		va := res.Get(path + "maxOmpSessions.value")
		if t.String() == "variable" {
			data.MaxOmpSessionsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MaxOmpSessions = types.Int64Value(va.Int())
		}
	}
	data.MultiTenant = types.BoolNull()
	data.MultiTenantVariable = types.StringNull()
	if t := res.Get(path + "multiTenant.optionType"); t.Exists() {
		va := res.Get(path + "multiTenant.value")
		if t.String() == "variable" {
			data.MultiTenantVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MultiTenant = types.BoolValue(va.Bool())
		}
	}
	data.TrackDefaultGateway = types.BoolNull()
	data.TrackDefaultGatewayVariable = types.StringNull()
	if t := res.Get(path + "trackDefaultGateway.optionType"); t.Exists() {
		va := res.Get(path + "trackDefaultGateway.value")
		if t.String() == "variable" {
			data.TrackDefaultGatewayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackDefaultGateway = types.BoolValue(va.Bool())
		}
	}
	data.AdminTechOnFailure = types.BoolNull()
	data.AdminTechOnFailureVariable = types.StringNull()
	if t := res.Get(path + "adminTechOnFailure.optionType"); t.Exists() {
		va := res.Get(path + "adminTechOnFailure.value")
		if t.String() == "variable" {
			data.AdminTechOnFailureVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdminTechOnFailure = types.BoolValue(va.Bool())
		}
	}
	data.IdleTimeout = types.Int64Null()
	data.IdleTimeoutVariable = types.StringNull()
	if t := res.Get(path + "idleTimeout.optionType"); t.Exists() {
		va := res.Get(path + "idleTimeout.value")
		if t.String() == "variable" {
			data.IdleTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IdleTimeout = types.Int64Value(va.Int())
		}
	}
	data.OnDemandEnable = types.BoolNull()
	data.OnDemandEnableVariable = types.StringNull()
	if t := res.Get(path + "onDemand.onDemandEnable.optionType"); t.Exists() {
		va := res.Get(path + "onDemand.onDemandEnable.value")
		if t.String() == "variable" {
			data.OnDemandEnableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OnDemandEnable = types.BoolValue(va.Bool())
		}
	}
	data.OnDemandIdleTimeout = types.Int64Null()
	data.OnDemandIdleTimeoutVariable = types.StringNull()
	if t := res.Get(path + "onDemand.onDemandIdleTimeout.optionType"); t.Exists() {
		va := res.Get(path + "onDemand.onDemandIdleTimeout.value")
		if t.String() == "variable" {
			data.OnDemandIdleTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OnDemandIdleTimeout = types.Int64Value(va.Int())
		}
	}
	data.TransportGateway = types.BoolNull()
	data.TransportGatewayVariable = types.StringNull()
	if t := res.Get(path + "transportGateway.optionType"); t.Exists() {
		va := res.Get(path + "transportGateway.value")
		if t.String() == "variable" {
			data.TransportGatewayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TransportGateway = types.BoolValue(va.Bool())
		}
	}
	data.EnhancedAppAwareRouting = types.StringNull()
	data.EnhancedAppAwareRoutingVariable = types.StringNull()
	if t := res.Get(path + "epfr.optionType"); t.Exists() {
		va := res.Get(path + "epfr.value")
		if t.String() == "variable" {
			data.EnhancedAppAwareRoutingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EnhancedAppAwareRouting = types.StringValue(va.String())
		}
	}
	data.SiteTypes = types.SetNull(types.StringType)
	data.SiteTypesVariable = types.StringNull()
	if t := res.Get(path + "siteType.optionType"); t.Exists() {
		va := res.Get(path + "siteType.value")
		if t.String() == "variable" {
			data.SiteTypesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SiteTypes = helpers.GetStringSet(va.Array())
		}
	}
	data.AffinityGroupNumber = types.Int64Null()
	data.AffinityGroupNumberVariable = types.StringNull()
	if t := res.Get(path + "affinityGroupNumber.optionType"); t.Exists() {
		va := res.Get(path + "affinityGroupNumber.value")
		if t.String() == "variable" {
			data.AffinityGroupNumberVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AffinityGroupNumber = types.Int64Value(va.Int())
		}
	}
	data.AffinityGroupPreferences = types.SetNull(types.Int64Type)
	data.AffinityGroupPreferencesVariable = types.StringNull()
	if t := res.Get(path + "affinityGroupPreference.optionType"); t.Exists() {
		va := res.Get(path + "affinityGroupPreference.value")
		if t.String() == "variable" {
			data.AffinityGroupPreferencesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AffinityGroupPreferences = helpers.GetInt64Set(va.Array())
		}
	}
	data.AffinityPreferenceAuto = types.BoolNull()
	data.AffinityPreferenceAutoVariable = types.StringNull()
	if t := res.Get(path + "affinityPreferenceAuto.optionType"); t.Exists() {
		va := res.Get(path + "affinityPreferenceAuto.value")
		if t.String() == "variable" {
			data.AffinityPreferenceAutoVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AffinityPreferenceAuto = types.BoolValue(va.Bool())
		}
	}
	if value := res.Get(path + "affinityPerVrf"); value.Exists() {
		data.AffinityPerVrfs = make([]SystemBasicAffinityPerVrfs, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := SystemBasicAffinityPerVrfs{}
			item.AffinityGroupNumber = types.Int64Null()
			item.AffinityGroupNumberVariable = types.StringNull()
			if t := v.Get("affinityGroupNumber.optionType"); t.Exists() {
				va := v.Get("affinityGroupNumber.value")
				if t.String() == "variable" {
					item.AffinityGroupNumberVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.AffinityGroupNumber = types.Int64Value(va.Int())
				}
			}
			item.VrfRange = types.StringNull()
			item.VrfRangeVariable = types.StringNull()
			if t := v.Get("vrfRange.optionType"); t.Exists() {
				va := v.Get("vrfRange.value")
				if t.String() == "variable" {
					item.VrfRangeVariable = types.StringValue(va.String())
				} else if t.String() == "global" {
					item.VrfRange = types.StringValue(va.String())
				}
			}
			data.AffinityPerVrfs = append(data.AffinityPerVrfs, item)
			return true
		})
	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *SystemBasic) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	data.Timezone = types.StringNull()
	data.TimezoneVariable = types.StringNull()
	if t := res.Get(path + "clock.timezone.optionType"); t.Exists() {
		va := res.Get(path + "clock.timezone.value")
		if t.String() == "variable" {
			data.TimezoneVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Timezone = types.StringValue(va.String())
		}
	}
	data.ConfigDescription = types.StringNull()
	data.ConfigDescriptionVariable = types.StringNull()
	if t := res.Get(path + "description.optionType"); t.Exists() {
		va := res.Get(path + "description.value")
		if t.String() == "variable" {
			data.ConfigDescriptionVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ConfigDescription = types.StringValue(va.String())
		}
	}
	data.Location = types.StringNull()
	data.LocationVariable = types.StringNull()
	if t := res.Get(path + "location.optionType"); t.Exists() {
		va := res.Get(path + "location.value")
		if t.String() == "variable" {
			data.LocationVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.Location = types.StringValue(va.String())
		}
	}
	data.GpsLongitude = types.Float64Null()
	data.GpsLongitudeVariable = types.StringNull()
	if t := res.Get(path + "gpsLocation.longitude.optionType"); t.Exists() {
		va := res.Get(path + "gpsLocation.longitude.value")
		if t.String() == "variable" {
			data.GpsLongitudeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GpsLongitude = types.Float64Value(va.Float())
		}
	}
	data.GpsLatitude = types.Float64Null()
	data.GpsLatitudeVariable = types.StringNull()
	if t := res.Get(path + "gpsLocation.latitude.optionType"); t.Exists() {
		va := res.Get(path + "gpsLocation.latitude.value")
		if t.String() == "variable" {
			data.GpsLatitudeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GpsLatitude = types.Float64Value(va.Float())
		}
	}
	data.GpsGeoFencingEnable = types.BoolNull()

	if t := res.Get(path + "gpsLocation.geoFencing.enable.optionType"); t.Exists() {
		va := res.Get(path + "gpsLocation.geoFencing.enable.value")
		if t.String() == "global" {
			data.GpsGeoFencingEnable = types.BoolValue(va.Bool())
		}
	}
	data.GpsGeoFencingRange = types.Int64Null()
	data.GpsGeoFencingRangeVariable = types.StringNull()
	if t := res.Get(path + "gpsLocation.geoFencing.range.optionType"); t.Exists() {
		va := res.Get(path + "gpsLocation.geoFencing.range.value")
		if t.String() == "variable" {
			data.GpsGeoFencingRangeVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.GpsGeoFencingRange = types.Int64Value(va.Int())
		}
	}
	data.GpsSmsEnable = types.BoolNull()

	if t := res.Get(path + "gpsLocation.geoFencing.sms.enable.optionType"); t.Exists() {
		va := res.Get(path + "gpsLocation.geoFencing.sms.enable.value")
		if t.String() == "global" {
			data.GpsSmsEnable = types.BoolValue(va.Bool())
		}
	}
	for i := range data.GpsSmsMobileNumbers {
		keys := [...]string{"number"}
		keyValues := [...]string{data.GpsSmsMobileNumbers[i].Number.ValueString()}
		keyValuesVariables := [...]string{data.GpsSmsMobileNumbers[i].NumberVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "gpsLocation.geoFencing.sms.mobileNumber").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.GpsSmsMobileNumbers[i].Number = types.StringNull()
		data.GpsSmsMobileNumbers[i].NumberVariable = types.StringNull()
		if t := r.Get("number.optionType"); t.Exists() {
			va := r.Get("number.value")
			if t.String() == "variable" {
				data.GpsSmsMobileNumbers[i].NumberVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.GpsSmsMobileNumbers[i].Number = types.StringValue(va.String())
			}
		}
	}
	data.DeviceGroups = types.SetNull(types.StringType)
	data.DeviceGroupsVariable = types.StringNull()
	if t := res.Get(path + "deviceGroups.optionType"); t.Exists() {
		va := res.Get(path + "deviceGroups.value")
		if t.String() == "variable" {
			data.DeviceGroupsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.DeviceGroups = helpers.GetStringSet(va.Array())
		}
	}
	data.ControllerGroups = types.SetNull(types.Int64Type)
	data.ControllerGroupsVariable = types.StringNull()
	if t := res.Get(path + "controllerGroupList.optionType"); t.Exists() {
		va := res.Get(path + "controllerGroupList.value")
		if t.String() == "variable" {
			data.ControllerGroupsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ControllerGroups = helpers.GetInt64Set(va.Array())
		}
	}
	data.OverlayId = types.Int64Null()
	data.OverlayIdVariable = types.StringNull()
	if t := res.Get(path + "overlayId.optionType"); t.Exists() {
		va := res.Get(path + "overlayId.value")
		if t.String() == "variable" {
			data.OverlayIdVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OverlayId = types.Int64Value(va.Int())
		}
	}
	data.PortOffset = types.Int64Null()
	data.PortOffsetVariable = types.StringNull()
	if t := res.Get(path + "portOffset.optionType"); t.Exists() {
		va := res.Get(path + "portOffset.value")
		if t.String() == "variable" {
			data.PortOffsetVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PortOffset = types.Int64Value(va.Int())
		}
	}
	data.PortHopping = types.BoolNull()
	data.PortHoppingVariable = types.StringNull()
	if t := res.Get(path + "portHop.optionType"); t.Exists() {
		va := res.Get(path + "portHop.value")
		if t.String() == "variable" {
			data.PortHoppingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.PortHopping = types.BoolValue(va.Bool())
		}
	}
	data.ControlSessionPps = types.Int64Null()
	data.ControlSessionPpsVariable = types.StringNull()
	if t := res.Get(path + "controlSessionPps.optionType"); t.Exists() {
		va := res.Get(path + "controlSessionPps.value")
		if t.String() == "variable" {
			data.ControlSessionPpsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ControlSessionPps = types.Int64Value(va.Int())
		}
	}
	data.TrackTransport = types.BoolNull()
	data.TrackTransportVariable = types.StringNull()
	if t := res.Get(path + "trackTransport.optionType"); t.Exists() {
		va := res.Get(path + "trackTransport.value")
		if t.String() == "variable" {
			data.TrackTransportVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackTransport = types.BoolValue(va.Bool())
		}
	}
	data.TrackInterfaceTag = types.Int64Null()
	data.TrackInterfaceTagVariable = types.StringNull()
	if t := res.Get(path + "trackInterfaceTag.optionType"); t.Exists() {
		va := res.Get(path + "trackInterfaceTag.value")
		if t.String() == "variable" {
			data.TrackInterfaceTagVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackInterfaceTag = types.Int64Value(va.Int())
		}
	}
	data.ConsoleBaudRate = types.StringNull()
	data.ConsoleBaudRateVariable = types.StringNull()
	if t := res.Get(path + "consoleBaudRate.optionType"); t.Exists() {
		va := res.Get(path + "consoleBaudRate.value")
		if t.String() == "variable" {
			data.ConsoleBaudRateVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.ConsoleBaudRate = types.StringValue(va.String())
		}
	}
	data.MaxOmpSessions = types.Int64Null()
	data.MaxOmpSessionsVariable = types.StringNull()
	if t := res.Get(path + "maxOmpSessions.optionType"); t.Exists() {
		va := res.Get(path + "maxOmpSessions.value")
		if t.String() == "variable" {
			data.MaxOmpSessionsVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MaxOmpSessions = types.Int64Value(va.Int())
		}
	}
	data.MultiTenant = types.BoolNull()
	data.MultiTenantVariable = types.StringNull()
	if t := res.Get(path + "multiTenant.optionType"); t.Exists() {
		va := res.Get(path + "multiTenant.value")
		if t.String() == "variable" {
			data.MultiTenantVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.MultiTenant = types.BoolValue(va.Bool())
		}
	}
	data.TrackDefaultGateway = types.BoolNull()
	data.TrackDefaultGatewayVariable = types.StringNull()
	if t := res.Get(path + "trackDefaultGateway.optionType"); t.Exists() {
		va := res.Get(path + "trackDefaultGateway.value")
		if t.String() == "variable" {
			data.TrackDefaultGatewayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TrackDefaultGateway = types.BoolValue(va.Bool())
		}
	}
	data.AdminTechOnFailure = types.BoolNull()
	data.AdminTechOnFailureVariable = types.StringNull()
	if t := res.Get(path + "adminTechOnFailure.optionType"); t.Exists() {
		va := res.Get(path + "adminTechOnFailure.value")
		if t.String() == "variable" {
			data.AdminTechOnFailureVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AdminTechOnFailure = types.BoolValue(va.Bool())
		}
	}
	data.IdleTimeout = types.Int64Null()
	data.IdleTimeoutVariable = types.StringNull()
	if t := res.Get(path + "idleTimeout.optionType"); t.Exists() {
		va := res.Get(path + "idleTimeout.value")
		if t.String() == "variable" {
			data.IdleTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.IdleTimeout = types.Int64Value(va.Int())
		}
	}
	data.OnDemandEnable = types.BoolNull()
	data.OnDemandEnableVariable = types.StringNull()
	if t := res.Get(path + "onDemand.onDemandEnable.optionType"); t.Exists() {
		va := res.Get(path + "onDemand.onDemandEnable.value")
		if t.String() == "variable" {
			data.OnDemandEnableVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OnDemandEnable = types.BoolValue(va.Bool())
		}
	}
	data.OnDemandIdleTimeout = types.Int64Null()
	data.OnDemandIdleTimeoutVariable = types.StringNull()
	if t := res.Get(path + "onDemand.onDemandIdleTimeout.optionType"); t.Exists() {
		va := res.Get(path + "onDemand.onDemandIdleTimeout.value")
		if t.String() == "variable" {
			data.OnDemandIdleTimeoutVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.OnDemandIdleTimeout = types.Int64Value(va.Int())
		}
	}
	data.TransportGateway = types.BoolNull()
	data.TransportGatewayVariable = types.StringNull()
	if t := res.Get(path + "transportGateway.optionType"); t.Exists() {
		va := res.Get(path + "transportGateway.value")
		if t.String() == "variable" {
			data.TransportGatewayVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.TransportGateway = types.BoolValue(va.Bool())
		}
	}
	data.EnhancedAppAwareRouting = types.StringNull()
	data.EnhancedAppAwareRoutingVariable = types.StringNull()
	if t := res.Get(path + "epfr.optionType"); t.Exists() {
		va := res.Get(path + "epfr.value")
		if t.String() == "variable" {
			data.EnhancedAppAwareRoutingVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.EnhancedAppAwareRouting = types.StringValue(va.String())
		}
	}
	data.SiteTypes = types.SetNull(types.StringType)
	data.SiteTypesVariable = types.StringNull()
	if t := res.Get(path + "siteType.optionType"); t.Exists() {
		va := res.Get(path + "siteType.value")
		if t.String() == "variable" {
			data.SiteTypesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.SiteTypes = helpers.GetStringSet(va.Array())
		}
	}
	data.AffinityGroupNumber = types.Int64Null()
	data.AffinityGroupNumberVariable = types.StringNull()
	if t := res.Get(path + "affinityGroupNumber.optionType"); t.Exists() {
		va := res.Get(path + "affinityGroupNumber.value")
		if t.String() == "variable" {
			data.AffinityGroupNumberVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AffinityGroupNumber = types.Int64Value(va.Int())
		}
	}
	data.AffinityGroupPreferences = types.SetNull(types.Int64Type)
	data.AffinityGroupPreferencesVariable = types.StringNull()
	if t := res.Get(path + "affinityGroupPreference.optionType"); t.Exists() {
		va := res.Get(path + "affinityGroupPreference.value")
		if t.String() == "variable" {
			data.AffinityGroupPreferencesVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AffinityGroupPreferences = helpers.GetInt64Set(va.Array())
		}
	}
	data.AffinityPreferenceAuto = types.BoolNull()
	data.AffinityPreferenceAutoVariable = types.StringNull()
	if t := res.Get(path + "affinityPreferenceAuto.optionType"); t.Exists() {
		va := res.Get(path + "affinityPreferenceAuto.value")
		if t.String() == "variable" {
			data.AffinityPreferenceAutoVariable = types.StringValue(va.String())
		} else if t.String() == "global" {
			data.AffinityPreferenceAuto = types.BoolValue(va.Bool())
		}
	}
	for i := range data.AffinityPerVrfs {
		keys := [...]string{"affinityGroupNumber"}
		keyValues := [...]string{strconv.FormatInt(data.AffinityPerVrfs[i].AffinityGroupNumber.ValueInt64(), 10)}
		keyValuesVariables := [...]string{data.AffinityPerVrfs[i].AffinityGroupNumberVariable.ValueString()}

		var r gjson.Result
		res.Get(path + "affinityPerVrf").ForEach(
			func(_, v gjson.Result) bool {
				found := false
				for ik := range keys {
					tt := v.Get(keys[ik] + ".optionType")
					vv := v.Get(keys[ik] + ".value")
					if tt.Exists() && vv.Exists() {
						if (tt.String() == "variable" && vv.String() == keyValuesVariables[ik]) || (tt.String() == "global" && vv.String() == keyValues[ik]) {
							found = true
							continue
						}
						found = false
						break
					}
					continue
				}
				if found {
					r = v
					return false
				}
				return true
			},
		)
		data.AffinityPerVrfs[i].AffinityGroupNumber = types.Int64Null()
		data.AffinityPerVrfs[i].AffinityGroupNumberVariable = types.StringNull()
		if t := r.Get("affinityGroupNumber.optionType"); t.Exists() {
			va := r.Get("affinityGroupNumber.value")
			if t.String() == "variable" {
				data.AffinityPerVrfs[i].AffinityGroupNumberVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.AffinityPerVrfs[i].AffinityGroupNumber = types.Int64Value(va.Int())
			}
		}
		data.AffinityPerVrfs[i].VrfRange = types.StringNull()
		data.AffinityPerVrfs[i].VrfRangeVariable = types.StringNull()
		if t := r.Get("vrfRange.optionType"); t.Exists() {
			va := r.Get("vrfRange.value")
			if t.String() == "variable" {
				data.AffinityPerVrfs[i].VrfRangeVariable = types.StringValue(va.String())
			} else if t.String() == "global" {
				data.AffinityPerVrfs[i].VrfRange = types.StringValue(va.String())
			}
		}
	}
}

// End of section. //template:end updateFromBody

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
	"strconv"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// End of section. //template:end imports

var MinCiscoSytemUpdateVersion = version.Must(version.NewVersion("20.12.0"))

// Section below is generated&owned by "gen/generator.go". //template:begin types
type CiscoSystem struct {
	Id                                types.String                           `tfsdk:"id"`
	Version                           types.Int64                            `tfsdk:"version"`
	TemplateType                      types.String                           `tfsdk:"template_type"`
	Name                              types.String                           `tfsdk:"name"`
	Description                       types.String                           `tfsdk:"description"`
	DeviceTypes                       types.Set                              `tfsdk:"device_types"`
	Timezone                          types.String                           `tfsdk:"timezone"`
	TimezoneVariable                  types.String                           `tfsdk:"timezone_variable"`
	Hostname                          types.String                           `tfsdk:"hostname"`
	HostnameVariable                  types.String                           `tfsdk:"hostname_variable"`
	SystemDescription                 types.String                           `tfsdk:"system_description"`
	SystemDescriptionVariable         types.String                           `tfsdk:"system_description_variable"`
	Location                          types.String                           `tfsdk:"location"`
	LocationVariable                  types.String                           `tfsdk:"location_variable"`
	Latitude                          types.Float64                          `tfsdk:"latitude"`
	LatitudeVariable                  types.String                           `tfsdk:"latitude_variable"`
	Longitude                         types.Float64                          `tfsdk:"longitude"`
	LongitudeVariable                 types.String                           `tfsdk:"longitude_variable"`
	GeoFencing                        types.Bool                             `tfsdk:"geo_fencing"`
	GeoFencingRange                   types.Int64                            `tfsdk:"geo_fencing_range"`
	GeoFencingRangeVariable           types.String                           `tfsdk:"geo_fencing_range_variable"`
	GeoFencingSms                     types.Bool                             `tfsdk:"geo_fencing_sms"`
	GeoFencingSmsPhoneNumbers         []CiscoSystemGeoFencingSmsPhoneNumbers `tfsdk:"geo_fencing_sms_phone_numbers"`
	DeviceGroups                      types.Set                              `tfsdk:"device_groups"`
	DeviceGroupsVariable              types.String                           `tfsdk:"device_groups_variable"`
	ControllerGroupList               types.List                             `tfsdk:"controller_group_list"`
	ControllerGroupListVariable       types.String                           `tfsdk:"controller_group_list_variable"`
	SystemIp                          types.String                           `tfsdk:"system_ip"`
	SystemIpVariable                  types.String                           `tfsdk:"system_ip_variable"`
	OverlayId                         types.Int64                            `tfsdk:"overlay_id"`
	OverlayIdVariable                 types.String                           `tfsdk:"overlay_id_variable"`
	SiteId                            types.Int64                            `tfsdk:"site_id"`
	SiteIdVariable                    types.String                           `tfsdk:"site_id_variable"`
	PortOffset                        types.Int64                            `tfsdk:"port_offset"`
	PortOffsetVariable                types.String                           `tfsdk:"port_offset_variable"`
	PortHopping                       types.Bool                             `tfsdk:"port_hopping"`
	PortHoppingVariable               types.String                           `tfsdk:"port_hopping_variable"`
	ControlSessionPps                 types.Int64                            `tfsdk:"control_session_pps"`
	ControlSessionPpsVariable         types.String                           `tfsdk:"control_session_pps_variable"`
	TrackTransport                    types.Bool                             `tfsdk:"track_transport"`
	TrackTransportVariable            types.String                           `tfsdk:"track_transport_variable"`
	TrackInterfaceTag                 types.Int64                            `tfsdk:"track_interface_tag"`
	TrackInterfaceTagVariable         types.String                           `tfsdk:"track_interface_tag_variable"`
	ConsoleBaudRate                   types.String                           `tfsdk:"console_baud_rate"`
	ConsoleBaudRateVariable           types.String                           `tfsdk:"console_baud_rate_variable"`
	MaxOmpSessions                    types.Int64                            `tfsdk:"max_omp_sessions"`
	MaxOmpSessionsVariable            types.String                           `tfsdk:"max_omp_sessions_variable"`
	MultiTenant                       types.Bool                             `tfsdk:"multi_tenant"`
	MultiTenantVariable               types.String                           `tfsdk:"multi_tenant_variable"`
	TrackDefaultGateway               types.Bool                             `tfsdk:"track_default_gateway"`
	TrackDefaultGatewayVariable       types.String                           `tfsdk:"track_default_gateway_variable"`
	AdminTechOnFailure                types.Bool                             `tfsdk:"admin_tech_on_failure"`
	AdminTechOnFailureVariable        types.String                           `tfsdk:"admin_tech_on_failure_variable"`
	IdleTimeout                       types.Int64                            `tfsdk:"idle_timeout"`
	IdleTimeoutVariable               types.String                           `tfsdk:"idle_timeout_variable"`
	Trackers                          []CiscoSystemTrackers                  `tfsdk:"trackers"`
	ObjectTrackers                    []CiscoSystemObjectTrackers            `tfsdk:"object_trackers"`
	OnDemandTunnel                    types.Bool                             `tfsdk:"on_demand_tunnel"`
	OnDemandTunnelVariable            types.String                           `tfsdk:"on_demand_tunnel_variable"`
	OnDemandTunnelIdleTimeout         types.Int64                            `tfsdk:"on_demand_tunnel_idle_timeout"`
	OnDemandTunnelIdleTimeoutVariable types.String                           `tfsdk:"on_demand_tunnel_idle_timeout_variable"`
	RegionId                          types.Int64                            `tfsdk:"region_id"`
	RegionIdVariable                  types.String                           `tfsdk:"region_id_variable"`
	SecondaryRegionId                 types.Int64                            `tfsdk:"secondary_region_id"`
	SecondaryRegionIdVariable         types.String                           `tfsdk:"secondary_region_id_variable"`
	Role                              types.String                           `tfsdk:"role"`
	RoleVariable                      types.String                           `tfsdk:"role_variable"`
	AffinityGroupNumber               types.Int64                            `tfsdk:"affinity_group_number"`
	AffinityGroupNumberVariable       types.String                           `tfsdk:"affinity_group_number_variable"`
	AffinityGroupPreference           types.List                             `tfsdk:"affinity_group_preference"`
	AffinityGroupPreferenceVariable   types.String                           `tfsdk:"affinity_group_preference_variable"`
	TransportGateway                  types.Bool                             `tfsdk:"transport_gateway"`
	TransportGatewayVariable          types.String                           `tfsdk:"transport_gateway_variable"`
	EnhancedAppAwareRouting           types.String                           `tfsdk:"enhanced_app_aware_routing"`
	EnableMrfMigration                types.String                           `tfsdk:"enable_mrf_migration"`
	MigrationBgpCommunity             types.Int64                            `tfsdk:"migration_bgp_community"`
}

type CiscoSystemGeoFencingSmsPhoneNumbers struct {
	Optional       types.Bool   `tfsdk:"optional"`
	Number         types.String `tfsdk:"number"`
	NumberVariable types.String `tfsdk:"number_variable"`
}

type CiscoSystemTrackers struct {
	Optional                types.Bool   `tfsdk:"optional"`
	Name                    types.String `tfsdk:"name"`
	NameVariable            types.String `tfsdk:"name_variable"`
	EndpointIp              types.String `tfsdk:"endpoint_ip"`
	EndpointIpVariable      types.String `tfsdk:"endpoint_ip_variable"`
	EndpointDnsName         types.String `tfsdk:"endpoint_dns_name"`
	EndpointDnsNameVariable types.String `tfsdk:"endpoint_dns_name_variable"`
	EndpointApiUrl          types.String `tfsdk:"endpoint_api_url"`
	EndpointApiUrlVariable  types.String `tfsdk:"endpoint_api_url_variable"`
	Elements                types.Set    `tfsdk:"elements"`
	ElementsVariable        types.String `tfsdk:"elements_variable"`
	Boolean                 types.String `tfsdk:"boolean"`
	BooleanVariable         types.String `tfsdk:"boolean_variable"`
	Threshold               types.Int64  `tfsdk:"threshold"`
	ThresholdVariable       types.String `tfsdk:"threshold_variable"`
	Interval                types.Int64  `tfsdk:"interval"`
	IntervalVariable        types.String `tfsdk:"interval_variable"`
	Multiplier              types.Int64  `tfsdk:"multiplier"`
	MultiplierVariable      types.String `tfsdk:"multiplier_variable"`
	Type                    types.String `tfsdk:"type"`
	TypeVariable            types.String `tfsdk:"type_variable"`
}

type CiscoSystemObjectTrackers struct {
	Optional             types.Bool                                `tfsdk:"optional"`
	ObjectNumber         types.Int64                               `tfsdk:"object_number"`
	ObjectNumberVariable types.String                              `tfsdk:"object_number_variable"`
	Interface            types.String                              `tfsdk:"interface"`
	InterfaceVariable    types.String                              `tfsdk:"interface_variable"`
	Sig                  types.String                              `tfsdk:"sig"`
	SigVariable          types.String                              `tfsdk:"sig_variable"`
	Ip                   types.String                              `tfsdk:"ip"`
	IpVariable           types.String                              `tfsdk:"ip_variable"`
	Mask                 types.String                              `tfsdk:"mask"`
	MaskVariable         types.String                              `tfsdk:"mask_variable"`
	VpnId                types.Int64                               `tfsdk:"vpn_id"`
	GroupTracksIds       []CiscoSystemObjectTrackersGroupTracksIds `tfsdk:"group_tracks_ids"`
	Boolean              types.String                              `tfsdk:"boolean"`
	BooleanVariable      types.String                              `tfsdk:"boolean_variable"`
}

type CiscoSystemObjectTrackersGroupTracksIds struct {
	Optional        types.Bool   `tfsdk:"optional"`
	TrackId         types.Int64  `tfsdk:"track_id"`
	TrackIdVariable types.String `tfsdk:"track_id_variable"`
}

// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data CiscoSystem) getModel() string {
	return "cisco_system"
}

// End of section. //template:end getModel
func (data CiscoSystem) toBody(ctx context.Context, version *version.Version) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_system")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."

	if !data.TimezoneVariable.IsNull() {
		body, _ = sjson.Set(body, path+"clock.timezone."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"clock.timezone."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"clock.timezone."+"vipVariableName", data.TimezoneVariable.ValueString())
	} else if data.Timezone.IsNull() {
		body, _ = sjson.Set(body, path+"clock.timezone."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"clock.timezone."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"clock.timezone."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"clock.timezone."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"clock.timezone."+"vipValue", data.Timezone.ValueString())
	}

	if !data.HostnameVariable.IsNull() {
		body, _ = sjson.Set(body, path+"host-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"host-name."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"host-name."+"vipVariableName", data.HostnameVariable.ValueString())
	} else if data.Hostname.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"host-name."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"host-name."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"host-name."+"vipValue", data.Hostname.ValueString())
	}

	if !data.SystemDescriptionVariable.IsNull() {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"description."+"vipVariableName", data.SystemDescriptionVariable.ValueString())
	} else if data.SystemDescription.IsNull() {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"description."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"description."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"description."+"vipValue", data.SystemDescription.ValueString())
	}

	if !data.LocationVariable.IsNull() {
		body, _ = sjson.Set(body, path+"location."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"location."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"location."+"vipVariableName", data.LocationVariable.ValueString())
	} else if data.Location.IsNull() {
		body, _ = sjson.Set(body, path+"location."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"location."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"location."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"location."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"location."+"vipValue", data.Location.ValueString())
	}

	if !data.LatitudeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"gps-location.latitude."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps-location.latitude."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"gps-location.latitude."+"vipVariableName", data.LatitudeVariable.ValueString())
	} else if data.Latitude.IsNull() {
		body, _ = sjson.Set(body, path+"gps-location.latitude."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps-location.latitude."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"gps-location.latitude."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps-location.latitude."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"gps-location.latitude."+"vipValue", data.Latitude.ValueFloat64())
	}

	if !data.LongitudeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"gps-location.longitude."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps-location.longitude."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"gps-location.longitude."+"vipVariableName", data.LongitudeVariable.ValueString())
	} else if data.Longitude.IsNull() {
		body, _ = sjson.Set(body, path+"gps-location.longitude."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps-location.longitude."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"gps-location.longitude."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps-location.longitude."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"gps-location.longitude."+"vipValue", data.Longitude.ValueFloat64())
	}
	if data.GeoFencing.IsNull() {
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.enable."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.enable."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.enable."+"vipValue", strconv.FormatBool(data.GeoFencing.ValueBool()))
	}

	if !data.GeoFencingRangeVariable.IsNull() {
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.range."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.range."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.range."+"vipVariableName", data.GeoFencingRangeVariable.ValueString())
	} else if data.GeoFencingRange.IsNull() {
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.range."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.range."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.range."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.range."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.range."+"vipValue", data.GeoFencingRange.ValueInt64())
	}
	if data.GeoFencingSms.IsNull() {
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.sms.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.sms.enable."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.sms.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.sms.enable."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.sms.enable."+"vipValue", strconv.FormatBool(data.GeoFencingSms.ValueBool()))
	}
	if len(data.GeoFencingSmsPhoneNumbers) > 0 {
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.sms.mobile-number."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.sms.mobile-number."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.sms.mobile-number."+"vipPrimaryKey", []string{"number"})
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.sms.mobile-number."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.sms.mobile-number."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.sms.mobile-number."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.sms.mobile-number."+"vipPrimaryKey", []string{"number"})
		body, _ = sjson.Set(body, path+"gps-location.geo-fencing.sms.mobile-number."+"vipValue", []interface{}{})
	}
	for _, item := range data.GeoFencingSmsPhoneNumbers {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "number")

		if !item.NumberVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "number."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "number."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "number."+"vipVariableName", item.NumberVariable.ValueString())
		} else if item.Number.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "number."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "number."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "number."+"vipValue", item.Number.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"gps-location.geo-fencing.sms.mobile-number."+"vipValue.-1", itemBody)
	}

	if !data.DeviceGroupsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"device-groups."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"device-groups."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"device-groups."+"vipVariableName", data.DeviceGroupsVariable.ValueString())
	} else if data.DeviceGroups.IsNull() {
		body, _ = sjson.Set(body, path+"device-groups."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"device-groups."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"device-groups."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"device-groups."+"vipType", "constant")
		var values []string
		data.DeviceGroups.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"device-groups."+"vipValue", values)
	}

	if !data.ControllerGroupListVariable.IsNull() {
		body, _ = sjson.Set(body, path+"controller-group-list."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"controller-group-list."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"controller-group-list."+"vipVariableName", data.ControllerGroupListVariable.ValueString())
	} else if data.ControllerGroupList.IsNull() {
		body, _ = sjson.Set(body, path+"controller-group-list."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"controller-group-list."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"controller-group-list."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"controller-group-list."+"vipType", "constant")
		var values []int64
		data.ControllerGroupList.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"controller-group-list."+"vipValue", values)
	}

	if !data.SystemIpVariable.IsNull() {
		body, _ = sjson.Set(body, path+"system-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"system-ip."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"system-ip."+"vipVariableName", data.SystemIpVariable.ValueString())
	} else if data.SystemIp.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"system-ip."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"system-ip."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"system-ip."+"vipValue", data.SystemIp.ValueString())
	}

	if !data.OverlayIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"overlay-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"overlay-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"overlay-id."+"vipVariableName", data.OverlayIdVariable.ValueString())
	} else if data.OverlayId.IsNull() {
		body, _ = sjson.Set(body, path+"overlay-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"overlay-id."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"overlay-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"overlay-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"overlay-id."+"vipValue", data.OverlayId.ValueInt64())
	}

	if !data.SiteIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"site-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"site-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"site-id."+"vipVariableName", data.SiteIdVariable.ValueString())
	} else if data.SiteId.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"site-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"site-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"site-id."+"vipValue", data.SiteId.ValueInt64())
	}

	if !data.PortOffsetVariable.IsNull() {
		body, _ = sjson.Set(body, path+"port-offset."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"port-offset."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"port-offset."+"vipVariableName", data.PortOffsetVariable.ValueString())
	} else if data.PortOffset.IsNull() {
		body, _ = sjson.Set(body, path+"port-offset."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"port-offset."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"port-offset."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"port-offset."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"port-offset."+"vipValue", data.PortOffset.ValueInt64())
	}

	if !data.PortHoppingVariable.IsNull() {
		body, _ = sjson.Set(body, path+"port-hop."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"port-hop."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"port-hop."+"vipVariableName", data.PortHoppingVariable.ValueString())
	} else if data.PortHopping.IsNull() {
		body, _ = sjson.Set(body, path+"port-hop."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"port-hop."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"port-hop."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"port-hop."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"port-hop."+"vipValue", strconv.FormatBool(data.PortHopping.ValueBool()))
	}

	if !data.ControlSessionPpsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"control-session-pps."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"control-session-pps."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"control-session-pps."+"vipVariableName", data.ControlSessionPpsVariable.ValueString())
	} else if data.ControlSessionPps.IsNull() {
		body, _ = sjson.Set(body, path+"control-session-pps."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"control-session-pps."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"control-session-pps."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"control-session-pps."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"control-session-pps."+"vipValue", data.ControlSessionPps.ValueInt64())
	}

	if !data.TrackTransportVariable.IsNull() {
		body, _ = sjson.Set(body, path+"track-transport."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"track-transport."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"track-transport."+"vipVariableName", data.TrackTransportVariable.ValueString())
	} else if data.TrackTransport.IsNull() {
		body, _ = sjson.Set(body, path+"track-transport."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"track-transport."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"track-transport."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"track-transport."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"track-transport."+"vipValue", strconv.FormatBool(data.TrackTransport.ValueBool()))
	}

	if !data.TrackInterfaceTagVariable.IsNull() {
		body, _ = sjson.Set(body, path+"track-interface-tag."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"track-interface-tag."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"track-interface-tag."+"vipVariableName", data.TrackInterfaceTagVariable.ValueString())
	} else if data.TrackInterfaceTag.IsNull() {
		body, _ = sjson.Set(body, path+"track-interface-tag."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"track-interface-tag."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"track-interface-tag."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"track-interface-tag."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"track-interface-tag."+"vipValue", data.TrackInterfaceTag.ValueInt64())
	}

	if !data.ConsoleBaudRateVariable.IsNull() {
		body, _ = sjson.Set(body, path+"console-baud-rate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"console-baud-rate."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"console-baud-rate."+"vipVariableName", data.ConsoleBaudRateVariable.ValueString())
	} else if data.ConsoleBaudRate.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"console-baud-rate."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"console-baud-rate."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"console-baud-rate."+"vipValue", data.ConsoleBaudRate.ValueString())
	}

	if !data.MaxOmpSessionsVariable.IsNull() {
		body, _ = sjson.Set(body, path+"max-omp-sessions."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"max-omp-sessions."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"max-omp-sessions."+"vipVariableName", data.MaxOmpSessionsVariable.ValueString())
	} else if data.MaxOmpSessions.IsNull() {
		body, _ = sjson.Set(body, path+"max-omp-sessions."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"max-omp-sessions."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"max-omp-sessions."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"max-omp-sessions."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"max-omp-sessions."+"vipValue", data.MaxOmpSessions.ValueInt64())
	}

	if !data.MultiTenantVariable.IsNull() {
		body, _ = sjson.Set(body, path+"multi-tenant."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"multi-tenant."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"multi-tenant."+"vipVariableName", data.MultiTenantVariable.ValueString())
	} else if data.MultiTenant.IsNull() {
	} else {
		body, _ = sjson.Set(body, path+"multi-tenant."+"vipObjectType", "node-only")
		body, _ = sjson.Set(body, path+"multi-tenant."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"multi-tenant."+"vipValue", strconv.FormatBool(data.MultiTenant.ValueBool()))
	}

	if !data.TrackDefaultGatewayVariable.IsNull() {
		body, _ = sjson.Set(body, path+"track-default-gateway."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"track-default-gateway."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"track-default-gateway."+"vipVariableName", data.TrackDefaultGatewayVariable.ValueString())
	} else if data.TrackDefaultGateway.IsNull() {
		body, _ = sjson.Set(body, path+"track-default-gateway."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"track-default-gateway."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"track-default-gateway."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"track-default-gateway."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"track-default-gateway."+"vipValue", strconv.FormatBool(data.TrackDefaultGateway.ValueBool()))
	}

	if !data.AdminTechOnFailureVariable.IsNull() {
		body, _ = sjson.Set(body, path+"admin-tech-on-failure."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"admin-tech-on-failure."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"admin-tech-on-failure."+"vipVariableName", data.AdminTechOnFailureVariable.ValueString())
	} else if data.AdminTechOnFailure.IsNull() {
		body, _ = sjson.Set(body, path+"admin-tech-on-failure."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"admin-tech-on-failure."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"admin-tech-on-failure."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"admin-tech-on-failure."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"admin-tech-on-failure."+"vipValue", strconv.FormatBool(data.AdminTechOnFailure.ValueBool()))
	}

	if !data.IdleTimeoutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"idle-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"idle-timeout."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"idle-timeout."+"vipVariableName", data.IdleTimeoutVariable.ValueString())
	} else if data.IdleTimeout.IsNull() {
		body, _ = sjson.Set(body, path+"idle-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"idle-timeout."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"idle-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"idle-timeout."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"idle-timeout."+"vipValue", data.IdleTimeout.ValueInt64())
	}
	if len(data.Trackers) > 0 {
		body, _ = sjson.Set(body, path+"tracker."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"tracker."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"tracker."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"tracker."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"tracker."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"tracker."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"tracker."+"vipPrimaryKey", []string{"name"})
		body, _ = sjson.Set(body, path+"tracker."+"vipValue", []interface{}{})
	}
	for _, item := range data.Trackers {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "name")

		if !item.NameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipVariableName", item.NameVariable.ValueString())
		} else if item.Name.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "name."+"vipValue", item.Name.ValueString())
		}
		itemAttributes = append(itemAttributes, "endpoint-ip")

		if !item.EndpointIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "endpoint-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "endpoint-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "endpoint-ip."+"vipVariableName", item.EndpointIpVariable.ValueString())
		} else if item.EndpointIp.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "endpoint-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "endpoint-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "endpoint-ip."+"vipValue", item.EndpointIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "endpoint-dns-name")

		if !item.EndpointDnsNameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "endpoint-dns-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "endpoint-dns-name."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "endpoint-dns-name."+"vipVariableName", item.EndpointDnsNameVariable.ValueString())
		} else if item.EndpointDnsName.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "endpoint-dns-name."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "endpoint-dns-name."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "endpoint-dns-name."+"vipValue", item.EndpointDnsName.ValueString())
		}
		itemAttributes = append(itemAttributes, "endpoint-api-url")

		if !item.EndpointApiUrlVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "endpoint-api-url."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "endpoint-api-url."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "endpoint-api-url."+"vipVariableName", item.EndpointApiUrlVariable.ValueString())
		} else if item.EndpointApiUrl.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "endpoint-api-url."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "endpoint-api-url."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "endpoint-api-url."+"vipValue", item.EndpointApiUrl.ValueString())
		}
		itemAttributes = append(itemAttributes, "elements")

		if !item.ElementsVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "elements."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "elements."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "elements."+"vipVariableName", item.ElementsVariable.ValueString())
		} else if item.Elements.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "elements."+"vipObjectType", "list")
			itemBody, _ = sjson.Set(itemBody, "elements."+"vipType", "constant")
			var values []string
			item.Elements.ElementsAs(ctx, &values, false)
			itemBody, _ = sjson.Set(itemBody, "elements."+"vipValue", values)
		}
		itemAttributes = append(itemAttributes, "boolean")

		if !item.BooleanVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "boolean."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "boolean."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "boolean."+"vipVariableName", item.BooleanVariable.ValueString())
		} else if item.Boolean.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "boolean."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "boolean."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "boolean."+"vipValue", item.Boolean.ValueString())
		}
		itemAttributes = append(itemAttributes, "threshold")

		if !item.ThresholdVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipVariableName", item.ThresholdVariable.ValueString())
		} else if item.Threshold.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "threshold."+"vipValue", item.Threshold.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "interval")

		if !item.IntervalVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipVariableName", item.IntervalVariable.ValueString())
		} else if item.Interval.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "interval."+"vipValue", item.Interval.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "multiplier")

		if !item.MultiplierVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipVariableName", item.MultiplierVariable.ValueString())
		} else if item.Multiplier.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "multiplier."+"vipValue", item.Multiplier.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "type")

		if !item.TypeVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "type."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "type."+"vipVariableName", item.TypeVariable.ValueString())
		} else if item.Type.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "type."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "type."+"vipValue", item.Type.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"tracker."+"vipValue.-1", itemBody)
	}
	if len(data.ObjectTrackers) > 0 {
		body, _ = sjson.Set(body, path+"object-track."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"object-track."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"object-track."+"vipPrimaryKey", []string{"object-number"})
		body, _ = sjson.Set(body, path+"object-track."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"object-track."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"object-track."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"object-track."+"vipPrimaryKey", []string{"object-number"})
		body, _ = sjson.Set(body, path+"object-track."+"vipValue", []interface{}{})
	}
	for _, item := range data.ObjectTrackers {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "object-number")

		if !item.ObjectNumberVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "object-number."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "object-number."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "object-number."+"vipVariableName", item.ObjectNumberVariable.ValueString())
		} else if item.ObjectNumber.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "object-number."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "object-number."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "object-number."+"vipValue", item.ObjectNumber.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "interface")

		if !item.InterfaceVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipVariableName", item.InterfaceVariable.ValueString())
		} else if item.Interface.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "interface."+"vipValue", item.Interface.ValueString())
		}
		itemAttributes = append(itemAttributes, "sig")

		if !item.SigVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "sig."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "sig."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "sig."+"vipVariableName", item.SigVariable.ValueString())
		} else if item.Sig.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "sig."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "sig."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "sig."+"vipValue", item.Sig.ValueString())
		}
		itemAttributes = append(itemAttributes, "ip")

		if !item.IpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipVariableName", item.IpVariable.ValueString())
		} else if item.Ip.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "ip."+"vipValue", item.Ip.ValueString())
		}
		itemAttributes = append(itemAttributes, "mask")

		if !item.MaskVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "mask."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "mask."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "mask."+"vipVariableName", item.MaskVariable.ValueString())
		} else if item.Mask.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "mask."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "mask."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "mask."+"vipValue", item.Mask.ValueString())
		}
		itemAttributes = append(itemAttributes, "vpn")
		if item.VpnId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "vpn."+"vipValue", item.VpnId.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "object")
		if len(item.GroupTracksIds) > 0 {
			itemBody, _ = sjson.Set(itemBody, "object."+"vipObjectType", "tree")
			itemBody, _ = sjson.Set(itemBody, "object."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "object."+"vipPrimaryKey", []string{"number"})
			itemBody, _ = sjson.Set(itemBody, "object."+"vipValue", []interface{}{})
		} else {
		}
		for _, childItem := range item.GroupTracksIds {
			itemChildBody := ""
			itemChildAttributes := make([]string, 0)
			itemChildAttributes = append(itemChildAttributes, "number")

			if !childItem.TrackIdVariable.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "number."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "number."+"vipType", "variableName")
				itemChildBody, _ = sjson.Set(itemChildBody, "number."+"vipVariableName", childItem.TrackIdVariable.ValueString())
			} else if childItem.TrackId.IsNull() {
			} else {
				itemChildBody, _ = sjson.Set(itemChildBody, "number."+"vipObjectType", "object")
				itemChildBody, _ = sjson.Set(itemChildBody, "number."+"vipType", "constant")
				itemChildBody, _ = sjson.Set(itemChildBody, "number."+"vipValue", childItem.TrackId.ValueInt64())
			}
			if !childItem.Optional.IsNull() {
				itemChildBody, _ = sjson.Set(itemChildBody, "vipOptional", childItem.Optional.ValueBool())
				itemChildBody, _ = sjson.Set(itemChildBody, "priority-order", itemChildAttributes)
			}
			itemBody, _ = sjson.SetRaw(itemBody, "object."+"vipValue.-1", itemChildBody)
		}
		itemAttributes = append(itemAttributes, "boolean")

		if !item.BooleanVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "boolean."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "boolean."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "boolean."+"vipVariableName", item.BooleanVariable.ValueString())
		} else if item.Boolean.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "boolean."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "boolean."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "boolean."+"vipValue", item.Boolean.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"object-track."+"vipValue.-1", itemBody)
	}

	if !data.OnDemandTunnelVariable.IsNull() {
		body, _ = sjson.Set(body, path+"on-demand.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"on-demand.enable."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"on-demand.enable."+"vipVariableName", data.OnDemandTunnelVariable.ValueString())
	} else if data.OnDemandTunnel.IsNull() {
		body, _ = sjson.Set(body, path+"on-demand.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"on-demand.enable."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"on-demand.enable."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"on-demand.enable."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"on-demand.enable."+"vipValue", strconv.FormatBool(data.OnDemandTunnel.ValueBool()))
	}

	if !data.OnDemandTunnelIdleTimeoutVariable.IsNull() {
		body, _ = sjson.Set(body, path+"on-demand.idle-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"on-demand.idle-timeout."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"on-demand.idle-timeout."+"vipVariableName", data.OnDemandTunnelIdleTimeoutVariable.ValueString())
	} else if data.OnDemandTunnelIdleTimeout.IsNull() {
		body, _ = sjson.Set(body, path+"on-demand.idle-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"on-demand.idle-timeout."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"on-demand.idle-timeout."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"on-demand.idle-timeout."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"on-demand.idle-timeout."+"vipValue", data.OnDemandTunnelIdleTimeout.ValueInt64())
	}

	if !data.RegionIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"region-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"region-id."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"region-id."+"vipVariableName", data.RegionIdVariable.ValueString())
	} else if data.RegionId.IsNull() {
		body, _ = sjson.Set(body, path+"region-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"region-id."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"region-id."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"region-id."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"region-id."+"vipValue", data.RegionId.ValueInt64())
	}

	if !data.SecondaryRegionIdVariable.IsNull() {
		body, _ = sjson.Set(body, path+"secondary-region."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"secondary-region."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"secondary-region."+"vipVariableName", data.SecondaryRegionIdVariable.ValueString())
	} else if data.SecondaryRegionId.IsNull() {
		body, _ = sjson.Set(body, path+"secondary-region."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"secondary-region."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"secondary-region."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"secondary-region."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"secondary-region."+"vipValue", data.SecondaryRegionId.ValueInt64())
	}

	if !data.RoleVariable.IsNull() {
		body, _ = sjson.Set(body, path+"role."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"role."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"role."+"vipVariableName", data.RoleVariable.ValueString())
	} else if data.Role.IsNull() {
		body, _ = sjson.Set(body, path+"role."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"role."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"role."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"role."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"role."+"vipValue", data.Role.ValueString())
	}

	if !data.AffinityGroupNumberVariable.IsNull() {
		body, _ = sjson.Set(body, path+"affinity-group.affinity-group-number."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"affinity-group.affinity-group-number."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"affinity-group.affinity-group-number."+"vipVariableName", data.AffinityGroupNumberVariable.ValueString())
	} else if data.AffinityGroupNumber.IsNull() {
		body, _ = sjson.Set(body, path+"affinity-group.affinity-group-number."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"affinity-group.affinity-group-number."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"affinity-group.affinity-group-number."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"affinity-group.affinity-group-number."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"affinity-group.affinity-group-number."+"vipValue", data.AffinityGroupNumber.ValueInt64())
	}

	if !data.AffinityGroupPreferenceVariable.IsNull() {
		body, _ = sjson.Set(body, path+"affinity-group.preference."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"affinity-group.preference."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"affinity-group.preference."+"vipVariableName", data.AffinityGroupPreferenceVariable.ValueString())
	} else if data.AffinityGroupPreference.IsNull() {
		body, _ = sjson.Set(body, path+"affinity-group.preference."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"affinity-group.preference."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"affinity-group.preference."+"vipObjectType", "list")
		body, _ = sjson.Set(body, path+"affinity-group.preference."+"vipType", "constant")
		var values []int64
		data.AffinityGroupPreference.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"affinity-group.preference."+"vipValue", values)
	}

	if !data.TransportGatewayVariable.IsNull() {
		body, _ = sjson.Set(body, path+"transport-gateway."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"transport-gateway."+"vipType", "variableName")
		body, _ = sjson.Set(body, path+"transport-gateway."+"vipVariableName", data.TransportGatewayVariable.ValueString())
	} else if data.TransportGateway.IsNull() {
		body, _ = sjson.Set(body, path+"transport-gateway."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"transport-gateway."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"transport-gateway."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"transport-gateway."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"transport-gateway."+"vipValue", strconv.FormatBool(data.TransportGateway.ValueBool()))
	}
	if version.LessThan(MinCiscoSytemUpdateVersion) {
	} else {
		if data.EnhancedAppAwareRouting.IsNull() {
			body, _ = sjson.Set(body, path+"epfr."+"vipObjectType", "object")
			body, _ = sjson.Set(body, path+"epfr."+"vipType", "ignore")
		} else {
			body, _ = sjson.Set(body, path+"epfr."+"vipObjectType", "object")
			body, _ = sjson.Set(body, path+"epfr."+"vipType", "constant")
			body, _ = sjson.Set(body, path+"epfr."+"vipValue", data.EnhancedAppAwareRouting.ValueString())
		}
	}
	if data.EnableMrfMigration.IsNull() {
		body, _ = sjson.Set(body, path+"enable-mrf-migration."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"enable-mrf-migration."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"enable-mrf-migration."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"enable-mrf-migration."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"enable-mrf-migration."+"vipValue", data.EnableMrfMigration.ValueString())
	}
	if data.MigrationBgpCommunity.IsNull() {
		body, _ = sjson.Set(body, path+"migration-bgp-community."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"migration-bgp-community."+"vipType", "ignore")
	} else {
		body, _ = sjson.Set(body, path+"migration-bgp-community."+"vipObjectType", "object")
		body, _ = sjson.Set(body, path+"migration-bgp-community."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"migration-bgp-community."+"vipValue", data.MigrationBgpCommunity.ValueInt64())
	}
	return body
}

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *CiscoSystem) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceTypes = helpers.GetStringSet(value.Array())
	} else {
		data.DeviceTypes = types.SetNull(types.StringType)
	}
	if value := res.Get("templateDescription"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateType"); value.Exists() {
		data.TemplateType = types.StringValue(value.String())
	} else {
		data.TemplateType = types.StringNull()
	}

	path := "templateDefinition."
	if value := res.Get(path + "clock.timezone.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Timezone = types.StringNull()

			v := res.Get(path + "clock.timezone.vipVariableName")
			data.TimezoneVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Timezone = types.StringNull()
			data.TimezoneVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "clock.timezone.vipValue")
			data.Timezone = types.StringValue(v.String())
			data.TimezoneVariable = types.StringNull()
		}
	} else {
		data.Timezone = types.StringNull()
		data.TimezoneVariable = types.StringNull()
	}
	if value := res.Get(path + "host-name.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Hostname = types.StringNull()

			v := res.Get(path + "host-name.vipVariableName")
			data.HostnameVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Hostname = types.StringNull()
			data.HostnameVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "host-name.vipValue")
			data.Hostname = types.StringValue(v.String())
			data.HostnameVariable = types.StringNull()
		}
	} else {
		data.Hostname = types.StringNull()
		data.HostnameVariable = types.StringNull()
	}
	if value := res.Get(path + "description.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SystemDescription = types.StringNull()

			v := res.Get(path + "description.vipVariableName")
			data.SystemDescriptionVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SystemDescription = types.StringNull()
			data.SystemDescriptionVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "description.vipValue")
			data.SystemDescription = types.StringValue(v.String())
			data.SystemDescriptionVariable = types.StringNull()
		}
	} else {
		data.SystemDescription = types.StringNull()
		data.SystemDescriptionVariable = types.StringNull()
	}
	if value := res.Get(path + "location.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Location = types.StringNull()

			v := res.Get(path + "location.vipVariableName")
			data.LocationVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Location = types.StringNull()
			data.LocationVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "location.vipValue")
			data.Location = types.StringValue(v.String())
			data.LocationVariable = types.StringNull()
		}
	} else {
		data.Location = types.StringNull()
		data.LocationVariable = types.StringNull()
	}
	if value := res.Get(path + "gps-location.latitude.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Latitude = types.Float64Null()

			v := res.Get(path + "gps-location.latitude.vipVariableName")
			data.LatitudeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Latitude = types.Float64Null()
			data.LatitudeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "gps-location.latitude.vipValue")
			data.Latitude = types.Float64Value(v.Float())
			data.LatitudeVariable = types.StringNull()
		}
	} else {
		data.Latitude = types.Float64Null()
		data.LatitudeVariable = types.StringNull()
	}
	if value := res.Get(path + "gps-location.longitude.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Longitude = types.Float64Null()

			v := res.Get(path + "gps-location.longitude.vipVariableName")
			data.LongitudeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Longitude = types.Float64Null()
			data.LongitudeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "gps-location.longitude.vipValue")
			data.Longitude = types.Float64Value(v.Float())
			data.LongitudeVariable = types.StringNull()
		}
	} else {
		data.Longitude = types.Float64Null()
		data.LongitudeVariable = types.StringNull()
	}
	if value := res.Get(path + "gps-location.geo-fencing.enable.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.GeoFencing = types.BoolNull()

		} else if value.String() == "ignore" {
			data.GeoFencing = types.BoolNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "gps-location.geo-fencing.enable.vipValue")
			data.GeoFencing = types.BoolValue(v.Bool())

		}
	} else {
		data.GeoFencing = types.BoolNull()

	}
	if value := res.Get(path + "gps-location.geo-fencing.range.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.GeoFencingRange = types.Int64Null()

			v := res.Get(path + "gps-location.geo-fencing.range.vipVariableName")
			data.GeoFencingRangeVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.GeoFencingRange = types.Int64Null()
			data.GeoFencingRangeVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "gps-location.geo-fencing.range.vipValue")
			data.GeoFencingRange = types.Int64Value(v.Int())
			data.GeoFencingRangeVariable = types.StringNull()
		}
	} else {
		data.GeoFencingRange = types.Int64Null()
		data.GeoFencingRangeVariable = types.StringNull()
	}
	if value := res.Get(path + "gps-location.geo-fencing.sms.enable.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.GeoFencingSms = types.BoolNull()

		} else if value.String() == "ignore" {
			data.GeoFencingSms = types.BoolNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "gps-location.geo-fencing.sms.enable.vipValue")
			data.GeoFencingSms = types.BoolValue(v.Bool())

		}
	} else {
		data.GeoFencingSms = types.BoolNull()

	}
	if value := res.Get(path + "gps-location.geo-fencing.sms.mobile-number.vipValue"); len(value.Array()) > 0 {
		data.GeoFencingSmsPhoneNumbers = make([]CiscoSystemGeoFencingSmsPhoneNumbers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoSystemGeoFencingSmsPhoneNumbers{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("number.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Number = types.StringNull()

					cv := v.Get("number.vipVariableName")
					item.NumberVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Number = types.StringNull()
					item.NumberVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("number.vipValue")
					item.Number = types.StringValue(cv.String())
					item.NumberVariable = types.StringNull()
				}
			} else {
				item.Number = types.StringNull()
				item.NumberVariable = types.StringNull()
			}
			data.GeoFencingSmsPhoneNumbers = append(data.GeoFencingSmsPhoneNumbers, item)
			return true
		})
	} else {
		if len(data.GeoFencingSmsPhoneNumbers) > 0 {
			data.GeoFencingSmsPhoneNumbers = []CiscoSystemGeoFencingSmsPhoneNumbers{}
		}
	}
	if value := res.Get(path + "device-groups.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.DeviceGroups = types.SetNull(types.StringType)

			v := res.Get(path + "device-groups.vipVariableName")
			data.DeviceGroupsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.DeviceGroups = types.SetNull(types.StringType)
			data.DeviceGroupsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "device-groups.vipValue")
			data.DeviceGroups = helpers.GetStringSet(v.Array())
			data.DeviceGroupsVariable = types.StringNull()
		}
	} else {
		data.DeviceGroups = types.SetNull(types.StringType)
		data.DeviceGroupsVariable = types.StringNull()
	}
	if value := res.Get(path + "controller-group-list.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.ControllerGroupList = types.ListNull(types.Int64Type)

			v := res.Get(path + "controller-group-list.vipVariableName")
			data.ControllerGroupListVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ControllerGroupList = types.ListNull(types.Int64Type)
			data.ControllerGroupListVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "controller-group-list.vipValue")
			data.ControllerGroupList = helpers.GetInt64List(v.Array())
			data.ControllerGroupListVariable = types.StringNull()
		}
	} else {
		data.ControllerGroupList = types.ListNull(types.Int64Type)
		data.ControllerGroupListVariable = types.StringNull()
	}
	if value := res.Get(path + "system-ip.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SystemIp = types.StringNull()

			v := res.Get(path + "system-ip.vipVariableName")
			data.SystemIpVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SystemIp = types.StringNull()
			data.SystemIpVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "system-ip.vipValue")
			data.SystemIp = types.StringValue(v.String())
			data.SystemIpVariable = types.StringNull()
		}
	} else {
		data.SystemIp = types.StringNull()
		data.SystemIpVariable = types.StringNull()
	}
	if value := res.Get(path + "overlay-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.OverlayId = types.Int64Null()

			v := res.Get(path + "overlay-id.vipVariableName")
			data.OverlayIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.OverlayId = types.Int64Null()
			data.OverlayIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "overlay-id.vipValue")
			data.OverlayId = types.Int64Value(v.Int())
			data.OverlayIdVariable = types.StringNull()
		}
	} else {
		data.OverlayId = types.Int64Null()
		data.OverlayIdVariable = types.StringNull()
	}
	if value := res.Get(path + "site-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SiteId = types.Int64Null()

			v := res.Get(path + "site-id.vipVariableName")
			data.SiteIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SiteId = types.Int64Null()
			data.SiteIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "site-id.vipValue")
			data.SiteId = types.Int64Value(v.Int())
			data.SiteIdVariable = types.StringNull()
		}
	} else {
		data.SiteId = types.Int64Null()
		data.SiteIdVariable = types.StringNull()
	}
	if value := res.Get(path + "port-offset.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PortOffset = types.Int64Null()

			v := res.Get(path + "port-offset.vipVariableName")
			data.PortOffsetVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PortOffset = types.Int64Null()
			data.PortOffsetVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "port-offset.vipValue")
			data.PortOffset = types.Int64Value(v.Int())
			data.PortOffsetVariable = types.StringNull()
		}
	} else {
		data.PortOffset = types.Int64Null()
		data.PortOffsetVariable = types.StringNull()
	}
	if value := res.Get(path + "port-hop.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.PortHopping = types.BoolNull()

			v := res.Get(path + "port-hop.vipVariableName")
			data.PortHoppingVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.PortHopping = types.BoolNull()
			data.PortHoppingVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "port-hop.vipValue")
			data.PortHopping = types.BoolValue(v.Bool())
			data.PortHoppingVariable = types.StringNull()
		}
	} else {
		data.PortHopping = types.BoolNull()
		data.PortHoppingVariable = types.StringNull()
	}
	if value := res.Get(path + "control-session-pps.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ControlSessionPps = types.Int64Null()

			v := res.Get(path + "control-session-pps.vipVariableName")
			data.ControlSessionPpsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ControlSessionPps = types.Int64Null()
			data.ControlSessionPpsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "control-session-pps.vipValue")
			data.ControlSessionPps = types.Int64Value(v.Int())
			data.ControlSessionPpsVariable = types.StringNull()
		}
	} else {
		data.ControlSessionPps = types.Int64Null()
		data.ControlSessionPpsVariable = types.StringNull()
	}
	if value := res.Get(path + "track-transport.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TrackTransport = types.BoolNull()

			v := res.Get(path + "track-transport.vipVariableName")
			data.TrackTransportVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TrackTransport = types.BoolNull()
			data.TrackTransportVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "track-transport.vipValue")
			data.TrackTransport = types.BoolValue(v.Bool())
			data.TrackTransportVariable = types.StringNull()
		}
	} else {
		data.TrackTransport = types.BoolNull()
		data.TrackTransportVariable = types.StringNull()
	}
	if value := res.Get(path + "track-interface-tag.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TrackInterfaceTag = types.Int64Null()

			v := res.Get(path + "track-interface-tag.vipVariableName")
			data.TrackInterfaceTagVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TrackInterfaceTag = types.Int64Null()
			data.TrackInterfaceTagVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "track-interface-tag.vipValue")
			data.TrackInterfaceTag = types.Int64Value(v.Int())
			data.TrackInterfaceTagVariable = types.StringNull()
		}
	} else {
		data.TrackInterfaceTag = types.Int64Null()
		data.TrackInterfaceTagVariable = types.StringNull()
	}
	if value := res.Get(path + "console-baud-rate.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.ConsoleBaudRate = types.StringNull()

			v := res.Get(path + "console-baud-rate.vipVariableName")
			data.ConsoleBaudRateVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.ConsoleBaudRate = types.StringNull()
			data.ConsoleBaudRateVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "console-baud-rate.vipValue")
			data.ConsoleBaudRate = types.StringValue(v.String())
			data.ConsoleBaudRateVariable = types.StringNull()
		}
	} else {
		data.ConsoleBaudRate = types.StringNull()
		data.ConsoleBaudRateVariable = types.StringNull()
	}
	if value := res.Get(path + "max-omp-sessions.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.MaxOmpSessions = types.Int64Null()

			v := res.Get(path + "max-omp-sessions.vipVariableName")
			data.MaxOmpSessionsVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.MaxOmpSessions = types.Int64Null()
			data.MaxOmpSessionsVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "max-omp-sessions.vipValue")
			data.MaxOmpSessions = types.Int64Value(v.Int())
			data.MaxOmpSessionsVariable = types.StringNull()
		}
	} else {
		data.MaxOmpSessions = types.Int64Null()
		data.MaxOmpSessionsVariable = types.StringNull()
	}
	if value := res.Get(path + "multi-tenant.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.MultiTenant = types.BoolNull()

			v := res.Get(path + "multi-tenant.vipVariableName")
			data.MultiTenantVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.MultiTenant = types.BoolNull()
			data.MultiTenantVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "multi-tenant.vipValue")
			data.MultiTenant = types.BoolValue(v.Bool())
			data.MultiTenantVariable = types.StringNull()
		}
	} else {
		data.MultiTenant = types.BoolNull()
		data.MultiTenantVariable = types.StringNull()
	}
	if value := res.Get(path + "track-default-gateway.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TrackDefaultGateway = types.BoolNull()

			v := res.Get(path + "track-default-gateway.vipVariableName")
			data.TrackDefaultGatewayVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TrackDefaultGateway = types.BoolNull()
			data.TrackDefaultGatewayVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "track-default-gateway.vipValue")
			data.TrackDefaultGateway = types.BoolValue(v.Bool())
			data.TrackDefaultGatewayVariable = types.StringNull()
		}
	} else {
		data.TrackDefaultGateway = types.BoolNull()
		data.TrackDefaultGatewayVariable = types.StringNull()
	}
	if value := res.Get(path + "admin-tech-on-failure.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AdminTechOnFailure = types.BoolNull()

			v := res.Get(path + "admin-tech-on-failure.vipVariableName")
			data.AdminTechOnFailureVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AdminTechOnFailure = types.BoolNull()
			data.AdminTechOnFailureVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "admin-tech-on-failure.vipValue")
			data.AdminTechOnFailure = types.BoolValue(v.Bool())
			data.AdminTechOnFailureVariable = types.StringNull()
		}
	} else {
		data.AdminTechOnFailure = types.BoolNull()
		data.AdminTechOnFailureVariable = types.StringNull()
	}
	if value := res.Get(path + "idle-timeout.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.IdleTimeout = types.Int64Null()

			v := res.Get(path + "idle-timeout.vipVariableName")
			data.IdleTimeoutVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.IdleTimeout = types.Int64Null()
			data.IdleTimeoutVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "idle-timeout.vipValue")
			data.IdleTimeout = types.Int64Value(v.Int())
			data.IdleTimeoutVariable = types.StringNull()
		}
	} else {
		data.IdleTimeout = types.Int64Null()
		data.IdleTimeoutVariable = types.StringNull()
	}
	if value := res.Get(path + "tracker.vipValue"); len(value.Array()) > 0 {
		data.Trackers = make([]CiscoSystemTrackers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoSystemTrackers{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Name = types.StringNull()

					cv := v.Get("name.vipVariableName")
					item.NameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Name = types.StringNull()
					item.NameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("name.vipValue")
					item.Name = types.StringValue(cv.String())
					item.NameVariable = types.StringNull()
				}
			} else {
				item.Name = types.StringNull()
				item.NameVariable = types.StringNull()
			}
			if cValue := v.Get("endpoint-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.EndpointIp = types.StringNull()

					cv := v.Get("endpoint-ip.vipVariableName")
					item.EndpointIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.EndpointIp = types.StringNull()
					item.EndpointIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("endpoint-ip.vipValue")
					item.EndpointIp = types.StringValue(cv.String())
					item.EndpointIpVariable = types.StringNull()
				}
			} else {
				item.EndpointIp = types.StringNull()
				item.EndpointIpVariable = types.StringNull()
			}
			if cValue := v.Get("endpoint-dns-name.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.EndpointDnsName = types.StringNull()

					cv := v.Get("endpoint-dns-name.vipVariableName")
					item.EndpointDnsNameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.EndpointDnsName = types.StringNull()
					item.EndpointDnsNameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("endpoint-dns-name.vipValue")
					item.EndpointDnsName = types.StringValue(cv.String())
					item.EndpointDnsNameVariable = types.StringNull()
				}
			} else {
				item.EndpointDnsName = types.StringNull()
				item.EndpointDnsNameVariable = types.StringNull()
			}
			if cValue := v.Get("endpoint-api-url.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.EndpointApiUrl = types.StringNull()

					cv := v.Get("endpoint-api-url.vipVariableName")
					item.EndpointApiUrlVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.EndpointApiUrl = types.StringNull()
					item.EndpointApiUrlVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("endpoint-api-url.vipValue")
					item.EndpointApiUrl = types.StringValue(cv.String())
					item.EndpointApiUrlVariable = types.StringNull()
				}
			} else {
				item.EndpointApiUrl = types.StringNull()
				item.EndpointApiUrlVariable = types.StringNull()
			}
			if cValue := v.Get("elements.vipType"); len(cValue.Array()) > 0 {
				if cValue.String() == "variableName" {
					item.Elements = types.SetNull(types.StringType)

					cv := v.Get("elements.vipVariableName")
					item.ElementsVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Elements = types.SetNull(types.StringType)
					item.ElementsVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("elements.vipValue")
					item.Elements = helpers.GetStringSet(cv.Array())
					item.ElementsVariable = types.StringNull()
				}
			} else {
				item.Elements = types.SetNull(types.StringType)
				item.ElementsVariable = types.StringNull()
			}
			if cValue := v.Get("boolean.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Boolean = types.StringNull()

					cv := v.Get("boolean.vipVariableName")
					item.BooleanVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Boolean = types.StringNull()
					item.BooleanVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("boolean.vipValue")
					item.Boolean = types.StringValue(cv.String())
					item.BooleanVariable = types.StringNull()
				}
			} else {
				item.Boolean = types.StringNull()
				item.BooleanVariable = types.StringNull()
			}
			if cValue := v.Get("threshold.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Threshold = types.Int64Null()

					cv := v.Get("threshold.vipVariableName")
					item.ThresholdVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Threshold = types.Int64Null()
					item.ThresholdVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("threshold.vipValue")
					item.Threshold = types.Int64Value(cv.Int())
					item.ThresholdVariable = types.StringNull()
				}
			} else {
				item.Threshold = types.Int64Null()
				item.ThresholdVariable = types.StringNull()
			}
			if cValue := v.Get("interval.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Interval = types.Int64Null()

					cv := v.Get("interval.vipVariableName")
					item.IntervalVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Interval = types.Int64Null()
					item.IntervalVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("interval.vipValue")
					item.Interval = types.Int64Value(cv.Int())
					item.IntervalVariable = types.StringNull()
				}
			} else {
				item.Interval = types.Int64Null()
				item.IntervalVariable = types.StringNull()
			}
			if cValue := v.Get("multiplier.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Multiplier = types.Int64Null()

					cv := v.Get("multiplier.vipVariableName")
					item.MultiplierVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Multiplier = types.Int64Null()
					item.MultiplierVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("multiplier.vipValue")
					item.Multiplier = types.Int64Value(cv.Int())
					item.MultiplierVariable = types.StringNull()
				}
			} else {
				item.Multiplier = types.Int64Null()
				item.MultiplierVariable = types.StringNull()
			}
			if cValue := v.Get("type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Type = types.StringNull()

					cv := v.Get("type.vipVariableName")
					item.TypeVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Type = types.StringNull()
					item.TypeVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("type.vipValue")
					item.Type = types.StringValue(cv.String())
					item.TypeVariable = types.StringNull()
				}
			} else {
				item.Type = types.StringNull()
				item.TypeVariable = types.StringNull()
			}
			data.Trackers = append(data.Trackers, item)
			return true
		})
	} else {
		if len(data.Trackers) > 0 {
			data.Trackers = []CiscoSystemTrackers{}
		}
	}
	if value := res.Get(path + "object-track.vipValue"); len(value.Array()) > 0 {
		data.ObjectTrackers = make([]CiscoSystemObjectTrackers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoSystemObjectTrackers{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("object-number.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ObjectNumber = types.Int64Null()

					cv := v.Get("object-number.vipVariableName")
					item.ObjectNumberVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.ObjectNumber = types.Int64Null()
					item.ObjectNumberVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("object-number.vipValue")
					item.ObjectNumber = types.Int64Value(cv.Int())
					item.ObjectNumberVariable = types.StringNull()
				}
			} else {
				item.ObjectNumber = types.Int64Null()
				item.ObjectNumberVariable = types.StringNull()
			}
			if cValue := v.Get("interface.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Interface = types.StringNull()

					cv := v.Get("interface.vipVariableName")
					item.InterfaceVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Interface = types.StringNull()
					item.InterfaceVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("interface.vipValue")
					item.Interface = types.StringValue(cv.String())
					item.InterfaceVariable = types.StringNull()
				}
			} else {
				item.Interface = types.StringNull()
				item.InterfaceVariable = types.StringNull()
			}
			if cValue := v.Get("sig.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Sig = types.StringNull()

					cv := v.Get("sig.vipVariableName")
					item.SigVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Sig = types.StringNull()
					item.SigVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("sig.vipValue")
					item.Sig = types.StringValue(cv.String())
					item.SigVariable = types.StringNull()
				}
			} else {
				item.Sig = types.StringNull()
				item.SigVariable = types.StringNull()
			}
			if cValue := v.Get("ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Ip = types.StringNull()

					cv := v.Get("ip.vipVariableName")
					item.IpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Ip = types.StringNull()
					item.IpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("ip.vipValue")
					item.Ip = types.StringValue(cv.String())
					item.IpVariable = types.StringNull()
				}
			} else {
				item.Ip = types.StringNull()
				item.IpVariable = types.StringNull()
			}
			if cValue := v.Get("mask.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Mask = types.StringNull()

					cv := v.Get("mask.vipVariableName")
					item.MaskVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Mask = types.StringNull()
					item.MaskVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("mask.vipValue")
					item.Mask = types.StringValue(cv.String())
					item.MaskVariable = types.StringNull()
				}
			} else {
				item.Mask = types.StringNull()
				item.MaskVariable = types.StringNull()
			}
			if cValue := v.Get("vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.VpnId = types.Int64Null()

				} else if cValue.String() == "ignore" {
					item.VpnId = types.Int64Null()

				} else if cValue.String() == "constant" {
					cv := v.Get("vpn.vipValue")
					item.VpnId = types.Int64Value(cv.Int())

				}
			} else {
				item.VpnId = types.Int64Null()

			}
			if cValue := v.Get("object.vipValue"); len(cValue.Array()) > 0 {
				item.GroupTracksIds = make([]CiscoSystemObjectTrackersGroupTracksIds, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := CiscoSystemObjectTrackersGroupTracksIds{}
					if ccValue := cv.Get("vipOptional"); ccValue.Exists() {
						cItem.Optional = types.BoolValue(ccValue.Bool())
					} else {
						cItem.Optional = types.BoolNull()
					}
					if ccValue := cv.Get("number.vipType"); ccValue.Exists() {
						if ccValue.String() == "variableName" {
							cItem.TrackId = types.Int64Null()

							ccv := cv.Get("number.vipVariableName")
							cItem.TrackIdVariable = types.StringValue(ccv.String())

						} else if ccValue.String() == "ignore" {
							cItem.TrackId = types.Int64Null()
							cItem.TrackIdVariable = types.StringNull()
						} else if ccValue.String() == "constant" {
							ccv := cv.Get("number.vipValue")
							cItem.TrackId = types.Int64Value(ccv.Int())
							cItem.TrackIdVariable = types.StringNull()
						}
					} else {
						cItem.TrackId = types.Int64Null()
						cItem.TrackIdVariable = types.StringNull()
					}
					item.GroupTracksIds = append(item.GroupTracksIds, cItem)
					return true
				})
			} else {
				if len(item.GroupTracksIds) > 0 {
					item.GroupTracksIds = []CiscoSystemObjectTrackersGroupTracksIds{}
				}
			}
			if cValue := v.Get("boolean.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.Boolean = types.StringNull()

					cv := v.Get("boolean.vipVariableName")
					item.BooleanVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.Boolean = types.StringNull()
					item.BooleanVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("boolean.vipValue")
					item.Boolean = types.StringValue(cv.String())
					item.BooleanVariable = types.StringNull()
				}
			} else {
				item.Boolean = types.StringNull()
				item.BooleanVariable = types.StringNull()
			}
			data.ObjectTrackers = append(data.ObjectTrackers, item)
			return true
		})
	} else {
		if len(data.ObjectTrackers) > 0 {
			data.ObjectTrackers = []CiscoSystemObjectTrackers{}
		}
	}
	if value := res.Get(path + "on-demand.enable.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.OnDemandTunnel = types.BoolNull()

			v := res.Get(path + "on-demand.enable.vipVariableName")
			data.OnDemandTunnelVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.OnDemandTunnel = types.BoolNull()
			data.OnDemandTunnelVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "on-demand.enable.vipValue")
			data.OnDemandTunnel = types.BoolValue(v.Bool())
			data.OnDemandTunnelVariable = types.StringNull()
		}
	} else {
		data.OnDemandTunnel = types.BoolNull()
		data.OnDemandTunnelVariable = types.StringNull()
	}
	if value := res.Get(path + "on-demand.idle-timeout.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.OnDemandTunnelIdleTimeout = types.Int64Null()

			v := res.Get(path + "on-demand.idle-timeout.vipVariableName")
			data.OnDemandTunnelIdleTimeoutVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.OnDemandTunnelIdleTimeout = types.Int64Null()
			data.OnDemandTunnelIdleTimeoutVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "on-demand.idle-timeout.vipValue")
			data.OnDemandTunnelIdleTimeout = types.Int64Value(v.Int())
			data.OnDemandTunnelIdleTimeoutVariable = types.StringNull()
		}
	} else {
		data.OnDemandTunnelIdleTimeout = types.Int64Null()
		data.OnDemandTunnelIdleTimeoutVariable = types.StringNull()
	}
	if value := res.Get(path + "region-id.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.RegionId = types.Int64Null()

			v := res.Get(path + "region-id.vipVariableName")
			data.RegionIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.RegionId = types.Int64Null()
			data.RegionIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "region-id.vipValue")
			data.RegionId = types.Int64Value(v.Int())
			data.RegionIdVariable = types.StringNull()
		}
	} else {
		data.RegionId = types.Int64Null()
		data.RegionIdVariable = types.StringNull()
	}
	if value := res.Get(path + "secondary-region.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.SecondaryRegionId = types.Int64Null()

			v := res.Get(path + "secondary-region.vipVariableName")
			data.SecondaryRegionIdVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.SecondaryRegionId = types.Int64Null()
			data.SecondaryRegionIdVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "secondary-region.vipValue")
			data.SecondaryRegionId = types.Int64Value(v.Int())
			data.SecondaryRegionIdVariable = types.StringNull()
		}
	} else {
		data.SecondaryRegionId = types.Int64Null()
		data.SecondaryRegionIdVariable = types.StringNull()
	}
	if value := res.Get(path + "role.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.Role = types.StringNull()

			v := res.Get(path + "role.vipVariableName")
			data.RoleVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.Role = types.StringNull()
			data.RoleVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "role.vipValue")
			data.Role = types.StringValue(v.String())
			data.RoleVariable = types.StringNull()
		}
	} else {
		data.Role = types.StringNull()
		data.RoleVariable = types.StringNull()
	}
	if value := res.Get(path + "affinity-group.affinity-group-number.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.AffinityGroupNumber = types.Int64Null()

			v := res.Get(path + "affinity-group.affinity-group-number.vipVariableName")
			data.AffinityGroupNumberVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AffinityGroupNumber = types.Int64Null()
			data.AffinityGroupNumberVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "affinity-group.affinity-group-number.vipValue")
			data.AffinityGroupNumber = types.Int64Value(v.Int())
			data.AffinityGroupNumberVariable = types.StringNull()
		}
	} else {
		data.AffinityGroupNumber = types.Int64Null()
		data.AffinityGroupNumberVariable = types.StringNull()
	}
	if value := res.Get(path + "affinity-group.preference.vipType"); len(value.Array()) > 0 {
		if value.String() == "variableName" {
			data.AffinityGroupPreference = types.ListNull(types.Int64Type)

			v := res.Get(path + "affinity-group.preference.vipVariableName")
			data.AffinityGroupPreferenceVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.AffinityGroupPreference = types.ListNull(types.Int64Type)
			data.AffinityGroupPreferenceVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "affinity-group.preference.vipValue")
			data.AffinityGroupPreference = helpers.GetInt64List(v.Array())
			data.AffinityGroupPreferenceVariable = types.StringNull()
		}
	} else {
		data.AffinityGroupPreference = types.ListNull(types.Int64Type)
		data.AffinityGroupPreferenceVariable = types.StringNull()
	}
	if value := res.Get(path + "transport-gateway.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.TransportGateway = types.BoolNull()

			v := res.Get(path + "transport-gateway.vipVariableName")
			data.TransportGatewayVariable = types.StringValue(v.String())

		} else if value.String() == "ignore" {
			data.TransportGateway = types.BoolNull()
			data.TransportGatewayVariable = types.StringNull()
		} else if value.String() == "constant" {
			v := res.Get(path + "transport-gateway.vipValue")
			data.TransportGateway = types.BoolValue(v.Bool())
			data.TransportGatewayVariable = types.StringNull()
		}
	} else {
		data.TransportGateway = types.BoolNull()
		data.TransportGatewayVariable = types.StringNull()
	}
	if value := res.Get(path + "epfr.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EnhancedAppAwareRouting = types.StringNull()

		} else if value.String() == "ignore" {
			data.EnhancedAppAwareRouting = types.StringNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "epfr.vipValue")
			data.EnhancedAppAwareRouting = types.StringValue(v.String())

		}
	} else {
		data.EnhancedAppAwareRouting = types.StringNull()

	}
	if value := res.Get(path + "enable-mrf-migration.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.EnableMrfMigration = types.StringNull()

		} else if value.String() == "ignore" {
			data.EnableMrfMigration = types.StringNull()

		} else if value.String() == "constant" {
			v := res.Get(path + "enable-mrf-migration.vipValue")
			data.EnableMrfMigration = types.StringValue(v.String())

		}
	} else {
		data.EnableMrfMigration = types.StringNull()

	}
	if value := res.Get(path + "migration-bgp-community.vipType"); value.Exists() {
		if value.String() == "variableName" {
			data.MigrationBgpCommunity = types.Int64Null()

		} else if value.String() == "ignore" {
			data.MigrationBgpCommunity = types.Int64Null()

		} else if value.String() == "constant" {
			v := res.Get(path + "migration-bgp-community.vipValue")
			data.MigrationBgpCommunity = types.Int64Value(v.Int())

		}
	} else {
		data.MigrationBgpCommunity = types.Int64Null()

	}
}

// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *CiscoSystem) hasChanges(ctx context.Context, state *CiscoSystem) bool {
	hasChanges := false
	if !data.Timezone.Equal(state.Timezone) {
		hasChanges = true
	}
	if !data.Hostname.Equal(state.Hostname) {
		hasChanges = true
	}
	if !data.SystemDescription.Equal(state.SystemDescription) {
		hasChanges = true
	}
	if !data.Location.Equal(state.Location) {
		hasChanges = true
	}
	if !data.Latitude.Equal(state.Latitude) {
		hasChanges = true
	}
	if !data.Longitude.Equal(state.Longitude) {
		hasChanges = true
	}
	if !data.GeoFencing.Equal(state.GeoFencing) {
		hasChanges = true
	}
	if !data.GeoFencingRange.Equal(state.GeoFencingRange) {
		hasChanges = true
	}
	if !data.GeoFencingSms.Equal(state.GeoFencingSms) {
		hasChanges = true
	}
	if len(data.GeoFencingSmsPhoneNumbers) != len(state.GeoFencingSmsPhoneNumbers) {
		hasChanges = true
	} else {
		for i := range data.GeoFencingSmsPhoneNumbers {
			if !data.GeoFencingSmsPhoneNumbers[i].Number.Equal(state.GeoFencingSmsPhoneNumbers[i].Number) {
				hasChanges = true
			}
		}
	}
	if !data.DeviceGroups.Equal(state.DeviceGroups) {
		hasChanges = true
	}
	if !data.ControllerGroupList.Equal(state.ControllerGroupList) {
		hasChanges = true
	}
	if !data.SystemIp.Equal(state.SystemIp) {
		hasChanges = true
	}
	if !data.OverlayId.Equal(state.OverlayId) {
		hasChanges = true
	}
	if !data.SiteId.Equal(state.SiteId) {
		hasChanges = true
	}
	if !data.PortOffset.Equal(state.PortOffset) {
		hasChanges = true
	}
	if !data.PortHopping.Equal(state.PortHopping) {
		hasChanges = true
	}
	if !data.ControlSessionPps.Equal(state.ControlSessionPps) {
		hasChanges = true
	}
	if !data.TrackTransport.Equal(state.TrackTransport) {
		hasChanges = true
	}
	if !data.TrackInterfaceTag.Equal(state.TrackInterfaceTag) {
		hasChanges = true
	}
	if !data.ConsoleBaudRate.Equal(state.ConsoleBaudRate) {
		hasChanges = true
	}
	if !data.MaxOmpSessions.Equal(state.MaxOmpSessions) {
		hasChanges = true
	}
	if !data.MultiTenant.Equal(state.MultiTenant) {
		hasChanges = true
	}
	if !data.TrackDefaultGateway.Equal(state.TrackDefaultGateway) {
		hasChanges = true
	}
	if !data.AdminTechOnFailure.Equal(state.AdminTechOnFailure) {
		hasChanges = true
	}
	if !data.IdleTimeout.Equal(state.IdleTimeout) {
		hasChanges = true
	}
	if len(data.Trackers) != len(state.Trackers) {
		hasChanges = true
	} else {
		for i := range data.Trackers {
			if !data.Trackers[i].Name.Equal(state.Trackers[i].Name) {
				hasChanges = true
			}
			if !data.Trackers[i].EndpointIp.Equal(state.Trackers[i].EndpointIp) {
				hasChanges = true
			}
			if !data.Trackers[i].EndpointDnsName.Equal(state.Trackers[i].EndpointDnsName) {
				hasChanges = true
			}
			if !data.Trackers[i].EndpointApiUrl.Equal(state.Trackers[i].EndpointApiUrl) {
				hasChanges = true
			}
			if !data.Trackers[i].Elements.Equal(state.Trackers[i].Elements) {
				hasChanges = true
			}
			if !data.Trackers[i].Boolean.Equal(state.Trackers[i].Boolean) {
				hasChanges = true
			}
			if !data.Trackers[i].Threshold.Equal(state.Trackers[i].Threshold) {
				hasChanges = true
			}
			if !data.Trackers[i].Interval.Equal(state.Trackers[i].Interval) {
				hasChanges = true
			}
			if !data.Trackers[i].Multiplier.Equal(state.Trackers[i].Multiplier) {
				hasChanges = true
			}
			if !data.Trackers[i].Type.Equal(state.Trackers[i].Type) {
				hasChanges = true
			}
		}
	}
	if len(data.ObjectTrackers) != len(state.ObjectTrackers) {
		hasChanges = true
	} else {
		for i := range data.ObjectTrackers {
			if !data.ObjectTrackers[i].ObjectNumber.Equal(state.ObjectTrackers[i].ObjectNumber) {
				hasChanges = true
			}
			if !data.ObjectTrackers[i].Interface.Equal(state.ObjectTrackers[i].Interface) {
				hasChanges = true
			}
			if !data.ObjectTrackers[i].Sig.Equal(state.ObjectTrackers[i].Sig) {
				hasChanges = true
			}
			if !data.ObjectTrackers[i].Ip.Equal(state.ObjectTrackers[i].Ip) {
				hasChanges = true
			}
			if !data.ObjectTrackers[i].Mask.Equal(state.ObjectTrackers[i].Mask) {
				hasChanges = true
			}
			if !data.ObjectTrackers[i].VpnId.Equal(state.ObjectTrackers[i].VpnId) {
				hasChanges = true
			}
			if len(data.ObjectTrackers[i].GroupTracksIds) != len(state.ObjectTrackers[i].GroupTracksIds) {
				hasChanges = true
			} else {
				for ii := range data.ObjectTrackers[i].GroupTracksIds {
					if !data.ObjectTrackers[i].GroupTracksIds[ii].TrackId.Equal(state.ObjectTrackers[i].GroupTracksIds[ii].TrackId) {
						hasChanges = true
					}
				}
			}
			if !data.ObjectTrackers[i].Boolean.Equal(state.ObjectTrackers[i].Boolean) {
				hasChanges = true
			}
		}
	}
	if !data.OnDemandTunnel.Equal(state.OnDemandTunnel) {
		hasChanges = true
	}
	if !data.OnDemandTunnelIdleTimeout.Equal(state.OnDemandTunnelIdleTimeout) {
		hasChanges = true
	}
	if !data.RegionId.Equal(state.RegionId) {
		hasChanges = true
	}
	if !data.SecondaryRegionId.Equal(state.SecondaryRegionId) {
		hasChanges = true
	}
	if !data.Role.Equal(state.Role) {
		hasChanges = true
	}
	if !data.AffinityGroupNumber.Equal(state.AffinityGroupNumber) {
		hasChanges = true
	}
	if !data.AffinityGroupPreference.Equal(state.AffinityGroupPreference) {
		hasChanges = true
	}
	if !data.TransportGateway.Equal(state.TransportGateway) {
		hasChanges = true
	}
	if !data.EnhancedAppAwareRouting.Equal(state.EnhancedAppAwareRouting) {
		hasChanges = true
	}
	if !data.EnableMrfMigration.Equal(state.EnableMrfMigration) {
		hasChanges = true
	}
	if !data.MigrationBgpCommunity.Equal(state.MigrationBgpCommunity) {
		hasChanges = true
	}
	return hasChanges
}

// End of section. //template:end hasChanges

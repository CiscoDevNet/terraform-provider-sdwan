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
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &TransportRoutingOSPFProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &TransportRoutingOSPFProfileParcelDataSource{}
)

func NewTransportRoutingOSPFProfileParcelDataSource() datasource.DataSource {
	return &TransportRoutingOSPFProfileParcelDataSource{}
}

type TransportRoutingOSPFProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *TransportRoutingOSPFProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transport_routing_ospf_feature"
}

func (d *TransportRoutingOSPFProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Transport Routing OSPF Feature.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Feature",
				Computed:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: "Feature Profile ID",
				Required:            true,
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: "Set OSPF router ID to override system IP address",
				Computed:            true,
			},
			"router_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"reference_bandwidth": schema.Int64Attribute{
				MarkdownDescription: "Set reference bandwidth method to assign OSPF cost",
				Computed:            true,
			},
			"reference_bandwidth_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"rfc_1583_compatible": schema.BoolAttribute{
				MarkdownDescription: "Calculate summary route cost based on RFC 1583",
				Computed:            true,
			},
			"rfc_1583_compatible_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"default_information_originate": schema.BoolAttribute{
				MarkdownDescription: "Distribute default external route into OSPF",
				Computed:            true,
			},
			"default_information_originate_always": schema.BoolAttribute{
				MarkdownDescription: "Always advertise default route",
				Computed:            true,
			},
			"default_information_originate_always_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"default_information_originate_metric": schema.Int64Attribute{
				MarkdownDescription: "Set metric used to generate default route <0..16777214>",
				Computed:            true,
			},
			"default_information_originate_metric_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"default_information_originate_metric_type": schema.StringAttribute{
				MarkdownDescription: "Set default route type",
				Computed:            true,
			},
			"default_information_originate_metric_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"distance_external": schema.Int64Attribute{
				MarkdownDescription: "Set distance for external routes",
				Computed:            true,
			},
			"distance_external_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"distance_inter_area": schema.Int64Attribute{
				MarkdownDescription: "Set distance for inter-area routes",
				Computed:            true,
			},
			"distance_inter_area_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"distance_intra_area": schema.Int64Attribute{
				MarkdownDescription: "Set distance for intra-area routes",
				Computed:            true,
			},
			"distance_intra_area_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"spf_calculation_delay": schema.Int64Attribute{
				MarkdownDescription: "Set delay from first change received until performing SPF calculation",
				Computed:            true,
			},
			"spf_calculation_delay_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"spf_initial_hold_time": schema.Int64Attribute{
				MarkdownDescription: "Set initial hold time between consecutive SPF calculations",
				Computed:            true,
			},
			"spf_initial_hold_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"spf_maximum_hold_time": schema.Int64Attribute{
				MarkdownDescription: "Set maximum hold time between consecutive SPF calculations",
				Computed:            true,
			},
			"spf_maximum_hold_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"redistributes": schema.ListNestedAttribute{
				MarkdownDescription: "Redistribute routes",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Set the protocol",
							Computed:            true,
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"nat_dia": schema.BoolAttribute{
							MarkdownDescription: "Enable NAT DIA for redistributed routes",
							Computed:            true,
						},
						"nat_dia_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"route_policy_id": schema.StringAttribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
			"router_lsas": schema.ListNestedAttribute{
				MarkdownDescription: "Advertise own router LSA with infinite distance",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							MarkdownDescription: "Set the router LSA advertisement type",
							Computed:            true,
						},
						"time": schema.Int64Attribute{
							MarkdownDescription: "Set how long to advertise maximum metric after router starts up",
							Computed:            true,
						},
						"time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
					},
				},
			},
			"route_policy_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"areas": schema.ListNestedAttribute{
				MarkdownDescription: "Configure OSPF area",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"area_number": schema.Int64Attribute{
							MarkdownDescription: "Set OSPF area number",
							Computed:            true,
						},
						"area_number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"area_type": schema.StringAttribute{
							MarkdownDescription: "set the area type",
							Computed:            true,
						},
						"no_summary": schema.BoolAttribute{
							MarkdownDescription: "Do not inject interarea routes into STUB or NSSA",
							Computed:            true,
						},
						"no_summary_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Computed:            true,
						},
						"interfaces": schema.ListNestedAttribute{
							MarkdownDescription: "Set OSPF interface parameters",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: "Set interface name",
										Computed:            true,
									},
									"name_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"hello_interval": schema.Int64Attribute{
										MarkdownDescription: "Set interval between OSPF hello packets",
										Computed:            true,
									},
									"hello_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"dead_interval": schema.Int64Attribute{
										MarkdownDescription: "Set interval after which neighbor is declared to be down",
										Computed:            true,
									},
									"dead_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"lsa_retransmit_interval": schema.Int64Attribute{
										MarkdownDescription: "Set time between retransmitting LSAs",
										Computed:            true,
									},
									"lsa_retransmit_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"cost": schema.Int64Attribute{
										MarkdownDescription: "Set cost of OSPF interface",
										Computed:            true,
									},
									"cost_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"designated_router_priority": schema.Int64Attribute{
										MarkdownDescription: "Set router’s priority to be elected as designated router",
										Computed:            true,
									},
									"designated_router_priority_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"network_type": schema.StringAttribute{
										MarkdownDescription: "Set the OSPF network type",
										Computed:            true,
									},
									"network_type_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"passive_interface": schema.BoolAttribute{
										MarkdownDescription: "Set the interface to advertise its address, but not to actively run OSPF",
										Computed:            true,
									},
									"passive_interface_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"authentication_type": schema.StringAttribute{
										MarkdownDescription: "Set OSPF interface authentication type",
										Computed:            true,
									},
									"authentication_type_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"message_digest_key_id": schema.Int64Attribute{
										MarkdownDescription: "Set MD5 message digest key",
										Computed:            true,
									},
									"message_digest_key_id_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"message_digest_key": schema.StringAttribute{
										MarkdownDescription: "Set MD5 authentication key",
										Computed:            true,
									},
									"message_digest_key_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
								},
							},
						},
						"ranges": schema.ListNestedAttribute{
							MarkdownDescription: "Summarize OSPF routes at an area boundary",
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										MarkdownDescription: "IP Address",
										Computed:            true,
									},
									"ip_address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"subnet_mask": schema.StringAttribute{
										MarkdownDescription: "Subnet Mask",
										Computed:            true,
									},
									"subnet_mask_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"cost": schema.Int64Attribute{
										MarkdownDescription: "Set cost for this range",
										Computed:            true,
									},
									"cost_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
									"no_advertise": schema.BoolAttribute{
										MarkdownDescription: "Do not advertise this range",
										Computed:            true,
									},
									"no_advertise_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *TransportRoutingOSPFProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *TransportRoutingOSPFProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config TransportRoutingOSPF

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read
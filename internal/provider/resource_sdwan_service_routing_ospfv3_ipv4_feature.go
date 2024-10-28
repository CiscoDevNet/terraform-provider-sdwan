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
	"regexp"
	"strings"
	"sync"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &ServiceRoutingOSPFv3IPv4ProfileParcelResource{}
var _ resource.ResourceWithImportState = &ServiceRoutingOSPFv3IPv4ProfileParcelResource{}

func NewServiceRoutingOSPFv3IPv4ProfileParcelResource() resource.Resource {
	return &ServiceRoutingOSPFv3IPv4ProfileParcelResource{}
}

type ServiceRoutingOSPFv3IPv4ProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *ServiceRoutingOSPFv3IPv4ProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_routing_ospfv3_ipv4_feature"
}

func (r *ServiceRoutingOSPFv3IPv4ProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Service Routing OSPFv3 IPv4 Feature.").AddMinimumVersionDescription("20.12.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Feature",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Feature",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Feature",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Feature",
				Optional:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Feature Profile ID").String,
				Required:            true,
			},
			"router_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set OSPF router ID to override system IP address").String,
				Optional:            true,
			},
			"router_id_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"distance": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Distance").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("110").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"distance_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"distance_external": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set distance for external routes").AddIntegerRangeDescription(1, 254).AddDefaultValueDescription("110").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 254),
				},
			},
			"distance_external_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"distance_inter_area": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set distance for inter-area routes").AddIntegerRangeDescription(1, 254).AddDefaultValueDescription("110").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 254),
				},
			},
			"distance_inter_area_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"distance_intra_area": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set distance for intra-area routes").AddIntegerRangeDescription(1, 254).AddDefaultValueDescription("110").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 254),
				},
			},
			"distance_intra_area_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"reference_bandwidth": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set reference bandwidth method to assign OSPF cost").AddIntegerRangeDescription(1, 4294967).AddDefaultValueDescription("100").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967),
				},
			},
			"reference_bandwidth_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"rfc_1583_compatible": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Calculate summary route cost based on RFC 1583").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"rfc_1583_compatible_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"default_information_originate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Distribute default external route into OSPF disabled").String,
				Optional:            true,
			},
			"default_information_originate_always": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Always advertise default route").String,
				Optional:            true,
			},
			"default_information_originate_always_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"default_information_originate_metric": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set metric used to generate default route <0..16777214>").AddIntegerRangeDescription(0, 16777214).String,
				Optional:            true,
			},
			"default_information_originate_metric_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"default_information_originate_metric_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set default route metric type").AddStringEnumDescription("type1", "type2").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("type1", "type2"),
				},
			},
			"default_information_originate_metric_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"spf_calculation_delay": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set delay from first change received until performing SPF calculation").AddIntegerRangeDescription(1, 600000).AddDefaultValueDescription("200").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 600000),
				},
			},
			"spf_calculation_delay_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"spf_initial_hold_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set initial hold time between consecutive SPF calculations").AddIntegerRangeDescription(1, 600000).AddDefaultValueDescription("1000").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 600000),
				},
			},
			"spf_initial_hold_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"spf_maximum_hold_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set maximum hold time between consecutive SPF calculations").AddIntegerRangeDescription(1, 600000).AddDefaultValueDescription("10000").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 600000),
				},
			},
			"spf_maximum_hold_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"route_policy_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
				},
			},
			"filter": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Table map filtered or not").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"filter_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"redistributes": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Redistribute routes").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the protocol").AddStringEnumDescription("connected", "static", "omp", "nat-route", "bgp", "eigrp").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("connected", "static", "omp", "nat-route", "bgp", "eigrp"),
							},
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"nat_dia": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Enable NAT DIA for redistributed routes").AddDefaultValueDescription("true").String,
							Optional:            true,
						},
						"nat_dia_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"route_policy_id": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`), ""),
							},
						},
					},
				},
			},
			"router_lsa_action": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Not advertise maximum metric Router LSA policy by default").String,
				Optional:            true,
			},
			"router_lsa_on_startup_time": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set how long to advertise maximum metric after router boot up").AddIntegerRangeDescription(5, 86400).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(5, 86400),
				},
			},
			"router_lsa_on_startup_time_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"areas": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure OSPFv3 IPv4 area").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"area_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set OSPF area number").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.AtMost(4294967295),
							},
						},
						"area_number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"area_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("stub area type").AddStringEnumDescription("stub").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("stub"),
							},
						},
						"no_summary": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Do not inject inter-area routes").String,
							Optional:            true,
						},
						"no_summary_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"always_translate": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Always translate type7 LSAs").String,
							Optional:            true,
						},
						"always_translate_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"interfaces": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set OSPF interface parameters").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set interface name").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.LengthBetween(3, 32),
											stringvalidator.RegexMatches(regexp.MustCompile(`(AppGigabitEthernet|BDI|BD-VIF|Virtual-Template|Dialer|Ethernet|FastEthernet|FiftyGigabitEthernet|FiveGigabitEthernet|FortyGigabitEthernet|FourHundredGigE|GigabitEthernet|HundredGigE|Loopback|Port-channel|TenGigabitEthernet|Tunnel|TwentyFiveGigE|TwoGigabitEthernet|TwoHundredGigE|Vlan|vmi)([0-9]*(. ?[1-9][0-9]*)*|[0-9/]+|[0-9]+/[0-9]+/[0-9]+:[0-9]+|[0-9]+/[0-9]+/[0-9]+|[0-9]+/[0-9]+|[0-9]+)`), ""),
										},
									},
									"name_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"hello_interval": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set interval between OSPF hello packets").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("10").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
									},
									"hello_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"dead_interval": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set interval after which neighbor is declared to be down").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("40").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
									},
									"dead_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"lsa_retransmit_interval": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set time between retransmitting LSAs").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("5").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
									},
									"lsa_retransmit_interval_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"cost": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set cost of OSPF interface").AddIntegerRangeDescription(1, 65535).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
									},
									"cost_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"network_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set the OSPF network type").AddStringEnumDescription("broadcast", "point-to-point", "non-broadcast", "point-to-multipoint").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("broadcast", "point-to-point", "non-broadcast", "point-to-multipoint"),
										},
									},
									"network_type_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"passive_interface": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set the interface to advertise its address, but not to actively run OSPF").AddDefaultValueDescription("false").String,
										Optional:            true,
									},
									"passive_interface_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"authentication_type": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("No Authentication by default").AddStringEnumDescription("no-auth").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("no-auth"),
										},
									},
									"authentication_spi": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set OSPF interface authentication IPSec SPI, range 256..4294967295").AddIntegerRangeDescription(256, 4294967295).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(256, 4294967295),
										},
									},
									"authentication_spi_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"authentication_key": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set OSPF interface authentication IPSEC key").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9a-fA-F]{40}$`), ""),
										},
									},
									"authentication_key_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
								},
							},
						},
						"ranges": schema.ListNestedAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Summarize OSPF routes at an area boundary").String,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip_address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").String,
										Optional:            true,
									},
									"ip_address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"subnet_mask": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("255.255.255.255", "255.255.255.254", "255.255.255.252", "255.255.255.248", "255.255.255.240", "255.255.255.224", "255.255.255.192", "255.255.255.128", "255.255.255.0", "255.255.254.0", "255.255.252.0", "255.255.248.0", "255.255.240.0", "255.255.224.0", "255.255.192.0", "255.255.128.0", "255.255.0.0", "255.254.0.0", "255.252.0.0", "255.240.0.0", "255.224.0.0", "255.192.0.0", "255.128.0.0", "255.0.0.0", "254.0.0.0", "252.0.0.0", "248.0.0.0", "240.0.0.0", "224.0.0.0", "192.0.0.0", "128.0.0.0", "0.0.0.0"),
										},
									},
									"subnet_mask_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"cost": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set cost for this range").AddIntegerRangeDescription(0, 16777214).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.AtMost(16777214),
										},
									},
									"cost_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"no_advertise": schema.BoolAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Do not advertise this range").AddDefaultValueDescription("false").String,
										Optional:            true,
									},
									"no_advertise_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
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

func (r *ServiceRoutingOSPFv3IPv4ProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *ServiceRoutingOSPFv3IPv4ProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ServiceRoutingOSPFv3IPv4

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("parcelId").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *ServiceRoutingOSPFv3IPv4ProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ServiceRoutingOSPFv3IPv4

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if res.Get("error.message").String() == "Invalid feature Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *ServiceRoutingOSPFv3IPv4ProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ServiceRoutingOSPFv3IPv4

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Name.ValueString()))

	body := plan.toBody(ctx)
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *ServiceRoutingOSPFv3IPv4ProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ServiceRoutingOSPFv3IPv4

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *ServiceRoutingOSPFv3IPv4ProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "service_routing_ospfv3_ipv4_feature_id" + ",feature_profile_id"
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q, %q", pattern, req.ID, count),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("feature_profile_id"), parts[1])...)
}

// End of section. //template:end import

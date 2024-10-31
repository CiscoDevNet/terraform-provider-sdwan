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
var _ resource.Resource = &CiscoOSPFFeatureTemplateResource{}
var _ resource.ResourceWithImportState = &CiscoOSPFFeatureTemplateResource{}

func NewCiscoOSPFFeatureTemplateResource() resource.Resource {
	return &CiscoOSPFFeatureTemplateResource{}
}

type CiscoOSPFFeatureTemplateResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *CiscoOSPFFeatureTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cisco_ospf_feature_template"
}

func (r *CiscoOSPFFeatureTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Cisco OSPF feature template.").AddMinimumVersionDescription("15.0.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the feature template",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the feature template",
				Computed:            true,
			},
			"template_type": schema.StringAttribute{
				MarkdownDescription: "The template type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the feature template",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the feature template",
				Required:            true,
			},
			"device_types": schema.SetAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("List of supported device types").AddStringEnumDescription("vedge-C8000V", "vedge-C8300-1N1S-4T2X", "vedge-C8300-1N1S-6T", "vedge-C8300-2N2S-6T", "vedge-C8300-2N2S-4T2X", "vedge-C8500-12X4QC", "vedge-C8500-12X", "vedge-C8500-20X6C", "vedge-C8500L-8S4X", "vedge-C8200-1N-4T", "vedge-C8200L-1N-4T").String,
				ElementType:         types.StringType,
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
			"auto_cost_reference_bandwidth": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set reference bandwidth method to assign OSPF cost").AddIntegerRangeDescription(1, 4294967).AddDefaultValueDescription("100").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 4294967),
				},
			},
			"auto_cost_reference_bandwidth_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"compatible_rfc1583": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Calculate summary route cost based on RFC 1583").AddDefaultValueDescription("true").String,
				Optional:            true,
			},
			"compatible_rfc1583_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"default_information_originate": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Distribute default external route into OSPF").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"default_information_originate_always": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Always advertise default route").AddDefaultValueDescription("false").String,
				Optional:            true,
			},
			"default_information_originate_always_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"default_information_originate_metric": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set metric used to generate default route <0..16777214>").AddIntegerRangeDescription(0, 16777214).String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 16777214),
				},
			},
			"default_information_originate_metric_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"default_information_originate_metric_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set default route type").AddStringEnumDescription("type1", "type2").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("type1", "type2"),
				},
			},
			"default_information_originate_metric_type_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"distance_external": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set distance for external routes").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("110").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"distance_external_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"distance_inter_area": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set distance for inter-area routes").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("110").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"distance_inter_area_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"distance_intra_area": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set distance for intra-area routes").AddIntegerRangeDescription(1, 255).AddDefaultValueDescription("110").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 255),
				},
			},
			"distance_intra_area_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"timers_spf_delay": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set delay from first change received until performing SPF calculation").AddIntegerRangeDescription(1, 600000).AddDefaultValueDescription("200").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 600000),
				},
			},
			"timers_spf_delay_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"timers_spf_initial_hold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set initial hold time between consecutive SPF calculations").AddIntegerRangeDescription(1, 600000).AddDefaultValueDescription("1000").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 600000),
				},
			},
			"timers_spf_initial_hold_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"timers_spf_max_hold": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set maximum hold time between consecutive SPF calculations").AddIntegerRangeDescription(1, 600000).AddDefaultValueDescription("10000").String,
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(1, 600000),
				},
			},
			"timers_spf_max_hold_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Optional:            true,
			},
			"redistribute": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Redistribute routes").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the protocol").AddStringEnumDescription("static", "connected", "bgp", "omp", "nat", "eigrp").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("static", "connected", "bgp", "omp", "nat", "eigrp"),
							},
						},
						"protocol_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"route_policy": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set route policy to apply to redistributed routes").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.LengthBetween(1, 32),
							},
						},
						"route_policy_variable": schema.StringAttribute{
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
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"max_metric_router_lsa": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Advertise own router LSA with infinite distance").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ad_type": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set the router LSA advertisement type").AddStringEnumDescription("administrative", "on-startup").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("administrative", "on-startup"),
							},
						},
						"time": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set how long to advertise maximum metric after router starts up").AddIntegerRangeDescription(5, 86400).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(5, 86400),
							},
						},
						"time_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"route_policies": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Set route policy to apply").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"direction": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set direction to apply policy").AddStringEnumDescription("in").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("in"),
							},
						},
						"direction_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"policy_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Name of route policy").String,
							Optional:            true,
						},
						"policy_name_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
			"areas": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Configure OSPF area").String,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"area_number": schema.Int64Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("Set OSPF area number").AddIntegerRangeDescription(0, 4294967295).String,
							Optional:            true,
							Validators: []validator.Int64{
								int64validator.Between(0, 4294967295),
							},
						},
						"area_number_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"stub": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Stub area").String,
							Optional:            true,
						},
						"stub_no_summary": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Do not inject interarea routes into stub").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"stub_no_summary_variable": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
							Optional:            true,
						},
						"nssa": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("NSSA area").String,
							Optional:            true,
						},
						"nssa_no_summary": schema.BoolAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("Do not inject interarea routes into NSSA").AddDefaultValueDescription("false").String,
							Optional:            true,
						},
						"nssa_no_summary_variable": schema.StringAttribute{
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
											stringvalidator.LengthBetween(1, 32),
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
									"retransmit_interval": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set time between retransmitting LSAs").AddIntegerRangeDescription(1, 65535).AddDefaultValueDescription("5").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 65535),
										},
									},
									"retransmit_interval_variable": schema.StringAttribute{
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
									"priority": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set router’s priority to be elected as designated router").AddIntegerRangeDescription(0, 255).AddDefaultValueDescription("1").String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 255),
										},
									},
									"priority_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"network": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set the OSPF network type").AddStringEnumDescription("broadcast", "point-to-point", "non-broadcast", "point-to-multipoint").AddDefaultValueDescription("broadcast").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("broadcast", "point-to-point", "non-broadcast", "point-to-multipoint"),
										},
									},
									"network_variable": schema.StringAttribute{
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
										MarkdownDescription: helpers.NewAttributeDescription("Set OSPF interface authentication type").AddStringEnumDescription("simple", "message-digest", "null").String,
										Optional:            true,
										Validators: []validator.String{
											stringvalidator.OneOf("simple", "message-digest", "null"),
										},
									},
									"authentication_type_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"authentication_message_digest_key_id": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set MD5 message digest key").AddIntegerRangeDescription(1, 255).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(1, 255),
										},
									},
									"authentication_message_digest_key_id_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"authentication_message_digest_key": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set MD5 authentication key").String,
										Optional:            true,
									},
									"authentication_message_digest_key_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
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
									"address": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set matching prefix").String,
										Optional:            true,
									},
									"address_variable": schema.StringAttribute{
										MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
										Optional:            true,
									},
									"cost": schema.Int64Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("Set cost for this range").AddIntegerRangeDescription(0, 16777214).String,
										Optional:            true,
										Validators: []validator.Int64{
											int64validator.Between(0, 16777214),
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
									"optional": schema.BoolAttribute{
										MarkdownDescription: "Indicates if list item is considered optional.",
										Optional:            true,
									},
								},
							},
						},
						"optional": schema.BoolAttribute{
							MarkdownDescription: "Indicates if list item is considered optional.",
							Optional:            true,
						},
					},
				},
			},
		},
	}
}

func (r *CiscoOSPFFeatureTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *CiscoOSPFFeatureTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CiscoOSPF

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post("/template/feature", body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("templateId").String())
	plan.Version = types.Int64Value(0)
	plan.TemplateType = types.StringValue(plan.getModel())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *CiscoOSPFFeatureTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CiscoOSPF

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get("/template/feature/object/" + url.QueryEscape(state.Id.ValueString()))
	if res.Get("error.message").String() == "Invalid Template Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *CiscoOSPFFeatureTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CiscoOSPF

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
	r.updateMutex.Lock()
	res, err := r.client.Put("/template/feature/"+url.QueryEscape(plan.Id.ValueString()), body)
	r.updateMutex.Unlock()
	if err != nil {
		if res.Get("error.message").String() == "Template locked in edit mode." {
			resp.Diagnostics.AddWarning("Client Warning", fmt.Sprintf("Failed to modify template due to template being locked by another change. Template changes will not be applied. Re-run 'terraform apply' to try again."))
		} else {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
			return
		}
	}

	if plan.hasChanges(ctx, &state) {
		plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)
	} else {
		plan.Version = state.Version
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *CiscoOSPFFeatureTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CiscoOSPF

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete("/template/feature/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *CiscoOSPFFeatureTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import

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
var _ resource.Resource = &PolicyObjectSecurityGeolocationListProfileParcelResource{}
var _ resource.ResourceWithImportState = &PolicyObjectSecurityGeolocationListProfileParcelResource{}

func NewPolicyObjectSecurityGeolocationListProfileParcelResource() resource.Resource {
	return &PolicyObjectSecurityGeolocationListProfileParcelResource{}
}

type PolicyObjectSecurityGeolocationListProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *PolicyObjectSecurityGeolocationListProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_object_security_geolocation_list_profile_parcel"
}

func (r *PolicyObjectSecurityGeolocationListProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Policy Object Security Geolocation List profile parcel.").AddMinimumVersionDescription("20.12.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the profile parcel",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the profile parcel",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the profile parcel",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the profile parcel",
				Optional:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Feature Profile ID").String,
				Required:            true,
			},
			"entries": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Geolocation  List").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"country": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("country names").AddStringEnumDescription("AFG", "ALB", "ATA", "DZA", "ASM", "AND", "AGO", "ATG", "AZE", "ARG", "AUS", "AUT", "BHS", "BHR", "BGD", "ARM", "BRB", "BEL", "BMU", "BTN", "BOL", "BIH", "BWA", "BVT", "BRA", "BLZ", "IOT", "SLB", "VGB", "BRN", "BGR", "MMR", "BDI", "BLR", "KHM", "CMR", "CAN", "CPV", "CYM", "CAF", "LKA", "TCD", "CHL", "CHN", "TWN", "CXR", "CCK", "COL", "COM", "MYT", "COG", "COD", "COK", "CRI", "HRV", "CUB", "CYP", "CZE", "BEN", "DNK", "DMA", "DOM", "ECU", "SLV", "GNQ", "ETH", "ERI", "EST", "FRO", "FLK", "SGS", "FJI", "FIN", "ALA", "FRA", "GUF", "PYF", "ATF", "DJI", "GAB", "GEO", "GMB", "PSE", "DEU", "GHA", "GIB", "KIR", "GRC", "GRL", "GRD", "GLP", "GUM", "GTM", "GIN", "GUY", "HTI", "HMD", "VAT", "HND", "HKG", "HUN", "ISL", "IND", "IDN", "IRN", "IRQ", "IRL", "ISR", "ITA", "CIV", "JAM", "JPN", "KAZ", "JOR", "KEN", "PRK", "KOR", "KWT", "KGZ", "LAO", "LBN", "LSO", "LVA", "LBR", "LBY", "LIE", "LTU", "LUX", "MAC", "MDG", "MWI", "MYS", "MDV", "MLI", "MLT", "MTQ", "MRT", "MUS", "MEX", "MCO", "MNG", "MDA", "MNE", "MSR", "MAR", "MOZ", "OMN", "NAM", "NRU", "NPL", "NLD", "ANT", "CUW", "ABW", "SXM", "BES", "NCL", "VUT", "NZL", "NIC", "NER", "NGA", "NIU", "NFK", "NOR", "MNP", "UMI", "FSM", "MHL", "PLW", "PAK", "PAN", "PNG", "PRY", "PER", "PHL", "PCN", "POL", "PRT", "GNB", "TLS", "PRI", "QAT", "REU", "ROU", "RUS", "RWA", "BLM", "SHN", "KNA", "AIA", "LCA", "MAF", "SPM", "VCT", "SMR", "STP", "SAU", "SEN", "SRB", "SYC", "SLE", "SGP", "SVK", "VNM", "SVN", "SOM", "ZAF", "ZWE", "ESP", "SSD", "ESH", "SDN", "SUR", "SJM", "SWZ", "SWE", "CHE", "SYR", "TJK", "THA", "TGO", "TKL", "TON", "TTO", "ARE", "TUN", "TUR", "TKM", "TCA", "TUV", "UGA", "UKR", "MKD", "EGY", "GBR", "GGY", "JEY", "IMN", "TZA", "USA", "VIR", "BFA", "URY", "UZB", "VEN", "WLF", "WSM", "YEM", "ZMB").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("AFG", "ALB", "ATA", "DZA", "ASM", "AND", "AGO", "ATG", "AZE", "ARG", "AUS", "AUT", "BHS", "BHR", "BGD", "ARM", "BRB", "BEL", "BMU", "BTN", "BOL", "BIH", "BWA", "BVT", "BRA", "BLZ", "IOT", "SLB", "VGB", "BRN", "BGR", "MMR", "BDI", "BLR", "KHM", "CMR", "CAN", "CPV", "CYM", "CAF", "LKA", "TCD", "CHL", "CHN", "TWN", "CXR", "CCK", "COL", "COM", "MYT", "COG", "COD", "COK", "CRI", "HRV", "CUB", "CYP", "CZE", "BEN", "DNK", "DMA", "DOM", "ECU", "SLV", "GNQ", "ETH", "ERI", "EST", "FRO", "FLK", "SGS", "FJI", "FIN", "ALA", "FRA", "GUF", "PYF", "ATF", "DJI", "GAB", "GEO", "GMB", "PSE", "DEU", "GHA", "GIB", "KIR", "GRC", "GRL", "GRD", "GLP", "GUM", "GTM", "GIN", "GUY", "HTI", "HMD", "VAT", "HND", "HKG", "HUN", "ISL", "IND", "IDN", "IRN", "IRQ", "IRL", "ISR", "ITA", "CIV", "JAM", "JPN", "KAZ", "JOR", "KEN", "PRK", "KOR", "KWT", "KGZ", "LAO", "LBN", "LSO", "LVA", "LBR", "LBY", "LIE", "LTU", "LUX", "MAC", "MDG", "MWI", "MYS", "MDV", "MLI", "MLT", "MTQ", "MRT", "MUS", "MEX", "MCO", "MNG", "MDA", "MNE", "MSR", "MAR", "MOZ", "OMN", "NAM", "NRU", "NPL", "NLD", "ANT", "CUW", "ABW", "SXM", "BES", "NCL", "VUT", "NZL", "NIC", "NER", "NGA", "NIU", "NFK", "NOR", "MNP", "UMI", "FSM", "MHL", "PLW", "PAK", "PAN", "PNG", "PRY", "PER", "PHL", "PCN", "POL", "PRT", "GNB", "TLS", "PRI", "QAT", "REU", "ROU", "RUS", "RWA", "BLM", "SHN", "KNA", "AIA", "LCA", "MAF", "SPM", "VCT", "SMR", "STP", "SAU", "SEN", "SRB", "SYC", "SLE", "SGP", "SVK", "VNM", "SVN", "SOM", "ZAF", "ZWE", "ESP", "SSD", "ESH", "SDN", "SUR", "SJM", "SWZ", "SWE", "CHE", "SYR", "TJK", "THA", "TGO", "TKL", "TON", "TTO", "ARE", "TUN", "TUR", "TKM", "TCA", "TUV", "UGA", "UKR", "MKD", "EGY", "GBR", "GGY", "JEY", "IMN", "TZA", "USA", "VIR", "BFA", "URY", "UZB", "VEN", "WLF", "WSM", "YEM", "ZMB"),
							},
						},
						"continent": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("continent name").AddStringEnumDescription("AF", "AN", "AS", "EU", "NA", "OC", "SA").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("AF", "AN", "AS", "EU", "NA", "OC", "SA"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *PolicyObjectSecurityGeolocationListProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *PolicyObjectSecurityGeolocationListProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PolicyObjectSecurityGeolocationList

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
func (r *PolicyObjectSecurityGeolocationListProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PolicyObjectSecurityGeolocationList

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
func (r *PolicyObjectSecurityGeolocationListProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PolicyObjectSecurityGeolocationList

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
func (r *PolicyObjectSecurityGeolocationListProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PolicyObjectSecurityGeolocationList

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
func (r *PolicyObjectSecurityGeolocationListProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// End of section. //template:end import

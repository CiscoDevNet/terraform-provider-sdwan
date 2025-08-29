//go:build ignore
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
	"strings"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
)
// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &{{camelCase .Name}}Resource{}
{{- if not .NoImport}}
var _ resource.ResourceWithImportState = &{{camelCase .Name}}Resource{}
{{- end}}

func New{{camelCase .Name}}Resource() resource.Resource {
	return &{{camelCase .Name}}Resource{}
}

type {{camelCase .Name}}Resource struct {
	client *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *{{camelCase .Name}}Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_{{snakeCase .Name}}"
}

func (r *{{camelCase .Name}}Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("{{.ResDescription}}"){{- if .MinimumVersion}}.AddMinimumVersionDescription("{{.MinimumVersion}}"){{- end}}.String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			{{- if .HasVersion}}
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			{{- end}}
			{{- if .TypeValue}}
			"type": schema.StringAttribute{
				MarkdownDescription: "Type",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			{{- end}}
			{{- range  .Attributes}}
			{{- if not .Value}}
			"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}{{if .ConditionalAttribute.Name}}, Attribute conditional on `{{.ConditionalAttribute.Name}}` being equal to `{{.ConditionalAttribute.Value}}`{{end}}")
					{{- if len .EnumValues -}}
					.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
					{{- end -}}
					{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
					.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
					{{- end -}}
					{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
					.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
					{{- end -}}
					{{- if .DefaultValue -}}
					.AddDefaultValueDescription("{{.DefaultValue}}")
					{{- end -}}
					.String,
				{{- if isListSet .}}
				ElementType:         types.{{.ElementType}}Type,
				{{- else if eq .Type "Versions"}}
				ElementType:         types.StringType,
				{{- end}}
				{{- if .Mandatory}}
				Required:            true,
				{{- else}}
				Optional:            true,
				{{- end}}
				{{- if len .DefaultValue}}
				Computed:            true,
				{{- end}}
				{{- if and (len .EnumValues) (not .IgnoreEnum)}}
				Validators: []validator.String{
					stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
				},
				{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
				Validators: []validator.String{
					{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
					stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
					{{- end}}
					{{- range .StringPatterns}}
					stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
					{{- end}}
				},
				{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
				Validators: []validator.Int64{
					int64validator.Between({{.MinInt}}, {{.MaxInt}}),
				},
				{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
				Validators: []validator.Float64{
					float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
				},
				{{- end}}
				{{- if and (len .DefaultValue) (eq .Type "Int64")}}
				Default:             int64default.StaticInt64({{.DefaultValue}}),
				{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
				Default:             booldefault.StaticBool({{.DefaultValue}}),
				{{- else if and (len .DefaultValue) (eq .Type "String")}}
				Default:             stringdefault.StaticString("{{.DefaultValue}}"),
				{{- end}}
				{{- if .RequiresReplace}}
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				{{- end}}
				{{- if isNestedListSet .}}
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						{{- range  .Attributes}}
						{{- if not .Value}}
						"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
							MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}{{if .ConditionalAttribute.Name}}, Attribute conditional on `{{.ConditionalAttribute.Name}}` being equal to `{{.ConditionalAttribute.Value}}`{{end}}")
								{{- if len .EnumValues -}}
								.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
								{{- end -}}
								{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
								.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
								{{- end -}}
								{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
								.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
								{{- end -}}
								{{- if .DefaultValue -}}
								.AddDefaultValueDescription("{{.DefaultValue}}")
								{{- end -}}
								.String,
							{{- if isListSet .}}
							ElementType:         types.{{.ElementType}}Type,
							{{- else if eq .Type "Versions"}}
							ElementType:         types.StringType,
							{{- end}}
							{{- if .Mandatory}}
							Required:            true,
							{{- else}}
							Optional:            true,
							{{- end}}
							{{- if len .DefaultValue}}
							Computed:            true,
							{{- end}}
							{{- if and (len .EnumValues) (not .IgnoreEnum)}}
							Validators: []validator.String{
								stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
							},
							{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
							Validators: []validator.String{
								{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
								stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
								{{- end}}
								{{- range .StringPatterns}}
								stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
								{{- end}}
							},
							{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
							Validators: []validator.Int64{
								int64validator.Between({{.MinInt}}, {{.MaxInt}}),
							},
							{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
							Validators: []validator.Float64{
								float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
							},
							{{- end}}
							{{- if and (len .DefaultValue) (eq .Type "Int64")}}
							Default:             int64default.StaticInt64({{.DefaultValue}}),
							{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
							Default:             booldefault.StaticBool({{.DefaultValue}}),
							{{- else if and (len .DefaultValue) (eq .Type "String")}}
							Default:             stringdefault.StaticString("{{.DefaultValue}}"),
							{{- end}}
							{{- if .RequiresReplace}}
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							{{- end}}
							{{- if isNestedListSet .}}
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									{{- range  .Attributes}}
									{{- if not .Value}}
									"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
										MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}{{if .ConditionalAttribute.Name}}, Attribute conditional on `{{.ConditionalAttribute.Name}}` being equal to `{{.ConditionalAttribute.Value}}`{{end}}")
											{{- if len .EnumValues -}}
											.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
											{{- end -}}
											{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
											.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
											{{- end -}}
											{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
											.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
											{{- end -}}
											{{- if .DefaultValue -}}
											.AddDefaultValueDescription("{{.DefaultValue}}")
											{{- end -}}
											.String,
										{{- if isListSet .}}
										ElementType:         types.{{.ElementType}}Type,
										{{- else if eq .Type "Versions"}}
										ElementType:         types.StringType,
										{{- end}}
										{{- if .Mandatory}}
										Required:            true,
										{{- else}}
										Optional:            true,
										{{- end}}
										{{- if len .DefaultValue}}
										Computed:            true,
										{{- end}}
										{{- if and (len .EnumValues) (not .IgnoreEnum)}}
										Validators: []validator.String{
											stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
										},
										{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
										Validators: []validator.String{
											{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
											stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
											{{- end}}
											{{- range .StringPatterns}}
											stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
											{{- end}}
										},
										{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
										Validators: []validator.Int64{
											int64validator.Between({{.MinInt}}, {{.MaxInt}}),
										},
										{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
										Validators: []validator.Float64{
											float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
										},
										{{- end}}
										{{- if and (len .DefaultValue) (eq .Type "Int64")}}
										Default:             int64default.StaticInt64({{.DefaultValue}}),
										{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
										Default:             booldefault.StaticBool({{.DefaultValue}}),
										{{- else if and (len .DefaultValue) (eq .Type "String")}}
										Default:             stringdefault.StaticString("{{.DefaultValue}}"),
										{{- end}}
										{{- if .RequiresReplace}}
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplace(),
										},
										{{- end}}
										{{- if isNestedListSet .}}
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												{{- range  .Attributes}}
												{{- if not .Value}}
												"{{.TfName}}": schema.{{if isNestedListSet .}}{{.Type}}Nested{{else if isList .}}List{{else if isSet .}}Set{{else if eq .Type "Versions"}}List{{else if eq .Type "Version"}}Int64{{else}}{{.Type}}{{end}}Attribute{
													MarkdownDescription: helpers.NewAttributeDescription("{{.Description}}{{if .ConditionalAttribute.Name}}, Attribute conditional on `{{.ConditionalAttribute.Name}}` being equal to `{{.ConditionalAttribute.Value}}`{{end}}")
														{{- if len .EnumValues -}}
														.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
														{{- end -}}
														{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
														.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
														{{- end -}}
														{{- if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0) -}}
														.AddFloatRangeDescription({{.MinFloat}}, {{.MaxFloat}})
														{{- end -}}
														{{- if .DefaultValue -}}
														.AddDefaultValueDescription("{{.DefaultValue}}")
														{{- end -}}
														.String,
													{{- if isListSet .}}
													ElementType:         types.{{.ElementType}}Type,
													{{- else if eq .Type "Versions"}}
													ElementType:         types.StringType,
													{{- end}}
													{{- if .Mandatory}}
													Required:            true,
													{{- else}}
													Optional:            true,
													{{- end}}
													{{- if len .DefaultValue}}
													Computed:            true,
													{{- end}}
													{{- if and (len .EnumValues) (not .IgnoreEnum)}}
													Validators: []validator.String{
														stringvalidator.OneOf({{range .EnumValues}}"{{.}}", {{end}}),
													},
													{{- else if or (len .StringPatterns) (ne .StringMinLength 0) (ne .StringMaxLength 0) }}
													Validators: []validator.String{
														{{- if or (ne .StringMinLength 0) (ne .StringMaxLength 0)}}
														stringvalidator.LengthBetween({{.StringMinLength}}, {{.StringMaxLength}}),
														{{- end}}
														{{- range .StringPatterns}}
														stringvalidator.RegexMatches(regexp.MustCompile(`{{.}}`), ""),
														{{- end}}
													},
													{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
													Validators: []validator.Int64{
														int64validator.Between({{.MinInt}}, {{.MaxInt}}),
													},
													{{- else if or (ne .MinFloat 0.0) (ne .MaxFloat 0.0)}}
													Validators: []validator.Float64{
														float64validator.Between({{.MinFloat}}, {{.MaxFloat}}),
													},
													{{- end}}
													{{- if and (len .DefaultValue) (eq .Type "Int64")}}
													Default:             int64default.StaticInt64({{.DefaultValue}}),
													{{- else if and (len .DefaultValue) (eq .Type "Bool")}}
													Default:             booldefault.StaticBool({{.DefaultValue}}),
													{{- else if and (len .DefaultValue) (eq .Type "String")}}
													Default:             stringdefault.StaticString("{{.DefaultValue}}"),
													{{- end}}
													{{- if .RequiresReplace}}
													PlanModifiers: []planmodifier.String{
														stringplanmodifier.RequiresReplace(),
													},
													{{- end}}
												},
												{{- end}}
												{{- end}}
											},
										},
										{{- if or (ne .MinList 0) (ne .MaxList 0)}}
										{{- if eq .Type "Set"}}
										Validators: []validator.Set{
											{{- if ne .MinList 0}}
											setvalidator.SizeAtLeast({{.MinList}}),
											{{- end}}
											{{- if ne .MaxList 0}}
											setvalidator.SizeAtMost({{.MaxList}}),
											{{- end}}
										},
										{{- else}}
										Validators: []validator.List{
											{{- if ne .MinList 0}}
											listvalidator.SizeAtLeast({{.MinList}}),
											{{- end}}
											{{- if ne .MaxList 0}}
											listvalidator.SizeAtMost({{.MaxList}}),
											{{- end}}
										},
										{{- end}}
										{{- end}}
										{{- end}}
									},
									{{- end}}
									{{- end}}
								},
							},
							{{- if or (ne .MinList 0) (ne .MaxList 0)}}
							{{- if eq .Type "Set"}}
							Validators: []validator.Set{
								{{- if ne .MinList 0}}
								setvalidator.SizeAtLeast({{.MinList}}),
								{{- end}}
								{{- if ne .MaxList 0}}
								setvalidator.SizeAtMost({{.MaxList}}),
								{{- end}}
							},
							{{- else}}
							Validators: []validator.List{
								{{- if ne .MinList 0}}
								listvalidator.SizeAtLeast({{.MinList}}),
								{{- end}}
								{{- if ne .MaxList 0}}
								listvalidator.SizeAtMost({{.MaxList}}),
								{{- end}}
							},
							{{- end}}
							{{- end}}
							{{- end}}
						},
						{{- end}}
						{{- end}}
					},
				},
				{{- if or (ne .MinList 0) (ne .MaxList 0)}}
				{{- if eq .Type "Set"}}
				Validators: []validator.Set{
					{{- if ne .MinList 0}}
					setvalidator.SizeAtLeast({{.MinList}}),
					{{- end}}
					{{- if ne .MaxList 0}}
					setvalidator.SizeAtMost({{.MaxList}}),
					{{- end}}
				},
				{{- else}}
				Validators: []validator.List{
					{{- if ne .MinList 0}}
					listvalidator.SizeAtLeast({{.MinList}}),
					{{- end}}
					{{- if ne .MaxList 0}}
					listvalidator.SizeAtMost({{.MaxList}}),
					{{- end}}
				},
				{{- end}}
				{{- end}}
				{{- end}}
			},
			{{- end}}
			{{- end}}
		},
	}
}

func (r *{{camelCase .Name}}Resource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}
// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *{{camelCase .Name}}Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan {{camelCase .Name}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	{{if hasName .Attributes}}
	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))
	{{end}}

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post({{if .PostRestEndpoint}}"{{.PostRestEndpoint}}"{{else}}plan.getPath(){{end}}, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}
	{{- if not .RemoveId}}
		{{- if .IdAttribute}}
		plan.Id = types.StringValue(res.Get("{{.IdAttribute}}").String())
		{{- else if .IdFromQueryPath}}
			{{- $id := getId .Attributes}}
	res, err = r.client.Get(plan.getPath())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}
	plan.Id = types.StringValue(res.Get("{{.IdFromQueryPath}}.#({{if $id.ResponseModelName}}{{$id.ResponseModelName}}{{else}}{{$id.ModelName}}{{end}}==\""+ plan.{{toGoName $id.TfName}}.Value{{$id.Type}}() +"\").{{if .IdFromQueryPathAttribute}}{{.IdFromQueryPathAttribute}}{{else}}id{{end}}").String())
	{{- end}}
	{{- end}}
	{{- if .HasVersion}}
	plan.Version = types.Int64Value(0)
	{{- end}}
	{{- if .TypeValue}}
	plan.Type = types.StringValue("{{.TypeValue}}")
	{{- end}}
	
	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", {{if hasName .Attributes}}plan.Name.ValueString(){{else}}plan.Id.ValueString(){{end}}))
	
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}
// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *{{camelCase .Name}}Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", {{if hasName .Attributes}}state.Name.ValueString(){{else}}state.Id.ValueString(){{end}}))

	res, err := r.client.Get({{if .GetRestEndpoint}}"{{.GetRestEndpoint}}"{{else}}state.getPath(){{end}}{{if not .RemoveId}} + url.QueryEscape(state.Id.ValueString()){{end}})
	if strings.Contains(res.Get("error.message").String(), "Failed to find specified resource") || strings.Contains(res.Get("error.message").String(), "Invalid template type") || strings.Contains(res.Get("error.message").String(), "Template definition not found") || strings.Contains(res.Get("error.message").String(), "Invalid Profile Id") || strings.Contains(res.Get("error.message").String(), "Invalid feature Id") || strings.Contains(res.Get("error.message").String(), "Invalid config group passed") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.fromBody(ctx, res)
	
	{{- if and (not .NoImport) (not (isUx20Feature .)) (or (.HasVersion) (.TypeValue) (hasVersionAttribute .Attributes))}}
	imp, diags := helpers.IsFlagImporting(ctx, req)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	if imp {
		state.processImport(ctx)
	}
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", {{if hasName .Attributes}}state.Name.ValueString(){{else}}state.Id.ValueString(){{end}}))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}
// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *{{camelCase .Name}}Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state {{camelCase .Name}}

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", {{if hasName .Attributes}}plan.Name.ValueString(){{else}}plan.Id.ValueString(){{end}}))

	if plan.hasChanges(ctx, &state) {
		body := plan.toBody(ctx)
		r.updateMutex.Lock()
		res, err := r.client.Put(plan.getPath(){{if not .RemoveId}} + url.QueryEscape(plan.Id.ValueString()){{end}}, body)
		r.updateMutex.Unlock()
		if err != nil {
			if strings.Contains(res.Get("error.message").String(), "Failed to acquire lock") {
				resp.Diagnostics.AddWarning("Client Warning", "Failed to modify policy due to policy being locked by another change. Policy changes will not be applied. Re-run 'terraform apply' to try again.")
			} else if strings.Contains(res.Get("error.message").String(), "Template locked in edit mode") {
				resp.Diagnostics.AddWarning("Client Warning", "Failed to modify template due to template being locked by another change. Template changes will not be applied. Re-run 'terraform apply' to try again.")
			} else {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
				return
			}
		}
	} else {
		tflog.Debug(ctx, fmt.Sprintf("%s: No changes detected", {{if hasName .Attributes}}plan.Name.ValueString(){{else}}plan.Id.ValueString(){{end}}))
	}

	{{- if .HasVersion}}
	plan.Version = types.Int64Value(state.Version.ValueInt64()+1)
	{{- end}}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", {{if hasName .Attributes}}plan.Name.ValueString(){{else}}plan.Id.ValueString(){{end}}))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}
// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *{{camelCase .Name}}Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", {{if hasName .Attributes}}state.Name.ValueString(){{else}}state.Id.ValueString(){{end}}))
	{{if .GetBeforeDelete}}
	_, _ = r.client.Get(state.getPath())
	{{- end}}
	{{- if .DeleteMutex}}
	r.updateMutex.Lock()
	{{- end}}
	{{- if not .RemoveId}}
	res, err := r.client.Delete(state.getPath() + url.QueryEscape(state.Id.ValueString()))
	{{- else}}
	res, err := r.client.Delete(state.getPath())
	{{- end}}
	{{- if .DeleteMutex}}
	r.updateMutex.Unlock()
	{{- end}}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", {{if hasName .Attributes}}state.Name.ValueString(){{else}}state.Id.ValueString(){{end}}))

	resp.State.RemoveResource(ctx)
}
// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
{{- if not .NoImport}}
func (r *{{camelCase .Name}}Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	{{- if hasReference .Attributes}}
	count := {{ countReferences .Attributes}}
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "{{snakeCase .Name}}_id"{{range .Attributes}}{{if .Reference}} + ",{{.TfName}}"{{end}}{{end}}
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q", pattern, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	{{- $count := 0}}
	{{- range .Attributes}}{{- if .Reference}}{{$count = add $count 1}}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("{{.TfName}}"), parts[{{$count}}])...)
	{{- end}}{{- end}}
	{{- else}}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	{{- end}}

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}
{{- end}}
// End of section. //template:end import

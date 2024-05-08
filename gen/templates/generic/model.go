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
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)
// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin types
{{- $name := camelCase .Name}}
type {{camelCase .Name}} struct {
	Id types.String `tfsdk:"id"`
{{- if .HasVersion}}
	Version types.Int64 `tfsdk:"version"`
{{- end}}
{{- if .TypeValue}}
Type types.String `tfsdk:"type"`
{{- end}}
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Versions"}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Version"}}
	{{toGoName .TfName}} types.Int64 `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}

{{ range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if isNestedListSet .}}
type {{$name}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Versions"}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Version"}}
	{{toGoName .TfName}} types.Int64 `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{ end}}

{{ range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if isNestedListSet .}}
{{ range .Attributes}}
{{- if not .Value}}
{{- $childChildName := toGoName .TfName}}
{{- if isNestedListSet .}}
type {{$name}}{{$childName}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{$name}}{{$childName}}{{$childChildName}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Versions"}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Version"}}
	{{toGoName .TfName}} types.Int64 `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{ end}}

{{ range .Attributes}}
{{- if not .Value}}
{{- $childName := toGoName .TfName}}
{{- if isNestedListSet .}}
{{ range .Attributes}}
{{- if not .Value}}
{{- $childChildName := toGoName .TfName}}
{{- if isNestedListSet .}}
{{ range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
type {{$name}}{{$childName}}{{$childChildName}}{{toGoName .TfName}} struct {
{{- range .Attributes}}
{{- if not .Value}}
{{- if eq .Type "Versions"}}
	{{toGoName .TfName}} types.List `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "Version"}}
	{{toGoName .TfName}} types.Int64 `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- end}}
{{- end}}
{{- end}}
}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{- end}}
{{ end}}
// End of section. //template:end types

// Section below is generated&owned by "gen/generator.go". //template:begin getPath
func (data {{camelCase .Name}}) getPath() string {
	{{- if hasReference .Attributes}}
		return fmt.Sprintf("{{.RestEndpoint}}"{{range .Attributes}}{{if .Reference}}, url.QueryEscape(data.{{toGoName .TfName}}.Value{{.Type}}()){{end}}{{end}})
	{{- else}}
		return "{{.RestEndpoint}}"
	{{- end}}
}
// End of section. //template:end getPath

// Section below is generated&owned by "gen/generator.go". //template:begin toBody
func (data {{camelCase .Name}}) toBody(ctx context.Context) string {
	body := ""
	{{- range .Attributes}}
	{{- if .Value}}
	if true{{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(data.{{toGoName .ConditionalListLength}}) > 0{{end}} {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
	}
	{{- else if and (not .TfOnly) (not .Reference)}}
	{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
	if {{if not .AlwaysInclude}}!data.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(data.{{toGoName .ConditionalListLength}}) > 0{{end}} {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}data.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
	}
	{{- else if eq .Type "Bool"}}
	if {{if not .AlwaysInclude}}!data.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(data.{{toGoName .ConditionalListLength}}) > 0{{end}} {
		if {{.BoolEmptyString}} && data.{{toGoName .TfName}}.Value{{.Type}}() {
			body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", "")
		}{{if not .BoolEmptyString}} else {
			body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}data.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
		}
		{{- end}}
	}
	{{- else if isListSet .}}
	if {{if not .AlwaysInclude}}!data.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(data.{{toGoName .ConditionalListLength}}) > 0{{end}} {
		var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
	}
	{{- else if isNestedListSet .}}
	if true{{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(data.{{toGoName .ConditionalListLength}}) > 0{{end}} {
		body, _ = sjson.Set(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
		for _, item := range data.{{toGoName .TfName}} {
			itemBody := ""
			{{- range .Attributes}}
			{{- if .Value}}
			if true{{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(item.{{toGoName .ConditionalListLength}}) > 0{{end}} {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
			}
			{{- else if and (not .TfOnly) (not .Reference)}}
			{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
			if {{if not .AlwaysInclude}}!item.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(item.{{toGoName .ConditionalListLength}}) > 0{{end}} {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}item.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
			}
			{{- else if eq .Type "Bool"}}
			if {{if not .AlwaysInclude}}!item.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(item.{{toGoName .ConditionalListLength}}) > 0{{end}} {
				if {{.BoolEmptyString}} && item.{{toGoName .TfName}}.Value{{.Type}}() {
					itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", "")
				}{{if not .BoolEmptyString}} else {
					itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}item.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
				}
				{{- end}}
			}
			{{- else if isListSet .}}
			if {{if not .AlwaysInclude}}!item.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(item.{{toGoName .ConditionalListLength}}) > 0{{end}} {
				var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
				item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
			}
			{{- else if isNestedListSet .}}
			if true{{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(item.{{toGoName .ConditionalListLength}}) > 0{{end}} {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
				for _, childItem := range item.{{toGoName .TfName}} {
					itemChildBody := ""
					{{- range .Attributes}}
					{{- if .Value}}
					if true{{if ne .ConditionalAttribute.Name ""}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(childItem.{{toGoName .ConditionalListLength}}) > 0{{end}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
					}
					{{- else if and (not .TfOnly) (not .Reference)}}
					{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
					if {{if not .AlwaysInclude}}!childItem.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(childItem.{{toGoName .ConditionalListLength}}) > 0{{end}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}childItem.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
					}
					{{- else if eq .Type "Bool"}}
					if {{if not .AlwaysInclude}}!childItem.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(childItem.{{toGoName .ConditionalListLength}}) > 0{{end}} {
						if {{.BoolEmptyString}} && childItem.{{toGoName .TfName}}.Value{{.Type}}() {
							itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", "")
						}{{if not .BoolEmptyString}} else {
							itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}childItem.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
						}
						{{- end}}
					}
					{{- else if isListSet .}}
					if {{if not .AlwaysInclude}}!childItem.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(childItem.{{toGoName .ConditionalListLength}}) > 0{{end}} {
						var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
						childItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
					}
					{{- else if isNestedListSet .}}
					if true{{if ne .ConditionalAttribute.Name ""}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(childItem.{{toGoName .ConditionalListLength}}) > 0{{end}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{})
						for _, childChildItem := range childItem.{{toGoName .TfName}} {
							itemChildChildBody := ""
							{{- range .Attributes}}
							{{- if .Value}}
							if true{{if ne .ConditionalAttribute.Name ""}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(childChildItem.{{toGoName .ConditionalListLength}}) > 0{{end}} {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
							}
							{{- else if and (not .TfOnly) (not .Reference)}}
							{{- if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64")}}
							if {{if not .AlwaysInclude}}!childChildItem.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(childChildItem.{{toGoName .ConditionalListLength}}) > 0{{end}} {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}childChildItem.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
							}
							{{- else if eq .Type "Bool"}}
							if {{if not .AlwaysInclude}}!childChildItem.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(childChildItem.{{toGoName .ConditionalListLength}}) > 0{{end}} {
								if {{.BoolEmptyString}} && childChildItem.{{toGoName .TfName}}.Value{{.Type}}() {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", "")
								}{{if not .BoolEmptyString}} else {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{if .ModelTypeString}}fmt.Sprint({{end}}childChildItem.{{toGoName .TfName}}.Value{{.Type}}(){{if .ModelTypeString}}){{end}})
								}
								{{- end}}
							}
							{{- else if isListSet .}}
							if {{if not .AlwaysInclude}}!childChildItem.{{toGoName .TfName}}.IsNull(){{else}}true{{end}}{{if ne .ConditionalAttribute.Name ""}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{else if ne .ConditionalListLength ""}} && len(childChildItem.{{toGoName .ConditionalListLength}}) > 0{{end}} {
								var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
								childChildItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", values)
							}
							{{- end}}
							{{- end}}
							{{- end}}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildChildBody)
						}
					}
					{{- end}}
					{{- end}}
					{{- end}}
					itemBody, _ = sjson.SetRaw(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildBody)
				}
			}
			{{- end}}
			{{- end}}
			{{- end}}
			body, _ = sjson.SetRaw(body, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemBody)
		}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	return body
}
// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *{{camelCase .Name}}) fromBody(ctx context.Context, res gjson.Result) {
	{{- if hasVersionAttribute .Attributes}}
	state := *data
	{{- end}}
	{{- range .Attributes}}
	{{- if and (not .TfOnly) (not .Value) (not .Reference)}}
	{{- $cname := toGoName .TfName}}
	{{- if eq .Type "String"}}
	if value := res.Get("{{getResponseModelPath .}}"); value.Exists(){{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && value.String() != ""{{end}} {
		data.{{toGoName .TfName}} = types.StringValue(value.String())
	} else {
		data.{{toGoName .TfName}} = types.StringNull()
	}
	{{- else if eq .Type "Int64"}}
	if value := res.Get("{{getResponseModelPath .}}"); value.Exists(){{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && value.String() != ""{{end}} {
		data.{{toGoName .TfName}} = types.Int64Value(value.Int())
	} else {
		data.{{toGoName .TfName}} = types.Int64Null()
	}
	{{- else if eq .Type "Float64"}}
	if value := res.Get("{{getResponseModelPath .}}"); value.Exists(){{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && value.String() != ""{{end}} {
		data.{{toGoName .TfName}} = types.Float64Value(value.Float())
	} else {
		data.{{toGoName .TfName}} = types.Float64Null()
	}
	{{- else if eq .Type "Bool"}}
	if value := res.Get("{{getResponseModelPath .}}"); value.Exists(){{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && value.String() != ""{{end}} {
		if {{.BoolEmptyString}} && value.String() == "" {
			data.{{toGoName .TfName}} = types.BoolValue(true)
		}{{if not .BoolEmptyString}} else {
			data.{{toGoName .TfName}} = types.BoolValue(value.Bool())
		}
		{{- end}}
	} else {
		data.{{toGoName .TfName}} = types.BoolNull()
	}
	{{- else if isListSet .}}
	if value := res.Get("{{getResponseModelPath .}}"); value.Exists(){{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && value.String() != ""{{end}} {
		data.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(value.Array())
	} else {
		data.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
	}
	{{- else if isNestedListSet .}}
	if value := res.Get("{{getResponseModelPath .}}"); value.Exists() && len(value.Array()) > 0{{if ne .ConditionalAttribute.Name ""}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			{{- range .Attributes}}
			{{- $ccname := toGoName .TfName}}
			{{- if and (not .TfOnly) (not .Value) (not .Reference)}}
			{{- if eq .Type "String"}}
			if cValue := v.Get("{{getResponseModelPath .}}"); cValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cValue.String() != ""{{end}} {
				item.{{toGoName .TfName}} = types.StringValue(cValue.String())
			} else {
				item.{{toGoName .TfName}} = types.StringNull()
			}
			{{- else if eq .Type "Int64"}}
			if cValue := v.Get("{{getResponseModelPath .}}"); cValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cValue.String() != ""{{end}} {
				item.{{toGoName .TfName}} = types.Int64Value(cValue.Int())
			} else {
				item.{{toGoName .TfName}} = types.Int64Null()
			}
			{{- else if eq .Type "Float64"}}
			if cValue := v.Get("{{getResponseModelPath .}}"); cValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cValue.String() != ""{{end}} {
				item.{{toGoName .TfName}} = types.Float64Value(cValue.Float())
			} else {
				item.{{toGoName .TfName}} = types.Float64Null()
			}
			{{- else if eq .Type "Bool"}}
			if cValue := v.Get("{{getResponseModelPath .}}"); cValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cValue.String() != ""{{end}} {
				if {{.BoolEmptyString}} && cValue.String() == "" {
					item.{{toGoName .TfName}} = types.BoolValue(true)
				}{{if not .BoolEmptyString}} else {
					item.{{toGoName .TfName}} = types.BoolValue(cValue.Bool())
				}
				{{- end}}
			} else {
				item.{{toGoName .TfName}} = types.BoolNull()
			}
			{{- else if isListSet .}}
			if cValue := v.Get("{{getResponseModelPath .}}"); cValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cValue.String() != ""{{end}} {
				item.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(cValue.Array())
			} else {
				item.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
			}
			{{- else if isNestedListSet .}}
			if cValue := v.Get("{{getResponseModelPath .}}"); cValue.Exists() && len(cValue.Array()) > 0{{if ne .ConditionalAttribute.Name ""}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
				item.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{toGoName .TfName}}, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := {{$name}}{{$cname}}{{toGoName .TfName}}{}
					{{- range .Attributes}}
					{{- if and (not .TfOnly) (not .Value) (not .Reference)}}
					{{- if eq .Type "String"}}
					if ccValue := cv.Get("{{getResponseModelPath .}}"); ccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && cItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && ccValue.String() != ""{{end}} {
						cItem.{{toGoName .TfName}} = types.StringValue(ccValue.String())
					} else {
						cItem.{{toGoName .TfName}} = types.StringNull()
					}
					{{- else if eq .Type "Int64"}}
					if ccValue := cv.Get("{{getResponseModelPath .}}"); ccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && cItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && ccValue.String() != ""{{end}} {
						cItem.{{toGoName .TfName}} = types.Int64Value(ccValue.Int())
					} else {
						cItem.{{toGoName .TfName}} = types.Int64Null()
					}
					{{- else if eq .Type "Float64"}}
					if ccValue := cv.Get("{{getResponseModelPath .}}"); ccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && cItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && ccValue.String() != ""{{end}} {
						cItem.{{toGoName .TfName}} = types.Float64Value(ccValue.Float())
					} else {
						cItem.{{toGoName .TfName}} = types.Float64Null()
					}
					{{- else if eq .Type "Bool"}}
					if ccValue := cv.Get("{{getResponseModelPath .}}"); ccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && cItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && ccValue.String() != ""{{end}} {
						if {{.BoolEmptyString}} && ccValue.String() == "" {
							cItem.{{toGoName .TfName}} = types.BoolValue(true)
						}{{if not .BoolEmptyString}} else {
							cItem.{{toGoName .TfName}} = types.BoolValue(ccValue.Bool())
						}
						{{- end}}
					} else {
						cItem.{{toGoName .TfName}} = types.BoolNull()
					}
					{{- else if isListSet .}}
					if ccValue := cv.Get("{{getResponseModelPath .}}"); ccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && cItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && ccValue.String() != ""{{end}} {
						cItem.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(ccValue.Array())
					} else {
						cItem.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
					}
					{{- else if isNestedListSet .}}
					if ccValue := cv.Get("{{getResponseModelPath .}}"); ccValue.Exists() && len(ccValue.Array()) > 0{{if ne .ConditionalAttribute.Name ""}} && cItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}} {
						cItem.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := {{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}{}
							{{- range .Attributes}}
							{{- if and (not .TfOnly) (not .Value) (not .Reference)}}
							{{- if eq .Type "String"}}
							if cccValue := ccv.Get("{{getResponseModelPath .}}"); cccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && ccItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cccValue.String() != ""{{end}} {
								ccItem.{{toGoName .TfName}} = types.StringValue(cccValue.String())
							} else {
								ccItem.{{toGoName .TfName}} = types.StringNull()
							}
							{{- else if eq .Type "Int64"}}
							if cccValue := ccv.Get("{{getResponseModelPath .}}"); cccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && ccItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cccValue.String() != ""{{end}} {
								ccItem.{{toGoName .TfName}} = types.Int64Value(cccValue.Int())
							} else {
								ccItem.{{toGoName .TfName}} = types.Int64Null()
							}
							{{- else if eq .Type "Float64"}}
							if cccValue := ccv.Get("{{getResponseModelPath .}}"); cccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && ccItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cccValue.String() != ""{{end}} {
								ccItem.{{toGoName .TfName}} = types.Float64Value(cccValue.Float())
							} else {
								ccItem.{{toGoName .TfName}} = types.Float64Null()
							}
							{{- else if eq .Type "Bool"}}
							if cccValue := ccv.Get("{{getResponseModelPath .}}"); cccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && ccItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cccValue.String() != ""{{end}} {
								if {{.BoolEmptyString}} && cccValue.String() == "" {
									ccItem.{{toGoName .TfName}} = types.BoolValue(true)
								}{{if not .BoolEmptyString}} else {
									ccItem.{{toGoName .TfName}} = types.BoolValue(cccValue.Bool())
								}
								{{- end}}
							} else {
								ccItem.{{toGoName .TfName}} = types.BoolNull()
							}
							{{- else if isListSet .}}
							if cccValue := ccv.Get("{{getResponseModelPath .}}"); cccValue.Exists(){{if ne .ConditionalAttribute.Name ""}} && ccItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}"{{end}}{{if .AlwaysInclude}} && cccValue.String() != ""{{end}} {
								ccItem.{{toGoName .TfName}} = helpers.Get{{.ElementType}}{{.Type}}(cccValue.Array())
							} else {
								ccItem.{{toGoName .TfName}} = types.{{.Type}}Null(types.{{.ElementType}}Type)
							}
							{{- end}}
							{{- end}}
							{{- end}}
							cItem.{{toGoName .TfName}} = append(cItem.{{toGoName .TfName}}, ccItem)
							return true
						})
					} else {
						if len(cItem.{{toGoName .TfName}}) > 0 {
							cItem.{{toGoName .TfName}} = []{{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}{}
						}
					}
					{{- end}}
					{{- end}}
					{{- end}}
					item.{{toGoName .TfName}} = append(item.{{toGoName .TfName}}, cItem)
					return true
				})
			} else {
				if len(item.{{toGoName .TfName}}) > 0 {
					item.{{toGoName .TfName}} = []{{$name}}{{$cname}}{{toGoName .TfName}}{}
				}
			}
			{{- end}}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
	} else {
		if len(data.{{toGoName .TfName}}) > 0 {
			data.{{toGoName .TfName}} = []{{$name}}{{toGoName .TfName}}{}
		}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if hasVersionAttribute .Attributes}}
	data.updateVersions(ctx, &state)
	{{- end}}
}
// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin hasChanges
func (data *{{camelCase .Name}}) hasChanges(ctx context.Context, state *{{camelCase .Name}}) bool {
	hasChanges := false
	{{- range .Attributes}}
	{{- $name := toGoName .TfName}}
	{{- if and (not .Value) (not .TfOnly)}}
	{{- if not (isNestedListSet .)}}
	if !data.{{toGoName .TfName}}.Equal(state.{{toGoName .TfName}}) {
		hasChanges = true
	}
	{{- else}}
	if len(data.{{toGoName .TfName}}) != len(state.{{toGoName .TfName}}) {
		hasChanges = true
	} else {
		for i := range data.{{toGoName .TfName}} {
			{{- range .Attributes}}
			{{- $cname := toGoName .TfName}}
			{{- if and (not .Value) (not .TfOnly)}}
			{{- if not (isNestedListSet .)}}
			if !data.{{$name}}[i].{{toGoName .TfName}}.Equal(state.{{$name}}[i].{{toGoName .TfName}}) {
				hasChanges = true
			}
			{{- else}}
			if len(data.{{$name}}[i].{{toGoName .TfName}}) != len(state.{{$name}}[i].{{toGoName .TfName}}) {
				hasChanges = true
			} else {
				for ii := range data.{{$name}}[i].{{toGoName .TfName}} {
					{{- range .Attributes}}
					{{- $ccname := toGoName .TfName}}
					{{- if and (not .Value) (not .TfOnly)}}
					{{- if not (isNestedListSet .)}}
					if !data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}.Equal(state.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}) {
						hasChanges = true
					}
					{{- else}}
					if len(data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}) != len(state.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}) {
						hasChanges = true
					} else {
						for iii := range data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}} {
							{{- range .Attributes}}
							{{- if and (not .Value) (not .TfOnly)}}
							if !data.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}}.Equal(state.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}}) {
								hasChanges = true
							}
							{{- end}}
							{{- end}}
						}
					}
					{{- end}}
					{{- end}}
					{{- end}}
				}
			}
			{{- end}}
			{{- end}}
			{{- end}}
		}
	}
	{{- end}}
	{{- end}}
	{{- end}}
	return hasChanges
}
// End of section. //template:end hasChanges

// Section below is generated&owned by "gen/generator.go". //template:begin updateVersions
{{if hasVersionAttribute .Attributes}}
func (data *{{camelCase .Name}}) updateVersions(ctx context.Context, state *{{camelCase .Name}}) {
	{{- range .Attributes}}
	{{- $name := toGoName .TfName}}
	{{- if eq .Type "Version"}}
	data.{{toGoName .TfName}} = state.{{toGoName .TfName}}
	{{- else if and (isNestedListSet .) (hasVersionAttribute .Attributes)}}
	for i := range data.{{toGoName .TfName}} {
		dataKeys := [...]string{ {{range .Attributes}}{{if .Id}}fmt.Sprintf("%v", data.{{$name}}[i].{{toGoName .TfName}}.{{if isStringListSet .}}String{{else}}Value{{.Type}}{{end}}()), {{end}}{{end}} }
		stateIndex := -1
		for j := range state.{{toGoName .TfName}} {
			stateKeys := [...]string{ {{range .Attributes}}{{if .Id}}fmt.Sprintf("%v", state.{{$name}}[j].{{toGoName .TfName}}.{{if isStringListSet .}}String{{else}}Value{{.Type}}{{end}}()), {{end}}{{end}} }
			if dataKeys == stateKeys {
				stateIndex = j
                break
			}
		}
		{{- range .Attributes}}
		{{- $cname := toGoName .TfName}}
		{{- if eq .Type "Version"}}
		if stateIndex > -1 {
			data.{{$name}}[i].{{toGoName .TfName}} = state.{{$name}}[stateIndex].{{toGoName .TfName}}
		} else {
			data.{{$name}}[i].{{toGoName .TfName}} = types.Int64Null()
		}
		{{- else if eq .Type "Versions"}}
		if stateIndex > -1 && !state.{{$name}}[stateIndex].{{toGoName .TfName}}.IsNull() {
			data.{{$name}}[i].{{toGoName .TfName}} = state.{{$name}}[stateIndex].{{toGoName .TfName}}
		} else {
			data.{{$name}}[i].{{toGoName .TfName}} = types.ListNull(types.StringType)
		}
		{{- else if and (isNestedListSet .) (hasVersionAttribute .Attributes)}}
		for ii := range data.{{$name}}[i].{{toGoName .TfName}} {
			cDataKeys := [...]string{ {{range .Attributes}}{{if .Id}}fmt.Sprintf("%v", data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}}.{{if isStringListSet .}}String{{else}}Value{{.Type}}{{end}}()), {{end}}{{end}} }
			cStateIndex := -1
			if stateIndex > -1 {
				for jj := range state.{{$name}}[stateIndex].{{toGoName .TfName}} {
					cStateKeys := [...]string{ {{range .Attributes}}{{if .Id}}fmt.Sprintf("%v", state.{{$name}}[stateIndex].{{$cname}}[jj].{{toGoName .TfName}}.{{if isStringListSet .}}String{{else}}Value{{.Type}}{{end}}()), {{end}}{{end}} }
					if cDataKeys == cStateKeys {
						cStateIndex = jj
						break
					}
				}
			}
			{{- range .Attributes}}
			{{- $ccname := toGoName .TfName}}
			{{- if eq .Type "Version"}}
			if cStateIndex > -1 {
				data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}} = state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{toGoName .TfName}}
			} else {
				data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}} = types.Int64Null()
			}
			{{- else if eq .Type "Versions"}}
			if cStateIndex > -1 && !state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{toGoName .TfName}}.IsNull() {
				data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}} = state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{toGoName .TfName}}
			} else {
				data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}} = types.ListNull(types.StringType)
			}
			{{- else if and (isNestedListSet .) (hasVersionAttribute .Attributes)}}
			for iii := range data.{{$name}}[i].{{$cname}}[ii].{{toGoName .TfName}} {
				ccDataKeys := [...]string{ {{range .Attributes}}{{if .Id}}fmt.Sprintf("%v", data.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}}.{{if isStringListSet .}}String{{else}}Value{{.Type}}{{end}}()), {{end}}{{end}} }
				ccStateIndex := -1
				if cStateIndex > -1 {
					for jjj := range state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{toGoName .TfName}} {
						ccStateKeys := [...]string{ {{range .Attributes}}{{if .Id}}fmt.Sprintf("%v", state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{$ccname}}[jjj].{{toGoName .TfName}}.{{if isStringListSet .}}String{{else}}Value{{.Type}}{{end}}()), {{end}}{{end}} }
						if ccDataKeys == ccStateKeys {
							ccStateIndex = jjj
							break
						}
					}
				}
				{{- range .Attributes}}
				{{- if eq .Type "Version"}}
				if ccStateIndex > -1 {
					data.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}} = state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{$ccname}}[ccStateIndex].{{toGoName .TfName}}
				} else {
					data.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}} = types.Int64Null()
				}
				{{- else if eq .Type "Versions"}}
				if ccStateIndex > -1 && !state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{$ccname}}[ccStateIndex].{{toGoName .TfName}}.IsNull() {
					data.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}} = state.{{$name}}[stateIndex].{{$cname}}[cStateIndex].{{$ccname}}[ccStateIndex].{{toGoName .TfName}}
				} else {
					data.{{$name}}[i].{{$cname}}[ii].{{$ccname}}[iii].{{toGoName .TfName}} = types.ListNull(types.StringType)
				}
				{{- end}}
				{{- end}}
			}
			{{- end}}
			{{- end}}
		}
		{{- end}}
		{{- end}}
	}
	{{- end}}
	{{- end}}
}
{{end}}
// End of section. //template:end updateVersions

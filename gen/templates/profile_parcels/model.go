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
	Version types.Int64 `tfsdk:"version"`
	Name types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
{{- range .Attributes}}
{{- if not .Value}}
{{- if isNestedListSet .}}
	{{toGoName .TfName}} []{{$name}}{{toGoName .TfName}} `tfsdk:"{{.TfName}}"`
{{- else if eq .Type "StringInt64"}}
	{{toGoName .TfName}} types.String `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
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
{{- else if eq .Type "StringInt64"}}
	{{toGoName .TfName}} types.String `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
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
{{- else if eq .Type "StringInt64"}}
	{{toGoName .TfName}} types.String `tfsdk:"{{.TfName}}"`
{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
{{- end}}
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
	{{- if eq .Type "StringInt64"}}
	{{toGoName .TfName}} types.String `tfsdk:"{{.TfName}}"`
	{{- else}}
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
	{{- end}}
{{- if .Variable}}
	{{toGoName .TfName}}Variable types.String `tfsdk:"{{.TfName}}_variable"`
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

// Section below is generated&owned by "gen/generator.go". //template:begin getModel
func (data {{camelCase .Name}}) getModel() string {
	return "{{.Model}}"
}
// End of section. //template:end getModel

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
	body, _ = sjson.Set(body, "name", data.Name.ValueString())
	body, _ = sjson.Set(body, "description", data.Description.ValueString())
	{{- if .TypeValue}}
	body, _ = sjson.Set(body, "data.type", "{{.TypeValue}}")
	{{- end}}
	path := "data."
	{{- range .Attributes}}
	{{- if .Value}}
	if true{{buildConditionalLogic .ConditionalAttribute $.Attributes "data"}} {
	body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", {{if .ValueType}}"{{.ValueType}}"{{else}}"default"{{end}})
	body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
	}
	{{- else if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "StringInt64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .TfOnly) (not .Reference)}}
	{{- if .NoOptionType}}
	if !data.{{toGoName .TfName}}.IsNull() {
		if true{{buildConditionalLogic .ConditionalAttribute $.Attributes "data"}} {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", data.{{toGoName .TfName}}.Value{{.Type}}())
		}
	}
	{{- else}}
	{{if .Variable}}
	if !data.{{toGoName .TfName}}Variable.IsNull() {
		if true{{buildConditionalLogic .ConditionalAttribute $.Attributes "data"}} {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "variable")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", data.{{toGoName .TfName}}Variable.ValueString())
		}
	} else {{end}}{{if and .DefaultValuePresent (not .ExcludeNull)}}if data.{{toGoName .TfName}}.IsNull() {{if and .DynamicDefault .DefaultValuePresent}}|| data.{{toGoName .TfName}}.ValueString() == "{{.DefaultValue}}"{{end}} {
		if true{{buildConditionalLogic .ConditionalAttribute $.Attributes "data"}} {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
		{{if or .DefaultValue .DefaultValueEmptyString}}body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}}){{end}}
		}
	} else {{else if .AlwaysIncludeParent }}if data.{{toGoName .TfName}}.IsNull() {
			body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}optionType", "default")
	} else {{else}}if !data.{{toGoName .TfName}}.IsNull(){{end}} {
		if true{{buildConditionalLogic .ConditionalAttribute $.Attributes "data"}} {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "global")
		{{- if isListSet .}}
		var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", values)
		{{- else}}
		{{- if eq .Type "StringInt64" }}
		if numValue, err := strconv.Atoi(data.{{toGoName .TfName}}.ValueString()); err != nil {
			body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", data.{{toGoName .TfName}}.ValueString())
		} else {
			body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", numValue)
		}
		{{- else}}
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", data.{{toGoName .TfName}}.Value{{.Type}}())
		{{- end}}
		{{- end}}
		}
	}
	{{- end}}
	{{- else if isNestedListSet .}}
	if true{{buildConditionalLogic .ConditionalAttribute $.Attributes "data"}} {
		{{if or .AlwaysInclude (and (not .MinList) (not .ExcludeNull))}}body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{}){{end}}
		for _, item := range data.{{toGoName .TfName}} {
			itemBody := ""
			{{- range .Attributes}}
			{{- if .Value}}
			if true{{buildConditionalLogic .ConditionalAttribute "item"}} {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", {{if .ValueType}}"{{.ValueType}}"{{else}}"default"{{end}})
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
			}
			{{- else if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "StringInt64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .TfOnly)}}
			{{if .Variable}}
			if !item.{{toGoName .TfName}}Variable.IsNull() {
				if true{{buildConditionalLogic .ConditionalAttribute "item"}} {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", item.{{toGoName .TfName}}Variable.ValueString())
				}
			} else {{end}}{{if and .DefaultValuePresent (not .ExcludeNull)}}if item.{{toGoName .TfName}}.IsNull() {{if and .DynamicDefault .DefaultValuePresent}}|| data.{{toGoName .TfName}}.ValueString() == "{{.DefaultValue}}"{{end}} {
				if true{{buildConditionalLogic .ConditionalAttribute "item"}} {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
				{{if or .DefaultValue .DefaultValueEmptyString}}itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}}){{end}}
				}
			} else {{else if .AlwaysInclude}}if item.{{toGoName .TfName}}.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}optionType", "default")
			} else {{else if .AlwaysIncludeParent}}if data.{{toGoName .TfName}}.IsNull() {
				body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}optionType", "default")
			} else {{else}}if !item.{{toGoName .TfName}}.IsNull(){{end}} {
				if true{{buildConditionalLogic .ConditionalAttribute "item"}} {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "global")
				{{- if isListSet .}}
				var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
				item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", values)
				{{- else}}
				{{- if eq .Type "StringInt64" }}
				if numValue, err := strconv.Atoi(item.{{toGoName .TfName}}.ValueString()); err != nil {
					itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", item.{{toGoName .TfName}}.ValueString())
				} else {
					itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", numValue)
				}
				{{- else}}
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", item.{{toGoName .TfName}}.Value{{.Type}}())
				{{- end}}
				{{- end}}
				}
			}
			{{- else if isNestedListSet .}}
				if true{{buildConditionalLogic .ConditionalAttribute "item"}} {
				{{if or .AlwaysInclude (and (not .MinList) (not .ExcludeNull))}}itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{}){{end}}
				for _, childItem := range item.{{toGoName .TfName}} {
					itemChildBody := ""
					{{- range .Attributes}}
					{{- if .Value}}
					if true{{buildConditionalLogic .ConditionalAttribute "childItem"}} {
					itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", {{if .ValueType}}"{{.ValueType}}"{{else}}"default"{{end}})
					itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
					}
					{{- else if or (eq .Type "String") (eq .Type "Int64") (eq .Type "StringInt64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)}}
					{{if .Variable}}
					if !childItem.{{toGoName .TfName}}Variable.IsNull() {
						if true{{buildConditionalLogic .ConditionalAttribute "childItem"}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "variable")
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", childItem.{{toGoName .TfName}}Variable.ValueString())
						}
					} else {{end}}{{if and .DefaultValuePresent (not .ExcludeNull)}}if childItem.{{toGoName .TfName}}.IsNull() {{if and .DynamicDefault .DefaultValuePresent}}|| data.{{toGoName .TfName}}.ValueString() == "{{.DefaultValue}}"{{end}} {
						if true{{buildConditionalLogic .ConditionalAttribute "childItem"}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
						{{if or .DefaultValue .DefaultValueEmptyString}}itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}}){{end}}
						}
					} else {{else if .AlwaysInclude }}if childItem.{{toGoName .TfName}}.IsNull() {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}optionType", "default")
					} else {{else if .AlwaysIncludeParent }}if data.{{toGoName .TfName}}.IsNull() {
						body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}optionType", "default")
					} else {{else}}if !childItem.{{toGoName .TfName}}.IsNull(){{end}} {
						if true{{buildConditionalLogic .ConditionalAttribute "childItem"}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "global")
						{{- if isListSet .}}
						var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
						childItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", values)
						{{- else}}
						{{- if eq .Type "StringInt64" }}
						if numValue, err := strconv.Atoi(childItem.{{toGoName .TfName}}.ValueString()); err != nil {
							itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", childItem.{{toGoName .TfName}}.ValueString())
						} else {
							itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", numValue)
						}
						{{- else}}
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", childItem.{{toGoName .TfName}}.Value{{.Type}}())
						{{- end}}
						{{- end}}
						}
					}
					{{- else if isNestedListSet .}}
					if true{{buildConditionalLogic .ConditionalAttribute "itemChildBody"}} {
						{{if or .AlwaysInclude (and (not .MinList) (not .ExcludeNull))}}itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{}){{end}}
						for _, childChildItem := range childItem.{{toGoName .TfName}} {
							itemChildChildBody := ""
							{{- range .Attributes}}
							{{- if .Value}}
							if true{{buildConditionalLogic .ConditionalAttribute "childChildItem"}} {
							itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", {{if .ValueType}}"{{.ValueType}}"{{else}}"default"{{end}})
							itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
							}
							{{- else if and (or (eq .Type "String") (eq .Type "StringInt64") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .TfOnly)}}
							{{if .Variable}}
							if !childChildItem.{{toGoName .TfName}}Variable.IsNull() {
								if true{{buildConditionalLogic .ConditionalAttribute "childChildItem"}} {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "variable")
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", childChildItem.{{toGoName .TfName}}Variable.ValueString())
								}
							} else {{end}}{{if and .DefaultValuePresent (not .ExcludeNull)}}if childChildItem.{{toGoName .TfName}}.IsNull() {{if and .DynamicDefault .DefaultValuePresent}}|| data.{{toGoName .TfName}}.ValueString() == "{{.DefaultValue}}"{{end}} {
								if true{{buildConditionalLogic .ConditionalAttribute "childChildItem"}} {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
								{{if or .DefaultValue .DefaultValueEmptyString}}itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}}){{end}}
							} else {{else if .AlwaysInclude }}if childChildItem.{{toGoName .TfName}}.IsNull() {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}optionType", "default")
							} else {{else if .AlwaysIncludeParent }}if data.{{toGoName .TfName}}.IsNull() {
								body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}optionType", "default")
							} else {{else}}if !childChildItem.{{toGoName .TfName}}.IsNull(){{end}} {
								if true{{buildConditionalLogic .ConditionalAttribute "childChildItem"}} {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "global")
								{{- if isListSet .}}
								var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
								childChildItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", values)
								{{- else}}
								{{- if eq .Type "StringInt64" }}
								if numValue, err := strconv.Atoi(childChildItem.{{toGoName .TfName}}.ValueString()); err != nil {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", childChildItem.{{toGoName .TfName}}.ValueString())
								} else {
									itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", numValue)
								}
								{{- else}}
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", childChildItem.{{toGoName .TfName}}.Value{{.Type}}())
								{{- end}}
								{{- end}}
								}
							}
							{{- end}}
							{{- end}}
							itemChildBody, _ = sjson.SetRaw(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildChildBody)
						}
					}
					{{- end}}
					{{- end}}
					itemBody, _ = sjson.SetRaw(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemChildBody)
				}
			}
			{{- end}}
			{{- end}}
			body, _ = sjson.SetRaw(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.-1", itemBody)
		}
	}
	{{- end}}
	{{- end}}
	return body
}
// End of section. //template:end toBody

// Section below is generated&owned by "gen/generator.go". //template:begin fromBody
func (data *{{camelCase .Name}}) fromBody(ctx context.Context, res gjson.Result, fullRead bool) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	{{- range .Attributes}}
	{{- $cname := toGoName .TfName}}
	{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "StringInt64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .Reference) (not .TfOnly) (not .WriteOnly) (not .Value)}}
	{{- if eq .Type "StringInt64"}}
	data.{{toGoName .TfName}} = types.StringNull()
	{{- else}}
	{{- if .DynamicDefault}}
	temp{{toGoName .TfName}} := data.{{toGoName .TfName}}
	{{- end}}
	data.{{toGoName .TfName}} = types.{{.Type}}Null({{if isListSet .}}types.{{.ElementType}}Type{{end}})
	{{- end}}
	{{ if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	{{- if .NoOptionType}}
	if va := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); va.Exists() {
		data.{{toGoName .TfName}} = types.{{.Type}}Value(va.{{getGjsonType .Type}}())
	}
	{{- else}}
	if t := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
		va := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
		{{if .Variable}}if t.String() == "variable" {
			data.{{toGoName .TfName}}Variable = types.StringValue(va.String())
		} else{{end}} if t.String() == "global" {{if .DynamicDefault}}|| (t.String() == "default" && (temp{{toGoName .TfName}}.ValueString() == "{{.DefaultValue}}" || (fullRead && temp{{toGoName .TfName}}.ValueString() == ""))){{end}} {
			{{- if eq .Type "StringInt64" }}
			data.{{toGoName .TfName}} = types.StringValue(va.String())
			{{- else}}
			data.{{toGoName .TfName}} = {{if isListSet .}}helpers.Get{{.ElementType}}{{.Type}}(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
			{{- end}}
		}
		{{- if hasTfOnlyConditional .ConditionalAttribute}}
		{{- range filterTfOnlyConditions .ConditionalAttribute.Conditions}}
		data.{{toGoName .Name}} = {{if eq .Type "Bool"}}types.BoolValue({{.Value}}){{else}}types.StringValue("{{.Value}}"){{end}}
		{{- end}}
		{{- end}}
	}
	{{- end}}
	{{- else if isNestedListSet .}}
	{{- $list := (toGoName .TfName)}}
	old{{$list}} := data.{{toGoName .TfName}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() && len(value.Array()) > 0 {
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			{{- range .Attributes}}
			{{- $ccname := toGoName .TfName}}
			{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "StringInt64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .Reference) (not .TfOnly) (not .WriteOnly) (not .Value)}}
			{{- if eq .Type "StringInt64"}}
			item.{{toGoName .TfName}} = types.StringNull()
			{{- else}}
			{{- if .DynamicDefault}}
			temp{{toGoName .TfName}} := item.{{toGoName .TfName}}
			{{- end}}
			item.{{toGoName .TfName}} = types.{{.Type}}Null({{if isListSet .}}types.{{.ElementType}}Type{{end}})
			{{- end}}
			{{ if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			if t := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
				va := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
				{{if .Variable}}if t.String() == "variable" {
					item.{{toGoName .TfName}}Variable = types.StringValue(va.String())
				} else{{end}} if t.String() == "global" {{if .DynamicDefault}}|| (t.String() == "default" && temp{{toGoName .TfName}}.ValueString() == "{{.DefaultValue}}"){{end}} {
					{{- if eq .Type "StringInt64" }}
					item.{{toGoName .TfName}} = types.StringValue(va.String())
					{{- else}}
					item.{{toGoName .TfName}} = {{if isListSet .}}helpers.Get{{.ElementType}}{{.Type}}(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
					{{- end}}
				}
				{{- if hasTfOnlyConditional .ConditionalAttribute}}
				{{- range filterTfOnlyConditions .ConditionalAttribute.Conditions}}
				item.{{toGoName .Name}} = {{if eq .Type "Bool"}}types.BoolValue({{.Value}}){{else}}types.StringValue("{{.Value}}"){{end}}
				{{- end}}
				{{- end}}
			}
			{{- else if isNestedListSet .}}
			{{- $clist := (toGoName .TfName)}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() && len(cValue.Array()) > 0 {
				item.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{toGoName .TfName}}, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := {{$name}}{{$cname}}{{toGoName .TfName}}{}
					{{- range .Attributes}}
					{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "StringInt64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .Reference) (not .TfOnly) (not .WriteOnly) (not .Value)}}
					{{- if eq .Type "StringInt64"}}
					cItem.{{toGoName .TfName}} = types.StringNull()
					{{- else}}
					{{- if .DynamicDefault}}
					temp{{toGoName .TfName}} := cItem.{{toGoName .TfName}}
					{{- end}}
					cItem.{{toGoName .TfName}} = types.{{.Type}}Null({{if isListSet .}}types.{{.ElementType}}Type{{end}})
					{{- end}}
					{{ if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
					if t := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
						va := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
						{{if .Variable}}if t.String() == "variable" {
							cItem.{{toGoName .TfName}}Variable = types.StringValue(va.String())
						} else{{end}} if t.String() == "global" {{if .DynamicDefault}}|| (t.String() == "default" && temp{{toGoName .TfName}}.ValueString() == "{{.DefaultValue}}"){{end}} {
							{{- if eq .Type "StringInt64" }}
							cItem.{{toGoName .TfName}} = types.StringValue(va.String())
							{{- else}}
							cItem.{{toGoName .TfName}} = {{if isListSet .}}helpers.Get{{.ElementType}}{{.Type}}(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
							{{- end}}
						}
						{{- if hasTfOnlyConditional .ConditionalAttribute}}
						{{- range filterTfOnlyConditions .ConditionalAttribute.Conditions}}
						cItem.{{toGoName .Name}} = {{if eq .Type "Bool"}}types.BoolValue({{.Value}}){{else}}types.StringValue("{{.Value}}"){{end}}
						{{- end}}
						{{- end}}
					}
					{{- else if isNestedListSet .}}
					{{- $cclist := (toGoName .TfName)}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() && len(ccValue.Array()) > 0{
						cItem.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := {{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}{}
							{{- range .Attributes}}
							{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "StringInt64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .Reference) (not .TfOnly) (not .WriteOnly) (not .Value)}}
							{{- if eq .Type "StringInt64"}}
							ccItem.{{toGoName .TfName}} = types.StringNull()
							{{- else}}
							{{- if .DynamicDefault}}
							temp{{toGoName .TfName}} := ccItem.{{toGoName .TfName}}
							{{- end}}
							ccItem.{{toGoName .TfName}} = types.{{.Type}}Null({{if isListSet .}}types.{{.ElementType}}Type{{end}})
							{{- end}}
							{{ if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
							if t := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
								va := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
								{{if .Variable}}if t.String() == "variable" {
									ccItem.{{toGoName .TfName}}Variable = types.StringValue(va.String())
								} else{{end}} if t.String() == "global" {{if .DynamicDefault}}|| (t.String() == "default" && temp{{toGoName .TfName}}.ValueString() == "{{.DefaultValue}}"){{end}} {
									{{- if eq .Type "StringInt64" }}
									ccItem.{{toGoName .TfName}} = types.StringValue(va.String())
									{{- else}}
									ccItem.{{toGoName .TfName}} = {{if isListSet .}}helpers.Get{{.ElementType}}{{.Type}}(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
									{{- end}}
								}
								{{- if hasTfOnlyConditional .ConditionalAttribute}}
								{{- range filterTfOnlyConditions .ConditionalAttribute.Conditions}}
								ccItem.{{toGoName .Name}} = {{if eq .Type "Bool"}}types.BoolValue({{.Value}}){{else}}types.StringValue("{{.Value}}"){{end}}
								{{- end}}
								{{- end}}
							}
							{{- end}}
							{{- end}}
							cItem.{{toGoName .TfName}} = append(cItem.{{toGoName .TfName}}, ccItem)
							return true
						})
						{{- if hasTfOnlyConditional .ConditionalAttribute}}
						{{- range filterTfOnlyConditions .ConditionalAttribute.Conditions}}
						cItem.{{toGoName .Name}} = {{if eq .Type "Bool"}}types.BoolValue({{.Value}}){{else}}types.StringValue("{{.Value}}"){{end}}
						{{- end}}
						{{- end}}
					}
					{{- end}}
					{{- end}}
					item.{{toGoName .TfName}} = append(item.{{toGoName .TfName}}, cItem)
					return true
				})
				{{- if hasTfOnlyConditional .ConditionalAttribute}}
				{{- range filterTfOnlyConditions .ConditionalAttribute.Conditions}}
				item.{{toGoName .Name}} = {{if eq .Type "Bool"}}types.BoolValue({{.Value}}){{else}}types.StringValue("{{.Value}}"){{end}}
				{{- end}}
				{{- end}}
			}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
		{{- if hasTfOnlyConditional .ConditionalAttribute}}
		{{- range filterTfOnlyConditions .ConditionalAttribute.Conditions}}
		data.{{toGoName .Name}} = {{if eq .Type "Bool"}}types.BoolValue({{.Value}}){{else}}types.StringValue("{{.Value}}"){{end}}
		{{- end}}
		{{- end}}
	} else {
		data.{{toGoName .TfName}} = nil
	}
	if !fullRead {
		{{- $noId := not (hasId .Attributes)}}
		result{{$list}} := make([]{{$name}}{{toGoName .TfName}}, 0, len(data.{{toGoName .TfName}}))
		matched{{$list}} := make([]bool, len(data.{{toGoName .TfName}}))
		for _, oldItem := range old{{$list}} {
			for ni := range data.{{toGoName .TfName}} {
				if matched{{$list}}[ni] {
					continue
				}
				keyMatch := true
				{{- range .Attributes}}
				{{- if or .Id $noId}}
				{{- if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String") (eq .Type "StringInt64")}}
				{{- if .Variable}}
				if keyMatch && (oldItem.{{toGoName .TfName}}Variable.ValueString() != "" || data.{{$list}}[ni].{{toGoName .TfName}}Variable.ValueString() != "") {
					if oldItem.{{toGoName .TfName}}Variable.ValueString() != data.{{$list}}[ni].{{toGoName .TfName}}Variable.ValueString() {
						keyMatch = false
					}
				} else if keyMatch {
				{{- else}}
				if keyMatch {
				{{- end}}
					{{- if eq .Type "Int64"}}
					if oldItem.{{toGoName .TfName}}.ValueInt64() != data.{{$list}}[ni].{{toGoName .TfName}}.ValueInt64() {
					{{- else if eq .Type "Bool"}}
					if oldItem.{{toGoName .TfName}}.ValueBool() != data.{{$list}}[ni].{{toGoName .TfName}}.ValueBool() {
					{{- else}}
					if oldItem.{{toGoName .TfName}}.ValueString() != data.{{$list}}[ni].{{toGoName .TfName}}.ValueString() {
					{{- end}}
						keyMatch = false
					}
				}
				{{- else if or (eq .Type "Set") (eq .Type "List")}}
				{{- if or (eq .ElementType "String") (eq .ElementType "Int64")}}
				if keyMatch {
					{{- if eq .Type "Set"}}
					if helpers.GetStringFromSet(oldItem.{{toGoName .TfName}}).ValueString() != helpers.GetStringFromSet(data.{{$list}}[ni].{{toGoName .TfName}}).ValueString() {
					{{- else}}
					if helpers.GetStringFromList(oldItem.{{toGoName .TfName}}).ValueString() != helpers.GetStringFromList(data.{{$list}}[ni].{{toGoName .TfName}}).ValueString() {
					{{- end}}
						keyMatch = false
					}
				}
				{{- end}}
				{{- end}}
				{{- end}}
				{{- end}}
				if keyMatch {
					matched{{$list}}[ni] = true
					{{- range .Attributes}}
					{{- if .Encrypted}}
					data.{{$list}}[ni].{{toGoName .TfName}} = oldItem.{{toGoName .TfName}}
					{{- if .Variable}}
					data.{{$list}}[ni].{{toGoName .TfName}}Variable = oldItem.{{toGoName .TfName}}Variable
					{{- end}}
					{{- end}}
					{{- if isNestedListSet .}}
					{{- $clist := (toGoName .TfName)}}
					{{- $ccname := toGoName .TfName}}
					{{- $cnoId := not (hasId .Attributes)}}
					{
						resultC := make([]{{$name}}{{$cname}}{{toGoName .TfName}}, 0, len(data.{{$list}}[ni].{{toGoName .TfName}}))
						matchedC := make([]bool, len(data.{{$list}}[ni].{{toGoName .TfName}}))
						for _, oldCItem := range oldItem.{{toGoName .TfName}} {
							for nci := range data.{{$list}}[ni].{{toGoName .TfName}} {
								if matchedC[nci] {
									continue
								}
								keyMatchC := true
								{{- range .Attributes}}
								{{- if or .Id $cnoId}}
								{{- if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String") (eq .Type "StringInt64")}}
								{{- if .Variable}}
								if keyMatchC && (oldCItem.{{toGoName .TfName}}Variable.ValueString() != "" || data.{{$list}}[ni].{{$clist}}[nci].{{toGoName .TfName}}Variable.ValueString() != "") {
									if oldCItem.{{toGoName .TfName}}Variable.ValueString() != data.{{$list}}[ni].{{$clist}}[nci].{{toGoName .TfName}}Variable.ValueString() {
										keyMatchC = false
									}
								} else if keyMatchC {
								{{- else}}
								if keyMatchC {
								{{- end}}
									{{- if eq .Type "Int64"}}
									if oldCItem.{{toGoName .TfName}}.ValueInt64() != data.{{$list}}[ni].{{$clist}}[nci].{{toGoName .TfName}}.ValueInt64() {
									{{- else if eq .Type "Bool"}}
									if oldCItem.{{toGoName .TfName}}.ValueBool() != data.{{$list}}[ni].{{$clist}}[nci].{{toGoName .TfName}}.ValueBool() {
									{{- else}}
									if oldCItem.{{toGoName .TfName}}.ValueString() != data.{{$list}}[ni].{{$clist}}[nci].{{toGoName .TfName}}.ValueString() {
									{{- end}}
										keyMatchC = false
									}
								}
								{{- else if or (eq .Type "Set") (eq .Type "List")}}
								{{- if or (eq .ElementType "String") (eq .ElementType "Int64")}}
								if keyMatchC {
									{{- if eq .Type "Set"}}
									if helpers.GetStringFromSet(oldCItem.{{toGoName .TfName}}).ValueString() != helpers.GetStringFromSet(data.{{$list}}[ni].{{$clist}}[nci].{{toGoName .TfName}}).ValueString() {
									{{- else}}
									if helpers.GetStringFromList(oldCItem.{{toGoName .TfName}}).ValueString() != helpers.GetStringFromList(data.{{$list}}[ni].{{$clist}}[nci].{{toGoName .TfName}}).ValueString() {
									{{- end}}
										keyMatchC = false
									}
								}
								{{- end}}
								{{- end}}
								{{- end}}
								{{- end}}
								if keyMatchC {
									matchedC[nci] = true
									{{- range .Attributes}}
									{{- if .Encrypted}}
									data.{{$list}}[ni].{{$clist}}[nci].{{toGoName .TfName}} = oldCItem.{{toGoName .TfName}}
									{{- if .Variable}}
									data.{{$list}}[ni].{{$clist}}[nci].{{toGoName .TfName}}Variable = oldCItem.{{toGoName .TfName}}Variable
									{{- end}}
									{{- end}}
									{{- if isNestedListSet .}}
									{{- $cclist := (toGoName .TfName)}}
									{{- $ccnoId := not (hasId .Attributes)}}
									{
										resultCC := make([]{{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}, 0, len(data.{{$list}}[ni].{{$clist}}[nci].{{toGoName .TfName}}))
										matchedCC := make([]bool, len(data.{{$list}}[ni].{{$clist}}[nci].{{toGoName .TfName}}))
										for _, oldCCItem := range oldCItem.{{toGoName .TfName}} {
											for ncci := range data.{{$list}}[ni].{{$clist}}[nci].{{toGoName .TfName}} {
												if matchedCC[ncci] {
													continue
												}
												keyMatchCC := true
												{{- range .Attributes}}
												{{- if or .Id $ccnoId}}
												{{- if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String") (eq .Type "StringInt64")}}
												{{- if .Variable}}
												if keyMatchCC && (oldCCItem.{{toGoName .TfName}}Variable.ValueString() != "" || data.{{$list}}[ni].{{$clist}}[nci].{{$cclist}}[ncci].{{toGoName .TfName}}Variable.ValueString() != "") {
													if oldCCItem.{{toGoName .TfName}}Variable.ValueString() != data.{{$list}}[ni].{{$clist}}[nci].{{$cclist}}[ncci].{{toGoName .TfName}}Variable.ValueString() {
														keyMatchCC = false
													}
												} else if keyMatchCC {
												{{- else}}
												if keyMatchCC {
												{{- end}}
													{{- if eq .Type "Int64"}}
													if oldCCItem.{{toGoName .TfName}}.ValueInt64() != data.{{$list}}[ni].{{$clist}}[nci].{{$cclist}}[ncci].{{toGoName .TfName}}.ValueInt64() {
													{{- else if eq .Type "Bool"}}
													if oldCCItem.{{toGoName .TfName}}.ValueBool() != data.{{$list}}[ni].{{$clist}}[nci].{{$cclist}}[ncci].{{toGoName .TfName}}.ValueBool() {
													{{- else}}
													if oldCCItem.{{toGoName .TfName}}.ValueString() != data.{{$list}}[ni].{{$clist}}[nci].{{$cclist}}[ncci].{{toGoName .TfName}}.ValueString() {
													{{- end}}
														keyMatchCC = false
													}
												}
												{{- else if or (eq .Type "Set") (eq .Type "List")}}
												{{- if or (eq .ElementType "String") (eq .ElementType "Int64")}}
												if keyMatchCC {
													{{- if eq .Type "Set"}}
													if helpers.GetStringFromSet(oldCCItem.{{toGoName .TfName}}).ValueString() != helpers.GetStringFromSet(data.{{$list}}[ni].{{$clist}}[nci].{{$cclist}}[ncci].{{toGoName .TfName}}).ValueString() {
													{{- else}}
													if helpers.GetStringFromList(oldCCItem.{{toGoName .TfName}}).ValueString() != helpers.GetStringFromList(data.{{$list}}[ni].{{$clist}}[nci].{{$cclist}}[ncci].{{toGoName .TfName}}).ValueString() {
													{{- end}}
														keyMatchCC = false
													}
												}
												{{- end}}
												{{- end}}
												{{- end}}
												{{- end}}
												if keyMatchCC {
													matchedCC[ncci] = true
													{{- range .Attributes}}
													{{- if .Encrypted}}
													data.{{$list}}[ni].{{$clist}}[nci].{{$cclist}}[ncci].{{toGoName .TfName}} = oldCCItem.{{toGoName .TfName}}
													{{- if .Variable}}
													data.{{$list}}[ni].{{$clist}}[nci].{{$cclist}}[ncci].{{toGoName .TfName}}Variable = oldCCItem.{{toGoName .TfName}}Variable
													{{- end}}
													{{- end}}
													{{- end}}
													resultCC = append(resultCC, data.{{$list}}[ni].{{$clist}}[nci].{{$cclist}}[ncci])
													break
												}
											}
										}
										for ncci := range data.{{$list}}[ni].{{$clist}}[nci].{{toGoName .TfName}} {
											if !matchedCC[ncci] {
												resultCC = append(resultCC, data.{{$list}}[ni].{{$clist}}[nci].{{$cclist}}[ncci])
											}
										}
										data.{{$list}}[ni].{{$clist}}[nci].{{toGoName .TfName}} = resultCC
									}
									{{- end}}
									{{- end}}
									resultC = append(resultC, data.{{$list}}[ni].{{$clist}}[nci])
									break
								}
							}
						}
						for nci := range data.{{$list}}[ni].{{toGoName .TfName}} {
							if !matchedC[nci] {
								resultC = append(resultC, data.{{$list}}[ni].{{$clist}}[nci])
							}
						}
						data.{{$list}}[ni].{{toGoName .TfName}} = resultC
					}
					{{- end}}
					{{- end}}
					result{{$list}} = append(result{{$list}}, data.{{$list}}[ni])
					break
				}
			}
		}
		for ni := range data.{{toGoName .TfName}} {
			if !matched{{$list}}[ni] {
				result{{$list}} = append(result{{$list}}, data.{{$list}}[ni])
			}
		}
		data.{{toGoName .TfName}} = result{{$list}}
	}
	{{- end}}
	{{- end}}
}
// End of section. //template:end fromBody

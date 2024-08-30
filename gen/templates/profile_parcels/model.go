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
	{{toGoName .TfName}} types.{{.Type}} `tfsdk:"{{.TfName}}"`
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
	path := "data."
	{{- range .Attributes}}
	{{- if .Value}}
	if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
	body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
	body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
	}
	{{- else if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .Reference)}}
	{{if .Variable}}
	if !data.{{toGoName .TfName}}Variable.IsNull() {
		if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "variable")
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", data.{{toGoName .TfName}}Variable.ValueString())
		}
	} else {{end}}{{if and .DefaultValuePresent (not .ExcludeNull)}}if data.{{toGoName .TfName}}.IsNull() {
		if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
		{{if or .DefaultValue .DefaultValueEmptyString}}body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}}){{end}}
		}
	} else {{else if .AlwaysIncludeParent }}if data.{{toGoName .TfName}}.IsNull() {
			body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}optionType", "default")
	} else {{else}}if !data.{{toGoName .TfName}}.IsNull(){{end}} {
		if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "global")
		{{- if isListSet .}}
		var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
		data.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", values)
		{{- else}}
		body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", data.{{toGoName .TfName}}.Value{{.Type}}())
		{{- end}}
		}
	}
	{{- else if isNestedListSet .}}
	if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && data.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
		{{if and (not .MinList) (not .ExcludeNull)}}body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{}){{end}}
		for _, item := range data.{{toGoName .TfName}} {
			itemBody := ""
			{{- range .Attributes}}
			{{- if .Value}}
			if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
			itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
			}
			{{- else if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)}}
			{{if .Variable}}
			if !item.{{toGoName .TfName}}Variable.IsNull() {
				if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "variable")
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", item.{{toGoName .TfName}}Variable.ValueString())
				}
			} else {{end}}{{if and .DefaultValuePresent (not .ExcludeNull)}}if item.{{toGoName .TfName}}.IsNull() {
				if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
				{{if or .DefaultValue .DefaultValueEmptyString}}itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}}){{end}}
				}
			} else {{else if .AlwaysIncludeParent }}if data.{{toGoName .TfName}}.IsNull() {
				body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}optionType", "default")
			} else {{else}}if !item.{{toGoName .TfName}}.IsNull(){{end}} {
				if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "global")
				{{- if isListSet .}}
				var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
				item.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", values)
				{{- else}}
				itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", item.{{toGoName .TfName}}.Value{{.Type}}())
				{{- end}}
				}
			}
			{{- else if isNestedListSet .}}
				if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && item.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
				{{if and (not .MinList) (not .ExcludeNull)}}itemBody, _ = sjson.Set(itemBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{}){{end}}
				for _, childItem := range item.{{toGoName .TfName}} {
					itemChildBody := ""
					{{- range .Attributes}}
					{{- if .Value}}
					if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
					itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
					itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
					}
					{{- else if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)}}
					{{if .Variable}}
					if !childItem.{{toGoName .TfName}}Variable.IsNull() {
						if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "variable")
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", childItem.{{toGoName .TfName}}Variable.ValueString())
						}
					} else {{end}}{{if and .DefaultValuePresent (not .ExcludeNull)}}if childItem.{{toGoName .TfName}}.IsNull() {
						if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
						{{if or .DefaultValue .DefaultValueEmptyString}}itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}}){{end}}
						}
					} else {{else if .AlwaysIncludeParent }}if data.{{toGoName .TfName}}.IsNull() {
						body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}optionType", "default")
					} else {{else}}if !childItem.{{toGoName .TfName}}.IsNull(){{end}} {
						if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && childItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "global")
						{{- if isListSet .}}
						var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
						childItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", values)
						{{- else}}
						itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", childItem.{{toGoName .TfName}}.Value{{.Type}}())
						{{- end}}
						}
					}
					{{- else if isNestedListSet .}}
					if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && itemChildBody.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && itemChildBody.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
						{{if and (not .MinList) (not .ExcludeNull)}}itemChildBody, _ = sjson.Set(itemChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", []interface{}{}){{end}}
						for _, childChildItem := range childItem.{{toGoName .TfName}} {
							itemChildChildBody := ""
							{{- range .Attributes}}
							{{- if .Value}}
							if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
							itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
							itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.Value}}{{if eq .Type "String"}}"{{end}})
							}
							{{- else if or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)}}
							{{if .Variable}}
							if !childChildItem.{{toGoName .TfName}}Variable.IsNull() {
								if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "variable")
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", childChildItem.{{toGoName .TfName}}Variable.ValueString())
								}
							} else {{end}}{{if and .DefaultValuePresent (not .ExcludeNull)}}if childChildItem.{{toGoName .TfName}}.IsNull() {
								if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "default")
								{{if or .DefaultValue .DefaultValueEmptyString}}itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", {{if eq .Type "String"}}"{{end}}{{.DefaultValue}}{{if eq .Type "String"}}"{{end}}){{end}}
							} else {{else if .AlwaysIncludeParent }}if data.{{toGoName .TfName}}.IsNull() {
								body, _ = sjson.Set(body, path+"{{range .DataPath}}{{.}}.{{end}}optionType", "default")
							} else {{else}}if !childChildItem.{{toGoName .TfName}}.IsNull(){{end}} {
								if true{{if ne .ConditionalAttribute.Name ""}} {{if eq .ConditionalAttribute.Type "Bool"}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueBool() == {{.ConditionalAttribute.Value}} {{else}} && childChildItem.{{toGoName .ConditionalAttribute.Name}}.ValueString() == "{{.ConditionalAttribute.Value}}" {{end}}{{end}} {
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType", "global")
								{{- if isListSet .}}
								var values []{{if isStringListSet .}}string{{else if isInt64ListSet .}}int64{{end}}
								childChildItem.{{toGoName .TfName}}.ElementsAs(ctx, &values, false)
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", values)
								{{- else}}
								itemChildChildBody, _ = sjson.Set(itemChildChildBody, "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value", childChildItem.{{toGoName .TfName}}.Value{{.Type}}())
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
func (data *{{camelCase .Name}}) fromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	{{- range .Attributes}}
	{{- $cname := toGoName .TfName}}
	{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .Reference) (not .WriteOnly) (not .Value)}}
	data.{{toGoName .TfName}} = types.{{.Type}}Null({{if isListSet .}}types.{{.ElementType}}Type{{end}})
	{{ if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	if t := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
		va := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
		{{if .Variable}}if t.String() == "variable" {
			data.{{toGoName .TfName}}Variable = types.StringValue(va.String())
		} else{{end}} if t.String() == "global" {
			data.{{toGoName .TfName}} = {{if isListSet .}}helpers.Get{{.ElementType}}{{.Type}}(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
		}
	}
	{{- else if isNestedListSet .}}
	if value := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); value.Exists() {
		data.{{toGoName .TfName}} = make([]{{$name}}{{toGoName .TfName}}, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := {{$name}}{{toGoName .TfName}}{}
			{{- range .Attributes}}
			{{- $ccname := toGoName .TfName}}
			{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .Reference) (not .WriteOnly) (not .Value)}}
			item.{{toGoName .TfName}} = types.{{.Type}}Null({{if isListSet .}}types.{{.ElementType}}Type{{end}})
			{{ if .Variable}}item.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			if t := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
				va := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
				{{if .Variable}}if t.String() == "variable" {
					item.{{toGoName .TfName}}Variable = types.StringValue(va.String())
				} else{{end}} if t.String() == "global" {
					item.{{toGoName .TfName}} = {{if isListSet .}}helpers.Get{{.ElementType}}{{.Type}}(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
				}
			}
			{{- else if isNestedListSet .}}
			if cValue := v.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); cValue.Exists() {
				item.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{toGoName .TfName}}, 0)
				cValue.ForEach(func(ck, cv gjson.Result) bool {
					cItem := {{$name}}{{$cname}}{{toGoName .TfName}}{}
					{{- range .Attributes}}
					{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .Reference) (not .WriteOnly) (not .Value)}}
					cItem.{{toGoName .TfName}} = types.{{.Type}}Null({{if isListSet .}}types.{{.ElementType}}Type{{end}})
					{{ if .Variable}}cItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
					if t := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
						va := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
						{{if .Variable}}if t.String() == "variable" {
							cItem.{{toGoName .TfName}}Variable = types.StringValue(va.String())
						} else{{end}} if t.String() == "global" {
							cItem.{{toGoName .TfName}} = {{if isListSet .}}helpers.Get{{.ElementType}}{{.Type}}(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
						}
					}
					{{- else if isNestedListSet .}}
					if ccValue := cv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}"); ccValue.Exists() {
						cItem.{{toGoName .TfName}} = make([]{{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}, 0)
						ccValue.ForEach(func(cck, ccv gjson.Result) bool {
							ccItem := {{$name}}{{$cname}}{{$ccname}}{{toGoName .TfName}}{}
							{{- range .Attributes}}
							{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .Reference) (not .WriteOnly) (not .Value)}}
							ccItem.{{toGoName .TfName}} = types.{{.Type}}Null({{if isListSet .}}types.{{.ElementType}}Type{{end}})
							{{ if .Variable}}ccItem.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
							if t := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
								va := ccv.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
								{{if .Variable}}if t.String() == "variable" {
									ccItem.{{toGoName .TfName}}Variable = types.StringValue(va.String())
								} else{{end}} if t.String() == "global" {
									ccItem.{{toGoName .TfName}} = {{if isListSet .}}helpers.Get{{.ElementType}}{{.Type}}(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
								}
							}
							{{- end}}
							{{- end}}
							cItem.{{toGoName .TfName}} = append(cItem.{{toGoName .TfName}}, ccItem)
							return true
						})
					}
					{{- end}}
					{{- end}}
					item.{{toGoName .TfName}} = append(item.{{toGoName .TfName}}, cItem)
					return true
				})
			}
			{{- end}}
			{{- end}}
			data.{{toGoName .TfName}} = append(data.{{toGoName .TfName}}, item)
			return true
		})
	}
	{{- end}}
	{{- end}}
}
// End of section. //template:end fromBody

// Section below is generated&owned by "gen/generator.go". //template:begin updateFromBody
func (data *{{camelCase .Name}}) updateFromBody(ctx context.Context, res gjson.Result) {
	data.Name = types.StringValue(res.Get("payload.name").String())
	if value := res.Get("payload.description"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	path := "payload.data."
	{{- range .Attributes}}
	{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .Reference) (not .WriteOnly) (not .Value)}}
	data.{{toGoName .TfName}} = types.{{.Type}}Null({{if isListSet .}}types.{{.ElementType}}Type{{end}})
	{{ if .Variable}}data.{{toGoName .TfName}}Variable = types.StringNull(){{end}}
	if t := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
		va := res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
		{{if .Variable}}if t.String() == "variable" {
			data.{{toGoName .TfName}}Variable = types.StringValue(va.String())
		} else{{end}} if t.String() == "global" {
			data.{{toGoName .TfName}} = {{if isListSet .}}helpers.Get{{.ElementType}}{{.Type}}(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
		}
	}
	{{- else if isNestedListSet .}}
	{{- $list := (toGoName .TfName)}}
	for i := range data.{{toGoName .TfName}} {
		keys := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{end}}{{end}}{{end}} }
		keyValues := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }
		keyValuesVariables := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if .Variable}}data.{{$list}}[i].{{toGoName .TfName}}Variable.ValueString(), {{else}}"", {{end}}{{end}}{{end}} }

		var r gjson.Result
		res.Get(path + "{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").ForEach(
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

		{{- range .Attributes}}
		{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .Reference) (not .WriteOnly) (not .Value)}}
		data.{{$list}}[i].{{toGoName .TfName}} = types.{{.Type}}Null({{if isListSet .}}types.{{.ElementType}}Type{{end}})
		{{ if .Variable}}data.{{$list}}[i].{{toGoName .TfName}}Variable = types.StringNull(){{end}}
		if t := r.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
			va := r.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
			{{if .Variable}}if t.String() == "variable" {
				data.{{$list}}[i].{{toGoName .TfName}}Variable = types.StringValue(va.String())
			} else{{end}} if t.String() == "global" {
				data.{{$list}}[i].{{toGoName .TfName}} = {{if isListSet .}}helpers.Get{{.ElementType}}{{.Type}}(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
			}
		}
		{{- else if isNestedListSet .}}
		{{- $clist := (toGoName .TfName)}}
		for ci := range data.{{$list}}[i].{{toGoName .TfName}} {
			keys := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{end}}{{end}}{{end}} }
			keyValues := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }
			keyValuesVariables := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if .Variable}}data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}Variable.ValueString(), {{else}}"", {{end}}{{end}}{{end}} }

			var cr gjson.Result
			r.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").ForEach(
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
						cr = v
						return false
					}
					return true
				},
			)

			{{- range .Attributes}}
			{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .Reference) (not .WriteOnly) (not .Value)}}
			data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = types.{{.Type}}Null({{if isListSet .}}types.{{.ElementType}}Type{{end}})
			{{ if .Variable}}data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}Variable = types.StringNull(){{end}}
			if t := cr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
				va := cr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
				{{if .Variable}}if t.String() == "variable" {
					data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}}Variable = types.StringValue(va.String())
				} else{{end}} if t.String() == "global" {
					data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} = {{if isListSet .}}helpers.Get{{.ElementType}}{{.Type}}(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
				}
			}
			{{- else if isNestedListSet .}}
			{{- $cclist := (toGoName .TfName)}}
			for cci := range data.{{$list}}[i].{{$clist}}[ci].{{toGoName .TfName}} {
				keys := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if or (eq .Type "Int64") (eq .Type "Bool") (eq .Type "String")}}"{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}", {{end}}{{end}}{{end}} }
				keyValues := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if eq .Type "Int64"}}strconv.FormatInt(data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.ValueInt64(), 10), {{else if eq .Type "Bool"}}strconv.FormatBool(data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.ValueBool()), {{else if eq .Type "String"}}data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}.Value{{.Type}}(), {{end}}{{end}}{{end}} }
				keyValuesVariables := [...]string{ {{$noId := not (hasId .Attributes)}}{{range .Attributes}}{{if or .Id $noId}}{{if .Variable}}data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}Variable.ValueString(), {{else}}"", {{end}}{{end}}{{end}} }

				var ccr gjson.Result
				cr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}").ForEach(
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
							ccr = v
							return false
						}
						return true
					},
				)

				{{- range .Attributes}}
				{{- if and (or (eq .Type "String") (eq .Type "Int64") (eq .Type "Float64") (eq .Type "Bool") (isListSet .)) (not .Reference) (not .WriteOnly) (not .Value)}}
				data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}} = types.{{.Type}}Null({{if isListSet .}}types.{{.ElementType}}Type{{end}})
				{{ if .Variable}}data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}Variable = types.StringNull(){{end}}
				if t := ccr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.optionType"); t.Exists() {
					va := ccr.Get("{{range .DataPath}}{{.}}.{{end}}{{.ModelName}}.value")
					{{if .Variable}}if t.String() == "variable" {
						data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}}Variable = types.StringValue(va.String())
					} else{{end}} if t.String() == "global" {
						data.{{$list}}[i].{{$clist}}[ci].{{$cclist}}[cci].{{toGoName .TfName}} = {{if isListSet .}}helpers.Get{{.ElementType}}{{.Type}}(va.Array()){{else}}types.{{.Type}}Value(va.{{getGjsonType .Type}}()){{end}}
					}
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
// End of section. //template:end updateFromBody

// Section below is generated&owned by "gen/generator.go". //template:begin isNull
func (data *{{camelCase .Name}}) isNull(ctx context.Context, res gjson.Result) bool {
	{{- range .Attributes}}
	{{- if not .Value}}
	{{- if isNestedListSet .}}
	if len(data.{{toGoName .TfName}}) > 0 {
		return false
	}
	{{- else}}
	if !data.{{toGoName .TfName}}.IsNull() {
		return false
	}
	{{- if .Variable}}
	if !data.{{toGoName .TfName}}Variable.IsNull() {
		return false
	}
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	return true
}
// End of section. //template:end isNull


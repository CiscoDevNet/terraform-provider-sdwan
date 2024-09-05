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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)
// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSource
func TestAccDataSourceSdwan{{camelCase .Name}}ProfileParcel(t *testing.T) {
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} && {{end}}os.Getenv("{{$e}}") == ""{{end}} {
		t.Skip("skipping test, set environment variable {{range $i, $e := .TestTags}}{{if $i}} or {{end}}{{$e}}{{end}}")
	}
	{{- end}}
	var checks []resource.TestCheckFunc
	{{- $config := . }}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Reference) (not .Value) (not .TestValue)}}
	{{- if isNestedListSet .}}
	{{- $list := .TfName }}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Reference) (not .Value) (not .TestValue)}}
	{{- if isNestedListSet .}}
	{{- $clist := .TfName }}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Reference) (not .Value) (not .TestValue)}}
	{{- if isNestedListSet .}}
	{{- $cclist := .TfName }}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	{{- range  .Attributes}}
	{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Reference) (not .Value) (not .TestValue) (not (isListSet .))}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_{{getProfileParcelName $config}}.test", "{{$list}}.0.{{$clist}}.0.{{$cclist}}.0.{{.TfName}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_{{getProfileParcelName $config}}.test", "{{$list}}.0.{{$clist}}.0.{{$cclist}}.0.{{.TfName}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if len .TestTags}}
	}
	{{- end}}
	{{- else if (not (isListSet .))}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_{{getProfileParcelName $config}}.test", "{{$list}}.0.{{$clist}}.0.{{.TfName}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_{{getProfileParcelName $config}}.test", "{{$list}}.0.{{$clist}}.0.{{.TfName}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if len .TestTags}}
	}
	{{- end}}
	{{- else if (not (isListSet .))}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_{{getProfileParcelName $config}}.test", "{{$list}}.0.{{.TfName}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_{{getProfileParcelName $config}}.test", "{{$list}}.0.{{.TfName}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	{{- if len .TestTags}}
	}
	{{- end}}
	{{- else if (not (isListSet .))}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_{{getProfileParcelName $config}}.test", "{{.TfName}}", "{{.Example}}"))
	}
	{{- else}}
	checks = append(checks, resource.TestCheckResourceAttr("data.sdwan_{{getProfileParcelName $config}}.test", "{{.TfName}}", "{{.Example}}"))
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: {{if .TestPrerequisites}}testAccDataSourceSdwan{{camelCase .Name}}PrerequisitesProfileParcelConfig+{{end}}testAccDataSourceSdwan{{camelCase .Name}}ProfileParcelConfig(),
				Check: resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}
// End of section. //template:end testAccDataSource

// Section below is generated&owned by "gen/generator.go". //template:begin testPrerequisites
{{- if .TestPrerequisites}}
const testAccDataSourceSdwan{{camelCase .Name}}PrerequisitesProfileParcelConfig = `
{{.TestPrerequisites}}
`
{{- end}}
// End of section. //template:end testPrerequisites

// Section below is generated&owned by "gen/generator.go". //template:begin testAccDataSourceConfig
func testAccDataSourceSdwan{{camelCase .Name}}ProfileParcelConfig() string {
	config := `resource "sdwan_{{getProfileParcelName $config}}" "test" {` + "\n"
	config += ` name = "TF_TEST"` + "\n"
	config += ` description = "Terraform integration test"` + "\n"
	{{- range  .Attributes}}
	{{- if and (not .ExcludeTest) (not .Value)}}
	{{- if isNestedListSet .}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
	{{- end}}
	config += `	{{.TfName}} = [{` + "\n"
		{{- range  .Attributes}}
		{{- if and (not .ExcludeTest) (not .Value)}}
		{{- if isNestedListSet .}}
		{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		{{- end}}
	config += `	  {{.TfName}} = [{` + "\n"
			{{- range  .Attributes}}
			{{- if and (not .ExcludeTest) (not .Value)}}
			{{- if isNestedListSet .}}
			{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
			{{- end}}
	config += `      {{.TfName}} = [{` + "\n"
				{{- range  .Attributes}}
				{{- if and (not .ExcludeTest) (not .Value)}}
				{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `			{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
		{{- if .SecondaryTestValue}}
		config += `		}, {` + "\n"
		config += `			{{.TfName}} = {{if .SecondaryTestValue}}{{.SecondaryTestValue}}{{end}}` + "\n"
		{{- end}}
	}
				{{- else}}
	config += `			{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	{{- if .SecondaryTestValue}}
	config += `		}, {` + "\n"
	config += `			{{.TfName}} = {{if .SecondaryTestValue}}{{.SecondaryTestValue}}{{end}}` + "\n"
	{{- end}}
				{{- end}}
				{{- end}}
				{{- end}}
	config += `		}]` + "\n"
			{{- if len .TestTags}}
	}
			{{- end}}
			{{- else}}
			{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `		{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
		{{- if .SecondaryTestValue}}
		config += `	}, {` + "\n"
		config += `			{{.TfName}} = {{if .SecondaryTestValue}}{{.SecondaryTestValue}}{{end}}` + "\n"
		{{- end}}
	}
			{{- else}}
	config += `		{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	{{- if .SecondaryTestValue}}
	config += `	}, {` + "\n"
	config += `		{{.TfName}} = {{if .SecondaryTestValue}}{{.SecondaryTestValue}}{{end}}` + "\n"
	{{- end}}
			{{- end}}
			{{- end}}
			{{- end}}
			{{- end}}
	config += `	}]` + "\n"
		{{- if len .TestTags}}
	}
		{{- end}}
		{{- else}}
		{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `	  {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
		{{- if .SecondaryTestValue}}
		config += `	}, {` + "\n"
		config += `	  {{.TfName}} = {{if .SecondaryTestValue}}{{.SecondaryTestValue}}{{end}}` + "\n"
		{{- end}}
	}
			{{- else}}
	config += `	  {{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	{{- if .SecondaryTestValue}}
	config += `	}, {` + "\n"
	config += `	  {{.TfName}} = {{if .SecondaryTestValue}}{{.SecondaryTestValue}}{{end}}` + "\n"
	{{- end}}
		{{- end}}
		{{- end}}
		{{- end}}
		{{- end}}
	config += `	}]` + "\n"
		{{- if len .TestTags}}
	}
		{{- end}}
	{{- else}}
	{{- if len .TestTags}}
	if {{range $i, $e := .TestTags}}{{if $i}} || {{end}}os.Getenv("{{$e}}") != ""{{end}} {
		config += `	{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	}
	{{- else}}
	config += `	{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}` + "\n"
	{{- end}}
	{{- end}}
	{{- end}}
	{{- end}}
	config += `}` + "\n"

	config += `
		data "sdwan_{{getProfileParcelName $config}}" "test" {
			id = sdwan_{{getProfileParcelName $config}}.test.id
			{{- range  .Attributes}}
			{{- if .Reference}}
			{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}
			{{- end}}
			{{- end}}
		}
	`
	return config
}
// End of section. //template:end testAccDataSourceConfig

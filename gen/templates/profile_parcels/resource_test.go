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

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwan{{camelCase .Name}}ProfileParcel(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwan{{camelCase .Name}}ProfileParcelConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					{{- $name := .Name }}
					{{- range  .Attributes}}
					{{- if and (.Mandatory) (not (isListSet .))}}
					resource.TestCheckResourceAttr("sdwan_{{snakeCase $name}}_profile_parcel.test", "{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
				),
			},
			{
				Config: testAccSdwan{{camelCase .Name}}ProfileParcelConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					{{- $name := .Name }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Reference)}}
					{{- if isNestedListSet .}}
					{{- $list := .TfName }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Reference)}}
					{{- if isNestedListSet .}}
					{{- $clist := .TfName }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Reference)}}
					{{- if isNestedListSet .}}
					{{- $cclist := .TfName }}
					{{- range  .Attributes}}
					{{- if and (not .WriteOnly) (not .ExcludeTest) (not .Reference) (not (isListSet .))}}
					resource.TestCheckResourceAttr("sdwan_{{snakeCase $name}}_profile_parcel.test", "{{$list}}.0.{{$clist}}.0.{{$cclist}}.0.{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- else if not (isListSet .)}}
					resource.TestCheckResourceAttr("sdwan_{{snakeCase $name}}_profile_parcel.test", "{{$list}}.0.{{$clist}}.0.{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
					{{- else if not (isListSet .)}}
					resource.TestCheckResourceAttr("sdwan_{{snakeCase $name}}_profile_parcel.test", "{{$list}}.0.{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
					{{- else if not (isListSet .)}}
					resource.TestCheckResourceAttr("sdwan_{{snakeCase $name}}_profile_parcel.test", "{{.TfName}}", "{{.Example}}"),
					{{- end}}
					{{- end}}
					{{- end}}
				),
			},
		},
	})
}

func testAccSdwan{{camelCase .Name}}ProfileParcelConfig_minimum() string {
	return `
	{{if .TestPrerequisites}}{{.TestPrerequisites}}{{end}}

	resource "sdwan_{{snakeCase $name}}_profile_parcel" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		{{- range .Attributes}}
		{{- if and (or (.ParcelMandatory) (.Reference)) (not (isNestedListSet .))}}
		{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}
		{{- end}}
		{{- end}}
	}
	`
}

func testAccSdwan{{camelCase .Name}}ProfileParcelConfig_all() string {
	return `
	{{if .TestPrerequisites}}{{.TestPrerequisites}}{{end}}

	resource "sdwan_{{snakeCase $name}}_profile_parcel" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		{{- range  .Attributes}}
		{{- if not .ExcludeTest}}
		{{- if isNestedListSet .}}
		{{.TfName}} = [{
			{{- range  .Attributes}}
			{{- if not .ExcludeTest}}
			{{- if isNestedListSet .}}
			{{.TfName}} = [{
				{{- range  .Attributes}}
				{{- if not .ExcludeTest}}
				{{- if isNestedListSet .}}
				{{.TfName}} = [{
					{{- range  .Attributes}}
					{{- if not .ExcludeTest}}
					{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}
					{{- end}}
					{{- end}}
				}]
				{{- else}}
				{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}
				{{- end}}
				{{- end}}
				{{- end}}
			}]
			{{- else}}
			{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}
			{{- end}}
			{{- end}}
			{{- end}}
		}]
		{{- else}}
		{{.TfName}} = {{if .TestValue}}{{.TestValue}}{{else}}{{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}{{end}}
		{{- end}}
		{{- end}}
		{{- end}}
	}
	`
}

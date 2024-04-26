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

//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v3"
)

const (
	featureTemplateDefinitionsPath = "./gen/definitions/feature_templates/"
	featureTemplateModelsPath      = "./gen/models/feature_templates/"
	profileParcelDefinitionsPath   = "./gen/definitions/profile_parcels/"
	profileParcelModelsPath        = "./gen/models/profile_parcels/"
	genericDefinitionsPath         = "./gen/definitions/generic/"
	providerTemplate               = "./gen/templates/provider.go"
	providerLocation               = "./internal/provider/provider.go"
	changelogTemplate              = "./gen/templates/changelog.md.tmpl"
	changelogLocation              = "./templates/guides/changelog.md.tmpl"
	changelogOriginal              = "./CHANGELOG.md"
)

type t struct {
	path   string
	prefix string
	suffix string
}

var featureTemplateTemplates = []t{
	{
		path:   "./gen/templates/feature_templates/model.go",
		prefix: "./internal/provider/model_sdwan_",
		suffix: "_feature_template.go",
	},
	{
		path:   "./gen/templates/feature_templates/data_source.go",
		prefix: "./internal/provider/data_source_sdwan_",
		suffix: "_feature_template.go",
	},
	{
		path:   "./gen/templates/feature_templates/data_source_test.go",
		prefix: "./internal/provider/data_source_sdwan_",
		suffix: "_feature_template_test.go",
	},
	{
		path:   "./gen/templates/feature_templates/resource.go",
		prefix: "./internal/provider/resource_sdwan_",
		suffix: "_feature_template.go",
	},
	{
		path:   "./gen/templates/feature_templates/resource_test.go",
		prefix: "./internal/provider/resource_sdwan_",
		suffix: "_feature_template_test.go",
	},
	{
		path:   "./gen/templates/feature_templates/data-source.tf",
		prefix: "./examples/data-sources/sdwan_",
		suffix: "_feature_template/data-source.tf",
	},
	{
		path:   "./gen/templates/feature_templates/resource.tf",
		prefix: "./examples/resources/sdwan_",
		suffix: "_feature_template/resource.tf",
	},
	{
		path:   "./gen/templates/feature_templates/import.sh",
		prefix: "./examples/resources/sdwan_",
		suffix: "_feature_template/import.sh",
	},
}

var profileParcelTemplates = []t{
	{
		path:   "./gen/templates/profile_parcels/model.go",
		prefix: "./internal/provider/model_sdwan_",
		suffix: "_profile_parcel.go",
	},
	{
		path:   "./gen/templates/profile_parcels/data_source.go",
		prefix: "./internal/provider/data_source_sdwan_",
		suffix: "_profile_parcel.go",
	},
	{
		path:   "./gen/templates/profile_parcels/data_source_test.go",
		prefix: "./internal/provider/data_source_sdwan_",
		suffix: "_profile_parcel_test.go",
	},
	{
		path:   "./gen/templates/profile_parcels/resource.go",
		prefix: "./internal/provider/resource_sdwan_",
		suffix: "_profile_parcel.go",
	},
	{
		path:   "./gen/templates/profile_parcels/resource_test.go",
		prefix: "./internal/provider/resource_sdwan_",
		suffix: "_profile_parcel_test.go",
	},
	{
		path:   "./gen/templates/profile_parcels/data-source.tf",
		prefix: "./examples/data-sources/sdwan_",
		suffix: "_profile_parcel/data-source.tf",
	},
	{
		path:   "./gen/templates/profile_parcels/resource.tf",
		prefix: "./examples/resources/sdwan_",
		suffix: "_profile_parcel/resource.tf",
	},
	{
		path:   "./gen/templates/profile_parcels/import.sh",
		prefix: "./examples/resources/sdwan_",
		suffix: "_profile_parcel/import.sh",
	},
}

var genericTemplates = []t{
	{
		path:   "./gen/templates/generic/model.go",
		prefix: "./internal/provider/model_sdwan_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/generic/data_source.go",
		prefix: "./internal/provider/data_source_sdwan_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/generic/data_source_test.go",
		prefix: "./internal/provider/data_source_sdwan_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/generic/resource.go",
		prefix: "./internal/provider/resource_sdwan_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/generic/resource_test.go",
		prefix: "./internal/provider/resource_sdwan_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/generic/data-source.tf",
		prefix: "./examples/data-sources/sdwan_",
		suffix: "/data-source.tf",
	},
	{
		path:   "./gen/templates/generic/resource.tf",
		prefix: "./examples/resources/sdwan_",
		suffix: "/resource.tf",
	},
	{
		path:   "./gen/templates/generic/import.sh",
		prefix: "./examples/resources/sdwan_",
		suffix: "/import.sh",
	},
}

type YamlConfig struct {
	Name                     string                `yaml:"name"`
	Model                    string                `yaml:"model"`
	RestEndpoint             string                `yaml:"rest_endpoint"`
	GetRestEndpoint          string                `yaml:"get_rest_endpoint"`
	PostRestEndpoint         string                `yaml:"post_rest_endpoint"`
	MinimumVersion           string                `yaml:"minimum_version"`
	DsDescription            string                `yaml:"ds_description"`
	ResDescription           string                `yaml:"res_description"`
	DocCategory              string                `yaml:"doc_category"`
	RootElement              string                `yaml:"root_element"`
	HasVersion               bool                  `yaml:"has_version"`
	IdAttribute              string                `yaml:"id_attribute"`
	IdFromQueryPath          string                `yaml:"id_from_query_path"`
	IdFromQueryPathAttribute string                `yaml:"id_from_query_path_attribute"`
	FeatureTemplateIsGlobal  bool                  `yaml:"feature_template_is_global"`
	SkipTemplates            []string              `yaml:"skip_templates"`
	Attributes               []YamlConfigAttribute `yaml:"attributes"`
	TestTags                 []string              `yaml:"test_tags"`
	TestPrerequisites        string                `yaml:"test_prerequisites"`
	RemoveId                 bool                  `yaml:"remove_id"`
}

type YamlConfigAttribute struct {
	ModelName             string                         `yaml:"model_name"`
	ResponseModelName     string                         `yaml:"response_model_name"`
	TfName                string                         `yaml:"tf_name"`
	Type                  string                         `yaml:"type"`
	ElementType           string                         `yaml:"element_type"`
	ObjectType            string                         `yaml:"object_type"`
	ModelTypeString       bool                           `yaml:"model_type_string"`
	BoolEmptyString       bool                           `yaml:"bool_empty_string"`
	DataPath              []string                       `yaml:"data_path"`
	Keys                  []string                       `yaml:"keys"`
	Id                    bool                           `yaml:"id"`
	Reference             bool                           `yaml:"reference"`
	Variable              bool                           `yaml:"variable"`
	Mandatory             bool                           `yaml:"mandatory"`
	ParcelMandatory       bool                           `yaml:"parcel_mandatory"`
	WriteOnly             bool                           `yaml:"write_only"`
	TfOnly                bool                           `yaml:"tf_only"`
	ExcludeTest           bool                           `yaml:"exclude_test"`
	ExcludeExample        bool                           `yaml:"exclude_example"`
	ExcludeIgnore         bool                           `yaml:"exclude_ignore"`
	ExcludeNull           bool                           `yaml:"exclude_null"`
	NodeOnlyContainer     bool                           `yaml:"node_only_container"`
	Description           string                         `yaml:"description"`
	Example               string                         `yaml:"example"`
	EnumValues            []string                       `yaml:"enum_values"`
	IgnoreEnum            bool                           `yaml:"ignore_enum"`
	MinList               int64                          `yaml:"min_list"`
	MaxList               int64                          `yaml:"max_list"`
	MinInt                int64                          `yaml:"min_int"`
	MaxInt                int64                          `yaml:"max_int"`
	MinFloat              float64                        `yaml:"min_float"`
	MaxFloat              float64                        `yaml:"max_float"`
	StringPatterns        []string                       `yaml:"string_patterns"`
	StringMinLength       int64                          `yaml:"string_min_length"`
	StringMaxLength       int64                          `yaml:"string_max_length"`
	DefaultValue          string                         `yaml:"default_value"`
	DefaultValuePresent   bool                           `yaml:"default_value_present"`
	Value                 string                         `yaml:"value"`
	TestValue             string                         `yaml:"test_value"`
	AlwaysInclude         bool                           `yaml:"always_include"`
	Attributes            []YamlConfigAttribute          `yaml:"attributes"`
	ConditionalAttribute  YamlConfigConditionalAttribute `yaml:"conditional_attribute"`
	ConditionalListLength string                         `yaml:"conditional_list_length"`
	QueryParam            bool                           `yaml:"query_param"`
	NoAugmentConfig       bool                           `yaml:"no_augment_config"`
	TestTags              []string                       `yaml:"test_tags"`
}

type YamlConfigConditionalAttribute struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

// Templating helper function to convert TF name to GO name
func ToGoName(s string) string {
	var g []string

	p := strings.Split(s, "_")

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	s = strings.Join(g, "")
	return s
}

// Templating helper function to convert string to camel case
func CamelCase(s string) string {
	var g []string

	s = strings.ReplaceAll(s, "-", " ")
	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	return strings.Join(g, "")
}

// Templating helper function to convert string to snake case
func SnakeCase(s string) string {
	var g []string

	if !strings.Contains(s, " ") && !strings.Contains(s, "-") {
		var words []string
		l := 0
		for w := s; w != ""; w = w[l:] {
			l = strings.IndexFunc(w[1:], unicode.IsUpper) + 1
			if l <= 0 {
				l = len(w)
			}
			words = append(words, strings.ToLower(w[:l]))
		}
		return strings.Join(words, "_")
	}

	s = strings.ReplaceAll(s, "-", " ")
	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

// Templating helper function to build a SJSON path
func BuildPath(s []string) string {
	return strings.Join(s, ".")
}

// Templating helper function to determine if attributes list contains one or more version attributes
func HasVersionAttribute(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Type == "Version" || attr.Type == "Versions" {
			return true
		} else if attr.Type == "List" || attr.Type == "Set" {
			if HasVersionAttribute(attr.Attributes) {
				return true
			}
		}
	}
	return false
}

// Templating helper function to return ResponseModelName if set, otherwise ModelName
func GetResponseModelName(attribute YamlConfigAttribute) string {
	if attribute.ResponseModelName != "" {
		return attribute.ResponseModelName
	} else {
		return attribute.ModelName
	}
}

// Templating helper function to return true if reference included in attributes
func HasReference(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Reference {
			return true
		}
	}
	return false
}

// Templating helper function to return GJSON type
func GetGjsonType(t string) string {
	if t == "String" {
		return "String"
	} else if t == "Int64" {
		return "Int"
	} else if t == "Float64" {
		return "Float"
	} else if t == "Bool" {
		return "Bool"
	} else {
		return ""
	}
}

// Templating helper function to return the ID attribute
func GetId(attributes []YamlConfigAttribute) YamlConfigAttribute {
	for _, attr := range attributes {
		if attr.Id {
			return attr
		}
	}
	return YamlConfigAttribute{}
}

// Templating helper function to return true if type is a list or set without nested elements
func IsListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType != "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list without nested elements
func IsList(attribute YamlConfigAttribute) bool {
	if attribute.Type == "List" && attribute.ElementType != "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a set without nested elements
func IsSet(attribute YamlConfigAttribute) bool {
	if attribute.Type == "Set" && attribute.ElementType != "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set of strings without nested elements
func IsStringListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "String" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set of integers without nested elements
func IsInt64ListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "Int64" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set with nested elements
func IsNestedListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list with nested elements
func IsNestedList(attribute YamlConfigAttribute) bool {
	if attribute.Type == "List" && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a set with nested elements
func IsNestedSet(attribute YamlConfigAttribute) bool {
	if attribute.Type == "Set" && attribute.ElementType == "" {
		return true
	}
	return false
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":             ToGoName,
	"camelCase":            CamelCase,
	"snakeCase":            SnakeCase,
	"sprintf":              fmt.Sprintf,
	"toLower":              strings.ToLower,
	"path":                 BuildPath,
	"hasVersionAttribute":  HasVersionAttribute,
	"getResponseModelName": GetResponseModelName,
	"hasReference":         HasReference,
	"getGjsonType":         GetGjsonType,
	"getId":                GetId,
	"isListSet":            IsListSet,
	"isList":               IsList,
	"isSet":                IsSet,
	"isStringListSet":      IsStringListSet,
	"isInt64ListSet":       IsInt64ListSet,
	"isNestedListSet":      IsNestedListSet,
	"isNestedList":         IsNestedList,
	"isNestedSet":          IsNestedSet,
	"contains":             contains,
}

func parseFeatureTemplateAttribute(attr *YamlConfigAttribute, model gjson.Result) {
	if attr.NoAugmentConfig {
		return
	}
	var r gjson.Result
	if model.Get("fields").Exists() {
		r = model.Get("fields.#(key==\"" + attr.ModelName + "\")#")
	} else {
		r = model.Get("children.#(key==\"" + attr.ModelName + "\")#")
	}
	if len(r.Array()) > 1 {
		found := false
		for _, item := range r.Array() {
			if len(item.Get("dataPath").Array()) == 0 {
				fmt.Printf("WARNING: Non-unique key without data_path: %s\n", attr.ModelName)
			}
			match := true
			for _, dp := range item.Get("dataPath").Array() {
				if !contains(attr.DataPath, dp.String()) {
					match = false
					break
				}
			}
			if match {
				r = item
				found = true
				break
			}
		}
		if !found {
			panic(fmt.Sprintf("Could not find element in schema: %v\n%v\n\n", attr.ModelName, model))
		}
	} else if len(r.Array()) == 1 {
		r = r.Array()[0]
	} else {
		panic(fmt.Sprintf("Could not find element in schema: %v\n%v\n\n", attr.ModelName, model))
	}

	if attr.Description == "" && r.Get("details").String() != "" {
		attr.Description = r.Get("details").String()
	}
	if attr.TfName == "" {
		attr.TfName = SnakeCase(attr.ModelName)
	}
	if len(attr.DataPath) == 0 && len(r.Get("dataPath").Array()) > 0 {
		paths := make([]string, len(r.Get("dataPath").Array()))
		for i, v := range r.Get("dataPath").Array() {
			paths[i] = v.String()
		}
		attr.DataPath = paths
	}
	if len(attr.Keys) == 0 && len(r.Get("primaryKeys").Array()) > 0 {
		keys := make([]string, len(r.Get("primaryKeys").Array()))
		for i, v := range r.Get("primaryKeys").Array() {
			keys[i] = v.String()
		}
		attr.Keys = keys
	}
	attr.ObjectType = r.Get("objectType").String()
	if r.Get("objectType").String() == "object" || r.Get("objectType").String() == "node-only" {
		t := r.Get("dataType.type").String()
		if r.Get("dataType").String() == "string" {
			t = "string"
		}
		if attr.Type != "" {
		} else if contains([]string{"string", "stringList", "passphrase", "restrictedPassphrase", "datetimelocal", "ip", "ipv4", "ipv6", "ipv4-prefix", "ipv6-prefix", "dnsHostName", "interfaceList", "tlocExtension", "xConnect", "mac", "remoteAS", "ike"}, t) {
			attr.Type = "String"
			if t != "passphrase" && t != "restrictedPassphrase" {
				if r.Get("dataType.minLength").Exists() && attr.StringMinLength == 0 {
					attr.StringMinLength = r.Get("dataType.minLength").Int()
				}
				if r.Get("dataType.maxLength").Exists() && attr.StringMaxLength == 0 {
					attr.StringMaxLength = r.Get("dataType.maxLength").Int()
				}
			}
		} else if t == "enum" {
			attr.Type = "String"
			for _, v := range r.Get("dataType.values").Array() {
				attr.EnumValues = append(attr.EnumValues, v.Get("key").String())
			}
		} else if t == "enumList" {
			attr.Type = "Set"
			attr.ElementType = "String"
		} else if t == "radioButtonList" {
			attr.Type = "String"
			for _, v := range r.Get("dataType.values").Array() {
				attr.EnumValues = append(attr.EnumValues, v.Get("value").String())
			}
		} else if contains([]string{"number", "numberFixedInterval", "holdTime", "bandwidth"}, t) {
			attr.Type = "Int64"
			if r.Get("dataType.min").Exists() && attr.MinInt == 0 {
				attr.MinInt = r.Get("dataType.min").Int()
			}
			if r.Get("dataType.max").Exists() && attr.MaxInt == 0 {
				attr.MaxInt = r.Get("dataType.max").Int()
			}
		} else if t == "float" {
			attr.Type = "Float64"
			if r.Get("dataType.min").Exists() && attr.MinFloat == 0 {
				attr.MinFloat = r.Get("dataType.min").Float()
			}
			if r.Get("dataType.max").Exists() && attr.MaxFloat == 0 {
				attr.MaxFloat = r.Get("dataType.max").Float()
			}
		} else if t == "boolean" {
			attr.Type = "Bool"
		}
	} else if r.Get("objectType").String() == "list" {
		attr.Type = "Set"
		if r.Get("dataType.type").String() == "number" {
			attr.ElementType = "Int64"

		} else {
			attr.ElementType = "String"
		}
	}
	if r.Get("dataType.default").Exists() {
		attr.DefaultValue = r.Get("dataType.default").String()
	}
	types := r.Get("optionType").Array()
	ignore := false
	variable := false
	for t := range types {
		if types[t].String() == "ignore" {
			ignore = true
		} else if types[t].String() == "variable" {
			variable = true
		}
	}
	if !ignore && r.Get("objectType").String() != "tree" && !attr.ExcludeIgnore {
		attr.Mandatory = true
	}
	if variable && r.Get("objectType").String() != "tree" && !attr.NodeOnlyContainer {
		attr.Variable = true
	}
	if r.Get("objectType").String() == "tree" && len(attr.Attributes) > 0 {
		attr.Type = "List"
		for a := range attr.Attributes {
			parseFeatureTemplateAttribute(&attr.Attributes[a], r)
		}
	}
}

func augmentFeatureTemplateConfig(config *YamlConfig) {
	if config.Model == "" {
		config.Model = SnakeCase(config.Name)
	}
	modelPath := featureTemplateModelsPath + config.Model + ".json"

	modelBytes, err := os.ReadFile(modelPath)
	if err != nil {
		log.Fatalf("Error reading model '%s': %v", modelPath, err)
	}

	model := gjson.ParseBytes(modelBytes)

	for ia := range config.Attributes {
		parseFeatureTemplateAttribute(&config.Attributes[ia], model)
	}

	if config.DsDescription == "" {
		config.DsDescription = fmt.Sprintf("This data source can read the %s feature template.", config.Name)
	}
	if config.ResDescription == "" {
		config.ResDescription = fmt.Sprintf("This resource can manage a %s feature template.", config.Name)
	}
}

func parseProfileParcelAttribute(attr *YamlConfigAttribute, model gjson.Result) {
	if attr.ModelName == "" {
		return
	}
	path := ""
	isOneOfAttribute := false
	for i, e := range attr.DataPath {
		// Check if the next element is a oneOf
		if model.Get("properties." + path + e + ".oneOf").Exists() {
			// We need to find the right element in oneOf which has the next element

			next := ""
			if i+1 < len(attr.DataPath) {
				next = attr.DataPath[i+1]
			} else {
				next = attr.ModelName
			}

			index := 0
			model.Get("properties." + path + e + ".oneOf").ForEach(func(k, v gjson.Result) bool {
				if v.Get("properties." + next).Exists() {
					path += fmt.Sprintf("%s.oneOf.%v.properties.", e, index)
					isOneOfAttribute = true
					return false // stop iterating
				}
				index += 1
				return true // keep iterating
			})

		} else if model.Get("properties." + path + e + ".items").Exists() {
			path += fmt.Sprintf("%s.items.properties.", e)

		} else {
			path += e + ".properties."
		}
	}
	path += attr.ModelName

	r := model.Get("properties." + path)

	if !r.Exists() {
		panic(fmt.Sprintf("Could not find element in schema: %v\n%v\n\n", attr.ModelName, model))
	}

	attr.Description = r.Get("description").String()
	if attr.TfName == "" {
		attr.TfName = SnakeCase(attr.ModelName)
	}

	if r.Get("type").String() == "object" && !r.Get("items").Exists() || !r.Get("type").Exists() {
		t := r.Get("oneOf.#(properties.optionType.enum.0=\"global\")")
		if value := r.Get("properties.optionType.enum.0"); value.String() == "global" {
			t = r
		}

		if t.Exists() {
			if t.Get("properties.value.type").String() == "string" || t.Get("properties.value.anyOf.0.type").String() == "string" || t.Get("properties.value.oneOf.0.type").String() == "string" {
				attr.Type = "String"
				if value := t.Get("properties.value.minLength"); value.Exists() {
					attr.StringMinLength = value.Int()
				}
				if value := t.Get("properties.value.maxLength"); value.Exists() {
					attr.StringMaxLength = value.Int()
				}
				if value := t.Get("properties.value.pattern"); value.Exists() {
					attr.StringPatterns = append(attr.StringPatterns, value.String())
				}
				if value := t.Get("properties.value.enum"); value.Exists() {
					for _, v := range value.Array() {
						attr.EnumValues = append(attr.EnumValues, v.String())
					}
				}
			} else if t.Get("properties.value.type").String() == "boolean" {
				attr.Type = "Bool"
			} else if t.Get("properties.value.type").String() == "integer" || t.Get("properties.value.type").String() == "number" || t.Get("properties.value.oneOf.0.type").String() == "integer" || t.Get("properties.value.oneOf.0.type").String() == "number" {
				attr.Type = "Int64"
				if value := t.Get("properties.value.minimum"); value.Exists() {
					attr.MinInt = value.Int()
				}
				if value := t.Get("properties.value.maximum"); value.Exists() {
					attr.MaxInt = value.Int()
				}
			} else if t.Get("properties.value.type").String() == "array" && t.Get("properties.value.items.type").String() == "string" {
				attr.Type = "Set"
				attr.ElementType = "String"
				// if value := t.Get("properties.value.items.minItems"); value.Exists() {
				//  attr.MinList = value.Int()
				// }
				// if value := t.Get("properties.value.items.maxItems"); value.Exists() {
				//  attr.MaxList = value.Int()
				// }
			} else if t.Get("properties.value.type").String() == "array" && t.Get("properties.value.items.type").String() == "integer" {
				attr.Type = "Set"
				attr.ElementType = "Int64"
				// if value := t.Get("properties.value.items.minimum"); value.Exists() {
				//  attr.MinInt = value.Int()
				// }
				// if value := t.Get("properties.value.items.maximum"); value.Exists() {
				//  attr.MaxInt = value.Int()
				// }
			} else if t.Get("properties.value.const").String() == "off" || t.Get("properties.value.const").String() == "on" {
				attr.Type = "String"
			} else {
				fmt.Printf("WARNING: Unsupported type: %s\n", t.Get("properties.value.type").String())
			}
		}
		if r.Get("oneOf.#(properties.optionType.enum.0=\"variable\")").Exists() && !isOneOfAttribute {
			attr.Variable = true
		}
		d := r.Get("oneOf.#(properties.optionType.enum.0=\"default\")")
		if value := r.Get("properties.optionType.enum.0=\"default\""); value.Exists() {
			d = r
		}
		if d.Exists() && !isOneOfAttribute {
			if value := d.Get("properties.value.enum.0"); value.Exists() {
				attr.DefaultValue = value.String()
				attr.DefaultValuePresent = true
			} else if value := d.Get("properties.value.default"); value.Exists() {
				attr.DefaultValue = value.String()
				attr.DefaultValuePresent = true
			} else if value := d.Get("properties.value.default"); value.Exists() {
				attr.DefaultValue = value.String()
				attr.DefaultValuePresent = true
			} else if value := d.Get("properties.value.minimum"); value.Exists() {
				attr.DefaultValue = value.String()
				attr.DefaultValuePresent = true
			} else {
				attr.ParcelMandatory = true
				if !attr.Variable {
					attr.Mandatory = true
				}
			}
		} else if isOneOfAttribute {
			attr.ParcelMandatory = true
			attr.ExcludeNull = true
		} else {
			attr.ParcelMandatory = true
			if !attr.Variable {
				attr.Mandatory = true
			}
		}
	} else if r.Get("type").String() == "array" && r.Get("items.type").String() == "object" && len(attr.Attributes) > 0 {
		attr.Type = "List"
		for a := range attr.Attributes {
			parseProfileParcelAttribute(&attr.Attributes[a], r.Get("items"))
		}
	}
}

func augmentProfileParcelConfig(config *YamlConfig) {
	if config.Model == "" {
		config.Model = SnakeCase(config.Name)
	}
	modelPath := profileParcelModelsPath + config.Model + ".json"

	modelBytes, err := os.ReadFile(modelPath)
	if err != nil {
		log.Fatalf("Error reading model '%s': %v", modelPath, err)
	}

	model := gjson.ParseBytes(modelBytes)

	for ia := range config.Attributes {
		parseProfileParcelAttribute(&config.Attributes[ia], model.Get("request.properties.data"))
	}

	if config.DsDescription == "" {
		config.DsDescription = fmt.Sprintf("This data source can read the %s profile parcel.", config.Name)
	}
	if config.ResDescription == "" {
		config.ResDescription = fmt.Sprintf("This resource can manage a %s profile parcel.", config.Name)
	}
}

func augmentGenericAttribute(attr *YamlConfigAttribute) {
	if attr.TfName == "" {
		attr.TfName = SnakeCase(attr.ModelName)
	}
	if (attr.Type == "List" || attr.Type == "Set") && attr.ElementType == "" {
		for a := range attr.Attributes {
			augmentGenericAttribute(&attr.Attributes[a])
		}
	}
}

func augmentGenericConfig(config *YamlConfig, type_ string) {
	for ia := range config.Attributes {
		augmentGenericAttribute(&config.Attributes[ia])
	}
	if config.DsDescription == "" {
		config.DsDescription = fmt.Sprintf("This data source can read the %s %s.", config.Name, type_)
	}
	if config.ResDescription == "" {
		config.ResDescription = fmt.Sprintf("This resource can manage a %s %s.", config.Name, type_)
	}
}

func renderTemplate(templatePath, outputPath string, config interface{}) {
	if c, ok := config.(YamlConfig); ok {
		for _, s := range c.SkipTemplates {
			if strings.Contains(templatePath, s) {
				return
			}
		}
	}

	file, err := os.Open(templatePath)
	if err != nil {
		log.Fatalf("Error opening template: %v", err)
	}
	defer file.Close()

	// skip first line with 'build-ignore' directive for go files
	scanner := bufio.NewScanner(file)
	if strings.HasSuffix(templatePath, ".go") {
		scanner.Scan()
	}
	var temp string
	for scanner.Scan() {
		temp = temp + scanner.Text() + "\n"
	}

	template, err := template.New(path.Base(templatePath)).Funcs(functions).Parse(temp)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// create output file
	outputFile := filepath.Join(outputPath)
	os.MkdirAll(filepath.Dir(outputFile), 0755)
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}

	output := new(bytes.Buffer)
	err = template.Execute(output, config)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	f.Write(output.Bytes())
}

func main() {
	featureTemplateFiles, _ := os.ReadDir(featureTemplateDefinitionsPath)
	featureTemplateConfigs := make([]YamlConfig, len(featureTemplateFiles))
	configs := make(map[string][]YamlConfig)

	// Load feature template configs
	for i, filename := range featureTemplateFiles {
		yamlFile, err := os.ReadFile(filepath.Join(featureTemplateDefinitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		featureTemplateConfigs[i] = config
	}

	for i := range featureTemplateConfigs {
		// Augment feature template config by model data
		augmentFeatureTemplateConfig(&featureTemplateConfigs[i])

		// Iterate over templates and render files
		for _, t := range featureTemplateTemplates {
			renderTemplate(t.path, t.prefix+SnakeCase(featureTemplateConfigs[i].Name)+t.suffix, featureTemplateConfigs[i])
		}
	}
	configs["FeatureTemplates"] = featureTemplateConfigs

	profileParcelFiles, _ := os.ReadDir(profileParcelDefinitionsPath)
	profileParcelConfigs := make([]YamlConfig, len(profileParcelFiles))

	// Load profile parcel configs
	for i, filename := range profileParcelFiles {
		yamlFile, err := os.ReadFile(filepath.Join(profileParcelDefinitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		profileParcelConfigs[i] = config
	}

	for i := range profileParcelConfigs {
		// Augment profile parcel config by model data
		augmentProfileParcelConfig(&profileParcelConfigs[i])

		// Iterate over templates and render files
		for _, t := range profileParcelTemplates {
			renderTemplate(t.path, t.prefix+SnakeCase(profileParcelConfigs[i].Name)+t.suffix, profileParcelConfigs[i])
		}
	}
	configs["ProfileParcels"] = profileParcelConfigs

	genericFiles, _ := os.ReadDir(genericDefinitionsPath)
	genericConfigs := make([]YamlConfig, len(genericFiles))

	// Load generic configs
	for i, filename := range genericFiles {
		yamlFile, err := os.ReadFile(filepath.Join(genericDefinitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		genericConfigs[i] = config
	}

	for i := range genericConfigs {
		// Augment generic config
		augmentGenericConfig(&genericConfigs[i], "")

		// Iterate over templates and render files
		for _, t := range genericTemplates {
			renderTemplate(t.path, t.prefix+SnakeCase(genericConfigs[i].Name)+t.suffix, genericConfigs[i])
		}
	}
	configs["Generic"] = genericConfigs

	// render provider.go
	renderTemplate(providerTemplate, providerLocation, configs)

	changelog, err := os.ReadFile(changelogOriginal)
	if err != nil {
		log.Fatalf("Error reading changelog: %v", err)
	}
	renderTemplate(changelogTemplate, changelogLocation, string(changelog))
}

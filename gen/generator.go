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
	"regexp"
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
		suffix: "%s.go",
	},
	{
		path:   "./gen/templates/profile_parcels/data_source.go",
		prefix: "./internal/provider/data_source_sdwan_",
		suffix: "%s.go",
	},
	{
		path:   "./gen/templates/profile_parcels/data_source_test.go",
		prefix: "./internal/provider/data_source_sdwan_",
		suffix: "%s_test.go",
	},
	{
		path:   "./gen/templates/profile_parcels/resource.go",
		prefix: "./internal/provider/resource_sdwan_",
		suffix: "%s.go",
	},
	{
		path:   "./gen/templates/profile_parcels/resource_test.go",
		prefix: "./internal/provider/resource_sdwan_",
		suffix: "%s_test.go",
	},
	{
		path:   "./gen/templates/profile_parcels/data-source.tf",
		prefix: "./examples/data-sources/sdwan_",
		suffix: "%s/data-source.tf",
	},
	{
		path:   "./gen/templates/profile_parcels/resource.tf",
		prefix: "./examples/resources/sdwan_",
		suffix: "%s/resource.tf",
	},
	{
		path:   "./gen/templates/profile_parcels/import.sh",
		prefix: "./examples/resources/sdwan_",
		suffix: "%s/import.sh",
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
	SkipMinimumTest          bool                  `yaml:"skip_minimum_test"`
	TestPrerequisites        string                `yaml:"test_prerequisites"`
	RemoveId                 bool                  `yaml:"remove_id"`
	TypeValue                string                `yaml:"type_value"`
	NoImport                 bool                  `yaml:"no_import"`
	NoResource               bool                  `yaml:"no_resource"`
	NoDataSource             bool                  `yaml:"no_data_source"`
	GetBeforeDelete          bool                  `yaml:"get_before_delete"`
	DeleteMutex              bool                  `yaml:"delete_mutex"`
	ParcelType               string                `yaml:"parcel_type"`
	FullUpdate               bool                  `yaml:"full_update"`
}

type YamlConfigAttribute struct {
	ModelName               string                         `yaml:"model_name"`
	ResponseModelName       string                         `yaml:"response_model_name"`
	TfName                  string                         `yaml:"tf_name"`
	Type                    string                         `yaml:"type"`
	ElementType             string                         `yaml:"element_type"`
	ObjectType              string                         `yaml:"object_type"`
	ModelTypeString         bool                           `yaml:"model_type_string"`
	BoolEmptyString         bool                           `yaml:"bool_empty_string"`
	DataPath                []string                       `yaml:"data_path"`
	ResponseDataPath        []string                       `yaml:"response_data_path"`
	Keys                    []string                       `yaml:"keys"`
	Id                      bool                           `yaml:"id"`
	Reference               bool                           `yaml:"reference"`
	Variable                bool                           `yaml:"variable"`
	Mandatory               bool                           `yaml:"mandatory"`
	IgnoreMandatory         bool                           `yaml:"ignore_mandatory"`
	Optional                bool                           `yaml:"optional"`
	WriteOnly               bool                           `yaml:"write_only"`
	TfOnly                  bool                           `yaml:"tf_only"`
	ExcludeTest             bool                           `yaml:"exclude_test"`
	ExcludeExample          bool                           `yaml:"exclude_example"`
	ExcludeIgnore           bool                           `yaml:"exclude_ignore"`
	ExcludeNull             bool                           `yaml:"exclude_null"`
	NodeOnlyContainer       bool                           `yaml:"node_only_container"`
	Description             string                         `yaml:"description"`
	Example                 string                         `yaml:"example"`
	EnumValues              []string                       `yaml:"enum_values"`
	IgnoreEnum              bool                           `yaml:"ignore_enum"`
	MinList                 int64                          `yaml:"min_list"`
	MaxList                 int64                          `yaml:"max_list"`
	MinInt                  int64                          `yaml:"min_int"`
	MaxInt                  int64                          `yaml:"max_int"`
	MinFloat                float64                        `yaml:"min_float"`
	MaxFloat                float64                        `yaml:"max_float"`
	StringPatterns          []string                       `yaml:"string_patterns"`
	StringMinLength         int64                          `yaml:"string_min_length"`
	StringMaxLength         int64                          `yaml:"string_max_length"`
	DefaultValue            string                         `yaml:"default_value"`
	DefaultValuePresent     bool                           `yaml:"default_value_present"`
	DefaultValueEmptyString bool                           `yaml:"default_value_empty_string"`
	Value                   string                         `yaml:"value"`
	TestValue               string                         `yaml:"test_value"`
	SecondaryTestValue      string                         `yaml:"secondary_test_value"`
	MinimumTestValue        string                         `yaml:"minimum_test_value"`
	AlwaysInclude           bool                           `yaml:"always_include"`
	AlwaysIncludeParent     bool                           `yaml:"always_include_parent"`
	Attributes              []YamlConfigAttribute          `yaml:"attributes"`
	ConditionalAttribute    YamlConfigConditionalAttribute `yaml:"conditional_attribute"`
	ConditionalListLength   string                         `yaml:"conditional_list_length"`
	QueryParam              bool                           `yaml:"query_param"`
	NoAugmentConfig         bool                           `yaml:"no_augment_config"`
	TestTags                []string                       `yaml:"test_tags"`
	RequiresConstAndVar     bool                           `yaml:"requires_const_and_var"`
	RequiresReplace         bool                           `yaml:"requires_replace"`
	DynamicDefault          bool                           `yaml:"dynamic_default"`
}

type YamlConfigConditionalAttribute struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
	Type  string `yaml:"type"`
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

func ToVersionName(s string) string {
	s = ToGoName(s)
	m := strings.ReplaceAll(s, "Versions", "Ids")
	m = strings.ReplaceAll(m, "Version", "Id")
	return m
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

// Templating helper function to return true if id included in attributes
func HasId(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Id {
			return true
		}
	}
	return false
}

// Templating helper function to return true if name included in attributes
func HasName(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.TfName == "name" {
			return true
		}
	}
	return false
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

// Templating helper function to return ResponseModelName if set, otherwise ModelName, including the path
func GetResponseModelPath(attribute YamlConfigAttribute) string {
	modelName := ""
	if attribute.ResponseModelName != "" {
		modelName = attribute.ResponseModelName
	} else {
		modelName = attribute.ModelName
	}
	if len(attribute.ResponseDataPath) > 0 {
		return strings.Join(attribute.ResponseDataPath, ".") + "." + modelName
	} else if len(attribute.DataPath) > 0 {
		return strings.Join(attribute.DataPath, ".") + "." + modelName
	}
	return modelName
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

// Templating helper function to return true if UX 2.0 related feature
func IsUx20Feature(config YamlConfig) bool {
	return !strings.HasPrefix(config.DocCategory, "(Classic)")
}

// Templating helper function to return number of reference included in attributes
func CountReferences(attributes []YamlConfigAttribute) int {
	count := 0
	for _, attr := range attributes {
		if attr.Reference {
			count++
		}
	}
	return count
}

// Templating helper function to add two integer values
func Add(x, y int) int {
	return x + y
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

// Templating helper function to return true if type is StringInt64
func IsStringInt64(attribute YamlConfigAttribute) bool {
	if attribute.Type == "StringInt64" {
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

// Templating helper function to return parent model name in case data_path is used
func GetParentModelName(attribute YamlConfigAttribute) string {
	if len(attribute.DataPath) > 0 {
		return attribute.DataPath[0]
	} else {
		return attribute.ModelName
	}
}

// Templating helper function to return the snake case profile parcel suffix
func GetProfileParcelSuffix(config YamlConfig) string {
	if config.ParcelType == "feature" {
		return "_feature"
	} else if config.ParcelType == "policy" {
		return "_policy"
	} else {
		return ""
	}
}

// Templating helper function to return the snake case profile parcel name
func GetProfileParcelName(config YamlConfig) string {
	return SnakeCase(config.Name) + GetProfileParcelSuffix(config)
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
	"toGoName":               ToGoName,
	"toVersionName":          ToVersionName,
	"camelCase":              CamelCase,
	"snakeCase":              SnakeCase,
	"sprintf":                fmt.Sprintf,
	"toLower":                strings.ToLower,
	"path":                   BuildPath,
	"hasId":                  HasId,
	"hasName":                HasName,
	"hasVersionAttribute":    HasVersionAttribute,
	"getResponseModelPath":   GetResponseModelPath,
	"hasReference":           HasReference,
	"countReferences":        CountReferences,
	"add":                    Add,
	"getGjsonType":           GetGjsonType,
	"getId":                  GetId,
	"isUx20Feature":          IsUx20Feature,
	"isListSet":              IsListSet,
	"isList":                 IsList,
	"isSet":                  IsSet,
	"isStringListSet":        IsStringListSet,
	"isInt64ListSet":         IsInt64ListSet,
	"isStringInt64":          IsStringInt64,
	"isNestedListSet":        IsNestedListSet,
	"isNestedList":           IsNestedList,
	"isNestedSet":            IsNestedSet,
	"getParentModelName":     GetParentModelName,
	"getProfileParcelSuffix": GetProfileParcelSuffix,
	"getProfileParcelName":   GetProfileParcelName,
	"contains":               contains,
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
			if attr.Type == "" {
				attr.Type = "Set"
			}
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
		if attr.Type == "" {
			attr.Type = "Set"
		}
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

func parseProfileParcelAttribute(attr *YamlConfigAttribute, model gjson.Result, isOneOfAttribute bool) {
	if attr.ModelName == "" {
		return
	}
	path := ""
	prefix := "properties."
	for i, e := range attr.DataPath {

		next := ""
		if i+1 < len(attr.DataPath) {
			next = attr.DataPath[i+1]
		} else {
			next = attr.ModelName
		}

		if model.Get("oneOf").Exists() {
			prefix = "oneOf."
		}

		if path == "" && prefix == "oneOf." {
			index := 0
			model.Get("oneOf").ForEach(func(k, v gjson.Result) bool {
				if v.Get(fmt.Sprintf("properties.%v.properties.%v", e, next)).Exists() {
					path += fmt.Sprintf("%v.properties.%v.properties.", index, e)
					return false // stop iterating
				}
				index += 1
				return true // keep iterating
			})
		} else if model.Get(prefix + path + e + ".oneOf").Exists() {
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

		} else if model.Get(prefix + path + e + ".items").Exists() {
			path += fmt.Sprintf("%s.items.properties.", e)

		} else {
			path += e + ".properties."
		}
	}

	path += attr.ModelName

	r := model.Get(prefix + path)

	if !r.Exists() {
		panic(fmt.Sprintf("Could not find element in schema: %v\n%v\n\n", attr.ModelName, model))
	}

	attr.Description = r.Get("description").String()
	if attr.TfName == "" {
		attr.TfName = SnakeCase(attr.ModelName)
	}

	if attr.ConditionalAttribute.Name != "" {
		attr.ConditionalAttribute.Name = SnakeCase(attr.ConditionalAttribute.Name)
	}

	if r.Get("type").String() == "object" || !r.Get("type").Exists() {
		noGlobal := false

		t := r.Get("oneOf.#(properties.optionType.enum.0=\"global\")")
		if value := r.Get("properties.optionType.enum.0"); value.String() == "global" {
			t = r
		}

		if t.Exists() {
			if attr.Type == "String" || attr.Type == "StringInt64" || t.Get("properties.value.type").String() == "string" || t.Get("properties.value.anyOf.0.type").String() == "string" || t.Get("properties.value.oneOf.0.type").String() == "string" {
				if attr.Type != "StringInt64" {
					attr.Type = "String"
				}
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
			} else if attr.Type == "Bool" || t.Get("properties.value.type").String() == "boolean" {
				attr.Type = "Bool"
			} else if attr.Type == "Int64" || t.Get("properties.value.type").String() == "integer" || t.Get("properties.value.type").String() == "number" || t.Get("properties.value.oneOf.0.type").String() == "integer" || t.Get("properties.value.oneOf.0.type").String() == "number" {

				if value := t.Get("properties.value.multipleOf"); value.Exists() {
					attr.Type = "Float64"
					if value := t.Get("properties.value.minimum"); value.Exists() {
						attr.MinFloat = value.Float()
					}
					if value := t.Get("properties.value.maximum"); value.Exists() {
						attr.MaxFloat = value.Float()
					}
				} else {
					attr.Type = "Int64"
					if value := t.Get("properties.value.minimum"); value.Exists() {
						attr.MinInt = value.Int()
					}
					if value := t.Get("properties.value.maximum"); value.Exists() {
						attr.MaxInt = value.Int()
					}
				}
			} else if attr.Type == "List" && attr.ElementType != "" {
				// Remove unsupported type error and enable model overrides
			} else if attr.Type == "Set" && attr.ElementType == "String" || t.Get("properties.value.type").String() == "array" && t.Get("properties.value.items.type").String() == "string" || t.Get("properties.value.type").String() == "array" && t.Get("properties.value.items.oneOf.0.type").String() == "string" {
				attr.Type = "Set"
				attr.ElementType = "String"
				// if value := t.Get("properties.value.items.minItems"); value.Exists() {
				//  attr.MinList = value.Int()
				// }
				// if value := t.Get("properties.value.items.maxItems"); value.Exists() {
				//  attr.MaxList = value.Int()
				// }
			} else if attr.Type == "Set" && attr.ElementType == "Int64" || t.Get("properties.value.type").String() == "array" && t.Get("properties.value.items.type").String() == "integer" {
				attr.Type = "Set"
				attr.ElementType = "Int64"
				// if value := t.Get("properties.value.items.minimum"); value.Exists() {
				//  attr.MinInt = value.Int()
				// }
				// if value := t.Get("properties.value.items.maximum"); value.Exists() {
				//  attr.MaxInt = value.Int()
				// }
			} else if t.Get("properties.value.const").String() == "true" || t.Get("properties.value.const").String() == "false" {
				attr.Type = "Bool"
			} else if t.Get("properties.value.const").String() == "off" || t.Get("properties.value.const").String() == "on" || t.Get("properties.value.const").Exists() {
				attr.Type = "String"
			} else {
				fmt.Printf("WARNING: Unsupported type: %s\n", t.Get("properties.value.type").String())
			}
		} else {
			noGlobal = true
		}
		if r.Get("oneOf.#(properties.optionType.enum.0=\"variable\")").Exists() {
			attr.Variable = true
		}

		d := r.Get("oneOf.#(properties.optionType.enum.0=\"default\")")
		if value := r.Get("properties.optionType.enum.0"); value.String() == "default" {
			d = r
		}
		if d.Exists() && (!isOneOfAttribute || attr.DefaultValuePresent == true) {
			attr.DefaultValuePresent = true
			if value := d.Get("properties.value.enum.0"); value.Exists() {
				if value.String() == "" {
					attr.DefaultValueEmptyString = true
				} else {
					if noGlobal {
						attr.Value = value.String()
					} else {
						attr.DefaultValue = value.String()
					}
				}
			} else if value := d.Get("properties.value.default"); value.Exists() {
				if value.String() == "" {
					attr.DefaultValueEmptyString = true
				} else {
					if noGlobal {
						attr.Value = value.String()
					} else {
						attr.DefaultValue = value.String()
					}
				}
			} else if value := d.Get("properties.value.minimum"); value.Exists() {
				if noGlobal {
					attr.Value = value.String()
				} else {
					attr.DefaultValue = value.String()
				}
			}
		} else if isOneOfAttribute {
			attr.ExcludeNull = true
		} else {
			if !attr.Variable && !attr.IgnoreMandatory {
				attr.Mandatory = true
			}
		}
	} else if r.Get("type").String() == "array" && r.Get("items.type").String() == "object" || r.Get("items.oneOf.0.type").String() == "object" && len(attr.Attributes) > 0 {
		attr.Type = "List"
		if r.Get("minItems").Exists() {
			attr.MinList = r.Get("minItems").Int()
		}
		if r.Get("maxItems").Exists() {
			attr.MaxList = r.Get("maxItems").Int()
		}
		if isOneOfAttribute {
			attr.ExcludeNull = true
		}
		for a := range attr.Attributes {
			parseProfileParcelAttribute(&attr.Attributes[a], r.Get("items"), isOneOfAttribute)
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
		if model.Get("request.properties.data.oneOf").Exists() {
			parseProfileParcelAttribute(&config.Attributes[ia], model.Get("request.properties.data.oneOf.0"), false)
		} else {
			parseProfileParcelAttribute(&config.Attributes[ia], model.Get("request.properties.data"), false)
		}
	}
	if config.DsDescription == "" {
		config.DsDescription = fmt.Sprintf("This data source can read the %s %s.", config.Name, CamelCase(config.ParcelType))
	}
	if config.ResDescription == "" {
		config.ResDescription = fmt.Sprintf("This resource can manage a %s %s.", config.Name, CamelCase(config.ParcelType))
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

func getTemplateSection(content, name string) string {
	scanner := bufio.NewScanner(strings.NewReader(content))
	result := ""
	foundSection := false
	beginRegex := regexp.MustCompile(`\/\/template:begin\s` + name + `$`)
	endRegex := regexp.MustCompile(`\/\/template:end\s` + name + `$`)
	for scanner.Scan() {
		line := scanner.Text()
		if !foundSection {
			match := beginRegex.MatchString(line)
			if match {
				foundSection = true
				result += line + "\n"
			}
		} else {
			result += line + "\n"
			match := endRegex.MatchString(line)
			if match {
				foundSection = false
			}
		}
	}
	return result
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

	output := new(bytes.Buffer)
	err = template.Execute(output, config)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	outputFile := filepath.Join(outputPath)
	existingFile, err := os.Open(outputPath)
	if err != nil {
		os.MkdirAll(filepath.Dir(outputFile), 0755)
	} else if strings.HasSuffix(templatePath, ".go") {
		existingScanner := bufio.NewScanner(existingFile)
		var newContent string
		currentSectionName := ""
		beginRegex := regexp.MustCompile(`\/\/template:begin\s(.*?)$`)
		endRegex := regexp.MustCompile(`\/\/template:end\s(.*?)$`)
		for existingScanner.Scan() {
			line := existingScanner.Text()
			if currentSectionName == "" {
				matches := beginRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] != "" {
					currentSectionName = matches[1]
				} else {
					newContent += line + "\n"
				}
			} else {
				matches := endRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] == currentSectionName {
					currentSectionName = ""
					newSection := getTemplateSection(string(output.Bytes()), matches[1])
					newContent += newSection
				}
			}
		}
		output = bytes.NewBufferString(newContent)
	}
	// write to output file
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
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
			parcelSuffix := GetProfileParcelSuffix(profileParcelConfigs[i])
			suffix := fmt.Sprintf(t.suffix, parcelSuffix)
			renderTemplate(t.path, t.prefix+SnakeCase(profileParcelConfigs[i].Name)+suffix, profileParcelConfigs[i])
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
			if (genericConfigs[i].NoImport && t.path == "./gen/templates/generic/import.sh") ||
				(genericConfigs[i].NoDataSource && t.path == "./gen/templates/generic/data_source.go") ||
				(genericConfigs[i].NoDataSource && t.path == "./gen/templates/generic/data_source_test.go") ||
				(genericConfigs[i].NoDataSource && t.path == "./gen/templates/generic/data-source.tf") ||
				(genericConfigs[i].NoResource && t.path == "./gen/templates/generic/resource.go") ||
				(genericConfigs[i].NoResource && t.path == "./gen/templates/generic/resource_test.go") ||
				(genericConfigs[i].NoResource && t.path == "./gen/templates/generic/resource.tf") ||
				(genericConfigs[i].NoResource && t.path == "./gen/templates/generic/import.sh") {
				continue
			}
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

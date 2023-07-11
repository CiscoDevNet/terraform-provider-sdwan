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
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	featureTemplateDefinitionsPath  = "./gen/definitions/feature_templates/"
	policyObjectDefinitionsPath     = "./gen/definitions/policy_objects/"
	policyDefinitionDefinitionsPath = "./gen/definitions/policy_definitions/"
)

type YamlConfig struct {
	Name string `yaml:"name"`
}

var docPaths = []string{"./docs/data-sources/", "./docs/resources/"}

var extraDocs = map[string]string{
	"cli_device_template":            "Device Templates",
	"feature_device_template":        "Device Templates",
	"attach_feature_device_template": "Device Templates",
	"localized_policy":               "Localized Policies",
}

func SnakeCase(s string) string {
	var g []string

	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

func main() {
	featureTemplateFiles, _ := ioutil.ReadDir(featureTemplateDefinitionsPath)
	featureTemplateConfigs := make([]YamlConfig, len(featureTemplateFiles))
	policyObjectFiles, _ := ioutil.ReadDir(policyObjectDefinitionsPath)
	policyObjectConfigs := make([]YamlConfig, len(policyObjectFiles))
	policyDefinitionFiles, _ := ioutil.ReadDir(policyDefinitionDefinitionsPath)
	policyDefinitionConfigs := make([]YamlConfig, len(policyDefinitionFiles))

	// Load feature template configs
	for i, filename := range featureTemplateFiles {
		yamlFile, err := ioutil.ReadFile(filepath.Join(featureTemplateDefinitionsPath, filename.Name()))
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

	// Update feature template doc category
	for i := range featureTemplateConfigs {
		for _, path := range docPaths {
			filename := path + SnakeCase(featureTemplateConfigs[i].Name) + "_feature_template.md"
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatalf("Error opening documentation: %v", err)
			}

			s := string(content)
			s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "Feature Templates"`)

			ioutil.WriteFile(filename, []byte(s), 0644)
		}
	}

	// Load policy object configs
	for i, filename := range policyObjectFiles {
		yamlFile, err := ioutil.ReadFile(filepath.Join(policyObjectDefinitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		policyObjectConfigs[i] = config
	}

	// Update policy object doc category
	for i := range policyObjectConfigs {
		for _, path := range docPaths {
			filename := path + SnakeCase(policyObjectConfigs[i].Name) + "_policy_object.md"
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatalf("Error opening documentation: %v", err)
			}

			s := string(content)
			s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "Policy Objects"`)

			ioutil.WriteFile(filename, []byte(s), 0644)
		}
	}

	// Load policy definition configs
	for i, filename := range policyDefinitionFiles {
		yamlFile, err := ioutil.ReadFile(filepath.Join(policyDefinitionDefinitionsPath, filename.Name()))
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		config := YamlConfig{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			log.Fatalf("Error parsing yaml: %v", err)
		}
		policyDefinitionConfigs[i] = config
	}

	// Update policy definition doc category
	for i := range policyDefinitionConfigs {
		for _, path := range docPaths {
			filename := path + SnakeCase(policyDefinitionConfigs[i].Name) + "_policy_definition.md"
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatalf("Error opening documentation: %v", err)
			}

			s := string(content)
			s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "Localized Policies"`)

			ioutil.WriteFile(filename, []byte(s), 0644)
		}
	}

	// Update extra (non-feature templates) doc categories
	for doc, cat := range extraDocs {
		for _, path := range docPaths {
			filename := path + doc + ".md"
			content, err := ioutil.ReadFile(filename)
			if err == nil {
				s := string(content)
				s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+cat+`"`)

				ioutil.WriteFile(filename, []byte(s), 0644)
			}
		}
	}
}

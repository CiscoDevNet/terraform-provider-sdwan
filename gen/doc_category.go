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
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	featureTemplateDefinitionsPath = "./gen/definitions/feature_templates/"
	profileParcelDefinitionsPath   = "./gen/definitions/profile_parcels/"
	genericDefinitionsPath         = "./gen/definitions/generic/"
)

type YamlConfig struct {
	Name        string `yaml:"name"`
	DocCategory string `yaml:"doc_category"`
	ParcelType  string `yaml:"parcel_type"`
}

var docPaths = []string{"./docs/data-sources/", "./docs/resources/"}

var extraDocs = map[string]string{
	"attach_feature_device_template": "(Classic) Device Templates",
	"activate_centralized_policy":    "(Classic) Centralized Policies",
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
	featureTemplateFiles, _ := os.ReadDir(featureTemplateDefinitionsPath)
	featureTemplateConfigs := make([]YamlConfig, len(featureTemplateFiles))
	profileParcelFiles, _ := os.ReadDir(profileParcelDefinitionsPath)
	profileParcelConfigs := make([]YamlConfig, len(profileParcelFiles))
	genericFiles, _ := os.ReadDir(genericDefinitionsPath)
	genericConfigs := make([]YamlConfig, len(genericFiles))

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

	// Update feature template doc category
	for i := range featureTemplateConfigs {
		for _, path := range docPaths {
			filename := path + SnakeCase(featureTemplateConfigs[i].Name) + "_feature_template.md"
			content, err := os.ReadFile(filename)
			if err != nil {
				log.Fatalf("Error opening documentation: %v", err)
			}

			s := string(content)
			s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "(Classic) Feature Templates"`)

			os.WriteFile(filename, []byte(s), 0644)
		}
	}

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

	// Update profile parcel doc category
	for i := range profileParcelConfigs {
		for _, path := range docPaths {
			extension := ""
			category := ""

			if profileParcelConfigs[i].ParcelType == "feature" {
				extension = "_feature.md"
				category = profileParcelConfigs[i].DocCategory
			} else if profileParcelConfigs[i].ParcelType == "policy" {
				extension = "_policy.md"
				category = "Policies"
			} else if profileParcelConfigs[i].ParcelType == "policy_object" {
				extension = ".md"
				category = "Policy Objects"
			}
			filename := path + SnakeCase(profileParcelConfigs[i].Name) + extension
			content, err := os.ReadFile(filename)
			if err != nil {
				log.Fatalf("Error opening documentation: %v", err)
			}

			s := string(content)
			s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+category+`"`)

			os.WriteFile(filename, []byte(s), 0644)
		}
	}

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

	// Update generic doc category
	for i := range genericConfigs {
		for _, path := range docPaths {
			filename := path + SnakeCase(genericConfigs[i].Name) + ".md"
			content, err := os.ReadFile(filename)
			if err == nil {
				cat := genericConfigs[i].DocCategory
				s := string(content)
				s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+cat+`"`)

				os.WriteFile(filename, []byte(s), 0644)
			}
		}
	}

	// Update extra doc categories
	for doc, cat := range extraDocs {
		for _, path := range docPaths {
			filename := path + doc + ".md"
			content, err := os.ReadFile(filename)
			if err == nil {
				s := string(content)
				s = strings.ReplaceAll(s, `subcategory: ""`, `subcategory: "`+cat+`"`)

				os.WriteFile(filename, []byte(s), 0644)
			}
		}
	}
}

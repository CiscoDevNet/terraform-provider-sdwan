package models

/* Sample
(For device template created from feature template)
{
    "templateName": "Sample_ISR4431_Template",
    "templateDescription": "For testing purpose",
    "deviceType": "vedge-ISR-4431",
    "deviceRole": "sdwan-edge",
    "configType": "template",
    "factoryDefault": false,
    "policyId": "",
    "featureTemplateUidRange": [],
    "connectionPreferenceRequired": true,
    "connectionPreference": true,
    "generalTemplates": [
        {
            "templateId": "6d4f554a-f930-4590-aca7-c3244a18f3be",
            "templateType": "cisco_vpn",
            "subTemplates": [
                {
                    "templateId": "99cfbde3-5fd9-47ed-be97-2bc8fd4c3e33",
                    "templateType": "cisco_vpn_interface"
                }
            ]
        },
        {
            "templateId": "bfcab33b-1c67-4179-9716-5ff753e29d3c",
            "templateType": "cedge_global"
        }
    ]
}

(For device template created from CLI)
{
    "templateName": "Sample_CLI_Device",
    "templateDescription": "For testing purpose",
    "deviceType": "vedge-cloud",
    "templateConfiguration": "string",
    "factoryDefault": false,
    "configType": "file"
}
*/

type Template struct {
	TemplateId   string `json:"templateId"`
	TemplateType string `json:"templateType"`
}

type GeneralTemplate struct {
	TemplateId   string      `json:"templateId"`
	TemplateType string      `json:"templateType"`
	SubTemplates []*Template `json:"subTemplates"`
}

type DeviceTemplate struct {
	TemplateName                 string             `json:"templateName"`
	TemplateDescription          string             `json:"templateDescription"`
	DeviceType                   string             `json:"deviceType"`
	DeviceRole                   string             `json:"deviceRole"`
	ConfigType                   string             `json:"configType"`
	FactoryDefault               bool               `json:"factoryDefault"`
	PolicyId                     string             `json:"policyId"`
	FeatureTemplateUidRange      []string           `json:"featureTemplateUidRange"`
	ConnectionPreferenceRequired bool               `json:"connectionPreferenceRequired"`
	ConnectionPreference         bool               `json:"connectionPreference"`
	GeneralTemplates             []*GeneralTemplate `json:"generalTemplates"`
	TemplateConfiguration        string             `json:"templateConfiguration"`
}

// ToMap - Returns map for Device Template model
func (dt *DeviceTemplate) ToMap() (map[string]interface{}, error) {
	deviceTemplateAttrMap := make(map[string]interface{})

	if dt.TemplateName != "" {
		deviceTemplateAttrMap["templateName"] = dt.TemplateName
	}

	if dt.TemplateDescription != "" {
		deviceTemplateAttrMap["templateDescription"] = dt.TemplateDescription
	}

	if dt.DeviceType != "" {
		deviceTemplateAttrMap["deviceType"] = dt.DeviceType
	}

	if dt.DeviceRole != "" {
		deviceTemplateAttrMap["deviceRole"] = dt.DeviceRole
	}

	if dt.ConfigType != "" {
		deviceTemplateAttrMap["configType"] = dt.ConfigType
	}

	deviceTemplateAttrMap["policyId"] = dt.PolicyId

	deviceTemplateAttrMap["featureTemplateUidRange"] = dt.FeatureTemplateUidRange

	deviceTemplateAttrMap["connectionPreferenceRequired"] = dt.ConnectionPreferenceRequired

	deviceTemplateAttrMap["connectionPreference"] = dt.ConnectionPreference

	if dt.TemplateConfiguration != "" {
		deviceTemplateAttrMap["templateConfiguration"] = dt.TemplateConfiguration
	}

	gts := []map[string]interface{}{}

	for _, generalTemplate := range dt.GeneralTemplates {
		gt := make(map[string]interface{})

		if generalTemplate.TemplateId != "" {
			gt["templateId"] = generalTemplate.TemplateId
		}

		if generalTemplate.TemplateType != "" {
			gt["templateType"] = generalTemplate.TemplateType
		}

		if generalTemplate.SubTemplates != nil {
			tps := []map[string]interface{}{}
			for _, subTemplate := range generalTemplate.SubTemplates {
				tp := make(map[string]interface{})
				if subTemplate.TemplateId != "" {
					tp["templateId"] = subTemplate.TemplateId
				}
				if subTemplate.TemplateType != "" {
					tp["templateType"] = subTemplate.TemplateType
				}
				tps = append(tps, tp)
			}
			gt["subTemplates"] = tps
		} 

		gts = append(gts, gt)
	}

	deviceTemplateAttrMap["generalTemplates"] = gts

	return deviceTemplateAttrMap, nil
}

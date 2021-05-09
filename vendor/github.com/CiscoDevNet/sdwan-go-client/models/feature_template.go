package models

type FeatureTemplate struct {
	TemplateName        string                 `json:"template_name"`
	TemplateDescription string                 `json:"template_description"`
	TemplateType        string                 `json:"template_type"`
	DeviceType          []string               `json:"device_type"`
	TemplateMinVersion  string                 `json:"template_min_version"`
	TemplateDefinition  map[string]interface{} `json:"template_definition"`
	FactoryDefault      bool                   `json:"factory_default"`
}

//ToMap - Returns map for FeatureTemplate model
func (ft *FeatureTemplate) ToMap() (map[string]interface{}, error) {
	ftAttrMap := make(map[string]interface{})

	if ft.TemplateName != "" {
		ftAttrMap["templateName"] = ft.TemplateName
	}

	if ft.TemplateDescription != "" {
		ftAttrMap["templateDescription"] = ft.TemplateDescription
	}

	if ft.TemplateType != "" {
		ftAttrMap["templateType"] = ft.TemplateType
	}

	ftAttrMap["deviceType"] = ft.DeviceType

	if ft.TemplateMinVersion != "" {
		ftAttrMap["templateMinVersion"] = ft.TemplateMinVersion
	}

	ftAttrMap["factoryDefault"] = ft.FactoryDefault

	ftAttrMap["templateDefinition"] = ft.TemplateDefinition

	return ftAttrMap, nil
}

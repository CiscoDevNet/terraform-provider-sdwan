package models

type Template struct {
	TemplateId   string `json:"templateId"`
	TemplateType string `json:"templateType"`
}

type GeneralTemplate struct {
	TemplateId   string     `json:"templateId"`
	TemplateType string     `json:"templateType"`
	SubTemplates []*Template `json:"subTemplates"`
}

type DeviceTemplate struct {
	TemplateName                 string            `json:"templateName"`
	TemplateDescription          string            `json:"templateDescription"`
	DeviceType                   string            `json:"deviceType"`
	DeviceRole                   string            `json:"deviceRole"`
	ConfigType                   string            `json:"configType"`
	FactoryDefault               bool              `json:"factoryDefault"`
	PolicyId                     string            `json:"policyId"`
	FeatureTemplateUidRange      []string          `json:"featureTemplateUidRange"`
	ConnectionPreferenceRequired bool              `json:"connectionPreferenceRequired"`
	ConnectionPreference         bool              `json:"connectionPreference"`
	GeneralTemplates             []*GeneralTemplate `json:"generalTemplates"`
	TemplateConfiguration        string            `json:"templateConfiguration"`
}

func (st *Template) MakeSubTemplate() (map[string]interface{}, error) {
	stMap := make(map[string]interface{})

	A(stMap, "templateId", st.TemplateId)

	A(stMap, "templateType", st.TemplateType)

	return stMap, nil
}

func (gt *GeneralTemplate) MakeGeneralTemplate() (map[string]interface{}, error) {
	gtMap := make(map[string]interface{})

	A(gtMap, "templateId", gt.TemplateId)

	A(gtMap, "templateType", gt.TemplateType)

	if len(gt.SubTemplates) > 0 {
		stList := make([]interface{}, 0)
		for _, st := range gt.SubTemplates {
			stMap, _ := st.MakeSubTemplate()
			stList = append(stList, stMap)
		}
		A(gtMap, "subTemplates", stList)
	}
	return gtMap, nil
}

// ToMap - Returns map for Device Template model
func (dt *DeviceTemplate) ToMap() (map[string]interface{}, error) {
	deviceTemplateAttrMap := make(map[string]interface{})

	A(deviceTemplateAttrMap, "templateName", dt.TemplateName)

	A(deviceTemplateAttrMap, "templateDescription", dt.TemplateDescription)

	A(deviceTemplateAttrMap, "deviceType", dt.DeviceType)

	A(deviceTemplateAttrMap, "deviceRole", dt.DeviceRole)

	A(deviceTemplateAttrMap, "configType", dt.ConfigType)

	A(deviceTemplateAttrMap, "factoryDefault", dt.FactoryDefault)

	A(deviceTemplateAttrMap, "policyId", dt.PolicyId)

	A(deviceTemplateAttrMap, "connectionPreferenceRequired", dt.ConnectionPreferenceRequired)

	A(deviceTemplateAttrMap, "connectionPreference", dt.ConnectionPreference)

	A(deviceTemplateAttrMap, "templateConfiguration", dt.TemplateConfiguration)

	if len(dt.GeneralTemplates) > 0 {
		gtList := make([]interface{}, 0)
		for _, gt := range dt.GeneralTemplates {
			gtMap, _ := gt.MakeGeneralTemplate()
			gtList = append(gtList, gtMap)
		}
		A(deviceTemplateAttrMap, "generalTemplates", gtList)
	}

	if len(dt.FeatureTemplateUidRange) > 0 {
		A(deviceTemplateAttrMap, "featureTemplateUidRange", dt.FeatureTemplateUidRange)
		
	} else {
		deviceTemplateAttrMap["featureTemplateUidRange"] = make([]string, 0, 1)
		
	}
	return deviceTemplateAttrMap, nil
}

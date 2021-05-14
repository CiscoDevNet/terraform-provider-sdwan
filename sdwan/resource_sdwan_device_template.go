package sdwan

import (
	"fmt"
	"log"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/CiscoDevNet/sdwan-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSDWANDeviceTemplate() *schema.Resource {
	return &schema.Resource{
		Create:        resourceSDWANDeviceTemplateCreate,
		Read:          resourceSDWANDeviceTemplateRead,
		Update:        resourceSDWANDeviceTemplateUpdate,
		Delete:        resourceSDWANDeviceTemplateDelete,
		Description:   "Manages Cisco SD-WAN Device Template",
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"template_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: NameValidation(),
				Description:  "Unique Device Template Name",
			},
			"template_description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Long Description of Device Template",
			},
			"device_type": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validDeviceTypesforDeviceType, false)),
				Description:      "Type of device for which Device Template should be created",
			},
			"device_role": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validDeviceRole, false)),
				Description:      "Device role of the device",
			},
			"config_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validConfigTypes, false)),
				Description:      "Configuration type for  Device Template",
			},
			"factory_default": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Boolean flag indicating whether Device Template is factory default or not",
			},
			"policy_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "policyId for  Device Template",
			},
			"feature_template_uid_range": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "featureTemplateUidRange for  Device Template",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"connection_preference_required": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "connectionPreferenceRequired flagfor Device Template",
			},
			"connection_preference": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "connectionPreference for Device Template",
			},
			"last_updated_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User who updated  Device Template last",
			},
			"last_updated_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Time when Device Template last updated",
			},
			"template_class": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Template Class type for  Device Template",
			},
			"template_configuration": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"template_configuration", "general_templates"},
				Description:  "Device Template configuration",
			},
			"template_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Device Template id for  Device Template",
			},
			"general_templates": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				MinItems:    1,
				Description: "List of Feature Templates and thier Sub Templates thourgh which Device Template is created",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Template Id of Feature Template",
						},
						"template_type": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validTemplateTypes, false)),
							Description:      "Template Type of Feature Template",
						},
						"sub_templates": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "List of Sub Templates associated with feature Template",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sub_template_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Sub Template Id of Feature Template",
									},
									"sub_template_type": {
										Type:             schema.TypeString,
										Required:         true,
										ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validTemplateTypes, false)),
										Description:      "Sub Template Type of Feature Template",
									},
								},
							},
						},
					},
				},
			},
			"created_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Time when  Device Template is created",
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User who created Device Template",
			},
			"rid": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Request ID for Device Template",
			},
			"feature": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Feature for Device Template",
			},
			"template_attached": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of templates attached to Device Template",
			},
		},
	}
}

func resourceSDWANDeviceTemplateCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Begining Create method ")

	sdwanClient := m.(*client.Client)

	deviceTemplate := &models.DeviceTemplate{
		TemplateName:        d.Get("template_name").(string),
		TemplateDescription: d.Get("template_description").(string),
		DeviceType:          d.Get("device_type").(string),
		ConfigType:          d.Get("config_type").(string),
		FactoryDefault:      d.Get("factory_default").(bool),
	}

	var tempID string

	if deviceTemplate.ConfigType == "template" {
		if connPref, ok := d.GetOk("connection_preference"); ok {
			deviceTemplate.ConnectionPreference = connPref.(bool)
		}

		if connPrefReq, ok := d.GetOk("connection_preference_required"); ok {
			deviceTemplate.ConnectionPreferenceRequired = connPrefReq.(bool)
		}

		if uidRange, ok := d.GetOk("feature_template_uid_range"); ok {
			deviceTemplate.FeatureTemplateUidRange = interfaceToStrList(uidRange)
		} else {
			deviceTemplate.FeatureTemplateUidRange = make([]string, 0, 1)
		}

		if polID, ok := d.GetOk("policy_id"); ok {
			deviceTemplate.PolicyId = polID.(string)
		} else {
			deviceTemplate.PolicyId = "{}"
		}

		if val, exists := d.GetOk("general_templates"); exists {
			deviceTemplate.GeneralTemplates = expandGeneralTemplates(val.(*schema.Set).List())
		}

		dURL := fmt.Sprintf("dataservice/template/device/feature")
		cont, err := sdwanClient.Save(dURL, deviceTemplate)
		if err != nil {
			return err
		}

		tempID = stripQuotes(cont.S("templateId").String())
	} else {

		if tconfig, ok := d.GetOk("template_configuration"); ok {
			deviceTemplate.TemplateConfiguration = tconfig.(string)
		}

		dURL := fmt.Sprintf("/dataservice/template/device/cli")
		cont, err := sdwanClient.Save(dURL, deviceTemplate)

		if err != nil {
			return err
		}

		tempID = stripQuotes(cont.S("templateId").String())
	}

	d.SetId(tempID)

	log.Println("[DEBUG] Ending Create method ", d.Id())

	return resourceSDWANDeviceTemplateRead(d, m)
}

func resourceSDWANDeviceTemplateRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Begining Read method ", d.Id())

	sdwanClient := m.(*client.Client)

	dtid := d.Id()

	dURL := fmt.Sprintf(`dataservice/template/device/object/%v`, dtid)

	data, err := sdwanClient.GetViaURL(dURL)
	if err != nil {
		return err
	}

	d.Set("template_name", stripQuotes(data.S("templateName").String()))
	d.Set("template_description", stripQuotes(data.S("templateDescription").String()))
	d.Set("device_type", stripQuotes(data.S("deviceType").String()))

	if data.Exists("deviceRole") {
		d.Set("device_role", stripQuotes(data.S("deviceRole").String()))
	}

	d.Set("config_type", stripQuotes(data.S("configType").String()))
	d.Set("factory_default", data.S("factoryDefault").Data().(bool))

	if data.Exists("policyId") {
		d.Set("policy_id", stripQuotes(data.S("policyId").String()))
	}

	if data.Exists("connectionPreferenceRequired") {
		d.Set("connection_preference_required", data.S("connectionPreferenceRequired").Data().(bool))
	}

	if data.Exists("connectionPreference") {
		d.Set("connection_preference", data.S("connectionPreference").Data().(bool))
	}

	if data.Exists("lastUpdatedBy") {
		d.Set("last_updated_by", stripQuotes(data.S("lastUpdatedBy").String()))
	}

	if data.Exists("lastUpdatedOn") {
		d.Set("last_updated_on", int((data.S("lastUpdatedOn").Data()).(float64)))
	}

	if data.Exists("templateClass") {
		d.Set("template_class", stripQuotes(data.S("templateClass").String()))
	}

	if data.Exists("templateConfiguration") {
		d.Set("template_configuration", data.S("templateConfiguration").Data().(string))
	}

	d.Set("template_id", stripQuotes(data.S("templateId").String()))

	if data.Exists("createdBy") {
		d.Set("created_by", stripQuotes(data.S("createdBy").String()))
	}

	if data.Exists("createdOn") {
		d.Set("created_on", int((data.S("createdOn").Data()).(float64)))
	}

	if data.Exists("feature") {
		d.Set("feature", stripQuotes(data.S("feature").String()))
	}

	if data.Exists("@rid") {
		d.Set("rid", int((data.S("@rid").Data()).(float64)))
	}

	if len(data.S("featureTemplateUidRange").Children()) > 0 {
		fturcontList := data.S("featureTemplateUidRange")
		fturList := make([]string, 0, 1)
		for i := 0; i < len(fturcontList.Data().([]interface{})); i++ {
			fturList = append(fturList, stripQuotes(fturcontList.Index(i).String()))
		}
		d.Set("feature_template_uid_range", fturList)
	} else {
		d.Set("feature_template_uid_range", make([]string, 0, 1))
	}

	if len(data.S("generalTemplates").Children()) > 0 {
		gtcontList := data.S("generalTemplates")
		gtList := make([]interface{}, 0, 1)
		for i := 0; i < len(gtcontList.Data().([]interface{})); i++ {
			gtMap := make(map[string]interface{})

			gtMap["template_id"] = stripQuotes(gtcontList.Index(i).S("templateId").String())

			gtMap["template_type"] = stripQuotes(gtcontList.Index(i).S("templateType").String())

			if len(gtcontList.Index(i).S("subTemplates").Children()) > 0 {
				stcontList := gtcontList.Index(i).S("subTemplates")
				stList := make([]interface{}, 0, 1)
				for j := 0; j < len(stcontList.Data().([]interface{})); j++ {
					stMap := make(map[string]string)

					stMap["sub_template_id"] = stripQuotes(stcontList.Index(j).S("templateId").String())

					stMap["sub_template_type"] = stripQuotes(stcontList.Index(j).S("templateType").String())

					stList = append(stList, stMap)
				}

				gtMap["sub_templates"] = stList
			} else {
				gtMap["sub_templates"] = make([]interface{}, 0, 1)
			}
			gtList = append(gtList, gtMap)
		}

		d.Set("general_templates", gtList)
	} else {
		d.Set("general_templates", make([]interface{}, 0, 1))
	}

	d.SetId(stripQuotes(data.S("templateId").String()))

	log.Println("[DEBUG] Ending Read method", d.Id())

	return nil
}

func resourceSDWANDeviceTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Begining Update method ", d.Id())

	devTempID := d.Id()

	sdwanClient := m.(*client.Client)

	deviceTemplate := &models.DeviceTemplate{
		TemplateName:        d.Get("template_name").(string),
		TemplateDescription: d.Get("template_description").(string),
		DeviceType:          d.Get("device_type").(string),
		ConfigType:          d.Get("config_type").(string),
		FactoryDefault:      d.Get("factory_default").(bool),
	}

	if deviceTemplate.ConfigType == "template" {
		if connPref, ok := d.GetOk("connection_preference"); ok {
			deviceTemplate.ConnectionPreference = connPref.(bool)
		}

		if connPrefReq, ok := d.GetOk("connection_preference_required"); ok {
			deviceTemplate.ConnectionPreferenceRequired = connPrefReq.(bool)
		}

		if uidRange, ok := d.GetOk("feature_template_uid_range"); ok {
			deviceTemplate.FeatureTemplateUidRange = interfaceToStrList(uidRange)
		} else {
			deviceTemplate.FeatureTemplateUidRange = make([]string, 0, 1)
		}

		if polID, ok := d.GetOk("policy_id"); ok {
			deviceTemplate.PolicyId = polID.(string)
		} else {
			deviceTemplate.PolicyId = "{}"
		}

		if val, exists := d.GetOk("general_templates"); exists {
			deviceTemplate.GeneralTemplates = expandGeneralTemplates(val.(*schema.Set).List())
		}

	} else {

		if tconfig, ok := d.GetOk("template_configuration"); ok {
			deviceTemplate.TemplateConfiguration = tconfig.(string)
		}

	}
	dURL := fmt.Sprintf("/dataservice/template/device/%v", devTempID)
	_, err := sdwanClient.Update(dURL, deviceTemplate)
	if err != nil {
		return err
	}

	d.SetId(devTempID)

	log.Println("[DEBUG] Ending Update method", d.Id())

	return resourceSDWANDeviceTemplateRead(d, m)
}

func resourceSDWANDeviceTemplateDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Begining Delete method ", d.Id())

	sdwanClient := m.(*client.Client)

	dtid := d.Id()

	dURL := fmt.Sprintf("/dataservice/template/device/%v", dtid)
	_, err := sdwanClient.Delete(dURL)
	if err != nil {
		return err
	}

	d.SetId("")

	log.Println("[DEBUG] Ending Create method")
	return nil
}

func expandGeneralTemplates(generalTemplates []interface{}) []*models.GeneralTemplate {
	GenTemplates := make([]*models.GeneralTemplate, 0, 1)

	for _, val := range generalTemplates {
		GenTemp := &models.GeneralTemplate{}

		GetGenTemp := val.(map[string]interface{})

		GenTemp.TemplateId = GetGenTemp["template_id"].(string)

		GenTemp.TemplateType = GetGenTemp["template_type"].(string)

		if GetGenTemp["sub_templates"] != nil {
			GenTemp.SubTemplates = expandSubTemplates(GetGenTemp["sub_templates"].(*schema.Set).List())
		}

		GenTemplates = append(GenTemplates, GenTemp)
	}

	return GenTemplates
}

func expandSubTemplates(subTemplates []interface{}) []*models.Template {
	genSubTemplates := make([]*models.Template, 0, 1)

	for _, subTemps := range subTemplates {
		subTemp := &models.Template{}

		getSubTemps := subTemps.(map[string]interface{})

		subTemp.TemplateId = getSubTemps["sub_template_id"].(string)

		subTemp.TemplateType = getSubTemps["sub_template_type"].(string)

		genSubTemplates = append(genSubTemplates, subTemp)
	}

	return genSubTemplates
}

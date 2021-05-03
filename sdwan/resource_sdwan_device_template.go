package sdwan

import (
	"log"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/CiscoDevNet/terraform-provider-sdwan/vendor/github.com/CiscoDevNet/sdwan-go-client/models"
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
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Unique Device Template Name",
			},
			"template_description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Long Description of Device Template",
			},
			"device_type": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validDeviceTypes, false)),
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
				Description: "connectionPreferenceRequired for Device Template",
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
				Description: "List of Feature Templates and thier Sub Templates thourgh which Device Template is created",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Template Id of Feature Template",
						},
						"template_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Template Type of Feature Template",
						},
						"sub_templates": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Description: "List of Sub Templates associated with feature Template",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"template_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Template Id of Feature Template",
									},
									"template_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Template Type of Feature Template",
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

	if deviceTemplate.ConfigType=="template"{
		if val,ok:=d.GetOk("connection_preference").(bool)
		deviceTemplate.ConnectionPreferenceRequired:=d.GetOk("")
		if val, exists := d.GetOk("general_templates"); exists {
			gts, errs := expandGeneralTemplates(val.(*schema.Set).List())
		}
	}else{

	}

	
	log.Println("[DEBUG] Ending Create method")
	return nil
}

func expandGeneralTemplates(generalTemplates []interface{}) ([]map[string]interface{}, error) {
	generalTemplate := []map[string]interface{}{}
	for _, gt := range generalTemplates {
		t := gt.(map[string]interface{})
		m := make(map[string]interface{})

		m["template_id"] = t["template_id"]
		m["template_type"] = t["template_type"]

		sts, errs := expandSubTemplates(t["sub_templates"].(*schema.Set).List())

	}
}

func expandSubTemplates(subTemplates []interface{}) ([]map[string]interface{}, error) {
subTemplate:
}

func resourceSDWANDeviceTemplateRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Begining Read method ")
	sdwanClient := m.(*client.Client)
	dtid := d.Get("template_id").(string)
	objurl := "/dataservice/template/device/object/" + dtid
	data, err := sdwanClient.GetViaURL(objurl)
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
		fturList := make([]string, 0)
		for i := 0; i < len(fturcontList.Data().([]interface{})); i++ {
			fturList = append(fturList, stripQuotes(fturcontList.Index(i).String()))
		}
		d.Set("feature_template_uid_range", fturList)
	} else {
		d.Set("feature_template_uid_range", make([]string, 0))
	}

	if len(data.S("generalTemplates").Children()) > 0 {
		gtcontList := data.S("generalTemplates")
		gtList := make([]interface{}, 0)
		for i := 0; i < len(gtcontList.Data().([]interface{})); i++ {
			gtMap := make(map[string]interface{})

			gtMap["template_id"] = stripQuotes(gtcontList.Index(i).S("templateId").String())

			gtMap["template_type"] = stripQuotes(gtcontList.Index(i).S("templateType").String())

			if len(gtcontList.Index(i).S("subTemplates").Children()) > 0 {
				stcontList := gtcontList.Index(i).S("subTemplates")
				stList := make([]interface{}, 0)
				for j := 0; j < len(stcontList.Data().([]interface{})); j++ {
					stMap := make(map[string]string)

					stMap["template_id"] = stripQuotes(stcontList.Index(j).S("templateId").String())

					stMap["template_type"] = stripQuotes(stcontList.Index(j).S("templateType").String())

					stList = append(stList, stMap)
				}

				gtMap["sub_templates"] = stList
			} else {
				gtMap["sub_templates"] = make([]interface{}, 0)
			}
			gtList = append(gtList, gtMap)
		}

		d.Set("general_templates", gtList)
	} else {
		d.Set("general_templates", make([]interface{}, 0))
	}
	log.Println("[DEBUG] Ending Create method")
	return nil
}

func resourceSDWANDeviceTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Begining Update method ")
	log.Println("[DEBUG] Ending Create method")
	return nil
}

func resourceSDWANDeviceTemplateDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Begining Delete method ")
	sdwanClient := m.(*client.Client)
	dtid := d.Get("template_id").(string)
	objurl := "/dataservice/template/device" + dtid
	_, err := sdwanClient.GetViaURL(objurl)
	if err != nil {
		return err
	}
	d.SetId("")
	log.Println("[DEBUG] Ending Create method")
	return nil
}

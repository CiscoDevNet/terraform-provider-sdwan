package sdwan

import (
	"fmt"
	"log"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/CiscoDevNet/sdwan-go-client/container"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceSDWANDeviceTemplate() *schema.Resource {
	return &schema.Resource{
		Read:          dataSourceSDWANDeviceTemplateRead,
		SchemaVersion: 1,
		Description:   "Represents Cisco SD-WAN Device Template",
		Schema: map[string]*schema.Schema{
			"template_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique Device Template Name",
			},
			"template_description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Long Description of Device Template",
			},
			"device_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of device supported by  Device Template",
			},
			"device_role": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Device role for the Device Template",
			},
			"config_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Configuration type for  Device Template",
			},
			"factory_default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Boolean flag indicating whether Device Template is factory default or not",
			},
			"policy_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "policyId for  Device Template",
			},
			"feature_template_uid_range": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "feature_template_uid_range for  Device Template",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"connection_preference_required": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "connection_preference_required for Device Template",
			},
			"connection_preference": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "connection_preference for Device Template",
			},
			"last_updated_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User who updated  Device Template last",
			},
			"template_class": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Template Class type for  Device Template",
			},
			"template_configuration": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Template Configuration for  Device Template",
			},
			"template_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Device Template id for  Device Template",
			},
			"created_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Time when  Device Template is created",
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
			"general_templates": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Feature Templates and thier Sub Templates thourgh which Device Template is created",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"template_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template Id of Feature Template",
						},
						"template_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template Type of Feature Template",
						},
						"sub_templates": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of Sub Templates associated with feature Template",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"template_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Template Id of Feature Template",
									},
									"template_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Template Type of Feature Template",
									},
								},
							},
						},
					},
				},
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User who created Device Template",
			},
			"last_updated_on": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Time when Device Template last updated",
			},
		},
	}
}

func dataSourceSDWANDeviceTemplateRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Begining Read method ")

	sdwanClient := m.(*client.Client)
	dtname := d.Get("template_name").(string)
	datas, err := sdwanClient.GetViaURL("/dataservice/template/device")
	if err != nil {
		return err
	}

	var (
		data     *container.Container = datas.S("data")
		selected *container.Container
		tid      string
	)

	for _, child := range data.Children() {
		if child.S("templateName").Data().(string) == dtname {
			tid = child.S("templateId").Data().(string)
			objurl := "dataservice/template/device/object/" + tid
			selected, err = sdwanClient.GetViaURL(objurl)
			if err != nil {
				return err
			} else {
				d.SetId(child.S("templateName").Data().(string))
			}
			break
		}
	}

	if selected == nil {
		return fmt.Errorf("There is no Device Template named %s", dtname)
	} else {
		d.Set("template_name", stripQuotes(selected.S("templateName").String()))
		d.Set("template_description", stripQuotes(selected.S("templateDescription").String()))
		d.Set("device_type", stripQuotes(selected.S("deviceType").String()))

		if selected.Exists("deviceRole") {
			d.Set("device_role", stripQuotes(selected.S("deviceRole").String()))
		}

		d.Set("config_type", stripQuotes(selected.S("configType").String()))
		d.Set("factory_default", selected.S("factoryDefault").Data().(bool))

		if selected.Exists("policyId") {
			d.Set("policy_id", stripQuotes(selected.S("policyId").String()))
		}

		if selected.Exists("connectionPreferenceRequired") {
			d.Set("connection_preference_required", selected.S("connectionPreferenceRequired").Data().(bool))
		}

		if selected.Exists("connectionPreference") {
			d.Set("connection_preference", selected.S("connectionPreference").Data().(bool))
		}

		if selected.Exists("lastUpdatedBy") {
			d.Set("last_updated_by", stripQuotes(selected.S("lastUpdatedBy").String()))
		}

		if selected.Exists("lastUpdatedOn") {
			d.Set("last_updated_on", int((selected.S("lastUpdatedOn").Data()).(float64)))
		}

		if selected.Exists("templateClass") {
			d.Set("template_class", stripQuotes(selected.S("templateClass").String()))
		}

		if selected.Exists("templateConfiguration") {
			d.Set("template_configuration", selected.S("templateConfiguration").Data().(string))
		}

		d.Set("template_id", stripQuotes(selected.S("templateId").String()))

		if selected.Exists("createdBy") {
			d.Set("created_by", stripQuotes(selected.S("createdBy").String()))
		}

		if selected.Exists("createdOn") {
			d.Set("created_on", int((selected.S("createdOn").Data()).(float64)))
		}

		if selected.Exists("feature") {
			d.Set("feature", stripQuotes(selected.S("feature").String()))
		}

		if selected.Exists("@rid") {
			d.Set("rid", int((selected.S("@rid").Data()).(float64)))
		}

		if len(selected.S("featureTemplateUidRange").Children()) > 0 {
			fturcontList := selected.S("featureTemplateUidRange")
			fturList := make([]string, 0)
			for i := 0; i < len(fturcontList.Data().([]interface{})); i++ {
				fturList = append(fturList, stripQuotes(fturcontList.Index(i).String()))
			}
			d.Set("feature_template_uid_range", fturList)
		} else {
			d.Set("feature_template_uid_range", make([]string, 0))
		}

		if len(selected.S("generalTemplates").Children()) > 0 {
			gtcontList := selected.S("generalTemplates")
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
		// 		for _, key := range []string{
		// 			"template_name",
		// 			"template_description",
		// 			"device_type",
		// 			"device_role",
		// 			"config_type",
		// 			"factory_default",
		// 			"policy_id",
		// 			"connection_preference_required",
		// 			"connection_preference",
		// 			"last_updated_by",
		// 			"last_updated_on",
		// 			"template_class",
		// 			"template_configuration",
		// 			"template_id",
		// 			"created_by",
		// 			"created_on",
		// 			"feature",
		// 			"template_attached",
		// 		} {
		// 			camelCase := snakeCaseToCamelCase(key)
		// 			if selected.Exists(camelCase) {
		// 				d.Set(key, selected.S(camelCase).Data())
		// 			}
		// 			if selected.Exists("@rid") {
		// 				d.Set("rid", selected.S("@rid").Data())
		// 			}
		// 		}

		// 		fturcc := snakeCaseToCamelCase("feature_template_uid_range")
		// 		if len(selected.S(fturcc).Children()) > 0 {
		// 			fturs := []string{}
		// 			for _, ftur := range selected.S(fturcc).Children() {
		// 				switch ftur.Data().(type) {
		// 				case int64:
		// 					{
		// 						val := ftur.Data().(int64)
		// 						str_val := strconv.FormatInt(val, 10)
		// 						fturs = append(fturs, string(str_val))
		// 					}
		// 				case float64:
		// 					{
		// 						val := ftur.Data().(float64)
		// 						str_val := strconv.FormatFloat(val, 'E', -1, 64)
		// 						fturs = append(fturs, string(str_val))
		// 					}
		// 				case string:
		// 					{
		// 						fturs = append(fturs, ftur.Data().(string))
		// 					}

		// 				}

		// 			}
		// 			d.Set("feature_template_uid_range", fturs)
		// 		}

		// 		gtcc := snakeCaseToCamelCase("general_templates")
		// 		if len(selected.S(gtcc).Children()) > 0 {
		// 			gts := containerToMapArrayDeviceTemplate(selected.S(gtcc), []string{"template_id", "template_type", "sub_templates"})
		// 			d.Set("general_templates", gts)
		// 		}
	}

	d.SetId(stripQuotes(selected.S("templateId").String()))

	log.Println("[DEBUG] Ending Read method ")
	return nil
}

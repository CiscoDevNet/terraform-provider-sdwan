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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Boolean flag indicating whether Device Template is factory default or not",
			},
			"policy_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "policy_id for  Device Template",
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
			"template_configuration_edited": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Previous Template Configuration for  Device Template",
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
				Type:        schema.TypeString,
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
		},
	}
}

func dataSourceSDWANDeviceTemplateRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Begining Read method ")
	sdwanClient := m.(*client.Client)
	log.Println("[DEBUG] checking %v", sdwanClient)
	dtname := d.Get("template_name").(string)
	datas, err := sdwanClient.GetViaURL("/dataservice/template/device")
	log.Println("[DEBUG] after GET call")
	if err != nil {
		return err
	}

	if datas != nil {
		log.Println("[DEBUG] container is nil")
	}
	l := len(datas.Children())
	log.Printf("[DEBUG] length of children %d", l)
	var data *container.Container = datas.S("data")
	// log.Println("[DEBUG] reterive data")

	var selected *container.Container
	for _, child := range data.Children() {
		if child.S("templateName").Data().(string) == dtname {
			selected = child
			d.SetId(child.S("template_name").Data().(string))
			break
		}
	}
	if selected == nil {
		return fmt.Errorf("There is no Device Template named %s", dtname)
	} else {
		for _, key := range []string{
			"template_name",
			"template_description",
			"device_type",
			"device_role",
			"config_type",
			"factory_default",
			"policy_id",
			"connection_preference_required",
			"connection_preference",
			"last_updated_by",
			"template_class",
			"template_configuration",
			"template_id",
			"created_on",
			"template_configuration_edited",
			"feature",
			"template_attached",
			"feature_template_uid_range",
			"general_templates",
		} {
			camelCase := snakeCaseToCamelCase(key)
			if selected.Exists(camelCase) {
				d.Set(key, selected.S(camelCase).Data())
			}
			if selected.Exists("@rid") {
				d.Set("rid", selected.S("@rid").Data())
			}
		}
		// fturcc:=snakeCaseToCamelCase("feature_template_uid_range")
		// if selected.Exists(fturcc) {
		// 	fturs := []string{}
		// 	for _, ftur := range selected.S(fturcc).Children() {
		// 		fturs = append(fturs, ftur.Data().(string))
		// 	}
		// 	d.Set("feature_template_uid_range", fturs)
		// }
		// gtcc:=snakeCaseToCamelCase("general_templates")
		// if selected.Exists(gtcc) {
		//     gts:=[]interface{}{}
		//     for _,gt:=range selected.S(gtcc).Children(){
		// 		for _,key:= range[]string{
		// 			"template_id",
		// 			"template_type",
		// 		}{
		// 			camelCase:=snakeCaseToCamelCase(key)
		// 			if gt.Exists(camelCase){
		// 				d.Set(key,gt.S(camelCase).Data())
		// 			}
		// 			stcc:=snakeCaseToCamelCase("sub_templates")
		// 			if gt.Exists(stcc){
		// 				sts:=[]interface{}{}
		// 				for _,st:=range gt.S(stcc).Children(){
		// 					for _,key:=range[]string{
		// 						"template_id",
		// 						"template_type",
		// 					}{
		// 						camelCase:=snakeCaseToCamelCase(key)
		// 						if gt.Exists(camelCase){

		// 						}
		// 					}
		// 				}
		// 			}
		// 		}
		//     }
		// }
		// if selected.Exists("feature_template_uid_range"){
		//     d.Set("feature_template_uid_range",)
		// }
	}
	return nil
}

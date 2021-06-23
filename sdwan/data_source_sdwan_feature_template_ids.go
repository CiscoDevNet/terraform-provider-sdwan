package sdwan

import (
	"fmt"
	"log"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceSDWANFeatureTemplateIds() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSDWANFeatureTemplateIdsRead,

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{

			"template_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			"template": {
				Type:     schema.TypeList,
				Computed: true,

				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"template_name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"template_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceSDWANFeatureTemplateIdsRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Read Method ")

	client := m.(*client.Client)
	templateType := d.Get("template_type").(string)

	cont, err := client.GetViaURL("/dataservice/template/feature")
	if err != nil {
		return err
	} else if cont == nil || !cont.Exists("data") {
		return fmt.Errorf("no feature template exists")
	}

	cont = cont.S("data")
	templatesList := make([]interface{}, 0)
	found := 0

	for _, template := range cont.Children() {

		templateMap := make(map[string]interface{})
		if template.Exists("templateType") && stripQuotes(template.S("templateType").String()) == templateType {

			if template.Exists("templateId") {
				templateMap["template_id"] = stripQuotes(template.S("templateId").String())
			}
			if template.Exists("templateName") {
				templateMap["template_name"] = stripQuotes(template.S("templateName").String())
			}

			found++
			templatesList = append(templatesList, templateMap)
		}
	}
	if found == 0 {
		return fmt.Errorf("no feature template of type %s exists", templateType)
	}

	d.Set("template", templatesList)

	d.SetId("template/feature")
	log.Println("[DEBUG] End of Read Method ")

	return nil
}

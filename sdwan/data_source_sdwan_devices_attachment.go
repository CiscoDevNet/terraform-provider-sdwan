package sdwan

import (
	"fmt"
	"log"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/CiscoDevNet/sdwan-go-client/container"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceSDWANDevicesAttachment() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSDWANDevicesAttachmentRead,

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{

			"device_template_id": {
				Type:     schema.TypeString,
				Required: true,
			},

			"template_type": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"devices_list": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"chassis_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
func datasourceSDWANDevicesAttachmentRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Read Method ")

	sdwanClient := m.(*client.Client)
	dtID := d.Get("device_template_id").(string)

	cont, err := getAttachedDevicesList(sdwanClient, dtID)
	if err != nil {
		if cont != nil {
			return fmt.Errorf("[ERROR] Data load failed in between")
		} else {
			return err
		}
	}

	log.Println("Container: ", cont)
	setDevicesList(d, cont.S("data"))

	dURL := fmt.Sprintf("dataservice/template/device/object/%s", dtID)
	cont, err = sdwanClient.GetViaURL(dURL)
	if err != nil {
		return err
	}

	templateType := stripQuotes(cont.S("configType").String())
	if templateType == "template" {
		templateType = "feature"
	} else {
		templateType = "cli"
	}

	d.Set("template_type", templateType)

	d.Set("device_template_id", dtID)

	d.SetId(dtID)

	log.Println("[DEBUG] End of Read Method ", d.Id())
	return nil
}

func getAttachedDevicesList(client *client.Client, id string) (*container.Container, error) {
	finalURL := fmt.Sprintf("/dataservice/template/device/config/attached/%s", id)

	cont, err := client.GetViaURL(finalURL)
	if err != nil {
		return nil, err
	}

	return cont, err
}

func setDevicesList(d *schema.ResourceData, cont *container.Container) *schema.ResourceData {

	log.Println("Data: ", cont)
	devicesList := make([]map[string]interface{}, 0, len(cont.Children()))

	for _, deviceCont := range cont.Children() {
		device := make(map[string]interface{})

		if deviceCont.Exists("uuid") {
			log.Println("Chassis Number: ", stripQuotes(deviceCont.S("uuid").String()))
			device["chassis_number"] = stripQuotes(deviceCont.S("uuid").String())
		}

		devicesList = append(devicesList, device)
	}

	d.Set("devices_list", devicesList)
	return d
}

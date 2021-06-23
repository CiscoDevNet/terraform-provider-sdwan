package sdwan

import (
	"fmt"
	"log"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceSDWANDeviceUUIDS() *schema.Resource {
	return &schema.Resource{
		Read:          datasourceSDWANDeviceUUIDSRead,
		SchemaVersion: 1,
		Description:   "Returns a list of SDWAN Devices present on vManage",
		Schema: map[string]*schema.Schema{

			"device": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"uuid": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_model": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func datasourceSDWANDeviceUUIDSRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Read Method")

	sdwanClient := m.(*client.Client)
	cont, err := sdwanClient.GetViaURL("dataservice/system/device/vedges")
	if err != nil {
		return err
	}

	if cont.Exists("data") {
		cont = cont.S("data")

		devicesList := make([]interface{}, 0)
		for _, deviceInfo := range cont.Children() {
			deviceMap := make(map[string]interface{})
			if deviceInfo.Exists("uuid") {
				deviceMap["uuid"] = stripQuotes(deviceInfo.S("uuid").String())
			}
			if deviceInfo.Exists("deviceModel") {
				deviceMap["device_model"] = stripQuotes(deviceInfo.S("deviceModel").String())
			}
			devicesList = append(devicesList, deviceMap)
		}
		d.Set("device", devicesList)
	} else {
		return fmt.Errorf("no device exists")
	}
	d.SetId("system/device/vedges")

	log.Println("[DEBUG] End of Read Method")

	return nil
}

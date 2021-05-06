package sdwan

import (
	"fmt"
	"testing"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/CiscoDevNet/sdwan-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestSDWANDeviceTemplate_Basic(t *testing.T) {
	var dt models.DeviceTemplate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSDWANDeviceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSDWANDeviceTemplateConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANDeviceTemplateExists("sdwan_device_template.test", &dt),
					testAccCheckSDWANDeviceTemplateAttributes("creation from terraform", &dt),
				),
			},
		},
	})

}

func TestSDWANDeviceTemplate_update(t *testing.T) {
	var dt models.DeviceTemplate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSDWANDeviceTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSDWANDeviceTemplateConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANDeviceTemplateExists("sdwan_device_template.test", &dt),
					testAccCheckSDWANDeviceTemplateAttributes("creation from terraform", &dt),
				),
			},
			{
				Config: testSDWANDeviceTemplateConfig_basic("updated from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANDeviceTemplateExists("sdwan_device_template.test", &dt),
					testAccCheckSDWANDeviceTemplateAttributes("updated from terraform", &dt),
				),
			},
		},
	})

}

func testSDWANDeviceTemplateConfig_basic(desc string) string {
	return fmt.Sprintf(`
	resource "sdwan_device_template" "test"{
		template_name = "example_name"
  		template_description="%s"
  		device_type="vedge-cloud"
  		device_role="sdwan-edge"
  		config_type="template"
  		factory_default=false
		connection_preference=false
		connection_preference_required=false
  		general_templates{
    		template_id="c8746871-cb8e-4ab7-a5f8-948c624f19ac"
    		template_type="system-vedge"
  		}
 		policy_id = ""
		feature_template_uid_range=["string"]
	}
	`, desc)
}

func testAccCheckSDWANDeviceTemplateExists(name string, dt *models.DeviceTemplate) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Device Template %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Device Template was set")
		}

		sdwanClient := testAccProvider.Meta().(*client.Client)

		cont, err := sdwanClient.GetViaURL(fmt.Sprintf("dataservice/template/device/object/%s", rs.Primary.ID))

		if err != nil {
			return err
		}

		dtGet := &models.DeviceTemplate{}

		dtGet.ConfigType = stripQuotes(cont.S("configType").String())
		dtGet.ConnectionPreference = cont.S("connectionPreference").Data().(bool)
		dtGet.ConnectionPreferenceRequired = cont.S("connectionPreferenceRequired").Data().(bool)
		dtGet.DeviceRole = stripQuotes(cont.S("deviceRole").String())
		dtGet.DeviceType = stripQuotes(cont.S("deviceType").String())
		dtGet.FactoryDefault = cont.S("factoryDefault").Data().(bool)

		fturList := make([]string, 0, 1)
		if len(cont.S("featureTemplateUidRange").Children()) > 0 {
			fturcontList := cont.S("featureTemplateUidRange")
			for i := 0; i < len(fturcontList.Data().([]interface{})); i++ {
				fturList = append(fturList, stripQuotes(fturcontList.Index(i).String()))
			}
		}

		dtGet.FeatureTemplateUidRange = fturList

		gtList := make([]*models.GeneralTemplate, 0, 1)
		if len(cont.S("generalTemplates").Children()) > 0 {
			gtcontList := cont.S("generalTemplates")
			for i := 0; i < len(gtcontList.Data().([]interface{})); i++ {
				gtMap := &models.GeneralTemplate{}

				gtMap.TemplateId = stripQuotes(gtcontList.Index(i).S("templateId").String())

				gtMap.TemplateType = stripQuotes(gtcontList.Index(i).S("templateType").String())

				if len(gtcontList.Index(i).S("subTemplates").Children()) > 0 {
					stcontList := gtcontList.Index(i).S("subTemplates")
					stList := make([]*models.Template, 0, 1)
					for j := 0; j < len(stcontList.Data().([]interface{})); j++ {
						stMap := &models.Template{}

						stMap.TemplateId = stripQuotes(stcontList.Index(j).S("templateId").String())

						stMap.TemplateType = stripQuotes(stcontList.Index(j).S("templateType").String())

						stList = append(stList, stMap)
					}

					gtMap.SubTemplates = stList
				}
				gtList = append(gtList, gtMap)
			}

		}

		dtGet.GeneralTemplates = gtList

		dtGet.PolicyId = stripQuotes(cont.S("policyId").String())
		dtGet.TemplateConfiguration = stripQuotes(cont.S("templateConfiguration").String())
		dtGet.TemplateDescription = stripQuotes(cont.S("templateDescription").String())
		dtGet.TemplateName = stripQuotes(cont.S("templateName").String())

		*dt = *dtGet
		return nil
	}
}

func testAccCheckSDWANDeviceTemplateDestroy(s *terraform.State) error {

	sdwanclient := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "sdwan_device_template" {
			dURL := fmt.Sprintf("/dataservice/template/device/%s", rs.Primary.ID)
			_, err := sdwanclient.GetViaURL(dURL)
			if err == nil {
				return fmt.Errorf("Device Template still exists")
			}
		}
	}
	return nil
}

func testAccCheckSDWANDeviceTemplateAttributes(desc string, dt *models.DeviceTemplate) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if desc != dt.TemplateDescription {
			return fmt.Errorf("Bad sdwan_device_template Templae Description %s", dt.TemplateDescription)
		}

		if "example_name" != dt.TemplateName {
			return fmt.Errorf("Bad sdwan_device_template Templae Name %s", dt.TemplateName)
		}

		if "vedge-cloud" != dt.DeviceType {
			return fmt.Errorf("Bad sdwan_device_template Device Type %s", dt.DeviceType)
		}

		if "template" != dt.ConfigType {
			return fmt.Errorf("Bad sdwan_device_template Config Type %s", dt.ConfigType)
		}

		if false != dt.FactoryDefault {
			return fmt.Errorf("Bad sdwan_device_template Factory Default value %v", dt.FactoryDefault)
		}

		if false != dt.ConnectionPreference {
			return fmt.Errorf("Bad sdwan_device_template connectionPreference value %v", dt.ConnectionPreference)
		}

		if false != dt.ConnectionPreferenceRequired {
			return fmt.Errorf("Bad sdwan_device_template connectionPreference value %v", dt.ConnectionPreferenceRequired)
		}

		gt := models.GeneralTemplate{}

		gt.TemplateId = "c8746871-cb8e-4ab7-a5f8-948c624f19ac"
		gt.TemplateType = "system-vedge"

		if gt.TemplateId != dt.GeneralTemplates[0].TemplateId {
			return fmt.Errorf("Bad sdwan_device_template genrealTemplates[0].TemplaId value %v", dt.GeneralTemplates[0].TemplateId)
		}

		if gt.TemplateType != dt.GeneralTemplates[0].TemplateType {
			return fmt.Errorf("Bad sdwan_device_template genrealTemplates[0].TemplaId value %v", dt.GeneralTemplates[0].TemplateType)
		}

		if "string" != dt.FeatureTemplateUidRange[0] {
			return fmt.Errorf("Bad sdwan_device_template featureTemplateUidRange[0] value %v", dt.FeatureTemplateUidRange[0])
		}

		return nil
	}
}

package sdwan

import (
	"fmt"
	"testing"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/CiscoDevNet/sdwan-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestSDWANSystemFeatureTemplate_Basic(t *testing.T) {
	var ft models.FeatureTemplate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSDWANSystemFeatureTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSDWANSystemFeatureTemplateConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANSystemFeatureTemplateExists("sdwan_system_feature_template.test", &ft),
					testAccCheckSDWANSystemFeatureTemplateAttributes("creation from terraform", &ft),
				),
			},
		},
	})

}

func TestSDWANSystemFeatureTemplate_update(t *testing.T) {
	var ft models.FeatureTemplate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSDWANSystemFeatureTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSDWANSystemFeatureTemplateConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANSystemFeatureTemplateExists("sdwan_system_feature_template.test", &ft),
					testAccCheckSDWANSystemFeatureTemplateAttributes("creation from terraform", &ft),
				),
			},
			{
				Config: testSDWANSystemFeatureTemplateConfig_basic("updated from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANSystemFeatureTemplateExists("sdwan_system_feature_template.test", &ft),
					testAccCheckSDWANSystemFeatureTemplateAttributes("updated from terraform", &ft),
				),
			},
		},
	})
}

func testSDWANSystemFeatureTemplateConfig_basic(desc string) string {
	fmt.Println("[DEBUG] Beginning of testSDWANSystemFeatureTemplateConfig_basic")
	return fmt.Sprintf(`
	resource "sdwan_system_feature_template" "test" {
  
		template_name = "sample"
		template_description = "%s"
		device_type = ["vedge-1000"]
		template_min_version = "15.0.0"
		template_definition {
			system_basic {  
				site_id = 5
				system_ip = "198.12.23.178"
				timezone = "UTC"
			}
		}
		factory_default = false
	}
	`, desc)
}

func testAccCheckSDWANSystemFeatureTemplateExists(name string, ft *models.FeatureTemplate) resource.TestCheckFunc {
	fmt.Println("[DEBUG] Beginning of testAccCheckSDWANSystemFeatureTemplateExists")
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("System feture template %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No System feature template was set")
		}

		sdwanClient := testAccProvider.Meta().(*client.Client)

		cont, err := sdwanClient.GetViaURL(fmt.Sprintf("dataservice/template/feature/object/%s", rs.Primary.ID))

		if err != nil {
			return err
		}

		ftGet := &models.FeatureTemplate{}

		devList := make([]string, 0, 1)

		if len(cont.S("deviceType").Children()) > 0 {
			devcontList := cont.S("deviceType")
			for i := 0; i < len(devcontList.Data().([]interface{})); i++ {
				devList = append(devList, stripQuotes(devcontList.Index(i).String()))
			}
		}

		contftdefinition := cont.S("templateDefinition").Data().(map[string]interface{})

		ftGet.DeviceType = devList
		ftGet.FactoryDefault = cont.S("factoryDefault").Data().(bool)
		ftGet.TemplateDefinition = contftdefinition
		ftGet.TemplateDescription = stripQuotes(cont.S("templateDescription").String())
		ftGet.TemplateMinVersion = stripQuotes(cont.S("templateMinVersion").String())
		ftGet.TemplateName = stripQuotes(cont.S("templateName").String())

		*ft = *ftGet
		fmt.Println("[DEBUG] End of testAccCheckSDWANSystemFeatureTemplateExists")
		return nil
	}
}

func testAccCheckSDWANSystemFeatureTemplateDestroy(s *terraform.State) error {
	fmt.Println("[DEBUG] Beginning of testAccCheckSDWANSystemFeatureTemplateDestroy")
	sdwanclient := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "sdwan_system_feature_template" {
			dURL := fmt.Sprintf("/dataservice/template/feature/object/%s", rs.Primary.ID)
			_, err := sdwanclient.GetViaURL(dURL)
			if err == nil {
				return fmt.Errorf("Feature template still exist")
			}
		}
	}
	fmt.Println("[DEBUG] End of testAccCheckSDWANSystemFeatureTemplateDestroy")
	return nil
}

func testAccCheckSDWANSystemFeatureTemplateAttributes(desc string, ft *models.FeatureTemplate) resource.TestCheckFunc {
	fmt.Println("[DEBUG] Beginning of testAccCheckSDWANSystemFeatureTemplateAttributes")
	return func(s *terraform.State) error {
		if desc != ft.TemplateDescription {
			return fmt.Errorf("Bad sdwan_system_feature_template Template Description %s", ft.TemplateDescription)
		}

		if "sample" != ft.TemplateName {
			return fmt.Errorf("Bad sdwan_system_feature_template Template Name %s", ft.TemplateName)
		}

		if "15.0.0" != ft.TemplateMinVersion {
			return fmt.Errorf("Bad sdwan_system_feature_template Template Min Version %s", ft.TemplateMinVersion)
		}

		if "vedge-1000" != ft.DeviceType[0] {
			return fmt.Errorf("Bad sdwan_system_feature_template Device Type %s", ft.DeviceType[0])
		}

		if false != ft.FactoryDefault {
			return fmt.Errorf("Bad sdwan_system_feature_template Factory Default %v", ft.FactoryDefault)
		}

		tempDef := map[string]interface{}{
			"system_basic": map[string]interface{}{
				"site_id":   5,
				"system_ip": "198.12.23.178",
				"timezone":  "UTC",
			},
		}

		if tempDef["system_basic"].(map[string]interface{})["site_id"].(int) != int(ft.TemplateDefinition["site-id"].(map[string]interface{})["vipValue"].(float64)) {
			return fmt.Errorf("Bad sdwan_system_feature_template site id %v", ft.TemplateDefinition["site-id"].(map[string]interface{})["vipValue"])
		}

		if tempDef["system_basic"].(map[string]interface{})["system_ip"].(string) != ft.TemplateDefinition["system-ip"].(map[string]interface{})["vipValue"].(string) {
			return fmt.Errorf("Bad sdwan_system_feature_template system ip %v", ft.TemplateDefinition["system-ip"].(map[string]interface{})["vipValue"])
		}

		if tempDef["system_basic"].(map[string]interface{})["timezone"].(string) != ft.TemplateDefinition["clock"].(map[string]interface{})["timezone"].(map[string]interface{})["vipValue"].(string) {
			return fmt.Errorf("Bad sdwan_system_feature_template timezone %v", ft.TemplateDefinition["timezone"].(map[string]interface{})["vipValue"])
		}

		return nil
	}
}

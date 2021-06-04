package sdwan

import (
	"fmt"
	"testing"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/CiscoDevNet/sdwan-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestSDWANVPNInterfaceFeatureTemplate_Basic(t *testing.T) {
	var ft models.FeatureTemplate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSDWANVPNInterfaceFeatureTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSDWANVPNInterfaceFeatureTemplateConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANVPNInterfaceFeatureTemplateExists("sdwan_vpn_interface_feature_template.test", &ft),
					testAccCheckSDWANVPNInterfaceFeatureTemplateAttributes("creation from terraform", &ft),
				),
			},
		},
	})
}

func TestSDWANVPNInterfaceFeatureTemplate_update(t *testing.T) {
	var ft models.FeatureTemplate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSDWANVPNInterfaceFeatureTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSDWANVPNInterfaceFeatureTemplateConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANVPNInterfaceFeatureTemplateExists("sdwan_vpn_interface_feature_template.test", &ft),
					testAccCheckSDWANVPNInterfaceFeatureTemplateAttributes("creation from terraform", &ft),
				),
			},
			{
				Config: testSDWANVPNInterfaceFeatureTemplateConfig_basic("updated from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANVPNInterfaceFeatureTemplateExists("sdwan_vpn_interface_feature_template.test", &ft),
					testAccCheckSDWANVPNInterfaceFeatureTemplateAttributes("updated from terraform", &ft),
				),
			},
		},
	})
}

func testSDWANVPNInterfaceFeatureTemplateConfig_basic(desc string) string {
	return fmt.Sprintf(`
	resource "sdwan_vpn_interface_feature_template" "test" {

		template_name = "example_test_10"
		template_description = "%s"
		device_type = ["vedge-100"]
		template_type = "vpn-vedge-interface"
		template_min_version = "15.0.0"
  		factory_default = false
		template_definition{
			vpn_interface_basic{
				description = "for testing"
			}
		}
	}
	`, desc)
}

func testAccCheckSDWANVPNInterfaceFeatureTemplateExists(name string, ft *models.FeatureTemplate) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("VPN interface feature template %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VPN inerface feature template was set")
		}

		sdwanClient := testAccProvider.Meta().(*client.Client)

		dURL := fmt.Sprintf("dataservice/template/feature/object/%s", rs.Primary.ID)

		cont, err := sdwanClient.GetViaURL(dURL)

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

		contftDefination := cont.S("templateDefinition").Data().(map[string]interface{})

		ftGet.DeviceType = devList
		ftGet.FactoryDefault = cont.S("factoryDefault").Data().(bool)
		ftGet.TemplateDefinition = contftDefination
		ftGet.TemplateDescription = stripQuotes(cont.S("templateDescription").String())
		ftGet.TemplateMinVersion = stripQuotes(cont.S("templateMinVersion").String())
		ftGet.TemplateName = stripQuotes(cont.S("templateName").String())

		*ft = *ftGet

		return nil
	}
}

func testAccCheckSDWANVPNInterfaceFeatureTemplateDestroy(s *terraform.State) error {
	sdwanClient := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "sdwan_vpn_interface_feature_template" {
			dURL := fmt.Sprintf("/dataservice/template/feature/object/%s", rs.Primary.ID)
			_, err := sdwanClient.GetViaURL(dURL)
			if err == nil {
				return fmt.Errorf("Feature template still exists")
			}
		}
	}
	return nil
}

func testAccCheckSDWANVPNInterfaceFeatureTemplateAttributes(desc string, ft *models.FeatureTemplate) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if desc != ft.TemplateDescription {
			return fmt.Errorf("Bad sdwan_vpn_interface_feature_template Template Description %s", ft.TemplateDescription)
		}

		if "example_test_10" != ft.TemplateName {
			return fmt.Errorf("Bad sdwan_vpn_interface_feature_template Template Name %s", ft.TemplateName)
		}

		if "15.0.0" != ft.TemplateMinVersion {
			return fmt.Errorf("Bad sdwan_vpn_interface_feature_template Template Min Version %s", ft.TemplateMinVersion)
		}

		if "vedge-100" != ft.DeviceType[0] {
			return fmt.Errorf("Bad sdwan_vpn_interface_feature_template Device Type %s", ft.DeviceType[0])
		}

		if false != ft.FactoryDefault {
			return fmt.Errorf("Bad sdwan_vpn_interface_feature_template Factory Default %v", ft.FactoryDefault)
		}

		tempDef := map[string]interface{}{
			"vpn_interface_basic": map[string]interface{}{
				"description": "for testing",
			},
		}

		if tempDef["vpn_interface_basic"].(map[string]interface{})["description"].(string) != ft.TemplateDefinition["description"].(map[string]interface{})["vipValue"].(string) {
			return fmt.Errorf("Bad sdwan_vpn_interface_featre_template inerface_name %s", ft.TemplateDefinition["description"].(map[string]interface{})["vipValue"].(string))
		}
		return nil
	}
}

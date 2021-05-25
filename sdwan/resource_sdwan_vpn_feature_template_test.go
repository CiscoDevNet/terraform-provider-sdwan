package sdwan

import (
	"fmt"
	"log"
	"testing"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/CiscoDevNet/sdwan-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestSDWANVPNFeatureTemplate_Basic(t *testing.T) {
	var ft models.FeatureTemplate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSDWANVPNFeatureTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSDWANVPNFeatureTemplateConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANVPNFeatureTemplateExists("sdwan_vpn_feature_template.test", &ft),
					testAccCheckSDWANVPNFeatureTemplateAttributes("creation from terraform", &ft),
				),
			},
		},
	})
}

func TestSDWANVPNFeatureTemplate_update(t *testing.T) {
	var ft models.FeatureTemplate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSDWANVPNFeatureTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSDWANVPNFeatureTemplateConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANVPNFeatureTemplateExists("sdwan_vpn_feature_template.test", &ft),
					testAccCheckSDWANVPNFeatureTemplateAttributes("creation from terraform", &ft),
				),
			},
			{
				Config: testSDWANVPNFeatureTemplateConfig_basic("updated from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANVPNFeatureTemplateExists("sdwan_vpn_feature_template.test", &ft),
					testAccCheckSDWANVPNFeatureTemplateAttributes("updated from terraform", &ft),
				),
			},
		},
	})
}

func testSDWANVPNFeatureTemplateConfig_basic(desc string) string {
	return fmt.Sprintf(`
	resource "sdwan_vpn_feature_template" "test" {
		template_name = "example"
		template_type = "vpn-vedge"
		factory_default = false
		device_type = [ "vedge-cloud" ]
		template_description = "%s"
		template_min_version = "15.0.0"
		template_definition {
		  vpn_basic {
			vpn_id = 0
		  }
		}
	  }
	`, desc)
}

func testAccCheckSDWANVPNFeatureTemplateExists(name string, ft *models.FeatureTemplate) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("VPN feature template %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No VPN feature template was set")
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
		ftGet.TemplateType = stripQuotes(cont.S("templateType").String())

		*ft = *ftGet
		return nil
	}
}

func testAccCheckSDWANVPNFeatureTemplateDestroy(s *terraform.State) error {
	sdwanclient := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "sdwan_vpn_feature_template" {
			dURL := fmt.Sprintf("/dataservice/template/feature/object/%s", rs.Primary.ID)
			_, err := sdwanclient.GetViaURL(dURL)
			if err == nil {
				return fmt.Errorf("Feature Template still exists")
			}
		}
	}
	return nil
}

func testAccCheckSDWANVPNFeatureTemplateAttributes(desc string, ft *models.FeatureTemplate) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if desc != ft.TemplateDescription {
			return fmt.Errorf("Bad sdwan_vpn_feature_template Template Description %s", ft.TemplateDescription)
		}

		if "example" != ft.TemplateName {
			return fmt.Errorf("Bad sdwan_vpn_feature_template Template Name %s", ft.TemplateName)
		}

		if "15.0.0" != ft.TemplateMinVersion {
			return fmt.Errorf("Bad sdwan_vpn_feature_template Template Min Version %s", ft.TemplateMinVersion)
		}

		if "vedge-cloud" != ft.DeviceType[0] {
			return fmt.Errorf("Bad sdwan_vpn_feature_template Device Type %s", ft.DeviceType[0])
		}

		if false != ft.FactoryDefault {
			return fmt.Errorf("Bad sdwan_vpn_feature_template Factory Default %v", ft.FactoryDefault)
		}

		if "vpn-vedge" != ft.TemplateType {
			log.Printf(ft.TemplateType)
			return fmt.Errorf("Bad sdwan_vpn_feature_template Template Type %s", ft.TemplateType)
		}

		tempDef := map[string]interface{}{
			"vpn_basic": map[string]interface{}{
				"vpn_id": 0,
			},
		}

		if tempDef["vpn_basic"].(map[string]interface{})["vpn_id"].(int) != int(ft.TemplateDefinition["vpn-id"].(map[string]interface{})["vipValue"].(float64)) {
			return fmt.Errorf("Bad sdwan_vpn_feature_template vpn_id %v", int(ft.TemplateDefinition["vpn-id"].(map[string]interface{})["vipValue"].(float64)))
		}

		return nil
	}
}
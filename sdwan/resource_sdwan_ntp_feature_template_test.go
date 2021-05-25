package sdwan

import (
	"fmt"
	"testing"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/CiscoDevNet/sdwan-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestSDWANNtpFeatureTemplate_Basic(t *testing.T) {
	var ft models.FeatureTemplate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSDWANNtpFeatureTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSDWANNtpFeatureTemplateConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANNtpFeatureTemplateExists("sdwan_ntp_feature_template.test", &ft),
					testAccCheckSDWANNtpFeatureTemplateAttributes("creation from terraform", &ft),
				),
			},
		},
	})

}

func TestSDWANNtpFeatureTemplate_update(t *testing.T) {
	var ft models.FeatureTemplate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSDWANNtpFeatureTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSDWANNtpFeatureTemplateConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANNtpFeatureTemplateExists("sdwan_ntp_feature_template.test", &ft),
					testAccCheckSDWANNtpFeatureTemplateAttributes("creation from terraform", &ft),
				),
			},
			{
				Config: testSDWANNtpFeatureTemplateConfig_basic("updated from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANNtpFeatureTemplateExists("sdwan_ntp_feature_template.test", &ft),
					testAccCheckSDWANNtpFeatureTemplateAttributes("updated from terraform", &ft),
				),
			},
		},
	})
}

func testSDWANNtpFeatureTemplateConfig_basic(desc string) string {
	fmt.Println("[DEBUG] Beginning of testSDWANNtpFeatureTemplateConfig_basic")
	return fmt.Sprintf(`
	resource "sdwan_ntp_feature_template" "test" {
		template_name = "sample"
		template_description = "%s"
		device_type = ["vedge-1000"]
		template_type = "ntp"
		template_min_version = "15.0.0"
		factory_default = false
		template_definition {}
	}
	`, desc)
}

func testAccCheckSDWANNtpFeatureTemplateExists(name string, ft *models.FeatureTemplate) resource.TestCheckFunc {
	fmt.Println("[DEBUG] Beginning of testAccCheckSDWANNtpFeatureTemplateExists")
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Ntp feture template %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Ntp feature template was set")
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
		fmt.Println("[DEBUG] End of testAccCheckSDWANNtpFeatureTemplateExists")
		return nil
	}
}

func testAccCheckSDWANNtpFeatureTemplateDestroy(s *terraform.State) error {
	fmt.Println("[DEBUG] Beginning of testAccCheckSDWANNtpFeatureTemplateDestroy")
	sdwanclient := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "sdwan_ntp_feature_template" {
			dURL := fmt.Sprintf("/dataservice/template/feature/object/%s", rs.Primary.ID)
			_, err := sdwanclient.GetViaURL(dURL)
			if err == nil {
				return fmt.Errorf("Feature template still exist")
			}
		}
	}
	fmt.Println("[DEBUG] End of testAccCheckSDWANNtpFeatureTemplateDestroy")
	return nil
}

func testAccCheckSDWANNtpFeatureTemplateAttributes(desc string, ft *models.FeatureTemplate) resource.TestCheckFunc {
	fmt.Println("[DEBUG] Beginning of testAccCheckSDWANNtpFeatureTemplateAttributes")
	return func(s *terraform.State) error {
		if desc != ft.TemplateDescription {
			return fmt.Errorf("Bad sdwan_Ntp_feature_template Template Description %s", ft.TemplateDescription)
		}

		if "sample" != ft.TemplateName {
			return fmt.Errorf("Bad sdwan_Ntp_feature_template Template Name %s", ft.TemplateName)
		}

		if "15.0.0" != ft.TemplateMinVersion {
			return fmt.Errorf("Bad sdwan_Ntp_feature_template Template Min Version %s", ft.TemplateMinVersion)
		}

		if "vedge-1000" != ft.DeviceType[0] {
			return fmt.Errorf("Bad sdwan_Ntp_feature_template Device Type %s", ft.DeviceType[0])
		}

		if "ntp" != ft.TemplateType {
			return fmt.Errorf("Bad sdwan_Ntp_feature_template Template Type %s", ft.TemplateType)
		}

		if false != ft.FactoryDefault {
			return fmt.Errorf("Bad sdwan_Ntp_feature_template Factory Default %v", ft.FactoryDefault)
		}

		return nil
	}
}
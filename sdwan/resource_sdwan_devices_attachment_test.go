package sdwan

import (
	"fmt"
	"testing"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestSDWANDeviceAttachment_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSDWANDeviceAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSDWANDeviceAttachmentConfig_basic(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANDeviceAttachmentExists("sdwan_devices_attachment.test"),
					testAccCheckSDWANDeviceAttachmentAttributes(1),
				),
			},
		},
	})
}

func TestSDWANDeviceAttachement_Update1(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSDWANDeviceAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSDWANDeviceAttachmentConfig_basic(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANDeviceAttachmentExists("sdwan_devices_attachment.test"),
					testAccCheckSDWANDeviceAttachmentAttributes(1),
				),
			},
			{
				Config: testSDWANDeviceAttachmentConfig_basic("devices_list {\nchassis_number = \"CSR-2AD10A1D-325E-48CC-9D80-A9B5B7BBF7CB\"\n}"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANDeviceAttachmentExists("sdwan_devices_attachment.test"),
					testAccCheckSDWANDeviceAttachmentAttributes(2),
				),
			},
		},
	})
}

func TestSDWANDeviceAttachement_Update2(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSDWANDeviceAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testSDWANDeviceAttachmentConfig_basic("devices_list {\nchassis_number = \"CSR-2AD10A1D-325E-48CC-9D80-A9B5B7BBF7CB\"\n}"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANDeviceAttachmentExists("sdwan_devices_attachment.test"),
					testAccCheckSDWANDeviceAttachmentAttributes(2),
				),
			},
			{
				Config: testSDWANDeviceAttachmentConfig_basic(""),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSDWANDeviceAttachmentExists("sdwan_devices_attachment.test"),
					testAccCheckSDWANDeviceAttachmentAttributes(1),
				),
			},
		},
	})
}

func testSDWANDeviceAttachmentConfig_basic(device string) string {

	return fmt.Sprintf(`	
	resource "sdwan_devices_attachment" "test" {
		timeout = 600
		device_template_id = "5477b564-db5e-4195-8d76-00b6d471a8b6"
		template_type = "feature"  
		file = "../examples/devices_attachment/Template1.csv"
		devices_list {
			chassis_number = "CSR-330E65C0-8D52-41B2-940A-398119726CCE"
		}
		%s
	}
	`, device)
}

var DevList []map[string]interface{}

func testAccCheckSDWANDeviceAttachmentExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Attachement %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No attchment found")
		}

		sdwanClient := testAccProvider.Meta().(*client.Client)

		StatusResult, err := getFinalStatus(sdwanClient, rs.Primary.ID)

		if err != nil {
			return err
		}

		for id, exists := range StatusResult.IDs {
			if exists {
				device := make(map[string]interface{})

				device["chassis_number"] = id

				if StatusResult.Data[id].Status == "Failure" || StatusResult.Data[id].Status == "In progress" {
					device["attach"] = false
				} else {
					device["attach"] = true
				}

				device["status"] = StatusResult.Data[id].Status

				device["activity"] = StatusResult.Data[id].Activity

				DevList = append(DevList, device)

			}
		}
		return nil
	}
}

func testAccCheckSDWANDeviceAttachmentDestroy(s *terraform.State) error {
	sdwanclient := testAccProvider.Meta().(*client.Client)
	for _, rs := range s.RootModule().Resources {

		dURL := fmt.Sprintf("/dataservice/device/action/status/%s", rs.Primary.ID)
		cont, err := sdwanclient.GetViaURL(dURL)
		if err != nil {
			return err
		}

		uuid := stripQuotes(cont.S("uuid").String())
		cont, err = sdwanclient.GetViaURL("/dataservice/system/device/vedges")
		if err != nil {
			return err
		}

		for _, device := range cont.S("data").Children() {
			if stripQuotes(device.S("uuid").String()) == uuid {
				if stripQuotes(device.S("configOperationMode").String()) == "vmanage" {
					return fmt.Errorf("attachment is still present")
				}
			}
		}
	}

	return nil
}

func testAccCheckSDWANDeviceAttachmentAttributes(required int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for num := 0; num < required; num++ {

			if DevList[num]["attach"] != true {
				return fmt.Errorf("device should have been attached")
			}
		}
		return nil
	}
}

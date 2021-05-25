package sdwan

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/CiscoDevNet/sdwan-go-client/container"
	"github.com/CiscoDevNet/sdwan-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSDWANDevicesAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceSDWANDevicesAttachmentCreate,
		Update: resourceSDWANDevicesAttachmentUpdate,
		Read:   resourceSDWANDevicesAttachmentRead,
		Delete: resourceSDWANDevicesAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  600,
			},

			"device_template_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"template_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice(
					[]string{
						"feature",
						"cli",
					}, false),
			},

			"file": {
				Type:     schema.TypeString,
				Required: true,
			},

			"devices_list": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"chassis_number": {
							Type:     schema.TypeString,
							Required: true,
						},

						"attach": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},

						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"activity": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
				Set: func(v interface{}) int {
					var buf bytes.Buffer
					m := v.(map[string]interface{})
					buf.WriteString(fmt.Sprintf("%s-", m["chassis_number"].(string)))
					return hashString(buf.String())
				},
			},
		},
	}
}

func resourceSDWANDevicesAttachmentCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Create method")

	start := time.Now()
	sdwanClient := m.(*client.Client)

	if val, ok := d.GetOk("timeout"); ok {
		OperationalTime = time.Duration(val.(int)-2) * time.Second
	}

	devicesData := d.Get("devices_list").(*schema.Set).List()
	var dList []string
	cnt := 0
	for _, deviceObject := range devicesData {
		uuid := deviceObject.(map[string]interface{})["chassis_number"].(string)

		if doesDeviceExists(sdwanClient, uuid) {
			dList = append(dList, uuid)
			cnt++
		} else {
			return fmt.Errorf("[ERROR] Device with Chassis Number: %s does not exist", uuid)
		}
	}

	err := attachDevices(d, sdwanClient, dList, cnt, OperationalTime)
	if err != nil {
		dur := time.Since(start)
		if dur > OperationalTime {
			log.Println("[ERROR] Process timed out, Check from vManage. \nReason:", err)
		} else {
			return err
		}
	}

	log.Println("[DEBUG] End of Create Method ", d.Id())
	return resourceSDWANDevicesAttachmentRead(d, m)
}

func resourceSDWANDevicesAttachmentUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Update method: ", d.Id())

	start := time.Now()

	sdwanClient := m.(*client.Client)

	if val, ok := d.GetOk("timeout"); ok {
		OperationalTime = time.Duration(val.(int)-2) * time.Second
	}

	if d.HasChange("devices_list") && !d.IsNewResource() {
		o, n := d.GetChange("devices_list")

		dList := []*models.Device{}
		devicesToDetach := 0

		for _, old := range o.(*schema.Set).List() {
			if !(n.(*schema.Set).Contains(old)) {

				devicesToDetach++

				deviceModel := models.Device{}

				deviceModel.DeviceID = old.(map[string]interface{})["chassis_number"].(string)

				dList = append(dList, &deviceModel)
			}
		}
		if devicesToDetach > 0 {
			err := detachDevices(d, sdwanClient, dList, OperationalTime)
			if err != nil {
				dur := time.Since(start)
				if dur > OperationalTime {
					return fmt.Errorf("[ERROR] Process timed out, %v", err)
				} else {
					return err
				}
			}
		}

		dur := time.Since(start)
		log.Println("Duration for attaching: ", dur)
		OperationalTime -= dur

		var devList []string
		cnt := 0
		for _, new := range n.(*schema.Set).List() {
			if !(o.(*schema.Set).Contains(new)) {
				uuid := new.(map[string]interface{})["chassis_number"].(string)
				if doesDeviceExists(sdwanClient, uuid) {
					devList = append(devList, uuid)
					cnt++
				} else {
					return fmt.Errorf("[ERROR] Device with Chassis Number: %s does not exist", uuid)
				}

			}
		}
		if cnt > 0 {
			err := attachDevices(d, sdwanClient, devList, cnt, OperationalTime)
			if err != nil {
				dur := time.Since(start)
				if dur > OperationalTime {
					return fmt.Errorf("[ERROR] Process timed out, %v", err)
				} else {
					return err
				}
			}
		}
	}

	log.Println("[DEBUG] End of Update Method ", d.Id())
	return resourceSDWANDevicesAttachmentRead(d, m)
}

func resourceSDWANDevicesAttachmentRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Read method: ", d.Id())

	sdwanClient := m.(*client.Client)
	processIds := d.Id()
	err := getFinalStatus(sdwanClient, processIds)
	if err != nil {
		return err
	}

	log.Println("Status: ", StatusResult)

	devList := make([]map[string]interface{}, 0, StatusResult.Summary.Total)
	for id, exists := range StatusResult.IDs {

		if exists {
			log.Println("Data: ", StatusResult.Data[id])
			device := make(map[string]interface{})

			device["chassis_number"] = id

			if StatusResult.Data[id].Status == "Failure" || StatusResult.Data[id].Status == "In progress" {
				device["attach"] = false
			} else {
				device["attach"] = true
			}

			device["status"] = StatusResult.Data[id].Status

			device["activity"] = StatusResult.Data[id].Activity

			devList = append(devList, device)
		}
	}

	d.Set("devices_list", devList)

	log.Println("[DEBUG] End of Read Method ", d.Id())
	return nil
}

func resourceSDWANDevicesAttachmentDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Delete method: ", d.Id())
	start := time.Now()

	sdwanClient := m.(*client.Client)

	if val, ok := d.GetOk("timeout"); ok {
		OperationalTime = time.Duration(val.(int)-2) * time.Second
	}

	devicesList := []*models.Device{}
	for _, deviceObject := range d.Get("devices_list").(*schema.Set).List() {
		deviceModel := &models.Device{
			DeviceID: deviceObject.(map[string]interface{})["chassis_number"].(string),
		}
		devicesList = append(devicesList, deviceModel)
	}

	err := detachDevices(d, sdwanClient, devicesList, OperationalTime)
	if err != nil {
		dur := time.Since(start)
		if dur > OperationalTime {
			return fmt.Errorf("[ERROR] Process timed out, %v", err)
		} else {
			return err
		}
	}

	d.SetId("")
	log.Println("[DEBUG] End of Delete Method ")
	return nil
}

var OperationalTime time.Duration
var StatusResult *models.StatusData
var TaskIDs []string

func getFinalStatus(client *client.Client, ids string) error {
	pIDs := strings.Split(ids, ":")

	StatusResult = &models.StatusData{
		Summary: &models.SummaryObject{
			Total: 0,
		},
	}

	for _, ptID := range pIDs {

		status, err := getStatusByID(client, ptID)
		if err != nil {
			return fmt.Errorf("[ERROR] Some error occured :%v", err)
		}

		if status == nil || status.Summary == nil {
			return fmt.Errorf("[ERROR] Status load failed in between")
		}

		if strings.HasPrefix(ptID, "push_feature_template_configuration") || strings.HasPrefix(ptID, "push_file_template_configuration") {

			// Update the StateResult after attaching devices
			for id, exists := range status.IDs {
				if exists {
					StatusResult.Summary.Total++

					if StatusResult.IDs == nil {
						StatusResult.IDs = make(map[string]bool)
					}

					StatusResult.IDs[id] = true

					if StatusResult.Data == nil {
						StatusResult.Data = make(map[string]*models.DataObject)
					}
					StatusResult.Data[id] = status.Data[id]
				}
			}
		} else {

			// Update the StateResult after detaching devices
			for id, exists := range status.IDs {
				if exists {
					StatusResult.Summary.Total--

					if StatusResult.IDs == nil {
						return fmt.Errorf("[ERROR] Cannot detach a device without attaching it")
					}

					StatusResult.IDs[id] = false
				}
			}
		}
	}
	return nil
}

func attachDevices(d *schema.ResourceData, client *client.Client, dList []string, cnt int, timeout time.Duration) error {

	log.Println("[DEBUG] Beginning Devices Attachment")

	cont1, err := generateDeviceConfig(d, client, dList)
	if err != nil {
		return err
	}

	dtID := d.Get("device_template_id").(string)

	templateType := d.Get("template_type").(string)

	devObjects := cont1.S("data").Data().([]interface{})

	if matchedCnt, err := strconv.Atoi(cont1.S("countSummary", "matchCount").String()); err == nil {
		if matchedCnt < cnt {
			return fmt.Errorf("[ERROR] Only %d Device values found. Need %d Device values",
				matchedCnt, cnt)
		}
	}

	configOpts := &models.DeviceConfig{
		DevicesCount: cnt,
		DeviceTemplateList: &models.DeviceTemplateList{
			TemplateID: dtID,
			Device:     devObjects,
		},
	}

	if templateType == "feature" {

		cont1, err = client.Save("/dataservice/template/device/config/attachfeature", configOpts)
		if err != nil {
			return err
		}
	} else {
		cont1, err = client.Save("/dataservice/template/device/config/attachcli", configOpts)
		if err != nil {
			return err
		}
	}

	ptID := stripQuotes(cont1.S("id").String())
	var processIds string
	if d.Id() == "" {
		processIds = ptID
	} else {
		processIds = d.Id() + ":" + ptID
	}

	d.Set("device_template_id", dtID)
	d.SetId(processIds)
	log.Println("[DEBUG] End of Devices Attachment")

	log.Println("[DEBUG] Waiting for device(s) to be attached, ", d.Id())

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"in_progress"},
		Refresh:    AttachmentStateRefreshFunc(client, ptID, []string{"Failure"}),
		Target:     []string{"done"},
		Timeout:    timeout,
		Delay:      2 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return err
	}

	log.Println("[DEBUG] End of attachment")
	return nil
}

func detachDevices(d *schema.ResourceData, client *client.Client, dList []*models.Device, timeout time.Duration) error {

	log.Println("[DEBUG] Beginning to detach the Device(s) from Template")
	detachOpts := &models.DeviceDetach{
		DeviceType: "vedge",
		Devices:    dList,
	}

	cont1, err := client.Save("/dataservice/template/config/device/mode/cli", detachOpts)
	if err != nil {
		return err
	}

	ptID := stripQuotes(cont1.S("id").String())

	var processIds string
	if d.Id() == "" {
		processIds = ptID
	} else {
		processIds = d.Id() + ":" + ptID
	}
	d.SetId(processIds)

	log.Println("[DEBUG] Waiting for Device(s) to be detached from Device Template: ", d.Id())

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"in_progress"},
		Refresh:    AttachmentStateRefreshFunc(client, ptID, []string{"Failure"}),
		Target:     []string{"done"},
		Timeout:    timeout,
		Delay:      2 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		cancelDetachment(client, ptID)
		return err
	}

	log.Println("[DEBUG] End of Detachment")
	return nil
}

func AttachmentStateRefreshFunc(client *client.Client, statusID string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		status, err := getStatusByID(client, statusID)

		if err != nil {
			log.Printf("Error on Attachment Status Refresh: %s", err)
			return nil, "", err
		}

		if status == nil || status.Summary == nil {
			return nil, "", nil
		}

		state := status.Summary.Status

		return status, state, nil
	}
}

func getStatusByID(client *client.Client, id string) (*models.StatusData, error) {

	statusURL := fmt.Sprintf("/dataservice/device/action/status/%s", id)
	cont, err := client.GetViaURL(statusURL)
	if err != nil {
		return nil, err
	}

	total, _ := strconv.Atoi(stripQuotes(cont.S("summary", "total").String()))

	summary := &models.SummaryObject{
		Total:  total,
		Status: stripQuotes(cont.S("summary", "status").String()),
		Count:  cont.S("summary", "count").Data().(map[string]interface{}),
	}

	ids := make(map[string]bool)

	results := make(map[string]*models.DataObject)
	for _, dataObject := range cont.S("data").Children() {

		uuid := stripQuotes(dataObject.S("uuid").String())
		ids[uuid] = true

		results[uuid] = &models.DataObject{
			Status:          stripQuotes(dataObject.S("status").String()),
			CurrentActivity: stripQuotes(dataObject.S("currentActivity").String()),
			Activity:        interfaceToStrList(dataObject.S("activity").Data()),
			ActionConfig:    stripQuotes(dataObject.S("actionConfig").String()),
		}
	}

	status := &models.StatusData{
		Summary: summary,
		IDs:     ids,
		Data:    results,
	}

	return status, nil
}

func generateDeviceConfig(d *schema.ResourceData, client *client.Client, dList []string) (*container.Container, error) {

	log.Println("[DEBUG] Beginning Config Generation")
	dtID := d.Get("device_template_id").(string)

	attachDataOpts := &models.DeviceAttachData{
		TemplateID: dtID,
		DeviceIDs:  dList,
	}

	pathofFile := d.Get("file").(string)

	attachOpts := &models.DeviceAttach{
		File: pathofFile,
		Data: attachDataOpts,
	}

	cont1, err := client.SaveWithFile("/dataservice/template/device/config/process/input/file", attachOpts)
	if err != nil {
		return nil, err
	}

	log.Println("[DEBUG] End of Config Generation")
	return cont1, nil
}

func cancelDetachment(client *client.Client, id string) {
	dURL := fmt.Sprintf("/dataservice/device/action/status/cancel/%s", id)
	client.Save(dURL, nil)
}

func doesDeviceExists(client *client.Client, id string) bool {

	cont, err := client.GetViaURL("dataservice/system/device/vedges")
	if err != nil {
		log.Println("vEdges Devices error:", err)
	}

	found := false

	if cont.Exists("data", "*", "chasisNumber") {
		result := cont.S("data", "*", "chasisNumber").Data()

		for _, device := range result.([]interface{}) {
			if device.(string) == id {
				found = true
			}
		}
	}

	return found
}

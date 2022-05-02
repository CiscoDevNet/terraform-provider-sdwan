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
					buf.WriteString(fmt.Sprintf("%v-", m["attach"].(bool)))
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

	var OperationalTime time.Duration
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

	var OperationalTime time.Duration
	if val, ok := d.GetOk("timeout"); ok {
		OperationalTime = time.Duration(val.(int)-2) * time.Second
	}

	if d.HasChange("devices_list") && !d.IsNewResource() {
		o, n := d.GetChange("devices_list")

		dList := []*models.Device{}
		devicesToDetach := 0

		for _, old := range o.(*schema.Set).List() {

			// Detach not found devices
			if !(n.(*schema.Set).Contains(old)) {

				// Only detach if it is attached
				if old.(map[string]interface{})["attach"].(bool) {
					devicesToDetach++

					deviceModel := models.Device{}

					deviceModel.DeviceID = old.(map[string]interface{})["chassis_number"].(string)

					dList = append(dList, &deviceModel)
				}
			}
		}
		if devicesToDetach > 0 {
			err := detachDevices(d, sdwanClient, dList, OperationalTime)
			if err != nil {
				dur := time.Since(start)
				if dur > OperationalTime {
					log.Println("[ERROR] Process timed out, Check from vManage. \nReason:", err)
				} else {
					return err
				}
			}
		}

		dur := time.Since(start)

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
					log.Println("[ERROR] Process timed out, Check from vManage. \nReason:", err)
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

	StatusResult, err := getFinalStatus(sdwanClient, processIds)
	if err != nil {
		return err
	}

	devList := make([]map[string]interface{}, 0, StatusResult.Summary.Total)
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

	var OperationalTime time.Duration
	if val, ok := d.GetOk("timeout"); ok {
		OperationalTime = time.Duration(val.(int)-2) * time.Second
	}

	devicesList := []*models.Device{}
	for _, deviceObject := range d.Get("devices_list").(*schema.Set).List() {
		if deviceObject.(map[string]interface{})["attach"].(bool) {
			deviceModel := &models.Device{
				DeviceID: deviceObject.(map[string]interface{})["chassis_number"].(string),
			}
			devicesList = append(devicesList, deviceModel)
		}
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

func getFinalStatus(client *client.Client, ids string) (*models.StatusData, error) {
	pIDs := strings.Split(ids, ":")

	StatusResult := &models.StatusData{
		Summary: &models.SummaryObject{
			Total: 0,
		},
	}

	for _, ptID := range pIDs {

		status, err := getStatusByID(client, ptID)
		if err != nil {
			return nil, fmt.Errorf("[ERROR] Some error occured :%v", err)
		}

		if status == nil || status.Summary == nil {
			return nil, fmt.Errorf("[ERROR] Status load failed in between")
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
						return nil, fmt.Errorf("[ERROR] Cannot detach a device without attaching it")
					}

					StatusResult.IDs[id] = false
				}
			}
		}
	}
	return StatusResult, nil
}

func attachDevices(d *schema.ResourceData, client *client.Client, dList []string, cnt int, timeout time.Duration) error {

	log.Println("[DEBUG] Beginning Devices Attachment")

	cont1, err := generateDeviceConfig(d, client, dList)
	if err != nil {
		return err
	}

	dtID := d.Get("device_template_id").(string)
	dURL := fmt.Sprintf("dataservice/template/device/object/%s", dtID)

	cont, err := client.GetViaURL(dURL)
	if err != nil {
		return err
	}

	templateType := stripQuotes(cont.S("configType").String())
	// TODO

	devObjects := cont1.S("data").Data().([]interface{})

	if matchedCnt, err := strconv.Atoi(cont1.S("countSummary", "matchCount").String()); err == nil {
		if matchedCnt < cnt {
			if matchedCnt == 0 {
				return fmt.Errorf("[ERROR] No such Device found in the mentioned file. Please check again")
			}
			return fmt.Errorf("[ERROR] Only %d such Device(s) found in the csv file. Need %d Device(s)",
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

	if templateType == "template" {

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

func processStatus(cont *container.Container) (*models.StatusData, error) {
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

func getStatusByID(client *client.Client, id string) (*models.StatusData, error) {

	statusURL := fmt.Sprintf("/dataservice/device/action/status/%s", id)
	cont, err := client.GetViaURL(statusURL)
	if err != nil {
		return nil, err
	}

	var status *models.StatusData
	if cont.Exists("summary", "total") {
		status, err = processStatus(cont)
		return status, err
	}
	// else { // Status was probably pruned by vManage automatically so we just have to assume it was a success
	// 	fake_data := `
	// 	{
	// 		"header": {
	// 		  "generatedOn": 1649425998320,
	// 		  "viewKeys": {
	// 			"uniqueKey": [],
	// 			"preferenceKey": "grid-DeviceActionDefault"
	// 		  },
	// 		  "columns": [
	// 			{
	// 			  "title": "Status",
	// 			  "property": "status",
	// 			  "display": "iconAndText",
	// 			  "iconProperty": "statusId",
	// 			  "icon": [
	// 				{
	// 				  "key": "success",
	// 				  "value": "images/action_success.png"
	// 				},
	// 				{
	// 				  "key": "success_warning",
	// 				  "value": "images/action_success_warning.png"
	// 				},
	// 				{
	// 				  "key": "failure",
	// 				  "value": "images/action_failure.png"
	// 				},
	// 				{
	// 				  "key": "in_progress",
	// 				  "value": "images/action_in_progress.png"
	// 				},
	// 				{
	// 				  "key": "scheduled",
	// 				  "value": "images/action_scheduled.png"
	// 				},
	// 				{
	// 				  "key": "time_out",
	// 				  "value": "images/action_timedout.png"
	// 				},
	// 				{
	// 				  "key": "skipped",
	// 				  "value": "images/action_skipped.png"
	// 				}
	// 			  ],
	// 			  "dataType": "string"
	// 			},
	// 			{
	// 			  "title": "Device IP",
	// 			  "property": "deviceID",
	// 			  "dataType": "ipv4"
	// 			},
	// 			{
	// 			  "title": "Message",
	// 			  "property": "currentActivity",
	// 			  "minWidth": 300,
	// 			  "dataType": "string"
	// 			},
	// 			{
	// 			  "title": "Start Time",
	// 			  "property": "startTime",
	// 			  "displayFormat": "DD MMM YYYY h:mm:ss A z",
	// 			  "inputFormat": "unix-time",
	// 			  "width": 220,
	// 			  "dataType": "date"
	// 			}
	// 		  ],
	// 		  "fields": [
	// 			{
	// 			  "property": "status",
	// 			  "dataType": "string",
	// 			  "display": "iconAndText"
	// 			},
	// 			{
	// 			  "property": "deviceID",
	// 			  "dataType": "ipv4"
	// 			},
	// 			{
	// 			  "property": "currentActivity",
	// 			  "dataType": "string"
	// 			},
	// 			{
	// 			  "property": "startTime",
	// 			  "dataType": "date"
	// 			},
	// 			{
	// 			  "property": "system-ip",
	// 			  "dataType": "string"
	// 			}
	// 		  ]
	// 		},
	// 		"data": [
	// 		  {
	// 			"local-system-ip": "3.3.4.1",
	// 			"statusType": "push_template_configuration",
	// 			"activity": [
	// 			  "[7-Apr-2022 20:31:39 UTC] Configuring device with feature template: EQUINIX_DHCP_DNS_ICGW_CSR1000V_Template_V02",
	// 			  "[7-Apr-2022 20:31:40 UTC] Checking and creating device in vManage",
	// 			  "[7-Apr-2022 20:31:41 UTC] Generating configuration from template",
	// 			  "[7-Apr-2022 20:31:46 UTC] Device is online",
	// 			  "[7-Apr-2022 20:31:46 UTC] Updating device configuration in vManage",
	// 			  "[7-Apr-2022 20:31:47 UTC] Sending configuration to device",
	// 			  "[7-Apr-2022 20:32:03 UTC] Completed template push to device.",
	// 			  "[7-Apr-2022 20:32:04 UTC] Template successfully attached to device"
	// 			],
	// 			"system-ip": "3.3.4.1",
	// 			"site-id": "3340",
	// 			"uuid": "CSR-7429AC11-5CA2-4C2C-9A09-86F8E4E4D706",
	// 			"tenant-id": "default",
	// 			"@rid": 4254,
	// 			"processId": "push_template_configuration-482dba5d-c442-4e51-92f4-ae504d8d888d",
	// 			"actionConfig": "{\"csv-deviceId\":\"CSR-7429AC11-5CA2-4C2C-9A09-86F8E4E4D706\",\"csv-deviceIP\":\"3.3.4.1\",\"csv-status\":\"complete\",\"csv-host-name\":\"us-east-1\",\"//system/host-name\":\"us-east-1\",\"//system/site-id\":\"3340\",\"//system/system-ip\":\"3.3.4.1\",\"/0/GigabitEthernet2//interface/tunnel-interface/color/value\":\"biz-internet\",\"/0/vpn-instance/dns/vpn_dns_primary/dns-addr\":\"156.154.70.1\",\"/0/vpn-instance/dns/vpn_dns_secondary/dns-addr\":\"156.154.71.1\",\"csv-templateId\":\"b8b378c8-3951-4bef-8495-a83193f037d7\"}",
	// 			"device-type": "vedge",
	// 			"action": "push_template_configuration",
	// 			"startTime": 1649363499559,
	// 			"order": 0,
	// 			"vmanageIP": "1.1.1.2",
	// 			"host-name": "us-east-1",
	// 			"deviceID": "CSR-7429AC11-5CA2-4C2C-9A09-86F8E4E4D706",
	// 			"statusId": "success",
	// 			"currentActivity": "Done - Push Template Configuration",
	// 			"deviceModel": "vedge-CSR-1000v",
	// 			"validity": "valid",
	// 			"requestStatus": "received",
	// 			"status": "Success"
	// 		  }
	// 		],
	// 		"validation": {
	// 		  "statusType": "push_template_configuration",
	// 		  "activity": [
	// 			"[7-Apr-2022 20:31:39 UTC] Starting Checks.",
	// 			"[7-Apr-2022 20:31:39 UTC] Validating if device scheduled for template push are active",
	// 			"[7-Apr-2022 20:31:39 UTC] Sending message to vmanage:1.1.1.2",
	// 			"[7-Apr-2022 20:31:39 UTC] Published messages to vmanage(s)",
	// 			"[7-Apr-2022 20:31:39 UTC] Checks completed."
	// 		  ],
	// 		  "vmanageIP": "1.1.1.2",
	// 		  "system-ip": "Validation",
	// 		  "deviceID": "Validation",
	// 		  "uuid": "Validation",
	// 		  "tenant-id": "default",
	// 		  "@rid": 4252,
	// 		  "statusId": "validation_success",
	// 		  "currentActivity": "Done - Validation",
	// 		  "actionConfig": "{}",
	// 		  "processId": "push_template_configuration-482dba5d-c442-4e51-92f4-ae504d8d888d",
	// 		  "action": "push_template_configuration",
	// 		  "startTime": 1649363499110,
	// 		  "requestStatus": "received",
	// 		  "status": "Validation success",
	// 		  "order": 0
	// 		},
	// 		"summary": {
	// 		  "action": "push_template_configuration",
	// 		  "name": "Push Template Configuration",
	// 		  "detailsURL": "/dataservice/device/action/status",
	// 		  "startTime": "1649363499366",
	// 		  "endTime": "1649363524548",
	// 		  "userSessionUserName": "william",
	// 		  "userSessionIP": "38.122.228.75",
	// 		  "tenantName": "DefaultTenant",
	// 		  "total": 1,
	// 		  "status": "done",
	// 		  "count": {
	// 			"Success": 1
	// 		  }
	// 		},
	// 		"isCancelEnabled": false,
	// 		"isParallelExecutionEnabled": false
	// 	  }`
	// 	cont, err = container.ParseJSON([]byte(fake_data))
	// 	status, err = processStatus(cont)
	// }
	// summary := &models.SummaryObject{
	// 	Total:  0,
	// 	Status: "",
	// 	Count:  make(map[string]interface{}),
	// }
	// ids := make(map[string]bool)
	// results := make(map[string]*models.DataObject)
	// status := &models.StatusData{
	// 	Summary: summary,
	// 	IDs:     ids,
	// 	Data:    results,
	// }
	// return status, nil
	// return nil, fmt.Errorf("[ERROR] No such process started\n%s", cont.BytesIndent("", "  "))

	return nil, fmt.Errorf("[ERROR] No device attachment job found")
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

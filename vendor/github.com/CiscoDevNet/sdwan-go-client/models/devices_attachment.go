package models

import (
	"fmt"
	"log"
)

type DeviceConfig struct {
	DevicesCount int
	*DeviceTemplateList
}

type DeviceTemplateList struct {
	TemplateID string        `json:"template_id"`
	Device     []interface{} `json:"device"`
}

type DeviceAttach struct {
	File string
	Data *DeviceAttachData
}

type DeviceAttachData struct {
	TemplateID string   `json:"template_id"`
	DeviceIDs  []string `json:"device_ids"`
}

type DeviceDetach struct {
	DeviceType string    `json:"device_type"`
	Devices    []*Device `json:"devices"`
}

type Device struct {
	DeviceID string `json:"device_id"`
	DeviceIP string `json:"device_ip,omitempty"`
}

type StatusData struct {
	Summary *SummaryObject
	IDs     map[string]bool
	Data    map[string]*DataObject
}

type DataObject struct {
	Status          string
	Activity        []string
	CurrentActivity string
	ActionConfig    string
}

type SummaryObject struct {
	Total  int
	Status string
	Count  map[string]interface{}
}

//ToMap - Returns map for Device Configuration Attachment model
func (dtConfig *DeviceConfig) ToMap() (map[string]interface{}, error) {
	dtAttrMap := make(map[string]interface{})

	tempList := make([]map[string]interface{}, 0, dtConfig.DevicesCount)
	if config, err := dtConfig.makeDeviceConfig(); err == nil {
		tempList = append(tempList, config)
	} else {
		return nil, fmt.Errorf("[ERROR] Unable to create device configuration")
	}

	A(dtAttrMap, "deviceTemplateList", tempList)

	log.Println("Device Template List: ", dtAttrMap)
	return dtAttrMap, nil
}

func (dtConfig *DeviceTemplateList) makeDeviceConfig() (map[string]interface{}, error) {
	dtAttrMap := make(map[string]interface{})

	A(dtAttrMap, "templateId", dtConfig.TemplateID)
	// devices := make([]interface{}, 0, 1)
	A(dtAttrMap, "device", dtConfig.Device)
	A(dtAttrMap, "isEdited", false)
	A(dtAttrMap, "isMasterEdited", false)

	log.Println("Device Configuration: ", dtAttrMap)
	return dtAttrMap, nil
}

//ToMap - Returns map for Device Attachment model
func (dtAttach *DeviceAttach) ToMap() (map[string]interface{}, error) {
	dtAttrMap := make(map[string]interface{})

	A(dtAttrMap, "file", dtAttach.File)

	devicesData, err := makeDeviceAttachData(dtAttach.Data)
	if err == nil {
		A(dtAttrMap, "data", devicesData)
	}

	return dtAttrMap, nil
}

func makeDeviceAttachData(dtAttachData *DeviceAttachData) (map[string]interface{}, error) {
	dtAttrMap := make(map[string]interface{})

	A(dtAttrMap, "templateId", dtAttachData.TemplateID)
	A(dtAttrMap, "devices", dtAttachData.DeviceIDs)
	A(dtAttrMap, "isEdited", false)
	A(dtAttrMap, "isMasterEdited", false)

	return dtAttrMap, nil
}

//ToMap - Returns map for Device Detachment model
func (dtDetach *DeviceDetach) ToMap() (map[string]interface{}, error) {
	dtAttrMap := make(map[string]interface{})

	A(dtAttrMap, "deviceType", dtDetach.DeviceType)

	if len(dtDetach.Devices) > 0 {
		devicesList := make([]interface{}, 0)

		for _, device := range dtDetach.Devices {
			deviceMap, _ := device.makeDevice()
			devicesList = append(devicesList, deviceMap)
		}
		A(dtAttrMap, "devices", devicesList)
	}

	return dtAttrMap, nil
}

func (dev *Device) makeDevice() (map[string]interface{}, error) {
	dMap := make(map[string]interface{})

	A(dMap, "deviceId", dev.DeviceID)
	A(dMap, "deviceIP", dev.DeviceIP)

	return dMap, nil
}

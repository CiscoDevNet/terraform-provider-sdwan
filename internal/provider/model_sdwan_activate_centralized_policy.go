// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/sjson"
)

type ActivateCentralizedPolicy struct {
	Id      types.String `tfsdk:"id"`
	Version types.Int64  `tfsdk:"version"`
}

func (data ActivateCentralizedPolicy) activatePolicy(ctx context.Context, client *sdwan.Client, isEdited bool, timeout *int64) error {
	var body string
	if isEdited {
		body = `{"isEdited":true}`
	} else {
		body = `{}`
	}
	res, err := client.Post("/template/policy/vsmart/activate/"+data.Id.ValueString()+"?confirm=true", body)
	if err != nil {
		return err
	}
	actionId := res.Get("id").String()
	err = helpers.WaitForActionToComplete(ctx, client, actionId, timeout)
	if err != nil {
		return err
	}
	return nil
}

func (data ActivateCentralizedPolicy) getPushBody(ctx context.Context, client *sdwan.Client) (string, error) {
	// Get all device templates
	res, err := client.Get("/template/device?feature=all")
	if err != nil {
		return "", err
	}
	// Filter for vSmart templates with devices attached
	var vsmartTemplateIds []string
	if res.Get("data").Exists() {
		for _, item := range res.Get("data").Array() {
			if item.Get("deviceType").String() == "vsmart" && item.Get("devicesAttached").Int() > 0 {
				templateId := item.Get("templateId").String()
				if templateId != "" {
					vsmartTemplateIds = append(vsmartTemplateIds, templateId)
				}
			}
		}
	}

	// For each vSmart template with devices attached, generate the attach payload
	var deviceTemplateList []interface{}
	for _, templateId := range vsmartTemplateIds {
		res, err := client.Get("/template/device/config/attached/" + templateId)
		if err != nil {
			return "", err
		}
		var deviceIds []string
		if res.Get("data").Exists() {
			for _, device := range res.Get("data").Array() {
				uuid := device.Get("uuid").String()
				if uuid != "" {
					deviceIds = append(deviceIds, uuid)
				}
			}
		}
		inputPayload, _ := sjson.Set("", "templateId", templateId)
		inputPayload, _ = sjson.Set(inputPayload, "deviceIds", deviceIds)
		inputPayload, _ = sjson.Set(inputPayload, "isEdited", true)
		inputPayload, _ = sjson.Set(inputPayload, "isMasterEdited", false)

		res, err = client.Post("/template/device/config/input/", inputPayload)
		if err != nil {
			return "", err
		}
		// Get device list from response
		deviceList := []interface{}{}
		if res.Get("data").Exists() {
			for _, data := range res.Get("data").Array() {
				devices := map[string]interface{}{}
				for key, value := range data.Map() {
					devices[key] = value.Value()
				}
				deviceList = append(deviceList, devices)
			}
		}
		templateAttachPayload := map[string]interface{}{
			"templateId":     templateId,
			"device":         deviceList,
			"isEdited":       true,
			"isMasterEdited": false,
		}
		deviceTemplateList = append(deviceTemplateList, templateAttachPayload)
	}
	attachPayload, _ := sjson.Set("", "deviceTemplateList", deviceTemplateList)
	return attachPayload, nil
}

// Section below is generated&owned by "gen/generator.go". //template:begin processImport

// End of section. //template:end processImport

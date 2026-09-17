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
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type WANEdgeCertificate struct {
	Id                types.String `tfsdk:"id"`
	ChassisNumber     types.String `tfsdk:"chassis_number"`
	SerialNumber      types.String `tfsdk:"serial_number"`
	Validity          types.String `tfsdk:"validity"`
	SendToControllers types.Bool   `tfsdk:"send_to_controllers"`
}

// getCertificate returns the WAN edge entry of the certificate list matching the chassis number.
func (data WANEdgeCertificate) getCertificate(ctx context.Context, client *sdwan.Client) (gjson.Result, bool, error) {
	res, err := client.Get("/certificate/vedge/list")
	if err != nil {
		return gjson.Result{}, false, err
	}
	var entry gjson.Result
	found := false
	res.Get("data").ForEach(func(_, v gjson.Result) bool {
		if v.Get("chasisNumber").String() == data.ChassisNumber.ValueString() {
			entry = v
			found = true
			return false
		}
		return true
	})
	return entry, found, nil
}

func (data WANEdgeCertificate) toBody(ctx context.Context, validity string, sendToControllers bool) string {
	body := `[]`
	body, _ = sjson.Set(body, "0.chasisNumber", data.ChassisNumber.ValueString())
	body, _ = sjson.Set(body, "0.validity", validity)
	body, _ = sjson.Set(body, "0.sendToControllers", sendToControllers)
	return body
}

func (data WANEdgeCertificate) legacyBody(ctx context.Context, validity string) string {
	body := `[]`
	body, _ = sjson.Set(body, "0.chasisNumber", data.ChassisNumber.ValueString())
	body, _ = sjson.Set(body, "0.validity", validity)
	return body
}

// setValidity stores validity and waits for Manager's propagation action when requested.
func (data WANEdgeCertificate) setValidity(ctx context.Context, client *sdwan.Client, validity string, sendToControllers bool, taskTimeout *int64) error {
	res, err := client.Post("/certificate/save/vedge/list", data.toBody(ctx, validity, sendToControllers))
	if err != nil {
		return fmt.Errorf("%s, %s", err, res.String())
	}
	if sendToControllers {
		return data.waitForAction(ctx, client, res, taskTimeout, "certificate validity update")
	}
	return nil
}

func (data WANEdgeCertificate) setValidityLegacy(ctx context.Context, client *sdwan.Client, validity string) error {
	res, err := client.Post("/certificate/save/vedge/list", data.legacyBody(ctx, validity))
	if err != nil {
		return fmt.Errorf("%s, %s", err, res.String())
	}
	return nil
}

func isSendToControllersUnsupported(err error) bool {
	message := strings.ToLower(err.Error())
	return strings.Contains(message, "sendtocontrollers") &&
		(strings.Contains(message, "impermissible") || strings.Contains(message, "incorrect") || strings.Contains(message, "unknown") || strings.Contains(message, "invalid") || strings.Contains(message, "unsupported") || strings.Contains(message, "unrecognized"))
}

// sendToControllers pushes the WAN edge list to all controllers and waits for the action to finish.
func (data WANEdgeCertificate) sendToControllers(ctx context.Context, client *sdwan.Client, taskTimeout *int64) error {
	res, err := client.Post("/certificate/vedge/list?action=push", "{}")
	if err != nil {
		return fmt.Errorf("%s, %s", err, res.String())
	}
	actionId := res.Get("sendToControllerId").String()
	if actionId == "" {
		actionId = res.Get("id").String()
	}
	if actionId == "" {
		return fmt.Errorf("certificate push returned no action ID: %s", res.String())
	}
	err, _ = helpers.WaitForActionToComplete(ctx, client, actionId, taskTimeout)
	return err
}

func (data WANEdgeCertificate) waitForAction(ctx context.Context, client *sdwan.Client, res gjson.Result, taskTimeout *int64, operation string) error {
	actionId := res.Get("sendToControllerId").String()
	if actionId == "" {
		actionId = res.Get("id").String()
	}
	if actionId == "" {
		return fmt.Errorf("%s returned no action ID: %s", operation, res.String())
	}
	return waitForCertificateAction(ctx, client, actionId, taskTimeout)
}

func waitForCertificateAction(ctx context.Context, client *sdwan.Client, actionId string, taskTimeout *int64) error {
	maxAttempts := *taskTimeout / 5
	for attempts := int64(0); ; attempts++ {
		time.Sleep(5 * time.Second)
		res, err := client.Get("/device/action/status/" + actionId)
		if err != nil {
			return err
		}
		status := strings.ToLower(res.Get("summary.status").String())
		switch status {
		case "done", "success", "successful", "complete", "completed":
			return certificateActionFailures(actionId, res)
		case "failure", "failed":
			return certificateActionFailures(actionId, res)
		}
		if attempts > maxAttempts {
			return fmt.Errorf("maximum waiting time for action '%s' reached", actionId)
		}
	}
}

func certificateActionFailures(actionId string, res gjson.Result) error {
	var failures []string
	res.Get("data").ForEach(func(_, v gjson.Result) bool {
		if strings.Contains(strings.ToLower(v.Get("statusId").String()), "failure") {
			failures = append(failures, fmt.Sprintf("Action %s for device %s failed. Activity log: %+v", actionId, v.Get("deviceID").String(), v.Get("activity").String()))
		}
		return true
	})
	if strings.Contains(strings.ToLower(res.Get("validation.status").String()), "failure") {
		failures = append(failures, fmt.Sprintf("Validation for action %s failed. Validation log: %+v", actionId, res.Get("validation.activity").String()))
	}
	if len(failures) > 0 {
		return errors.New(strings.Join(failures, "\n"))
	}
	return nil
}

func (data *WANEdgeCertificate) fromBody(ctx context.Context, res gjson.Result) {
	data.Id = types.StringValue(res.Get("chasisNumber").String())
	data.ChassisNumber = types.StringValue(res.Get("chasisNumber").String())
	if value := res.Get("serialNumber"); value.Exists() {
		data.SerialNumber = types.StringValue(value.String())
	}
	if value := res.Get("validity"); value.Exists() {
		data.Validity = types.StringValue(value.String())
	}
}

func (data *WANEdgeCertificate) processImport(ctx context.Context) {
	if data.SendToControllers.IsNull() {
		data.SendToControllers = types.BoolValue(true)
	}
}

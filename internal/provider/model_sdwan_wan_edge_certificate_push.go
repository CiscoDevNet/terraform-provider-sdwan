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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
)

type WANEdgeCertificatePush struct {
	Id       types.String `tfsdk:"id"`
	Triggers types.Map    `tfsdk:"triggers"`
	Version  types.Int64  `tfsdk:"version"`
}

// push sends the current WAN edge list to all controllers and waits for the action to finish.
func (data WANEdgeCertificatePush) push(ctx context.Context, client *sdwan.Client, taskTimeout *int64) error {
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
	return waitForCertificatePushAction(ctx, client, actionId, taskTimeout)
}

func waitForCertificatePushAction(ctx context.Context, client *sdwan.Client, actionId string, taskTimeout *int64) error {
	maxAttempts := *taskTimeout / 5
	for attempts := int64(0); ; attempts++ {
		time.Sleep(5 * time.Second)
		res, err := client.Get("/device/action/status/" + actionId)
		if err != nil {
			return err
		}
		status := strings.ToLower(res.Get("summary.status").String())
		switch status {
		case "done", "success", "successful", "complete", "completed", "failure", "failed":
			return certificatePushActionFailures(actionId, res)
		}
		if attempts > maxAttempts {
			return fmt.Errorf("maximum waiting time for action '%s' reached", actionId)
		}
	}
}

func certificatePushActionFailures(actionId string, res gjson.Result) error {
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

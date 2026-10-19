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

package helpers

import (
	"testing"

	"github.com/tidwall/gjson"
)

func TestIsDeployInProgress(t *testing.T) {
	cases := []struct {
		name, body string
		want       bool
	}{
		// A second deploy on a config group while one runs.
		{"config group CFGRP0017", `{"error":{"code":"CFGRP0017","details":"A deployment for this Configuration Group is in-progress. Please wait for it to complete.","type":"error"}}`, true},
		// Policy group: a disassociate refused while a deploy runs.
		{"policy group PLGRP0017", `{"error":{"message":"Policy Group State validation error","code":"PLGRP0017","details":"A deployment for Policy Group [pg1] is in-progress. Please wait for it to complete.","type":"error"}}`, true},
		{"other code, in-progress text", `{"error":{"code":"XYZ0001","details":"A deployment for this Policy Group is in-progress."}}`, true},
		{"association conflict", `{"error":{"code":"CFGRP0018","details":"Device(s) already associated to group(s)"}}`, false},
		{"server error", `{"error":{"code":"SCHVALID9999","message":"Unexpected Error"}}`, false},
		{"empty", ``, false},
	}
	for _, c := range cases {
		if got := IsDeployInProgress(gjson.Parse(c.body)); got != c.want {
			t.Errorf("%s: got %v, want %v", c.name, got, c.want)
		}
	}
}

// Deploy task results as the Manager returns them, shortened.
const (
	taskSkippedByController = `{"summary":{"status":"done","total":0},"data":[],"validation":{"statusId":"failure","status":"Failure","activity":["[20-Oct-2026 12:38:09 UTC] Starting Checks.","[20-Oct-2026 12:38:10 UTC] \"Deploy policy group to associated devices\" operation is already in progress for device:10.0.0.104 since Oct 20 12:38:09 2026 UTC. Wait until that operation completes before executing \"Deploy policy group to associated devices\""]}}`
	taskControllerConflict  = `{"summary":{"status":"done","total":6},"validation":{"statusId":"success","status":"Success"},"data":[` +
		`{"deviceID":"C8K-83E4","statusId":"success","activity":["Done"]},` +
		`{"deviceID":"5b8d9988","statusId":"failure","activity":["Failed to deploy device with Policy Group - pg1, Transaction 36293 conflicts with transaction 36275 started by user vmanage-admin: /ncs:devices/device{vsmart-5b8d9988}/config/viptela-system:system/vbond/port"]},` +
		`{"deviceID":"6fe90e6c","statusId":"failure","activity":["Failed to deploy device with Policy Group - pg1, Transaction 36297 conflicts with transaction 36260 started by user vmanage-admin"]}]}`
	taskGenerationCrash  = `{"summary":{"status":"done","total":1},"validation":{"statusId":"success"},"data":[{"deviceID":"C8K-B5EE","statusId":"failure","activity":["Failed to deploy device with Policy Group - pg1, Error while generating cli. Error: NullPointerException"]}]}`
	taskRouterRejected   = `{"summary":{"status":"done","total":1},"validation":{"statusId":"success"},"data":[{"deviceID":"C8K-83E4","statusId":"failure","activity":["Device refused one or more commands: ip address 172.16.1.44 255.255.255.0 % 172.16.1.0 overlaps with GigabitEthernet1","Failed to deploy Config Group cg1"]}]}`
	taskConflictAndOther = `{"summary":{"status":"done","total":2},"validation":{"statusId":"success"},"data":[` +
		`{"deviceID":"5b8d9988","statusId":"failure","activity":["Transaction 1 conflicts with transaction 2"]},` +
		`{"deviceID":"C8K-83E4","statusId":"failure","activity":["% 172.16.1.0 overlaps with GigabitEthernet1"]}]}`
	taskSuccess = `{"summary":{"status":"done","total":5},"validation":{"statusId":"success","status":"Success"},"data":[{"deviceID":"C8K-B5EE","statusId":"success"}]}`
)

func TestDeployCollision(t *testing.T) {
	cases := []struct {
		name, body string
		collision  bool
	}{
		{"controller pre-check skip, nothing deployed", taskSkippedByController, true},
		{"controller transaction conflict", taskControllerConflict, true},
		{"generation crash is not a collision", taskGenerationCrash, false},
		{"router rejected the config", taskRouterRejected, false},
		{"conflict plus another failure", taskConflictAndOther, false},
		{"success", taskSuccess, false},
	}
	for _, c := range cases {
		if got := DeployCollision(gjson.Parse(c.body)) != ""; got != c.collision {
			t.Errorf("%s: collision %v, want %v", c.name, got, c.collision)
		}
	}
}

func TestDeployOutcome(t *testing.T) {
	cases := []struct {
		name, body string
		wantErr    bool
	}{
		// WaitForActionToComplete misses this one ("Failure" with a capital F).
		{"validation failure, capital F", taskSkippedByController, true},
		{"device failure", taskRouterRejected, true},
		{"success", taskSuccess, false},
	}
	for _, c := range cases {
		err, _ := deployOutcome("deploy_x", gjson.Parse(c.body))
		if (err != nil) != c.wantErr {
			t.Errorf("%s: error %v, want error %v", c.name, err, c.wantErr)
		}
	}
}

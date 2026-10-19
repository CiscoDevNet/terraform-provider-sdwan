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
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
)

// IsDeployInProgress reports whether a request was refused because a deploy is running on the
// same group. The Manager allows one deploy per group at a time and answers CFGRP0017 (config
// group) or PLGRP0017 (policy group), "A deployment for ... is in-progress".
func IsDeployInProgress(res sdwan.Res) bool {
	switch res.Get("error.code").String() {
	case "CFGRP0017", "PLGRP0017":
		return true
	}
	details := strings.ToLower(res.Get("error.details").String() + " " + res.Get("error.message").String())
	return strings.Contains(details, "in-progress") || strings.Contains(details, "in progress")
}

// RetryWhileDeployInProgress runs call, and while the Manager refuses it because a deploy is
// running on the same group, waits and retries, up to taskTimeout seconds. The Manager refuses
// not only a second deploy but also associate/disassociate calls on a group whose deploy is in
// progress (CFGRP0017 "A deployment for this Group is in-progress"). With several
// states owning devices on one group, those calls otherwise fail instead of queueing.
func RetryWhileDeployInProgress(ctx context.Context, taskTimeout *int64, call func() (sdwan.Res, error)) (sdwan.Res, error) {
	deadline := time.Now().Add(time.Duration(*taskTimeout) * time.Second)
	for {
		res, err := call()
		if err == nil || !IsDeployInProgress(res) || time.Now().After(deadline) {
			return res, err
		}
		// Spread the retries, so states waiting on the same group don't all retry at once.
		wait := 10*time.Second + time.Duration(rand.Intn(5000))*time.Millisecond
		tflog.Info(ctx, fmt.Sprintf("A deploy is running on this group (%s), retrying in %s", res.Get("error.code").String(), wait))
		select {
		case <-ctx.Done():
			return res, ctx.Err()
		case <-time.After(wait):
		}
	}
}

// PostDeploy sends a group deploy request, waiting while another deploy runs on the same group.
func PostDeploy(ctx context.Context, client *sdwan.Client, path, body string, taskTimeout *int64) (sdwan.Res, error) {
	return RetryWhileDeployInProgress(ctx, taskTimeout, func() (sdwan.Res, error) { return client.Post(path, body) })
}

// DeployCollision reports why a finished deploy task failed only because another deploy ran at
// the same time, or "" if it failed (or succeeded) for any other reason. Manager can accept two
// deploys on one policy group sent close together, before it has registered the first one, and
// the second one then fails inside its task:
//   - the controller pre-check skips it, nothing deployed: validation failure with "operation is
//     already in progress for device: <controller>";
//   - its controller writes conflict with the other deploy's: every failed device reports
//     "Transaction … conflicts with transaction …".
//
// Such a deploy can be sent again. Any other failure (a router rejecting the config, Manager
// failing to generate it) is not a collision.
func DeployCollision(task gjson.Result) string {
	if validationFailed(task) {
		for _, line := range task.Get("validation.activity").Array() {
			if strings.Contains(line.String(), "already in progress") {
				return line.String()
			}
		}
		return ""
	}
	var conflict string
	for _, d := range task.Get("data").Array() {
		if !strings.Contains(strings.ToLower(d.Get("statusId").String()), "failure") {
			continue
		}
		activity := d.Get("activity").String()
		if !strings.Contains(activity, "conflicts with transaction") {
			return ""
		}
		if conflict == "" {
			conflict = activity
		}
	}
	return conflict
}

func validationFailed(task gjson.Result) bool {
	return strings.Contains(strings.ToLower(task.Get("validation.statusId").String()+" "+task.Get("validation.status").String()), "failure")
}

// DeployAndWait sends a group deploy (PostDeploy: waits while Manager refuses it because
// another deploy runs) and waits for its task. A task that failed only by colliding with another
// deploy (DeployCollision) is sent again after 10–15 s; when it still collides at taskTimeout,
// that is an error. It returns the POST response and error, or the task error, and the task's
// warnings.
//
// Unlike WaitForActionToComplete, a failed validation is always an error here (Manager reports
// "status":"Failure" with a capital F, which that check does not match): for the group-device
// resources a deploy that did nothing must not pass.
func DeployAndWait(ctx context.Context, client *sdwan.Client, path, body string, taskTimeout *int64) (postRes sdwan.Res, postErr, taskErr error, warnings string) {
	deadline := time.Now().Add(time.Duration(*taskTimeout) * time.Second)
	for {
		res, err := PostDeploy(ctx, client, path, body, taskTimeout)
		if err != nil {
			return res, err, nil, ""
		}
		id := res.Get("parentTaskId").String()
		task, err := waitForTask(ctx, client, id, taskTimeout)
		if err != nil {
			return res, nil, err, ""
		}
		collision := DeployCollision(task)
		if collision == "" {
			taskErr, warnings = deployOutcome(id, task)
			return res, nil, taskErr, warnings
		}
		if time.Now().After(deadline) {
			return res, nil, fmt.Errorf("Deploy %s still collided with another deploy on the group when task_timeout was reached, nothing more was sent: %s", id, collision), ""
		}
		wait := 10*time.Second + time.Duration(rand.Intn(5000))*time.Millisecond
		tflog.Warn(ctx, fmt.Sprintf("Deploy %s collided with another deploy on the group, sending it again in %s: %s", id, wait, collision))
		select {
		case <-ctx.Done():
			return res, nil, ctx.Err(), ""
		case <-time.After(wait):
		}
	}
}

// waitForTask polls /device/action/status/{id} every 5 s until the task is done, up to taskTimeout
// (as WaitForActionToComplete), and returns the final status.
func waitForTask(ctx context.Context, client *sdwan.Client, id string, taskTimeout *int64) (gjson.Result, error) {
	maxAttempts := *taskTimeout / 5
	for attempts := int64(0); ; attempts++ {
		select {
		case <-ctx.Done():
			return gjson.Result{}, ctx.Err()
		case <-time.After(5 * time.Second):
		}
		res, err := client.Get("/device/action/status/" + id)
		if err != nil {
			return gjson.Result{}, err
		}
		if res.Get("summary.status").String() == "done" {
			return res, nil
		}
		tflog.Debug(ctx, fmt.Sprintf("[DEBUG] Waiting for action '%s' to complete.", id))
		if attempts > maxAttempts {
			return gjson.Result{}, fmt.Errorf("Maximum waiting time for action '%s' reached.", id)
		}
	}
}

// deployOutcome turns a finished, non-colliding deploy task into an error (failed devices, failed
// validation) and warnings, with the same messages as WaitForActionToComplete.
func deployOutcome(id string, task gjson.Result) (error, string) {
	var failures []string
	task.Get("data").ForEach(func(_, v gjson.Result) bool {
		if strings.Contains(strings.ToLower(v.Get("statusId").String()), "failure") {
			failures = append(failures, fmt.Sprintf("Action %s for device %s failed. Activity log: %+v", id, v.Get("deviceID").String(), v.Get("activity").String()))
		}
		return true
	})
	if validationFailed(task) {
		failures = append(failures, fmt.Sprintf("Validation for action %s failed. Validation log: %+v", id, task.Get("validation.activity").String()))
	}
	if len(failures) > 0 {
		return errors.New(strings.Join(failures, "\n")), ""
	}
	var warnings string
	task.Get("validation.activity").ForEach(func(_, line gjson.Result) bool {
		if strings.Contains(line.String(), "operation is already in progress") {
			warnings = "The configuration update is skipped because another operation is already in progress for this device. Some changes may not have been applied. Run 'terraform apply' again to ensure all changes are applied."
			return false
		}
		return true
	})
	return nil, warnings
}

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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func GetStringFromList(list basetypes.ListValue) types.String {
	v := make([]string, len(list.Elements()))
	for r := range list.Elements() {
		v[r] = list.Elements()[r].String()
	}
	return types.StringValue("[" + strings.Join(v, ",") + "]")
}

func GetStringFromSet(set basetypes.SetValue) types.String {
	v := make([]string, len(set.Elements()))
	for r := range set.Elements() {
		v[r] = set.Elements()[r].String()
	}
	return types.StringValue("[" + strings.Join(v, ",") + "]")
}

func GetStringList(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.ListValueMust(types.StringType, v)
}

func GetInt64List(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.ListValueMust(types.Int64Type, v)
}

func GetStringSet(result []gjson.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.SetValueMust(types.StringType, v)
}

func GetInt64Set(result []gjson.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.SetValueMust(types.Int64Type, v)
}

func WaitForActionToComplete(ctx context.Context, client *sdwan.Client, id string, taskTimeout *int64) error {
	var MaxActionAttempts = *taskTimeout / 5

	for attempts := 0; ; attempts++ {
		time.Sleep(5 * time.Second)
		res, err := client.Get("/device/action/status/" + id)
		if err != nil {
			return err
		}
		if res.Get("summary.status").String() == "done" {
			var failures []string
			res.Get("data").ForEach(func(k, v gjson.Result) bool {
				if strings.Contains(v.Get("statusId").String(), "failure") {
					failures = append(failures, fmt.Sprintf("Action %s for device %s failed. Activity log: %+v", id, v.Get("deviceID").String(), v.Get("activity").String()))
				}
				return true
			})
			if strings.Contains(res.Get("validation.status").String(), "failure") {
				failures = append(failures, fmt.Sprintf("Validation for action %s failed. Validation log: %+v", id, res.Get("validation.activity").String()))
			}
			if len(failures) > 0 {
				return errors.New(strings.Join(failures, "\n"))
			}
			break
		} else {
			tflog.Debug(ctx, fmt.Sprintf("[DEBUG] Waiting for action '%s' to complete.", id))
		}
		if attempts > int(MaxActionAttempts) {
			tflog.Debug(ctx, fmt.Sprintf("[DEBUG] Maximum number of attempts reached for action '%s'.", id))
			return fmt.Errorf("Maximum waiting time for action '%s' reached.", id)
		}
	}
	return nil
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

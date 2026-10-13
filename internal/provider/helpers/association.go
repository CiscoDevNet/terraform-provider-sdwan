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
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
)

// The Manager's DELETE .../{group}/device/associate/ removes a device from whatever group it is
// on, not only from the group in the URL. A release sent for a device that has meanwhile
// moved to another group therefore pulls it out of that group. The helpers below guard both sides:
// the releasing side re-reads its group before the DELETE, and the receiving side checks the
// devices it just associated are still there.

// reassociateWait is how long the receiving side waits before checking (and between retries), so
// that a release from the device's old group, sent in the same apply, has landed.
var reassociateWait = 2 * time.Second

// AssociatedIds returns the device IDs listed in a GET .../device/associate/ response.
func AssociatedIds(res sdwan.Res) map[string]bool {
	ids := map[string]bool{}
	res.Get("devices").ForEach(func(_, v gjson.Result) bool {
		ids[v.Get("id").String()] = true
		return true
	})
	return ids
}

// StillAssociated re-reads the group's associations (associatePath) and returns the subset of
// ids still listed there. Devices that moved on are skipped and logged. A group that no longer
// exists has nothing to release.
func StillAssociated(ctx context.Context, client *sdwan.Client, associatePath string, ids []string) ([]string, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	res, err := client.Get(associatePath)
	if err != nil {
		if strings.Contains(res.Get("error.message").String(), "group passed") { // "Invalid config/policy group passed"
			return nil, nil
		}
		return nil, err
	}
	on := AssociatedIds(res)
	var keep []string
	for _, id := range ids {
		if on[id] {
			keep = append(keep, id)
		} else {
			tflog.Info(ctx, fmt.Sprintf("%s: device %s is no longer on this group, not releasing it", associatePath, id))
		}
	}
	return keep, nil
}

// EnsureAssociated checks that the devices just associated (added) are still on the group and
// associates the missing ones again (associate), up to three times, checking after each. It
// returns the last GET response, which callers reuse for the deploy selection, and the IDs still
// missing at the end.
func EnsureAssociated(ctx context.Context, client *sdwan.Client, associatePath string, added []string, associate func(missing []string) (sdwan.Res, error)) (sdwan.Res, []string, error) {
	var res sdwan.Res
	var missing []string
	for attempt := 0; ; attempt++ {
		select {
		case <-ctx.Done():
			return res, missing, ctx.Err()
		case <-time.After(reassociateWait):
		}
		var err error
		res, err = client.Get(associatePath)
		if err != nil {
			return res, missing, err
		}
		on := AssociatedIds(res)
		missing = missing[:0]
		for _, id := range added {
			if !on[id] {
				missing = append(missing, id)
			}
		}
		if len(missing) == 0 {
			return res, nil, nil
		}
		if attempt == 3 {
			return res, missing, nil
		}
		tflog.Info(ctx, fmt.Sprintf("%s: devices %v were removed again right after association (a release from their old group), associating them again", associatePath, missing))
		if r, err := associate(missing); err != nil {
			return r, missing, err
		}
	}
}

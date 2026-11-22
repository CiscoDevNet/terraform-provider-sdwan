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
	"strings"

	"github.com/tidwall/gjson"
)

// IsOutOfSync reports whether the Manager lists a device as "Out of Sync" in a GET
// .../device/associate/ response. That is what a deploy the router rejected leaves behind: the
// group up-to-date flag stays true (the variables and the group revision were accepted), only
// the status says the router doesn't run the config. Such a device needs a deploy.
func IsOutOfSync(device gjson.Result) bool {
	return strings.EqualFold(strings.TrimSpace(device.Get("configStatusMessage").String()), "Out of Sync")
}

// NeedsRedeploy reports whether Read should flag a device for redeploy: the Manager lists it
// out of date (upToDateKey false) or out of sync. An offline or not-yet-synced device always
// reports out of date with "Sync Pending"; that is not a real out-of-date condition, so it is
// skipped, otherwise an offline device would be flagged on every refresh.
func NeedsRedeploy(device gjson.Result, upToDateKey string) bool {
	if IsOutOfSync(device) {
		return true
	}
	return !device.Get(upToDateKey).Bool() && device.Get("configStatusMessage").String() != "Sync Pending"
}

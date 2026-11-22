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

func TestNeedsRedeploy(t *testing.T) {
	cases := []struct {
		name, body    string
		outOfSync     bool
		needsRedeploy bool
	}{
		{"in sync, up to date", `{"configGroupUpToDate":true,"configStatusMessage":"In Sync"}`, false, false},
		{"group changed", `{"configGroupUpToDate":false,"configStatusMessage":"In Sync"}`, false, true},
		// Router rejected the config: up to date stays true.
		{"router rejected", `{"configGroupUpToDate":true,"configStatusMessage":"Out of Sync"}`, true, true},
		{"out of sync, other case", `{"configGroupUpToDate":true,"configStatusMessage":"out of sync"}`, true, true},
		{"offline", `{"configGroupUpToDate":false,"configStatusMessage":"Sync Pending"}`, false, false},
		{"no status", `{"configGroupUpToDate":true}`, false, false},
	}
	for _, c := range cases {
		v := gjson.Parse(c.body)
		if got := IsOutOfSync(v); got != c.outOfSync {
			t.Errorf("%s: IsOutOfSync got %v, want %v", c.name, got, c.outOfSync)
		}
		if got := NeedsRedeploy(v, "configGroupUpToDate"); got != c.needsRedeploy {
			t.Errorf("%s: NeedsRedeploy got %v, want %v", c.name, got, c.needsRedeploy)
		}
	}
}

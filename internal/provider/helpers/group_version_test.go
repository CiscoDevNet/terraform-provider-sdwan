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

import "testing"

func TestGroupVersionForcesDeploy(t *testing.T) {
	for v, want := range map[string]bool{
		"20.12.3": true,
		"20.14.9": true,
		"20.15.0": false,
		"20.18.4": false,
		"26.1.1":  false,
		"":        true,
	} {
		if got := GroupVersionForcesDeploy(v); got != want {
			t.Errorf("GroupVersionForcesDeploy(%q) = %v, want %v", v, got, want)
		}
	}
}

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

	"github.com/hashicorp/go-version"
)

func TestDefaultForVersion(t *testing.T) {
	defaults := []VersionDefault{{MinVersion: "20.18.5", Value: false}, {MinVersion: "26.1.1", Value: "x"}}
	cases := []struct {
		ver  string
		want any
	}{
		{"20.15.4", true},
		{"20.18.4", true},
		{"20.18.5", false},
		{"20.18.5.1", false},
		{"26.1.1", "x"},
	}
	for _, c := range cases {
		got := DefaultForVersion(version.Must(version.NewVersion(c.ver)), true, defaults...)
		if got != c.want {
			t.Errorf("version %s: got %v, want %v", c.ver, got, c.want)
		}
	}
	if got := DefaultForVersion(nil, true, defaults...); got != true {
		t.Errorf("nil version: got %v, want true", got)
	}
}

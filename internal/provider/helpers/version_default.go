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
	"github.com/hashicorp/go-version"
)

// VersionDefault is the default value SD-WAN Manager expects from MinVersion onwards.
type VersionDefault struct {
	MinVersion string
	Value      any
}

// DefaultForVersion returns the default value to send for an unset attribute: the Value of the
// last VersionDefault whose MinVersion the Manager version meets, otherwise base.
// VersionDefaults are listed in ascending MinVersion order. A nil ver returns base.
func DefaultForVersion(ver *version.Version, base any, defaults ...VersionDefault) any {
	value := base
	if ver == nil {
		return value
	}
	for _, d := range defaults {
		if ver.GreaterThanOrEqual(version.Must(version.NewVersion(d.MinVersion))) {
			value = d.Value
		}
	}
	return value
}

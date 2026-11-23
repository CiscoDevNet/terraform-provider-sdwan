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

// minGroupUpdateVersion is the Manager version from which a group change no longer needs to
// deploy every device: the Manager itself marks the affected devices out of date.
var minGroupUpdateVersion = version.Must(version.NewVersion("20.15.0"))

// GroupVersionForcesDeploy reports whether a change of the wired group version must deploy every
// device of a _devices resource. From 20.15 the Manager marks exactly the devices whose
// configuration a group change affects as out of date, so the version change only makes the
// resource run and the Manager's flags select the devices. Below 20.15 (or if
// the version can't be read) the version change deploys all devices.
func GroupVersionForcesDeploy(managerVersion string) bool {
	v, err := version.NewVersion(managerVersion)
	if err != nil {
		return true
	}
	return v.LessThan(minGroupUpdateVersion)
}

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
	"slices"
	"testing"

	"github.com/tidwall/gjson"
)

func TestEncryptedVariables(t *testing.T) {
	res := gjson.Parse(`{"devices":[
		{"device-id":"C8K-A","variables":[{"name":"aaa_password","value":"$CRYPT_CLUSTER$abc"},{"name":"host_name","value":"r1"}]},
		{"device-id":"C8K-B","variables":[{"name":"aaa_password","value":"$crypt_cluster$def"}]},
		{"device-id":"C8K-OTHER-STATE","variables":[{"name":"aaa_password","value":"$CRYPT_CLUSTER$x"}]}]}`)
	got := EncryptedVariables(res, []string{"C8K-A", "C8K-B"})
	if len(got) != 2 || !slices.Equal(got["C8K-A"], []string{"aaa_password"}) || !slices.Equal(got["C8K-B"], []string{"aaa_password"}) {
		t.Errorf("unexpected: %v", got)
	}
}

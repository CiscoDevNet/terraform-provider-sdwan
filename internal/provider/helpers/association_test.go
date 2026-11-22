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

func TestAssociatedIds(t *testing.T) {
	res := gjson.Parse(`{"devices":[{"id":"C8K-A","configGroupUpToDate":true},{"id":"C8K-B","groupTopologyLabel":"primary"}]}`)
	ids := AssociatedIds(res)
	if len(ids) != 2 || !ids["C8K-A"] || !ids["C8K-B"] || ids["C8K-C"] {
		t.Errorf("unexpected ids: %v", ids)
	}
	if got := AssociatedIds(gjson.Parse(`{}`)); len(got) != 0 {
		t.Errorf("empty response: got %v", got)
	}
}

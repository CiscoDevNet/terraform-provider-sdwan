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
	"encoding/json"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
)

// importEncryptedKey is the private state key holding, per device, the variables the Manager
// returned encrypted ($CRYPT_CLUSTER) at import. Read can't keep their value (there is no earlier
// state), so the first plan sees the clear-text config value as a change; the first Update after
// the import must not redeploy for that alone.
const importEncryptedKey = "import_encrypted"

// EncryptedVariables returns, per device ID in ids, the names of variables whose value in a
// GET .../device/variables/ response is encrypted.
func EncryptedVariables(res sdwan.Res, ids []string) map[string][]string {
	out := map[string][]string{}
	res.Get("devices").ForEach(func(_, d gjson.Result) bool {
		id := d.Get("device-id").String()
		if !Contains(ids, id) {
			return true
		}
		d.Get("variables").ForEach(func(_, v gjson.Result) bool {
			if strings.Contains(strings.ToLower(v.Get("value").String()), "$crypt_cluster") {
				out[id] = append(out[id], v.Get("name").String())
			}
			return true
		})
		return true
	})
	return out
}

// GetKeyer is something like ReadRequest.Private or UpdateRequest.Private.
type GetKeyer interface {
	GetKey(ctx context.Context, key string) ([]byte, diag.Diagnostics)
}

// GetImportEncrypted returns the encrypted-at-import variables recorded in private state.
func GetImportEncrypted(ctx context.Context, gk GetKeyer) (map[string][]string, diag.Diagnostics) {
	b, diags := gk.GetKey(ctx, importEncryptedKey)
	out := map[string][]string{}
	if len(b) > 0 {
		if err := json.Unmarshal(b, &out); err != nil {
			diags.AddWarning("Invalid private state", "Ignoring unreadable import_encrypted map: "+err.Error())
			return map[string][]string{}, diags
		}
	}
	return out, diags
}

// SetImportEncrypted stores the map; an empty map clears it.
func SetImportEncrypted(ctx context.Context, sk SetKeyer, m map[string][]string) diag.Diagnostics {
	if len(m) == 0 {
		return sk.SetKey(ctx, importEncryptedKey, nil)
	}
	b, _ := json.Marshal(m)
	return sk.SetKey(ctx, importEncryptedKey, b)
}

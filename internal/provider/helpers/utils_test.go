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

func TestIsEncryptedValue(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{"full masked secret", "$CRYPT_CLUSTER$abc==$def==", true},
		{"bare marker", "$CRYPT_CLUSTER", true},
		{"lowercase marker", "$crypt_cluster$abc==", true},
		{"mixed case marker", "$Crypt_Cluster$abc==", true},
		{"leading whitespace", "   $CRYPT_CLUSTER$abc==", true},
		{"trailing whitespace", "$CRYPT_CLUSTER$abc==   ", true},
		{"leading and trailing whitespace", "\t\n $CRYPT_CLUSTER$abc== \n", true},
		// The Other TrustSec parcel returns this shorter form. Observed against a
		// live Manager 20.18.4; the original $CRYPT_CLUSTER-only prefix missed it,
		// which would have written ciphertext into state on import.
		{"short crypt form", "$CRYPT$qlawtUpSzXQ=", true},
		{"short crypt form lowercase", "$crypt$qlawtUpSzXQ=", true},
		{"short crypt bare marker", "$CRYPT", true},
		{"cleartext password", "cisco123", false},
		{"empty string", "", false},
		{"whitespace only", "   ", false},
		{"marker not at start", "not$CRYPT_CLUSTER", false},
		{"marker embedded", "abc $CRYPT_CLUSTER$def", false},
		// Prefix-match semantics: anything starting with "$CRYPT" is treated as
		// ciphertext. Pinned so a future regex rewrite does not silently narrow it.
		{"partial marker", "$CRYPT_CLUS", true},
		{"similar prefix", "$CRYPT_CLUSTERED", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEncryptedValue(tt.value); got != tt.expected {
				t.Errorf("IsEncryptedValue(%q) = %v, want %v", tt.value, got, tt.expected)
			}
		})
	}
}

func TestJSONEqual(t *testing.T) {
	tests := []struct {
		name     string
		a, b     string
		expected bool
	}{
		{"identical strings", `{"a":1}`, `{"a":1}`, true},
		{"key order differs", `{"a":1,"b":2}`, `{"b":2,"a":1}`, true},
		{"whitespace differs", `{"a": 1}`, `{"a":1}`, true},
		{"nested key order differs", `{"x":{"p":1,"q":2}}`, `{"x":{"q":2,"p":1}}`, true},
		// The migration case: the secret moves spelling but lands on the same path.
		{"same secret via either spelling", `{"password":{"optionType":"global","value":"s"}}`, `{"password":{"value":"s","optionType":"global"}}`, true},
		{"different value", `{"a":1}`, `{"a":2}`, false},
		{"rotated secret", `{"password":{"value":"old"}}`, `{"password":{"value":"new"}}`, false},
		{"extra key", `{"a":1}`, `{"a":1,"b":2}`, false},
		{"both empty", "", "", true},
		// Unparseable input must fall through to the request, never be skipped.
		{"invalid left", `{`, `{"a":1}`, false},
		{"invalid right", `{"a":1}`, `{`, false},
		{"both invalid but unequal", `{`, `}`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JSONEqual(tt.a, tt.b); got != tt.expected {
				t.Errorf("JSONEqual(%q, %q) = %v, want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

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

package provider

import (
	"context"
	"fmt"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-sdwan"
)

type WANEdgeCertificatePush struct {
	Id      types.String `tfsdk:"id"`
	Version types.Int64  `tfsdk:"version"`
}

// push sends the current WAN edge list to all controllers and waits for the action to finish.
func (data WANEdgeCertificatePush) push(ctx context.Context, client *sdwan.Client, taskTimeout *int64) error {
	res, err := client.Post("/certificate/vedge/list?action=push", "{}")
	if err != nil {
		return fmt.Errorf("%s, %s", err, res.String())
	}
	actionId := res.Get("sendToControllerId").String()
	if actionId == "" {
		actionId = res.Get("id").String()
	}
	if actionId == "" {
		return fmt.Errorf("certificate push returned no action ID: %s", res.String())
	}
	err, _ = helpers.WaitForActionToComplete(ctx, client, actionId, taskTimeout)
	return err
}

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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-sdwan"
)

type ActivateTopologyGroup struct {
	Id      types.String `tfsdk:"id"`
	Version types.Int64  `tfsdk:"version"`
}

func (data ActivateTopologyGroup) getPath() string {
	return "/v1/topology-group/"
}

func (data ActivateTopologyGroup) activateTopologyGroup(ctx context.Context, client *sdwan.Client, taskTimeout *int64) error {
	res, err := client.Post("/v1/topology-group/"+data.Id.ValueString()+"/device/deploy", "{}")
	if err != nil {
		return err
	}
	actionId := res.Get("parentTaskId").String()
	err, _ = helpers.WaitForActionToComplete(ctx, client, actionId, taskTimeout)
	if err != nil {
		return err
	}
	return nil
}

func (data ActivateTopologyGroup) deactivateTopologyGroup(ctx context.Context, client *sdwan.Client, taskTimeout *int64) error {
	res, err := client.Post("/v1/topology-group/"+data.Id.ValueString()+"/device/deploy", `{"deactivateTopology":true}`)
	if err != nil {
		return err
	}
	actionId := res.Get("parentTaskId").String()
	err, _ = helpers.WaitForActionToComplete(ctx, client, actionId, taskTimeout)
	if err != nil {
		return err
	}
	return nil
}

func (data *ActivateTopologyGroup) processImport(ctx context.Context) {
	data.Version = types.Int64Value(0)
}

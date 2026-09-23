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
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-sdwan"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type WANEdgeCertificate struct {
	Id            types.String `tfsdk:"id"`
	ChassisNumber types.String `tfsdk:"chassis_number"`
	SerialNumber  types.String `tfsdk:"serial_number"`
	Validity      types.String `tfsdk:"validity"`
	Version       types.Int64  `tfsdk:"version"`
}

// getCertificate returns the WAN edge entry of the certificate list matching the chassis number.
func (data WANEdgeCertificate) getCertificate(ctx context.Context, client *sdwan.Client) (gjson.Result, bool, error) {
	res, err := client.Get("/certificate/vedge/list")
	if err != nil {
		return gjson.Result{}, false, err
	}
	var entry gjson.Result
	found := false
	res.Get("data").ForEach(func(_, v gjson.Result) bool {
		if v.Get("chasisNumber").String() == data.ChassisNumber.ValueString() {
			entry = v
			found = true
			return false
		}
		return true
	})
	return entry, found, nil
}

func (data WANEdgeCertificate) toBody(ctx context.Context, validity string) string {
	body := `[]`
	body, _ = sjson.Set(body, "0.chasisNumber", data.ChassisNumber.ValueString())
	body, _ = sjson.Set(body, "0.validity", validity)
	return body
}

func (data WANEdgeCertificate) legacyBody(ctx context.Context, validity string) string {
	body := `[]`
	body, _ = sjson.Set(body, "0.chasisNumber", data.ChassisNumber.ValueString())
	body, _ = sjson.Set(body, "0.validity", validity)
	return body
}

// setValidity stores the certificate validity. It does not push the updated WAN edge list to
// the controllers; use the sdwan_send_wan_edge_list_to_controllers resource for that.
func (data WANEdgeCertificate) setValidity(ctx context.Context, client *sdwan.Client, validity string) error {
	res, err := client.Post("/certificate/save/vedge/list", data.toBody(ctx, validity))
	if err != nil {
		if isSendToControllersUnsupported(err) {
			return data.setValidityLegacy(ctx, client, validity)
		}
		return fmt.Errorf("%s, %s", err, res.String())
	}
	return nil
}

func (data WANEdgeCertificate) setValidityLegacy(ctx context.Context, client *sdwan.Client, validity string) error {
	res, err := client.Post("/certificate/save/vedge/list", data.legacyBody(ctx, validity))
	if err != nil {
		return fmt.Errorf("%s, %s", err, res.String())
	}
	return nil
}

// isSendToControllersUnsupported reports whether Manager rejected the sendToControllers field
// in the certificate save payload (older Manager versions do not support it).
func isSendToControllersUnsupported(err error) bool {
	message := strings.ToLower(err.Error())
	return strings.Contains(message, "sendtocontrollers") &&
		(strings.Contains(message, "impermissible") || strings.Contains(message, "incorrect") || strings.Contains(message, "unknown") || strings.Contains(message, "invalid") || strings.Contains(message, "unsupported") || strings.Contains(message, "unrecognized"))
}

func (data *WANEdgeCertificate) fromBody(ctx context.Context, res gjson.Result) {
	data.Id = types.StringValue(res.Get("chasisNumber").String())
	data.ChassisNumber = types.StringValue(res.Get("chasisNumber").String())
	if value := res.Get("serialNumber"); value.Exists() {
		data.SerialNumber = types.StringValue(value.String())
	}
	if value := res.Get("validity"); value.Exists() {
		data.Validity = types.StringValue(value.String())
	}
}

func (data *WANEdgeCertificate) processImport(ctx context.Context) {
	if data.Version.IsNull() || data.Version.IsUnknown() {
		data.Version = types.Int64Value(1)
	}
}

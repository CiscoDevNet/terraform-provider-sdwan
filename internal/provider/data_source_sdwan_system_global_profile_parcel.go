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

// Section below is generated&owned by "gen/generator.go". //template:begin imports
import (
	"context"
	"fmt"
	"net/url"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &SystemGlobalProfileParcelDataSource{}
	_ datasource.DataSourceWithConfigure = &SystemGlobalProfileParcelDataSource{}
)

func NewSystemGlobalProfileParcelDataSource() datasource.DataSource {
	return &SystemGlobalProfileParcelDataSource{}
}

type SystemGlobalProfileParcelDataSource struct {
	client *sdwan.Client
}

func (d *SystemGlobalProfileParcelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_global_profile_parcel"
}

func (d *SystemGlobalProfileParcelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the System Global profile parcel.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the profile parcel",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the profile parcel",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the profile parcel",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the profile parcel",
				Computed:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: "Feature Profile ID",
				Required:            true,
			},
			"http_server": schema.BoolAttribute{
				MarkdownDescription: "Set a HTTP Server",
				Computed:            true,
			},
			"http_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"https_server": schema.BoolAttribute{
				MarkdownDescription: "Set a HTTPS Server",
				Computed:            true,
			},
			"https_server_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ftp_passive": schema.BoolAttribute{
				MarkdownDescription: "Set Passive FTP",
				Computed:            true,
			},
			"ftp_passive_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"domain_lookup": schema.BoolAttribute{
				MarkdownDescription: "Configure Domain-Lookup",
				Computed:            true,
			},
			"domain_lookup_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"arp_proxy": schema.BoolAttribute{
				MarkdownDescription: "Set ARP Proxy",
				Computed:            true,
			},
			"arp_proxy_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"rsh_rcp": schema.BoolAttribute{
				MarkdownDescription: "Set RSH/RCP",
				Computed:            true,
			},
			"rsh_rcp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"line_vty": schema.BoolAttribute{
				MarkdownDescription: "Configure Telnet (Outbound)",
				Computed:            true,
			},
			"line_vty_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"cdp": schema.BoolAttribute{
				MarkdownDescription: "Configure CDP",
				Computed:            true,
			},
			"cdp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"lldp": schema.BoolAttribute{
				MarkdownDescription: "Configure LLDP",
				Computed:            true,
			},
			"lldp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"source_interface": schema.StringAttribute{
				MarkdownDescription: "Specify interface for source address in all HTTP(S) client connections",
				Computed:            true,
			},
			"source_interface_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tcp_keepalives_in": schema.BoolAttribute{
				MarkdownDescription: "Configure tcp-keepalives-in",
				Computed:            true,
			},
			"tcp_keepalives_in_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tcp_keepalives_out": schema.BoolAttribute{
				MarkdownDescription: "Configure tcp-keepalives-out",
				Computed:            true,
			},
			"tcp_keepalives_out_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"tcp_small_servers": schema.BoolAttribute{
				MarkdownDescription: "Configure tcp-small-servers",
				Computed:            true,
			},
			"tcp_small_servers_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"udp_small_servers": schema.BoolAttribute{
				MarkdownDescription: "Configure udp-small-servers",
				Computed:            true,
			},
			"udp_small_servers_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"console_logging": schema.BoolAttribute{
				MarkdownDescription: "Configure Console Logging",
				Computed:            true,
			},
			"console_logging_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ip_source_routing": schema.BoolAttribute{
				MarkdownDescription: "Set Source Route",
				Computed:            true,
			},
			"ip_source_routing_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"vty_line_logging": schema.BoolAttribute{
				MarkdownDescription: "Configure VTY Line Logging",
				Computed:            true,
			},
			"vty_line_logging_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"snmp_ifindex_persist": schema.BoolAttribute{
				MarkdownDescription: "Configure SNMP Ifindex Persist",
				Computed:            true,
			},
			"snmp_ifindex_persist_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ignore_bootp": schema.BoolAttribute{
				MarkdownDescription: "Configure Ignore BOOTP",
				Computed:            true,
			},
			"ignore_bootp_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat64_udp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Set NAT64 UDP session timeout, in seconds",
				Computed:            true,
			},
			"nat64_udp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"nat64_tcp_timeout": schema.Int64Attribute{
				MarkdownDescription: "Set NAT64 TCP session timeout, in seconds",
				Computed:            true,
			},
			"nat64_tcp_timeout_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"http_authentication": schema.StringAttribute{
				MarkdownDescription: "Set preference for HTTP Authentication",
				Computed:            true,
			},
			"http_authentication_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
			"ssh_version": schema.StringAttribute{
				MarkdownDescription: "Set SSH version",
				Computed:            true,
			},
			"ssh_version_variable": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Variable name").String,
				Computed:            true,
			},
		},
	}
}

func (d *SystemGlobalProfileParcelDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (d *SystemGlobalProfileParcelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config SystemGlobal

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get(config.getPath() + "/" + url.QueryEscape(config.Id.ValueString()))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Name.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end read

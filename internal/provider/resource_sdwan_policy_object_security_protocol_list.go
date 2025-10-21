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
	"strings"
	"sync"

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// End of section. //template:end imports

// Section below is generated&owned by "gen/generator.go". //template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &PolicyObjectSecurityProtocolListProfileParcelResource{}
var _ resource.ResourceWithImportState = &PolicyObjectSecurityProtocolListProfileParcelResource{}

func NewPolicyObjectSecurityProtocolListProfileParcelResource() resource.Resource {
	return &PolicyObjectSecurityProtocolListProfileParcelResource{}
}

type PolicyObjectSecurityProtocolListProfileParcelResource struct {
	client      *sdwan.Client
	updateMutex *sync.Mutex
}

func (r *PolicyObjectSecurityProtocolListProfileParcelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policy_object_security_protocol_list"
}

func (r *PolicyObjectSecurityProtocolListProfileParcelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Policy Object Security Protocol List Policy_object.").AddMinimumVersionDescription("20.12.0").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the Policy_object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the Policy_object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the Policy_object",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the Policy_object",
				Optional:            true,
			},
			"feature_profile_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Feature Profile ID").String,
				Required:            true,
			},
			"entries": schema.ListNestedAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"protocol_name": schema.StringAttribute{
							MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("snmp", "icmp", "tcp", "udp", "echo", "telnet", "wins", "n2h2server", "nntp", "pptp", "rtsp", "bootpc", "gdoi", "tacacs", "gopher", "icabrowser", "skinny", "sunrpc", "biff", "router", "ircs", "orasrv", "ms-cluster-net", "kermit", "isakmp", "sshell", "realsecure", "ircu", "appleqtc", "pwdgen", "rdb-dbs-disp", "creativepartnr", "finger", "ftps", "giop", "rsvd", "hp-alarm-mgr", "uucp", "kerberos", "imap", "time", "bootps", "tftp", "oracle", "snmptrap", "http", "qmtp", "radius", "oracle-em-vp", "tarantella", "pcanywheredata", "ldap", "mgcp", "sqlsrv", "hsrp", "cisco-net-mgmt", "smtp", "pcanywherestat", "exec", "send", "stun", "syslog", "ms-sql-m", "citrix", "creativeserver", "cifs", "cisco-sys", "cisco-tna", "ms-dotnetster", "gtpv1", "gtpv0", "imap3", "fcip-port", "netbios-dgm", "sip-tls", "pop3s", "cisco-fna", "802-11-iapp", "oem-agent", "cisco-tdp", "tr-rsrb", "r-winsock", "sql-net", "syslog-conn", "tacacs-ds", "h225ras", "ace-svr", "dhcp-failover", "igmpv3lite", "irc-serv", "entrust-svcs", "dbcontrol_agent", "cisco-svcs", "ipsec-msft", "microsoft-ds", "ms-sna", "rsvp_tunnel", "rsvp-encap", "hp-collector", "netbios-ns", "msexch-routing", "h323", "l2tp", "ldap-admin", "pop3", "h323callsigalt", "ms-sql", "iscsi-target", "webster", "lotusnote", "ipx", "entrust-svc-hand", "citriximaclient", "rtc-pm-port", "ftp", "aol", "xdmcp", "oraclenames", "login", "iscsi", "ttc", "imaps", "socks", "ssh", "dnsix", "daytime", "sip", "discard", "ntp", "ldaps", "https", "vdolive", "ica", "net8-cman", "cuseeme", "netstat", "sms", "streamworks", "rtelnet", "who", "kazaa", "ssp", "dbase", "timed", "cddbp", "telnets", "ymsgr", "ident", "bgp", "ddns-v3", "vqp", "irc", "ipass", "x11", "dns", "lotusmtap", "mysql", "nfs", "msnmsgr", "netshow", "sqlserv", "hp-managed-node", "ncp", "shell", "realmedia", "msrpc", "clp").String,
							Optional:            true,
							Validators: []validator.String{
								stringvalidator.OneOf("snmp", "icmp", "tcp", "udp", "echo", "telnet", "wins", "n2h2server", "nntp", "pptp", "rtsp", "bootpc", "gdoi", "tacacs", "gopher", "icabrowser", "skinny", "sunrpc", "biff", "router", "ircs", "orasrv", "ms-cluster-net", "kermit", "isakmp", "sshell", "realsecure", "ircu", "appleqtc", "pwdgen", "rdb-dbs-disp", "creativepartnr", "finger", "ftps", "giop", "rsvd", "hp-alarm-mgr", "uucp", "kerberos", "imap", "time", "bootps", "tftp", "oracle", "snmptrap", "http", "qmtp", "radius", "oracle-em-vp", "tarantella", "pcanywheredata", "ldap", "mgcp", "sqlsrv", "hsrp", "cisco-net-mgmt", "smtp", "pcanywherestat", "exec", "send", "stun", "syslog", "ms-sql-m", "citrix", "creativeserver", "cifs", "cisco-sys", "cisco-tna", "ms-dotnetster", "gtpv1", "gtpv0", "imap3", "fcip-port", "netbios-dgm", "sip-tls", "pop3s", "cisco-fna", "802-11-iapp", "oem-agent", "cisco-tdp", "tr-rsrb", "r-winsock", "sql-net", "syslog-conn", "tacacs-ds", "h225ras", "ace-svr", "dhcp-failover", "igmpv3lite", "irc-serv", "entrust-svcs", "dbcontrol_agent", "cisco-svcs", "ipsec-msft", "microsoft-ds", "ms-sna", "rsvp_tunnel", "rsvp-encap", "hp-collector", "netbios-ns", "msexch-routing", "h323", "l2tp", "ldap-admin", "pop3", "h323callsigalt", "ms-sql", "iscsi-target", "webster", "lotusnote", "ipx", "entrust-svc-hand", "citriximaclient", "rtc-pm-port", "ftp", "aol", "xdmcp", "oraclenames", "login", "iscsi", "ttc", "imaps", "socks", "ssh", "dnsix", "daytime", "sip", "discard", "ntp", "ldaps", "https", "vdolive", "ica", "net8-cman", "cuseeme", "netstat", "sms", "streamworks", "rtelnet", "who", "kazaa", "ssp", "dbase", "timed", "cddbp", "telnets", "ymsgr", "ident", "bgp", "ddns-v3", "vqp", "irc", "ipass", "x11", "dns", "lotusmtap", "mysql", "nfs", "msnmsgr", "netshow", "sqlserv", "hp-managed-node", "ncp", "shell", "realmedia", "msrpc", "clp"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *PolicyObjectSecurityProtocolListProfileParcelResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*SdwanProviderData).Client
	r.updateMutex = req.ProviderData.(*SdwanProviderData).UpdateMutex
}

// End of section. //template:end model

// Section below is generated&owned by "gen/generator.go". //template:begin create
func (r *PolicyObjectSecurityProtocolListProfileParcelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PolicyObjectSecurityProtocolList

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Name.ValueString()))

	// Create object
	body := plan.toBody(ctx)

	res, err := r.client.Post(plan.getPath(), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = types.StringValue(res.Get("parcelId").String())
	plan.Version = types.Int64Value(0)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end create

// Section below is generated&owned by "gen/generator.go". //template:begin read
func (r *PolicyObjectSecurityProtocolListProfileParcelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PolicyObjectSecurityProtocolList

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Name.String()))

	res, err := r.client.Get(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if res.Get("error.message").String() == "Invalid feature Id" {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	state.fromBody(ctx, res)
	if state.Version.IsNull() {
		state.Version = types.Int64Value(0)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Name.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)

	helpers.SetFlagImporting(ctx, false, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end read

// Section below is generated&owned by "gen/generator.go". //template:begin update
func (r *PolicyObjectSecurityProtocolListProfileParcelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PolicyObjectSecurityProtocolList

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Name.ValueString()))

	body := plan.toBody(ctx)
	res, err := r.client.Put(plan.getPath()+"/"+url.QueryEscape(plan.Id.ValueString()), body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	plan.Version = types.Int64Value(state.Version.ValueInt64() + 1)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Name.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

// End of section. //template:end update

// Section below is generated&owned by "gen/generator.go". //template:begin delete
func (r *PolicyObjectSecurityProtocolListProfileParcelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PolicyObjectSecurityProtocolList

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Name.ValueString()))

	res, err := r.client.Delete(state.getPath() + "/" + url.QueryEscape(state.Id.ValueString()))
	if err != nil && res.Get("error.message").String() != "Invalid Template Id" {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Name.ValueString()))

	resp.State.RemoveResource(ctx)
}

// End of section. //template:end delete

// Section below is generated&owned by "gen/generator.go". //template:begin import
func (r *PolicyObjectSecurityProtocolListProfileParcelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	count := 1
	parts := strings.SplitN(req.ID, ",", (count + 1))

	pattern := "policy_object_security_protocol_list_id" + ",feature_profile_id"
	if len(parts) != (count + 1) {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier", fmt.Sprintf("Expected import identifier with the format: %s. Got: %q", pattern, req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), parts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("feature_profile_id"), parts[1])...)

	helpers.SetFlagImporting(ctx, true, resp.Private, &resp.Diagnostics)
}

// End of section. //template:end import

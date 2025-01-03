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

// Section below is generated&owned by "gen/generator.go". //template:begin provider
import (
	"context"
	"os"
	"strconv"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-sdwan"
)

// SdwanProvider defines the provider implementation.
type SdwanProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// SdwanProviderModel describes the provider data model.
type SdwanProviderModel struct {
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	URL      types.String `tfsdk:"url"`
	Insecure types.Bool   `tfsdk:"insecure"`
	Retries  types.Int64  `tfsdk:"retries"`
}

// SdwanProviderData describes the data maintained by the provider.
type SdwanProviderData struct {
	Client      *sdwan.Client
	UpdateMutex *sync.Mutex
}

// Metadata returns the provider type name.
func (p *SdwanProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "sdwan"
	resp.Version = p.version
}

func (p *SdwanProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				MarkdownDescription: "Username for the SD-WAN Manager account. This can also be set as the `SDWAN_USERNAME` environment variable.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the SD-WAN Manager account. This can also be set as the `SDWAN_PASSWORD` environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "URL of the Cisco SD-WAN Manager device. This can also be set as the `SDWAN_URL` environment variable.",
				Optional:            true,
			},
			"insecure": schema.BoolAttribute{
				MarkdownDescription: "Allow insecure HTTPS client. This can also be set as the `SDWAN_INSECURE` environment variable. Defaults to `true`.",
				Optional:            true,
			},
			"retries": schema.Int64Attribute{
				MarkdownDescription: "Number of retries for REST API calls. This can also be set as the `SDWAN_RETRIES` environment variable. Defaults to `3`.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 9),
				},
			},
		},
	}
}

func (p *SdwanProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config SdwanProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a username to the provider
	var username string
	if config.Username.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}

	if config.Username.IsNull() {
		username = os.Getenv("SDWAN_USERNAME")
	} else {
		username = config.Username.ValueString()
	}

	if username == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find username",
			"Username cannot be an empty string",
		)
		return
	}

	// User must provide a password to the provider
	var password string
	if config.Password.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Password.IsNull() {
		password = os.Getenv("SDWAN_PASSWORD")
	} else {
		password = config.Password.ValueString()
	}

	if password == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find password",
			"Password cannot be an empty string",
		)
		return
	}

	// User must provide a username to the provider
	var url string
	if config.URL.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as url",
		)
		return
	}

	if config.URL.IsNull() {
		url = os.Getenv("SDWAN_URL")
	} else {
		url = config.URL.ValueString()
	}

	if url == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find url",
			"URL cannot be an empty string",
		)
		return
	}

	var insecure bool
	if config.Insecure.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as insecure",
		)
		return
	}

	if config.Insecure.IsNull() {
		insecureStr := os.Getenv("SDWAN_INSECURE")
		if insecureStr == "" {
			insecure = true
		} else {
			insecure, _ = strconv.ParseBool(insecureStr)
		}
	} else {
		insecure = config.Insecure.ValueBool()
	}

	var retries int64
	if config.Retries.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as retries",
		)
		return
	}

	if config.Retries.IsNull() {
		retriesStr := os.Getenv("SDWAN_RETRIES")
		if retriesStr == "" {
			retries = 3
		} else {
			retries, _ = strconv.ParseInt(retriesStr, 0, 64)
		}
	} else {
		retries = config.Retries.ValueInt64()
	}

	// Create a new NX-OS client and set it to the provider client
	c, err := sdwan.NewClient(url, username, password, insecure, sdwan.MaxRetries(int(retries)))
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create sdwan client:\n\n"+err.Error(),
		)
		return
	}

	data := SdwanProviderData{Client: &c, UpdateMutex: &sync.Mutex{}}
	resp.DataSourceData = &data
	resp.ResourceData = &data
}

func (p *SdwanProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewCEdgeAAAFeatureTemplateResource,
		NewCEdgeGlobalFeatureTemplateResource,
		NewCEdgeIGMPFeatureTemplateResource,
		NewCEdgeMulticastFeatureTemplateResource,
		NewCEdgePIMFeatureTemplateResource,
		NewCellularCEdgeProfileFeatureTemplateResource,
		NewCellularControllerFeatureTemplateResource,
		NewCellularProfileFeatureTemplateResource,
		NewCiscoBannerFeatureTemplateResource,
		NewCiscoBFDFeatureTemplateResource,
		NewCiscoBGPFeatureTemplateResource,
		NewCiscoDHCPServerFeatureTemplateResource,
		NewCiscoLoggingFeatureTemplateResource,
		NewCiscoNTPFeatureTemplateResource,
		NewCiscoOMPFeatureTemplateResource,
		NewCiscoOSPFFeatureTemplateResource,
		NewCiscoOSPFv3FeatureTemplateResource,
		NewCiscoSecureInternetGatewayFeatureTemplateResource,
		NewCiscoSecurityFeatureTemplateResource,
		NewCiscoSIGCredentialsFeatureTemplateResource,
		NewCiscoSNMPFeatureTemplateResource,
		NewCiscoSystemFeatureTemplateResource,
		NewCiscoThousandEyesFeatureTemplateResource,
		NewCiscoTrustSecFeatureTemplateResource,
		NewCiscoVPNFeatureTemplateResource,
		NewCiscoVPNInterfaceFeatureTemplateResource,
		NewCiscoVPNInterfaceGREFeatureTemplateResource,
		NewCiscoVPNInterfaceIPSecFeatureTemplateResource,
		NewCiscoWirelessLANFeatureTemplateResource,
		NewCLITemplateFeatureTemplateResource,
		NewEigrpFeatureTemplateResource,
		NewGpsFeatureTemplateResource,
		NewSwitchportFeatureTemplateResource,
		NewSecurityAppHostingFeatureTemplateResource,
		NewVPNInterfaceCellularFeatureTemplateResource,
		NewVPNInterfaceEthernetPPPoEFeatureTemplateResource,
		NewVPNInterfaceDSLIPoEFeatureTemplateResource,
		NewVPNInterfaceMultilinkFeatureTemplateResource,
		NewVPNInterfaceDSLPPPoAFeatureTemplateResource,
		NewVPNInterfaceDSLPPPoEFeatureTemplateResource,
		NewVPNInterfaceSVIFeatureTemplateResource,
		NewVPNInterfaceT1E1SerialFeatureTemplateResource,
		NewApplicationPriorityQoSProfileParcelResource,
		NewApplicationPriorityTrafficPolicyProfileParcelResource,
		NewOtherThousandEyesProfileParcelResource,
		NewOtherUCSEProfileParcelResource,
		NewPolicyObjectAppProbeClassProfileParcelResource,
		NewPolicyObjectApplicationListProfileParcelResource,
		NewPolicyObjectASPathListProfileParcelResource,
		NewPolicyObjectClassMapProfileParcelResource,
		NewPolicyObjectColorListProfileParcelResource,
		NewPolicyObjectDataIPv4PrefixListProfileParcelResource,
		NewPolicyObjectDataIPv6PrefixListProfileParcelResource,
		NewPolicyObjectExpandedCommunityListProfileParcelResource,
		NewPolicyObjectExtendedCommunityListProfileParcelResource,
		NewPolicyObjectIPv4PrefixListProfileParcelResource,
		NewPolicyObjectIPv6PrefixListProfileParcelResource,
		NewPolicyObjectMirrorProfileParcelResource,
		NewPolicyObjectPolicerProfileParcelResource,
		NewPolicyObjectPreferredColorGroupProfileParcelResource,
		NewPolicyObjectSecurityDataIPv4PrefixListProfileParcelResource,
		NewPolicyObjectSecurityFQDNListProfileParcelResource,
		NewPolicyObjectSecurityGeolocationListProfileParcelResource,
		NewPolicyObjectSecurityIdentityListProfileParcelResource,
		NewPolicyObjectSecurityIPSSignatureProfileParcelResource,
		NewPolicyObjectSecurityLocalApplicationListProfileParcelResource,
		NewPolicyObjectSecurityLocalDomainListProfileParcelResource,
		NewPolicyObjectSecurityPortListProfileParcelResource,
		NewPolicyObjectSecurityScalableGroupTagListProfileParcelResource,
		NewPolicyObjectSecurityURLAllowListProfileParcelResource,
		NewPolicyObjectSecurityURLBlockListProfileParcelResource,
		NewPolicyObjectSLAClassListProfileParcelResource,
		NewPolicyObjectStandardCommunityListProfileParcelResource,
		NewPolicyObjectTLOCListProfileParcelResource,
		NewPolicyObjectVPNGroupProfileParcelResource,
		NewServiceDHCPServerProfileParcelResource,
		NewServiceIPv4ACLProfileParcelResource,
		NewServiceIPv6ACLProfileParcelResource,
		NewServiceLANVPNProfileParcelResource,
		NewServiceLANVPNInterfaceEthernetProfileParcelResource,
		NewServiceLANVPNInterfaceGREProfileParcelResource,
		NewServiceLANVPNInterfaceIPSecProfileParcelResource,
		NewServiceLANVPNInterfaceSVIProfileParcelResource,
		NewServiceMulticastProfileParcelResource,
		NewServiceObjectTrackerProfileParcelResource,
		NewServiceObjectTrackerGroupProfileParcelResource,
		NewServiceRoutePolicyProfileParcelResource,
		NewServiceRoutingBGPProfileParcelResource,
		NewServiceRoutingEIGRPProfileParcelResource,
		NewServiceRoutingOSPFProfileParcelResource,
		NewServiceRoutingOSPFv3IPv4ProfileParcelResource,
		NewServiceRoutingOSPFv3IPv6ProfileParcelResource,
		NewServiceSwitchportProfileParcelResource,
		NewServiceTrackerProfileParcelResource,
		NewServiceTrackerGroupProfileParcelResource,
		NewServiceWirelessLANProfileParcelResource,
		NewSystemAAAProfileParcelResource,
		NewSystemBannerProfileParcelResource,
		NewSystemBasicProfileParcelResource,
		NewSystemBFDProfileParcelResource,
		NewSystemFlexiblePortSpeedProfileParcelResource,
		NewSystemGlobalProfileParcelResource,
		NewSystemIPv4DeviceAccessProfileParcelResource,
		NewSystemIPv6DeviceAccessProfileParcelResource,
		NewSystemLoggingProfileParcelResource,
		NewSystemMRFProfileParcelResource,
		NewSystemNTPProfileParcelResource,
		NewSystemOMPProfileParcelResource,
		NewSystemPerformanceMonitoringProfileParcelResource,
		NewSystemRemoteAccessProfileParcelResource,
		NewSystemSecurityProfileParcelResource,
		NewSystemSNMPProfileParcelResource,
		NewTransportCellularControllerProfileParcelResource,
		NewTransportCellularProfileProfileParcelResource,
		NewTransportGPSProfileParcelResource,
		NewTransportIPv4ACLProfileParcelResource,
		NewTransportIPv6ACLProfileParcelResource,
		NewTransportIPv6TrackerProfileParcelResource,
		NewTransportIPv6TrackerGroupProfileParcelResource,
		NewTransportManagementVPNProfileParcelResource,
		NewTransportManagementVPNInterfaceEthernetProfileParcelResource,
		NewTransportRoutePolicyProfileParcelResource,
		NewTransportRoutingBGPProfileParcelResource,
		NewTransportRoutingOSPFProfileParcelResource,
		NewTransportRoutingOSPFv3IPv4ProfileParcelResource,
		NewTransportRoutingOSPFv3IPv6ProfileParcelResource,
		NewTransportT1E1ControllerProfileParcelResource,
		NewTransportTrackerProfileParcelResource,
		NewTransportTrackerGroupProfileParcelResource,
		NewTransportWANVPNProfileParcelResource,
		NewTransportWANVPNInterfaceCellularProfileParcelResource,
		NewTransportWANVPNInterfaceEthernetProfileParcelResource,
		NewTransportWANVPNInterfaceGREProfileParcelResource,
		NewTransportWANVPNInterfaceIPSECProfileParcelResource,
		NewTransportWANVPNInterfaceT1E1SerialProfileParcelResource,
		NewAdvancedInspectionProfilePolicyDefinitionResource,
		NewAdvancedMalwareProtectionPolicyDefinitionResource,
		NewAllowURLListPolicyObjectResource,
		NewAppProbeClassPolicyObjectResource,
		NewApplicationAwareRoutingPolicyDefinitionResource,
		NewApplicationListPolicyObjectResource,
		NewApplicationPriorityFeatureProfileResource,
		NewASPathListPolicyObjectResource,
		NewBlockURLListPolicyObjectResource,
		NewCentralizedPolicyResource,
		NewCflowdPolicyDefinitionResource,
		NewClassMapPolicyObjectResource,
		NewCLIConfigFeatureResource,
		NewCLIDeviceTemplateResource,
		NewCLIFeatureProfileResource,
		NewColorListPolicyObjectResource,
		NewConfigurationGroupResource,
		NewCustomControlTopologyPolicyDefinitionResource,
		NewDataFQDNPrefixListPolicyObjectResource,
		NewDataIPv4PrefixListPolicyObjectResource,
		NewDataIPv6PrefixListPolicyObjectResource,
		NewDNSSecurityFeatureProfileResource,
		NewDNSSecurityPolicyDefinitionResource,
		NewDomainListPolicyObjectResource,
		NewEmbeddedSecurityFeatureProfileResource,
		NewExpandedCommunityListPolicyObjectResource,
		NewExtendedCommunityListPolicyObjectResource,
		NewFeatureDeviceTemplateResource,
		NewGeoLocationListPolicyObjectResource,
		NewHubAndSpokeTopologyPolicyDefinitionResource,
		NewIntrusionPreventionPolicyDefinitionResource,
		NewIPSSignatureListPolicyObjectResource,
		NewIPv4ACLPolicyDefinitionResource,
		NewIPv4DeviceACLPolicyDefinitionResource,
		NewIPv4PrefixListPolicyObjectResource,
		NewIPv6ACLPolicyDefinitionResource,
		NewIPv6DeviceACLPolicyDefinitionResource,
		NewIPv6PrefixListPolicyObjectResource,
		NewLocalApplicationListPolicyObjectResource,
		NewLocalizedPolicyResource,
		NewMeshTopologyPolicyDefinitionResource,
		NewMirrorPolicyObjectResource,
		NewObjectGroupPolicyDefinitionResource,
		NewOtherFeatureProfileResource,
		NewPolicerPolicyObjectResource,
		NewPolicyObjectFeatureProfileResource,
		NewPortListPolicyObjectResource,
		NewPreferredColorGroupPolicyObjectResource,
		NewProtocolListPolicyObjectResource,
		NewQoSMapPolicyDefinitionResource,
		NewRegionListPolicyObjectResource,
		NewRewriteRulePolicyDefinitionResource,
		NewRoutePolicyDefinitionResource,
		NewRuleSetPolicyDefinitionResource,
		NewSecurityPolicyResource,
		NewServiceFeatureProfileResource,
		NewServiceLANVPNInterfaceEthernetFeatureAssociateDHCPServerFeatureResource,
		NewServiceLANVPNInterfaceIPSecFeatureAssociateDHCPServerFeatureResource,
		NewServiceLANVPNInterfaceSVIFeatureAssociateDHCPServerFeatureResource,
		NewServiceLANVPNFeatureAssociateMulticastFeatureResource,
		NewServiceLANVPNFeatureAssociateRoutingBGPFeatureResource,
		NewServiceLANVPNFeatureAssociateRoutingEIGRPFeatureResource,
		NewServiceLANVPNFeatureAssociateRoutingOSPFFeatureResource,
		NewServiceLANVPNFeatureAssociateRoutingOSPFv3IPv4FeatureResource,
		NewServiceLANVPNFeatureAssociateRoutingOSPFv3IPv6FeatureResource,
		NewSIGSecurityFeatureProfileResource,
		NewSiteListPolicyObjectResource,
		NewSLAClassPolicyObjectResource,
		NewStandardCommunityListPolicyObjectResource,
		NewSystemFeatureProfileResource,
		NewTLOCListPolicyObjectResource,
		NewTLSSSLDecryptionPolicyDefinitionResource,
		NewTLSSSLProfilePolicyDefinitionResource,
		NewTrafficDataPolicyDefinitionResource,
		NewTransportFeatureProfileResource,
		NewTransportWANVPNFeatureAssociateRoutingBGPFeatureResource,
		NewTransportWANVPNFeatureAssociateRoutingOSPFFeatureResource,
		NewTransportWANVPNFeatureAssociateRoutingOSPFv3IPv4FeatureResource,
		NewTransportWANVPNFeatureAssociateRoutingOSPFv3IPv6FeatureResource,
		NewTransportWANVPNInterfaceCellularFeatureAssociateTrackerFeatureResource,
		NewTransportWANVPNInterfaceEthernetFeatureAssociateIPv6TrackerFeatureResource,
		NewTransportWANVPNInterfaceEthernetFeatureAssociateIPv6TrackerGroupFeatureResource,
		NewTransportWANVPNInterfaceEthernetFeatureAssociateTrackerFeatureResource,
		NewTransportWANVPNInterfaceEthernetFeatureAssociateTrackerGroupFeatureResource,
		NewTransportWANVPNInterfaceGREFeatureAssociateTrackerFeatureResource,
		NewTransportWANVPNInterfaceIPSECFeatureAssociateTrackerFeatureResource,
		NewURLFilteringPolicyDefinitionResource,
		NewVPNListPolicyObjectResource,
		NewVPNMembershipPolicyDefinitionResource,
		NewZoneBasedFirewallPolicyDefinitionResource,
		NewZoneListPolicyObjectResource,
		NewAttachFeatureDeviceTemplateResource,
		NewActivateCentralizedPolicyResource,
	}
}

func (p *SdwanProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewCEdgeAAAFeatureTemplateDataSource,
		NewCEdgeGlobalFeatureTemplateDataSource,
		NewCEdgeIGMPFeatureTemplateDataSource,
		NewCEdgeMulticastFeatureTemplateDataSource,
		NewCEdgePIMFeatureTemplateDataSource,
		NewCellularCEdgeProfileFeatureTemplateDataSource,
		NewCellularControllerFeatureTemplateDataSource,
		NewCellularProfileFeatureTemplateDataSource,
		NewCiscoBannerFeatureTemplateDataSource,
		NewCiscoBFDFeatureTemplateDataSource,
		NewCiscoBGPFeatureTemplateDataSource,
		NewCiscoDHCPServerFeatureTemplateDataSource,
		NewCiscoLoggingFeatureTemplateDataSource,
		NewCiscoNTPFeatureTemplateDataSource,
		NewCiscoOMPFeatureTemplateDataSource,
		NewCiscoOSPFFeatureTemplateDataSource,
		NewCiscoOSPFv3FeatureTemplateDataSource,
		NewCiscoSecureInternetGatewayFeatureTemplateDataSource,
		NewCiscoSecurityFeatureTemplateDataSource,
		NewCiscoSIGCredentialsFeatureTemplateDataSource,
		NewCiscoSNMPFeatureTemplateDataSource,
		NewCiscoSystemFeatureTemplateDataSource,
		NewCiscoThousandEyesFeatureTemplateDataSource,
		NewCiscoTrustSecFeatureTemplateDataSource,
		NewCiscoVPNFeatureTemplateDataSource,
		NewCiscoVPNInterfaceFeatureTemplateDataSource,
		NewCiscoVPNInterfaceGREFeatureTemplateDataSource,
		NewCiscoVPNInterfaceIPSecFeatureTemplateDataSource,
		NewCiscoWirelessLANFeatureTemplateDataSource,
		NewCLITemplateFeatureTemplateDataSource,
		NewEigrpFeatureTemplateDataSource,
		NewGpsFeatureTemplateDataSource,
		NewSwitchportFeatureTemplateDataSource,
		NewSecurityAppHostingFeatureTemplateDataSource,
		NewVPNInterfaceCellularFeatureTemplateDataSource,
		NewVPNInterfaceEthernetPPPoEFeatureTemplateDataSource,
		NewVPNInterfaceDSLIPoEFeatureTemplateDataSource,
		NewVPNInterfaceMultilinkFeatureTemplateDataSource,
		NewVPNInterfaceDSLPPPoAFeatureTemplateDataSource,
		NewVPNInterfaceDSLPPPoEFeatureTemplateDataSource,
		NewVPNInterfaceSVIFeatureTemplateDataSource,
		NewVPNInterfaceT1E1SerialFeatureTemplateDataSource,
		NewApplicationPriorityQoSProfileParcelDataSource,
		NewApplicationPriorityTrafficPolicyProfileParcelDataSource,
		NewOtherThousandEyesProfileParcelDataSource,
		NewOtherUCSEProfileParcelDataSource,
		NewPolicyObjectAppProbeClassProfileParcelDataSource,
		NewPolicyObjectApplicationListProfileParcelDataSource,
		NewPolicyObjectASPathListProfileParcelDataSource,
		NewPolicyObjectClassMapProfileParcelDataSource,
		NewPolicyObjectColorListProfileParcelDataSource,
		NewPolicyObjectDataIPv4PrefixListProfileParcelDataSource,
		NewPolicyObjectDataIPv6PrefixListProfileParcelDataSource,
		NewPolicyObjectExpandedCommunityListProfileParcelDataSource,
		NewPolicyObjectExtendedCommunityListProfileParcelDataSource,
		NewPolicyObjectIPv4PrefixListProfileParcelDataSource,
		NewPolicyObjectIPv6PrefixListProfileParcelDataSource,
		NewPolicyObjectMirrorProfileParcelDataSource,
		NewPolicyObjectPolicerProfileParcelDataSource,
		NewPolicyObjectPreferredColorGroupProfileParcelDataSource,
		NewPolicyObjectSecurityDataIPv4PrefixListProfileParcelDataSource,
		NewPolicyObjectSecurityFQDNListProfileParcelDataSource,
		NewPolicyObjectSecurityGeolocationListProfileParcelDataSource,
		NewPolicyObjectSecurityIdentityListProfileParcelDataSource,
		NewPolicyObjectSecurityIPSSignatureProfileParcelDataSource,
		NewPolicyObjectSecurityLocalApplicationListProfileParcelDataSource,
		NewPolicyObjectSecurityLocalDomainListProfileParcelDataSource,
		NewPolicyObjectSecurityPortListProfileParcelDataSource,
		NewPolicyObjectSecurityScalableGroupTagListProfileParcelDataSource,
		NewPolicyObjectSecurityURLAllowListProfileParcelDataSource,
		NewPolicyObjectSecurityURLBlockListProfileParcelDataSource,
		NewPolicyObjectSLAClassListProfileParcelDataSource,
		NewPolicyObjectStandardCommunityListProfileParcelDataSource,
		NewPolicyObjectTLOCListProfileParcelDataSource,
		NewPolicyObjectVPNGroupProfileParcelDataSource,
		NewServiceDHCPServerProfileParcelDataSource,
		NewServiceIPv4ACLProfileParcelDataSource,
		NewServiceIPv6ACLProfileParcelDataSource,
		NewServiceLANVPNProfileParcelDataSource,
		NewServiceLANVPNInterfaceEthernetProfileParcelDataSource,
		NewServiceLANVPNInterfaceGREProfileParcelDataSource,
		NewServiceLANVPNInterfaceIPSecProfileParcelDataSource,
		NewServiceLANVPNInterfaceSVIProfileParcelDataSource,
		NewServiceMulticastProfileParcelDataSource,
		NewServiceObjectTrackerProfileParcelDataSource,
		NewServiceObjectTrackerGroupProfileParcelDataSource,
		NewServiceRoutePolicyProfileParcelDataSource,
		NewServiceRoutingBGPProfileParcelDataSource,
		NewServiceRoutingEIGRPProfileParcelDataSource,
		NewServiceRoutingOSPFProfileParcelDataSource,
		NewServiceRoutingOSPFv3IPv4ProfileParcelDataSource,
		NewServiceRoutingOSPFv3IPv6ProfileParcelDataSource,
		NewServiceSwitchportProfileParcelDataSource,
		NewServiceTrackerProfileParcelDataSource,
		NewServiceTrackerGroupProfileParcelDataSource,
		NewServiceWirelessLANProfileParcelDataSource,
		NewSystemAAAProfileParcelDataSource,
		NewSystemBannerProfileParcelDataSource,
		NewSystemBasicProfileParcelDataSource,
		NewSystemBFDProfileParcelDataSource,
		NewSystemFlexiblePortSpeedProfileParcelDataSource,
		NewSystemGlobalProfileParcelDataSource,
		NewSystemIPv4DeviceAccessProfileParcelDataSource,
		NewSystemIPv6DeviceAccessProfileParcelDataSource,
		NewSystemLoggingProfileParcelDataSource,
		NewSystemMRFProfileParcelDataSource,
		NewSystemNTPProfileParcelDataSource,
		NewSystemOMPProfileParcelDataSource,
		NewSystemPerformanceMonitoringProfileParcelDataSource,
		NewSystemRemoteAccessProfileParcelDataSource,
		NewSystemSecurityProfileParcelDataSource,
		NewSystemSNMPProfileParcelDataSource,
		NewTransportCellularControllerProfileParcelDataSource,
		NewTransportCellularProfileProfileParcelDataSource,
		NewTransportGPSProfileParcelDataSource,
		NewTransportIPv4ACLProfileParcelDataSource,
		NewTransportIPv6ACLProfileParcelDataSource,
		NewTransportIPv6TrackerProfileParcelDataSource,
		NewTransportIPv6TrackerGroupProfileParcelDataSource,
		NewTransportManagementVPNProfileParcelDataSource,
		NewTransportManagementVPNInterfaceEthernetProfileParcelDataSource,
		NewTransportRoutePolicyProfileParcelDataSource,
		NewTransportRoutingBGPProfileParcelDataSource,
		NewTransportRoutingOSPFProfileParcelDataSource,
		NewTransportRoutingOSPFv3IPv4ProfileParcelDataSource,
		NewTransportRoutingOSPFv3IPv6ProfileParcelDataSource,
		NewTransportT1E1ControllerProfileParcelDataSource,
		NewTransportTrackerProfileParcelDataSource,
		NewTransportTrackerGroupProfileParcelDataSource,
		NewTransportWANVPNProfileParcelDataSource,
		NewTransportWANVPNInterfaceCellularProfileParcelDataSource,
		NewTransportWANVPNInterfaceEthernetProfileParcelDataSource,
		NewTransportWANVPNInterfaceGREProfileParcelDataSource,
		NewTransportWANVPNInterfaceIPSECProfileParcelDataSource,
		NewTransportWANVPNInterfaceT1E1SerialProfileParcelDataSource,
		NewAdvancedInspectionProfilePolicyDefinitionDataSource,
		NewAdvancedMalwareProtectionPolicyDefinitionDataSource,
		NewAllowURLListPolicyObjectDataSource,
		NewAppProbeClassPolicyObjectDataSource,
		NewApplicationAwareRoutingPolicyDefinitionDataSource,
		NewApplicationListPolicyObjectDataSource,
		NewApplicationPriorityFeatureProfileDataSource,
		NewASPathListPolicyObjectDataSource,
		NewBlockURLListPolicyObjectDataSource,
		NewCentralizedPolicyDataSource,
		NewCflowdPolicyDefinitionDataSource,
		NewClassMapPolicyObjectDataSource,
		NewCLIConfigFeatureDataSource,
		NewCLIDeviceTemplateDataSource,
		NewCLIFeatureProfileDataSource,
		NewColorListPolicyObjectDataSource,
		NewConfigurationGroupDataSource,
		NewCustomControlTopologyPolicyDefinitionDataSource,
		NewDataFQDNPrefixListPolicyObjectDataSource,
		NewDataIPv4PrefixListPolicyObjectDataSource,
		NewDataIPv6PrefixListPolicyObjectDataSource,
		NewDeviceDataSource,
		NewDNSSecurityFeatureProfileDataSource,
		NewDNSSecurityPolicyDefinitionDataSource,
		NewDomainListPolicyObjectDataSource,
		NewEmbeddedSecurityFeatureProfileDataSource,
		NewExpandedCommunityListPolicyObjectDataSource,
		NewExtendedCommunityListPolicyObjectDataSource,
		NewFeatureDeviceTemplateDataSource,
		NewGeoLocationListPolicyObjectDataSource,
		NewHubAndSpokeTopologyPolicyDefinitionDataSource,
		NewIntrusionPreventionPolicyDefinitionDataSource,
		NewIPSSignatureListPolicyObjectDataSource,
		NewIPv4ACLPolicyDefinitionDataSource,
		NewIPv4DeviceACLPolicyDefinitionDataSource,
		NewIPv4PrefixListPolicyObjectDataSource,
		NewIPv6ACLPolicyDefinitionDataSource,
		NewIPv6DeviceACLPolicyDefinitionDataSource,
		NewIPv6PrefixListPolicyObjectDataSource,
		NewLocalApplicationListPolicyObjectDataSource,
		NewLocalizedPolicyDataSource,
		NewMeshTopologyPolicyDefinitionDataSource,
		NewMirrorPolicyObjectDataSource,
		NewObjectGroupPolicyDefinitionDataSource,
		NewOtherFeatureProfileDataSource,
		NewPolicerPolicyObjectDataSource,
		NewPolicyObjectFeatureProfileDataSource,
		NewPortListPolicyObjectDataSource,
		NewPreferredColorGroupPolicyObjectDataSource,
		NewProtocolListPolicyObjectDataSource,
		NewQoSMapPolicyDefinitionDataSource,
		NewRegionListPolicyObjectDataSource,
		NewRewriteRulePolicyDefinitionDataSource,
		NewRoutePolicyDefinitionDataSource,
		NewRuleSetPolicyDefinitionDataSource,
		NewSecurityPolicyDataSource,
		NewServiceFeatureProfileDataSource,
		NewServiceLANVPNInterfaceEthernetFeatureAssociateDHCPServerFeatureDataSource,
		NewServiceLANVPNInterfaceIPSecFeatureAssociateDHCPServerFeatureDataSource,
		NewServiceLANVPNInterfaceSVIFeatureAssociateDHCPServerFeatureDataSource,
		NewServiceLANVPNFeatureAssociateMulticastFeatureDataSource,
		NewServiceLANVPNFeatureAssociateRoutingBGPFeatureDataSource,
		NewServiceLANVPNFeatureAssociateRoutingEIGRPFeatureDataSource,
		NewServiceLANVPNFeatureAssociateRoutingOSPFFeatureDataSource,
		NewServiceLANVPNFeatureAssociateRoutingOSPFv3IPv4FeatureDataSource,
		NewServiceLANVPNFeatureAssociateRoutingOSPFv3IPv6FeatureDataSource,
		NewSIGSecurityFeatureProfileDataSource,
		NewSiteListPolicyObjectDataSource,
		NewSLAClassPolicyObjectDataSource,
		NewStandardCommunityListPolicyObjectDataSource,
		NewSystemFeatureProfileDataSource,
		NewTLOCListPolicyObjectDataSource,
		NewTLSSSLDecryptionPolicyDefinitionDataSource,
		NewTLSSSLProfilePolicyDefinitionDataSource,
		NewTrafficDataPolicyDefinitionDataSource,
		NewTransportFeatureProfileDataSource,
		NewTransportWANVPNFeatureAssociateRoutingBGPFeatureDataSource,
		NewTransportWANVPNFeatureAssociateRoutingOSPFFeatureDataSource,
		NewTransportWANVPNFeatureAssociateRoutingOSPFv3IPv4FeatureDataSource,
		NewTransportWANVPNFeatureAssociateRoutingOSPFv3IPv6FeatureDataSource,
		NewTransportWANVPNInterfaceCellularFeatureAssociateTrackerFeatureDataSource,
		NewTransportWANVPNInterfaceEthernetFeatureAssociateIPv6TrackerFeatureDataSource,
		NewTransportWANVPNInterfaceEthernetFeatureAssociateIPv6TrackerGroupFeatureDataSource,
		NewTransportWANVPNInterfaceEthernetFeatureAssociateTrackerFeatureDataSource,
		NewTransportWANVPNInterfaceEthernetFeatureAssociateTrackerGroupFeatureDataSource,
		NewTransportWANVPNInterfaceGREFeatureAssociateTrackerFeatureDataSource,
		NewTransportWANVPNInterfaceIPSECFeatureAssociateTrackerFeatureDataSource,
		NewURLFilteringPolicyDefinitionDataSource,
		NewVEdgeInventoryDataSource,
		NewVPNListPolicyObjectDataSource,
		NewVPNMembershipPolicyDefinitionDataSource,
		NewZoneBasedFirewallPolicyDefinitionDataSource,
		NewZoneListPolicyObjectDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &SdwanProvider{
			version: version,
		}
	}
}

// End of section. //template:end provider

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

func TestOSPFRoutePolicyVersion(t *testing.T) {
	tests := []struct {
		name       string
		manager    string
		global     types.String
		variable   types.String
		wantOption string
		wantValue  string
	}{
		{"old-global", "20.15.4", types.StringValue("policy-id"), types.StringNull(), "global", "policy-id"},
		{"new-global", "20.18.5", types.StringValue("policy-id"), types.StringNull(), "global", "policy-id"},
		{"old-variable", "20.15.4", types.StringNull(), types.StringValue("{{policy}}"), "", ""},
		{"new-variable", "20.18.5", types.StringNull(), types.StringValue("{{policy}}"), "variable", "{{policy}}"},
		{"old-unset", "20.15.4", types.StringNull(), types.StringNull(), "", ""},
		{"new-unset", "20.18.5", types.StringNull(), types.StringNull(), "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver := version.Must(version.NewVersion(tt.manager))
			ctx := context.Background()
			bodies := []struct {
				name    string
				body    string
				topPath string
			}{
				{
					"ospf",
					(ServiceRoutingOSPF{
						RoutePolicyId: tt.global, RoutePolicyIdVariable: tt.variable,
						Redistributes: []ServiceRoutingOSPFRedistributes{{RoutePolicyId: tt.global, RoutePolicyIdVariable: tt.variable}},
					}).toBody(ctx, ver),
					"data.routePolicy.refId",
				},
				{
					"ospfv3-ipv4",
					(ServiceRoutingOSPFv3IPv4{
						RoutePolicyId: tt.global, RoutePolicyIdVariable: tt.variable,
						Redistributes: []ServiceRoutingOSPFv3IPv4Redistributes{{RoutePolicyId: tt.global, RoutePolicyIdVariable: tt.variable}},
					}).toBody(ctx, ver),
					"data.advanced.policyName.refId",
				},
				{
					"ospfv3-ipv6",
					(ServiceRoutingOSPFv3IPv6{
						RoutePolicyId: tt.global, RoutePolicyIdVariable: tt.variable,
						Redistributes: []ServiceRoutingOSPFv3IPv6Redistributes{{RoutePolicyId: tt.global, RoutePolicyIdVariable: tt.variable}},
					}).toBody(ctx, ver),
					"data.advanced.policyName.refId",
				},
			}
			for _, body := range bodies {
				t.Run(body.name, func(t *testing.T) {
					for _, path := range []string{body.topPath, "data.redistribute.0.routePolicy.refId"} {
						option := gjson.Get(body.body, path+".optionType")
						value := gjson.Get(body.body, path+".value")
						if tt.wantOption == "" {
							if option.Exists() || value.Exists() {
								t.Errorf("%s unexpectedly present: %s", path, gjson.Get(body.body, path).Raw)
							}
						} else if option.String() != tt.wantOption || value.String() != tt.wantValue {
							t.Errorf("%s: got option %q, value %q; want %q, %q", path, option.String(), value.String(), tt.wantOption, tt.wantValue)
						}
					}
				})
			}
		})
	}
}

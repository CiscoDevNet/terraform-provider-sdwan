package sdwan

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/CiscoDevNet/sdwan-go-client/container"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceSDWANVPNFeatureTemplate() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSDWANVPNFeatureTemplateRead,

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"template_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"template_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_type": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"template_min_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"factory_default": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"attached_masters_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"devices_attached": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_on": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"g_template_class": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_on": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rid": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"feature": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template_definition": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// VPN Basic configuration
						"vpn_basic": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpn_id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ecmp_hash_key": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},

						// DNS configuration

						"vpn_dns": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"primary_dns_ipv4": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"secondary_dns_ipv4": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"primary_dns_ipv6": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"secondary_dns_ipv6": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"vpn_host": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hostname": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"ip": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},

						// IPv4 route

						"ipv4_route": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"gateway_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"next_hop": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"next_hop_address": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"next_hop_distance": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"null0_distance": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						//IPv6 route
						"ipv6_route": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"gateway_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"next_hop": {
										Type:     schema.TypeSet,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"next_hop_address": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"next_hop_distance": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"null0_distance": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						// TE service checker
						"te_service_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						// service route

						"service_route": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						// NAT64

						"nat64": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pool_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"start_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"end_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"overload": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func datasourceSDWANVPNFeatureTemplateRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Beginning Read Method %v", d.Id())

	sdwanClient := m.(*client.Client)

	ftName := d.Get("template_name").(string)

	cont, err := sdwanClient.GetViaURL("dataservice/template/feature")

	if err != nil {
		return err
	}

	var (
		data     *container.Container = cont.S("data")
		selected *container.Container
		tid      string
	)

	for _, child := range data.Children() {
		if child.S("templateName").Data().(string) == ftName {
			tid = child.S("templateId").Data().(string)
			dURL := fmt.Sprintf("dataservice/template/feature/object/%s", tid)
			selected, err = sdwanClient.GetViaURL(dURL)
			if err != nil {
				return err
			}
		}
	}

	if selected == nil {
		return fmt.Errorf("There is no VPN type of Feature Template named %s", ftName)
	}

	d.Set("template_name", stripQuotes(selected.S("templateName").String()))

	d.Set("template_description", stripQuotes(selected.S("templateDescription").String()))

	d.Set("template_min_version", stripQuotes(selected.S("templateMinVersion").String()))

	d.Set("template_type", stripQuotes(selected.S("templateType").String()))

	d.Set("factory_default", selected.S("factoryDefault").Data().(bool))

	d.Set("device_type", interfaceToStrList(selected.S("deviceType").Data()))

	if selected.Exists("attachedMastersCount") {
		d.Set("attached_masters_count", int((selected.S("attachedMastersCount").Data()).(float64)))
	}

	if selected.Exists("devicesAttached") {
		d.Set("devices_attached", int((selected.S("devicesAttached").Data()).(float64)))
	}

	if selected.Exists("lastUpdatedBy") {
		d.Set("last_updated_by", stripQuotes(selected.S("lastUpdatedBy").String()))
	}

	if selected.Exists("lastUpdatedOn") {
		d.Set("last_updated_on", int((selected.S("lastUpdatedOn").Data().(float64))))
	}

	if selected.Exists("gTemplateClass") {
		d.Set("g_template_class", stripQuotes(selected.S("gTemplateClass").String()))
	}

	if selected.Exists("configType") {
		d.Set("config_type", stripQuotes(selected.S("configType").String()))
	}

	if selected.Exists("createdOn") {
		d.Set("created_on", int((selected.S("createdOn").Data()).(float64)))
	}

	if selected.Exists("createdBy") {
		d.Set("created_by", stripQuotes(selected.S("createdBy").String()))
	}

	if selected.Exists("feature") {
		d.Set("feature", stripQuotes(selected.S("feature").String()))
	}

	d.Set("template_id", stripQuotes(selected.S("templateId").String()))

	if selected.Exists("@rid") {
		d.Set("rid", int((selected.S("@rid").Data()).(float64)))
	}

	ftDef := selected.S("templateDefinition")
	FTdefinition := getVPNdefinition(ftDef)

	d.Set("template_definition", FTdefinition)

	d.SetId(ftName)

	log.Printf("[DEBUG] Ending of Read Method %s", d.Id())
	return nil
}

func getVPNdefinition(cont *container.Container) []map[string]interface{} {
	definition := make([]map[string]interface{}, 0, 1)

	VPNDef := make(map[string]interface{})

	VPNDef["vpn_basic"] = getVPNBasic(cont)

	if cont.Exists("dns", "vipValue") {
		VPNDef["vpn_dns"] = getVPNDNS(cont)
	}

	if cont.Exists("ip", "route", "vipValue") {
		VPNDef["ipv4_route"] = getVPNIPv4Route(cont)
	}

	if cont.Exists("ipv6", "route", "vipValue") {
		VPNDef["ipv6_route"] = getVPNIPv6Route(cont)
	}

	if stripQuotes(cont.S("service", "vipType").String()) == "constant" {
		VPNDef["te_service_enabled"] = true
	} else {
		VPNDef["te_service_enabled"] = false
	}

	if cont.Exists("ip", "service-route", "vipValue") {
		VPNDef["service_route"] = getVPNServiceRoute(cont)
	}

	if cont.Exists("nat64", "v4", "pool", "vipValue") {
		VPNDef["nat64"] = getNAT(cont)
	}

	definition = append(definition, VPNDef)

	return definition
}

func getVPNBasic(cont *container.Container) []map[string]interface{} {

	basic := make([]map[string]interface{}, 0, 1)

	result := make(map[string]interface{})

	if cont.Exists("vpn-id", "vipValue") {
		if value, err := strconv.Atoi(cont.S("vpn-id", "vipValue").String()); err == nil {
			result["vpn_id"] = value
		}
	}

	if cont.Exists("name", "vipValue") {
		result["name"] = stripQuotes(cont.S("name", "vipValue").String())
	}

	if cont.Exists("ecmp-hash-key", "layer4", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(cont.S("ecmp-hash-key", "layer4", "vipValue").String())); err == nil {
			result["ecmp_hash_key"] = value
		}
	}

	basic = append(basic, result)
	return basic
}

func getVPNDNS(cont *container.Container) []map[string]interface{} {

	dns := make(map[string]interface{})

	result := make([]map[string]interface{}, 0, 1)

	if cont.Exists("dns", "vipValue") {
		for _, dnsobj := range cont.S("dns", "vipValue").Children() {

			if dnsobj.Exists("role", "vipValue") && stripQuotes(dnsobj.S("role", "vipValue").String()) == "primary" {
				dns["primary_dns_ipv4"] = stripQuotes(dnsobj.S("dns-addr", "vipValue").String())
			}

			if dnsobj.Exists("role", "vipValue") && stripQuotes(dnsobj.S("role", "vipValue").String()) == "secondary" {
				dns["secondary_dns_ipv4"] = stripQuotes(dnsobj.S("dns-addr", "vipValue").String())
			}

			if dnsobj.Exists("role", "vipValue") && stripQuotes(dnsobj.S("role", "vipValue").String()) == "primaryv6" {
				dns["primary_dns_ipv6"] = stripQuotes(dnsobj.S("dns-addr", "vipValue").String())
			}

			if dnsobj.Exists("role", "vipValue") && stripQuotes(dnsobj.S("role", "vipValue").String()) == "secondaryv6" {
				dns["secondary_dns_ipv6"] = stripQuotes(dnsobj.S("dns-addr", "vipValue").String())
			}
		}
	}

	hosts := make([]map[string]interface{}, 0, 1)
	if cont.Exists("host", "vipValue") {
		for _, hostobj := range cont.S("host", "vipValue").Children() {
			host := make(map[string]interface{})

			if hostobj.Exists("hostname", "vipValue") {
				host["hostname"] = stripQuotes(hostobj.S("hostname", "vipValue").String())
			}

			host["ip"] = interfaceToStrList(hostobj.S("ip", "vipValue").Data())

			hosts = append(hosts, host)
		}
		dns["vpn_host"] = hosts
	}
	result = append(result, dns)

	return result
}

func getVPNIPv4Route(cont *container.Container) []map[string]interface{} {

	routes := make([]map[string]interface{}, 0, 1)

	if cont.Exists("ip", "route", "vipValue") {
		for _, routeobj := range cont.S("ip", "route", "vipValue").Children() {
			route := make(map[string]interface{})

			if routeobj.Exists("prefix", "vipValue") {
				route["prefix"] = stripQuotes(routeobj.S("prefix", "vipValue").String())
			}

			if routeobj.Exists("dhcp") {
				route["gateway_type"] = "dhcp"
			} else if routeobj.Exists("null0") {
				route["gateway_type"] = "null0"
				if routeobj.Exists("distance", "vipValue") {
					route["null0_distance"] = stripQuotes(routeobj.S("distance", "vipValue").String())
				}
			} else if routeobj.Exists("next-hop") {
				route["gateway_type"] = "next-hop"
				nexthops := make([]map[string]interface{}, 0, 1)
				if routeobj.Exists("next-hop", "vipValue") {
					for _, nexthopobj := range routeobj.S("next-hop", "vipValue").Children() {
						nexthop := make(map[string]interface{})

						if nexthopobj.Exists("address", "vipValue") {
							nexthop["next_hop_address"] = stripQuotes(nexthopobj.S("address", "vipValue").String())
						}

						if nexthopobj.Exists("distance", "vipValue") {
							nexthop["next_hop_distance"] = stripQuotes(nexthopobj.S("distance", "vipValue").String())
						}
						nexthops = append(nexthops, nexthop)
					}
				}
				route["next_hop"] = nexthops
			}

			routes = append(routes, route)
		}
		return routes
	}
	return nil
}

func getVPNIPv6Route(cont *container.Container) []map[string]interface{} {
	routes := make([]map[string]interface{}, 0, 1)
	if cont.Exists("ipv6", "route", "vipValue") {
		for _, routeobj := range cont.S("ipv6", "route", "vipValue").Children() {
			route := make(map[string]interface{})

			if routeobj.Exists("prefix", "vipValue") {
				route["prefix"] = stripQuotes(routeobj.S("prefix", "vipValue").String())
			}

			if routeobj.Exists("null0") {
				route["gateway_type"] = "null0"
				if routeobj.Exists("distance", "vipValue") {
					route["null0_distance"] = stripQuotes(routeobj.S("distance", "vipValue").String())
				}
			} else if routeobj.Exists("next-hop") {
				route["gateway_type"] = "next-hop"
				nexthops := make([]map[string]interface{}, 0, 1)

				if routeobj.Exists("next-hop", "vipValue") {
					for _, nexthopobj := range routeobj.S("next-hop", "vipValue").Children() {

						nexthop := make(map[string]interface{})

						if nexthopobj.Exists("address", "vipValue") {
							nexthop["next_hop_address"] = stripQuotes(nexthopobj.S("address", "vipValue").String())
						}

						if nexthopobj.Exists("distance", "vipValue") {
							nexthop["next_hop_distance"] = stripQuotes(nexthopobj.S("distance", "vipValue").String())
						}

						nexthops = append(nexthops, nexthop)
					}
				}

				route["next_hop"] = nexthops
			}
			routes = append(routes, route)
		}
		return routes
	}
	return nil
}

func getVPNServiceRoute(cont *container.Container) []map[string]interface{} {
	serroutes := make([]map[string]interface{}, 0, 1)

	if cont.Exists("ip", "service-route", "vipValue") {
		for _, serrouteobj := range cont.S("ip", "service-route", "vipValue").Children() {
			serroute := make(map[string]interface{})

			if serrouteobj.Exists("prefix", "vipValue") {
				serroute["prefix"] = stripQuotes(serrouteobj.S("prefix", "vipValue").String())
			}

			if serrouteobj.Exists("service", "vipValue") {
				serroute["service_type"] = stripQuotes(serrouteobj.S("service", "vipValue").String())
			}

			serroutes = append(serroutes, serroute)
		}

		return serroutes
	}
	return nil
}

func getNAT(cont *container.Container) []map[string]interface{} {

	nats := make([]map[string]interface{}, 0, 1)

	if cont.Exists("nat64", "v4", "pool", "vipValue") {
		for _, natobj := range cont.S("nat64", "v4", "pool", "vipValue").Children() {
			nat := make(map[string]interface{})

			if natobj.Exists("name", "vipValue") {
				nat["pool_name"] = stripQuotes(natobj.S("name", "vipValue").String())
			}

			if natobj.Exists("start-address", "vipValue") {
				nat["start_address"] = stripQuotes(natobj.S("start-address", "vipValue").String())
			}

			if natobj.Exists("end-address", "vipValue") {
				nat["end_address"] = stripQuotes(natobj.S("end-address", "vipValue").String())
			}

			if natobj.Exists("overload", "vipValue") {
				if value, err := strconv.ParseBool(stripQuotes(natobj.S("overload", "vipValue").String())); err == nil {
					nat["overload"] = value
				}
			}

			nats = append(nats, nat)
		}
		return nats
	}

	return nil
}

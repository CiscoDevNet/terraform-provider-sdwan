package sdwan

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/CiscoDevNet/sdwan-go-client/container"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceSDWANSystemFeatureTemplate() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSDWANFeatureTemplateRead,

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"template_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"template_description": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"device_type": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"template_type": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"template_min_version": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"template_definition": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						//For Basic Configuration
						"system_basic": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"overlay_id": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"timezone": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"location": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_groups": {
										Type: schema.TypeList,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Computed: true,
									},

									"controller_groups": {
										Type: schema.TypeList,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Computed: true,
									},

									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"console_baud_rate": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"max_omp_sessions": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tcp_optimization_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},

						//For GPS Configuration
						"system_gps": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"latitude": {
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"longitude": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},

						//For Tracker configuration
						"system_trackers": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"tracker_name": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"tracker_type": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"tracker_endpoint_type": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"tracker_endpoint_value": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"tracker_threshold": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tracker_interval": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tracker_multiplier": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						//For Advanced Configuration
						"system_advanced": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"control_session_pps": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"system_tunnel_mtu": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"port_hop": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									"port_offset": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"dns_cache_timeout": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"track_transport": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									"vbond_local": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									"track_interface_tag": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"multicast_buffer_percent": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"usb_controller": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									"track_default_gateway": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									"host_policer_pps": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"icmp_error_pps": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"allow_same_site_tunnels": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									"route_consistency_check": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									"admin_tech_on_failure": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									"idle_timeout": {
										Type:     schema.TypeInt,
										Computed: true,
									},

									"eco_friendly_mode": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									"on_demand_enable": {
										Type:     schema.TypeBool,
										Computed: true,
									},

									"on_demand_idle_timeout": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
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

			"rid": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"feature": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func datasourceSDWANFeatureTemplateRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Read Method ", d.Id())

	sdwanClient := m.(*client.Client)

	ftName := d.Get("template_name").(string)
	cont, err := getFeatureTemplateByName(sdwanClient, ftName)
	if err != nil {
		if cont != nil {
			return fmt.Errorf("[ERROR] Data load failed in between")
		} else {
			return err
		}
	}
	if strings.Compare(stripQuotes(cont.S("templateType").String()), "system-vedge") != 0 &&
		strings.Compare(stripQuotes(cont.S("templateType").String()), "cisco_system") != 0 {
		return fmt.Errorf("[ERROR] Invalid Template Type")
	}
	setSystemTemplateAttributes(d, cont)
	d.SetId(ftName)

	log.Println("[DEBUG] End of Read Method ", d.Id())
	return nil
}

func getFeatureTemplateByName(client *client.Client, ftName string) (*container.Container, error) {
	cont, err := client.GetViaURL("/dataservice/template/feature")
	if err != nil {
		return nil, err
	} else {

		if cont != nil {
			cont = cont.S("data")
			for _, template := range cont.Children() {
				if stripQuotes(template.S("templateName").String()) == ftName {

					ftID := stripQuotes(template.S("templateId").String())
					return getFeatureTemplateByID(client, ftID)
				}
			}
		}

	}

	return nil, fmt.Errorf("[ERROR] No such Record exists")
}

func getFeatureTemplateByID(client *client.Client, ftID string) (*container.Container, error) {
	cont, err := client.GetViaURL(fmt.Sprintf("/dataservice/template/feature/object/%s", ftID))
	if err != nil {
		return nil, err
	}

	return cont, err
}

func setSystemTemplateAttributes(d *schema.ResourceData, cont *container.Container) *schema.ResourceData {

	d.Set("template_name", stripQuotes(cont.S("templateName").String()))

	d.Set("template_description", stripQuotes(cont.S("templateDescription").String()))

	d.Set("device_type", interfaceToStrList(cont.S("deviceType").Data()))

	templateType := stripQuotes(cont.S("templateType").String())
	if templateType == "system-vedge" {
		templateType = "System"
	} else if templateType == "cisco_system" {
		templateType = "Cisco System"
	} else {
		log.Println("[ERROR] Only System and Cisco System are supported")
	}
	d.Set("template_type", templateType)

	d.Set("template_min_version", stripQuotes(cont.S("templateMinVersion").String()))

	ftDef := cont.S("templateDefinition")
	FTDefinition := getSystemFTDefinition(ftDef)

	d.Set("template_definition", FTDefinition)

	if value, err := strconv.ParseBool(stripQuotes(cont.S("factoryDefault").String())); err == nil {
		d.Set("factory_default", value)
	}

	if cont.Exists("attachedMastersCount") {
		if value, err := strconv.Atoi(stripQuotes(cont.S("attachedMastersCount").String())); err == nil {
			d.Set("attached_masters_count", value)
		}
	}

	if cont.Exists("devicesAttached") {
		if value, err := strconv.Atoi(stripQuotes(cont.S("devicesAttached").String())); err == nil {
			d.Set("devices_attached", value)
		}
	}

	if cont.Exists(("lastUpdatedBy")) {
		d.Set("last_updated_by", stripQuotes(cont.S("lastUpdatedBy").String()))
	}

	if cont.Exists("lastUpdatedOn") {
		if value, err := strconv.Atoi(stripQuotes(cont.S("lastUpdatedOn").String())); err == nil {
			d.Set("last_updated_on", value)
		}
	}

	if cont.Exists(("gTemplateClass")) {
		d.Set("g_template_class", stripQuotes(cont.S("gTemplateClass").String()))
	}

	if cont.Exists(("configType")) {
		d.Set("config_type", stripQuotes(cont.S("configType").String()))
	}

	if cont.Exists("createdOn") {
		if value, err := strconv.Atoi(stripQuotes(cont.S("createdOn").String())); err == nil {
			d.Set("created_on", value)
		}
	}
	if cont.Exists("@rid") {
		if value, err := strconv.Atoi(stripQuotes(cont.S("@rid").String())); err == nil {
			d.Set("rid", value)
		}
	}

	if cont.Exists("feature") {
		d.Set("feature", stripQuotes(cont.S("feature").String()))
	}

	if cont.Exists("createdBy") {
		d.Set("created_by", stripQuotes(cont.S("createdBy").String()))
	}

	d.Set("template_id", stripQuotes(cont.S("templateId").String()))

	return d
}

func getSystemFTDefinition(cont *container.Container) []map[string]interface{} {
	definition := make([]map[string]interface{}, 0, 1)

	systemDef := make(map[string]interface{})

	systemDef["system_basic"] = getSystemBasic(cont)

	systemDef["system_gps"] = getSystemGPS(cont)

	if cont.Exists("tracker", "vipValue") {
		systemDef["system_trackers"] = getSystemTrackers(cont)
	}

	systemDef["system_advanced"] = getSystemAdvanced(cont)

	definition = append(definition, systemDef)

	return definition
}

func getSystemBasic(cont *container.Container) []map[string]interface{} {
	basic := make([]map[string]interface{}, 0, 1)

	result := make(map[string]interface{})

	if cont.Exists("overlay-id", "vipValue") {
		if value, err := strconv.Atoi(cont.S("overlay-id", "vipValue").String()); err == nil {
			result["overlay_id"] = value
		}
	}

	if cont.Exists("clock", "timezone", "vipValue") {
		result["timezone"] = stripQuotes(cont.S("clock", "timezone", "vipValue").String())
	}

	if cont.Exists("location", "vipValue") {
		result["location"] = stripQuotes(cont.S("location", "vipValue").String())
	}

	if cont.Exists("device-groups", "vipValue") {
		val := cont.S("device-groups", "vipValue").Data().([]interface{})
		if len(val) > 0 {
			result["device_groups"] = interfaceToStrList(cont.S("device-groups", "vipValue").Data())
		}
	}

	if cont.Exists("controller-group-list", "vipValue") {
		val := cont.S("controller-group-list", "vipValue").Data().([]interface{})
		if len(val) > 0 {
			result["controller_groups"] = interfaceToStrList(cont.S("controller-group-list", "vipValue").Data())
		}
	}

	if cont.Exists("description", "vipValue") {
		result["description"] = stripQuotes(cont.S("description", "vipValue").String())
	}

	if cont.Exists("console-baud-rate", "vipValue") {
		result["console_baud_rate"] = stripQuotes(cont.S("console-baud-rate", "vipValue").String())
	}

	if cont.Exists("max-omp-sessions", "vipValue") {
		if value, err := strconv.Atoi(cont.S("max-omp-sessions", "vipValue").String()); err == nil {
			result["max_omp_sessions"] = value
		}
	}

	if cont.Exists("tcp-optimization-enabled", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(cont.S("tcp-optimization-enabled", "vipValue").String())); err == nil {
			result["tcp_optimization_enabled"] = value
		}
	}

	basic = append(basic, result)
	return basic
}

func getSystemGPS(cont *container.Container) []map[string]interface{} {
	gps := make([]map[string]interface{}, 0, 1)

	result := make(map[string]interface{})
	isPresent := 0

	if cont.Exists("gps-location", "latitude", "vipValue") {
		if value, err := strconv.Atoi(cont.S("gps-location", "latitude", "vipValue").String()); err == nil {
			result["latitude"] = value
			isPresent += 1
		}
	}

	if cont.Exists("gps-location", "longitude", "vipValue") {
		if value, err := strconv.Atoi(cont.S("gps-location", "longitude", "vipValue").String()); err == nil {
			result["longitude"] = value
			isPresent += 1
		}
	}

	if isPresent > 0 {
		gps = append(gps, result)
		return gps
	}
	return nil
}

func getSystemTrackers(cont *container.Container) []map[string]interface{} {
	trackers := make([]map[string]interface{}, 0, 8)

	if cont.Exists("tracker", "vipValue") {
		for _, trackerObject := range cont.S("tracker", "vipValue").Children() {
			tracker := make(map[string]interface{})

			if trackerObject.Exists("name", "vipValue") {
				tracker["tracker_name"] = stripQuotes(trackerObject.S("name", "vipValue").String())
			}

			if trackerObject.Exists("type", "vipValue") {
				tracker["tracker_type"] = stripQuotes(trackerObject.S("type", "vipValue").String())
			}

			if trackerObject.Exists("endpoint-ip", "vipValue") {
				tracker["tracker_endpoint_type"] = "ip-address"

				tracker["tracker_endpoint_value"] = stripQuotes(trackerObject.S("endpoint-ip", "vipValue").String())
			}

			if trackerObject.Exists("endpoint-dns-name", "vipValue") {
				tracker["tracker_endpoint_type"] = "dns-name"

				tracker["tracker_endpoint_value"] = stripQuotes(trackerObject.S("endpoint-dns-name", "vipValue").String())
			}

			if trackerObject.Exists("endpoint-api-url", "vipValue") {
				tracker["tracker_endpoint_type"] = "api-url"

				tracker["tracker_endpoint_value"] = stripQuotes(trackerObject.S("endpoint-api-url", "vipValue").String())
			}

			if trackerObject.Exists("threshold", "vipValue") {
				if value, err := strconv.Atoi(trackerObject.S("threshold", "vipValue").String()); err == nil {
					tracker["tracker_threshold"] = value
				}
			}

			if trackerObject.Exists("interval", "vipValue") {
				if value, err := strconv.Atoi(trackerObject.S("interval", "vipValue").String()); err == nil {
					tracker["tracker_interval"] = value
				}
			}

			if trackerObject.Exists("multiplier", "vipValue") {
				if value, err := strconv.Atoi(trackerObject.S("multiplier", "vipValue").String()); err == nil {
					tracker["tracker_multiplier"] = value
				}
			}
			trackers = append(trackers, tracker)
		}
		return trackers
	}
	return nil
}

func getSystemAdvanced(cont *container.Container) []map[string]interface{} {
	advanced := make([]map[string]interface{}, 0, 1)

	result := make(map[string]interface{})
	isPresent := 0

	if cont.Exists("control-session-pps", "vipValue") {
		if value, err := strconv.Atoi(cont.S("control-session-pps", "vipValue").String()); err == nil {
			result["control_session_pps"] = value
			isPresent += 1
		}
	}

	if cont.Exists("system-tunnel-mtu", "vipValue") {
		if value, err := strconv.Atoi(cont.S("system-tunnel-mtu", "vipValue").String()); err == nil {
			result["system_tunnel_mtu"] = value
			isPresent += 1
		}
	}

	if cont.Exists("port-hop", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(cont.S("port-hop", "vipValue").String())); err == nil {
			result["port_hop"] = value
			isPresent += 1
		}
	}

	if cont.Exists("port-offset", "vipValue") {
		if value, err := strconv.Atoi(cont.S("port-offset", "vipValue").String()); err == nil {
			result["port_offset"] = value
			isPresent += 1
		}
	}

	if cont.Exists("timer", "dns-cache-timeout", "vipValue") {
		if value, err := strconv.Atoi(cont.S("timer", "dns-cache-timeout", "vipValue").String()); err == nil {
			result["dns_cache_timeout"] = value
			isPresent += 1
		}
	}

	if cont.Exists("track-transport", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(cont.S("track-transport", "vipValue").String())); err == nil {
			result["track_transport"] = value
			isPresent += 1
		}
	}

	if cont.Exists("vbond", "local", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(cont.S("vbond", "local", "vipValue").String())); err == nil {
			result["vbond_local"] = value
			isPresent += 1
		}
	}

	if cont.Exists("track-interface-tag", "vipValue") {
		if value, err := strconv.Atoi(cont.S("track-interface-tag", "vipValue").String()); err == nil {
			result["track_interface_tag"] = value
			isPresent += 1
		}
	}

	if cont.Exists("multicast-buffer-percent", "vipValue") {
		if value, err := strconv.Atoi(cont.S("multicast-buffer-percent", "vipValue").String()); err == nil {
			result["multicast_buffer_percent"] = value
			isPresent += 1
		}
	}

	if cont.Exists("usb-controller", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(cont.S("usb-controller", "vipValue").String())); err == nil {
			result["usb_controller"] = value
			isPresent += 1
		}
	}

	if cont.Exists("track-default-gateway", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(cont.S("track-default-gateway", "vipValue").String())); err == nil {
			result["track_default_gateway"] = value
			isPresent += 1
		}
	}

	if cont.Exists("host-policer-pps", "vipValue") {
		if value, err := strconv.Atoi(cont.S("host-policer-pps", "vipValue").String()); err == nil {
			result["host_policer_pps"] = value
			isPresent += 1
		}
	}

	if cont.Exists("icmp-error-pps", "vipValue") {
		if value, err := strconv.Atoi(cont.S("icmp-error-pps", "vipValue").String()); err == nil {
			result["icmp_error_pps"] = value
			isPresent += 1
		}
	}

	if cont.Exists("allow-same-site-tunnels", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(cont.S("allow-same-site-tunnels", "vipValue").String())); err == nil {
			result["allow_same_site_tunnels"] = value
			isPresent += 1
		}
	}

	if cont.Exists("route-consistency-check", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(cont.S("route-consistency-check", "vipValue").String())); err == nil {
			result["route_consistency_check"] = value
			isPresent += 1
		}
	}

	if cont.Exists("admin-tech-on-failure", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(cont.S("admin-tech-on-failure", "vipValue").String())); err == nil {
			result["admin_tech_on_failure"] = value
			isPresent += 1
		}
	}

	if cont.Exists("idle-timeout", "vipValue") {
		if value, err := strconv.Atoi(cont.S("idle-timeout", "vipValue").String()); err == nil {
			result["idle_timeout"] = value
			isPresent += 1
		}
	}

	if cont.Exists("eco-friendly-mode", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(cont.S("eco-friendly-mode", "vipValue").String())); err == nil {
			result["eco_friendly_mode"] = value
			isPresent += 1
		}
	}

	if cont.Exists("on-demand", "enable", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(cont.S("on-demand", "enable", "vipValue").String())); err == nil {
			result["on_demand_enable"] = value
			isPresent += 1
		}
	}

	if cont.Exists("on-demand", "idle-timeout", "vipValue") {
		if value, err := strconv.Atoi(cont.S("on-demand", "idle-timeout", "vipValue").String()); err == nil {
			result["on_demand_idle_timeout"] = value
			isPresent += 1
		}
	}
	if isPresent > 0 {
		advanced = append(advanced, result)
		return advanced
	}
	return nil
}

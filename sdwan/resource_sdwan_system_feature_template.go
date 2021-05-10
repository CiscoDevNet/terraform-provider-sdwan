package sdwan

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/CiscoDevNet/sdwan-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSDWANSystemFeatureTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSDWANSystemFeatureTemplateCreate,
		Update: resourceSDWANSystemFeatureTemplateUpdate,
		Read:   resourceSDWANSystemFeatureTemplateRead,
		Delete: resourceSDWANSystemFeatureTemplateDelete,

		Schema: map[string]*schema.Schema{
			"template_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: NameValidation(),
			},

			"template_description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"device_type": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"vbond",
						"vedge-100",
						"vedge-100-B",
						"vedge-100-M",
						"vedge-100-WM",
						"vedge-1000",
						"vedge-2000",
						"vedge-5000",
						"vedge-cloud",
						"vedge-ISR1100-4G",
						"vedge-ISR1100-4GLTE",
						"vedge-ISR1100-6G",
						"vedge-ISR1100X-4G",
						"vedge-ISR1100X-6G",
						"vedge-nfvis-CSP-5444",
						"vedge-nfvis-CSP-5456",
						"vedge-nfvis-CSP2100",
						"vedge-nfvis-CSP2100-X1",
						"vedge-nfvis-CSP2100-X2",
						"vedge-nfvis-UCSC-E",
						"vedge-nfvis-UCSC-M5",
					}, false),
				},
			},

			"template_min_version": {
				Type:     schema.TypeString,
				Required: true,
			},

			"template_definition": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"system_basic": {
							Type:     schema.TypeSet,
							Required: true,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"site_id": {
										Type:         schema.TypeFloat,
										Required:     true,
										ValidateFunc: validation.FloatBetween(1, 4294967295),
									},

									"system_ip": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.IsIPv4Address,
									},

									"overlay_id": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  1,
									},

									"timezone": {
										Type:     schema.TypeString,
										Required: true,
									},

									"hostname": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringLenBetween(1, 32),
									},

									"location": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringLenBetween(1, 128),
									},

									"device_groups": {
										Type: schema.TypeList,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Optional: true,
									},

									"controller_groups": {
										Type: schema.TypeList,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Optional: true,
									},

									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"console_baud_rate": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "_empty",
										ValidateFunc: validation.StringInSlice(
											[]string{
												"1200",
												"2400",
												"4800",
												"9600",
												"19200",
												"38400",
												"57600",
												"115200",
											}, false),
									},

									"max_omp_sessions": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      2,
										ValidateFunc: validation.IntBetween(0, 100),
									},

									"tcp_optimization_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
								},
							},
						},

						"system_gps": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"latitude": {
										Type:     schema.TypeFloat,
										Optional: true,
									},

									"longitude": {
										Type:     schema.TypeFloat,
										Optional: true,
									},
								},
							},
						},

						"system_trackers": {
							Type:     schema.TypeSet,
							MaxItems: 8,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tracker_name": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: NameValidation(),
									},

									"tracker_threshold": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      300,
										ValidateFunc: validation.IntBetween(100, 1000),
									},

									"tracker_interval": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      60,
										ValidateFunc: validation.IntBetween(10, 600),
									},

									"tracker_multiplier": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      3,
										ValidateFunc: validation.IntBetween(1, 10),
									},

									"tracker_type": {
										Type:     schema.TypeString,
										Required: true,
										ValidateFunc: validation.StringInSlice([]string{
											"ip-address",
											"dns-name",
											"api-url",
										}, false),
									},

									"tracker_endpoint_value": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"system_advanced": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"control_session_pps": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      300,
										ValidateFunc: validation.IntBetween(1, 65535),
									},

									"system_tunnel_mtu": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      1024,
										ValidateFunc: validation.IntBetween(500, 2000),
									},

									"port_hop": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},

									"port_offset": {
										Type:         schema.TypeInt,
										Optional:     true,
										ValidateFunc: validation.IntBetween(0, 19),
									},

									"dns_cache_timeout": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      2,
										ValidateFunc: validation.IntBetween(1, 30),
									},

									"track_transport": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},

									"vbond_local": {
										Type:     schema.TypeBool,
										Default:  false,
										Optional: true,
									},

									"vbond_remote": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"track_interface_tag": {
										Type:         schema.TypeFloat,
										Optional:     true,
										ValidateFunc: validation.FloatBetween(1, 4294967295),
									},

									"multicast_buffer_percent": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      20,
										ValidateFunc: validation.IntBetween(5, 100),
									},

									"usb_controller": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},

									"track_default_gateway": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},

									"host_policer_pps": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      20000,
										ValidateFunc: validation.IntBetween(1000, 25000),
									},

									"icmp_error_pps": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      100,
										ValidateFunc: validation.IntBetween(1, 200),
									},

									"allow_same_site_tunnels": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},

									"route_consistency_check": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},

									"admin_tech_on_failure": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},

									"idle_timeout": {
										Type:         schema.TypeInt,
										Optional:     true,
										ValidateFunc: validation.IntBetween(0, 300),
									},

									"eco_friendly_mode": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},

									"on_demand_enable": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},

									"on_demand_idle_timeout": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  10,
									},
								},
							},
						},
					},
				},
			},

			"factory_default": {
				Type:     schema.TypeBool,
				Required: true,
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

func resourceSDWANSystemFeatureTemplateCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Create method!")

	sdwanClient := m.(*client.Client)

	fTemplate := models.FeatureTemplate{}

	fTemplate.TemplateType = "system-vedge"
	fTemplate.TemplateName = d.Get("template_name").(string)
	fTemplate.TemplateDescription = d.Get("template_description").(string)
	fTemplate.DeviceType = interfaceToStrList(d.Get("device_type"))
	fTemplate.TemplateMinVersion = d.Get("template_min_version").(string)
	fTemplate.FactoryDefault = d.Get("factory_default").(bool)

	ftDefinition := d.Get("template_definition").(*schema.Set).List()
	if def, err := createSystemFTDefinition(ftDefinition); err == nil {
		fTemplate.TemplateDefinition = def
	}

	cont, err := sdwanClient.Save("/dataservice/template/feature", &fTemplate)
	if err != nil {
		return err
	}

	ftID := stripQuotes(cont.S("templateId").String())
	d.Set("template_id", ftID)
	d.SetId(ftID)

	return resourceSDWANSystemFeatureTemplateRead(d, m)
}

func resourceSDWANSystemFeatureTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Update Method ", d.Id())

	sdwanClient := m.(*client.Client)

	ftID := d.Id()

	fTemplate := models.FeatureTemplate{}

	fTemplate.TemplateType = "system-vedge"
	fTemplate.TemplateName = d.Get("template_name").(string)
	fTemplate.TemplateDescription = d.Get("template_description").(string)
	fTemplate.DeviceType = interfaceToStrList(d.Get("device_type"))
	fTemplate.TemplateMinVersion = d.Get("template_min_version").(string)
	fTemplate.FactoryDefault = d.Get("factory_default").(bool)

	ftDefinition := d.Get("template_definition").(*schema.Set).List()
	if def, err := createSystemFTDefinition(ftDefinition); err == nil {
		fTemplate.TemplateDefinition = def
	}

	ftURL := fmt.Sprintf("/dataservice/template/feature/%s", ftID)
	_, err := sdwanClient.Update(ftURL, &fTemplate)

	if err != nil {
		return err
	}
	d.SetId(ftID)

	log.Println("[DEBUG] End of Update Method ", d.Id())
	return resourceSDWANSystemFeatureTemplateRead(d, m)
}

func resourceSDWANSystemFeatureTemplateRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Read Method ", d.Id())

	sdwanClient := m.(*client.Client)

	tID := d.Id()
	cont, err := getFeatureTemplateByID(sdwanClient, tID)
	if err != nil {
		if cont != nil {
			log.Println("Data load failed in between!")
		} else {
			return err
		}
	}

	setSystemTemplateAttributes(d, cont)

	log.Println("[DEBUG] End of Read Method ", d.Id())
	return nil
}

func resourceSDWANSystemFeatureTemplateDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Delete Method ", d.Id())

	sdwanClient := m.(*client.Client)
	ftID := d.Id()

	_, err := sdwanClient.Delete(fmt.Sprintf("/dataservice/template/feature/%s", ftID))
	if err != nil {
		return fmt.Errorf("[ERROR] Error occured while destroying")
	}

	d.SetId("")
	log.Println("[DEBUG] End of Delete Method!")
	return nil
}

func createSystemFTDefinition(ftDefinitions []interface{}) (map[string]interface{}, error) {
	definition := make(map[string]interface{})

	ftDefinition := ftDefinitions[0]
	ftDefMap := ftDefinition.(map[string]interface{})

	systemBasicMap := (ftDefMap["system_basic"].(*schema.Set).List())[0].(map[string]interface{})
	createSystemBasic(definition, systemBasicMap)

	GPSSet := ftDefMap["system_gps"].(*schema.Set).List()

	if len(GPSSet) > 0 {
		systemGPSMap := GPSSet[0].(map[string]interface{})
		createSystemGPS(definition, systemGPSMap)
	}

	TrackerSet := ftDefMap["system_trackers"].(*schema.Set).List()

	if len(TrackerSet) > 0 {
		trackers := make([]interface{}, 0, 8)

		for _, systemTrackerMap := range TrackerSet {
			tracker := make(map[string]interface{})
			createSystemTracker(tracker, systemTrackerMap.(map[string]interface{}))

			trackers = append(trackers, tracker)
		}

		TrackerMap := make(map[string]interface{})
		TrackerMap["vipType"] = "constant"
		TrackerMap["vipObjectType"] = "tree"
		TrackerMap["vipValue"] = trackers
		TrackerMap["vipPrimaryKey"] = []string{
			"name",
		}

		definition["tracker"] = TrackerMap
	}

	AdvancedSet := ftDefMap["system_advanced"].(*schema.Set).List()

	if len(AdvancedSet) > 0 {
		systemAdvancedMap := AdvancedSet[0].(map[string]interface{})
		createSystemAdvanced(definition, systemAdvancedMap)
	}

	return definition, nil
}

func createSystemBasic(defMap map[string]interface{}, input map[string]interface{}) {

	if input["site_id"] != nil {
		siteId := make(map[string]interface{})
		siteId["vipObjectType"] = "object"
		siteId["vipType"] = "constant"
		siteId["vipValue"] = input["site_id"].(float64)
		siteId["vipVariableName"] = "system_site_id"
		defMap["site-id"] = siteId
	}

	if input["system_ip"] != nil && input["system_ip"] != "" {
		systemIp := make(map[string]interface{})
		systemIp["vipObjectType"] = "object"
		systemIp["vipType"] = "constant"
		systemIp["vipValue"] = input["system_ip"].(string)
		systemIp["vipVariableName"] = "system_system_ip"
		defMap["system-ip"] = systemIp
	}

	if input["overlay_id"] != nil {
		overlayId := make(map[string]interface{})
		overlayId["vipObjectType"] = "object"
		overlayId["vipType"] = "constant"
		overlayId["vipValue"] = input["overlay_id"].(int)
		overlayId["vipVariableName"] = "system_overlay_id"
		defMap["overlay-id"] = overlayId
	}

	if input["timezone"] != nil && input["timezone"] != "" {
		timezone := make(map[string]interface{})
		timezone["vipObjectType"] = "object"
		timezone["vipType"] = "constant"
		timezone["vipValue"] = input["timezone"].(string)
		timezone["vipVariableName"] = "system_timezone"
		defMap["clock"] = map[string]interface{}{
			"timezone": timezone,
		}
	}

	if input["hostname"] != nil && input["hostname"] != "" {
		hostname := make(map[string]interface{})
		hostname["vipObjectType"] = "object"
		hostname["vipType"] = "constant"
		hostname["vipValue"] = input["hostname"].(string)
		hostname["vipVariableName"] = "system_hostname"
		defMap["host-name"] = hostname
	}

	if input["location"] != nil && input["location"] != "" {
		location := make(map[string]interface{})
		location["vipObjectType"] = "object"
		location["vipType"] = "constant"
		location["vipValue"] = input["location"].(string)
		location["vipVariableName"] = "system_location"
		defMap["location"] = location
	}

	if input["device_groups"] != nil {
		if len(input["device_groups"].([]interface{})) > 0 {
			deviceGroups := make(map[string]interface{})
			deviceGroups["vipObjectType"] = "list"
			deviceGroups["vipType"] = "constant"
			deviceGroups["vipValue"] = input["device_groups"].([]interface{})
			deviceGroups["vipVariableName"] = "system_device_groups"
			defMap["device-groups"] = deviceGroups
		}
	}

	if input["controller_groups"] != nil {
		if len(input["controller_groups"].([]interface{})) > 0 {
			controllerGroups := make(map[string]interface{})
			controllerGroups["vipObjectType"] = "list"
			controllerGroups["vipType"] = "constant"
			controllerGroups["vipValue"] = input["controller_groups"].([]interface{})
			controllerGroups["vipVariableName"] = "system_controller_group_list"
			defMap["controller-group-list"] = controllerGroups
		}
	}

	if input["description"] != nil && input["description"] != "" {
		description := make(map[string]interface{})
		description["vipObjectType"] = "object"
		description["vipType"] = "constant"
		description["vipValue"] = input["description"].(string)
		description["vipVariableName"] = "system_description"
		defMap["description"] = description
	}

	if input["console_baud_rate"] != nil {
		consoleBaudRate := make(map[string]interface{})
		consoleBaudRate["vipObjectType"] = "object"
		consoleBaudRate["vipType"] = "constant"
		consoleBaudRate["vipValue"] = input["console_baud_rate"].(string)
		consoleBaudRate["vipVariableName"] = "system_console_baud_rate"
		defMap["console-baud-rate"] = consoleBaudRate
	}

	if input["max_omp_sessions"] != nil {
		maxOmpSessions := make(map[string]interface{})
		maxOmpSessions["vipObjectType"] = "object"
		maxOmpSessions["vipType"] = "constant"
		maxOmpSessions["vipValue"] = input["max_omp_sessions"].(int)
		maxOmpSessions["vipVariableName"] = "system_max_omp_sessions"
		defMap["max-omp-sessions"] = maxOmpSessions
	}

	if input["tcp_optimization_enabled"] != nil {
		tcpOptimizationEnabled := make(map[string]interface{})
		tcpOptimizationEnabled["vipObjectType"] = "object"
		tcpOptimizationEnabled["vipType"] = "constant"
		tcpOptimizationEnabled["vipValue"] = strconv.FormatBool(input["tcp_optimization_enabled"].(bool))
		tcpOptimizationEnabled["vipVariableName"] = "system_tcp_optimization_enabled"
		defMap["tcp-optimization-enabled"] = tcpOptimizationEnabled
	}

}

func createSystemGPS(defMap map[string]interface{}, input map[string]interface{}) {
	gps := make(map[string]interface{})
	if input["latitude"] != nil {
		latitude := make(map[string]interface{})
		latitude["vipObjectType"] = "object"
		latitude["vipType"] = "constant"
		latitude["vipValue"] = input["latitude"].(float64)
		latitude["vipVariableName"] = "system_latitude"
		gps["latitude"] = latitude
	}

	if input["longitude"] != nil {
		longitude := make(map[string]interface{})
		longitude["vipObjectType"] = "object"
		longitude["vipType"] = "constant"
		longitude["vipValue"] = input["longitude"].(float64)
		longitude["vipVariableName"] = "system_longitude"
		gps["longitude"] = longitude
	}
	defMap["gps-location"] = gps
}

func createSystemTracker(defMap map[string]interface{}, input map[string]interface{}) {

	if input["tracker_name"] != nil && input["tracker_name"] != "" {
		trackerName := make(map[string]interface{})
		trackerName["vipObjectType"] = "object"
		trackerName["vipType"] = "constant"
		trackerName["vipValue"] = input["tracker_name"].(string)
		trackerName["vipVariableName"] = "tracker_name"
		defMap["name"] = trackerName
	}

	if input["tracker_type"] != nil && input["tracker_type"] != "" {
		trackerEndpointValue := make(map[string]interface{})

		tType := input["tracker_type"].(string)
		if tType == "ip-address" {

			trackerEndpointValue["vipObjectType"] = "object"
			trackerEndpointValue["vipType"] = "constant"
			trackerEndpointValue["vipValue"] = input["tracker_endpoint_value"].(string)
			trackerEndpointValue["vipVariableName"] = "tracker_endpoint-ip"
			defMap["endpoint-ip"] = trackerEndpointValue

		} else if tType == "dns-name" {

			trackerEndpointValue["vipObjectType"] = "object"
			trackerEndpointValue["vipType"] = "constant"
			trackerEndpointValue["vipValue"] = input["tracker_endpoint_value"].(string)
			trackerEndpointValue["vipVariableName"] = "tracker_endpoint-dns-name"
			defMap["endpoint-dns-name"] = trackerEndpointValue

		} else if tType == "api-url" {

			trackerEndpointValue["vipObjectType"] = "object"
			trackerEndpointValue["vipType"] = "constant"
			trackerEndpointValue["vipValue"] = input["tracker_endpoint_value"].(string)
			trackerEndpointValue["vipVariableName"] = "tracker_endpoint-api-url"
			defMap["endpoint-api-url"] = trackerEndpointValue

		} else {
			fmt.Errorf("[ERROR] No such Tracker type found")
		}

	}

	if input["tracker_threshold"] != nil {
		trackerThreshold := make(map[string]interface{})
		trackerThreshold["vipObjectType"] = "object"
		trackerThreshold["vipType"] = "constant"
		trackerThreshold["vipValue"] = input["tracker_threshold"].(int)
		trackerThreshold["vipVariableName"] = "tracker_threshold"
		defMap["threshold"] = trackerThreshold
	}

	if input["tracker_interval"] != nil {
		trackerInterval := make(map[string]interface{})
		trackerInterval["vipObjectType"] = "object"
		trackerInterval["vipType"] = "constant"
		trackerInterval["vipValue"] = input["tracker_interval"].(int)
		trackerInterval["vipVariableName"] = "tracker_interval"
		defMap["interval"] = trackerInterval
	}

	if input["tracker_multiplier"] != nil {
		trackerMultiplier := make(map[string]interface{})
		trackerMultiplier["vipObjectType"] = "object"
		trackerMultiplier["vipType"] = "constant"
		trackerMultiplier["vipValue"] = input["tracker_multiplier"].(int)
		trackerMultiplier["vipVariableName"] = "tracker_multiplier"
		defMap["multiplier"] = trackerMultiplier
	}
	defMap["priority-order"] = []string{
		"name",
		"endpoint-ip",
		"endpoint-dns-name",
		"endpoint-api-url",
		"threshold",
		"interval",
		"multiplier",
	}
}

func createSystemAdvanced(defMap map[string]interface{}, input map[string]interface{}) {

	if input["control_session_pps"] != nil {
		controlSessionPps := make(map[string]interface{})
		controlSessionPps["vipObjectType"] = "object"
		controlSessionPps["vipType"] = "constant"
		controlSessionPps["vipValue"] = input["control_session_pps"].(int)
		controlSessionPps["vipVariableName"] = "system_control_session_pps"
		defMap["control-session-pps"] = controlSessionPps
	}

	if input["system_tunnel_mtu"] != nil {
		systemTunnelMtu := make(map[string]interface{})
		systemTunnelMtu["vipObjectType"] = "object"
		systemTunnelMtu["vipType"] = "constant"
		systemTunnelMtu["vipValue"] = input["system_tunnel_mtu"].(int)
		systemTunnelMtu["vipVariableName"] = "system_system_tunnel_mtu"
		defMap["system-tunnel-mtu"] = systemTunnelMtu
	}

	if input["port_hop"] != nil {
		portHop := make(map[string]interface{})
		portHop["vipObjectType"] = "object"
		portHop["vipType"] = "constant"
		portHop["vipValue"] = strconv.FormatBool(input["port_hop"].(bool))
		portHop["vipVariableName"] = "system_port_hop"
		defMap["port-hop"] = portHop
	}

	if input["port_offset"] != nil {
		portOffset := make(map[string]interface{})
		portOffset["vipObjectType"] = "object"
		portOffset["vipType"] = "constant"
		portOffset["vipValue"] = input["port_offset"].(int)
		portOffset["vipVariableName"] = "system_port_offset"
		defMap["port-offset"] = portOffset
	}

	if input["dns_cache_timeout"] != nil {
		dnsCacheTimeout := make(map[string]interface{})
		dnsCacheTimeout["vipObjectType"] = "object"
		dnsCacheTimeout["vipType"] = "constant"
		dnsCacheTimeout["vipValue"] = input["dns_cache_timeout"].(int)
		dnsCacheTimeout["vipVariableName"] = "system_dns_cache_timeout"
		defMap["timer"] = map[string]interface{}{
			"dns-cache-timeout": dnsCacheTimeout,
		}
	}

	if input["track_transport"] != nil {
		trackTransport := make(map[string]interface{})
		trackTransport["vipObjectType"] = "object"
		trackTransport["vipType"] = "constant"
		trackTransport["vipValue"] = strconv.FormatBool(input["track_transport"].(bool))
		trackTransport["vipVariableName"] = "system_track_transport"
		defMap["track-transport"] = trackTransport
	}

	if input["vbond_local"] != nil {
		vbondLocal := make(map[string]interface{})
		vbond := make(map[string]interface{})
		vbondLocal["vipObjectType"] = "object"
		vbondLocal["vipType"] = "constant"
		vbondLocal["vipValue"] = strconv.FormatBool(input["vbond_local"].(bool))
		vbondLocal["vipVariableName"] = "system_vbond_local"
		vbond["local"] = vbondLocal

		if input["vbond_remote"] != nil && input["vbond_remote"] != "" {
			if input["vbond_local"].(bool) {
				vbondRemote := make(map[string]interface{})
				vbondRemote["vipObjectType"] = "object"
				vbondRemote["vipType"] = "constant"
				vbondRemote["vipValue"] = input["vbond_remote"].(string)
				vbondRemote["vipVariableName"] = "system_vbond_remote"
				vbond["remote"] = vbondRemote
			} else {
				log.Panic("[ERROR] vbond_local should be true to set vbond_remote")
			}
		}
		defMap["vbond"] = vbond
	}

	if input["track_interface_tag"] != nil {
		trackInterfaceTag := make(map[string]interface{})
		trackInterfaceTag["vipObjectType"] = "object"
		trackInterfaceTag["vipType"] = "constant"
		trackInterfaceTag["vipValue"] = input["track_interface_tag"].(float64)
		trackInterfaceTag["vipVariableName"] = "system_track_interface_tag"
		defMap["track-interface-tag"] = trackInterfaceTag
	}

	if input["multicast_buffer_percent"] != nil {
		multicastBufferPercent := make(map[string]interface{})
		multicastBufferPercent["vipObjectType"] = "object"
		multicastBufferPercent["vipType"] = "constant"
		multicastBufferPercent["vipValue"] = input["multicast_buffer_percent"].(int)
		multicastBufferPercent["vipVariableName"] = "system_multicast_buffer_percent"
		defMap["multicast-buffer-percent"] = multicastBufferPercent
	}

	if input["usb_controller"] != nil {
		usbController := make(map[string]interface{})
		usbController["vipObjectType"] = "object"
		usbController["vipType"] = "constant"
		usbController["vipValue"] = strconv.FormatBool(input["usb_controller"].(bool))
		usbController["vipVariableName"] = "system_usb_controller"
		defMap["usb-controller"] = usbController
	}

	if input["track_default_gateway"] != nil {
		trackDefaultGateway := make(map[string]interface{})
		trackDefaultGateway["vipObjectType"] = "object"
		trackDefaultGateway["vipType"] = "constant"
		trackDefaultGateway["vipValue"] = strconv.FormatBool(input["track_default_gateway"].(bool))
		trackDefaultGateway["vipVariableName"] = "system_track_default_gateway"
		defMap["track-default-gateway"] = trackDefaultGateway
	}

	if input["host_policer_pps"] != nil {
		hostPolicerPps := make(map[string]interface{})
		hostPolicerPps["vipObjectType"] = "object"
		hostPolicerPps["vipType"] = "constant"
		hostPolicerPps["vipValue"] = input["host_policer_pps"].(int)
		hostPolicerPps["vipVariableName"] = "system_host_policer_pps"
		defMap["host-policer-pps"] = hostPolicerPps
	}

	if input["icmp_error_pps"] != nil {
		icmpErrorPps := make(map[string]interface{})
		icmpErrorPps["vipObjectType"] = "object"
		icmpErrorPps["vipType"] = "constant"
		icmpErrorPps["vipValue"] = input["icmp_error_pps"].(int)
		icmpErrorPps["vipVariableName"] = "system_icmp_error_pps"
		defMap["icmp-error-pps"] = icmpErrorPps
	}

	if input["allow_same_site_tunnels"] != nil {
		allowSameSiteTunnels := make(map[string]interface{})
		allowSameSiteTunnels["vipObjectType"] = "object"
		allowSameSiteTunnels["vipType"] = "constant"
		allowSameSiteTunnels["vipValue"] = strconv.FormatBool(input["allow_same_site_tunnels"].(bool))
		allowSameSiteTunnels["vipVariableName"] = "system_allow_same_site_tunnels"
		defMap["allow-same-site-tunnels"] = allowSameSiteTunnels
	}

	if input["route_consistency_check"] != nil {
		routeConsistencyCheck := make(map[string]interface{})
		routeConsistencyCheck["vipObjectType"] = "object"
		routeConsistencyCheck["vipType"] = "constant"
		routeConsistencyCheck["vipValue"] = strconv.FormatBool(input["route_consistency_check"].(bool))
		routeConsistencyCheck["vipVariableName"] = "system_route_consistency_check"
		defMap["route-consistency-check"] = routeConsistencyCheck
	}

	if input["admin_tech_on_failure"] != nil {
		adminTechOnFailure := make(map[string]interface{})
		adminTechOnFailure["vipObjectType"] = "object"
		adminTechOnFailure["vipType"] = "constant"
		adminTechOnFailure["vipValue"] = strconv.FormatBool(input["admin_tech_on_failure"].(bool))
		adminTechOnFailure["vipVariableName"] = "system_admin_tech_on_failure"
		defMap["admin-tech-on-failure"] = adminTechOnFailure
	}

	if input["idle_timeout"] != nil {
		idleTimeout := make(map[string]interface{})
		idleTimeout["vipObjectType"] = "object"
		idleTimeout["vipType"] = "constant"
		idleTimeout["vipValue"] = input["idle_timeout"].(int)
		idleTimeout["vipVariableName"] = "system_idle_timeout"
		defMap["idle-timeout"] = idleTimeout
	}

	if input["eco_friendly_mode"] != nil {
		ecoFriendlyMode := make(map[string]interface{})
		ecoFriendlyMode["vipObjectType"] = "object"
		ecoFriendlyMode["vipType"] = "constant"
		ecoFriendlyMode["vipValue"] = strconv.FormatBool(input["eco_friendly_mode"].(bool))
		ecoFriendlyMode["vipVariableName"] = "system_eco_friendly_mode"
		defMap["eco-friendly-mode"] = ecoFriendlyMode
	}

	onDemand := make(map[string]interface{})
	if input["on_demand_enable"] != nil {
		onDemandEnable := make(map[string]interface{})
		onDemandEnable["vipObjectType"] = "object"
		onDemandEnable["vipType"] = "constant"
		onDemandEnable["vipValue"] = strconv.FormatBool(input["on_demand_enable"].(bool))
		onDemandEnable["vipVariableName"] = "system_on_demand_enable"
		onDemand["enable"] = onDemandEnable
	}

	if input["on_demand_idle_timeout"] != nil {
		onDemandIdleTimeout := make(map[string]interface{})
		onDemandIdleTimeout["vipObjectType"] = "object"
		onDemandIdleTimeout["vipType"] = "constant"
		onDemandIdleTimeout["vipValue"] = input["on_demand_idle_timeout"].(int)
		onDemandIdleTimeout["vipVariableName"] = "system_on_demand_idle_timeout"
		onDemand["idle-timeout"] = onDemandIdleTimeout
	}
	defMap["on-demand"] = onDemand

}

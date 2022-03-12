package sdwan

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var (
	validDeviceTypeforVPN = []string{
		"vedge-1000",
		"vedge-2000",
		"vedge-cloud",
		"vedge-5000",
		"vedge-ISR1100-6G",
		"vedge-100-B",
		"vedge-ISR1100-4G",
		"vedge-100",
		"vedge-ISR1100-4GLTE",
		"vedge-100-WM",
		"vbond",
		"vedge-100-M",
		"vedge-ISR1100X-6G",
		"vedge-ISR1100X-4G",
		"vedge-C8500-12X4QC",
		"vedge-C1111-4PLTEEA",
		"vedge-C1161-8P",
		"vedge-C1117-4PLTEEAW",
		"vedge-C1121X-8P",
		"vedge-C8200-1N-4T",
		"vedge-ISR-4331",
		"vedge-ISRv",
		"vedge-C1127X-8PMLTEP",
		"vedge-C1117-4PMLTEEAWE",
		"vedge-ISR-4451-X",
		"vedge-C1113-8PLTEEA",
		"vedge-IR-1821",
		"vedge-ASR-1001-X",
		"vedge-ISR-4321",
		"vedge-C1116-4PLTEEAWE",
		"vedge-C1109-4PLTE2P",
		"vedge-C1121-8P",
		"vedge-ASR-1002-HX",
		"vedge-C1111-8PLTEEAW",
		"vedge-C1112-8PWE",
		"vedge-C1101-4PLTEP",
		"vedge-ISR1100-4GLTENA-XE",
		"vedge-C1111-8PLTELA",
		"vedge-IR-1835",
		"vedge-C1121X-8PLTEP",
		"vedge-IR-1833",
		"vedge-C8300-1N1S-4T2X",
		"vedge-C1121-4P",
		"vedge-ISR-4351",
		"vedge-C1117-4PLTELA",
		"vedge-C1116-4PWE",
		"vedge-nfvis-C8200-UCPE",
		"vedge-C1113-8PM",
		"vedge-IR-1831",
		"vedge-C1127-8PLTEP",
		"vedge-C1121-8PLTEPW",
		"vedge-C1113-8PW",
		"vedge-ASR-1001-HX",
		"vedge-C1128-8PLTEP",
		"vedge-C1113-8PLTEEAW",
		"vedge-C1117-4PW",
		"vedge-C1116-4P",
		"vedge-C1113-8PMLTEEA",
		"vedge-C1112-8P",
		"vedge-ISR-4461",
		"vedge-C1116-4PLTEEA",
		"vedge-ISR-4221",
		"vedge-C1117-4PM",
		"vedge-C1113-8PLTELAWZ",
		"vedge-C1117-4PMWE",
		"vedge-C1109-2PLTEVZ",
		"vedge-C1113-8P",
		"vedge-C1117-4P",
		"vedge-nfvis-ENCS5400",
		"vedge-C8300-2N2S-6T",
		"vedge-C1127-8PMLTEP",
		"vedge-ESR-6300",
		"vedge-ISR-4221X",
		"vedge-ISR1100-4GLTEGB-XE",
		"vedge-C8500-12X",
		"vedge-C1109-2PLTEGB",
		"vedge-CSR-1000v",
		"vedge-C1113-8PLTEW",
		"vedge-C1121X-8PLTEPW",
		"vedge-ISR1100-6G-XE",
		"vedge-C1121-4PLTEP",
		"vedge-C1111-8PLTEEA",
		"vedge-C1117-4PLTEEA",
		"vedge-C1127X-8PLTEP",
		"vedge-C1109-2PLTEUS",
		"vedge-C1112-8PLTEEAWE",
		"vedge-C1161X-8P",
		"vedge-C8500L-8S4X",
		"vedge-C1111-8PW",
		"vedge-C1161X-8PLTEP",
		"vedge-C1101-4PLTEPW",
		"vedge-ISR1100X-4G-XE",
		"vedge-IR-1101",
		"vedge-C1111-4P",
		"vedge-C1111-4PW",
		"vedge-C1111-8P",
		"vedge-C1117-4PMLTEEA",
		"vedge-C1113-8PLTELA",
		"vedge-C1111-8PLTELAW",
		"vedge-C1161-8PLTEP",
		"vedge-ISR1100X-6G-XE",
		"vedge-ISR-4431",
		"vedge-C1101-4P",
		"vedge-C1109-4PLTE2PW",
		"vedge-C1113-8PMWE",
		"vedge-C1118-8P",
		"vedge-C1126-8PLTEP",
		"vedge-C8300-1N1S-6T",
		"vedge-C1121-8PLTEP",
		"vedge-C8300-2N2S-4T2X",
		"vedge-C1112-8PLTEEA",
		"vedge-C1111-4PLTELA",
		"vedge-ASR-1002-X",
		"vedge-C1111X-8P",
		"vedge-C1126X-8PLTEP",
		"vedge-ASR-1006-X",
		"vedge-C8000V",
		"vedge-ISR1100-4G-XE",
		"vedge-C1117-4PLTELAWZ"}
	vpnMinValidId        = 0
	vpnMaxValidId        = 512
	allowedGetewayTypev4 = []string{"next-hop", "dhcp", "null0"}
	allowedGatewayTypev6 = []string{"next-hop", "null0"}
	isValidServiceType   = []string{"sig", "FW", "IDS", "IDP", "netsvc1", "netsvc2", "netsvc3", "netsvc4", "TE"}
	validVPNFTTypes      = []string{"vpn-vedge", "cisco_vpn"}
)

func resouceSDWANVPNFeatureTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSDWANVPNFeatureTemplateCreate,
		Update: resourceSDWANVPNFeatureTemplateUpdate,
		Read:   resouceSDWANVPNFeatureTemplateRead,
		Delete: resourceSDWANVPNFeatureTemplateDelete,

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"template_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: NameValidation(),
				ForceNew:     true,
			},
			"template_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice(validVPNFTTypes, false),
			},
			"template_description": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: DescValidation(),
			},
			"device_type": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice(validDeviceTypeforVPN, false),
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
						"vpn_basic": {
							Type:     schema.TypeSet,
							Required: true,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpn_id": {
										Type:     schema.TypeInt,
										Required: true,
										ValidateFunc: validation.All(
											validation.IntAtLeast(vpnMinValidId),
											validation.IntAtMost(vpnMaxValidId),
										),
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ecmp_hash_key": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
								},
							},
						},

						"vpn_dns": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"primary_dns_ipv4": {
										Type:         schema.TypeString,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.IsIPv4Address,
									},
									"secondary_dns_ipv4": {
										Type:         schema.TypeString,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.IsIPv4Address,
									},
									"primary_dns_ipv6": {
										Type:         schema.TypeString,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.IsIPv6Address,
									},
									"secondary_dns_ipv6": {
										Type:         schema.TypeString,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.IsIPv6Address,
									},
									"vpn_host": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hostname": {
													Type:         schema.TypeString,
													Required:     true,
													ValidateFunc: validation.StringLenBetween(1, 128),
												},
												"ip": {
													Type:     schema.TypeList,
													Required: true,
													MaxItems: 8,
													MinItems: 1,
													Elem: &schema.Schema{
														Type:         schema.TypeString,
														ValidateFunc: validation.IsIPAddress,
													},
												},
											},
										},
									},
								},
							},
						},

						"ipv4_route": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.IsCIDR,
									},
									"gateway_type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice(allowedGetewayTypev4, false),
									},
									"next_hop": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"next_hop_address": {
													Type:         schema.TypeString,
													Required:     true,
													ValidateFunc: validation.IsIPv4Address,
												},
												"next_hop_distance": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ValidateDiagFunc: isStringInRange(1, 255),
												},
											},
										},
									},
									"null0_distance": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ValidateDiagFunc: isStringInRange(1, 255),
									},
								},
							},
						},

						"ipv6_route": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: IPv6SubNetValidate(),
									},
									"gateway_type": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice(allowedGatewayTypev6, false),
									},
									"next_hop": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"next_hop_address": {
													Type:         schema.TypeString,
													Required:     true,
													ValidateFunc: validation.IsIPv6Address,
												},
												"next_hop_distance": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ValidateDiagFunc: isStringInRange(1, 255),
												},
											},
										},
									},
									"null0_distance": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ValidateDiagFunc: isStringInRange(1, 255),
									},
								},
							},
						},

						"te_service_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},

						"service_route": {
							Type:     schema.TypeSet,
							Computed: true,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"prefix": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.IsCIDR,
									},
									"service_type": {
										Type:         schema.TypeString,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.StringInSlice(isValidServiceType, false),
									},
								},
							},
						},

						"nat64": {
							Type:     schema.TypeSet,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pool_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"start_address": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.IsIPv4Address,
									},
									"end_address": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.IsIPv4Address,
									},
									"overload": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
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

func resourceSDWANVPNFeatureTemplateCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Create method")

	sdwanClient := m.(*client.Client)

	ftMap := make(map[string]interface{})

	ftMap["templateType"] = d.Get("template_type").(string)
	ftMap["templateName"] = d.Get("template_name").(string)
	ftMap["templateDescription"] = d.Get("template_description").(string)
	if !verifyDeviceTypeVPN(ftMap["templateType"].(string), interfaceToStrList(d.Get("device_type"))) {
		return fmt.Errorf("Please use a valid Device Type for Template: %s", ftMap["templateType"].(string))
	}
	ftMap["deviceType"] = interfaceToStrList(d.Get("device_type"))
	ftMap["templateMinVersion"] = d.Get("template_min_version").(string)
	ftMap["factoryDefault"] = d.Get("factory_default").(bool)

	ftdefinition := d.Get("template_definition").(*schema.Set).List()

	if def, err := createVPNFTdefinition(ftdefinition, ftMap["templateType"].(string)); err == nil {
		ftMap["templateDefinition"] = def
	} else {
		return err
	}

	dURL := fmt.Sprintf("dataservice/template/feature")
	bodyBytes, err := marshalJSON(ftMap)
	if err != nil {
		return err
	}

	cont, err := sdwanClient.SaveviaBytes(dURL, bodyBytes)
	if err != nil {
		return err
	}

	ftID := stripQuotes(cont.S("templateId").String())
	d.Set("template_id", ftID)
	d.SetId(ftID)

	return resouceSDWANVPNFeatureTemplateRead(d, m)
}

func resourceSDWANVPNFeatureTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Update Method ", d.Id())

	sdwanClient := m.(*client.Client)

	ftID := d.Id()

	ftMap := make(map[string]interface{})

	ftMap["templateType"] = d.Get("template_type").(string)
	ftMap["templateName"] = d.Get("template_name").(string)
	ftMap["templateDescription"] = d.Get("template_description").(string)
	if !verifyDeviceTypeVPN(ftMap["templateType"].(string), interfaceToStrList(d.Get("device_type"))) {
		return fmt.Errorf("Please use a valid Device Type for Template: %s", ftMap["templateType"].(string))
	}
	ftMap["deviceType"] = interfaceToStrList(d.Get("device_type"))
	ftMap["templateMinVersion"] = d.Get("template_min_version").(string)
	ftMap["factoryDefault"] = d.Get("factory_default").(bool)

	ftdefinition := d.Get("template_definition").(*schema.Set).List()

	if def, err := createVPNFTdefinition(ftdefinition, ftMap["templateType"].(string)); err == nil {
		ftMap["templateDefinition"] = def
	} else {
		return err
	}

	bodyBytes, err := marshalJSON(ftMap)
	if err != nil {
		return err
	}

	dURL := fmt.Sprintf("dataservice/template/feature/%s", ftID)
	_, err = sdwanClient.UpdateviaBytes(dURL, bodyBytes)
	if err != nil {
		return err
	}

	d.SetId(ftID)

	log.Println("[DEBUG] End of Update Method ", d.Id())
	return resouceSDWANVPNFeatureTemplateRead(d, m)
}

func resouceSDWANVPNFeatureTemplateRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Read Method ", d.Id())

	sdwanClient := m.(*client.Client)

	tID := d.Id()

	dURL := fmt.Sprintf("dataservice/template/feature/object/%v", tID)

	data, err := sdwanClient.GetViaURL(dURL)

	if err != nil {
		return err
	}

	d.Set("template_name", stripQuotes(data.S("templateName").String()))

	d.Set("template_description", stripQuotes(data.S("templateDescription").String()))

	d.Set("template_min_version", stripQuotes(data.S("templateMinVersion").String()))

	d.Set("template_type", stripQuotes(data.S("templateType").String()))

	d.Set("factory_default", data.S("factoryDefault").Data().(bool))

	d.Set("device_type", interfaceToStrList(data.S("deviceType").Data()))

	if data.Exists("attachedMastersCount") {
		d.Set("attached_masters_count", int((data.S("attachedMastersCount").Data()).(float64)))
	}

	if data.Exists("devicesAttached") {
		d.Set("devices_attached", int((data.S("devicesAttached").Data()).(float64)))
	}

	if data.Exists("lastUpdatedBy") {
		d.Set("last_updated_by", stripQuotes(data.S("lastUpdatedBy").String()))
	}

	if data.Exists("lastUpdatedOn") {
		d.Set("last_updated_on", int((data.S("lastUpdatedOn").Data().(float64))))
	}

	if data.Exists("gTemplateClass") {
		d.Set("g_template_class", stripQuotes(data.S("gTemplateClass").String()))
	}

	if data.Exists("configType") {
		d.Set("config_type", stripQuotes(data.S("configType").String()))
	}

	if data.Exists("createdOn") {
		d.Set("created_on", int((data.S("createdOn").Data()).(float64)))
	}

	if data.Exists("createdBy") {
		d.Set("created_by", stripQuotes(data.S("createdBy").String()))
	}

	if data.Exists("feature") {
		d.Set("feature", stripQuotes(data.S("feature").String()))
	}

	d.Set("template_id", stripQuotes(data.S("templateId").String()))

	if data.Exists("@rid") {
		d.Set("rid", int((data.S("@rid").Data()).(float64)))
	}

	ftDef := data.S("templateDefinition")
	FTdefinition := getVPNdefinition(ftDef)

	d.Set("template_definition", FTdefinition)

	d.SetId(stripQuotes(data.S("templateId").String()))

	log.Println("[DEBUG] End of Read Method ", d.Id())
	return nil
}

func resourceSDWANVPNFeatureTemplateDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Delete Method ", d.Id())

	sdwanClient := m.(*client.Client)
	ftID := d.Id()

	dURL := fmt.Sprintf("/dataservice/template/feature/%s", ftID)
	_, err := sdwanClient.Delete(dURL)
	if err != nil {
		return err
	}

	d.SetId("")

	log.Println("[DEBUG] End of Delete Method!")
	return nil
}

func createVPNFTdefinition(ftDefinitions []interface{}, ftType string) (map[string]interface{}, error) {
	definition := make(map[string]interface{})

	ftDefinition := ftDefinitions[0]
	ftDefMap := ftDefinition.(map[string]interface{})
	VPNBasicMap := (ftDefMap["vpn_basic"].(*schema.Set).List())[0].(map[string]interface{})
	createVPNBasic(definition, VPNBasicMap, ftType)

	createRouteImport(definition)

	createRouteExport(definition)

	createOMP(definition)

	definition["nat64-global"] = map[string]interface{}{
		"prefix": map[string]interface{}{
			"stateful": map[string]interface{}{},
		},
	}

	DNSset := ftDefMap["vpn_dns"].(*schema.Set).List()

	if len(DNSset) > 0 {
		vpnDNSMap := DNSset[0].(map[string]interface{})
		err := createVPNDNS(definition, vpnDNSMap)

		if err != nil {
			return nil, err
		}
	} else {
		hostMap := make(map[string]interface{})

		hostMap["vipType"] = "ignore"
		hostMap["vipValue"] = make([]map[string]interface{}, 0, 1)
		hostMap["vipObjectType"] = "tree"
		hostMap["vipPrimaryKey"] = []string{"hostname"}
		definition["host"] = hostMap
	}

	ipMap := make(map[string]interface{})

	if ftDefMap["ipv4_route"] != nil {
		ipv4routes := ftDefMap["ipv4_route"].(*schema.Set).List()

		if len(ipv4routes) > 0 {
			ipv4routesVal := make([]map[string]interface{}, 0, len(ipv4routes))

			for _, ipv4routeMap := range ipv4routes {
				ipv4route := make(map[string]interface{})

				err := createVPNIPv4Routes(ipv4route, ipv4routeMap.(map[string]interface{}))

				if err != nil {
					return nil, err
				}
				ipv4routesVal = append(ipv4routesVal, ipv4route)
			}

			routeMap := make(map[string]interface{})

			routeMap["vipType"] = "constant"
			routeMap["vipObjectType"] = "tree"
			routeMap["vipPrimaryKey"] = []string{"prefix"}
			routeMap["vipValue"] = ipv4routesVal
			ipMap["route"] = routeMap
		}
	}

	if ftDefMap["ipv6_route"] != nil {
		ipv6routes := ftDefMap["ipv6_route"].(*schema.Set).List()

		if len(ipv6routes) > 0 {
			ipv6routesVal := make([]map[string]interface{}, 0, len(ipv6routes))

			for _, ipv6routeMap := range ipv6routes {
				ipv6route := make(map[string]interface{})
				err := createVPNIPv6Routes(ipv6route, ipv6routeMap.(map[string]interface{}), ftType)

				if err != nil {
					return nil, err
				}
				ipv6routesVal = append(ipv6routesVal, ipv6route)
			}
			routeMap := make(map[string]interface{})
			routeMap["vipType"] = "constant"
			routeMap["vipValue"] = ipv6routesVal
			routeMap["vipObjectType"] = "tree"
			routeMap["vipPrimaryKey"] = []string{"prefix"}

			definition["ipv6"] = map[string]interface{}{
				"route": routeMap,
			}
		} else {
			definition["ipv6"] = make(map[string]interface{})
		}
	}

	if ftDefMap["te_service_enabled"] == true && VPNBasicMap["vpn_id"] == 512 {
		return nil, fmt.Errorf("TE service cannot be enabled in VPN512")
	} else if ftDefMap["te_service_enabled"] == true {
		enableTEService(definition)
	} else {
		createDefaultService(definition)
	}

	if ftDefMap["service_route"] != nil {
		serviceRoutes := ftDefMap["service_route"].(*schema.Set).List()

		if len(serviceRoutes) > 0 {
			serviceroutesVal := make([]map[string]interface{}, 0, 1)

			for _, serviceRouteMap := range serviceRoutes {
				serviceRoute := make(map[string]interface{})
				err := createServiceRoutes(serviceRoute, serviceRouteMap.(map[string]interface{}), ftType)
				if err != nil {
					return nil, err
				}
				serviceroutesVal = append(serviceroutesVal, serviceRoute)
			}

			serviceRouteMap := make(map[string]interface{})
			serviceRouteMap["vipType"] = "constant"
			serviceRouteMap["vipValue"] = serviceroutesVal
			serviceRouteMap["vipObjectType"] = "tree"
			serviceRouteMap["vipPrimaryKey"] = []string{"prefix"}

			ipMap["service-route"] = serviceRouteMap
			ipMap["ipsec-route"] = make(map[string]interface{})
			ipMap["gre-route"] = make(map[string]interface{})
		} else {
			ipMap["service-route"] = make(map[string]interface{})
			ipMap["ipsec-route"] = make(map[string]interface{})
			ipMap["gre-route"] = make(map[string]interface{})
		}

	}

	definition["ip"] = ipMap

	if ftDefMap["nat64"] != nil {
		nats := ftDefMap["nat64"].(*schema.Set).List()

		if len(nats) > 0 {
			natsVal := make([]map[string]interface{}, 0, 1)

			for _, natMap := range nats {
				nat := make(map[string]interface{})
				err := createNATs(nat, natMap.(map[string]interface{}), ftType)
				if err != nil {
					return nil, err
				}
				natsVal = append(natsVal, nat)
			}

			natsMap := make(map[string]interface{})
			natsMap["vipType"] = "constant"
			natsMap["vipValue"] = natsVal
			natsMap["vipObjectType"] = "tree"
			natsMap["vipPrimaryKey"] = []string{"name"}

			definition["nat64"] = map[string]interface{}{
				"v4": map[string]interface{}{
					"pool": natsMap,
				},
			}
		} else {
			natsMap := make(map[string]interface{})
			natsMap["vipType"] = "ignore"
			natsMap["vipValue"] = make([]map[string]interface{}, 0, 1)
			natsMap["vipObjectType"] = "tree"
			natsMap["vipPrimaryKey"] = []string{"name"}
			definition["nat64"] = map[string]interface{}{
				"v4": map[string]interface{}{
					"pool": natsMap,
				},
			}
		}
	}

	return definition, nil
}

func createVPNBasic(defMap map[string]interface{}, input map[string]interface{}, ftType string) {

	if input["vpn_id"] != nil {
		vpnId := make(map[string]interface{})
		vpnId["vipObjectType"] = "object"
		vpnId["vipType"] = "constant"
		vpnId["vipValue"] = input["vpn_id"]
		defMap["vpn-id"] = vpnId
	}

	if input["name"] != nil && input["name"] != "" {
		vpnName := make(map[string]interface{})
		vpnName["vipObjectType"] = "object"
		vpnName["vipType"] = "constant"
		vpnName["vipValue"] = input["name"]
		vpnName["vipVariableName"] = "vpn_name"
		defMap["name"] = vpnName
	} else {
		vpnName := make(map[string]interface{})
		vpnName["vipObjectType"] = "object"
		vpnName["vipType"] = "ignore"
		vpnName["vipVariableName"] = "vpn_name"
		defMap["name"] = vpnName
	}

	if input["ecmp_hash_key"] != nil {
		ehk := make(map[string]interface{})
		ehk["vipObjectType"] = "object"
		ehk["vipType"] = "constant"
		ehk["vipValue"] = strconv.FormatBool(input["ecmp_hash_key"].(bool))
		ehk["vipVariableName"] = "vpn_layer4"
		layer := make(map[string]interface{})
		defMap["ecmp-hash-key"] = layer
		layer["layer4"] = ehk
	}

	if ftType == "vpn-vedge" {
		tcp := make(map[string]interface{})
		tcp["vipObjectType"] = "node-only"
		tcp["vipType"] = "ignore"
		tcp["vipValue"] = "false"
		tcp["vipVariableName"] = "vpn_tcp_optimization"
		defMap["tcp-optimization"] = tcp
	} else {
		nat := map[string]interface{}{
			"natpool": map[string]interface{}{
				"vipObjectType": "tree",
				"vipPrimaryKey": []string{"name"},
				"vipType":       "ignore",
				"vipValue":      make([]map[string]interface{}, 0, 1),
			},
			"port-forward": map[string]interface{}{
				"vipObjectType": "tree",
				"vipPrimaryKey": []string{"source-port", "translate-port"},
				"vipType":       "ignore",
				"vipValue":      make([]map[string]interface{}, 0, 1),
			},
			"static": map[string]interface{}{
				"vipObjectType": "tree",
				"vipPrimaryKey": []string{"source-ip", "translate-ip"},
				"vipType":       "ignore",
				"vipValue":      make([]map[string]interface{}, 0, 1),
			},
		}
		defMap["nat"] = nat
	}
}

func createRouteExport(defMap map[string]interface{}) {
	re := make(map[string]interface{})
	re["vipObjectType"] = "tree"
	re["vipType"] = "ignore"
	re["vipValue"] = make([]map[string]interface{}, 0, 1)
	re["vipPrimaryKey"] = []string{"protocol"}
	defMap["route-export"] = re
}

func createOMP(defMap map[string]interface{}) {
	adv := make(map[string]interface{})
	adv["vipType"] = "ignore"
	adv["vipValue"] = make([]map[string]interface{}, 0, 1)
	adv["vipObjectType"] = "tree"
	adv["vipPrimaryKey"] = []string{"protocol"}

	v6adv := make(map[string]interface{})
	v6adv["vipType"] = "ignore"
	v6adv["vipValue"] = make([]map[string]interface{}, 0, 1)
	v6adv["vipObjectType"] = "tree"
	v6adv["vipPrimaryKey"] = []string{"protocol"}

	defMap["omp"] = map[string]interface{}{
		"advertise":      adv,
		"ipv6-advertise": v6adv,
	}
}

func createRouteImport(defMap map[string]interface{}) {
	ri := make(map[string]interface{})
	ri["vipObjectType"] = "tree"
	ri["vipType"] = "ignore"
	ri["vipValue"] = make([]map[string]interface{}, 0, 1)
	ri["vipPrimaryKey"] = []string{"protocol"}
	defMap["route-import"] = ri
}

func createDefaultService(defMap map[string]interface{}) {
	outerMap := make(map[string]interface{})

	outerMap["vipType"] = "ignore"
	outerMap["vipObjectType"] = "tree"
	outerMap["vipPrimaryKey"] = []string{"svc-type"}
	outerMap["vipValue"] = make([]map[string]interface{}, 0, 1)
	defMap["service"] = outerMap
}

func createVPNDNS(defMap map[string]interface{}, input map[string]interface{}) error {

	dns := make(map[string]interface{})

	dnsval := make([]map[string]interface{}, 0, 4)

	cnt := 0

	if input["primary_dns_ipv4"] != nil && input["primary_dns_ipv4"] != "" {
		dnsobj := make(map[string]interface{})

		role := make(map[string]interface{})
		role["vipType"] = "constant"
		role["vipValue"] = "primary"
		role["vipObjectType"] = "object"
		dnsobj["role"] = role

		address := make(map[string]interface{})
		address["vipType"] = "constant"
		address["vipValue"] = input["primary_dns_ipv4"]
		address["vipObjectType"] = "object"
		dnsobj["dns-addr"] = address

		dnsobj["priority-order"] = []string{"dns-addr", "role"}

		dnsval = append(dnsval, dnsobj)

		cnt++

		if input["secondary_dns_ipv4"] != nil && input["secondary_dns_ipv4"] != "" {
			dnsobj := make(map[string]interface{})
			role := make(map[string]interface{})
			role["vipType"] = "constant"
			role["vipValue"] = "secondary"
			role["vipObjectType"] = "object"
			dnsobj["role"] = role

			address := make(map[string]interface{})
			address["vipType"] = "constant"
			address["vipValue"] = input["secondary_dns_ipv4"]
			address["vipObjectType"] = "object"
			dnsobj["dns-addr"] = address

			dnsobj["priority-order"] = []string{"dns-addr", "role"}

			dnsval = append(dnsval, dnsobj)

			cnt++
		}

	} else if input["secondary_dns_ipv4"] != nil && input["secondary_dns_ipv4"] != "" {
		return fmt.Errorf("secondary_dns_ipv4 should be provided if and only if primary_dns_ipv4 is set")
	}

	if input["primary_dns_ipv6"] != nil && input["primary_dns_ipv6"] != "" {
		dnsobj := make(map[string]interface{})
		role := make(map[string]interface{})
		role["vipType"] = "constant"
		role["vipValue"] = "primaryv6"
		role["vipObjectType"] = "object"
		role["ipType"] = "ipv6"
		dnsobj["role"] = role

		address := make(map[string]interface{})
		address["vipType"] = "constant"
		address["vipValue"] = input["primary_dns_ipv6"]
		address["vipObjectType"] = "object"
		address["ipType"] = "ipv6"
		dnsobj["dns-addr"] = address

		dnsobj["priority-order"] = []string{"dns-addr", "role"}

		dnsval = append(dnsval, dnsobj)

		cnt++

		if input["secondary_dns_ipv6"] != nil && input["secondary_dns_ipv6"] != "" {
			dnsobj := make(map[string]interface{})
			role := make(map[string]interface{})
			role["vipType"] = "constant"
			role["vipValue"] = "secondaryv6"
			role["vipObjectType"] = "object"
			role["ipType"] = "ipv6"
			dnsobj["role"] = role

			address := make(map[string]interface{})
			address["vipType"] = "constant"
			address["vipValue"] = input["secondary_dns_ipv6"]
			address["vipObjectType"] = "object"
			address["ipType"] = "ipv6"
			dnsobj["dns-addr"] = address

			dnsobj["priority-order"] = []string{"dns-addr", "role"}

			dnsval = append(dnsval, dnsobj)

			cnt++
		}
	} else if input["secondary_dns_ipv6"] != nil && input["secondary_dns_ipv6"] != "" {
		return fmt.Errorf("secondary_dns_ipv6 should be provided if and only if primary_dns_ipv6 is set")
	}

	if cnt > 0 {
		dns["vipType"] = "constant"
		dns["vipValue"] = dnsval
		dns["vipObjectType"] = "tree"
		dns["vipPrimaryKey"] = []string{"dns-addr"}
		defMap["dns"] = dns
	}

	if input["vpn_host"] != nil {

		hostList := input["vpn_host"].(*schema.Set).List()
		if len(hostList) > 0 {
			hostVal := make([]map[string]interface{}, 0, 1)

			for _, hostMap := range hostList {
				host := make(map[string]interface{})
				createVPNHost(host, hostMap.(map[string]interface{}))

				hostVal = append(hostVal, host)
			}

			hostMap := make(map[string]interface{})
			hostMap["vipType"] = "constant"
			hostMap["vipValue"] = hostVal
			hostMap["vipObjectType"] = "tree"
			hostMap["vipPrimaryKey"] = []string{"hostname"}
			defMap["host"] = hostMap
		}
	}

	return nil
}

func createVPNHost(defMap map[string]interface{}, input map[string]interface{}) {
	if input["hostname"] != nil {
		hostName := make(map[string]interface{})
		hostName["vipObjectType"] = "object"
		hostName["vipType"] = "constant"
		hostName["vipValue"] = input["hostname"]
		hostName["vipVariableName"] = "vpn_host_hostname"
		defMap["hostname"] = hostName
	}

	ipMap := make(map[string]interface{})
	ipMap["vipObjectType"] = "list"
	ipMap["vipType"] = "constant"
	ipMap["vipValue"] = interfaceToStrList(input["ip"])
	ipMap["vipVariableName"] = "vpn_host_ip"
	defMap["ip"] = ipMap

	defMap["priority-order"] = []string{"hostname", "ip"}
}

func createVPNIPv4Routes(defMap map[string]interface{}, input map[string]interface{}) error {

	nexthopList := input["next_hop"].(*schema.Set).List()

	if input["gateway_type"] == "dhcp" {
		if input["null0_distance"].(string) != "" {
			return fmt.Errorf("null0_distance should be set for gateway_type null0")
		} else if len(nexthopList) > 0 {
			return fmt.Errorf("next_hop should be set for gateway_type next-hop")
		}

		dhcpMap := make(map[string]interface{})
		dhcpMap["vipType"] = "constant"
		dhcpMap["vipObjectType"] = "node-only"
		dhcpMap["vipValue"] = "true"

		defMap["dhcp"] = dhcpMap

		defMap["priority-order"] = []string{"prefix", "dhcp"}

	} else if input["gateway_type"] == "null0" {
		if len(nexthopList) > 0 {
			return fmt.Errorf("next_hop should be set for gateway_type next_hop")
		}
		null0Map := make(map[string]interface{})
		null0Map["vipType"] = "constant"
		null0Map["vipObjectType"] = "node-only"
		null0Map["vipValue"] = "true"

		defMap["null0"] = null0Map

		if input["null0_distance"].(string) != "" {
			nullDis := make(map[string]interface{})
			nullDis["vipType"] = "constant"
			nullDis["vipObjectType"] = "object"
			nullDis["vipValue"], _ = getInt(input["null0_distance"])
			defMap["distance"] = nullDis
		}

	} else if input["gateway_type"] == "next-hop" {
		if input["null0_distance"].(string) != "" {
			return fmt.Errorf("null0_distance should be set for gateway_type null0")
		}
		if len(nexthopList) > 0 {
			nexthopVal := make([]map[string]interface{}, 0, len(nexthopList))

			for _, nextHopMap := range nexthopList {
				nextHop := make(map[string]interface{})
				createNextHop(nextHop, nextHopMap.(map[string]interface{}))
				nexthopVal = append(nexthopVal, nextHop)
			}

			nextHop := make(map[string]interface{})
			nextHop["vipType"] = "constant"
			nextHop["vipValue"] = nexthopVal
			nextHop["vipObjectType"] = "tree"
			nextHop["vipPrimaryKey"] = []string{"address"}
			defMap["priority-order"] = []string{"prefix", "next-hop"}
			defMap["next-hop"] = nextHop
		} else {
			return fmt.Errorf("there should be atleast one next_hop")
		}

	}

	if input["prefix"] != nil {
		prefixMap := make(map[string]interface{})
		prefixMap["vipObjectType"] = "object"
		prefixMap["vipType"] = "constant"
		prefixMap["vipValue"] = input["prefix"]
		defMap["prefix"] = prefixMap
	}

	return nil
}

func createNextHop(defMap map[string]interface{}, input map[string]interface{}) {

	if input["next_hop_address"] != nil {
		address := make(map[string]interface{})
		address["vipObjectType"] = "object"
		address["vipType"] = "constant"
		address["vipValue"] = input["next_hop_address"]
		defMap["address"] = address
	}

	if input["next_hop_distance"] != nil && input["next_hop_distance"] != "" {
		distance := make(map[string]interface{})
		distance["vipObjectType"] = "object"
		distance["vipType"] = "constant"
		distance["vipValue"], _ = getInt(input["next_hop_distance"])
		defMap["distance"] = distance
	}

	defMap["priority-order"] = []string{"address", "distance"}
}

func createVPNIPv6Routes(defMap map[string]interface{}, input map[string]interface{}, ftType string) error {

	nexthopList := input["next_hop"].(*schema.Set).List()

	if input["gateway_type"] == "null0" {
		if len(nexthopList) > 0 {
			return fmt.Errorf("next_hop should be set for gateway_type next_hop")
		}

		if ftType == "cisco_vpn" && input["null0_distance"] != "" {
			return fmt.Errorf("null0_distance should not be set for feature template_type cisco_vpn")
		}
		null0Map := make(map[string]interface{})
		null0Map["vipType"] = "constant"
		null0Map["vipObjectType"] = "node-only"
		null0Map["vipValue"] = "true"

		defMap["null0"] = null0Map

		if input["null0_distance"] != nil && input["null0_distance"] != "" {
			nullDis := make(map[string]interface{})
			nullDis["vipObjectType"] = "object"
			nullDis["vipType"] = "constant"
			nullDis["vipValue"], _ = getInt(input["null0_distance"])
			defMap["distance"] = nullDis
		}

	} else if input["gateway_type"] == "next-hop" {
		if input["null0_distance"] != "" {
			return fmt.Errorf("null0_distance should be set for gateway_type null0")
		}
		if len(nexthopList) > 0 {
			nexthopVal := make([]map[string]interface{}, 0, 1)

			for _, nextHopMap := range nexthopList {
				nextHop := make(map[string]interface{})
				createNextHop(nextHop, nextHopMap.(map[string]interface{}))
				nexthopVal = append(nexthopVal, nextHop)
			}

			nextHop := make(map[string]interface{})
			nextHop["vipType"] = "constant"
			nextHop["vipValue"] = nexthopVal
			nextHop["vipObjectType"] = "tree"
			nextHop["vipPrimaryKey"] = []string{"address"}
			defMap["priority-order"] = []string{"prefix", "next-hop"}
			defMap["next-hop"] = nextHop
		} else {
			return fmt.Errorf("there should be atleast one next_hop")
		}

	}

	if input["prefix"] != nil {
		prefixMap := make(map[string]interface{})
		prefixMap["vipObjectType"] = "object"
		prefixMap["vipType"] = "constant"
		prefixMap["vipValue"] = input["prefix"]
		prefixMap["vipVariableName"] = "vpn_ipv6_ipv6_prefix"
		defMap["prefix"] = prefixMap
	}
	return nil
}

func enableTEService(defMap map[string]interface{}) {
	outerMap := make(map[string]interface{})

	outerMap["vipType"] = "constant"
	outerMap["vipObjectType"] = "tree"
	outerMap["vipPrimaryKey"] = []string{"svc-type"}
	result := make([]map[string]interface{}, 0, 1)
	valueMap := make(map[string]interface{})
	getServiceValueMap(valueMap)
	result = append(result, valueMap)
	outerMap["vipValue"] = result
	defMap["service"] = outerMap
}

func getServiceValueMap(defMap map[string]interface{}) {
	svcMap := make(map[string]interface{})

	svcMap["vipObjectType"] = "object"
	svcMap["vipType"] = "constant"
	svcMap["vipValue"] = "TE"

	defMap["svc-type"] = svcMap
	defMap["priority-order"] = []string{"svc-type"}
}

func createServiceRoutes(defMap map[string]interface{}, input map[string]interface{}, ftType string) error {

	if input["prefix"] != nil {
		prefix := make(map[string]interface{})
		prefix["vipObjectType"] = "object"
		prefix["vipType"] = "constant"
		prefix["vipValue"] = input["prefix"]
		prefix["vipVariableName"] = "vpn_serviceroute_prefix"
		defMap["prefix"] = prefix
	}

	if input["service_type"] != "" {
		service := make(map[string]interface{})
		service["vipObjectType"] = "object"
		service["vipType"] = "constant"
		if input["service_type"] == "sig" && ftType == "cisco_vpn" {
			return fmt.Errorf("sig is not allowed service type for cisco_vpn")
		}
		service["vipValue"] = input["service_type"]
		defMap["service"] = service
	}

	defMap["priority-order"] = []string{"prefix", "service"}

	return nil
}

func createNATs(defMap map[string]interface{}, input map[string]interface{}, ftType string) error {

	if ftType == "cisco_vpn" {
		return fmt.Errorf("nat64 cannot be configured for temlate_type cisco_vpn")
	}

	if input["pool_name"] != nil {
		name := make(map[string]interface{})
		name["vipObjectType"] = "object"
		name["vipType"] = "constant"
		name["vipValue"] = input["pool_name"]
		defMap["name"] = name
	}

	if input["start_address"] != nil {
		saddress := make(map[string]interface{})
		saddress["vipObjectType"] = "object"
		saddress["vipType"] = "constant"
		saddress["vipValue"] = input["start_address"]
		saddress["vipVariableName"] = "vpn_interface_start_address"
		defMap["start-address"] = saddress
	}

	if input["end_address"] != nil {
		eaddress := make(map[string]interface{})
		eaddress["vipObjectType"] = "object"
		eaddress["vipType"] = "constant"
		eaddress["vipValue"] = input["end_address"]
		eaddress["vipVariableName"] = "vpn_interface_end_address"
		defMap["end-address"] = eaddress
	}

	if input["overload"] != nil {
		overload := make(map[string]interface{})
		overload["vipObjectType"] = "node-only"
		overload["vipType"] = "constant"
		overload["vipValue"] = strconv.FormatBool(input["overload"].(bool))
		overload["vipVariableName"] = "vpn_interface_overload"
		defMap["overload"] = overload
	}

	defMap["priority-order"] = []string{"name", "start-address", "end-address", "overload"}

	return nil
}

func verifyDeviceTypeVPN(templateType string, deviceType []string) bool {
	var found = 0
	if templateType == "vpn-vedge" {
		for _, dType := range deviceType {
			if vpnDeviceTypes[dType] {
				found++
			}
		}
	} else {
		for _, dType := range deviceType {
			if ciscoVPNDeviceTypes[dType] {
				found++
			}
		}
	}

	return found == len(deviceType)
}

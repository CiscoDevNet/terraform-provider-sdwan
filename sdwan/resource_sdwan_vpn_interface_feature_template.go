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
	IPv6                   = make(map[string]interface{})
	isNATPresent           int
	isTLOCPresent          int
	template_name          string
	vedgeMaxControlConnect = []int{0, 8}
	ciscoMaxControlConnect = []int{0, 100}
	vedgeTimer             = []int{1, 3600}
	ciscoTimer             = []int{100, 40950}
	vedgeTimerDefault      = 1
	ciscoTimerDefault      = 100
	vedgeTCPMSS            = []int{552, 1460}
	ciscoTCPMSS            = []int{500, 1460}
	vedgeARPTimeOut        = []int{0, 2678400}
	ciscoARPTimeOut        = []int{0, 2147483}
)

func resourceSDWANVPNInterfaceFeatureTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSDWANVPNInterfaceFeatureTemplateCreate,
		Read:   resourceSDWANVPNInterfaceFeatureTemplateRead,
		Update: resourceSDWANVPNInterfaceFeatureTemplateUpdate,
		Delete: resourceSDWANVPNInterfaceFeatureTemplateDelete,

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
				ValidateFunc: DescValidation(),
			},
			"device_type": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"vedge-100",
						"vedge-100-WM",
						"vedge-ISR1100-4GLTE",
						"vedge-100-B",
						"vedge-1000",
						"vedge-ISR1100-6G",
						"vedge-100-M",
						"vedge-2000",
						"vedge-ISR1100X-4G",
						"vedge-ISR1100-4G",
						"vedge-5000",
						"vedge-ISR1100X-6G",
						"vedge-cloud",
						"vedge-C8500-12X4QC",
						"vedge-ISRv",
						"vedge-C1111-4PLTEEA",
						"vedge-C1161-8P",
						"vedge-C1117-4PLTEEAW",
						"vedge-C1121X-8P",
						"vedge-C8200-1N-4T",
						"vedge-ISR-4331",
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
						"vedge-C1117-4PLTELAWZ",
					}, false),
				},
			},
			"template_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"vpn-vedge-interface",
					"cisco_vpn_interface",
				}, false),
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
						"vpn_interface_basic": {
							Type:     schema.TypeSet,
							Required: true,
							MinItems: 1,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"shutdown": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ipv4": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"primary_address": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"secondary_address": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													MaxItems: 4,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"address": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},
												"dhcp_distance": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ValidateDiagFunc: isStringInRange(1, 65536),
												},
												"dhcp_helper": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type:         schema.TypeString,
														ValidateFunc: DHCPHelperValidation(),
													},
												},
											},
										},
									},
									"ipv6": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"primary_address": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"secondary_address": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													MaxItems: 2,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"address": {
																Type:     schema.TypeString,
																Required: true,
															},
														},
													},
												},
												"dhcp_distance": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ValidateDiagFunc: isStringInRange(1, 65536),
												},
												"dhcp_rapid_commit": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"dhcp_helper": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													MaxItems: 8,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"address": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validation.IsIPv6Address,
															},
															"vpn": {
																Type:         schema.TypeInt,
																Optional:     true,
																Computed:     true,
																ValidateFunc: validation.IntBetween(1, 65536),
															},
														},
													},
												},
											},
										},
									},
									"block_non_source_ip": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"bandwidth_upstream": {
										Type:         schema.TypeFloat,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.FloatBetween(1, 2147483647),
									},
									"bandwidth_downstream": {
										Type:         schema.TypeFloat,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.FloatBetween(1, 2147483647),
									},
								},
							},
						},

						"vpn_interface_tunnel": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"per_tunnel_qos": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"per_tunnel_qos_aggregator": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"tunnels_bandwidth_percent": {
										Type:         schema.TypeInt,
										Optional:     true,
										ValidateFunc: validation.IntBetween(1, 99),
									},
									"color": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ValidateFunc: validation.StringInSlice([]string{
											"3g",
											"biz-internet",
											"blue",
											"bronze",
											"custom1",
											"custom2",
											"custom3",
											"default",
											"gold",
											"green",
											"Ite",
											"metro-ethernet",
											"mpls",
											"private1",
											"private2",
											"private3",
											"private4",
											"private5",
											"private6",
											"public-internet",
											"red",
											"silver",
										}, false),
									},
									"restrict": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"groups": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type:         schema.TypeFloat,
											ValidateFunc: validation.FloatBetween(1, 4294967295),
										},
									},
									"border": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"control_connection": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"maximum_control_connections": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ValidateDiagFunc: isStringInRange(0, 100),
									},
									"vbond_as_stun_server": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"exclude_controller_group_list": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type:         schema.TypeInt,
											ValidateFunc: validation.IntBetween(0, 100),
										},
									},
									"vmanage_connection_preference": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      5,
										ValidateFunc: validation.IntBetween(0, 8),
									},
									"port_hop": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},
									"low_bandwidth_link": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"allow_service": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"all": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"bgp": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"dhcp": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
												},
												"dns": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
												},
												"icmp": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
												},
												"netconf": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"ntp": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"ospf": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"ssh": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"stun": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"https": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
												},
												"snmp": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
											},
										},
									},
									"encapsulation": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"gre": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"gre_preference": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ValidateDiagFunc: isStringInRange(0, 4294967295),
												},
												"gre_weight": {
													Type:         schema.TypeInt,
													Optional:     true,
													Default:      1,
													ValidateFunc: validation.IntBetween(1, 255),
												},
												"ipsec": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
												},
												"ipsec_preference": {
													Type:             schema.TypeString,
													Optional:         true,
													Computed:         true,
													ValidateDiagFunc: isStringInRange(0, 4294967295),
												},
												"ipsec_weight": {
													Type:         schema.TypeInt,
													Optional:     true,
													Default:      1,
													ValidateFunc: validation.IntBetween(1, 255),
												},
												"carrier": {
													Type:     schema.TypeString,
													Optional: true,
													Default:  "default",
													ValidateFunc: validation.StringInSlice([]string{
														"carrier1",
														"carrier2",
														"carrier3",
														"carrier4",
														"carrier5",
														"carrier6",
														"carrier7",
														"carrier8",
														"default",
													}, false),
												},
												"bind_loopback_tunnel": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"last_resort_circuit": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"hold_time": {
													Type:         schema.TypeInt,
													Optional:     true,
													Default:      7000,
													ValidateFunc: validation.IntBetween(100, 10000),
												},
												"nat_refresh_interval": {
													Type:         schema.TypeInt,
													Optional:     true,
													Default:      5,
													ValidateFunc: validation.IntBetween(1, 60),
												},
												"hello_interval": {
													Type:         schema.TypeInt,
													Optional:     true,
													Default:      1000,
													ValidateFunc: validation.IntBetween(100, 600000),
												},
												"hello_tolerance": {
													Type:         schema.TypeInt,
													Optional:     true,
													Default:      12,
													ValidateFunc: validation.IntBetween(12, 6000),
												},
												"gre_tunnel_destination_ip": {
													Type:         schema.TypeString,
													Optional:     true,
													Computed:     true,
													ValidateFunc: validation.IsIPv4Address,
												},
											},
										},
									},
								},
							},
						},

						"vpn_interface_nat": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nat_type": {
													Type:     schema.TypeString,
													Optional: true,
													Default:  "interface",
													ValidateFunc: validation.StringInSlice([]string{
														"interface",
														"pool",
														"loopback",
													}, false),
												},
												"refresh_mode": {
													Type:     schema.TypeString,
													Optional: true,
													Default:  "outbound",
													ValidateFunc: validation.StringInSlice([]string{
														"bi-directional",
														"outbound",
													}, false),
												},
												"log_nat_flow_creations_or_deletions": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"udp_timeout": {
													Type:         schema.TypeInt,
													Optional:     true,
													Default:      1,
													ValidateFunc: validation.IntBetween(1, 65536),
												},
												"tcp_timeout": {
													Type:         schema.TypeInt,
													Optional:     true,
													Default:      60,
													ValidateFunc: validation.IntBetween(1, 65536),
												},
												"block_icmp": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
												},
												"respond_to_ping": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"nat_pool_range_start": {
													Type:         schema.TypeString,
													Optional:     true,
													Computed:     true,
													ValidateFunc: validation.IsIPv4Address,
												},
												"nat_pool_range_end": {
													Type:         schema.TypeString,
													Optional:     true,
													Computed:     true,
													ValidateFunc: validation.IsIPv4Address,
												},
												"nat_pool_prefix_length": {
													Type:         schema.TypeInt,
													Optional:     true,
													Default:      60,
													ValidateFunc: validation.IntBetween(1, 65536),
												},
												"overload": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
												},
												"nat_inside_source_loopback_interface": {
													Type:         schema.TypeString,
													Optional:     true,
													Computed:     true,
													ValidateFunc: NATLoopBackInterfaceValidation(),
												},
												"port_forward": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"port_start_range": {
																Type:         schema.TypeInt,
																Required:     true,
																ValidateFunc: validation.IntBetween(0, 65535),
															},
															"port_end_range": {
																Type:         schema.TypeInt,
																Required:     true,
																ValidateFunc: validation.IntBetween(0, 65535),
															},
															"protocol": {
																Type:     schema.TypeString,
																Required: true,
																ValidateFunc: validation.StringInSlice([]string{
																	"tcp",
																	"udp",
																}, false),
															},
															"vpn": {
																Type:         schema.TypeInt,
																Required:     true,
																ValidateFunc: validation.IntBetween(0, 65535),
															},
															"private_ip": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validation.IsIPv4Address,
															},
														},
													},
												},
												"static_nat": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"source_ip": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validation.IsIPv4Address,
															},
															"translated_source_ip": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validation.IsIPv4Address,
															},
															"source_vpn_id": {
																Type:         schema.TypeInt,
																Optional:     true,
																Default:      0,
																ValidateFunc: validation.IntBetween(0, 65530),
															},
															"static_nat_direction": {
																Type:     schema.TypeString,
																Optional: true,
																Default:  "inside",
																ValidateFunc: validation.StringInSlice([]string{
																	"inside",
																	"outside",
																}, false),
															},
															"protocol": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ValidateFunc: validation.StringInSlice([]string{
																	"tcp",
																	"udp",
																}, false),
															},
															"source_port": {
																Type:             schema.TypeString,
																Optional:         true,
																Computed:         true,
																ValidateDiagFunc: isStringInRange(0, 65535),
															},
															"translate_port": {
																Type:             schema.TypeString,
																Optional:         true,
																Computed:         true,
																ValidateDiagFunc: isStringInRange(0, 65535),
															},
														},
													},
												},
											},
										},
									},
									"ipv6": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nat64": {
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

						"vpn_interface_vrrp": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										MaxItems: 5,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"group_id": {
													Type:         schema.TypeInt,
													Required:     true,
													ValidateFunc: validation.IntBetween(1, 255),
												},
												"priority": {
													Type:         schema.TypeInt,
													Optional:     true,
													Default:      100,
													ValidateFunc: validation.IntBetween(1, 254),
												},
												"timer": {
													Type:         schema.TypeInt,
													Optional:     true,
													Computed:     true,
													ValidateFunc: validation.IntBetween(1, 40950),
												},
												"track_omp": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"track_prefix_list": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"ip_address": {
													Type:         schema.TypeString,
													Optional:     true,
													Computed:     true,
													ValidateFunc: validation.IsIPv4Address,
												},
											},
										},
									},
									"ipv6": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"group_id": {
													Type:         schema.TypeInt,
													Required:     true,
													ValidateFunc: validation.IntBetween(1, 255),
												},
												"priority": {
													Type:         schema.TypeInt,
													Optional:     true,
													Default:      100,
													ValidateFunc: validation.IntBetween(1, 254),
												},
												"timer": {
													Type:         schema.TypeInt,
													Optional:     true,
													Computed:     true,
													ValidateFunc: validation.IntBetween(1, 40950),
												},
												"track_omp": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  false,
												},
												"track_prefix_list": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"link_local_ipv6_address": {
													Type:         schema.TypeString,
													Required:     true,
													ValidateFunc: validation.IsIPv6Address,
												},
												"global_ipv6_prefix": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"vpn_interface_acl_qos": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"adapt_period": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      15,
										ValidateFunc: validation.IntBetween(1, 720),
									},
									"shaping_rate_upstream_min": {
										Type:         schema.TypeFloat,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.FloatBetween(8, 100000000),
									},
									"shaping_rate_upstream_max": {
										Type:         schema.TypeFloat,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.FloatBetween(8, 100000000),
									},
									"shaping_rate_upstream_default": {
										Type:         schema.TypeFloat,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.FloatBetween(8, 100000000),
									},
									"shaping_rate_downstream_min": {
										Type:         schema.TypeFloat,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.FloatBetween(8, 100000000),
									},
									"shaping_rate_downstream_max": {
										Type:         schema.TypeFloat,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.FloatBetween(8, 100000000),
									},
									"shaping_rate_downstream_default": {
										Type:         schema.TypeFloat,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.FloatBetween(8, 100000000),
									},
									"shaping_rate": {
										Type:         schema.TypeFloat,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.FloatBetween(8, 100000000),
									},
									"qos_map": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rewrite_rule": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ipv4_ingress_access_list": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ipv4_egress_access_list": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ipv6_ingress_access_list": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ipv6_egress_access_list": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ingress_policer_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"egress_policer_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},

						"vpn_interface_arp": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.IsIPv4Address,
									},
									"mac_address": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.IsMACAddress,
									},
								},
							},
						},

						"vpn_interface_8021x": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"radius_server": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"account_interim_interval": {
										Type:         schema.TypeInt,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.IntBetween(1, 1440),
									},
									"nas_identifier": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"nas_ip": {
										Type:         schema.TypeString,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.IsIPv4Address,
									},
									"wake_on_lan": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"control_direction": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "in-and-out",
										ValidateFunc: validation.StringInSlice([]string{
											"in-and-out",
											"in-only",
										}, false),
									},
									"reauthentication": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      0,
										ValidateFunc: validation.IntBetween(0, 1440),
									},
									"inactivity": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      60,
										ValidateFunc: validation.IntBetween(0, 1440),
									},
									"host_mode": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "single-host",
										ValidateFunc: validation.StringInSlice([]string{
											"multi-auth",
											"multi-host",
											"single-host",
										}, false),
									},
									"advanced_options": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vlan": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"authentication_fail_vlan": {
																Type:         schema.TypeInt,
																Optional:     true,
																Computed:     true,
																ValidateFunc: validation.IntBetween(1, 65536),
															},
															"guest_vlan": {
																Type:         schema.TypeInt,
																Optional:     true,
																Computed:     true,
																ValidateFunc: validation.IntBetween(1, 65536),
															},
															"authentication_reject_vlan": {
																Type:         schema.TypeInt,
																Optional:     true,
																Computed:     true,
																ValidateFunc: validation.IntBetween(1, 65536),
															},
															"default_vlan": {
																Type:         schema.TypeInt,
																Optional:     true,
																Computed:     true,
																ValidateFunc: validation.IntBetween(1, 65536),
															},
														},
													},
												},
												"dynamic_authentication_server": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"das_port": {
																Type:         schema.TypeInt,
																Optional:     true,
																Default:      3799,
																ValidateFunc: validation.IntBetween(1, 65535),
															},
															"client": {
																Type:         schema.TypeString,
																Optional:     true,
																Computed:     true,
																ValidateFunc: validation.IsIPv4Address,
															},
															"secret_key": {
																Type:      schema.TypeString,
																Optional:  true,
																Computed:  true,
																Sensitive: true,
															},
															"time_window": {
																Type:         schema.TypeInt,
																Optional:     true,
																Default:      300,
																ValidateFunc: validation.IntBetween(10, 1000),
															},
															"required_timestamp": {
																Type:     schema.TypeBool,
																Optional: true,
																Default:  false,
															},
															"vpn": {
																Type:         schema.TypeInt,
																Optional:     true,
																Default:      0,
																ValidateFunc: validation.IntBetween(0, 65535),
															},
														},
													},
												},
												"mac_authentication_bypass": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"server": {
																Type:     schema.TypeBool,
																Optional: true,
																Default:  false,
															},
															"mac_authentication_bypass_entries": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 8,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"request_attributes": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"authentication": {
																Type:     schema.TypeSet,
																Optional: true,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": {
																			Type:         schema.TypeInt,
																			Required:     true,
																			ValidateFunc: validation.IntBetween(1, 64),
																		},
																		"syntax_choice": {
																			Type:     schema.TypeString,
																			Required: true,
																			ValidateFunc: validation.StringInSlice([]string{
																				"string",
																				"integer",
																				"octet",
																			}, false),
																		},
																		"value": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},
															"accounting": {
																Type:     schema.TypeSet,
																Optional: true,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": {
																			Type:         schema.TypeInt,
																			Required:     true,
																			ValidateFunc: validation.IntBetween(1, 64),
																		},
																		"syntax_choice": {
																			Type:     schema.TypeString,
																			Required: true,
																			ValidateFunc: validation.StringInSlice([]string{
																				"string",
																				"integer",
																				"octet",
																			}, false),
																		},
																		"value": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},

						"vpn_interface_trustsec": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_sgt_propagation": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"propagate": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"security_group_tag": {
										Type:         schema.TypeInt,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.IntBetween(2, 65519),
									},
									"trusted": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
								},
							},
						},

						"vpn_interface_advanced": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"duplex": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ValidateFunc: validation.StringInSlice([]string{
											"auto",
											"full",
											"half",
										}, false),
									},
									"mac_address": {
										Type:         schema.TypeString,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.IsMACAddress,
									},
									"ip_mtu": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      1500,
										ValidateFunc: validation.IntBetween(576, 2000),
									},
									"pmtu_discovery": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"flow_control": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "autoneg",
										ValidateFunc: validation.StringInSlice([]string{
											"autoneg",
											"both",
											"egress",
											"ingress",
											"none",
										}, false),
									},
									"tcp_mss": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ValidateDiagFunc: isStringInRange(500, 1460),
									},
									"speed": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ValidateFunc: validation.StringInSlice([]string{
											"10",
											"100",
											"1000",
											"10000",
											"2500",
										}, false),
									},
									"clear_dont_fragment": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"static_ingess_qos": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ValidateDiagFunc: isStringInRange(0, 7),
									},
									"arp_timeout": {
										Type:         schema.TypeInt,
										Optional:     true,
										Default:      1200,
										ValidateFunc: validation.IntBetween(0, 2678400),
									},
									"autonegotiation": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},
									"tloc_extension": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"power_over_ethernet": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"load_interval": {
										Type:         schema.TypeInt,
										Optional:     true,
										Computed:     true,
										ValidateFunc: LoadIntervalValidation(),
									},
									"tracker": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"icmp_redirect_disable": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
									"gre_tunnel_source_ip": {
										Type:         schema.TypeString,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.IsIPv4Address,
									},
									"xconnect": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"ip_directed_broadcast": {
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

func resourceSDWANVPNInterfaceFeatureTemplateCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Create method!")

	sdwanClient := m.(*client.Client)

	ftMap := make(map[string]interface{})

	ftMap["templateType"] = d.Get("template_type").(string)
	ftMap["templateName"] = d.Get("template_name").(string)
	template_name = d.Get("template_name").(string)
	ftMap["templateDescription"] = d.Get("template_description").(string)
	device_type := interfaceToStrList(d.Get("device_type"))
	check := validateVPNInterfaceDeviceType(device_type, ftMap["templateType"].(string))
	if !check {
		return fmt.Errorf("[ERROR] Device Type is not compatible with Template type!")
	}

	ftMap["deviceType"] = interfaceToStrList(d.Get("device_type"))
	ftMap["templateMinVersion"] = d.Get("template_min_version").(string)
	ftMap["factoryDefault"] = d.Get("factory_default").(bool)

	ftDefinition := d.Get("template_definition").(*schema.Set).List()
	if def, err := createVPNInterfaceFTDefinition(ftDefinition, ftMap["templateType"].(string)); err == nil {
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

	log.Println("[DEBUG] End of Create Method ", d.Id())

	return resourceSDWANVPNInterfaceFeatureTemplateRead(d, m)
}

func resourceSDWANVPNInterfaceFeatureTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Update Method ", d.Id())

	sdwanClient := m.(*client.Client)

	ftID := d.Id()

	ftMap := make(map[string]interface{})

	ftMap["templateType"] = d.Get("template_type").(string)
	ftMap["templateName"] = d.Get("template_name").(string)
	template_name = d.Get("template_name").(string)
	ftMap["templateDescription"] = d.Get("template_description").(string)
	device_type := interfaceToStrList(d.Get("device_type"))
	check := validateVPNInterfaceDeviceType(device_type, ftMap["templateType"].(string))
	if !check {
		return fmt.Errorf("[ERROR] Device Type is not compatible with Template type!")
	}

	ftMap["deviceType"] = interfaceToStrList(d.Get("device_type"))
	ftMap["templateMinVersion"] = d.Get("template_min_version").(string)
	ftMap["factoryDefault"] = d.Get("factory_default").(bool)

	ftDefinition := d.Get("template_definition").(*schema.Set).List()
	if def, err := createVPNInterfaceFTDefinition(ftDefinition, ftMap["templateType"].(string)); err == nil {
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

	return resourceSDWANVPNInterfaceFeatureTemplateRead(d, m)
}

func resourceSDWANVPNInterfaceFeatureTemplateRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Begining Read method ", d.Id())

	sdwanClient := m.(*client.Client)

	ftid := d.Id()

	cont, err := getFeatureTemplateByID(sdwanClient, ftid)

	if err != nil {
		if cont != nil {
			return fmt.Errorf("[ERROR] Data load failed in between!")
		} else {
			return err
		}
	}

	setVPNInterfaceTemplateAttributes(d, cont)

	log.Println("[DEBUG] End of Read Method ", d.Id())
	return nil
}

func resourceSDWANVPNInterfaceFeatureTemplateDelete(d *schema.ResourceData, m interface{}) error {
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

func createVPNInterfaceFTDefinition(ftDefinitions []interface{}, ftType string) (map[string]interface{}, error) {
	definition := make(map[string]interface{})

	ftDefinition := ftDefinitions[0]
	ftDefMap := ftDefinition.(map[string]interface{})

	vpnInterfaceBasicMap := (ftDefMap["vpn_interface_basic"].(*schema.Set).List())[0].(map[string]interface{})
	createVPNInterfaceBasic(definition, vpnInterfaceBasicMap)

	if len(ftDefMap["vpn_interface_tunnel"].(*schema.Set).List()) > 0 {
		vpnInterfaceTunnelMap := (ftDefMap["vpn_interface_tunnel"].(*schema.Set).List())[0].(map[string]interface{})
		err := createVPNInterfaceTunnel(definition, vpnInterfaceTunnelMap, ftType)
		if err != nil {
			return nil, err
		}
	}

	natSet := ftDefMap["vpn_interface_nat"].(*schema.Set).List()
	if len(natSet) > 0 {
		vpnInterfaceNATMap := (ftDefMap["vpn_interface_nat"].(*schema.Set).List())[0].(map[string]interface{})
		err := createVPNInterfaceNAT(definition, vpnInterfaceNATMap, ftType)
		if err != nil {
			return nil, err
		}
	}

	if len(ftDefMap["vpn_interface_vrrp"].(*schema.Set).List()) > 0 {
		vpnInterfaceVRRPMap := (ftDefMap["vpn_interface_vrrp"].(*schema.Set).List())[0].(map[string]interface{})
		err := createVPNInterfaceVRRP(definition, vpnInterfaceVRRPMap, ftType)
		if err != nil {
			return nil, err
		}
	}

	if len(ftDefMap["vpn_interface_acl_qos"].(*schema.Set).List()) > 0 {
		vpnInterfaceACLMap := (ftDefMap["vpn_interface_acl_qos"].(*schema.Set).List())[0].(map[string]interface{})
		createVPNInterfaceACL(definition, vpnInterfaceACLMap)
	} else {
		temp := make([]interface{}, 0, 1)
		ipv4AccessListMap := make(map[string]interface{})
		ipv4AccessListMap["vipType"] = "ignore"
		ipv4AccessListMap["vipObjectType"] = "tree"
		ipv4AccessListMap["vipValue"] = temp
		ipv4AccessListMap["vipPrimaryKey"] = []string{
			"direction",
		}

		definition["access-list"] = ipv4AccessListMap

		// vManage v20.6.2.2 (possibly others) breaks badly if any interface templates exist without qos-map
		definition["qos-map"] = createVIPObject("object", "ignore", nil, "qos_map", nil)

		// vManage v20.6.2.2 (possibly others) breaks badly if any interface templates exist without rewrite-rule
		rewriteRule := make(map[string]interface{})
		rewriteRule["rule-name"] = createVIPObject("object", "ignore", nil, "rewrite_rule_name", nil)
		definition["rewrite-rule"] = rewriteRule
	}

	arpLen := len(ftDefMap["vpn_interface_arp"].(*schema.Set).List())
	if arpLen > 0 {

		arpSet := ftDefMap["vpn_interface_arp"].(*schema.Set).List()

		arps := make([]interface{}, 0, arpLen)

		for _, arpMap := range arpSet {
			arp := make(map[string]interface{})
			createVPNInterfaceARP(arp, arpMap.(map[string]interface{}))
			arps = append(arps, arp)
		}

		IPMap := make(map[string]interface{})
		IPMap["vipType"] = "constant"
		IPMap["vipObjectType"] = "tree"
		IPMap["vipValue"] = arps
		IPMap["vipPrimaryKey"] = []string{
			"addr",
		}

		ARPMap := make(map[string]interface{})
		ARPMap["ip"] = IPMap

		definition["arp"] = ARPMap
	} else {
		arps := make([]interface{}, 0, 1)
		IPMap := make(map[string]interface{})
		IPMap["vipType"] = "ignore"
		IPMap["vipObjectType"] = "tree"
		IPMap["vipValue"] = arps
		IPMap["vipPrimaryKey"] = []string{
			"addr",
		}

		ARPMap := make(map[string]interface{})
		ARPMap["ip"] = IPMap

		definition["arp"] = ARPMap
	}

	if ftType == "cisco_vpn_interface" {
		if len(ftDefMap["vpn_interface_trustsec"].(*schema.Set).List()) > 0 {
			vpnInterfaceTrustSecMap := (ftDefMap["vpn_interface_trustsec"].(*schema.Set).List())[0].(map[string]interface{})
			createVPNInterfaceTrustSec(definition, vpnInterfaceTrustSecMap)
		}
	}

	if len(ftDefMap["vpn_interface_advanced"].(*schema.Set).List()) > 0 {
		vpnInterfaceAdvancedMap := (ftDefMap["vpn_interface_advanced"].(*schema.Set).List())[0].(map[string]interface{})
		err := createVPNInterfaceAdvanced(definition, vpnInterfaceAdvancedMap, ftType)
		if err != nil {
			return nil, err
		}
	}

	if len(ftDefMap["vpn_interface_8021x"].(*schema.Set).List()) > 0 {
		vpnInterface8021XMap := (ftDefMap["vpn_interface_8021x"].(*schema.Set).List())[0].(map[string]interface{})
		err := createVPNInterface8021X(definition, vpnInterface8021XMap)
		if err != nil {
			return nil, err
		}
	} else {
		dot1xMap := make(map[string]interface{})
		dot1xMap["vipType"] = "ignore"
		dot1xMap["vipObjectType"] = "node-only"
		definition["dot1x"] = dot1xMap
	}

	return definition, nil
}

func createVPNInterfaceBasic(defMap map[string]interface{}, input map[string]interface{}) {

	if input["shutdown"] != nil {
		defMap["shutdown"] = createVIPObject("object", "constant", input["shutdown"], "vpn_if_shutdown", nil)
	}

	template_name = "vpn_if_name_" + template_name
	defMap["if-name"] = createVIPObject("object", "variableName", "", template_name, nil)

	if input["description"] != nil && input["description"] != "" {
		defMap["description"] = createVIPObject("object", "constant", input["description"], "vpn_if_description", nil)
	} else {
		defMap["description"] = createVIPObject("object", "ignore", nil, "vpn_if_description", nil)
	}

	//	ipv4 configuration

	if input["ipv4"] != nil && len(input["ipv4"].(*schema.Set).List()) > 0 {
		vpnInterfaceBasicIPv4Map := (input["ipv4"].(*schema.Set).List())[0].(map[string]interface{})
		createVPNInterfaceBasicIPv4(defMap, vpnInterfaceBasicIPv4Map)
	}

	//ipv6 configuration

	if input["ipv6"] != nil && len(input["ipv6"].(*schema.Set).List()) > 0 {
		vpnInterfaceBasicIPv6Map := (input["ipv6"].(*schema.Set).List())[0].(map[string]interface{})
		createVPNInterfaceBasicIPv6(defMap, vpnInterfaceBasicIPv6Map)
	}

	if input["block_non_source_ip"] != nil {
		defMap["block-non-source-ip"] = createVIPObject("object", "constant", input["block_non_source_ip"], "vpn_if_block_non_source_ip", nil)
	}

	if input["bandwidth_upstream"] != nil && input["bandwidth_upstream"] != 0.0 {
		defMap["bandwidth-upstream"] = createVIPObject("object", "constant", input["bandwidth_upstream"], "vpn_if_bandwidth_upstream", nil)
	} else {
		defMap["bandwidth-upstream"] = createVIPObject("object", "ignore", nil, "vpn_if_bandwidth_upstream", nil)
	}

	if input["bandwidth_downstream"] != nil && input["bandwidth_downstream"] != 0.0 {
		defMap["bandwidth-downstream"] = createVIPObject("object", "constant", input["bandwidth_downstream"], "vpn_if_bandwidth_downstream", nil)
	} else {
		defMap["bandwidth-downstream"] = createVIPObject("object", "ignore", nil, "vpn_if_bandwidth_downstream", nil)
	}
}

func createVPNInterfaceBasicIPv4(defMap map[string]interface{}, input map[string]interface{}) {

	ipv4 := make(map[string]interface{})
	isPresent := 0
	if input["primary_address"] != nil && input["primary_address"] != "" {
		ipv4["address"] = createVIPObject("object", "constant", input["primary_address"], "vpn_if_ipv4_address", nil)
		isPresent = isPresent + 1
	}

	//secondary_address configuration
	secAddrSet := input["secondary_address"].(*schema.Set).List()

	if len(secAddrSet) > 0 {
		secAddrs := []interface{}{}
		SecAddrMap := make(map[string]interface{})
		count := 0
		for _, secAddrMap := range secAddrSet {
			secAddr := make(map[string]interface{})
			createVPNInterfaceBasicIPv4Secondary(secAddr, secAddrMap.(map[string]interface{}), count)
			count = count + 1
			secAddrs = append(secAddrs, secAddr)
		}

		SecAddrMap["vipType"] = "constant"
		SecAddrMap["vipObjectType"] = "tree"
		SecAddrMap["vipValue"] = secAddrs
		SecAddrMap["vipPrimaryKey"] = []string{
			"address",
		}

		ipv4["secondary-address"] = SecAddrMap
		isPresent = isPresent + 1
	}

	if input["dhcp_distance"] != nil && input["dhcp_distance"] != "" {
		val, _ := getInt(input["dhcp_distance"])
		ipv4["dhcp-distance"] = createVIPObject("object", "constant", val, "vpn_if_ipv4_dhcp_distance", nil)
		ipv4["dhcp-client"] = createVIPObject("object", "constant", "true", nil, nil)
		isPresent = isPresent + 1
	}
	if isPresent > 0 {
		defMap["ip"] = ipv4
	}

	if input["dhcp_helper"] != nil && len(input["dhcp_helper"].([]interface{})) > 0 {
		dhcpHelper := make(map[string]interface{})
		dhcpHelper["vipObjectType"] = "list"
		dhcpHelper["vipType"] = "constant"
		dhcpHelper["vipValue"] = input["dhcp_helper"].([]interface{})
		dhcpHelper["vipVariableName"] = "vpn_if_dhcp_helper"
		defMap["dhcp-helper"] = dhcpHelper
	}
}

func createVPNInterfaceBasicIPv4Secondary(defMap map[string]interface{}, input map[string]interface{}, count int) {
	if input["address"] != nil && input["address"] != "" {
		vname := "vpn_if_secondary_address_" + strconv.Itoa(count)
		defMap["address"] = createVIPObject("object", "constant", input["address"], vname, nil)
	}

	defMap["priority-order"] = []string{
		"address",
	}
}

func createVPNInterfaceBasicIPv6(defMap map[string]interface{}, input map[string]interface{}) {
	isPresent := 0
	if input["primary_address"] != nil && input["primary_address"] != "" {
		IPv6["address"] = createVIPObject("object", "constant", input["primary_address"], "vpn_if_ipv6_ipv6_address", nil)
		isPresent = isPresent + 1
	}

	//ipv6 secondary_address configuration
	secAddrSet := input["secondary_address"].(*schema.Set).List()

	if len(secAddrSet) > 0 {
		secAddrs := []interface{}{}
		SecAddrMap := make(map[string]interface{})
		count := 0
		for _, secAddrMap := range secAddrSet {
			secAddr := make(map[string]interface{})
			createVPNInterfaceBasicIPv6Secondary(secAddr, secAddrMap.(map[string]interface{}), count)
			count = count + 1
			secAddrs = append(secAddrs, secAddr)
		}

		SecAddrMap["vipType"] = "constant"
		SecAddrMap["vipObjectType"] = "tree"
		SecAddrMap["vipValue"] = secAddrs
		SecAddrMap["vipPrimaryKey"] = []string{
			"address",
		}

		IPv6["secondary-address"] = SecAddrMap
		isPresent = isPresent + 1
	}

	if input["dhcp_distance"] != nil && input["dhcp_distance"] != "" {
		val, _ := getInt(input["dhcp_distance"])
		IPv6["dhcp-distance"] = createVIPObject("object", "constant", val, "vpn_if_ipv6_ipv6_dhcp_distance", nil)
		IPv6["dhcp-client"] = createVIPObject("object", "constant", "true", nil, nil)
		isPresent = isPresent + 1
	}

	if input["dhcp_rapid_commit"] != nil {
		IPv6["dhcp-rapid-commit"] = createVIPObject("object", "constant", input["dhcp_rapid_commit"], "vpn_if_ipv6_ipv6_dhcp_rapid_commit", nil)
		isPresent = isPresent + 1
	}

	//ipv6 dhcp_helper configuration
	dhcpHelperSet := input["dhcp_helper"].(*schema.Set).List()

	if len(dhcpHelperSet) > 0 {
		helpers := []interface{}{}
		HelperMap := make(map[string]interface{})
		count := 0
		for _, helperMap := range dhcpHelperSet {
			helper := make(map[string]interface{})
			createVPNInterfaceBasicIPv6Helper(helper, helperMap.(map[string]interface{}), count)
			count = count + 1
			helpers = append(helpers, helper)
		}

		HelperMap["vipType"] = "constant"
		HelperMap["vipObjectType"] = "tree"
		HelperMap["vipValue"] = helpers
		HelperMap["vipPrimaryKey"] = []string{
			"address",
		}

		IPv6["dhcp-helper-v6"] = HelperMap
		isPresent = isPresent + 1
	}

	if isPresent > 0 {
		defMap["ipv6"] = IPv6
	}
}

func createVPNInterfaceBasicIPv6Secondary(defMap map[string]interface{}, input map[string]interface{}, count int) {
	if input["address"] != nil && input["address"] != "" {
		vname := "vpn_if_ipv6_secondary_address_ipv6_" + strconv.Itoa(count)
		defMap["address"] = createVIPObject("object", "constant", input["address"], vname, nil)
	}

	defMap["priority-order"] = []string{
		"address",
	}
}

func createVPNInterfaceBasicIPv6Helper(defMap map[string]interface{}, input map[string]interface{}, count int) {
	if input["address"] != nil && input["address"] != "" {
		vname := "vpn_if_dhcp_helper_" + strconv.Itoa(count)
		defMap["address"] = createVIPObject("object", "constant", input["address"], vname, nil)
	}

	if input["vpn"] != nil {
		vname := "vpn_if_ipv6_dhcp_vpn_" + strconv.Itoa(count)
		defMap["vpn"] = createVIPObject("object", "constant", input["vpn"], vname, nil)
	}

	defMap["priority-order"] = []string{
		"address",
		"vpn",
	}
}

func createVPNInterfaceTunnel(defMap map[string]interface{}, input map[string]interface{}, ftType string) error {

	tunnel := make(map[string]interface{})
	isPresent := 0

	tunnelQOS := make(map[string]interface{})
	if input["per_tunnel_qos"] != nil && input["per_tunnel_qos"] == true {
		if input["per_tunnel_qos_aggregator"] == nil || input["per_tunnel_qos_aggregator"] == false {
			tunnelQOS["mode"] = createVIPObject("object", "notIgnore", "spoke", nil, nil)
			tunnel["tunnel-qos"] = tunnelQOS
			isPresent = isPresent + 1
		} else if input["per_tunnel_qos_aggregator"] != nil && input["per_tunnel_qos_aggregator"] == true {
			tunnelQOS["mode"] = createVIPObject("object", "constant", "hub", nil, nil)
			tunnel["tunnel-qos"] = tunnelQOS
			isPresent = isPresent + 1
		}
	}

	if input["tunnels_bandwidth_percent"] != nil && input["tunnels_bandwidth_percent"] != 0 {
		if input["per_tunnel_qos"] == true && input["per_tunnel_qos_aggregator"] == true {
			tunnel["tunnels-bandwidth"] = createVIPObject("object", "constant", input["tunnels_bandwidth_percent"], "vpn_if_tunnel_tunnels-bandwidth", nil)
			isPresent = isPresent + 1
		}
	}

	color := make(map[string]interface{})
	if input["color"] != nil && input["color"] != "" {
		color["value"] = createVIPObject("object", "constant", input["color"], "vpn_if_tunnel_color_value", nil)
		isPresent = isPresent + 1
	}

	if input["restrict"] != nil {
		color["restrict"] = createVIPObject("node-only", "constant", input["restrict"], "vpn_if_tunnel_color_restrict", nil)
		isPresent = isPresent + 1
	}

	tunnel["color"] = color

	if input["groups"] != nil && len(input["groups"].([]interface{})) > 0 {
		groups := make(map[string]interface{})
		groups["vipObjectType"] = "list"
		groups["vipType"] = "constant"
		groups["vipValue"] = input["groups"].([]interface{})
		groups["vipVariableName"] = "vpn_if_tunnel_group"
		tunnel["group"] = groups
		isPresent = isPresent + 1
	}

	if input["border"] != nil {
		tunnel["border"] = createVIPObject("object", "constant", input["border"], "vpn_if_tunnel_border", nil)
		isPresent = isPresent + 1
	}

	if input["control_connection"] != nil {
		tunnel["control-connections"] = createVIPObject("object", "constant", input["control_connection"], "control_connections", nil)
		isPresent = isPresent + 1
	}

	if input["maximum_control_connections"] != nil && input["maximum_control_connections"] != "" {
		val, _ := getInt(input["maximum_control_connections"])
		if ftType == "vpn-vedge-interface" {
			if val >= vedgeMaxControlConnect[0] && val <= vedgeMaxControlConnect[1] {
				tunnel["max-control-connections"] = createVIPObject("object", "constant", val, "vpn_if_tunnel_max_control_connections", nil)
				isPresent = isPresent + 1
			} else {
				return fmt.Errorf("value of maximum_control_connections is out of range for vpn-vedge-interface type feature template")
			}
		} else if ftType == "cisco_vpn_interface" {
			if val >= ciscoMaxControlConnect[0] && val <= ciscoMaxControlConnect[1] {
				tunnel["max-control-connections"] = createVIPObject("object", "constant", val, "vpn_if_tunnel_max_control_connections", nil)
				isPresent = isPresent + 1
			} else {
				return fmt.Errorf("value of maximum_control_connections is out of range for cisco_vpn_interface type feature template")
			}
		}
	}

	if input["vbond_as_stun_server"] != nil {
		tunnel["vbond-as-stun-server"] = createVIPObject("object", "constant", input["vbond_as_stun_server"], "vpn_if_tunnel_vbond_as_stun_server", nil)
		isPresent = isPresent + 1
	}

	if input["exclude_controller_group_list"] != nil && len(input["exclude_controller_group_list"].([]interface{})) > 0 {
		excludeGroup := make(map[string]interface{})
		excludeGroup["vipObjectType"] = "list"
		excludeGroup["vipType"] = "constant"
		excludeGroup["vipValue"] = input["exclude_controller_group_list"].([]interface{})
		excludeGroup["vipVariableName"] = "vpn_if_tunnel_exclude_controller_group_list"
		tunnel["exclude-controller-group-list"] = excludeGroup
		isPresent = isPresent + 1
	}

	if input["vmanage_connection_preference"] != nil {
		tunnel["vmanage-connection-preference"] = createVIPObject("object", "constant", input["vmanage_connection_preference"], "vpn_if_tunnel_vmanage_connection_preference", nil)
		isPresent = isPresent + 1
	}

	if input["port_hop"] != nil {
		tunnel["port-hop"] = createVIPObject("object", "constant", input["port_hop"], "vpn_if_tunnel_port_hop", nil)
		isPresent = isPresent + 1
	}

	if input["low_bandwidth_link"] != nil {
		tunnel["low-bandwidth-link"] = createVIPObject("object", "constant", input["low_bandwidth_link"], "vpn_if_tunnel_low_bandwidth_link", nil)
		isPresent = isPresent + 1
	}

	if len(input["allow_service"].(*schema.Set).List()) > 0 {

		serviceSet := input["allow_service"].(*schema.Set).List()[0].(map[string]interface{})
		allowService := make(map[string]interface{})

		if serviceSet["all"] != nil {
			allowService["all"] = createVIPObject("object", "constant", serviceSet["all"], "vpn_if_tunnel_all", nil)
			isPresent = isPresent + 1
		}

		if serviceSet["bgp"] != nil {
			allowService["bgp"] = createVIPObject("object", "constant", serviceSet["bgp"], "vpn_if_tunnel_bgp", nil)
			isPresent = isPresent + 1
		}

		if serviceSet["dhcp"] != nil {
			allowService["dhcp"] = createVIPObject("object", "constant", serviceSet["dhcp"], "vpn_if_tunnel_dhcp", nil)
			isPresent = isPresent + 1
		}

		if serviceSet["dns"] != nil {
			allowService["dns"] = createVIPObject("object", "constant", serviceSet["dns"], "vpn_if_tunnel_dns", nil)
			isPresent = isPresent + 1
		}

		if serviceSet["icmp"] != nil {
			allowService["icmp"] = createVIPObject("object", "constant", serviceSet["icmp"], "vpn_if_tunnel_icmp", nil)
			isPresent = isPresent + 1
		}

		if serviceSet["netconf"] != nil {
			allowService["netconf"] = createVIPObject("object", "constant", serviceSet["netconf"], "vpn_if_tunnel_netconf", nil)
			isPresent = isPresent + 1
		}

		if serviceSet["ntp"] != nil {
			allowService["ntp"] = createVIPObject("object", "constant", serviceSet["ntp"], "vpn_if_tunnel_ntp", nil)
			isPresent = isPresent + 1
		}

		if serviceSet["ospf"] != nil {
			allowService["ospf"] = createVIPObject("object", "constant", serviceSet["ospf"], "vpn_if_tunnel_ospf", nil)
			isPresent = isPresent + 1
		}

		if serviceSet["ssh"] != nil {
			allowService["sshd"] = createVIPObject("object", "constant", serviceSet["ssh"], "vpn_if_tunnel_sshd", nil)
			isPresent = isPresent + 1
		}

		if serviceSet["stun"] != nil {
			allowService["stun"] = createVIPObject("object", "constant", serviceSet["stun"], "vpn_if_tunnel_stun", nil)
			isPresent = isPresent + 1
		}

		if serviceSet["https"] != nil {
			allowService["https"] = createVIPObject("object", "constant", serviceSet["https"], "vpn_if_tunnel_https", nil)
			isPresent = isPresent + 1
		}

		if ftType == "cisco_vpn_interface" {
			if serviceSet["snmp"] != nil {
				allowService["snmp"] = createVIPObject("object", "constant", serviceSet["snmp"], "vpn_if_tunnel_snmp", nil)
				isPresent = isPresent + 1
			}
		}

		tunnel["allow-service"] = allowService
	}

	if len(input["encapsulation"].(*schema.Set).List()) > 0 {

		encapSet := input["encapsulation"].(*schema.Set).List()[0].(map[string]interface{})
		encapsulation := make([]map[string]interface{}, 0, 2)

		gre := make(map[string]interface{})

		ipsec := make(map[string]interface{})

		isEncapPresent := 0
		if encapSet["gre"] != nil && encapSet["gre"] == true {
			gre["encap"] = createVIPObject("object", "constant", "gre", nil, nil)
			isPresent = isPresent + 1
			isEncapPresent = isEncapPresent + 1

			if encapSet["gre_preference"] != nil && encapSet["gre_preference"] != "" {
				val, _ := getInt(encapSet["gre_preference"])
				gre["preference"] = createVIPObject("object", "constant", val, "vpn_if_tunnel_gre_preference", nil)
				isPresent = isPresent + 1
				isEncapPresent = isEncapPresent + 1
			} else {
				gre["preference"] = createVIPObject("object", "ignore", nil, "vpn_if_tunnel_gre_preference", nil)
				isPresent = isPresent + 1
				isEncapPresent = isEncapPresent + 1
			}

			if encapSet["gre_weight"] != nil {
				gre["weight"] = createVIPObject("object", "constant", encapSet["gre_weight"], "vpn_if_tunnel_gre_weight", nil)
				isPresent = isPresent + 1
				isEncapPresent = isEncapPresent + 1
			}

			gre["priority-order"] = []string{
				"encap",
				"preference",
				"weight",
			}

			encapsulation = append(encapsulation, gre)
		}

		if encapSet["ipsec"] != nil && encapSet["ipsec"] == true {
			ipsec["encap"] = createVIPObject("object", "constant", "ipsec", nil, nil)
			isPresent = isPresent + 1
			isEncapPresent = isEncapPresent + 1

			if encapSet["ipsec_preference"] != nil && encapSet["ipsec_preference"] != "" {
				val, _ := getInt(encapSet["ipsec_preference"])
				ipsec["preference"] = createVIPObject("object", "constant", val, "vpn_if_tunnel_ipsec_preference", nil)
				isPresent = isPresent + 1
				isEncapPresent = isEncapPresent + 1
			} else {
				ipsec["preference"] = createVIPObject("object", "ignore", nil, "vpn_if_tunnel_ipsec_preference", nil)
				isPresent = isPresent + 1
				isEncapPresent = isEncapPresent + 1
			}

			if encapSet["ipsec_weight"] != nil {
				ipsec["weight"] = createVIPObject("object", "constant", encapSet["ipsec_weight"], "vpn_if_tunnel_ipsec_weight", nil)
				isPresent = isPresent + 1
				isEncapPresent = isEncapPresent + 1
			}

			ipsec["priority-order"] = []string{
				"encap",
				"preference",
				"weight",
			}

			encapsulation = append(encapsulation, ipsec)
		}

		var vitType string
		if isEncapPresent > 0 {
			vitType = "constant"
		} else {
			vitType = "ignore"
		}

		encapsulationMap := make(map[string]interface{})
		encapsulationMap["vipObjectType"] = "tree"
		encapsulationMap["vipType"] = vitType
		encapsulationMap["vipValue"] = encapsulation
		encapsulationMap["vipPrimaryKey"] = []string{
			"encap",
		}
		tunnel["encapsulation"] = encapsulationMap
		isPresent = isPresent + 1

		if encapSet["carrier"] != nil && encapSet["carrier"] != "" {
			tunnel["carrier"] = createVIPObject("object", "constant", encapSet["carrier"], "vpn_if_tunnel_carrier", nil)
			isPresent = isPresent + 1
		}

		if encapSet["bind_loopback_tunnel"] != nil && encapSet["bind_loopback_tunnel"] != "" {
			tunnel["bind"] = createVIPObject("object", "constant", encapSet["bind_loopback_tunnel"], "vpn_if_tunnel_bind", nil)
			isPresent = isPresent + 1
		}

		if encapSet["last_resort_circuit"] != nil {
			tunnel["last-resort-circuit"] = createVIPObject("object", "constant", encapSet["last_resort_circuit"], "vpn_if_tunnel_last_resort_circuit", nil)
			isPresent = isPresent + 1
		}

		if encapSet["hold_time"] != nil {
			tunnel["hold-time"] = createVIPObject("object", "constant", encapSet["hold_time"], "hold-time", nil)
			isPresent = isPresent + 1
		}

		if encapSet["nat_refresh_interval"] != nil {
			tunnel["nat-refresh-interval"] = createVIPObject("object", "constant", encapSet["nat_refresh_interval"], "vpn_if_tunnel_nat_refresh_interval", nil)
			isPresent = isPresent + 1
		}

		if encapSet["hello_interval"] != nil {
			tunnel["hello-interval"] = createVIPObject("object", "constant", encapSet["hello_interval"], "vpn_if_tunnel_hello_interval", nil)
			isPresent = isPresent + 1
		}

		if encapSet["hello_tolerance"] != nil {
			tunnel["hello-tolerance"] = createVIPObject("object", "constant", encapSet["hello_tolerance"], "vpn_if_tunnel_hello_tolerance", nil)
			isPresent = isPresent + 1
		}

		if encapSet["gre_tunnel_destination_ip"] != nil && encapSet["gre_tunnel_destination_ip"] != "" {
			destIP := make(map[string]interface{})
			destIP["dst-ip"] = createVIPObject("object", "constant", encapSet["gre_tunnel_destination_ip"], "vpn_if_tunnel_tloc_ext_gre_to_dst_ip", nil)
			tunnel["tloc-extension-gre-to"] = destIP
			isPresent = isPresent + 1
		}
	}

	if isPresent > 0 {
		defMap["tunnel-interface"] = tunnel
	}
	return nil
}

func createVPNInterfaceNAT(defMap map[string]interface{}, input map[string]interface{}, ftType string) error {

	nat := make(map[string]interface{})
	isPresent := 0

	//ipv4
	if len(input["ipv4"].(*schema.Set).List()) > 0 {
		ipv4Set := input["ipv4"].(*schema.Set).List()[0].(map[string]interface{})

		if ipv4Set["nat_type"] != nil && ipv4Set["nat_type"] != "" {
			nat["nat-choice"] = createVIPObject("object", "constant", ipv4Set["nat_type"], "nat-type", nil)
			isPresent = isPresent + 1
		}

		if ipv4Set["refresh_mode"] != nil && ipv4Set["refresh_mode"] != "" {
			nat["refresh"] = createVIPObject("object", "constant", ipv4Set["refresh_mode"], "nat_refresh", nil)
			isPresent = isPresent + 1
		}

		if ipv4Set["log_nat_flow_creations_or_deletions"] != nil {
			nat["log-translations"] = createVIPObject("object", "constant", ipv4Set["log_nat_flow_creations_or_deletions"], "nat_log_translations", nil)
			isPresent = isPresent + 1
		}

		if ipv4Set["udp_timeout"] != nil {
			nat["udp-timeout"] = createVIPObject("object", "constant", ipv4Set["udp_timeout"], "nat_udp_timeout", nil)
			isPresent = isPresent + 1
		}

		if ipv4Set["tcp_timeout"] != nil {
			nat["tcp-timeout"] = createVIPObject("object", "constant", ipv4Set["tcp_timeout"], "nat_tcp_timeout", nil)
			isPresent = isPresent + 1
		}

		if ipv4Set["block_icmp"] != nil {
			nat["block-icmp-error"] = createVIPObject("object", "constant", ipv4Set["block_icmp"], "nat_block_icmp_error", nil)
			isPresent = isPresent + 1
		}

		if ipv4Set["respond_to_ping"] != nil {
			nat["respond-to-ping"] = createVIPObject("object", "constant", ipv4Set["respond_to_ping"], "nat_respond_to_ping", nil)
			isPresent = isPresent + 1
		}

		natpool := make(map[string]interface{})

		isNatpoolPresent := 0

		if ipv4Set["nat_pool_range_start"] != nil && ipv4Set["nat_pool_range_start"] != "" {
			natpool["range-start"] = createVIPObject("object", "constant", ipv4Set["nat_pool_range_start"], "nat_range_start", nil)
			isPresent = isPresent + 1
			isNatpoolPresent = isNatpoolPresent + 1
		}

		if ipv4Set["nat_pool_range_end"] != nil && ipv4Set["nat_pool_range_end"] != "" {
			natpool["range-end"] = createVIPObject("object", "constant", ipv4Set["nat_pool_range_end"], "nat_range_end", nil)
			isPresent = isPresent + 1
			isNatpoolPresent = isNatpoolPresent + 1
		}

		if ipv4Set["nat_pool_prefix_length"] != nil {
			natpool["prefix-length"] = createVIPObject("object", "constant", ipv4Set["nat_pool_prefix_length"], "nat_prefix_length", nil)
			isPresent = isPresent + 1
			isNatpoolPresent = isNatpoolPresent + 1
		}

		if isNatpoolPresent > 0 {
			nat["natpool"] = natpool
		}

		if ipv4Set["overload"] != nil {
			nat["overload"] = createVIPObject("object", "constant", ipv4Set["overload"], "nat_overload", nil)
			isPresent = isPresent + 1
		}

		loopbackInterface := make(map[string]interface{})

		if ipv4Set["nat_inside_source_loopback_interface"] != nil && ipv4Set["nat_inside_source_loopback_interface"] != "" {
			loopbackInterface["loopback-interface"] = createVIPObject("object", "constant", ipv4Set["nat_inside_source_loopback_interface"], "nat_loopback", nil)
			isPresent = isPresent + 1
			nat["interface"] = loopbackInterface
		}

		// port forwarding
		PortForwardingSet := ipv4Set["port_forward"].(*schema.Set).List()

		if len(PortForwardingSet) > 0 {

			rules := make([]interface{}, 0, 128)

			for _, ruleMap := range PortForwardingSet {
				rule := make(map[string]interface{})
				createPortForwardingRule(rule, ruleMap.(map[string]interface{}))

				rules = append(rules, rule)
			}

			RuleMap := make(map[string]interface{})
			RuleMap["vipType"] = "constant"
			RuleMap["vipObjectType"] = "tree"
			RuleMap["vipValue"] = rules
			RuleMap["vipPrimaryKey"] = []string{
				"port-start",
				"port-end",
				"proto",
			}

			nat["port-forward"] = RuleMap
		}

		// static net rule
		StaticNATSet := ipv4Set["static_nat"].(*schema.Set).List()

		if len(StaticNATSet) > 0 {

			rules := []interface{}{}

			for _, ruleMap := range StaticNATSet {
				rule := make(map[string]interface{})
				err := createStaticNATRule(rule, ruleMap.(map[string]interface{}), ftType)
				if err != nil {
					return err
				}

				rules = append(rules, rule)
			}

			RuleMap := make(map[string]interface{})
			RuleMap["vipType"] = "constant"
			RuleMap["vipObjectType"] = "tree"
			RuleMap["vipValue"] = rules
			RuleMap["vipPrimaryKey"] = []string{
				"source-ip",
				"translate-ip",
			}

			nat["static"] = RuleMap
		}
	}

	//ipv6
	if len(input["ipv6"].(*schema.Set).List()) > 0 {
		ipv6Set := input["ipv6"].(*schema.Set).List()[0].(map[string]interface{})

		if ipv6Set["nat64"] != nil {
			defMap["nat64"] = createVIPObject("node-only", "constant", ipv6Set["nat64"], nil, nil)
		}
	}

	if isPresent > 0 {
		defMap["nat"] = nat
		isNATPresent = isNATPresent + 1
	}

	return nil
}

func createPortForwardingRule(defMap map[string]interface{}, input map[string]interface{}) {

	if input["port_start_range"] != nil {
		defMap["port-start"] = createVIPObject("object", "constant", input["port_start_range"], nil, nil)
	}

	if input["port_end_range"] != nil {
		defMap["port-end"] = createVIPObject("object", "constant", input["port_end_range"], nil, nil)
	}

	if input["protocol"] != nil && input["protocol"] != "" {
		defMap["proto"] = createVIPObject("object", "constant", input["protocol"], nil, nil)
	}

	if input["vpn"] != nil {
		defMap["private-vpn"] = createVIPObject("object", "constant", input["vpn"], "vpn_interface_private_vpn", nil)
	}

	if input["private_ip"] != nil && input["private_ip"] != "" {
		defMap["private-ip-address"] = createVIPObject("object", "constant", input["private_ip"], "vpn_interface_private_ip_address", nil)
	}

	defMap["priority-order"] = []string{
		"port-start",
		"port-end",
		"proto",
		"private-vpn",
		"private-ip-address",
	}
}

func createStaticNATRule(defMap map[string]interface{}, input map[string]interface{}, ftType string) error {

	if input["source_ip"] != nil && input["source_ip"] != "" {
		defMap["source-ip"] = createVIPObject("object", "constant", input["source_ip"], "vpn_static_nat_source_ip", nil)
	}

	if input["translated_source_ip"] != nil && input["translated_source_ip"] != "" {
		defMap["translate-ip"] = createVIPObject("object", "constant", input["translated_source_ip"], "vpn_static_nat_translate_ip", nil)
	}

	if input["source_vpn_id"] != nil {
		defMap["source-vpn"] = createVIPObject("object", "constant", input["source_vpn_id"], "vpn_static_nat_source_vpn", nil)
	}

	if input["static_nat_direction"] != nil && input["static_nat_direction"] != "" {
		val := input["static_nat_direction"].(string)
		if ftType == "cisco_vpn_interface" {
			if val == "inside" {
				defMap["static-nat-direction"] = createVIPObject("object", "constant", val, "vpn_static_nat_static_nat_direction", nil)
			} else {
				return fmt.Errorf("value of static_nat_direction is inside and it can't be change for cisco_vpn_interface type feature template")
			}
		} else {
			defMap["static-nat-direction"] = createVIPObject("object", "constant", val, "vpn_static_nat_static_nat_direction", nil)
		}
	}

	if input["protocol"] != nil && input["protocol"] != "" {
		defMap["proto"] = createVIPObject("object", "constant", input["protocol"], "vpn_static_nat_proto", nil)
	}

	if input["source_port"] != nil && input["source_port"] != "" {
		val, _ := getInt(input["source_port"])
		defMap["source-port"] = createVIPObject("object", "constant", val, "vpn_static_nat_source_port", nil)
	}

	if input["translate_port"] != nil && input["translate_port"] != "" {
		val, _ := getInt(input["translate_port"])
		defMap["translate-port"] = createVIPObject("object", "constant", val, "vpn_static_nat_translate_port", nil)
	}

	defMap["priority-order"] = []string{
		"source-ip",
		"translate-ip",
		"source-vpn",
		"static-nat-direction",
		"proto",
		"source-port",
		"translate-port",
	}

	return nil
}

func createVPNInterfaceVRRP(defMap map[string]interface{}, input map[string]interface{}, ftType string) error {

	//ipv4
	VRRPSet := input["ipv4"].(*schema.Set).List()
	if len(VRRPSet) > 0 {

		vrrps := make([]interface{}, 0, 5)

		ciscoIPv4VRRP := 1
		for _, vrrpMap := range VRRPSet {
			vrrp := make(map[string]interface{})
			err := createIPv4VRRP(vrrp, vrrpMap.(map[string]interface{}), ftType)
			if err != nil {
				return err
			}

			if ftType == "cisco_vpn_interface" {
				if ciscoIPv4VRRP <= 1 {
					vrrps = append(vrrps, vrrp)
					ciscoIPv4VRRP += 1
				} else {
					return fmt.Errorf("maximum number of VRRP IPv4 entries for cisco_vpn_interface type feature template is 1!")
				}
			} else {
				vrrps = append(vrrps, vrrp)
			}
		}

		VRRPMap := make(map[string]interface{})
		VRRPMap["vipType"] = "constant"
		VRRPMap["vipObjectType"] = "tree"
		VRRPMap["vipValue"] = vrrps
		VRRPMap["vipPrimaryKey"] = []string{
			"grp-id",
		}

		defMap["vrrp"] = VRRPMap
	}

	//ipv6
	VRRPIPv6Set := input["ipv6"].(*schema.Set).List()
	if len(VRRPIPv6Set) > 0 {

		ipv6vrrps := make([]interface{}, 0, 1)

		for _, ipv6vrrpMap := range VRRPIPv6Set {
			ipv6vrrp := make(map[string]interface{})
			err := createIPv6VRRP(ipv6vrrp, ipv6vrrpMap.(map[string]interface{}), ftType)
			if err != nil {
				return err
			}
			ipv6vrrps = append(ipv6vrrps, ipv6vrrp)
		}

		IPv6VRRPMap := make(map[string]interface{})
		IPv6VRRPMap["vipType"] = "constant"
		IPv6VRRPMap["vipObjectType"] = "tree"
		IPv6VRRPMap["vipValue"] = ipv6vrrps
		IPv6VRRPMap["vipPrimaryKey"] = []string{
			"grp-id",
		}

		defMap["ipv6-vrrp"] = IPv6VRRPMap
	}
	return nil
}

func createIPv4VRRP(defMap map[string]interface{}, input map[string]interface{}, ftType string) error {

	if input["group_id"] != nil {
		defMap["grp-id"] = createVIPObject("object", "constant", input["group_id"], "vpn_if_vrrp_grpid", nil)
	}

	if input["priority"] != nil {
		defMap["priority"] = createVIPObject("object", "constant", input["priority"], "vpn_if_vrrp_priority", nil)
	}

	if input["timer"] != nil && input["timer"] != 0 {
		val := input["timer"].(int)
		if ftType == "vpn-vedge-interface" {
			if val >= vedgeTimer[0] && val <= vedgeTimer[1] {
				defMap["timer"] = createVIPObject("object", "constant", val, "vpn_if_vrrp_timer", nil)
			} else {
				return fmt.Errorf("value of VRRP IPv4 timer is out of range for vpn-vedge-interface type feature template")
			}
		} else if ftType == "cisco_vpn_interface" {
			if val >= ciscoTimer[0] && val <= ciscoTimer[1] {
				defMap["timer"] = createVIPObject("object", "constant", val, "vpn_if_vrrp_timer", nil)
			} else {
				return fmt.Errorf("value of VRRP IPv4 timer is out of range for cisco_vpn_interface type feature template")
			}
		}
	} else {
		if ftType == "vpn-vedge-interface" {
			defMap["timer"] = createVIPObject("object", "constant", vedgeTimerDefault, "vpn_if_vrrp_timer", nil)
		} else if ftType == "cisco_vpn_interface" {
			defMap["timer"] = createVIPObject("object", "constant", ciscoTimerDefault, "vpn_if_vrrp_timer", nil)
		}
	}

	if input["track_omp"] != nil {
		defMap["track-omp"] = createVIPObject("node-only", "constant", input["track_omp"], "vpn_if_vrrp_track_omp", nil)
	}

	if input["track_prefix_list"] != nil && input["track_prefix_list"] != "" {
		defMap["track-prefix-list"] = createVIPObject("object", "constant", input["track_prefix_list"], "vpn_if_vrrp_track_prefix_list", nil)
	}

	if input["ip_address"] != nil && input["ip_address"] != "" {
		ipv4 := make(map[string]interface{})
		ipv4["address"] = createVIPObject("object", "constant", input["ip_address"], "vpn_if_vrrp_vrrp_ipaddress", nil)
		defMap["ipv4"] = ipv4
	}

	defMap["priority-order"] = []string{
		"grp-id",
		"priority",
		"timer",
		"track-omp",
		"track-prefix-list",
		"ipv4",
	}

	return nil
}

func createIPv6VRRP(defMap map[string]interface{}, input map[string]interface{}, ftType string) error {

	if input["group_id"] != nil {
		groupID := make(map[string]interface{})
		groupID["vipObjectType"] = "object"
		groupID["vipType"] = "constant"
		groupID["vipValue"] = input["group_id"].(int)
		groupID["vipVariableName"] = "vpn_if_vrrp_ipv6_grpid_ipv6"
		groupID["ipType"] = "ipv6"
		defMap["grp-id"] = groupID
	}

	if input["priority"] != nil {
		priority := make(map[string]interface{})
		priority["vipObjectType"] = "object"
		priority["vipType"] = "constant"
		priority["vipValue"] = input["priority"].(int)
		priority["vipVariableName"] = "vpn_if_vrrp_ipv6_priority_ipv6"
		priority["ipType"] = "ipv6"
		defMap["priority"] = priority
	}

	if input["timer"] != nil && input["timer"] != 0 {
		val := input["timer"].(int)
		if ftType == "vpn-vedge-interface" {
			if val >= vedgeTimer[0] && val <= vedgeTimer[1] {
				timer := make(map[string]interface{})
				timer["vipObjectType"] = "object"
				timer["vipType"] = "constant"
				timer["vipValue"] = val
				timer["vipVariableName"] = "vpn_if_vrrp_ipv6_timer_ipv6"
				timer["ipType"] = "ipv6"
				defMap["timer"] = timer
			} else {
				return fmt.Errorf("value of VRRP IPv6 timer is out of range for vpn-vedge-interface type feature template")
			}
		} else if ftType == "cisco_vpn_interface" {
			if val >= ciscoTimer[0] && val <= ciscoTimer[1] {
				timer := make(map[string]interface{})
				timer["vipObjectType"] = "object"
				timer["vipType"] = "constant"
				timer["vipValue"] = val
				timer["vipVariableName"] = "vpn_if_vrrp_ipv6_timer_ipv6"
				timer["ipType"] = "ipv6"
				defMap["timer"] = timer
			} else {
				return fmt.Errorf("value of VRRP IPv6 timer is out of range for cisco_vpn_interface type feature template")
			}
		}
	} else {
		if ftType == "vpn-vedge-interface" {
			timer := make(map[string]interface{})
			timer["vipObjectType"] = "object"
			timer["vipType"] = "constant"
			timer["vipValue"] = vedgeTimerDefault
			timer["vipVariableName"] = "vpn_if_vrrp_ipv6_timer_ipv6"
			timer["ipType"] = "ipv6"
			defMap["timer"] = timer
		} else if ftType == "cisco_vpn_interface" {
			timer := make(map[string]interface{})
			timer["vipObjectType"] = "object"
			timer["vipType"] = "constant"
			timer["vipValue"] = ciscoTimerDefault
			timer["vipVariableName"] = "vpn_if_vrrp_ipv6_timer_ipv6"
			timer["ipType"] = "ipv6"
			defMap["timer"] = timer
		}
	}

	if input["track_omp"] != nil {
		trackOmp := make(map[string]interface{})
		trackOmp["vipObjectType"] = "node-only"
		trackOmp["vipType"] = "constant"
		trackOmp["vipValue"] = strconv.FormatBool(input["track_omp"].(bool))
		trackOmp["vipVariableName"] = "vpn_if_vrrp_ipv6_track_omp_ipv6"
		trackOmp["ipType"] = "ipv6"
		defMap["track-omp"] = trackOmp
	}

	if input["track_prefix_list"] != nil && input["track_prefix_list"] != "" {
		trackPreList := make(map[string]interface{})
		trackPreList["vipObjectType"] = "object"
		trackPreList["vipType"] = "constant"
		trackPreList["vipValue"] = input["track_prefix_list"].(string)
		trackPreList["vipVariableName"] = "vpn_if_vrrp_ipv6_track_prefix_list_ipv6"
		trackPreList["ipType"] = "ipv6"
		defMap["track-prefix-list"] = trackPreList
	}

	ipv6 := make(map[string]interface{})

	isIPV6Present := 0
	if input["link_local_ipv6_address"] != nil && input["link_local_ipv6_address"] != "" {
		localIPv6 := make(map[string]interface{})
		localIPv6["vipObjectType"] = "object"
		localIPv6["vipType"] = "constant"
		localIPv6["vipValue"] = input["link_local_ipv6_address"].(string)
		localIPv6["vipVariableName"] = "vpn_if_vrrp_ipv6_vrrp_link_local_address_ipv6"
		localIPv6["ipType"] = "ipv6"
		ipv6["ipv6-link-local"] = localIPv6
		isIPV6Present = isIPV6Present + 1
	}

	if input["global_ipv6_prefix"] != nil && input["global_ipv6_prefix"] != "" {
		globalIPv6 := make(map[string]interface{})
		globalIPv6["vipObjectType"] = "object"
		globalIPv6["vipType"] = "constant"
		globalIPv6["vipValue"] = input["global_ipv6_prefix"].(string)
		globalIPv6["vipVariableName"] = "vpn_if_vrrp_ipv6_vrrp_ip_prefix_ipv6"
		globalIPv6["ipType"] = "ipv6"
		ipv6["prefix"] = globalIPv6
		isIPV6Present = isIPV6Present + 1
	}

	if isIPV6Present > 0 {
		ipv6["priority-order"] = []string{
			"ipv6-link-local",
			"prefix",
		}

		ipv6set := make([]map[string]interface{}, 0, 1)
		ipv6set = append(ipv6set, ipv6)

		ipv6Map := make(map[string]interface{})
		ipv6Map["vipObjectType"] = "tree"
		ipv6Map["vipType"] = "constant"
		ipv6Map["vipValue"] = ipv6set
		ipv6Map["vipPrimaryKey"] = []string{
			"ipv6-link-local",
		}

		defMap["ipv6"] = ipv6Map
	}

	defMap["priority-order"] = []string{
		"grp-id",
		"priority",
		"timer",
		"track-omp",
		"track-prefix-list",
		"ipv6",
	}
	return nil
}

func createVPNInterfaceACL(defMap map[string]interface{}, input map[string]interface{}) {

	//qos-adaptive
	qos_adaptive := make(map[string]interface{})
	isQOSAdPresent := 0

	if input["adapt_period"] != nil && input["adapt_period"] != 0 {
		qos_adaptive["period"] = createVIPObject("object", "constant", input["adapt_period"], "qos_adaptive_period", nil)
		isQOSAdPresent += 1
	} else {
		qos_adaptive["period"] = createVIPObject("object", "ignore", nil, "qos_adaptive_period", nil)
	}

	//upstream
	upstream := make(map[string]interface{})
	isUpstreamPresent := 0

	if input["shaping_rate_upstream_default"] != nil && input["shaping_rate_upstream_default"] != 0.0 {
		upstream["bandwidth-up"] = createVIPObject("object", "constant", input["shaping_rate_upstream_default"], "qos_adaptive_upstream_bandwidth_up", nil)
		isUpstreamPresent += 1
	}

	urange := make(map[string]interface{})
	isURangePresent := 0

	if input["shaping_rate_upstream_min"] != nil && input["shaping_rate_upstream_min"] != 0.0 {
		urange["umin"] = createVIPObject("object", "constant", input["shaping_rate_upstream_min"], "qos_adaptive_upstream_range_umin", nil)
		isURangePresent += 1
	}

	if input["shaping_rate_upstream_max"] != nil && input["shaping_rate_upstream_max"] != 0.0 {
		urange["umax"] = createVIPObject("object", "constant", input["shaping_rate_upstream_max"], "qos_adaptive_upstream_range_umax", nil)
		isURangePresent += 1
	}

	if isURangePresent > 0 {
		upstream["range"] = urange
		isUpstreamPresent += 1
	}

	//downstream
	downstream := make(map[string]interface{})
	isDownstreamPresent := 0

	if input["shaping_rate_downstream_default"] != nil && input["shaping_rate_downstream_default"] != 0.0 {
		downstream["bandwidth-down"] = createVIPObject("object", "constant", input["shaping_rate_downstream_default"], "qos_adaptive_downstream_bandwidth_down", nil)
		isDownstreamPresent += 1
	}

	drange := make(map[string]interface{})
	isDRangePresent := 0

	if input["shaping_rate_downstream_min"] != nil && input["shaping_rate_downstream_min"] != 0.0 {
		drange["dmin"] = createVIPObject("object", "constant", input["shaping_rate_downstream_min"], "qos_adaptive_downstream_range_dmin", nil)
		isDRangePresent += 1
	}

	if input["shaping_rate_downstream_max"] != nil && input["shaping_rate_downstream_max"] != 0.0 {
		drange["dmax"] = createVIPObject("object", "constant", input["shaping_rate_downstream_max"], "qos_adaptive_downstream_range_dmax", nil)
		isDRangePresent += 1
	}

	if isDRangePresent > 0 {
		downstream["range"] = drange
		isDownstreamPresent += 1
	}

	if isUpstreamPresent > 0 {
		qos_adaptive["upstream"] = upstream
		isQOSAdPresent += 1
	}

	if isDownstreamPresent > 0 {
		qos_adaptive["downstream"] = downstream
		isQOSAdPresent += 1
	}

	if isQOSAdPresent > 0 {
		defMap["qos-adaptive"] = qos_adaptive
	}

	if input["shaping_rate"] != nil && input["shaping_rate"] != "" && input["shaping_rate"] != 0.0 {
		defMap["shaping-rate"] = createVIPObject("object", "constant", input["shaping_rate"], "qos_shaping_rate", nil)
	}

	// vManage v20.6.2.2 (possibly others) breaks badly if any interface templates exist without qos-map
	if input["qos_map"] != nil && input["qos_map"] != "" {
		defMap["qos-map"] = createVIPObject("object", "constant", input["qos_map"], "qos_map", nil)
	} else {
		defMap["qos-map"] = createVIPObject("object", "ignore", nil, "qos_map", nil)
	}

	// vManage v20.6.2.2 (possibly others) breaks badly if any interface templates exist without rewrite-rule
	rewriteRule := make(map[string]interface{})
	if input["rewrite_rule"] != nil && input["rewrite_rule"] != "" {
		rewriteRule["rule-name"] = createVIPObject("object", "constant", input["rewrite_rule"], "rewrite_rule_name", nil)
	} else {
		rewriteRule["rule-name"] = createVIPObject("object", "ignore", nil, "rewrite_rule_name", nil)
	}
	defMap["rewrite-rule"] = rewriteRule

	//ipv4 access-list
	ipv4AccessList := make([]map[string]interface{}, 0, 2)
	isIPv4AccessListPresent := 0

	if input["ipv4_ingress_access_list"] != nil && input["ipv4_ingress_access_list"] != "" {
		ingressAccessList := make(map[string]interface{})
		ingressAccessList["acl-name"] = createVIPObject("object", "constant", input["ipv4_ingress_access_list"], "access_list_ingress_acl_name_ipv4", nil)

		ingressAccessList["direction"] = createVIPObject("object", "constant", "in", nil, nil)

		ingressAccessList["priority-order"] = []string{
			"direction",
			"acl-name",
		}
		ipv4AccessList = append(ipv4AccessList, ingressAccessList)

		isIPv4AccessListPresent = isIPv4AccessListPresent + 1
	}

	if input["ipv4_egress_access_list"] != nil && input["ipv4_egress_access_list"] != "" {
		egressAccessList := make(map[string]interface{})
		egressAccessList["acl-name"] = createVIPObject("object", "constant", input["ipv4_egress_access_list"], "ipv4_access_list_egress_acl_name_ipv4", nil)

		egressAccessList["direction"] = createVIPObject("object", "constant", "out", nil, nil)

		egressAccessList["priority-order"] = []string{
			"direction",
			"acl-name",
		}
		ipv4AccessList = append(ipv4AccessList, egressAccessList)

		isIPv4AccessListPresent = isIPv4AccessListPresent + 1
	}

	if isIPv4AccessListPresent > 0 {
		ipv4AccessListMap := make(map[string]interface{})
		ipv4AccessListMap["vipType"] = "constant"
		ipv4AccessListMap["vipObjectType"] = "tree"
		ipv4AccessListMap["vipValue"] = ipv4AccessList
		ipv4AccessListMap["vipPrimaryKey"] = []string{
			"direction",
		}

		defMap["access-list"] = ipv4AccessListMap

	}

	//ipv6 access-list
	ipv6AccessList := make([]map[string]interface{}, 0, 2)
	isIPv6AccessListPresent := 0

	if input["ipv6_ingress_access_list"] != nil && input["ipv6_ingress_access_list"] != "" {
		ingressAccessList := make(map[string]interface{})
		ingressAccessList["acl-name"] = createVIPObject("object", "constant", input["ipv6_ingress_access_list"], "access_list_ingress_acl_name_ipv6", nil)

		ingressAccessList["direction"] = createVIPObject("object", "constant", "in", nil, nil)

		ingressAccessList["priority-order"] = []string{
			"direction",
			"acl-name",
		}
		ipv6AccessList = append(ipv6AccessList, ingressAccessList)

		isIPv6AccessListPresent = isIPv6AccessListPresent + 1
	}

	if input["ipv6_egress_access_list"] != nil && input["ipv6_egress_access_list"] != "" {
		egressAccessList := make(map[string]interface{})
		egressAccessList["acl-name"] = createVIPObject("object", "constant", input["ipv6_egress_access_list"], "access_list_egress_acl_name_ipv6", nil)

		egressAccessList["direction"] = createVIPObject("object", "constant", "out", nil, nil)

		egressAccessList["priority-order"] = []string{
			"direction",
			"acl-name",
		}
		ipv6AccessList = append(ipv6AccessList, egressAccessList)

		isIPv6AccessListPresent = isIPv6AccessListPresent + 1
	}

	if isIPv6AccessListPresent > 0 {
		ipv6AccessListMap := make(map[string]interface{})
		ipv6AccessListMap["vipType"] = "constant"
		ipv6AccessListMap["vipObjectType"] = "tree"
		ipv6AccessListMap["vipValue"] = ipv6AccessList
		ipv6AccessListMap["vipPrimaryKey"] = []string{
			"direction",
		}

		IPv6["access-list"] = ipv6AccessListMap
		defMap["ipv6"] = IPv6
	}

	//Policer
	policer := make([]map[string]interface{}, 0, 2)
	isPolicerPresent := 0

	if input["ingress_policer_name"] != nil && input["ingress_policer_name"] != "" {
		ingressPolicer := make(map[string]interface{})
		ingressPolicer["policer-name"] = createVIPObject("object", "constant", input["ingress_policer_name"], "policer_ingress_policer_name", nil)

		ingressPolicer["direction"] = createVIPObject("object", "constant", "in", nil, nil)

		ingressPolicer["priority-order"] = []string{
			"policer-name",
			"direction",
		}
		policer = append(policer, ingressPolicer)

		isPolicerPresent = isPolicerPresent + 1
	}

	if input["egress_policer_name"] != nil && input["egress_policer_name"] != "" {
		egressPolicer := make(map[string]interface{})
		egressPolicer["policer-name"] = createVIPObject("object", "constant", input["egress_policer_name"], "policer_egress_policer_name", nil)

		egressPolicer["direction"] = createVIPObject("object", "constant", "out", nil, nil)

		egressPolicer["priority-order"] = []string{
			"policer-name",
			"direction",
		}
		policer = append(policer, egressPolicer)

		isPolicerPresent = isPolicerPresent + 1
	}

	if isPolicerPresent > 0 {
		policerMap := make(map[string]interface{})
		policerMap["vipType"] = "constant"
		policerMap["vipObjectType"] = "tree"
		policerMap["vipValue"] = policer
		policerMap["vipPrimaryKey"] = []string{
			"policer-name",
			"direction",
		}

		defMap["policer"] = policerMap
	}
}

func createVPNInterfaceARP(defMap map[string]interface{}, input map[string]interface{}) {

	if input["ip_address"] != nil && input["ip_address"] != "" {
		defMap["addr"] = createVIPObject("object", "constant", input["ip_address"], "vpn_bridge_arp_addr", nil)
	}

	if input["mac_address"] != nil && input["mac_address"] != "" {
		defMap["mac"] = createVIPObject("object", "constant", input["mac_address"], "vpn_bridge_arp_mac", nil)
	}

	defMap["priority-order"] = []string{
		"addr",
		"mac",
	}
}

func createVPNInterfaceAdvanced(defMap map[string]interface{}, input map[string]interface{}, ftType string) error {

	if input["duplex"] != nil && input["duplex"] != "" {
		val := input["duplex"].(string)
		if ftType == "vpn-vedge-interface" {
			if val == "auto" {
				return fmt.Errorf("duplex value auto is not allowed for vpn-vedge-interface type feature template")
			}
			defMap["duplex"] = createVIPObject("object", "constant", val, "vpn_if_duplex", nil)
		} else if ftType == "cisco_vpn_interface" {
			defMap["duplex"] = createVIPObject("object", "constant", val, "vpn_if_duplex", nil)
		}
	} else {
		defMap["duplex"] = createVIPObject("object", "ignore", "_empty", "vpn_if_duplex", nil)
	}

	if input["mac_address"] != nil && input["mac_address"] != "" {
		defMap["mac-address"] = createVIPObject("object", "constant", input["mac_address"], "vpn_if_mac_address", nil)
	}

	if input["ip_mtu"] != nil {
		defMap["mtu"] = createVIPObject("object", "constant", input["ip_mtu"], "vpn_if_ip_mtu", nil)
	}

	if input["pmtu_discovery"] != nil {
		defMap["pmtu"] = createVIPObject("object", "constant", input["pmtu_discovery"], "vpn_if_pmtu", nil)
	}

	if input["flow_control"] != nil && input["flow_control"] != "" {
		defMap["flow-control"] = createVIPObject("object", "constant", input["flow_control"], "vpn_if_flow_control", nil)
	}

	if input["tcp_mss"] != nil && input["tcp_mss"] != "" {
		val, _ := getInt(input["tcp_mss"])
		if ftType == "vpn-vedge-interface" {
			if val >= vedgeTCPMSS[0] && val <= vedgeTCPMSS[1] {
				defMap["tcp-mss-adjust"] = createVIPObject("object", "constant", val, "vpn_if_tcp_mss_adjust", nil)
			} else {
				return fmt.Errorf("value of tcp_mss is out of range for vpn-vedge-interface type feature template")
			}
		} else if ftType == "cisco_vpn_interface" {
			if val >= ciscoTCPMSS[0] && val <= ciscoTCPMSS[1] {
				defMap["tcp-mss-adjust"] = createVIPObject("object", "constant", val, "vpn_if_tcp_mss_adjust", nil)
			} else {
				return fmt.Errorf("value of tcp_mss is out of range for cisco_vpn_interface type feature template")
			}
		}
	}

	if input["speed"] != nil && input["speed"] != "" {
		val := input["speed"].(string)
		if ftType == "vpn-vedge-interface" {
			if val == "2500" {
				return fmt.Errorf("speed value 2500 is not allowed for vpn-vedge-interface type feature template")
			}
			defMap["speed"] = createVIPObject("object", "constant", input["speed"], "vpn_if_speed", nil)
		} else if ftType == "cisco_vpn_interface" {
			defMap["speed"] = createVIPObject("object", "constant", input["speed"], "vpn_if_speed", nil)
		}
	}

	if input["clear_dont_fragment"] != nil {
		defMap["clear-dont-fragment"] = createVIPObject("object", "constant", input["clear_dont_fragment"], "vpn_if_clear_dont_fragment", nil)
	}

	if input["static_ingess_qos"] != nil && input["static_ingess_qos"] != "" {
		val, _ := getInt(input["static_ingess_qos"])
		defMap["static-ingress-qos"] = createVIPObject("object", "constant", val, "vpn_if_static_ingress_qos", nil)
	}

	if input["arp_timeout"] != nil {
		val := input["arp_timeout"].(int)
		if ftType == "vpn-vedge-interface" {
			if val >= vedgeARPTimeOut[0] && val <= vedgeARPTimeOut[1] {
				defMap["arp-timeout"] = createVIPObject("object", "constant", val, "vpn_if_arp_timeout", nil)
			} else {
				return fmt.Errorf("value of arp_timeout is out of range for vpn-vedge-interface type feature template")
			}
		} else if ftType == "cisco_vpn_interface" {
			if val >= ciscoARPTimeOut[0] && val <= ciscoARPTimeOut[1] {
				defMap["arp-timeout"] = createVIPObject("object", "constant", val, "vpn_if_arp_timeout", nil)
			} else {
				return fmt.Errorf("value of arp_timeout is out of range for cisco_vpn_interface type feature template")
			}
		}
	}

	if input["autonegotiation"] != nil {
		defMap["autonegotiate"] = createVIPObject("object", "constant", input["autonegotiation"], "vpn_if_autonegotiate", nil)
	}

	if input["tloc_extension"] != nil && input["tloc_extension"] != "" {
		defMap["tloc-extension"] = createVIPObject("object", "constant", input["tloc_extension"], "vpn_if_tloc_extension", nil)
		isTLOCPresent = isTLOCPresent + 1
	}

	if input["power_over_ethernet"] != nil {
		defMap["poe"] = createVIPObject("object", "constant", input["power_over_ethernet"], "vpn_if_poe", nil)
	}

	if ftType == "cisco_vpn_interface" {
		if input["load_interval"] != nil && input["load_interval"] != 0 {
			defMap["load-interval"] = createVIPObject("object", "constant", input["load_interval"], "vpn_if_load_interval", nil)
		}
	}

	if input["tracker"] != nil && len(input["tracker"].([]interface{})) > 0 {
		tracker := make(map[string]interface{})
		tracker["vipObjectType"] = "list"
		tracker["vipType"] = "constant"
		// tracker["vipValue"] = interfaceToStrList(input["tracker"])
		tracker["vipValue"] = input["tracker"].([]interface{})
		tracker["vipVariableName"] = "vpn_if_tracker"
		defMap["tracker"] = tracker
	}

	if input["icmp_redirect_disable"] != nil {
		defMap["icmp-redirect-disable"] = createVIPObject("object", "constant", input["icmp_redirect_disable"], "vpn_if_icmp_redirect_disable", nil)
	}

	gre := make(map[string]interface{})
	isGREPresent := 0
	if input["gre_tunnel_source_ip"] != nil && input["gre_tunnel_source_ip"] != "" {
		gre["src-ip"] = createVIPObject("object", "constant", input["gre_tunnel_source_ip"], "vpn_if_tloc-ext_gre_from_src_ip", nil)
		isGREPresent = isGREPresent + 1
	}

	if input["xconnect"] != nil && input["xconnect"] != "" {
		gre["xconnect"] = createVIPObject("object", "constant", input["xconnect"], "vpn_if_tloc-ext_gre_from_xconnect", nil)
		isGREPresent = isGREPresent + 1
	}

	if isGREPresent > 0 {
		defMap["tloc-extension-gre-from"] = gre
	}

	if input["ip_directed_broadcast"] != nil {
		defMap["ip-directed-broadcast"] = createVIPObject("object", "constant", input["ip_directed_broadcast"], "vpn_if_ip-directed-broadcast", nil)
	}
	return nil
}

func createVPNInterfaceTrustSec(defMap map[string]interface{}, input map[string]interface{}) {

	trustSec := make(map[string]interface{})
	isPresent := 0

	if input["enable_sgt_propagation"] != nil {
		trustSec["enable"] = createVIPObject("object", "constant", input["enable_sgt_propagation"], "trusted-enable", nil)
		isPresent = isPresent + 1
	}

	if input["propagate"] != nil {
		propagate := make(map[string]interface{})
		propagate["sgt"] = createVIPObject("object", "constant", input["propagate"], nil, nil)
		trustSec["propagate"] = propagate
		isPresent = isPresent + 1
	}

	sgt := make(map[string]interface{})
	isSGTPresent := 0
	if input["security_group_tag"] != nil && input["security_group_tag"] != 0 {
		sgt["sgt"] = createVIPObject("object", "constant", input["security_group_tag"], "trusted__sgt", nil)
		isPresent = isPresent + 1
		isSGTPresent = isSGTPresent + 1
	}

	if input["trusted"] != nil {
		sgt["trusted"] = createVIPObject("node-only", "constant", input["trusted"], nil, nil)
		isPresent = isPresent + 1
		isSGTPresent = isSGTPresent + 1
	}

	if isSGTPresent > 0 {
		trustSec["static"] = sgt
	}

	if isPresent > 0 {
		defMap["trustsec"] = trustSec
	}
}

func createVPNInterface8021X(defMap map[string]interface{}, input map[string]interface{}) error {

	dot1x := make(map[string]interface{})
	isPresent := 0

	if input["radius_server"] != nil && len(input["radius_server"].([]interface{})) > 0 {
		// dot1x["radius-servers"] = createVIPObject("list", "constant", input["radius_server"], "802.1X_radius_server", nil)
		radiusServer := make(map[string]interface{})
		radiusServer["vipObjectType"] = "list"
		radiusServer["vipType"] = "constant"
		radiusServer["vipValue"] = input["radius_server"].([]interface{})
		radiusServer["vipVariableName"] = "802.1X_radius_server"
		dot1x["radius-servers"] = radiusServer
		isPresent = isPresent + 1
	}

	if input["account_interim_interval"] != nil && input["account_interim_interval"] != 0 {
		dot1x["accounting-interval"] = createVIPObject("object", "constant", input["account_interim_interval"], "802.1X_accounting_interval", nil)
		isPresent = isPresent + 1
	}

	if input["nas_identifier"] != nil && input["nas_identifier"] != "" {
		dot1x["nas-identifier"] = createVIPObject("object", "constant", input["nas_identifier"], "802.1X_nas_identifier", nil)
		isPresent = isPresent + 1
	}

	if input["nas_ip"] != nil && input["nas_ip"] != "" {
		dot1x["nas-ip-address"] = createVIPObject("object", "constant", input["nas_ip"], "802.1X_nas_ip_address", nil)
		isPresent = isPresent + 1
	}

	if input["wake_on_lan"] != nil {
		dot1x["wake-on-lan"] = createVIPObject("node-only", "constant", input["wake_on_lan"], "802.1X_wake_on_lan", nil)
		isPresent = isPresent + 1
	}

	if input["control_direction"] != nil && input["control_direction"] != "" {
		dot1x["control-direction"] = createVIPObject("object", "constant", input["control_direction"], "802.1X_control_direction", nil)
		isPresent = isPresent + 1
	}

	if input["reauthentication"] != nil {
		dot1x["reauthentication"] = createVIPObject("object", "constant", input["reauthentication"], "802.1X_reauthentication", nil)
		isPresent = isPresent + 1
	}

	if input["inactivity"] != nil {
		timeout := make(map[string]interface{})
		timeout["inactivity"] = createVIPObject("object", "constant", input["inactivity"], "802.1X_timeout_inactivity", nil)
		dot1x["timeout"] = timeout
		isPresent = isPresent + 1
	}

	if input["host_mode"] != nil && input["host_mode"] != "" {
		dot1x["host-mode"] = createVIPObject("object", "constant", input["host_mode"], "802.1X_host_mode", nil)
		isPresent = isPresent + 1
	}

	//advanced configuration

	if len(input["advanced_options"].(*schema.Set).List()) > 0 {
		advancedSet := input["advanced_options"].(*schema.Set).List()[0].(map[string]interface{})

		if len(advancedSet["vlan"].(*schema.Set).List()) > 0 {

			vlanSet := advancedSet["vlan"].(*schema.Set).List()[0].(map[string]interface{})

			if vlanSet["authentication_fail_vlan"] != nil && vlanSet["authentication_fail_vlan"] != 0 {
				dot1x["auth-fail-vlan"] = createVIPObject("object", "constant", vlanSet["authentication_fail_vlan"], "802.1X_auth_fail_vlan", nil)
				isPresent = isPresent + 1
			}

			if vlanSet["guest_vlan"] != nil && vlanSet["guest_vlan"] != 0 {
				dot1x["guest-vlan"] = createVIPObject("object", "constant", vlanSet["guest_vlan"], "802.1X_guest_vlan", nil)
				isPresent = isPresent + 1
			}

			if vlanSet["authentication_reject_vlan"] != nil && vlanSet["authentication_reject_vlan"] != 0 {
				dot1x["auth-reject-vlan"] = createVIPObject("object", "constant", vlanSet["authentication_reject_vlan"], "802.1X_auth_reject_vlan", nil)
				isPresent = isPresent + 1
			}

			if vlanSet["default_vlan"] != nil && vlanSet["default_vlan"] != 0 {
				dot1x["default-vlan"] = createVIPObject("object", "constant", vlanSet["default_vlan"], "802.1X_default_vlan", nil)
				isPresent = isPresent + 1
			}
		}

		if len(advancedSet["dynamic_authentication_server"].(*schema.Set).List()) > 0 {
			das := make(map[string]interface{})
			isDASPresent := 0

			dasSet := advancedSet["dynamic_authentication_server"].(*schema.Set).List()[0].(map[string]interface{})

			if dasSet["das_port"] != nil {
				das["port"] = createVIPObject("object", "constant", dasSet["das_port"], "802.1X_das_port", nil)
				isDASPresent = isDASPresent + 1
			}

			if dasSet["client"] != nil && dasSet["client"] != "" {
				das["client"] = createVIPObject("object", "constant", dasSet["client"], "802.1X_das_client", nil)
				isDASPresent = isDASPresent + 1
			}

			if dasSet["secret_key"] != nil && dasSet["secret_key"] != "" {
				das["secret-key"] = createVIPObject("object", "constant", dasSet["secret_key"], "802.1X_das_secret_key", nil)
				isDASPresent = isDASPresent + 1
			}

			if dasSet["time_window"] != nil {
				das["time-window"] = createVIPObject("object", "constant", dasSet["time_window"], "802.1X_das_time_window", nil)
				isDASPresent = isDASPresent + 1
			}

			if dasSet["required_timestamp"] != nil {
				das["require-timestamp"] = createVIPObject("node-only", "constant", dasSet["required_timestamp"], "802.1X_das_require_timestamp", nil)
				isDASPresent = isDASPresent + 1
			}

			if dasSet["vpn"] != nil {
				das["vpn"] = createVIPObject("object", "constant", dasSet["vpn"], "802.1X_das_vpn", nil)
				isDASPresent = isDASPresent + 1
			}

			if isDASPresent > 0 {
				dot1x["das"] = das
			}
		}

		if len(advancedSet["mac_authentication_bypass"].(*schema.Set).List()) > 0 {
			mac := make(map[string]interface{})
			isMACPresent := 0

			macSet := advancedSet["mac_authentication_bypass"].(*schema.Set).List()[0].(map[string]interface{})

			if macSet["server"] != nil {
				mac["server"] = createVIPObject("object", "constant", macSet["server"], "802.1X_mab_server", nil)
				isMACPresent = isMACPresent + 1
			}

			if macSet["mac_authentication_bypass_entries"] != nil && len(macSet["mac_authentication_bypass_entries"].([]interface{})) > 0 {
				bypassEntries := make(map[string]interface{})
				bypassEntries["vipObjectType"] = "list"
				bypassEntries["vipType"] = "constant"
				bypassEntries["vipValue"] = macSet["mac_authentication_bypass_entries"].([]interface{})
				bypassEntries["vipVariableName"] = "802.1X_allow_list"
				mac["allow"] = bypassEntries
				isMACPresent = isMACPresent + 1
			}

			if isMACPresent > 0 {
				dot1x["mac-authentication-bypass"] = mac
			}
		}

		if len(advancedSet["request_attributes"].(*schema.Set).List()) > 0 {
			reqAttrSet := advancedSet["request_attributes"].(*schema.Set).List()[0].(map[string]interface{})

			if len(reqAttrSet["authentication"].(*schema.Set).List()) > 0 {
				authSet := reqAttrSet["authentication"].(*schema.Set).List()
				auths := []interface{}{}
				AuthMap := make(map[string]interface{})
				for _, authMap := range authSet {
					auth := make(map[string]interface{})
					err := createVPNInterface8021XAuth(auth, authMap.(map[string]interface{}))
					if err != nil {
						return err
					}

					auths = append(auths, auth)
				}

				AuthMap["vipType"] = "constant"
				AuthMap["vipObjectType"] = "tree"
				AuthMap["vipValue"] = auths
				AuthMap["vipPrimaryKey"] = []string{
					"id",
				}

				dot1x["auth-req-attr"] = AuthMap
				isPresent = isPresent + 1
			}

			if len(reqAttrSet["accounting"].(*schema.Set).List()) > 0 {
				accountSet := reqAttrSet["accounting"].(*schema.Set).List()
				accounts := []interface{}{}
				AccountingMap := make(map[string]interface{})

				for _, accountMap := range accountSet {
					account := make(map[string]interface{})
					err := createVPNInterface8021XAccount(account, accountMap.(map[string]interface{}))
					if err != nil {
						return err
					}
					accounts = append(accounts, account)
				}

				AccountingMap["vipType"] = "constant"
				AccountingMap["vipObjectType"] = "tree"
				AccountingMap["vipValue"] = accounts
				AccountingMap["vipPrimaryKey"] = []string{
					"id",
				}

				dot1x["acct-req-attr"] = AccountingMap
				isPresent = isPresent + 1
			}
		}
	}

	if isPresent > 0 {
		defMap["dot1x"] = dot1x
	} else {
		dot1xMap := make(map[string]interface{})
		dot1xMap["vipType"] = "ignore"
		dot1xMap["vipObjectType"] = "node-only"
		defMap["dot1x"] = dot1xMap
	}
	return nil
}

func createVPNInterface8021XAuth(defMap map[string]interface{}, input map[string]interface{}) error {

	if input["id"] != nil {
		defMap["id"] = createVIPObject("object", "constant", input["id"], "vpn_interface_auth_id", nil)
	}

	if input["syntax_choice"] != nil && input["syntax_choice"] != "" {
		switch input["syntax_choice"] {
		case "string":
			if input["value"] != nil && input["value"] != "" {
				defMap["string"] = createVIPObject("object", "constant", input["value"], "vpn_interface_auth_string", nil)
				defMap["priority-order"] = []string{
					"id",
					"string",
				}
			}
		case "integer":
			if input["value"] != nil {
				value := input["value"].(string)
				if val, err := strconv.Atoi(value); err == nil {
					defMap["integer"] = createVIPObject("object", "constant", val, "vpn_interface_auth_integer", nil)
					defMap["priority-order"] = []string{
						"id",
						"integer",
					}
				} else {
					return fmt.Errorf("[ERROR] Invalid interger value for 802.1X Authentication value")
				}
			}
		case "octet":
			if input["value"] != nil && input["value"] != "" {
				val := input["value"].(string)
				check := HexaValidation(val)
				if check {
					defMap["octet"] = createVIPObject("object", "constant", input["value"], "vpn_interface_auth_octet", nil)
					defMap["priority-order"] = []string{
						"id",
						"octet",
					}
				} else {
					return fmt.Errorf("[ERROR] Invalid Hexa Decimal Number for 802.1X Authentication value")
				}
			}
		default:
			return fmt.Errorf("[ERROR] Invalid Syntax Choice for 802.1X Authentication syntax_choice")
		}
	}
	return nil
}

func createVPNInterface8021XAccount(defMap map[string]interface{}, input map[string]interface{}) error {

	if input["id"] != nil {
		defMap["id"] = createVIPObject("object", "constant", input["id"], "vpn_interface_acct_id", nil)
	}

	if input["syntax_choice"] != nil && input["syntax_choice"] != "" {
		switch input["syntax_choice"] {
		case "string":
			if input["value"] != nil && input["value"] != "" {
				defMap["string"] = createVIPObject("object", "constant", input["value"], "vpn_interface_acct_string{{isUpdateFlow", nil)
				defMap["priority-order"] = []string{
					"id",
					"string",
				}
			}
		case "integer":
			if input["value"] != nil && input["value"] != "" {
				value := input["value"].(string)
				if val, err := strconv.Atoi(value); err == nil {
					defMap["integer"] = createVIPObject("object", "constant", val, "vpn_interface_acct_integer{{isUpdateFlow", nil)
					defMap["priority-order"] = []string{
						"id",
						"integer",
					}
				} else {
					return fmt.Errorf("[ERROR] Invalid interger value for 802.1X Accounting value")
				}
			}
		case "octet":
			if input["value"] != nil && input["value"] != "" {
				val := input["value"].(string)
				check := HexaValidation(val)
				if check {
					defMap["octet"] = createVIPObject("object", "constant", input["value"], "vpn_interface_acct_octet{{isUpdateFlow", nil)
					defMap["priority-order"] = []string{
						"id",
						"octet",
					}
				} else {
					return fmt.Errorf("[ERROR] Invalid Hexa Decimal Number for 802.1X Accounting value")
				}
			}
		default:
			fmt.Errorf("[ERROR] No such Syntax Choice found")
		}
	}
	return nil
}

func validateVPNInterfaceDeviceType(deviceType []string, templateType string) bool {

	if templateType == "vpn-vedge-interface" {
		var check bool
		for _, device := range deviceType {
			check = belongsToVPNVEdgeInterface(device)
			if !check {
				return false
			}
		}
		return true
	}

	if templateType == "cisco_vpn_interface" {
		var check bool
		for _, device := range deviceType {
			check = belongsToCiscoVPNInterface(device)
			if !check {
				return false
			}
		}
		return true
	}

	return false
}

package sdwan

import (
	"fmt"
	"log"
	"strconv"

	"github.com/CiscoDevNet/sdwan-go-client/client"
	"github.com/CiscoDevNet/sdwan-go-client/container"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceSDWANVPNInterfaceFeatureTemplate() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSDWANVPNInterfaceFeatureTemplateRead,

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
						"vpn_interface_basic": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"shutdown": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"interface_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},

									"ipv4": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"primary_address": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"secondary_address": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"address": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"dhcp_distance": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"dhcp_helper": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"ipv6": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"primary_address": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"secondary_address": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"address": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"dhcp_distance": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"dhcp_rapid_commit": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"dhcp_helper": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"address": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"vpn": {
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},

									"block_non_source_ip": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"bandwidth_upstream": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"bandwidth_downstream": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},

						"vpn_interface_tunnel": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"per_tunnel_qos": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"per_tunnel_qos_aggregator": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"tunnels_bandwidth_percent": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"color": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"restrict": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"groups": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"border": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"control_connection": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"maximum_control_connections": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"vbond_as_stun_server": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"exclude_controller_group_list": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"vmanage_connection_preference": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"port_hop": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"low_bandwidth_link": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"allow_service": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"all": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"bgp": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"dhcp": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"dns": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"icmp": {
													Type:     schema.TypeBool,
													Optional: true,
													Default:  true,
												},
												"netconf": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"ntp": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"ospf": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"ssh": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"stun": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"https": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"snmp": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"encapsulation": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"gre": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"gre_preference": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"gre_weight": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"ipsec": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"ipsec_preference": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"ipsec_weight": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"carrier": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"bind_loopback_tunnel": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"last_resort_circuit": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"hold_time": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"nat_refresh_interval": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"hello_interval": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"hello_tolerance": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"gre_tunnel_destination_ip": {
													Type:     schema.TypeString,
													Computed: true,
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nat_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"refresh_mode": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"log_nat_flow_creations_or_deletions": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"udp_timeout": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"tcp_timeout": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"block_icmp": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"respond_to_ping": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"nat_pool_range_start": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"nat_pool_range_end": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"nat_pool_prefix_length": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"overload": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"nat_inside_source_loopback_interface": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"port_forward": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"port_start_range": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"port_end_range": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"protocol": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"vpn": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"private_ip": {
																Type:     schema.TypeString,
																Computed: true,
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
																Type:     schema.TypeString,
																Computed: true,
															},
															"translated_source_ip": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_vpn_id": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"static_nat_direction": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"protocol": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"source_port": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"translate_port": {
																Type:     schema.TypeInt,
																Computed: true,
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
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"nat64": {
													Type:     schema.TypeBool,
													Computed: true,
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"group_id": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"priority": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"timer": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"track_omp": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"track_prefix_list": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"ip_address": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"ipv6": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"group_id": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"priority": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"timer": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"track_omp": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"track_prefix_list": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"link_local_ipv6_address": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"global_ipv6_prefix": {
													Type:     schema.TypeString,
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"adapt_period": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"shapinng_rate_upstream_min": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"shapinng_rate_upstream_max": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"shapinng_rate_upstream_default": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"shapinng_rate_downstream_min": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"shapinng_rate_downstream_max": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"shapinng_rate_downstream_default": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"shapinng_rate": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"qos_map": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"rewrite_rule": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ipv4_ingress_access_list": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ipv4_egress_access_list": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ipv6_ingress_access_list": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ipv6_egress_access_list": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ingress_policer_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"egress_policer_name": {
										Type:     schema.TypeString,
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
										Type:     schema.TypeString,
										Computed: true,
									},
									"mac_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"vpn_interface_8021x": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"radius_server": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"account_interim_interval": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"nas_identifier": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"nas_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"wake_on_lan": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"control_direction": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"reauthentication": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"inactivity": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"host_mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"advanced_options": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vlan": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"authentication_fail_vlan": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"guest_vlan": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"authentication_reject_vlan": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"default_vlan": {
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"dynamic_authentication_server": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"das_port": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"client": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"secret_key": {
																Type:      schema.TypeString,
																Sensitive: true,
																Computed:  true,
															},
															"time_window": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"required_timestamp": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"vpn": {
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
												"mac_authentication_bypass": {
													Type:     schema.TypeSet,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"server": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"mac_authentication_bypass_entries": {
																Type:     schema.TypeList,
																Computed: true,
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
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"authentication": {
																Type:     schema.TypeSet,
																Optional: true,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"id": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"syntax_choice": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"value": {
																			Type:     schema.TypeString,
																			Computed: true,
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
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"syntax_choice": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"value": {
																			Type:     schema.TypeString,
																			Computed: true,
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_sgt_propagation": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"propagate": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"security_group_tag": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"trusted": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},

						"vpn_interface_advanced": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"duplex": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"mac_address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_mtu": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"pmtu_discovery": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"flow_control": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"tcp_mss": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"speed": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"clear_dont_fragment": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"static_ingess_qos": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"arp_timeout": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"autonegotiation": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"tloc_extension": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"power_over_ethernet": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"load_interval": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"tracker": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"icmp_redirect_disable": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"gre_tunnel_source_ip": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"xconnect": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_directed_broadcast": {
										Type:     schema.TypeBool,
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

func datasourceSDWANVPNInterfaceFeatureTemplateRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Read Method ", d.Id())

	sdwanClient := m.(*client.Client)

	ftName := d.Get("template_name").(string)
	cont, err := getFeatureTemplateByName(sdwanClient, ftName)
	if err != nil {
		if cont != nil {
			return fmt.Errorf("[ERROR] Data load failed in between!")
		} else {
			return err
		}
	}

	templateType := stripQuotes(cont.S("templateType").String())
	if templateType != "vpn-vedge-interface" && templateType != "cisco_vpn_interface" {
		return fmt.Errorf("[ERROR] Invalid Template Type")
	}

	setVPNInterfaceTemplateAttributes(d, cont)
	d.SetId(ftName)

	log.Println("[DEBUG] End of Read Method ", d.Id())
	return nil
}

func setVPNInterfaceTemplateAttributes(d *schema.ResourceData, cont *container.Container) *schema.ResourceData {
	log.Println("[DEBUG] Beginning set attribute Method ")

	d.Set("template_name", stripQuotes(cont.S("templateName").String()))

	d.Set("template_description", stripQuotes(cont.S("templateDescription").String()))

	d.Set("device_type", interfaceToStrList(cont.S("deviceType").Data()))

	d.Set("template_type", stripQuotes(cont.S("templateType").String()))

	d.Set("template_min_version", stripQuotes(cont.S("templateMinVersion").String()))

	ftDef := cont.S("templateDefinition")
	FTDefinition := getVPNInterfaceFTDefinition(ftDef)

	d.Set("template_definition", FTDefinition)

	log.Println("End of template definition: ", FTDefinition)

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
	if cont.Exists("rid") {
		if value, err := strconv.Atoi(stripQuotes(cont.S("rid").String())); err == nil {
			d.Set("@rid", value)
		}
	}

	if cont.Exists(("feature")) {
		d.Set("feature", stripQuotes(cont.S("feature").String()))
	}

	if cont.Exists(("createdBy")) {
		d.Set("created_by", stripQuotes(cont.S("createdBy").String()))
	}

	d.Set("template_id", stripQuotes(cont.S("templateId").String()))

	return d
}

//getVPNInterfaceFTDefinition
func getVPNInterfaceFTDefinition(cont *container.Container) []map[string]interface{} {
	log.Println("[DEBUG] Beginning get vpn interface definition Method")

	definition := make([]map[string]interface{}, 0, 1)

	VPNInterfaceDef := make(map[string]interface{})

	VPNInterfaceDef["vpn_interface_basic"] = getVPNInterfaceBasic(cont)

	if cont.Exists("tunnel-interface") {
		VPNInterfaceDef["vpn_interface_tunnel"] = getVPNInterfaceTunnel(cont)
	}

	if cont.Exists("nat") || cont.Exists("nat64") {
		VPNInterfaceDef["vpn_interface_nat"] = getVPNInterfaceNAT(cont)
	}

	if cont.Exists("vrrp") || cont.Exists("ipv6-vrrp") {
		if stripQuotes(cont.S("vrrp", "vipType").String()) != "ignore" || stripQuotes(cont.S("ipv6-vrrp", "vipType").String()) != "ignore" {
			VPNInterfaceDef["vpn_interface_vrrp"] = getVPNInterfaceVRRP(cont)
		}
	}

	VPNInterfaceDef["vpn_interface_acl_qos"] = getVPNInterfaceACL(cont)

	if cont.Exists("arp", "ip") {
		if stripQuotes(cont.S("arp", "ip", "vipType").String()) != "ignore" {
			VPNInterfaceDef["vpn_interface_arp"] = getVPNInterfaceARP(cont)
		}
	}

	if cont.Exists("dot1x") {
		if stripQuotes(cont.S("dot1x", "vipType").String()) != "ignore" {
			VPNInterfaceDef["vpn_interface_8021x"] = getVPNInterface8021X(cont)
		}
	}

	if cont.Exists("trustsec") {
		VPNInterfaceDef["vpn_interface_trustsec"] = getVPNInterfaceTrustsec(cont)
	}

	VPNInterfaceDef["vpn_interface_advanced"] = getVPNInterfaceAdvanced(cont)

	definition = append(definition, VPNInterfaceDef)

	log.Println("[DEBUG] End of get vpn interface definition Method")
	log.Println("FTDefnition : ", definition)
	return definition
}

func getVPNInterfaceBasic(basicCont *container.Container) []map[string]interface{} {
	basic := make([]map[string]interface{}, 0, 1)

	result := make(map[string]interface{})

	if basicCont.Exists("shutdown", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(basicCont.S("shutdown", "vipValue").String())); err == nil {
			result["shutdown"] = value
		}
	}

	if basicCont.Exists("if-name", "vipValue") {
		result["interface_name"] = stripQuotes(basicCont.S("if-name", "vipValue").String())
	}

	if basicCont.Exists("description", "vipValue") {
		result["description"] = stripQuotes(basicCont.S("description", "vipValue").String())
	}

	if basicCont.Exists("ip") {
		log.Println("Inside ip")
		temp := make(map[string]interface{})
		temp1 := make([]map[string]interface{}, 0, 1)
		isIPv4Present := 0
		if basicCont.Exists("ip", "address", "vipValue") {
			temp["primary_address"] = stripQuotes(basicCont.S("ip", "address", "vipValue").String())
			isIPv4Present += 1
		}

		if basicCont.Exists("ip", "secondary-address", "vipValue") {
			address := []map[string]interface{}{}
			addresses := basicCont.S("ip", "secondary-address", "vipValue")

			isSecAddPresent := 0
			for _, val := range addresses.Children() {
				temp2 := make(map[string]interface{})

				if val.Exists("address", "vipValue") {
					temp2["address"] = stripQuotes(val.S("address", "vipValue").String())
					isSecAddPresent += 1
				}

				address = append(address, temp2)
			}

			if isSecAddPresent > 0 {
				temp["secondary_address"] = address
				isIPv4Present += 1
			}
		}

		if basicCont.Exists("ip", "dhcp-distance") {
			if value, err := strconv.Atoi(stripQuotes(basicCont.S("ip", "dhcp-distance", "vipValue").String())); err == nil {
				temp["dhcp_distance"] = value
				isIPv4Present += 1
			}
		}

		if basicCont.Exists("dhcp-helper", "vipValue") {
			temp["dhcp_helper"] = interfaceToStrList(basicCont.S("dhcp-helper", "vipValue").Data())
			isIPv4Present += 1
		}

		if isIPv4Present > 0 {
			temp1 = append(temp1, temp)
			result["ipv4"] = temp1
		}
	}

	if basicCont.Exists("ipv6") {
		temp := map[string]interface{}{}
		temp1 := make([]map[string]interface{}, 0, 1)

		isIPv6Present := 0
		if basicCont.Exists("ipv6", "address", "vipValue") {
			temp["primary_address"] = stripQuotes(basicCont.S("ipv6", "address", "vipValue").String())
			isIPv6Present += 1
		}

		if basicCont.Exists("ipv6", "secondary-address", "vipValue") {
			address := make([]map[string]interface{}, 0, 1)
			addresses := basicCont.S("ipv6", "secondary-address", "vipValue")

			isIPv6SecAddPres := 0
			for _, val := range addresses.Children() {
				temp1 := make(map[string]interface{})

				if val.Exists("address", "vipValue") {
					temp1["address"] = stripQuotes(val.S("address", "vipValue").String())
					isIPv6SecAddPres += 1
				}

				address = append(address, temp1)
			}

			if isIPv6SecAddPres > 0 {
				temp["secondary_address"] = address
				isIPv6Present += 1
			}
		}

		if basicCont.Exists("ipv6", "dhcp-distance", "vipValue") {
			if value, err := strconv.Atoi(stripQuotes(basicCont.S("ipv6", "dhcp-distance", "vipValue").String())); err == nil {
				temp["dhcp_distance"] = value
				isIPv6Present += 1
			}
		}

		if basicCont.Exists("ipv6", "dhcp-rapid-commit", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(basicCont.S("ipv6", "dhcp-rapid-commit", "vipValue").String())); err == nil {
				temp["dhcp_rapid_commit"] = value
				isIPv6Present += 1
			}
		}

		if basicCont.Exists("ipv6", "dhcp-helper-v6", "vipValue") {
			helper := make([]map[string]interface{}, 0, 1)
			helpers := basicCont.S("ipv6", "dhcp-helper-v6", "vipValue")

			isHelperPresent := 0
			for _, val := range helpers.Children() {
				temp2 := make(map[string]interface{})

				if val.Exists("address", "vipValue") {
					temp2["address"] = stripQuotes(val.S("address", "vipValue").String())
					isHelperPresent += 1
				}

				if val.Exists("vpn", "vipValue") {
					vpn, err := strconv.Atoi(val.S("vpn", "vipValue").String())
					if err == nil {
						temp2["vpn"] = vpn
						isHelperPresent += 1
					}
				}

				helper = append(helper, temp2)
			}

			if isHelperPresent > 0 {
				temp["dhcp_helper"] = helper
				isIPv6Present += 1
			}
		}

		if isIPv6Present > 0 {
			temp1 = append(temp1, temp)
			result["ipv6"] = temp1
		}
	}

	if basicCont.Exists("block-non-source-ip", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(basicCont.S("block-non-source-ip", "vipValue").String())); err == nil {
			result["block_non_source_ip"] = value
		}
	}

	if basicCont.Exists("bandwidth-upstream", "vipValue") {
		if value, err := strconv.Atoi(basicCont.S("bandwidth-upstream", "vipValue").String()); err == nil {
			result["bandwidth_upstream"] = value
		}
	}

	if basicCont.Exists("bandwidth-downstream", "vipValue") {
		if value, err := strconv.Atoi(basicCont.S("bandwidth-downstream", "vipValue").String()); err == nil {
			result["bandwidth_downstream"] = value
		}
	}

	basic = append(basic, result)
	return basic
}

func getVPNInterfaceTunnel(tunnelCont *container.Container) []map[string]interface{} {
	tunnel := []map[string]interface{}{}

	result := map[string]interface{}{}
	isPresent := 0

	if tunnelCont.Exists("tunnel-interface", "tunnel-qos", "mode", "vipValue") {
		mode := stripQuotes(tunnelCont.S("tunnel-interface", "tunnel-qos", "mode", "vipValue").String())

		if mode == "hub" {
			result["per_tunnel_qos"] = true
			result["per_tunnel_qos_aggregator"] = true
		} else if mode == "spoke" {
			result["per_tunnel_qos"] = true
			result["per_tunnel_qos_aggregator"] = false
		}
		isPresent += 1
	}

	if tunnelCont.Exists("tunnel-interface", "tunnels-bandwidth", "vipValue") {
		if value, err := strconv.Atoi(tunnelCont.S("tunnel-interface", "tunnels-bandwidth", "vipValue").String()); err == nil {
			result["tunnels_bandwidth_percent"] = value
			isPresent += 1
		}
	}

	if tunnelCont.Exists("tunnel-interface", "color") {
		if tunnelCont.Exists("tunnel-interface", "color", "value", "vipValue") {
			result["color"] = stripQuotes(tunnelCont.S("tunnel-interface", "color", "value", "vipValue").String())
			isPresent += 1
		}

		if tunnelCont.Exists("tunnel-interface", "color", "restrict", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(tunnelCont.S("tunnel-interface", "color", "restrict", "vipValue").String())); err == nil {
				result["restrict"] = value
				isPresent += 1
			}
		}
	}

	if tunnelCont.Exists("tunnel-interface", "group", "vipValue") {
		result["groups"] = interfaceToStrList(tunnelCont.S("tunnel-interface", "group", "vipValue").Data())
		isPresent += 1
	}

	if tunnelCont.Exists("tunnel-interface", "border", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(tunnelCont.S("tunnel-interface", "border", "vipValue").String())); err == nil {
			result["border"] = value
			isPresent += 1
		}
	}

	if tunnelCont.Exists("tunnel-interface", "control-connections", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(tunnelCont.S("tunnel-interface", "control-connections", "vipValue").String())); err == nil {
			result["control_connection"] = value
			isPresent += 1
		}
	}

	if tunnelCont.Exists("tunnel-interface", "max-control-connections", "vipValue") {
		if value, err := strconv.Atoi(tunnelCont.S("tunnel-interface", "max-control-connections", "vipValue").String()); err == nil {
			result["maximum_control_connections"] = value
			isPresent += 1
		}
	}

	if tunnelCont.Exists("tunnel-interface", "vbond-as-stun-server", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(tunnelCont.S("tunnel-interface", "vbond-as-stun-server", "vipValue").String())); err == nil {
			result["vbond_as_stun_server"] = value
			isPresent += 1
		}
	}

	if tunnelCont.Exists("tunnel-interface", "exclude-controller-group-list", "vipValue") {
		result["exclude_controller_group_list"] = interfaceToStrList(tunnelCont.S("tunnel-interface", "exclude-controller-group-list", "vipValue").Data())
		isPresent += 1
	}

	if tunnelCont.Exists("tunnel-interface", "vmanage-connection-preference", "vipValue") {
		if value, err := strconv.Atoi(tunnelCont.S("tunnel-interface", "vmanage-connection-preference", "vipValue").String()); err == nil {
			result["vmanage_connection_preference"] = value
			isPresent += 1
		}
	}

	if tunnelCont.Exists("tunnel-interface", "port-hop", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(tunnelCont.S("tunnel-interface", "port-hop", "vipValue").String())); err == nil {
			result["port_hop"] = value
			isPresent += 1
		}
	}

	if tunnelCont.Exists("tunnel-interface", "low-bandwidth-link", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(tunnelCont.S("tunnel-interface", "low-bandwidth-link", "vipValue").String())); err == nil {
			result["low_bandwidth_link"] = value
			isPresent += 1
		}
	}

	if tunnelCont.Exists("tunnel-interface", "allow-service") {
		temp := map[string]interface{}{}
		temp1 := make([]map[string]interface{}, 0, 1)
		service := tunnelCont.S("tunnel-interface", "allow-service")

		isAllowPresent := 0
		if service.Exists("all", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(service.S("all", "vipValue").String())); err == nil {
				temp["all"] = value
				isAllowPresent += 1
			}
		}

		if service.Exists("bgp", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(service.S("bgp", "vipValue").String())); err == nil {
				temp["bgp"] = value
				isAllowPresent += 1
			}
		}

		if service.Exists("dhcp", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(service.S("dhcp", "vipValue").String())); err == nil {
				temp["dhcp"] = value
				isAllowPresent += 1
			}
		}

		if service.Exists("dns", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(service.S("dns", "vipValue").String())); err == nil {
				temp["dns"] = value
				isAllowPresent += 1
			}
		}

		if service.Exists("icmp", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(service.S("icmp", "vipValue").String())); err == nil {
				temp["icmp"] = value
				isAllowPresent += 1
			}
		}

		if service.Exists("netconf", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(service.S("netconf", "vipValue").String())); err == nil {
				temp["netconf"] = value
				isAllowPresent += 1
			}
		}

		if service.Exists("ntp", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(service.S("ntp", "vipValue").String())); err == nil {
				temp["ntp"] = value
				isAllowPresent += 1
			}
		}

		if service.Exists("ospf", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(service.S("ospf", "vipValue").String())); err == nil {
				temp["ospf"] = value
				isAllowPresent += 1
			}
		}

		if service.Exists("sshd", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(service.S("sshd", "vipValue").String())); err == nil {
				temp["ssh"] = value
				isAllowPresent += 1
			}
		}

		if service.Exists("stun", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(service.S("stun", "vipValue").String())); err == nil {
				temp["stun"] = value
				isAllowPresent += 1
			}
		}

		if service.Exists("https", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(service.S("https", "vipValue").String())); err == nil {
				temp["https"] = value
				isAllowPresent += 1
			}
		}

		if service.Exists("snmp", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(service.S("snmp", "vipValue").String())); err == nil {
				temp["snmp"] = value
				isAllowPresent += 1
			}
		}

		if isAllowPresent > 0 {
			temp1 = append(temp1, temp)
			result["allow_service"] = temp1
			isPresent += 1
		}
	}

	if tunnelCont.Exists("tunnel-interface", "encapsulation") {
		temp := map[string]interface{}{}
		temp3 := make([]map[string]interface{}, 0, 1)
		encap := tunnelCont.S("tunnel-interface", "encapsulation", "vipValue")
		isEncapPresent := 0
		if len(encap.Children()) > 0 {

			temp1 := encap.Children()[0]

			if stripQuotes(temp1.S("encap", "vipValue").String()) == "gre" {
				gre := encap.Children()[0]
				temp["gre"] = true
				if gre.Exists("preference", "vipValue") {
					if value, err := strconv.Atoi(gre.S("preference", "vipValue").String()); err == nil {
						temp["gre_preference"] = value
						isEncapPresent += 1
					}
				}

				if gre.Exists("weight", "vipValue") {
					if value, err := strconv.Atoi(gre.S("weight", "vipValue").String()); err == nil {
						temp["gre_weight"] = value
						isEncapPresent += 1
					}
				}
			}
			if stripQuotes(temp1.S("encap", "vipValue").String()) == "ipsec" {
				ipsec := encap.Children()[0]
				temp["ipsec"] = true
				if ipsec.Exists("preference", "vipValue") {
					if value, err := strconv.Atoi(ipsec.S("preference", "vipValue").String()); err == nil {
						temp["ipsec_preference"] = value
						isEncapPresent += 1
					}
				}

				if ipsec.Exists("weight", "vipValue") {
					if value, err := strconv.Atoi(ipsec.S("weight", "vipValue").String()); err == nil {
						temp["ipsec_weight"] = value
						isEncapPresent += 1
					}
				}
			}

			if len(encap.Children()) > 1 {

				temp2 := encap.Children()[1]

				if stripQuotes(temp2.S("encap", "vipValue").String()) == "gre" {
					gre := encap.Children()[1]
					temp["gre"] = true
					if gre.Exists("preference", "vipValue") {
						if value, err := strconv.Atoi(gre.S("preference", "vipValue").String()); err == nil {
							temp["gre_preference"] = value
							isEncapPresent += 1
						}
					}

					if gre.Exists("weight", "vipValue") {
						if value, err := strconv.Atoi(gre.S("weight", "vipValue").String()); err == nil {
							temp["gre_weight"] = value
							isEncapPresent += 1
						}
					}
				}
				if stripQuotes(temp2.S("encap", "vipValue").String()) == "ipsec" {
					ipsec := encap.Children()[1]
					temp["ipsec"] = true
					if ipsec.Exists("preference", "vipValue") {
						if value, err := strconv.Atoi(ipsec.S("preference", "vipValue").String()); err == nil {
							temp["ipsec_preference"] = value
							isEncapPresent += 1
						}
					}

					if ipsec.Exists("weight", "vipValue") {
						if value, err := strconv.Atoi(ipsec.S("weight", "vipValue").String()); err == nil {
							temp["ipsec_weight"] = value
							isEncapPresent += 1
						}
					}
				}
			}
		}

		if tunnelCont.Exists("tunnel-interface", "carrier", "vipValue") {
			temp["carrier"] = stripQuotes(tunnelCont.S("tunnel-interface", "carrier", "vipValue").String())
			isEncapPresent += 1
		}

		if tunnelCont.Exists("tunnel-interface", "bind", "vipValue") {
			temp["bind_loopback_tunnel"] = stripQuotes(tunnelCont.S("tunnel-interface", "bind", "vipValue").String())
			isEncapPresent += 1
		}

		if tunnelCont.Exists("tunnel-interface", "last-resort-circuit", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(tunnelCont.S("tunnel-interface", "last-resort-circuit", "vipValue").String())); err == nil {
				temp["last_resort_circuit"] = value
				isEncapPresent += 1
			}
		}

		if tunnelCont.Exists("tunnel-interface", "hold-time", "vipValue") {
			if value, err := strconv.Atoi(tunnelCont.S("tunnel-interface", "hold-time", "vipValue").String()); err == nil {
				temp["hold_time"] = value
				isEncapPresent += 1
			}
		}

		if tunnelCont.Exists("tunnel-interface", "nat-refresh-interval", "vipValue") {
			if value, err := strconv.Atoi(tunnelCont.S("tunnel-interface", "nat-refresh-interval", "vipValue").String()); err == nil {
				temp["nat_refresh_interval"] = value
				isEncapPresent += 1
			}
		}

		if tunnelCont.Exists("tunnel-interface", "hello-interval", "vipValue") {
			if value, err := strconv.Atoi(tunnelCont.S("tunnel-interface", "hello-interval", "vipValue").String()); err == nil {
				temp["hello_interval"] = value
				isEncapPresent += 1
			}
		}

		if tunnelCont.Exists("tunnel-interface", "hello-tolerance", "vipValue") {
			if value, err := strconv.Atoi(tunnelCont.S("tunnel-interface", "hello-tolerance", "vipValue").String()); err == nil {
				temp["hello_tolerance"] = value
				isEncapPresent += 1
			}
		}

		if tunnelCont.Exists("tunnel-interface", "tloc-extension-gre-to", "dst-ip", "vipValue") {
			temp["gre_tunnel_destination_ip"] = stripQuotes(tunnelCont.S("tunnel-interface", "tloc-extension-gre-to", "dst-ip", "vipValue").String())
			isEncapPresent += 1
		}

		if isEncapPresent > 0 {
			temp3 = append(temp3, temp)
			result["encapsulation"] = temp3
			isPresent += 1
		}
	}

	if isPresent > 0 {
		tunnel = append(tunnel, result)
		return tunnel
	}
	return nil
}

func getVPNInterfaceNAT(natCont *container.Container) []map[string]interface{} {
	nat := []map[string]interface{}{}

	NATDef := map[string]interface{}{}

	isNATPresent := 0
	if natCont.Exists("nat") {
		if natCont.Exists("nat", "vipObjectType") {
			if stripQuotes(natCont.S("nat", "vipObjectType").String()) != "node-only" {
				NATDef["ipv4"] = getVPNInterfaceNATIPv4(natCont)
				isNATPresent += 1
			}
		} else {
			NATDef["ipv4"] = getVPNInterfaceNATIPv4(natCont)
			isNATPresent += 1
		}

	}

	if natCont.Exists("nat64") {
		NATDef["ipv6"] = getVPNInterfaceNATIPv6(natCont)
		isNATPresent += 1
	}

	if isNATPresent > 0 {
		nat = append(nat, NATDef)
		return nat
	}
	return nil
}

func getVPNInterfaceNATIPv4(natCont *container.Container) []map[string]interface{} {
	ipv4 := []map[string]interface{}{}

	result := map[string]interface{}{}

	natIPv4Cont := natCont.S("nat")

	if natIPv4Cont.Exists("nat-choice", "vipValue") {
		result["nat_type"] = stripQuotes(natIPv4Cont.S("nat-choice", "vipValue").String())
	}

	if natIPv4Cont.Exists("refresh", "vipValue") {
		result["refresh_mode"] = stripQuotes(natIPv4Cont.S("refresh", "vipValue").String())
	}

	if natIPv4Cont.Exists("log-translations", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(natIPv4Cont.S("log-translations", "vipValue").String())); err == nil {
			result["log_nat_flow_creations_or_deletions"] = value
		}
	}

	if natIPv4Cont.Exists("udp-timeout", "vipValue") {
		if value, err := strconv.Atoi(natIPv4Cont.S("udp-timeout", "vipValue").String()); err == nil {
			result["udp_timeout"] = value
		}
	}

	if natIPv4Cont.Exists("tcp-timeout", "vipValue") {
		if value, err := strconv.Atoi(natIPv4Cont.S("tcp-timeout", "vipValue").String()); err == nil {
			result["tcp_timeout"] = value
		}
	}

	if natIPv4Cont.Exists("block-icmp-error", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(natIPv4Cont.S("block-icmp-error", "vipValue").String())); err == nil {
			result["block_icmp"] = value
		}
	}

	if natIPv4Cont.Exists("respond-to-ping", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(natIPv4Cont.S("respond-to-ping", "vipValue").String())); err == nil {
			result["respond_to_ping"] = value
		}
	}

	if natIPv4Cont.Exists("natpool", "range-start", "vipValue") {
		result["nat_pool_range_start"] = stripQuotes(natIPv4Cont.S("natpool", "range-start", "vipValue").String())
	}

	if natIPv4Cont.Exists("natpool", "range-end", "vipValue") {
		result["nat_pool_range_end"] = stripQuotes(natIPv4Cont.S("natpool", "range-end", "vipValue").String())
	}

	if natIPv4Cont.Exists("natpool", "prefix-length", "vipValue") {
		if value, err := strconv.Atoi(natIPv4Cont.S("natpool", "prefix-length", "vipValue").String()); err == nil {
			result["nat_pool_prefix_length"] = value
		}
	}

	if natIPv4Cont.Exists("overload", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(natIPv4Cont.S("overload", "vipValue").String())); err == nil {
			result["overload"] = value
		}
	}

	if natIPv4Cont.Exists("interface", "loopback-interface", "vipValue") {
		result["nat_inside_source_loopback_interface"] = stripQuotes(natIPv4Cont.S("interface", "loopback-interface", "vipValue").String())
	}

	//port forwarding rule

	if natIPv4Cont.Exists("port-forward", "vipValue") {
		rule := []map[string]interface{}{}
		rules := natIPv4Cont.S("port-forward", "vipValue")

		for _, val := range rules.Children() {
			temp := make(map[string]interface{})

			if val.Exists("port-start", "vipValue") {
				if value, err := strconv.Atoi(val.S("port-start", "vipValue").String()); err == nil {
					temp["port_start_range"] = value
				}
			}

			if val.Exists("port-end", "vipValue") {
				if value, err := strconv.Atoi(val.S("port-end", "vipValue").String()); err == nil {
					temp["port_end_range"] = value
				}
			}

			if val.Exists("proto", "vipValue") {
				temp["protocol"] = stripQuotes(val.S("proto", "vipValue").String())
			}

			if val.Exists("private-vpn", "vipValue") {
				if value, err := strconv.Atoi(val.S("private-vpn", "vipValue").String()); err == nil {
					temp["vpn"] = value
				}
			}
			if val.Exists("private-ip-address", "vipValue") {
				temp["private_ip"] = stripQuotes(val.S("private-ip-address", "vipValue").String())
			}

			rule = append(rule, temp)
		}

		result["port_forward"] = rule
	}

	//static nat rule

	if natIPv4Cont.Exists("static", "vipValue") {
		rule := []map[string]interface{}{}
		rules := natIPv4Cont.S("static", "vipValue")

		for _, val := range rules.Children() {
			temp := make(map[string]interface{})

			if val.Exists("source-ip", "vipValue") {
				temp["source_ip"] = stripQuotes(val.S("source-ip", "vipValue").String())
			}

			if val.Exists("translate-ip", "vipValue") {
				temp["translated_source_ip"] = stripQuotes(val.S("translate-ip", "vipValue").String())
			}

			if val.Exists("source-vpn", "vipValue") {
				if value, err := strconv.Atoi(val.S("source-vpn", "vipValue").String()); err == nil {
					temp["source_vpn_id"] = value
				}
			}

			if val.Exists("static-nat-direction", "vipValue") {
				temp["static_nat_direction"] = stripQuotes(val.S("static-nat-direction", "vipValue").String())
			}

			if val.Exists("proto", "vipValue") {
				temp["protocol"] = stripQuotes(val.S("proto", "vipValue").String())
			}

			if val.Exists("source-port", "vipValue") {
				if value, err := strconv.Atoi(val.S("source-port", "vipValue").String()); err == nil {
					temp["source_port"] = value
				}
			}

			if val.Exists("translate-port", "vipValue") {
				if value, err := strconv.Atoi(val.S("translate-port", "vipValue").String()); err == nil {
					temp["translate_port"] = value
				}
			}

			rule = append(rule, temp)
		}

		result["static_nat"] = rule
	}

	ipv4 = append(ipv4, result)
	return ipv4
}

func getVPNInterfaceNATIPv6(natCont *container.Container) []map[string]interface{} {
	ipv6 := []map[string]interface{}{}

	result := map[string]interface{}{}

	if natCont.Exists("nat64", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(natCont.S("nat64", "vipValue").String())); err == nil {
			result["nat64"] = value
		}

		ipv6 = append(ipv6, result)
		return ipv6

	} else if natCont.Exists("nat64", "vipType") {
		if stripQuotes(natCont.S("nat64", "vipType").String()) == "ignore" {
			result["nat64"] = false
		}

		ipv6 = append(ipv6, result)
		return ipv6
	}

	return nil
}

func getVPNInterfaceVRRP(vrrpCont *container.Container) []map[string]interface{} {
	vrrp := []map[string]interface{}{}

	VRRPDef := map[string]interface{}{}

	VRRPDef["ipv4"] = getVPNInterfaceVRRPIpv4(vrrpCont)

	VRRPDef["ipv6"] = getVPNInterfaceVRRPIpv6(vrrpCont)

	vrrp = append(vrrp, VRRPDef)

	return vrrp
}

func getVPNInterfaceVRRPIpv4(vrrpCont *container.Container) []map[string]interface{} {
	ipv4 := []map[string]interface{}{}

	if vrrpCont.Exists("vrrp", "vipValue") {
		ipv4s := vrrpCont.S("vrrp", "vipValue")

		for _, val := range ipv4s.Children() {
			temp := make(map[string]interface{})

			if val.Exists("grp-id", "vipValue") {
				if value, err := strconv.Atoi(val.S("grp-id", "vipValue").String()); err == nil {
					temp["group_id"] = value
				}
			}

			if val.Exists("priority", "vipValue") {
				if value, err := strconv.Atoi(val.S("priority", "vipValue").String()); err == nil {
					temp["priority"] = value
				}
			}

			if val.Exists("timer", "vipValue") {
				if value, err := strconv.Atoi(val.S("timer", "vipValue").String()); err == nil {
					temp["timer"] = value
				}
			}

			if val.Exists("track-omp", "vipValue") {
				if value, err := strconv.ParseBool(stripQuotes(val.S("track-omp", "vipValue").String())); err == nil {
					temp["track_omp"] = value
				}
			}

			if val.Exists("track-prefix-list", "vipValue") {
				temp["track_prefix_list"] = stripQuotes(val.S("track-prefix-list", "vipValue").String())
			}

			if val.Exists("ipv4", "address", "vipValue") {
				temp["ip_address"] = stripQuotes(val.S("ipv4", "address", "vipValue").String())
			}

			ipv4 = append(ipv4, temp)
		}
		return ipv4
	}
	return nil
}

func getVPNInterfaceVRRPIpv6(vrrpCont *container.Container) []map[string]interface{} {
	ipv6 := []map[string]interface{}{}

	if vrrpCont.Exists("ipv6-vrrp", "vipValue") {
		ipv6s := vrrpCont.S("ipv6-vrrp", "vipValue")

		for _, val := range ipv6s.Children() {
			temp := make(map[string]interface{})

			if val.Exists("grp-id", "vipValue") {
				if value, err := strconv.Atoi(val.S("grp-id", "vipValue").String()); err == nil {
					temp["group_id"] = value
				}
			}

			if val.Exists("priority", "vipValue") {
				if value, err := strconv.Atoi(val.S("priority", "vipValue").String()); err == nil {
					temp["priority"] = value
				}
			}

			if val.Exists("timer", "vipValue") {
				if value, err := strconv.Atoi(val.S("timer", "vipValue").String()); err == nil {
					temp["timer"] = value
				}
			}

			if val.Exists("track-omp", "vipValue") {
				if value, err := strconv.ParseBool(stripQuotes(val.S("track-omp", "vipValue").String())); err == nil {
					temp["track_omp"] = value
				}
			}

			if val.Exists("track-prefix-list", "vipValue") {
				temp["track_prefix_list"] = stripQuotes(val.S("track-prefix-list", "vipValue").String())
			}

			ipv6Cont := val.S("ipv6", "vipValue").Children()[0]
			if ipv6Cont.Exists("ipv6-link-local", "vipValue") {
				temp["link_local_ipv6_address"] = stripQuotes(ipv6Cont.S("ipv6-link-local", "vipValue").String())
			}

			if ipv6Cont.Exists("prefix", "vipValue") {
				temp["global_ipv6_prefix"] = stripQuotes(ipv6Cont.S("prefix", "vipValue").String())
			}

			ipv6 = append(ipv6, temp)
		}

	}

	return ipv6
}

func getVPNInterfaceACL(aclCont *container.Container) []map[string]interface{} {
	acl := []map[string]interface{}{}

	result := map[string]interface{}{}
	isPresent := 0

	if aclCont.Exists("qos-adaptive") {
		qosAdp := aclCont.S("qos-adaptive")
		if qosAdp.Exists("period", "vipValue") {
			if value, err := strconv.Atoi(stripQuotes(qosAdp.S("qos-adaptive", "period", "vipValue").String())); err == nil {
				result["adapt_period"] = value
				isPresent += 1
			}
		}

		if qosAdp.Exists("downstream") {
			if qosAdp.Exists("downstream", "bandwidth-down", "vipValue") {
				if value, err := strconv.Atoi(stripQuotes(qosAdp.S("downstream", "bandwidth-down", "vipValue").String())); err == nil {
					result["shapinng_rate_downstream_default"] = value
					isPresent += 1
				}
			}

			if qosAdp.Exists("downstream", "range", "dmin", "vipValue") {
				if value, err := strconv.Atoi(stripQuotes(qosAdp.S("downstream", "range", "dmin", "vipValue").String())); err == nil {
					result["shapinng_rate_downstream_min"] = value
					isPresent += 1
				}
			}
			if qosAdp.Exists("downstream", "range", "dmax", "vipValue") {
				if value, err := strconv.Atoi(stripQuotes(qosAdp.S("downstream", "range", "dmax", "vipValue").String())); err == nil {
					result["shapinng_rate_downstream_max"] = value
					isPresent += 1
				}
			}
		}

		if qosAdp.Exists("upstream") {
			if qosAdp.Exists("upstream", "bandwidth-up", "vipValue") {
				if value, err := strconv.Atoi(stripQuotes(qosAdp.S("upstream", "bandwidth-up", "vipValue").String())); err == nil {
					result["shapinng_rate_upstream_default"] = value
					isPresent += 1
				}
			}

			if qosAdp.Exists("upstream", "range", "umin", "vipValue") {
				if value, err := strconv.Atoi(stripQuotes(qosAdp.S("upstream", "range", "umin", "vipValue").String())); err == nil {
					result["shapinng_rate_upstream_min"] = value
					isPresent += 1
				}
			}
			if qosAdp.Exists("upstream", "range", "umax", "vipValue") {
				if value, err := strconv.Atoi(stripQuotes(qosAdp.S("upstream", "range", "umax", "vipValue").String())); err == nil {
					result["shapinng_rate_upstream_max"] = value
					isPresent += 1
				}
			}
		}

	}

	if aclCont.Exists("shaping-rate", "vipValue") {
		if value, err := strconv.Atoi(stripQuotes(aclCont.S("shaping-rate", "vipValue").String())); err == nil {
			result["shapinng_rate"] = value
			isPresent += 1
		}
	}

	if aclCont.Exists("qos-map", "vipValue") {
		result["qos_map"] = stripQuotes(aclCont.S("qos-map", "vipValue").String())
		isPresent += 1
	}

	if aclCont.Exists("rewrite-rule", "rule-name", "vipValue") {
		result["rewrite_rule"] = stripQuotes(aclCont.S("rewrite-rule", "rule-name", "vipValue").String())
		isPresent += 1
	}

	//IPv4 access-list
	accessList := aclCont.S("access-list", "vipValue")

	if len(accessList.Children()) > 0 {

		temp1 := accessList.Children()[0]

		if stripQuotes(temp1.S("direction", "vipValue").String()) == "in" {

			if temp1.Exists("acl-name", "vipValue") {
				result["ipv4_ingress_access_list"] = stripQuotes(temp1.S("acl-name", "vipValue").String())
				isPresent += 1
			}
		}
		if stripQuotes(temp1.S("direction", "vipValue").String()) == "out" {

			if temp1.Exists("acl-name", "vipValue") {
				result["ipv4_egress_access_list"] = stripQuotes(temp1.S("acl-name", "vipValue").String())
				isPresent += 1
			}
		}

		if len(accessList.Children()) > 1 {

			temp2 := accessList.Children()[1]

			if stripQuotes(temp2.S("direction", "vipValue").String()) == "in" {

				if temp2.Exists("acl-name", "vipValue") {
					result["ipv4_ingress_access_list"] = stripQuotes(temp2.S("acl-name", "vipValue").String())
					isPresent += 1
				}
			}
			if stripQuotes(temp2.S("direction", "vipValue").String()) == "out" {

				if temp2.Exists("acl-name", "vipValue") {
					result["ipv4_egress_access_list"] = stripQuotes(temp2.S("acl-name", "vipValue").String())
					isPresent += 1
				}
			}
		}
	}

	//IPv6 access-list
	accessListIPv6 := aclCont.S("ipv6", "access-list", "vipValue")

	if len(accessListIPv6.Children()) > 0 {

		temp1 := accessListIPv6.Children()[0]

		if stripQuotes(temp1.S("direction", "vipValue").String()) == "in" {

			if temp1.Exists("acl-name", "vipValue") {
				result["ipv6_ingress_access_list"] = stripQuotes(temp1.S("acl-name", "vipValue").String())
				isPresent += 1
			}
		}

		if stripQuotes(temp1.S("direction", "vipValue").String()) == "out" {

			if temp1.Exists("acl-name", "vipValue") {
				result["ipv6_egress_access_list"] = stripQuotes(temp1.S("acl-name", "vipValue").String())
				isPresent += 1
			}
		}

		if len(accessListIPv6.Children()) > 1 {

			temp2 := accessListIPv6.Children()[1]

			if stripQuotes(temp2.S("direction", "vipValue").String()) == "in" {

				if temp2.Exists("acl-name", "vipValue") {
					result["ipv6_ingress_access_list"] = stripQuotes(temp2.S("acl-name", "vipValue").String())
					isPresent += 1
				}
			}
			if stripQuotes(temp2.S("direction", "vipValue").String()) == "out" {

				if temp2.Exists("acl-name", "vipValue") {
					result["ipv6_egress_access_list"] = stripQuotes(temp2.S("acl-name", "vipValue").String())
					isPresent += 1
				}
			}
		}
	}

	//Policer
	policer := aclCont.S("policer", "vipValue")

	if len(policer.Children()) > 0 {

		temp1 := policer.Children()[0]

		if stripQuotes(temp1.S("direction", "vipValue").String()) == "in" {

			if temp1.Exists("policer-name", "vipValue") {
				result["ingress_policer_name"] = stripQuotes(temp1.S("policer-name", "vipValue").String())
				isPresent += 1
			}
		}

		if stripQuotes(temp1.S("direction", "vipValue").String()) == "out" {

			if temp1.Exists("policer-name", "vipValue") {
				result["egress_policer_name"] = stripQuotes(temp1.S("policer-name", "vipValue").String())
				isPresent += 1
			}
		}

		if len(policer.Children()) > 1 {

			temp2 := policer.Children()[1]

			if stripQuotes(temp2.S("direction", "vipValue").String()) == "in" {

				if temp2.Exists("policer-name", "vipValue") {
					result["ingress_policer_name"] = stripQuotes(temp2.S("policer-name", "vipValue").String())
					isPresent += 1
				}
			}
			if stripQuotes(temp2.S("direction", "vipValue").String()) == "out" {

				if temp2.Exists("policer-name", "vipValue") {
					result["egress_policer_name"] = stripQuotes(temp2.S("policer-name", "vipValue").String())
					isPresent += 1
				}
			}
		}
	}

	if isPresent > 0 {
		acl = append(acl, result)
		return acl
	}
	return nil
}

func getVPNInterfaceARP(arpCont *container.Container) []map[string]interface{} {
	arp := []map[string]interface{}{}

	if arpCont.Exists("arp", "ip", "vipValue") {
		arps := arpCont.S("arp", "ip", "vipValue")

		for _, val := range arps.Children() {
			temp := make(map[string]interface{})

			if val.Exists("addr", "vipValue") {
				temp["ip_address"] = stripQuotes(val.S("addr", "vipValue").String())
			}

			if val.Exists("mac", "vipValue") {
				temp["mac_address"] = stripQuotes(val.S("mac", "vipValue").String())
			}

			arp = append(arp, temp)
		}
	}

	return arp
}

func getVPNInterface8021X(dot1Cont *container.Container) []map[string]interface{} {

	dot1xCont := dot1Cont.S("dot1x")

	dot1x := []map[string]interface{}{}

	result := map[string]interface{}{}

	if dot1xCont.Exists("radius-servers", "vipValue") {
		result["radius_server"] = interfaceToStrList(dot1xCont.S("radius-servers", "vipValue").Data())
	}

	if dot1xCont.Exists("accounting-interval", "vipValue") {
		if value, err := strconv.Atoi(stripQuotes(dot1xCont.S("accounting-interval", "vipValue").String())); err == nil {
			result["account_interim_interval"] = value
		}
	}

	if dot1xCont.Exists("nas-identifier", "vipValue") {
		result["nas_identifier"] = stripQuotes(dot1xCont.S("nas-identifier", "vipValue").String())
	}

	if dot1xCont.Exists("nas-ip-address", "vipValue") {
		result["nas_ip"] = stripQuotes(dot1xCont.S("nas-ip-address", "vipValue").String())
	}

	if dot1xCont.Exists("wake-on-lan", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(dot1xCont.S("wake-on-lan", "vipValue").String())); err == nil {
			result["wake_on_lan"] = value
		}
	}

	if dot1xCont.Exists("control-direction", "vipValue") {
		result["control_direction"] = stripQuotes(dot1xCont.S("control-direction", "vipValue").String())
	}

	if dot1xCont.Exists("reauthentication", "vipValue") {
		if value, err := strconv.Atoi(stripQuotes(dot1xCont.S("reauthentication", "vipValue").String())); err == nil {
			result["reauthentication"] = value
		}
	}

	if dot1xCont.Exists("timeout", "inactivity", "vipValue") {
		if value, err := strconv.Atoi(stripQuotes(dot1xCont.S("timeout", "inactivity", "vipValue").String())); err == nil {
			result["inactivity"] = value
		}
	}

	if dot1xCont.Exists("host-mode", "vipValue") {
		result["host_mode"] = stripQuotes(dot1xCont.S("host-mode", "vipValue").String())
	}

	//advanced option
	advanced := map[string]interface{}{}
	advancedTemp := make([]map[string]interface{}, 0, 1)
	isPresent := 0

	//vlan
	vlan := map[string]interface{}{}
	vlanTemp := make([]map[string]interface{}, 0, 1)

	if dot1xCont.Exists("auth-fail-vlan", "vipValue") {
		if value, err := strconv.Atoi(stripQuotes(dot1xCont.S("auth-fail-vlan", "vipValue").String())); err == nil {
			vlan["authentication_fail_vlan"] = value
			isPresent += 1
		}
	}

	if dot1xCont.Exists("guest-vlan", "vipValue") {
		if value, err := strconv.Atoi(stripQuotes(dot1xCont.S("guest-vlan", "vipValue").String())); err == nil {
			vlan["guest_vlan"] = value
			isPresent += 1
		}
	}

	if dot1xCont.Exists("auth-reject-vlan", "vipValue") {
		if value, err := strconv.Atoi(stripQuotes(dot1xCont.S("auth-reject-vlan", "vipValue").String())); err == nil {
			vlan["authentication_reject_vlan"] = value
			isPresent += 1
		}
	}

	if dot1xCont.Exists("default-vlan", "vipValue") {
		if value, err := strconv.Atoi(stripQuotes(dot1xCont.S("default-vlan", "vipValue").String())); err == nil {
			vlan["default_vlan"] = value
			isPresent += 1
		}
	}

	if isPresent > 0 {
		vlanTemp = append(vlanTemp, vlan)
		advanced["vlan"] = vlanTemp
	}

	//dynamic_authentication_server option
	if dot1xCont.Exists("das") {
		isPresent += 1
		dasCont := dot1xCont.S("das")

		das := map[string]interface{}{}
		dasTemp := make([]map[string]interface{}, 0, 1)

		if dasCont.Exists("port", "vipValue") {
			if value, err := strconv.Atoi(stripQuotes(dasCont.S("port", "vipValue").String())); err == nil {
				das["das_port"] = value
			}
		}

		if dasCont.Exists("client", "vipValue") {
			das["client"] = stripQuotes(dasCont.S("client", "vipValue").String())
		}

		if dasCont.Exists("secret-key", "vipValue") {
			das["secret_key"] = stripQuotes(dasCont.S("secret-key", "vipValue").String())
		}

		if dasCont.Exists("time-window", "vipValue") {
			if value, err := strconv.Atoi(stripQuotes(dasCont.S("time-window", "vipValue").String())); err == nil {
				das["time_window"] = value
			}
		}

		if dasCont.Exists("require-timestamp", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(dasCont.S("require-timestamp", "vipValue").String())); err == nil {
				das["required_timestamp"] = value
			}
		}

		if dasCont.Exists("vpn", "vipValue") {
			if value, err := strconv.Atoi(stripQuotes(dasCont.S("vpn", "vipValue").String())); err == nil {
				das["vpn"] = value
			}
		}

		dasTemp = append(dasTemp, das)
		advanced["dynamic_authentication_server"] = dasTemp
	}

	//mac_authentication_bypass option
	if dot1xCont.Exists("mac-authentication-bypass") {
		isPresent += 1
		mabCont := dot1xCont.S("mac-authentication-bypass")

		mab := map[string]interface{}{}
		mabTemp := make([]map[string]interface{}, 0, 1)

		if mabCont.Exists("server", "vipValue") {
			if value, err := strconv.ParseBool(stripQuotes(mabCont.S("server", "vipValue").String())); err == nil {
				mab["server"] = value
			}
		}

		if mabCont.Exists("allow", "vipValue") {
			mab["mac_authentication_bypass_entries"] = interfaceToStrList(mabCont.S("allow", "vipValue").Data())
		}

		mabTemp = append(mabTemp, mab)
		advanced["mac_authentication_bypass"] = mabTemp
	}

	//request_attributes option
	if dot1xCont.Exists("auth-req-attr") || dot1xCont.Exists("acct-req-attr") {
		isPresent += 1
		reqAttrDef := map[string]interface{}{}
		reqAttrDefTemp := make([]map[string]interface{}, 0, 1)

		if dot1xCont.Exists("auth-req-attr") {
			authCont := dot1xCont.S("auth-req-attr")
			reqAttrDef["authentication"] = getVPNInterface8021XReqAttrAuth(authCont)
		}

		if dot1xCont.Exists("acct-req-attr") {
			accCont := dot1xCont.S("acct-req-attr")
			reqAttrDef["accounting"] = getVPNInterface8021XReqAttrAcc(accCont)
		}

		reqAttrDefTemp = append(reqAttrDefTemp, reqAttrDef)
		advanced["request_attributes"] = reqAttrDefTemp
	}

	if isPresent > 0 {
		advancedTemp = append(advancedTemp, advanced)
		result["advanced_options"] = advancedTemp
	}

	dot1x = append(dot1x, result)
	return dot1x
}

func getVPNInterface8021XReqAttrAuth(authCont *container.Container) []map[string]interface{} {
	auth := []map[string]interface{}{}

	if authCont.Exists("vipValue") {
		auths := authCont.S("vipValue")

		for _, val := range auths.Children() {
			temp := make(map[string]interface{})

			var syntax_choice string
			if val.Exists("id", "vipValue") {
				if value, err := strconv.Atoi(val.S("id", "vipValue").String()); err == nil {
					temp["id"] = value
				}
			}

			if val.Exists("priority-order") {
				temp["syntax_choice"] = interfaceToStrList(val.S("priority-order").Data())[1]
				syntax_choice = interfaceToStrList(val.S("priority-order").Data())[1]
			}

			switch {
			case syntax_choice == "string":
				if val.Exists(syntax_choice, "vipValue") {
					temp["value"] = stripQuotes(val.S(syntax_choice, "vipValue").String())
				}
			case syntax_choice == "integer":
				if val.Exists(syntax_choice, "vipValue") {
					temp["value"] = stripQuotes(val.S(syntax_choice, "vipValue").String())
				}
			case syntax_choice == "octet":
				if val.Exists(syntax_choice, "vipValue") {
					temp["value"] = stripQuotes(val.S(syntax_choice, "vipValue").String())
				}
			default:
				fmt.Println("Invalid Syntax")
			}

			auth = append(auth, temp)
		}
	}

	return auth
}

func getVPNInterface8021XReqAttrAcc(accCont *container.Container) []map[string]interface{} {
	acc := []map[string]interface{}{}

	if accCont.Exists("vipValue") {
		accs := accCont.S("vipValue")

		for _, val := range accs.Children() {
			temp := make(map[string]interface{})

			var syntax_choice string
			if val.Exists("id", "vipValue") {
				if value, err := strconv.Atoi(val.S("id", "vipValue").String()); err == nil {
					temp["id"] = value
				}
			}

			if val.Exists("priority-order") {
				temp["syntax_choice"] = interfaceToStrList(val.S("priority-order").Data())[1]
				syntax_choice = interfaceToStrList(val.S("priority-order").Data())[1]
			}

			switch {
			case syntax_choice == "string":
				if val.Exists(syntax_choice, "vipValue") {
					temp["value"] = stripQuotes(val.S(syntax_choice, "vipValue").String())
				}
			case syntax_choice == "integer":
				if val.Exists(syntax_choice, "vipValue") {
					temp["value"] = stripQuotes(val.S(syntax_choice, "vipValue").String())
				}
			case syntax_choice == "octet":
				if val.Exists(syntax_choice, "vipValue") {
					temp["value"] = stripQuotes(val.S(syntax_choice, "vipValue").String())
				}
			default:
				fmt.Println("Invalid Syntax")
			}

			acc = append(acc, temp)
		}
	}

	return acc
}

func getVPNInterfaceTrustsec(trustSecCont *container.Container) []map[string]interface{} {
	trustsec := []map[string]interface{}{}

	result := map[string]interface{}{}

	trustCont := trustSecCont.S("trustsec")
	isPresent := 0

	if trustCont.Exists("enable", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(trustCont.S("enable", "vipValue").String())); err == nil {
			result["enable_sgt_propagation"] = value
			isPresent += 1
		}
	}

	if trustCont.Exists("propagate", "sgt", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(trustCont.S("propagate", "sgt", "vipValue").String())); err == nil {
			result["propagate"] = value
			isPresent += 1
		}
	}

	if trustCont.Exists("static", "sgt", "vipValue") {
		if value, err := strconv.Atoi(trustCont.S("static", "sgt", "vipValue").String()); err == nil {
			result["security_group_tag"] = value
			isPresent += 1
		}
	}

	if trustCont.Exists("static", "trusted", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(trustCont.S("static", "trusted", "vipValue").String())); err == nil {
			result["trusted"] = value
			isPresent += 1
		}
	}

	if isPresent > 0 {
		trustsec = append(trustsec, result)
		return trustsec
	}

	return nil
}

func getVPNInterfaceAdvanced(advancedCont *container.Container) []map[string]interface{} {
	advanced := []map[string]interface{}{}

	result := map[string]interface{}{}
	isPresent := 0

	if advancedCont.Exists("duplex", "vipValue") {
		result["duplex"] = stripQuotes(advancedCont.S("duplex", "vipValue").String())
		isPresent += 1
	}

	if advancedCont.Exists("mac-address", "vipValue") {
		result["mac_address"] = stripQuotes(advancedCont.S("mac-address", "vipValue").String())
		isPresent += 1
	}

	if advancedCont.Exists("mtu", "vipValue") {
		if value, err := strconv.Atoi(advancedCont.S("mtu", "vipValue").String()); err == nil {
			result["ip_mtu"] = value
			isPresent += 1
		}
	}

	if advancedCont.Exists("pmtu", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(advancedCont.S("pmtu", "vipValue").String())); err == nil {
			result["pmtu_discovery"] = value
			isPresent += 1
		}
	}

	if advancedCont.Exists("flow-control", "vipValue") {
		result["flow_control"] = stripQuotes(advancedCont.S("flow-control", "vipValue").String())
		isPresent += 1
	}

	if advancedCont.Exists("tcp-mss-adjust", "vipValue") {
		if value, err := strconv.Atoi(advancedCont.S("tcp-mss-adjust", "vipValue").String()); err == nil {
			result["tcp_mss"] = value
			isPresent += 1
		}
	}

	if advancedCont.Exists("speed", "vipValue") {
		result["speed"] = stripQuotes(advancedCont.S("speed", "vipValue").String())
		isPresent += 1
	}

	if advancedCont.Exists("clear-dont-fragment", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(advancedCont.S("clear-dont-fragment", "vipValue").String())); err == nil {
			result["clear_dont_fragment"] = value
			isPresent += 1
		}
	}

	if advancedCont.Exists("static-ingress-qos", "vipValue") {
		if value, err := strconv.Atoi(advancedCont.S("static-ingress-qos", "vipValue").String()); err == nil {
			result["static_ingess_qos"] = value
			isPresent += 1
		}
	}

	if advancedCont.Exists("arp-timeout", "vipValue") {
		if value, err := strconv.Atoi(advancedCont.S("arp-timeout", "vipValue").String()); err == nil {
			result["arp_timeout"] = value
			isPresent += 1
		}
	}

	if advancedCont.Exists("autonegotiate", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(advancedCont.S("autonegotiate", "vipValue").String())); err == nil {
			result["autonegotiation"] = value
			isPresent += 1
		}
	}

	if advancedCont.Exists("tloc-extension", "vipValue") {
		result["tloc_extension"] = stripQuotes(advancedCont.S("tloc-extension", "vipValue").String())
		isPresent += 1
	}

	if advancedCont.Exists("poe", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(advancedCont.S("poe", "vipValue").String())); err == nil {
			result["power_over_ethernet"] = value
			isPresent += 1
		}
	}

	if advancedCont.Exists("load-interval", "vipValue") {
		if value, err := strconv.Atoi(advancedCont.S("load-interval", "vipValue").String()); err == nil {
			result["load_interval"] = value
			isPresent += 1
		}
	}

	if advancedCont.Exists("tracker", "vipValue") {
		result["tracker"] = interfaceToStrList(advancedCont.S("tracker", "vipValue").Data())
		isPresent += 1
	}

	if advancedCont.Exists("icmp-redirect-disable", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(advancedCont.S("icmp-redirect-disable", "vipValue").String())); err == nil {
			result["icmp_redirect_disable"] = value
			isPresent += 1
		}
	}

	if advancedCont.Exists("tloc-extension-gre-from", "src-ip", "vipValue") {
		result["gre_tunnel_source_ip"] = stripQuotes(advancedCont.S("tloc-extension-gre-from", "src-ip", "vipValue").String())
		isPresent += 1
	}

	if advancedCont.Exists("tloc-extension-gre-from", "xconnect", "vipValue") {
		result["xconnect"] = stripQuotes(advancedCont.S("tloc-extension-gre-from", "xconnect", "vipValue").String())
		isPresent += 1
	}

	if advancedCont.Exists("ip-directed-broadcast", "vipValue") {
		if value, err := strconv.ParseBool(stripQuotes(advancedCont.S("ip-directed-broadcast", "vipValue").String())); err == nil {
			result["ip_directed_broadcast"] = value
			isPresent += 1
		}
	}

	if isPresent > 0 {
		advanced = append(advanced, result)
		return advanced
	}
	return nil
}

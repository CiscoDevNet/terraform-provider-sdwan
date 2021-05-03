package sdwan

// import (
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
// )

// // Structure for Resource

// func generateSystemTemplateSchema() map[string]*schema.Schema {
// 	return map[string]*schema.Schema{
// 		//For Basic Configuration
// 		"system_basic": {
// 			Type:     schema.TypeList,
// 			Optional: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"site_id": {
// 						Type:         schema.TypeInt,
// 						Required:     true,
// 						ValidateFunc: validation.IntBetween(1, 4294967295),
// 					},
// 					"system_ip": {
// 						Type:         schema.TypeString,
// 						Required:     true,
// 						ValidateFunc: validation.IsIPv4Address,
// 					},
// 					"overlay_id": {
// 						Type:     schema.TypeInt,
// 						Optional: true,
// 					},
// 					"timezone": {
// 						Type:     schema.TypeString,
// 						Optional: true,
// 						Default:  "UTC",
// 						//TODO
// 						ValidateFunc: validation.StringInSlice([]string{
// 							"UTC",
// 							"Europe/Zurich",
// 							"Asia/Dubai",
// 						}, false),
// 					},
// 					"hostname": {
// 						Type:     schema.TypeString,
// 						Optional: true,
// 					},
// 					"location": {
// 						Type:     schema.TypeString,
// 						Optional: true,
// 					},
// 					"device_groups": {
// 						Type: schema.TypeList,
// 						Elem: &schema.Schema{
// 							Type: schema.TypeString,
// 						},
// 						Optional: true,
// 					},
// 					"controller_groups": {
// 						Type: schema.TypeList,
// 						Elem: &schema.Schema{
// 							Type: schema.TypeString,
// 						},
// 						Optional: true,
// 					},
// 					"description": {
// 						Type:     schema.TypeString,
// 						Optional: true,
// 					},
// 					"console_baud_rate": {
// 						Type:     schema.TypeInt,
// 						Optional: true,
// 					},
// 					"max_omp_sessions": {
// 						Type:     schema.TypeInt,
// 						Optional: true,
// 					},
// 					"tcp_optimization": {
// 						Type:     schema.TypeBool,
// 						Optional: true,
// 						Default:  false,
// 					},
// 				},
// 			},
// 		},

// 		//For GPS Configuration
// 		"system_gps": {
// 			Type:     schema.TypeList,
// 			Optional: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"gps_latitude": {
// 						Type:     schema.TypeFloat,
// 						Optional: true,
// 					},
// 					"gps_longitude": {
// 						Type:     schema.TypeFloat,
// 						Optional: true,
// 					},
// 				},
// 			},
// 		},

// 		//For Tracker configuration
// 		"system_trackers": {
// 			Type:     schema.TypeList,
// 			MaxItems: 8,
// 			Optional: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"tracker_name": {
// 						Type:         schema.TypeString,
// 						Optional:     true,
// 						ValidateFunc: NameValidation(),
// 					},
// 					"tracker_endpoint_ip": {
// 						Type:     schema.TypeString,
// 						Optional: true,
// 					},
// 					"tracker_endpoint_dns_name": {
// 						Type:     schema.TypeString,
// 						Optional: true,
// 					},
// 					"tracker_endpoint_api_url": {
// 						Type:     schema.TypeString,
// 						Optional: true,
// 					},
// 					"tracker_threshold": {
// 						Type:         schema.TypeInt,
// 						Optional:     true,
// 						Default:      300,
// 						ValidateFunc: validation.IntBetween(100, 1000),
// 					},
// 					"tracker_interval": {
// 						Type:         schema.TypeInt,
// 						Optional:     true,
// 						Default:      60,
// 						ValidateFunc: validation.IntBetween(10, 600),
// 					},
// 					"tracker_multiplier": {
// 						Type:         schema.TypeInt,
// 						Optional:     true,
// 						Default:      3,
// 						ValidateFunc: validation.IntBetween(1, 10),
// 					},
// 					// "tracker_type": {
// 					// 	Type:     schema.TypeString,
// 					// 	Optional: true,
// 					// },
// 				},
// 			},
// 		},

// 		//For Advanced Configuration

// 		"system_advanced": {
// 			Type:     schema.TypeList,
// 			Optional: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"control_session_pps": {
// 						Type:         schema.TypeInt,
// 						Optional:     true,
// 						Default:      300,
// 						ValidateFunc: validation.IntBetween(1, 65535),
// 					},
// 					"system_tunnel_mtu": {
// 						Type:         schema.TypeInt,
// 						Optional:     true,
// 						Default:      1024,
// 						ValidateFunc: validation.IntBetween(500, 2000),
// 					},
// 					"port_hop": {
// 						Type:     schema.TypeBool,
// 						Optional: true,
// 						Default:  true,
// 					},
// 					"port_offset": {
// 						Type:         schema.TypeInt,
// 						Optional:     true,
// 						Default:      0,
// 						ValidateFunc: validation.IntBetween(0, 19),
// 					},
// 					"timer_dns_cache_timeout": {
// 						Type:         schema.TypeInt,
// 						Optional:     true,
// 						Default:      30,
// 						ValidateFunc: validation.IntBetween(1, 30),
// 					},
// 					"track_transport": {
// 						Type:     schema.TypeBool,
// 						Optional: true,
// 						Default:  true,
// 					},
// 					"vbond_local": {
// 						Type:     schema.TypeBool,
// 						Default:  false,
// 						Optional: true,
// 					},
// 					"vbond_remote": {
// 						Type:     schema.TypeString,
// 						Optional: true,
// 					},
// 					"track_interface_tag": {
// 						Type:         schema.TypeInt,
// 						Optional:     true,
// 						ValidateFunc: validation.IntBetween(1, 4294967295),
// 					},
// 					"multicast_buffer_percent": {
// 						Type:         schema.TypeInt,
// 						Optional:     true,
// 						Default:      20,
// 						ValidateFunc: validation.IntBetween(5, 100),
// 					},
// 					"usb_controller": {
// 						Type:     schema.TypeBool,
// 						Optional: true,
// 						Default:  false,
// 					},
// 					"track_default_gateway": {
// 						Type:     schema.TypeBool,
// 						Optional: true,
// 						Default:  true,
// 					},
// 					"host_policier_pps": {
// 						Type:         schema.TypeInt,
// 						Optional:     true,
// 						Default:      5000,
// 						ValidateFunc: validation.IntBetween(1000, 20000),
// 					},
// 					"icmp_error_pps": {
// 						Type:         schema.TypeInt,
// 						Optional:     true,
// 						Default:      100,
// 						ValidateFunc: validation.IntBetween(1, 200),
// 					},
// 					"allow_same_site_tunnels": {
// 						Type:     schema.TypeBool,
// 						Optional: true,
// 						Default:  false,
// 					},
// 					"route_consistency_check": {
// 						Type:     schema.TypeBool,
// 						Optional: true,
// 						Default:  false,
// 					},
// 					"admin_tech_on_failure": {
// 						Type:     schema.TypeBool,
// 						Optional: true,
// 						Default:  true,
// 					},
// 					"idle_timeout": {
// 						Type:         schema.TypeInt,
// 						Optional:     true,
// 						ValidateFunc: validation.IntBetween(0, 300),
// 					},
// 					"eco_friendly_mode": {
// 						Type:     schema.TypeBool,
// 						Optional: true,
// 						Default:  false,
// 					},
// 					"on_demand_enable": {
// 						Type:     schema.TypeBool,
// 						Optional: true,
// 						Default:  false,
// 					},
// 					"on_demand_idle_timeout": {
// 						Type:     schema.TypeInt,
// 						Optional: true,
// 						Default:  10,
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func generateNTPTemplateSchema() map[string]*schema.Schema {
// 	return map[string]*schema.Schema{
// 		// For Server Configuration
// 		"servers": {
// 			Type:     schema.TypeList,
// 			Optional: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"hostname": {
// 						Type:     schema.TypeString,
// 						Required: true,
// 					},
// 					"key": {
// 						Type:     schema.TypeInt,
// 						Required: true,
// 					},
// 					"vpn": {
// 						Type:         schema.TypeInt,
// 						Required:     true,
// 						ValidateFunc: validation.IntBetween(0, 65530),
// 					},
// 					"version": {
// 						Type:         schema.TypeInt,
// 						Optional:     true,
// 						Default:      4,
// 						ValidateFunc: validation.IntBetween(1, 4),
// 					},
// 					"source_interface": {
// 						Type:     schema.TypeString,
// 						Optional: true,
// 					},
// 					"prefer": {
// 						Type:     schema.TypeBool,
// 						Optional: true,
// 					},
// 				},
// 			},
// 		},

// 		// For Authentication Configuration
// 		"authentication_keys": {
// 			Type:     schema.TypeList,
// 			Optional: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"id": {
// 						Type:     schema.TypeInt,
// 						Required: true,
// 					},
// 					"value": {
// 						Type:     schema.TypeString,
// 						Required: true,
// 					},
// 				},
// 			},
// 		},

// 		"trusted_key": {
// 			Type: schema.TypeList,
// 			Elem: &schema.Schema{
// 				Type: schema.TypeString,
// 			},
// 			Optional: true,
// 		},
// 	}
// }

// func generateVPNTemplateSchema() map[string]*schema.Schema {
// 	return map[string]*schema.Schema{
// 		"vpn_basic": {
// 			Type:     schema.TypeList,
// 			Optional: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"vpn": {
// 						Type:         schema.TypeInt,
// 						Required:     true,
// 						ValidateFunc: validation.IntBetween(0, 65530),
// 					},
// 					"name": {
// 						Type:         schema.TypeString,
// 						Optional:     true,
// 						ValidateFunc: NameValidation(),
// 					},
// 					"enhace_ecmp_keying": {
// 						Type:     schema.TypeBool,
// 						Optional: true,
// 						Default:  false,
// 					},
// 					"enhance_tcp_optimization": {
// 						Type:     schema.TypeBool,
// 						Optional: true,
// 						Default:  false,
// 					},
// 				},
// 			},
// 		},
// 		"vpn_dns": {
// 			Type:     schema.TypeList,
// 			Optional: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"dns_ipv4_primary": {
// 						Type:         schema.TypeString,
// 						Optional:     true,
// 						ValidateFunc: validation.IsIPv4Address,
// 					},
// 					"dns_ipv4_secondary": {
// 						Type:         schema.TypeString,
// 						Optional:     true,
// 						ValidateFunc: validation.IsIPv4Address,
// 					},
// 					"dns_ipv6_primary": {
// 						Type:         schema.TypeString,
// 						Optional:     true,
// 						ValidateFunc: validation.IsIPv6Address,
// 					},
// 					"dns_ipv6_secondary": {
// 						Type:         schema.TypeString,
// 						Optional:     true,
// 						ValidateFunc: validation.IsIPv6Address,
// 					},
// 					"host": {
// 						Type:     schema.TypeList,
// 						Optional: true,
// 						Elem: &schema.Resource{
// 							Schema: map[string]*schema.Schema{
// 								"hostname": {
// 									Type:     schema.TypeString,
// 									Required: true,
// 								},
// 								"ip_addresses": {
// 									Type:     schema.TypeList,
// 									MaxItems: 8,
// 									Elem: &schema.Schema{
// 										Type: schema.TypeString,
// 									},
// 									Required: true,
// 								},
// 							},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func generateVPNInterfaceTemplateSchema() map[string]*schema.Schema {
// 	return map[string]*schema.Schema{
// 		//For Basic Configuration
// 		"vpn_interface_basic": {
// 			Type:     schema.TypeList,
// 			Required: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"shutdown": {
// 						Type:     schema.TypeBool,
// 						Required: true,
// 					},
// 					"interface_name": {
// 						Type:         schema.TypeString,
// 						Required:     true,
// 						ValidateFunc: validation.StringInSlice(validInterfaceTypes, false),
// 					},
// 					"description": {
// 						Type:     schema.TypeString,
// 						Optional: true,
// 					},
// 					//TODO:Check for validation
// 					"static_ipv4": {
// 						Type:     schema.TypeList,
// 						Optional: true,
// 						Elem: &schema.Resource{
// 							Schema: map[string]*schema.Schema{
// 								"primary_address": {
// 									Type:         schema.TypeString,
// 									Optional:     true,
// 									ValidateFunc: validation.IsCIDR,
// 								},
// 								"secondary_addresses": {
// 									Type:     schema.TypeList,
// 									Optional: true,
// 									MaxItems: 4,
// 									Elem: &schema.Schema{
// 										Type:         schema.TypeString,
// 										ValidateFunc: validation.IsCIDR,
// 									},
// 								},
// 								"dhcp_helper": {
// 									Type:         schema.TypeString,
// 									Optional:     true,
// 									ValidateFunc: validation.IsCIDR,
// 								},
// 							},
// 						},
// 					},
// 					"dynamic_ipv4": {
// 						Type:     schema.TypeList,
// 						Optional: true,
// 						Elem: &schema.Resource{
// 							Schema: map[string]*schema.Schema{},
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// // Structure for Data Source starts here

// func generateSystemTemplateSchemaforDS() map[string]*schema.Schema {
// 	return map[string]*schema.Schema{
// 		//For Basic Configuration
// 		"system_basic": {
// 			Type:     schema.TypeList,
// 			Computed: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"site_id": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"system_ip": {
// 						Type:     schema.TypeString,
// 						Computed: true,
// 					},
// 					"overlay_id": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"timezone": {
// 						Type:     schema.TypeString,
// 						Computed: true,
// 					},
// 					"hostname": {
// 						Type:     schema.TypeString,
// 						Computed: true,
// 					},
// 					"location": {
// 						Type:     schema.TypeString,
// 						Computed: true,
// 					},
// 					"device_groups": {
// 						Type: schema.TypeList,
// 						Elem: &schema.Schema{
// 							Type: schema.TypeString,
// 						},
// 						Computed: true,
// 					},
// 					"controller_groups": {
// 						Type: schema.TypeList,
// 						Elem: &schema.Schema{
// 							Type: schema.TypeString,
// 						},
// 						Computed: true,
// 					},
// 					"description": {
// 						Type:     schema.TypeString,
// 						Computed: true,
// 					},
// 					"console_baud_rate": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"max_omp_sessions": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"tcp_optimization_enabled": {
// 						Type:     schema.TypeBool,
// 						Computed: true,
// 					},
// 				},
// 			},
// 		},

// 		//For GPS Configuration
// 		"system_gps": {
// 			Type:     schema.TypeList,
// 			Computed: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"gps_latitude": {
// 						Type:     schema.TypeFloat,
// 						Computed: true,
// 					},
// 					"gps_longitude": {
// 						Type:     schema.TypeFloat,
// 						Computed: true,
// 					},
// 				},
// 			},
// 		},

// 		//For Tracker configuration
// 		"system_trackers": {
// 			Type:     schema.TypeList,
// 			Computed: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"tracker_name": {
// 						Type:     schema.TypeString,
// 						Computed: true,
// 					},
// 					"tracker_endpoint_ip": {
// 						Type:     schema.TypeString,
// 						Computed: true,
// 					},
// 					"tracker_endpoint_dns_name": {
// 						Type:     schema.TypeString,
// 						Computed: true,
// 					},
// 					"tracker_endpoint_api_url": {
// 						Type:     schema.TypeString,
// 						Computed: true,
// 					},
// 					"tracker_threshold": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"tracker_interval": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"tracker_multiplier": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					// "tracker_type": {
// 					// 	Type:     schema.TypeString,
// 					// 	Computed: true,
// 					// },
// 				},
// 			},
// 		},

// 		//For Advanced Configuration
// 		"system_advanced": {
// 			Type:     schema.TypeList,
// 			Computed: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"control_session_pps": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"system_tunnel_mtu": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"port_hop": {
// 						Type:     schema.TypeBool,
// 						Computed: true,
// 					},
// 					"port_offset": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"dns_cache_timeout": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"track_transport": {
// 						Type:     schema.TypeBool,
// 						Computed: true,
// 					},
// 					"vbond_local": {
// 						Type:     schema.TypeBool,
// 						Computed: true,
// 					},
// 					"vbond_remote": {
// 						Type:     schema.TypeString,
// 						Computed: true,
// 					},
// 					"track_interface_tag": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 						// ValidateFunc: validation.IntBetween(1, 4294967295),
// 					},
// 					"multicast_buffer_percent": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"usb_controller": {
// 						Type:     schema.TypeBool,
// 						Computed: true,
// 					},
// 					"track_default_gateway": {
// 						Type:     schema.TypeBool,
// 						Computed: true,
// 					},
// 					"host_policier_pps": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"icmp_error_pps": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"allow_same_site_tunnels": {
// 						Type:     schema.TypeBool,
// 						Computed: true,
// 					},
// 					"route_consistency_check": {
// 						Type:     schema.TypeBool,
// 						Computed: true,
// 					},
// 					"admin_tech_on_failure": {
// 						Type:     schema.TypeBool,
// 						Computed: true,
// 					},
// 					"idle_timeout": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"eco_friendly_mode": {
// 						Type:     schema.TypeBool,
// 						Computed: true,
// 					},
// 					"on_demand_enable": {
// 						Type:     schema.TypeBool,
// 						Computed: true,
// 					},
// 					"on_demand_idle_timeout": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func generateNTPTemplateSchemaforDS() map[string]*schema.Schema {
// 	return map[string]*schema.Schema{
// 		// For Server Configuration
// 		"servers": {
// 			Type:     schema.TypeList,
// 			Computed: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"hostname": {
// 						Type:     schema.TypeString,
// 						Computed: true,
// 					},
// 					"key": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"vpn": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"version": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"source_interface": {
// 						Type:     schema.TypeString,
// 						Computed: true,
// 					},
// 					"prefer": {
// 						Type:     schema.TypeBool,
// 						Computed: true,
// 					},
// 				},
// 			},
// 		},

// 		// For Authentication Configuration
// 		"authentication_keys": {
// 			Type:     schema.TypeList,
// 			Computed: true,
// 			Elem: &schema.Resource{
// 				Schema: map[string]*schema.Schema{
// 					"id": {
// 						Type:     schema.TypeInt,
// 						Computed: true,
// 					},
// 					"value": {
// 						Type:     schema.TypeString,
// 						Computed: true,
// 					},
// 				},
// 			},
// 		},

// 		"trusted_key": {
// 			Type: schema.TypeList,
// 			Elem: &schema.Schema{
// 				Type: schema.TypeString,
// 			},
// 			Computed: true,
// 		},
// 	}
// }

// func generateVPNTemplateSchemaforDS() map[string]*schema.Schema {
// 	return map[string]*schema.Schema{}
// }

// func generateVPNInterfaceTemplateSchemaforDS() map[string]*schema.Schema {
// 	return map[string]*schema.Schema{}
// }

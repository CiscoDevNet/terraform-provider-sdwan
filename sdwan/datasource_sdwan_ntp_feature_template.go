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

func datasourceSDWANNtpFeatureTemplate() *schema.Resource {
	return &schema.Resource{
		Read: datasourceSDWANNtpFeatureTemplateRead,

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
			"template_min_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"template_definition": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hostname": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"vpn": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"version": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"source_interface": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"prefer": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"authentication": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"value": {
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
								},
							},
						},
						"trusted": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed: true,
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

func datasourceSDWANNtpFeatureTemplateRead(d *schema.ResourceData, m interface{}) error {
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
	if strings.Compare(stripQuotes(cont.S("templateType").String()), "ntp") != 0 {
		return fmt.Errorf("[ERROR] Invalid Template Type")
	}

	setTemplateAttributes(d, cont)
	d.SetId(ftName)

	log.Println("[DEBUG] End of Read Method ", d.Id())
	return nil
}

func setTemplateAttributes(d *schema.ResourceData, cont *container.Container) *schema.ResourceData {

	d.Set("template_name", stripQuotes(cont.S("templateName").String()))

	d.Set("template_description", stripQuotes(cont.S("templateDescription").String()))

	d.Set("device_type", interfaceToStrList(cont.S("deviceType").Data()))

	d.Set("template_min_version", stripQuotes(cont.S("templateMinVersion").String()))

	ftDef := cont.S("templateDefinition")
	FTDefinition := getNTPFTDefinition(ftDef)

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

func getNTPFTDefinition(cont *container.Container) []map[string]interface{} {
	definition := []map[string]interface{}{}

	ntpDef := map[string]interface{}{}

	ntpDef["server"] = getNtpServer(cont)

	ntpDef["authentication"] = getNtpAuth(cont)

	if cont.Exists("keys", "trusted", "vipValue") {
		ntpDef["trusted"] = interfaceToStrList(cont.S("keys", "trusted", "vipValue").Data())
	}

	definition = append(definition, ntpDef)

	return definition
}

func getNtpServer(serverCont *container.Container) []map[string]interface{} {
	server := make([]map[string]interface{}, 0, 1)

	if serverCont.Exists("server", "vipValue") {
		servers := serverCont.S("server", "vipValue")

		for _, val := range servers.Children() {
			temp := make(map[string]interface{})

			temp["hostname"] = stripQuotes(val.S("name", "vipValue").String())

			keyID, err := strconv.Atoi(val.S("key", "vipValue").String())
			if err == nil {
				temp["key"] = keyID
			}

			vpn, err := strconv.Atoi(val.S("vpn", "vipValue").String())
			if err == nil {
				temp["vpn"] = vpn
			}

			version, err := strconv.Atoi(val.S("version", "vipValue").String())
			if err == nil {
				temp["version"] = version
			}

			if val.Exists("source-interface", "vipValue") {
				temp["source_interface"] = stripQuotes(val.S("source-interface", "vipValue").String())
			}

			if value, err := strconv.ParseBool(stripQuotes(val.S("prefer", "vipValue").String())); err == nil {
				temp["prefer"] = value
			}

			server = append(server, temp)
		}
	}

	return server
}

func getNtpAuth(authCont *container.Container) []map[string]interface{} {
	authKeys := []map[string]interface{}{}

	if authCont.Exists("keys", "authentication") {
		auth := authCont.S("keys", "authentication", "vipValue")

		for _, val := range auth.Children() {
			temp := make(map[string]interface{})

			keyID, err := strconv.Atoi(val.S("number", "vipValue").String())
			if err == nil {
				temp["id"] = keyID
			}

			temp["value"] = stripQuotes(val.S("md5", "vipValue").String())

			authKeys = append(authKeys, temp)
		}
	}

	return authKeys
}

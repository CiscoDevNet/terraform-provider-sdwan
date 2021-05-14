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

func resourceSDWANNtpFeatureTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSDWANNtpFeatureTemplateCreate,
		Read:   resourceSDWANNtpFeatureTemplateRead,
		Update: resourceSDWANNtpFeatureTemplateUpdate,
		Delete: resourceSDWANNtpFeatureTemplateDelete,

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
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							MaxItems: 4,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hostname": {
										Type:     schema.TypeString,
										Required: true,
									},
									"key": {
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validation.IntBetween(1, 65535),
									},
									"vpn": {
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validation.IntBetween(0, 65530),
									},
									"version": {
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validation.IntBetween(1, 4),
									},
									"source_interface": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"prefer": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  false,
									},
								},
							},
						},
						"authentication": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validation.IntBetween(1, 65535),
									},
									"value": {
										Type:      schema.TypeString,
										Required:  true,
										Sensitive: true,
									},
								},
							},
						},

						"trusted": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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

func resourceSDWANNtpFeatureTemplateCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Create method!")

	sdwanClient := m.(*client.Client)

	fTemplate := models.FeatureTemplate{}

	fTemplate.TemplateType = "ntp"
	fTemplate.TemplateName = d.Get("template_name").(string)
	fTemplate.TemplateDescription = d.Get("template_description").(string)
	fTemplate.DeviceType = interfaceToStrList(d.Get("device_type"))
	fTemplate.TemplateMinVersion = d.Get("template_min_version").(string)
	fTemplate.FactoryDefault = d.Get("factory_default").(bool)

	ftDefinition := d.Get("template_definition").(*schema.Set).List()
	if def, err := createNtpFTDefinition(ftDefinition); err == nil {
		fTemplate.TemplateDefinition = def
	}

	log.Println("ftDefinition ", fTemplate)

	cont, err := sdwanClient.Save("/dataservice/template/feature", &fTemplate)
	if err != nil {
		return err
	}

	ftID := stripQuotes(cont.S("templateId").String())
	d.Set("template_id", ftID)
	d.SetId(ftID)

	log.Println("[DEBUG] End of Create Method ", d.Id())

	return resourceSDWANNtpFeatureTemplateRead(d, m)
}

func resourceSDWANNtpFeatureTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Update Method ", d.Id())

	sdwanClient := m.(*client.Client)

	ftID := d.Id()

	fTemplate := models.FeatureTemplate{}

	fTemplate.TemplateType = "ntp"
	fTemplate.TemplateName = d.Get("template_name").(string)
	fTemplate.TemplateDescription = d.Get("template_description").(string)
	fTemplate.DeviceType = interfaceToStrList(d.Get("device_type"))
	fTemplate.TemplateMinVersion = d.Get("template_min_version").(string)
	fTemplate.FactoryDefault = d.Get("factory_default").(bool)

	ftDefinition := d.Get("template_definition").(*schema.Set).List()
	if def, err := createNtpFTDefinition(ftDefinition); err == nil {
		fTemplate.TemplateDefinition = def
	}

	ftURL := fmt.Sprintf("/dataservice/template/feature/%s", ftID)
	_, err := sdwanClient.Update(ftURL, &fTemplate)

	if err != nil {
		return err
	}

	d.SetId(ftID)

	log.Println("[DEBUG] End of Update Method ", d.Id())

	return resourceSDWANNtpFeatureTemplateRead(d, m)
}

func resourceSDWANNtpFeatureTemplateRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Begining Read method ", d.Id())

	sdwanClient := m.(*client.Client)

	ntpftid := d.Id()

	cont, err := getFeatureTemplateByID(sdwanClient, ntpftid)

	if err != nil {
		if cont != nil {
			return fmt.Errorf("[ERROR] Data load failed in between!")
		} else {
			return err
		}
	}

	setTemplateAttributes(d, cont)

	log.Println("[DEBUG] End of Read Method ", d.Id())
	return nil
}

func resourceSDWANNtpFeatureTemplateDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Beginning Delete Method ", d.Id())

	sdwanClient := m.(*client.Client)
	ntpftID := d.Id()

	_, err := sdwanClient.Delete(fmt.Sprintf("/dataservice/template/feature/%s", ntpftID))
	if err != nil {
		return fmt.Errorf("[ERROR] Error occured while destroying")
	}

	d.SetId("")

	log.Println("[DEBUG] End of Delete Method!")
	return nil
}
func createNtpFTDefinition(ftDefinitions []interface{}) (map[string]interface{}, error) {
	definition := make(map[string]interface{})

	if ftDefinitions != nil && len(ftDefinitions) > 0 {
		ftDefinition := ftDefinitions[0]
		ftDefMap := ftDefinition.(map[string]interface{})

		//server configuration
		ServerSet := ftDefMap["server"].(*schema.Set).List()

		if len(ServerSet) > 0 {

			servers := make([]interface{}, 0, 4)

			for _, serverMap := range ServerSet {
				server := make(map[string]interface{})
				createNtpServer(server, serverMap.(map[string]interface{}))

				servers = append(servers, server)
			}

			ServerMap := make(map[string]interface{})
			ServerMap["vipType"] = "constant"
			ServerMap["vipObjectType"] = "tree"
			ServerMap["vipValue"] = servers
			ServerMap["vipPrimaryKey"] = []string{
				"name",
			}

			definition["server"] = ServerMap
		}

		//trusted configuration
		keys := make(map[string]interface{})

		temp := 0

		trusted := make(map[string]interface{})
		if ftDefMap["trusted"] != nil && len(interfaceToStrList(ftDefMap["trusted"])) > 0 {
			trusted["vipObjectType"] = "list"
			trusted["vipType"] = "constant"
			trusted["vipValue"] = interfaceToStrList(ftDefMap["trusted"])
			trusted["vipVariableName"] = "trusted_key"
			keys["trusted"] = trusted
			temp = temp + 1
		}

		//authentication configuration
		AuthSet := ftDefMap["authentication"].(*schema.Set).List()

		if len(AuthSet) > 0 {
			auths := []interface{}{}
			AuthMap := make(map[string]interface{})
			for _, authMap := range AuthSet {
				auth := make(map[string]interface{})
				createNtpAuth(auth, authMap.(map[string]interface{}))

				auths = append(auths, auth)
			}

			AuthMap["vipType"] = "constant"
			AuthMap["vipObjectType"] = "tree"
			AuthMap["vipValue"] = auths
			AuthMap["vipPrimaryKey"] = []string{
				"number",
			}

			keys["authentication"] = AuthMap
			temp = temp + 1
		}

		if temp != 0 {
			definition["keys"] = keys
		}

	}

	return definition, nil
}

func createNtpServer(defMap map[string]interface{}, input map[string]interface{}) {

	if input["hostname"] != nil && input["hostname"] != "" {
		hostname := make(map[string]interface{})
		hostname["vipObjectType"] = "object"
		hostname["vipType"] = "constant"
		hostname["vipValue"] = input["hostname"].(string)
		hostname["vipVariableName"] = "ntp_server_host"
		defMap["name"] = hostname
	}

	if input["key"] != nil {
		key := make(map[string]interface{})
		key["vipObjectType"] = "object"
		key["vipType"] = "constant"
		key["vipValue"] = input["key"].(int)
		key["vipVariableName"] = "ntp_server_server_auth_key"
		defMap["key"] = key

	}

	if input["vpn"] != nil {
		vpn := make(map[string]interface{})
		vpn["vipObjectType"] = "object"
		vpn["vipType"] = "constant"
		vpn["vipValue"] = input["vpn"].(int)
		vpn["vipVariableName"] = "ntp_server_vpn"
		defMap["vpn"] = vpn

	}

	if input["version"] != nil {
		version := make(map[string]interface{})
		version["vipObjectType"] = "object"
		version["vipType"] = "constant"
		version["vipValue"] = input["version"].(int)
		version["vipVariableName"] = "ntp_server_version"
		defMap["version"] = version

	}

	if input["source_interface"] != nil && input["source_interface"] != "" {
		sourceInterface := make(map[string]interface{})
		sourceInterface["vipObjectType"] = "object"
		sourceInterface["vipType"] = "constant"
		sourceInterface["vipValue"] = input["source_interface"].(string)
		sourceInterface["vipVariableName"] = "ntp_server_source_interface"
		defMap["source-interface"] = sourceInterface

	}

	if input["prefer"] != nil {
		prefer := make(map[string]interface{})
		prefer["vipObjectType"] = "object"
		prefer["vipType"] = "constant"
		prefer["vipValue"] = strconv.FormatBool(input["prefer"].(bool))
		prefer["vipVariableName"] = "ntp_server_prefer"
		defMap["prefer"] = prefer

	}

	defMap["priority-order"] = []string{
		"name",
		"key",
		"vpn",
		"version",
		"source-interface",
		"prefer",
	}
}

func createNtpAuth(defMap map[string]interface{}, input map[string]interface{}) {

	if input["id"] != nil {
		id := make(map[string]interface{})
		id["vipObjectType"] = "object"
		id["vipType"] = "constant"
		id["vipValue"] = input["id"].(int)
		id["vipVariableName"] = "ntp_auth_auth_key"
		defMap["number"] = id
	}

	if input["value"] != nil {
		value := make(map[string]interface{})
		value["vipObjectType"] = "object"
		value["vipType"] = "constant"
		value["vipValue"] = input["value"].(string)
		value["vipVariableName"] = "ntp_auth_md5"
		defMap["md5"] = value
	}

	defMap["priority-order"] = []string{
		"number",
		"md5",
	}
}

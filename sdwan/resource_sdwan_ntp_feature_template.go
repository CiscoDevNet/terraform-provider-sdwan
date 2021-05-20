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
						"vedge-C8500-12X4QC",
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
						"vedge-ISRv",
					}, false),
				},
			},
			"template_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice(
					[]string{
						"ntp",
						"cisco_ntp",
					}, false),
			},
			"template_min_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"template_definition": {
				Type:     schema.TypeSet,
				Required: true,
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
						"master": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"stratum": {
										Type:         schema.TypeInt,
										Optional:     true,
										Computed:     true,
										ValidateFunc: validation.IntBetween(1, 15),
									},
									"source": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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

	fTemplate.TemplateType = d.Get("template_type").(string)
	fTemplate.TemplateName = d.Get("template_name").(string)
	fTemplate.TemplateDescription = d.Get("template_description").(string)
	fTemplate.DeviceType = interfaceToStrList(d.Get("device_type"))
	fTemplate.TemplateMinVersion = d.Get("template_min_version").(string)
	fTemplate.FactoryDefault = d.Get("factory_default").(bool)

	check := validateDeviceType(fTemplate.DeviceType, fTemplate.TemplateType)
	if !check {
		return fmt.Errorf("[ERROR] Device Type is not competible with Template type!")
	}

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

	fTemplate.TemplateType = d.Get("template_type").(string)
	fTemplate.TemplateName = d.Get("template_name").(string)
	fTemplate.TemplateDescription = d.Get("template_description").(string)
	fTemplate.DeviceType = interfaceToStrList(d.Get("device_type"))
	fTemplate.TemplateMinVersion = d.Get("template_min_version").(string)
	fTemplate.FactoryDefault = d.Get("factory_default").(bool)

	check := validateDeviceType(fTemplate.DeviceType, fTemplate.TemplateType)
	if !check {
		return fmt.Errorf("[ERROR] Device Type is not competible with Template type!")
	}

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

		//master configuration
		MasterSet := ftDefMap["master"].(*schema.Set).List()

		if len(MasterSet) > 0 {
			masterMap := MasterSet[0].(map[string]interface{})
			createNtpMaster(definition, masterMap)
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

func createNtpMaster(defMap map[string]interface{}, input map[string]interface{}) {
	master := make(map[string]interface{})

	if input["enable"] != nil {
		enable := make(map[string]interface{})
		enable["vipObjectType"] = "object"
		enable["vipType"] = "constant"
		enable["vipValue"] = strconv.FormatBool(input["enable"].(bool))
		enable["vipVariableName"] = "enable"
		master["enable"] = enable
	}

	if input["stratum"] != nil {
		stratum := make(map[string]interface{})
		stratum["vipObjectType"] = "object"
		stratum["vipType"] = "constant"
		stratum["vipValue"] = input["stratum"].(int)
		stratum["vipVariableName"] = "stratum"
		master["stratum"] = stratum
	}

	if input["source"] != nil && input["source"] != "" {
		source := make(map[string]interface{})
		source["vipObjectType"] = "object"
		source["vipType"] = "constant"
		source["vipValue"] = input["source"].(string)
		source["vipVariableName"] = "source"
		master["source"] = source
	}

	defMap["master"] = master
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

func validateDeviceType(deviceType []string, templateType string) bool {

	if templateType == "ntp" {
		var check bool
		for _, device := range deviceType {
			check = belongsToNTP(device)
			if !check {
				return false
			}
		}
		return true
	}

	if templateType == "cisco_ntp" {
		var check bool
		for _, device := range deviceType {
			check = belongsToCiscoNTP(device)
			if !check {
				return false
			}
		}
		return true
	}

	return false
}

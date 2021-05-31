package sdwan

import (
	"fmt"
	"hash/crc32"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func NameValidation() schema.SchemaValidateFunc {
	return validation.All(
		validation.StringLenBetween(1, 128),
		validation.StringIsNotWhiteSpace,
		validation.StringMatch(regexp.MustCompile("^[^<>!&\"@% ]*$"), "must be alphanumeric!"),
		// <, >, !, &, ", or white space are not allowed
	)
}

func DescValidation() schema.SchemaValidateFunc {
	return validation.All(
		validation.StringLenBetween(1, 2048),
		validation.StringIsNotWhiteSpace,
	)
}

func DHCPHelperValidation() schema.SchemaValidateFunc {
	return validation.Any(
		validation.IsIPv4Address,
		validation.IsIPv6Address,
	)
}

func LoadIntervalValidation() schema.SchemaValidateFunc {
	return validation.All(
		validation.IntBetween(30, 600),
		validation.IntDivisibleBy(30),
	)
}

var (
	validDeviceTypesforDeviceType = []string{
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
		"vedge-ASR-1001-HX",
		"vedge-ASR-1001-X",
		"vedge-ASR-1002-HX",
		"vedge-ASR-1002-X",
		"vedge-C1101-4P",
		"vedge-C1101-4PLTEP",
		"vedge-C1101-4PLTEPW",
		"vedge-C1109-2PLTEGB",
		"vedge-C1109-2PLTEUS",
		"vedge-C1109-2PLTEVZ",
		"vedge-C1109-4PLTE2PW",
		"vedge-C1109-4PLTE2P",
		"vedge-C1111-4P",
		"vedge-C1111-4PW",
		"vedge-C1111-4PLTELA",
		"vedge-C1111-8PLTEEA",
		"vedge-C1111-8PW",
		"vedge-C1111-8P",
		"vedge-C1111-8PLTELAW",
		"vedge-C1111-8PLTEEAW",
		"vedge-C1111-8PLTELA",
		"vedge-C1111X-8P",
		"vedge-C1113-8PLTEEAW",
		"vedge-C1113-8PLTELA",
		"vedge-C1113-8PMLTEEA",
		"vedge-C1116-4P",
		"vedge-C1116-4PLTEEA",
		"vedge-C1117-4P",
		"vedge-C1117-4PLTEEA",
		"vedge-C1117-4PLTELA",
		"vedge-C1117-4PM",
		"vedge-C1117-4PMLTEEA",
		"vedge-C1121-4P",
		"vedge-C1121-4PLTEP",
		"vedge-C1121-8P",
		"vedge-C1121-8PLTEP",
		"vedge-C1121-8PLTEPW",
		"vedge-C1121X-8P",
		"vedge-C1121X-8PLTEP",
		"vedge-C1121X-8PLTEPW",
		"vedge-C1126-8PLTEP",
		"vedge-C1126X-8PLTEP",
		"vedge-C1127-8PLTEP",
		"vedge-C1127-8PMLTEP",
		"vedge-C1127X-8PLTEP",
		"vedge-C1127X-8PMLTEP",
		"vedge-C1128-8PLTEP",
		"vedge-C1161-8P",
		"vedge-C1161X-8P",
		"vedge-C1161X-8PLTEP",
		"vedge-ISR-4221",
		"vedge-ISR-4221X",
		"vedge-ISR-4321",
		"vedge-ISR-4331",
		"vedge-ISR-4351",
		"vedge-ISR-4431",
		"vedge-ISR-4451-X",
		"vedge-ISR-4461",
		"vedge-ISRv",
		"vedge-CSR-1000v"}
	validDeviceRole = []string{
		"sdwan-edge",
		"service-node"}
	validConfigTypes = []string{
		"template",
		"file"}
	validTemplateTypes = []string{
		"system-vedge",
		"cisco_system",
		"ntp",
		"cisco_ntp",
		"vpn-vedge",
		"cisco_vpn",
		"vpn-vedge-interface",
		"cisco_vpn_interface"}

	sysDeviceTypes = map[string]bool{
		"vbond":                  true,
		"vedge-100":              true,
		"vedge-100-B":            true,
		"vedge-100-M":            true,
		"vedge-100-WM":           true,
		"vedge-1000":             true,
		"vedge-2000":             true,
		"vedge-5000":             true,
		"vedge-cloud":            true,
		"vedge-ISR1100-4G":       true,
		"vedge-ISR1100-4GLTE":    true,
		"vedge-ISR1100-6G":       true,
		"vedge-ISR1100X-4G":      true,
		"vedge-ISR1100X-6G":      true,
		"vedge-nfvis-CSP-5444":   true,
		"vedge-nfvis-CSP-5456":   true,
		"vedge-nfvis-CSP2100":    true,
		"vedge-nfvis-CSP2100-X1": true,
		"vedge-nfvis-CSP2100-X2": true,
		"vedge-nfvis-UCSC-E":     true,
		"vedge-nfvis-UCSC-M5":    true,
	}

	ciscoSysDeviceTypes = map[string]bool{
		"vedge-C8500-12X4QC":       true,
		"vedge-C1111-4PLTEEA":      true,
		"vedge-C1161-8P":           true,
		"vedge-C1117-4PLTEEAW":     true,
		"vedge-C1121X-8P":          true,
		"vedge-C8200-1N-4T":        true,
		"vedge-ISR-4331":           true,
		"vedge-ISRv":               true,
		"cellular-gateway-CG522-E": true,
		"vedge-C1127X-8PMLTEP":     true,
		"vedge-C1117-4PMLTEEAWE":   true,
		"vedge-ISR-4451-X":         true,
		"vedge-C1113-8PLTEEA":      true,
		"vedge-IR-1821":            true,
		"vedge-ASR-1001-X":         true,
		"vedge-ISR-4321":           true,
		"vedge-C1116-4PLTEEAWE":    true,
		"vedge-C1109-4PLTE2P":      true,
		"vedge-C1121-8P":           true,
		"vedge-ASR-1002-HX":        true,
		"cellular-gateway-CG418-E": true,
		"vedge-C1111-8PLTEEAW":     true,
		"vedge-C1112-8PWE":         true,
		"vedge-C1101-4PLTEP":       true,
		"vedge-ISR1100-4GLTENA-XE": true,
		"vedge-C1111-8PLTELA":      true,
		"vedge-IR-1835":            true,
		"vedge-C1121X-8PLTEP":      true,
		"vedge-IR-1833":            true,
		"vedge-C8300-1N1S-4T2X":    true,
		"vedge-C1121-4P":           true,
		"vedge-ISR-4351":           true,
		"vedge-C1117-4PLTELA":      true,
		"vedge-C1116-4PWE":         true,
		"vedge-nfvis-C8200-UCPE":   true,
		"vedge-C1113-8PM":          true,
		"vedge-IR-1831":            true,
		"vedge-C1127-8PLTEP":       true,
		"vedge-C1121-8PLTEPW":      true,
		"vedge-C1113-8PW":          true,
		"vedge-ASR-1001-HX":        true,
		"vedge-C1128-8PLTEP":       true,
		"vedge-C1113-8PLTEEAW":     true,
		"vedge-C1117-4PW":          true,
		"vedge-C1116-4P":           true,
		"vedge-C1113-8PMLTEEA":     true,
		"vedge-C1112-8P":           true,
		"vedge-ISR-4461":           true,
		"vedge-C1116-4PLTEEA":      true,
		"vedge-ISR-4221":           true,
		"vedge-C1117-4PM":          true,
		"vedge-C1113-8PLTELAWZ":    true,
		"vedge-C1117-4PMWE":        true,
		"vedge-C1109-2PLTEVZ":      true,
		"vedge-C1113-8P":           true,
		"vedge-C1117-4P":           true,
		"vedge-nfvis-ENCS5400":     true,
		"vedge-C8300-2N2S-6T":      true,
		"vedge-C1127-8PMLTEP":      true,
		"vedge-ESR-6300":           true,
		"vedge-ISR-4221X":          true,
		"vedge-ISR1100-4GLTEGB-XE": true,
		"vedge-C8500-12X":          true,
		"vedge-C1109-2PLTEGB":      true,
		"vedge-CSR-1000v":          true,
		"vedge-C1113-8PLTEW":       true,
		"vedge-C1121X-8PLTEPW":     true,
		"vedge-ISR1100-6G-XE":      true,
		"vedge-C1121-4PLTEP":       true,
		"vedge-C1111-8PLTEEA":      true,
		"vedge-C1117-4PLTEEA":      true,
		"vedge-C1127X-8PLTEP":      true,
		"vedge-C1109-2PLTEUS":      true,
		"vedge-C1112-8PLTEEAWE":    true,
		"vedge-C1161X-8P":          true,
		"vedge-C8500L-8S4X":        true,
		"vedge-C1111-8PW":          true,
		"vedge-C1161X-8PLTEP":      true,
		"vedge-C1101-4PLTEPW":      true,
		"vedge-ISR1100X-4G-XE":     true,
		"vedge-IR-1101":            true,
		"vedge-C1111-4P":           true,
		"vedge-C1111-4PW":          true,
		"vedge-C1111-8P":           true,
		"vedge-C1117-4PMLTEEA":     true,
		"vedge-C1113-8PLTELA":      true,
		"vedge-C1111-8PLTELAW":     true,
		"vedge-C1161-8PLTEP":       true,
		"vedge-ISR1100X-6G-XE":     true,
		"vedge-ISR-4431":           true,
		"vedge-C1101-4P":           true,
		"vedge-C1109-4PLTE2PW":     true,
		"vedge-C1113-8PMWE":        true,
		"vedge-C1118-8P":           true,
		"vedge-C1126-8PLTEP":       true,
		"vedge-C8300-1N1S-6T":      true,
		"vedge-C1121-8PLTEP":       true,
		"vedge-C8300-2N2S-4T2X":    true,
		"vedge-C1112-8PLTEEA":      true,
		"vedge-C1111-4PLTELA":      true,
		"vedge-ASR-1002-X":         true,
		"vedge-C1111X-8P":          true,
		"vedge-C1126X-8PLTEP":      true,
		"vedge-ASR-1006-X":         true,
		"vedge-C8000V":             true,
		"vedge-ISR1100-4G-XE":      true,
		"vedge-C1117-4PLTELAWZ":    true,
	}
)

func stripQuotes(word string) string {
	if strings.HasPrefix(word, "\"") && strings.HasSuffix(word, "\"") {
		return strings.TrimSuffix(strings.TrimPrefix(word, "\""), "\"")
	}
	return word
}

func interfaceToStrList(data interface{}) []string {
	values := data.([]interface{})

	strList := make([]string, 0, 1)
	for _, val := range values {
		strList = append(strList, val.(string))
	}

	return strList
}

func interfaceToIntList(data interface{}) []int {
	values := data.([]interface{})

	strList := make([]string, 0, 1)
	for _, val := range values {
		strList = append(strList, val.(string))
	}

	intList := make([]int, 0, 1)
	for _, val := range strList {
		if value, err := strconv.Atoi(val); err == nil {
			intList = append(intList, value)
		}
	}

	return intList
}

func interfaceToFloatList(data interface{}) []float64 {
	values := data.([]interface{})

	strList := make([]string, 0, 1)
	for _, val := range values {
		strList = append(strList, val.(string))
	}

	floatList := make([]float64, 0, 1)
	for _, val := range strList {
		if value, err := strconv.ParseFloat(val, 64); err == nil {
			floatList = append(floatList, value)
		}
	}

	return floatList
}

func intInterfaceToStrList(data interface{}) []string {
	values := data.([]interface{})

	intList := make([]int, 0, 1)
	for _, val := range values {
		intList = append(intList, val.(int))
	}

	strList := make([]string, 0, 1)
	for _, val := range intList {
		i := fmt.Sprint(val)
		strList = append(strList, i)
	}

	return strList
}

func floatInterfaceToStrList(data interface{}) []string {
	values := data.([]interface{})

	floatList := make([]float64, 0, 1)
	for _, val := range values {
		floatList = append(floatList, val.(float64))
	}

	strList := make([]string, 0, 1)
	for _, val := range floatList {
		i := fmt.Sprint(val)
		strList = append(strList, i)
	}

	return strList
}

func createVIPObject(vot interface{}, vt interface{}, vv interface{}, vvn interface{}, vpk interface{}) map[string]interface{} {
	object := make(map[string]interface{})

	if vot != nil {
		object["vipObjectType"] = vot.(string)
	}
	if vt != nil {
		object["vipType"] = vt.(string)
	}
	if vv != nil {
		switch vv.(type) {
		case string:
			object["vipValue"] = vv.(string)
		case bool:
			object["vipValue"] = strconv.FormatBool(vv.(bool))
		case int:
			object["vipValue"] = vv.(int)
		case float32:
			object["vipValue"] = vv.(float32)
		case float64:
			object["vipValue"] = vv.(float64)
		case []interface{}:
			log.Println("Case Interface!")
			object["vipValue"] = vv
			log.Println("Value: ", vv)
		default:
			log.Println("Case default!")
			object["vipValue"] = vv
			return nil
		}
	}
	if vvn != nil {
		object["vipVariableName"] = vvn.(string)
	}
	if vpk != nil {
		object["vipPrimaryKey"] = interfaceToStrList(vpk)
	}
	return object
}

func belongsToNTP(lookup string) bool {
	switch lookup {
	case
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
		"vedge-100-M",
		"vedge-ISR1100X-6G",
		"vedge-ISR1100X-4G":
		return true
	}
	return false
}

func belongsToCiscoNTP(lookup string) bool {
	switch lookup {
	case
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
		"vedge-ISRv":
		return true
	}
	return false
}

func isStringInRange(min, max int) schema.SchemaValidateDiagFunc {
	return func(v interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics
		k, err := strconv.Atoi(v.(string))

		if err != nil {
			diags = append(diags, diag.Errorf("expected integer value, got: %v", v)...)
			return diags
		}

		if min > k || k > max {
			diags = append(diags, diag.Errorf("expected to be in range %v-%v, got: %v", min, max, v)...)
		}

		return diags
	}
}

func IPv6SubNetValidate() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))(\\/((1(1[0-9]|2[0-8]))|([0-9][0-9])|([0-9])))?$"), "Enter valid IPv6 prefix")
}

func getInt(v interface{}) (int, bool) {
	switch v := v.(type) {
	case float64:
		in := int(v)
		return in, true

	case string:
		if v == "" {
			return 0, false
		}
		in, err := strconv.Atoi(v)
		return in, err == nil

	case int:
		return v, true

	default:
		return 0, false
	}

}

var (
	vpnDeviceTypes = map[string]bool{
		"vedge-1000":          true,
		"vedge-2000":          true,
		"vedge-cloud":         true,
		"vedge-5000":          true,
		"vedge-ISR1100-6G":    true,
		"vedge-100-B":         true,
		"vedge-ISR1100-4G":    true,
		"vedge-100":           true,
		"vedge-ISR1100-4GLTE": true,
		"vedge-100-WM":        true,
		"vbond":               true,
		"vedge-100-M":         true,
		"vedge-ISR1100X-6G":   true,
		"vedge-ISR1100X-4G":   true,
	}
	ciscoVPNDeviceTypes = map[string]bool{
		"vedge-C8500-12X4QC":       true,
		"vedge-C1111-4PLTEEA":      true,
		"vedge-C1161-8P":           true,
		"vedge-C1117-4PLTEEAW":     true,
		"vedge-C1121X-8P":          true,
		"vedge-C8200-1N-4T":        true,
		"vedge-ISR-4331":           true,
		"vedge-ISRv":               true,
		"vedge-C1127X-8PMLTEP":     true,
		"vedge-C1117-4PMLTEEAWE":   true,
		"vedge-ISR-4451-X":         true,
		"vedge-C1113-8PLTEEA":      true,
		"vedge-IR-1821":            true,
		"vedge-ASR-1001-X":         true,
		"vedge-ISR-4321":           true,
		"vedge-C1116-4PLTEEAWE":    true,
		"vedge-C1109-4PLTE2P":      true,
		"vedge-C1121-8P":           true,
		"vedge-ASR-1002-HX":        true,
		"vedge-C1111-8PLTEEAW":     true,
		"vedge-C1112-8PWE":         true,
		"vedge-C1101-4PLTEP":       true,
		"vedge-ISR1100-4GLTENA-XE": true,
		"vedge-C1111-8PLTELA":      true,
		"vedge-IR-1835":            true,
		"vedge-C1121X-8PLTEP":      true,
		"vedge-IR-1833":            true,
		"vedge-C8300-1N1S-4T2X":    true,
		"vedge-C1121-4P":           true,
		"vedge-ISR-4351":           true,
		"vedge-C1117-4PLTELA":      true,
		"vedge-C1116-4PWE":         true,
		"vedge-nfvis-C8200-UCPE":   true,
		"vedge-C1113-8PM":          true,
		"vedge-IR-1831":            true,
		"vedge-C1127-8PLTEP":       true,
		"vedge-C1121-8PLTEPW":      true,
		"vedge-C1113-8PW":          true,
		"vedge-ASR-1001-HX":        true,
		"vedge-C1128-8PLTEP":       true,
		"vedge-C1113-8PLTEEAW":     true,
		"vedge-C1117-4PW":          true,
		"vedge-C1116-4P":           true,
		"vedge-C1113-8PMLTEEA":     true,
		"vedge-C1112-8P":           true,
		"vedge-ISR-4461":           true,
		"vedge-C1116-4PLTEEA":      true,
		"vedge-ISR-4221":           true,
		"vedge-C1117-4PM":          true,
		"vedge-C1113-8PLTELAWZ":    true,
		"vedge-C1117-4PMWE":        true,
		"vedge-C1109-2PLTEVZ":      true,
		"vedge-C1113-8P":           true,
		"vedge-C1117-4P":           true,
		"vedge-nfvis-ENCS5400":     true,
		"vedge-C8300-2N2S-6T":      true,
		"vedge-C1127-8PMLTEP":      true,
		"vedge-ESR-6300":           true,
		"vedge-ISR-4221X":          true,
		"vedge-ISR1100-4GLTEGB-XE": true,
		"vedge-C8500-12X":          true,
		"vedge-C1109-2PLTEGB":      true,
		"vedge-CSR-1000v":          true,
		"vedge-C1113-8PLTEW":       true,
		"vedge-C1121X-8PLTEPW":     true,
		"vedge-ISR1100-6G-XE":      true,
		"vedge-C1121-4PLTEP":       true,
		"vedge-C1111-8PLTEEA":      true,
		"vedge-C1117-4PLTEEA":      true,
		"vedge-C1127X-8PLTEP":      true,
		"vedge-C1109-2PLTEUS":      true,
		"vedge-C1112-8PLTEEAWE":    true,
		"vedge-C1161X-8P":          true,
		"vedge-C8500L-8S4X":        true,
		"vedge-C1111-8PW":          true,
		"vedge-C1161X-8PLTEP":      true,
		"vedge-C1101-4PLTEPW":      true,
		"vedge-ISR1100X-4G-XE":     true,
		"vedge-IR-1101":            true,
		"vedge-C1111-4P":           true,
		"vedge-C1111-4PW":          true,
		"vedge-C1111-8P":           true,
		"vedge-C1117-4PMLTEEA":     true,
		"vedge-C1113-8PLTELA":      true,
		"vedge-C1111-8PLTELAW":     true,
		"vedge-C1161-8PLTEP":       true,
		"vedge-ISR1100X-6G-XE":     true,
		"vedge-ISR-4431":           true,
		"vedge-C1101-4P":           true,
		"vedge-C1109-4PLTE2PW":     true,
		"vedge-C1113-8PMWE":        true,
		"vedge-C1118-8P":           true,
		"vedge-C1126-8PLTEP":       true,
		"vedge-C8300-1N1S-6T":      true,
		"vedge-C1121-8PLTEP":       true,
		"vedge-C8300-2N2S-4T2X":    true,
		"vedge-C1112-8PLTEEA":      true,
		"vedge-C1111-4PLTELA":      true,
		"vedge-ASR-1002-X":         true,
		"vedge-C1111X-8P":          true,
		"vedge-C1126X-8PLTEP":      true,
		"vedge-ASR-1006-X":         true,
		"vedge-C8000V":             true,
		"vedge-ISR1100-4G-XE":      true,
		"vedge-C1117-4PLTELAWZ":    true,
	}
)

func hashString(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func belongsToVPNVEdgeInterface(lookup string) bool {
	switch lookup {
	case
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
		"vedge-cloud":
		return true
	}
	return false
}

func belongsToCiscoVPNInterface(lookup string) bool {
	switch lookup {
	case
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
		"vedge-C1117-4PLTELAWZ":
		return true
	}
	return false
}

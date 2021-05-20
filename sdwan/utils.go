package sdwan

import (
	"regexp"
	"strings"

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
		// <, >, !, &, ", or white space are not allowed
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
		"vmanage",
		"vsmart"}
	validDeviceRole = []string{
		"sdwan-edge",
		"service-node"}
	validConfigTypes = []string{
		"template",
		"file"}
	validTemplateTypes = []string{
		"system-vedge",
		"ntp",
		"vpn-vedge",
		"vpn-vedge-interface"}
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

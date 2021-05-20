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

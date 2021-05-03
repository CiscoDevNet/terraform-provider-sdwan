package sdwan

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// func containerToMapArray(data *container.Container, keys []string) []map[string]interface{} {
// 	ms := []map[string]interface{}{}

// 	for _, child := range data.Children() {
// 		ms = append(ms, containerToMap(child, keys))
// 	}

// 	return ms
// }

// func containerToMap(data *container.Container, keys []string) map[string]interface{} {
// 	m := map[string]interface{}{}

// 	for _, k := range keys {
// 		camelCase := snakeCaseToCamelCase(k)
// 		if data.Exists(camelCase) && data.S(camelCase).Data() != nil {
// 			m[k] = data.S(camelCase).Data()
// 		} else {
// 			m[k] = nil
// 		}
// 	}

// 	return m
// }

// func containerToMapArrayDeviceTemplate(data *container.Container, keys []string) []map[string]interface{} {
// 	ms := []map[string]interface{}{}

// 	for _, child := range data.Children() {
// 		ms = append(ms, containerToMapDeviceTemplate(child, keys))
// 	}

// 	return ms
// }

// func containerToMapDeviceTemplate(data *container.Container, keys []string) map[string]interface{} {
// 	m := map[string]interface{}{}

// 	for _, k := range keys {
// 		camelCase := snakeCaseToCamelCase(k)
// 		if data.Exists(camelCase) && data.S(camelCase).Data() != nil && camelCase != "subTemplates" {
// 			m[k] = data.S(camelCase).Data()
// 		} else if camelCase == "subTemplates" {
// 			m[k] = containerToMapArray(data.S("subTemplates"), []string{"template_id", "template_type"})
// 		} else {
// 			m[k] = nil
// 		}
// 	}

// 	return m
// }

// func snakeCaseToCamelCase(inputUnderScoreStr string) (camelCase string) {

// 	isToUpper := false

// 	for _, v := range inputUnderScoreStr {
// 		if isToUpper {
// 			camelCase += strings.ToUpper(string(v))
// 			isToUpper = false
// 		} else {
// 			if v == '_' {
// 				isToUpper = true
// 			} else {
// 				camelCase += string(v)
// 			}
// 		}
// 	}
// 	return

// }

func NameValidation() schema.SchemaValidateFunc {
	return validation.All(
		validation.StringLenBetween(1, 128),
		validation.StringMatch(regexp.MustCompile("^[^<>!&\" ]*$"), "must be alphanumeric!"),
		// <, >, !, &, ", or white space are not allowed
	)
}

var (
	validDeviceTypes = []string{
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
		"vedge-nfvis-UCSC-M5"}
	validDeviceRole = []string{
		"sdwan-edge",
		"service-node"}
	validConfigTypes = []string{
		"template",
		"file"}
	validInterfaceTypes = []string{
		"ge0/",
		"ge1/",
		"ge2/",
		"ge3/",
		"10ge0/",
		"10ge1/",
		"10ge2/",
		"10ge3/",
		"GigabitEthernet0",
		"GigabitEthernet0/",
		"GigabitEthernet1/",
		"GigabitEthernet2/",
		"Cellular0/",
		"eth",
		"irb",
		"loopback",
		"Loopback",
		"mgmt0"}
)

func stripQuotes(word string) string {
	if strings.HasPrefix(word, "\"") && strings.HasSuffix(word, "\"") {
		return strings.TrimSuffix(strings.TrimPrefix(word, "\""), "\"")
	}
	return word
}

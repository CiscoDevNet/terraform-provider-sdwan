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
		validation.StringMatch(regexp.MustCompile("^[^<>!&\"@%]*$"), "must be alphanumeric!"),
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

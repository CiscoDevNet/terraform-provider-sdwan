package sdwan

import (
	"github.com/CiscoDevNet/sdwan-go-client/container"
	"strings"
)

func conainerToMapArray(data *container.Container, keys []string) []map[string]interface{} {
	ms := []map[string]interface{}{}

	for _, child := range data.Children() {
		ms = append(ms, containerToMap(child, keys))
	}

	return ms
}

func containerToMap(data *container.Container, keys []string) map[string]interface{} {
	m := map[string]interface{}{}

	for _, k := range keys {
		camelCase := snakeCaseToCamelCase(k)
		if data.Exists(camelCase) && data.S(camelCase).Data() != nil {
			m[k] = data.S(camelCase).Data()
		} else {
			m[k] = nil
		}
	}

	return m
}

func snakeCaseToCamelCase(inputUnderScoreStr string) (camelCase string) {

	isToUpper := false

	for k, v := range inputUnderScoreStr {
		if k == 0 {
			camelCase = strings.ToUpper(string(inputUnderScoreStr[0]))
		} else {
			if isToUpper {
				camelCase += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '_' {
					isToUpper = true
				} else {
					camelCase += string(v)
				}
			}
		}
	}
	return

}

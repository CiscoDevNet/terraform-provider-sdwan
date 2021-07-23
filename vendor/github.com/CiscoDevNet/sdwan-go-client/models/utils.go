package models

func A(data map[string]interface{}, key string, value interface{}) {

	if value != "" {
		data[key] = value
	}

	if value == "{}" {
		data[key] = ""
	}

	if value == nil {
		data[key] = ""
	}
}
{{- if eq .Name "Configuration Group Devices"}}# Expected import identifier with the format: "configuration_group_id,[deviceId1,deviceId2,...]"
terraform import sdwan_{{snakeCase .Name}}.example "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac,[C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B]"
{{- else if eq .Name "Policy Group Devices"}}# Expected import identifier with the format: "policy_group_id,[deviceId1,deviceId2,...]"
terraform import sdwan_{{snakeCase .Name}}.example "f6dd22c8-0b4f-496c-9a0b-6813d1f8b8ac,[C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B]"
{{- else if hasReference .Attributes}}# Expected import identifier with the format: "{{snakeCase .Name}}_id{{range .Attributes}}{{if .Reference}},{{.TfName}}{{end}}{{end}}"
terraform import sdwan_{{snakeCase .Name}}.example "f6b2c44c-693c-4763-b010-895aa3d236bd{{range .Attributes}}{{if .Reference}},{{.Example}}{{end}}{{end}}"
{{- else}}terraform import sdwan_{{snakeCase .Name}}.example "f6b2c44c-693c-4763-b010-895aa3d236bd"{{- end}}
data "sdwan_{{snakeCase .Name}}_profile_parcel" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
{{- range  .Attributes}}
{{- if .Reference}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else if eq .Type "Int64List"}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
}

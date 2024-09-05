data "sdwan_{{getProfileParcelName .}}" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
{{- range  .Attributes}}
{{- if .Reference}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
}

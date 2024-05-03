data "sdwan_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
{{- if or .QueryParam .Reference}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
{{- if not .RemoveId}}
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
{{- end}}
}
resource "sdwan_{{snakeCase .Name}}_feature_template" "example" {
	name = "Example"
	description = "My Example"
	device_types = ["vedge-C8000V"]
{{- range  .Attributes}}
{{- if and (not .ExcludeTest) (not .ExcludeExample)}}
{{- if isNestedListSet .}}
  {{.TfName}} = [
    {
      {{- range  .Attributes}}
      {{- if and (not .ExcludeTest) (not .ExcludeExample)}}
      {{- if isNestedListSet .}}
      {{.TfName}} = [
        {
          {{- range  .Attributes}}
          {{- if and (not .ExcludeTest) (not .ExcludeExample)}}
          {{- if isNestedListSet .}}
          {{.TfName}} = [
            {
              {{- range  .Attributes}}
              {{- if and (not .ExcludeTest) (not .ExcludeExample)}}
              {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
              {{- end}}
              {{- end}}
            }
          ]
          {{- else}}
          {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
          {{- end}}
          {{- end}}
          {{- end}}
        }
      ]
      {{- else}}
      {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
      {{- end}}
      {{- end}}
      {{- end}}
    }
  ]
{{- else}}
  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if isStringListSet .}}["{{.Example}}"]{{else if isInt64ListSet .}}[{{.Example}}]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
{{- end}}
}

resource "sdwan_{{snakeCase .Name}}_feature_template" "example" {
	name = "Example"
	description = "My Example"
	device_types = ["vedge-C8000V"]
{{- range  .Attributes}}
{{- if and (not .ExcludeTest) (not .ExcludeExample)}}
{{- if eq .Type "List"}}
  {{.TfName}} = [
    {
      {{- range  .Attributes}}
      {{- if and (not .ExcludeTest) (not .ExcludeExample)}}
      {{- if eq .Type "List"}}
      {{.TfName}} = [
        {
          {{- range  .Attributes}}
          {{- if and (not .ExcludeTest) (not .ExcludeExample)}}
          {{- if eq .Type "List"}}
          {{.TfName}} = [
            {
              {{- range  .Attributes}}
              {{- if and (not .ExcludeTest) (not .ExcludeExample)}}
              {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
              {{- end}}
              {{- end}}
            }
          ]
          {{- else}}
          {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
          {{- end}}
          {{- end}}
          {{- end}}
        }
      ]
      {{- else}}
      {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
      {{- end}}
      {{- end}}
      {{- end}}
    }
  ]
{{- else}}
  {{.TfName}} = {{if eq .Type "String"}}"{{end}}{{.Example}}{{if eq .Type "String"}}"{{end}}
{{- end}}
{{- end}}
{{- end}}
}

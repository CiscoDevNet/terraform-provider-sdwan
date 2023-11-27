resource "sdwan_{{snakeCase .Name}}_profile_parcel" "example" {
	name = "Example"
	description = "My Example"
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

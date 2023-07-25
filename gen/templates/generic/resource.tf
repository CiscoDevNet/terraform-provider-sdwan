resource "sdwan_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
{{- if and (not .ExcludeTest) (not .ExcludeExample) (not .TfOnly) (not .Value)}}
{{- if or (eq .Type "List") (eq .Type "Set")}}
  {{.TfName}} = [
    {
      {{- range  .Attributes}}
      {{- if and (not .ExcludeTest) (not .ExcludeExample) (not .TfOnly) (not .Value)}}
      {{- if or (eq .Type "List") (eq .Type "Set")}}
        {{.TfName}} = [
          {
          {{- range  .Attributes}}
          {{- if and (not .ExcludeTest) (not .ExcludeExample) (not .TfOnly) (not .Value)}}
          {{- if or (eq .Type "List") (eq .Type "Set")}}
            {{.TfName}} = [
              {
                {{- range  .Attributes}}
                {{- if and (not .ExcludeTest) (not .ExcludeExample) (not .TfOnly) (not .Value)}}
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

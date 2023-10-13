data "sdwan_{{snakeCase .Name}}" "example" {
{{- range  .Attributes}}
{{- if and (not .ExcludeTest) (not .ExcludeExample) (not .TfOnly) (not .Value) (or .QueryParam)}}
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
                  {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else}}{{.Example}}{{end}}
                  {{- end}}
                {{- end}}
                }
                ]
        {{- else}}
          {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else}}{{.Example}}{{end}}
        {{- end}}
      {{- end}}
    {{- end}}
    }
    ]
  {{- else}}
    {{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else}}{{.Example}}{{end}}
  {{- end}}
{{- end}}
{{- end}}
}
]
{{- else}}
{{.TfName}} = {{if eq .Type "String"}}"{{.Example}}"{{else if eq .Type "StringList"}}["{{.Example}}"]{{else}}{{.Example}}{{end}}
{{- end}}
{{- end}}
{{- end}}
{{- if not .RemoveId}}
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
{{- end}}
}

/*
{{ Summary }}

{{range $section := Sections -}}
{{ template "section" $section -}}
{{ end }}

*/
package conv

{{ define "section" -}}
{{ .Name }}

{{ range $group := .Groups -}}
{{ .Description }}
{{ template "group" $group }}
{{- end }}
{{ end }}

{{ define "group" }}
{{- range $assertion := .Assertions }}
{{- if $assertion.Expects }}
{{- $expect := index ($assertion.Expects) 0 }}
	{{ CodeConvFunc $assertion.Interface $expect }}		// {{ CodeValue $expect }}
{{- else }}
	{{ $assertion.Interface }}
{{- end }}
{{ end }}
{{ end }}

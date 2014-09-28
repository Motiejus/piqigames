package templates

import "html/template"

var Details = template.Must(template.New("details").Parse(`<html>
{{ range . }}

<h1>Module '{{ .Module }}'</h1>

<h2>Type Definitions</h2>

{{ range .PiqiTypedef }}

{{ if .Record }}
<h3>Record '{{ .Record.Name }}'</h3>
<ul>
  {{ range .Record.Field }}
  <li>Field '{{ .Name }}' ({{ .Type}})</li>
  {{ end }} {{/* Field */}}
</ul>
{{ else if .Alias }}
<h3>Alias '{{ .Alias.Name }}'</h3>
Type: {{ .Alias.Type }}
{{ else if .Variant }}
<h3>Variant '{{ .Variant.Name }}'</h3>
Not implemented
{{ else if .List }}
<h3>List '{{ .List.Name }}'</h3>
{{ end }}


{{ end }} {{/* PiqiTypedef */}}

<h2>Methods</h2>
{{ range .Function }}
<h3>{{ .Name }}</h3>
<h4>Input</h4>
{{ .Input }}

<h4>Output</h4>
{{ .Output }}

<h4>Error</h4>
{{ if .Error }}
{{ .Error }}
{{ else }}
N/A
{{ end }} {{/* if error */}}

{{ end }} {{/* function */}}
{{ end }} {{/* toplevel PiqiList */}}
</html>`))

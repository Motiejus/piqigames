package templates
var Module = `

<html>
{{ with $piqi := . }}{{/* top-level 'Piqi' elements */}}

<h1>Module : {{ .Module }}</h1>

<h2>Type Definitions</h2>
{{ range .PiqiTypedef }} {{ with $t := . }}

  {{ if .Record }}
    <h3 id="module_{{ $piqi.Module }}_{{ .Record.Name }}">
      Record : {{ .Record.Name }}</h3>
    <ul>
      {{ range .Record.Field }}
      <li>{{ .Name }} ({{ .Type }})</li>
      {{ end }} {{/* Field */}}
    </ul>
  {{ else if .Alias }}
    <h3 id="module_{{ $piqi.Module }}_{{ .Alias.Name }}">
      Alias : {{ .Alias.Name }}</h3>
    {{ if builtin .Alias.Name }}
    Builtin!!!
    {{ end }}
    Type: {{ .Alias.Type }}
  {{ else if .Variant }}
    <h3 id="module_{{ $piqi.Module }}_{{ .Variant.Name }}">
      Variant : {{ .Variant.Name }}</h3>
      <ul>
        {{ range .Variant.Option }}
        <li>Option {{ .Name }} ({{ .Type }})</li>
        {{ end }}{{/* range .Variant.Option */}}
      </ul>
  {{ else if .List }}
  <h3 id="module_{{ $piqi.Module }}_{{ .List.Name }}">
    List : {{ .List.Name }}</h3>
  {{ end }} {{/* else if .List */}}

{{ end }} {{/* range .PiqiTypedef */}}
{{ end }}

<h2>Methods</h2>
{{ range .Function }}
  <h3>Function : {{ .Name }}</h3>
  <h4>Input</h4>
  {{ .Input }}

  <h4>Output</h4>
  {{ .Output }}

  <h4>Error</h4>
  {{ if .Error }}
  {{ .Error }}
  {{ else }}
  N/A
  {{ end }} {{/* if .Error */}}

  {{ end }} {{/* range .Function */}}
{{ end }} {{/* toplevel Piqi */}}
</html>

`

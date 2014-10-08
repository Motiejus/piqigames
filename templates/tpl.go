package templates
var Module = `

<html>
<head>
  <style type="text/css">
    div.doc {
      border: 1px grey solid;
      background-color: #eee;
    }
  </style>
</head>
<body>

{{ range . }} {{/* Iterate over modules */}}
{{ with $piqi := . }}{{/* top-level 'Piqi' elements */}}

<h1>Module : {{ .Module }}</h1>

<h2>Type Definitions</h2>
{{ range .PiqiTypedef }} {{ with $t := . }}

  {{ if .Record }}
    <h3 id="module_{{ $piqi.Module }}_{{ .Record.Name }}">
      Record : {{ .Record.Name }}</h3>
    <ul>
      {{ if .Record.Doc }}
      <div class="doc">{{ .Record.Doc }}</div>
      {{ end }}{{/* if .Record.Doc */}}

      {{ range .Record.Field }}
      <li>{{ .Name }} ({{ hreftype $piqi.Module .Type }})</li>
      {{ end }} {{/* Field */}}
    </ul>
  {{ else if .Alias }}
    <h3 id="module_{{ $piqi.Module }}_{{ .Alias.Name }}">
      Alias : {{ .Alias.Name }}</h3>
    Type: {{ hreftype $piqi.Module .Alias.Type }}
  {{ else if .Variant }}
    <h3 id="module_{{ $piqi.Module }}_{{ .Variant.Name }}">
      Variant : {{ .Variant.Name }}</h3>
      <ul>
        {{ range .Variant.Option }}
        <li>Option {{ .Name }} ({{ hreftype $piqi.Module .Type }})</li>
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
  {{ if .Input }}
  <h4>Input</h4>
    {{ hreftype $piqi.Module .Output }}
  {{ end }}{{/* if .Input */}}

  {{ if .Output }}
  <h4>Output</h4>
    {{ hreftype $piqi.Module .Output }}
  {{ end }}{{/* if .Output */}}

  {{ if .Error }}
    {{ hreftype $piqi.Module .Error }}
  {{ end }}{{/* if .Error */}}

  {{ end }} {{/* range .Function */}}
{{ end }} {{/* toplevel Piqi (in piqiList) */}}

{{ end }} {{/* toplevel range . */}}
</body>
</html>
`

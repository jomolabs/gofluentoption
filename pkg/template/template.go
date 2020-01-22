package template

import (
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/jomolabs/gofluentoption/pkg/parse"
)

var (
	methodTemplate = `package {{ .Package }}
{{ $root := . }}
{{- range $index, $info := .Targets }}
{{- if $root.MakeCreateMethods }}
func New{{ $info.Name }}() {{ $info.ReceiverText }} {
	return {{ $info.CreationText }}{}
}

func New{{ $info.Name }}WithValues({{ toArgumentList $info.Fields }}) {{ $info.ReceiverText }} {
	return {{ $info.CreationText }}{
{{- range $index, $field := $info.Fields }}
		{{ $field.Name }}: {{ $field.LowerName }},
{{- end }}
	}
}
{{- end }}
{{ range $index, $field := $info.Fields }}
func With{{ $field.Name }}({{ $field.LowerName }} {{ $field.Type }}) {{ $info.ReceiverText }} {
    return {{ $info.CreationText }}{
        {{ $field.Name }}: {{ $field.LowerName }},
    }
}

func ({{ $info.Letter }} {{ $info.ReceiverText }}) With{{ $field.Name }}({{ $field.LowerName }} {{ $field.Type }}) {{ $info.ReceiverText }} {
    {{ $info.Letter }}.{{ $field.Name }} = {{ $field.LowerName }}
    return {{ $info.Letter }}
}

{{ end }}
func Merge{{ $info.Name }}({{ $info.Letter }} ...{{ $info.ReceiverText }}) {{ $info.ReceiverText }} {
    root := {{ $info.CreationText }}{}
    for _, item := range {{ $info.Letter }} {
{{- range $index, $field := $info.Fields }}
        root.{{ $field.Name }} = item.{{ $field.Name }}
{{- end }}
    }

    return root
}
{{- end }}
`
)

func toArgumentList(fields []parse.Field) string {
	allTypes := make([]string, len(fields))
	for idx, field := range fields {
		allTypes[idx] = fmt.Sprintf("%s %s", field.LowerName, field.Type)
	}
	return strings.Join(allTypes, ", ")
}

var (
	funcMap = map[string]interface{}{
		"toArgumentList": toArgumentList,
	}
)

func Render(typeInfo *parse.TypeInfo, wr io.Writer) error {
	tpl, err := template.New("").Funcs(funcMap).Parse(methodTemplate)
	if err != nil {
		return fmt.Errorf("error creating new template: %s", err)
	}

	err = tpl.Execute(wr, typeInfo)
	if err != nil {
		return fmt.Errorf("error rendering template: %s", err)
	}

	return nil
}

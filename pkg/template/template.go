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
{{- if $root.Options.MakeCreateMethods }}
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
func With{{ $field.Name }}{{ possiblyAddSuffix $info.Name $root.Options.UseSuffixes $root.Options.UseTypeNameAsSuffix }}({{ $field.LowerName }} {{ $field.Type }}) {{ $info.ReceiverText }} {
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
{{ end }}
`
)

func toArgumentList(fields []parse.Field) string {
	allTypes := make([]string, len(fields))
	for idx, field := range fields {
		allTypes[idx] = fmt.Sprintf("%s %s", field.LowerName, field.Type)
	}
	return strings.Join(allTypes, ", ")
}

func possiblyAddSuffix(typeName string, suffixList string, useTypeNameAsSuffix bool) (string, error) {
	suffixListItems := strings.Split(suffixList, ",")
	for _, suffixListItem := range suffixListItems {
		fields := strings.Split(suffixListItem, ":")
		if len(fields) != 2 {
			return "", fmt.Errorf("badly formatted suffix entry: \"%s\"", suffixListItem)
		}
		if typeName == fields[0] {
			return fields[1], nil
		}
	}

	if useTypeNameAsSuffix {
		return typeName, nil
	}

	return "", nil
}

var (
	funcMap = map[string]interface{}{
		"toArgumentList":    toArgumentList,
		"possiblyAddSuffix": possiblyAddSuffix,
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

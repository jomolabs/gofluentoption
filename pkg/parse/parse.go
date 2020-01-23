package parse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"unicode"

	"github.com/jomolabs/gofluentoption/pkg/options"
)

func isPrivate(name string) bool {
	return unicode.IsLower(rune(name[0]))
}

func isStructureMatch(name string, structures []string) bool {
	if len(structures) == 0 {
		return true
	}

	for _, structure := range structures {
		if name == structure {
			return true
		}
	}

	return false
}

func checkAllowable(name string, structures []string, opts *options.Options) bool {
	if len(structures) == 0 {
		if isPrivate(name) {
			return opts.AllowPrivateStructures
		}
		return true
	}

	return isStructureMatch(name, structures)
}

func resolveType(t ast.Expr) string {
	switch v := t.(type) {
	case *ast.Ident:
		return v.String()
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", resolveType(v.X))
	case *ast.FuncType:
		params := make([]string, len(v.Params.List))
		returns := make([]string, len(v.Results.List))
		returnsStr := ""
		for i, param := range v.Params.List {
			params[i] = resolveType(param.Type)
		}
		for i, ret := range v.Results.List {
			returns[i] = resolveType(ret.Type)
		}
		if len(returns) == 1 {
			returnsStr = fmt.Sprintf(" %s", returns[0])
		} else if len(returns) > 1 {
			returnsStr = fmt.Sprintf(" (%s)", strings.Join(returns, ", "))
		}

		return fmt.Sprintf("func(%s)%s", strings.Join(params, ", "), returnsStr)
	default:
		panic(fmt.Sprintf("don't know how to handle type! %#v", t))
	}
}

func possiblyModify(name, prefix string, modify bool) string {
	if modify {
		return prefix + name
	}
	return name
}

func possiblyPointerize(name string, pointerize bool) string {
	return possiblyModify(name, "*", pointerize)
}

func possiblyAddressize(name string, pointerize bool) string {
	return possiblyModify(name, "&", pointerize)
}

func parseSuppressedFields(suppressionList string) []suppressedField {
	suppressions := strings.Split(suppressionList, ",")
	suppressedFields := make([]suppressedField, len(suppressions))

	for _, suppression := range suppressions {
		structure := ""
		field := ""
		if strings.Contains(suppression, ":") {
			fields := strings.Split(suppression, ":")
			structure = fields[0]
			field = fields[1]
		} else {
			field = suppression
		}

		suppressedFields = append(suppressedFields, suppressedField{
			Structure: structure,
			Field:     field,
		})
	}

	return suppressedFields
}

func isSuppressedField(structureName, fieldName string, suppressionList []suppressedField, allowPrivateFields bool) bool {
	if unicode.IsLower(rune(fieldName[0])) && !allowPrivateFields {
		return true
	}

	for _, suppressionEntry := range suppressionList {
		if fieldName == suppressionEntry.Field && (suppressionEntry.Structure == "" || suppressionEntry.Structure == structureName) {
			return true
		}
	}

	return false
}

func ParseSource(file string, findStructures []string, opts *options.Options) (*TypeInfo, error) {
	fs := token.NewFileSet()
	top, err := parser.ParseFile(fs, file, nil, parser.AllErrors)
	if err != nil {
		return nil, fmt.Errorf("error parsing source file: %s", err)
	}

	typeInfo := &TypeInfo{
		Package: top.Name.Name,
		Targets: make([]Target, 0),
		Options: TypeInfoOptions{
			MakeCreateMethods:   opts.MakeCreateMethods,
			UseTypeNameAsSuffix: opts.UseTypeNameAsSuffix,
			UseSuffixes:         opts.UseSuffixes,
		},
	}
	suppressedFields := parseSuppressedFields(opts.IgnoreFields)

	for _, decl := range top.Decls {
		if d, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range d.Specs {
				if s, ok := spec.(*ast.TypeSpec); ok {
					if t, ok := s.Type.(*ast.StructType); ok {
						name := s.Name.Name
						if checkAllowable(name, findStructures, opts) {
							target := Target{
								Name:         name,
								ReceiverText: possiblyPointerize(name, opts.Pointerize),
								CreationText: possiblyAddressize(name, opts.Pointerize),
								Letter:       strings.ToLower(string(name[0])),
								Fields:       make([]Field, 0),
							}
							for _, f := range t.Fields.List {
								if !isSuppressedField(name, f.Names[0].Name, suppressedFields, opts.AllowPrivateFields) {
									target.Fields = append(target.Fields, Field{
										Type:      resolveType(f.Type),
										Name:      f.Names[0].Name,
										LowerName: fmt.Sprintf("%s%s", strings.ToLower(string(f.Names[0].Name[0])), f.Names[0].Name[1:len(f.Names[0].Name)]),
										Options:   ParseTagIntoOption(f.Tag),
									})
								}
							}
							typeInfo.Targets = append(typeInfo.Targets, target)
						}
					}
				}
			}
		}
	}

	return typeInfo, nil
}

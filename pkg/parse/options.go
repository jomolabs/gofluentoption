package parse

import (
	"go/ast"
	"strings"
)

type Option struct {
	TestData string
}

func ParseTagIntoOption(tag *ast.BasicLit) Option {
	opt := Option{}

	if tag == nil {
		return opt
	}

	for _, tagline := range strings.Split(tag.Value, " ") {
		tagline := strings.ReplaceAll(tagline, "`", "")

		if strings.HasPrefix(tagline, "goption:") {
			keyValue := strings.Split(tagline, ":")
			if len(keyValue) == 2 {
				fields := strings.Split(keyValue[1], ",")
				for _, field := range fields {
					field = strings.TrimPrefix(strings.TrimSuffix(field, "\""), "\"")
					if strings.Contains(field, "=") {
						fieldKeyValue := strings.Split(field, "=")
						if len(fieldKeyValue) == 2 {
							switch fieldKeyValue[0] {
							case "test.data":
								opt.TestData = fieldKeyValue[1]
							}
						}
					}
				}
			}
		}
	}

	return opt
}

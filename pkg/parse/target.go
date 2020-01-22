package parse

type Field struct {
	Type      string
	Name      string
	LowerName string
	Options   Option
}

type Target struct {
	Name         string
	ReceiverText string
	CreationText string
	Letter       string
	Fields       []Field
}

type TypeInfo struct {
	Package           string
	Targets           []Target
	MakeCreateMethods bool
}

type suppressedField struct {
	Structure string
	Field     string
}

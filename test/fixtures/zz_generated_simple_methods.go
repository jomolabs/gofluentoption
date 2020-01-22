package fixtures

func NewSimpleOptions() SimpleOptions {
	return SimpleOptions{}
}

func NewSimpleOptionsWithValues(name string, floatable float64, pointerToString *string, hidden bool, funcPointer func(int) error) SimpleOptions {
	return SimpleOptions{
		Name:            name,
		Floatable:       floatable,
		PointerToString: pointerToString,
		hidden:          hidden,
		FuncPointer:     funcPointer,
	}
}

func WithName(name string) SimpleOptions {
	return SimpleOptions{
		Name: name,
	}
}

func (s SimpleOptions) WithName(name string) SimpleOptions {
	s.Name = name
	return s
}

func WithFloatable(floatable float64) SimpleOptions {
	return SimpleOptions{
		Floatable: floatable,
	}
}

func (s SimpleOptions) WithFloatable(floatable float64) SimpleOptions {
	s.Floatable = floatable
	return s
}

func WithPointerToString(pointerToString *string) SimpleOptions {
	return SimpleOptions{
		PointerToString: pointerToString,
	}
}

func (s SimpleOptions) WithPointerToString(pointerToString *string) SimpleOptions {
	s.PointerToString = pointerToString
	return s
}

func Withhidden(hidden bool) SimpleOptions {
	return SimpleOptions{
		hidden: hidden,
	}
}

func (s SimpleOptions) Withhidden(hidden bool) SimpleOptions {
	s.hidden = hidden
	return s
}

func WithFuncPointer(funcPointer func(int) error) SimpleOptions {
	return SimpleOptions{
		FuncPointer: funcPointer,
	}
}

func (s SimpleOptions) WithFuncPointer(funcPointer func(int) error) SimpleOptions {
	s.FuncPointer = funcPointer
	return s
}

func MergeSimpleOptions(s ...SimpleOptions) SimpleOptions {
	root := SimpleOptions{}
	for _, item := range s {
		root.Name = item.Name
		root.Floatable = item.Floatable
		root.PointerToString = item.PointerToString
		root.hidden = item.hidden
		root.FuncPointer = item.FuncPointer
	}

	return root
}

package fixtures

type SimpleOptions struct {
	Name            string
	Floatable       float64
	PointerToString *string
	hidden          bool
	FuncPointer     func(int) error
}

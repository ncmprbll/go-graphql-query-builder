package qb

import "fmt"

type Field struct {
	fieldName string
	fields    []*Field

	depth uint8
}

func NewField(fieldName string) *Field {
	return &Field{
		fieldName: fieldName,
		fields:    []*Field{},
	}
}

func (f *Field) Fields(fields ...*Field) *Field {
	f.fields = append(f.fields, fields...)

	return f
}

func (f *Field) depthUpdate() {
	for _, field := range f.fields { 
		field.depth = f.depth + 1
		field.depthUpdate()
	}
}

func (f *Field) String() string {
	spaces := f.depth * PRETTY_PRINT_SPACES

	if len(f.fields) == 0 {
		return fmt.Sprintf("%*c%s", spaces, ' ', f.fieldName)
	}

	fields := ""

	for _, v := range f.fields { 
		fields += fmt.Sprintf("%s\n", v.String())
	}

	return fmt.Sprintf("%*c%s {\n%s%*c}", spaces, ' ', f.fieldName, fields, spaces, ' ')
}

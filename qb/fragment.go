package qb

import "fmt"

type Fragment struct {
	name, typ string

	fields []*Field
}

func NewFragment(name, typ string) *Fragment {
	return &Fragment{name: name, typ: typ}
}

func (f *Fragment) Fields(fields ...*Field) *Fragment {
	for _, field := range fields {
		field.depth = 1
		field.depthUpdate()
		f.fields = append(f.fields, field)
	}

	return f
}

func (f *Fragment) InlineToField() *Field {
	return &Field{fieldName: "... on " + f.typ, fields: f.fields}
}

func (f *Fragment) ToField() *Field {
	return NewField("..." + f.name)
}

func (f *Fragment) String() string {
	fields := ""

	for _, field := range f.fields { 
		fields += fmt.Sprintf("%s\n", field.String())
	}

	return fmt.Sprintf("fragment %s on %s {\n%s}", f.name, f.typ, fields)
}

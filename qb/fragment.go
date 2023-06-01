package qb

import "fmt"

type Fragment struct {
	name string
	typ  any

	fields []*Field
}

func NewFragment(name, typ string) *Fragment {
	return &Fragment{name: name, typ: typ}
}

func UseFragment(name string) *Field {
	return &Field{fieldName: "..." + name}
}

func (f *Fragment) Fields(fields ...*Field) *Fragment {
	for _, field := range fields {
		field.depth = 1
		field.depthUpdate()
		f.fields = append(f.fields, field)
	}

	return f
}

func (f *Fragment) String() string {
	fields := ""

	for _, field := range f.fields { 
		fields += fmt.Sprintf("%s\n", field.String())
	}

	return fmt.Sprintf("fragment %s on %s {\n%s}", f.name, f.typ, fields)
}

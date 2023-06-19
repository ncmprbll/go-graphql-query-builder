package qb

import (
	"fmt"
	"strings"
)

type Fragment struct {
	name, typ string

	fields []*Field
}

func NewFragment(name, typ string) *Fragment {
	return &Fragment{name: name, typ: typ}
}

func (f *Fragment) Fields(fields ...*Field) *Fragment {
	f.fields = append(f.fields, fields...)

	return f
}

func (f *Fragment) InlineToField() *Field {
	return &Field{name: "... on " + f.typ, fields: f.fields}
}

func (f *Fragment) ToField() *Field {
	return NewField("..." + f.name)
}

func (f *Fragment) String(spaces int) (string, error) {
	var b strings.Builder

	// Name and type
	fmt.Fprintf(&b, "fragment %s on %s {\n", f.name, f.typ)

	// Fields
	for _, field := range f.fields {
		s, err := field.String(spaces)

		if err != nil {
			return "", err
		}

		fmt.Fprintf(&b, "%s\n", s)
	}

	b.WriteString("}")

	return b.String(), nil
}

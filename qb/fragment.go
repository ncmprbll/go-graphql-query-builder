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
	var b strings.Builder

	// Fields depth reset
	for _, field := range f.fields {
		field.depth = 1
		field.depthUpdate()
	}

	// Name and type
	fmt.Fprintf(&b, "fragment %s on %s {\n", f.name, f.typ)

	// Fields
	for _, field := range f.fields { 
		fmt.Fprintf(&b, "%s\n", field.String())
	}

	b.WriteString("}")

	return b.String()
}

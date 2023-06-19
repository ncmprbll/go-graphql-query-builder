package qb

import (
	"fmt"
	"strings"
)

type Operation struct {
	name string
	typ  int

	fields    []*Field
	vars      []*Var
	fragments []*Fragment
}

func NewOperation(name string, typ int) *Operation {
	return &Operation{
		name:   name,
		typ:    typ,
		fields: []*Field{},
	}
}

func NewQuery(name string) *Operation {
	return NewOperation(name, TYPE_QUERY)
}

func NewMutation(name string) *Operation {
	return NewOperation(name, TYPE_MUTATION)
}

func (o *Operation) Vars(vars ...*Var) *Operation {
	o.vars = append(o.vars, vars...)

	return o
}

func (o *Operation) Fields(fields ...*Field) *Operation {
	o.fields = append(o.fields, fields...)

	return o
}

func (o *Operation) Fragments(fragments ...*Fragment) *Operation {
	o.fragments = append(o.fragments, fragments...)

	return o
}

func (o *Operation) String() (string, error) {
	s, err := o.PrettyString(0)

	if err != nil {
		return "", err
	}

	s = strings.ReplaceAll(s, "\n}", "}")
	s = strings.ReplaceAll(s, "}\n", "} ")
	s = strings.ReplaceAll(s, " }", "}")
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "{ ", "{")
	s = strings.ReplaceAll(s, "  ", " ")

	return s, nil
}

func (o *Operation) PrettyString(spaces int) (string, error) {
	var b strings.Builder

	// Operation type (query, mutation)
	b.WriteString(typeDescriptor[o.typ])
	b.WriteString(" ")

	// Operation name with arguments
	if o.name != "" {
		b.WriteString(o.name)

		if len(o.vars) > 0 {
			b.WriteString("(")

			for i := 0; i < len(o.vars)-1; i++ {
				fmt.Fprintf(&b, "%s, ", o.vars[i].String())
			}

			b.WriteString(o.vars[len(o.vars)-1].String())
			b.WriteString(")")
		}

		b.WriteString(" ")
	}

	// Fields
	if len(o.fields) > 0 {
		b.WriteString("{\n")

		for _, field := range o.fields {
			s, err := field.String(spaces)

			if err != nil {
				return "", err
			}

			fmt.Fprintf(&b, "%s\n", s)
		}

		b.WriteString("}")
	}

	// Fragments
	if len(o.fragments) > 0 {
		b.WriteString("\n\n")

		for i := 0; i < len(o.fragments)-1; i++ {
			s, err := o.fragments[i].String(spaces)

			if err != nil {
				return "", err
			}

			fmt.Fprintf(&b, "%s\n\n", s)
		}

		s, err := o.fragments[len(o.fragments)-1].String(spaces)

		if err != nil {
			return "", err
		}

		b.WriteString(s)
	}

	return b.String(), nil
}

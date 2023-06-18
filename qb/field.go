package qb

import (
	"errors"
	"fmt"
	"strings"
)

type Field struct {
	fieldName string

	args       []*Arg
	directives []*Directive
	fields     []*Field
	alias      string
}

func NewField(fieldName string) *Field {
	return &Field{
		fieldName: fieldName,
		fields:    []*Field{},
	}
}

func (f *Field) Args(args ...*Arg) *Field {
	f.args = append(f.args, args...)

	return f
}

func (f *Field) Fields(fields ...*Field) *Field {
	f.fields = append(f.fields, fields...)

	return f
}

func (f *Field) Alias(alias string) *Field {
	f.alias = alias

	return f
}

func (f *Field) Directives(directives ...*Directive) *Field {
	f.directives = directives

	return f
}

func (f *Field) SkipIf(what string) *Field {
	f.directives = append(f.directives, NewDirective("@skip").Args(NewArg("if", what)))

	return f
}

func (f *Field) IncludeIf(what string) *Field {
	f.directives = append(f.directives, NewDirective("@include").Args(NewArg("if", what)))

	return f
}

func (f *Field) prettyString(spaces, inc int, visited map[*Field]struct{}) (string, error) {
	var b strings.Builder

	if _, ok := visited[f]; ok {
		return "", errors.New("cycle detected")
	}

	visited[f] = struct{}{}

	// Alias
	if f.alias != "" {
		b.WriteString(f.alias)
		b.WriteString(": ")
	}

	// Field name with arguments
	fmt.Fprintf(&b, "%*c%s", spaces, ' ', f.fieldName)

	if len(f.args) > 0 {
		b.WriteString("(")

		for i := 0; i < len(f.args) - 1; i++ {
			fmt.Fprintf(&b, "%s, ", f.args[i].String())
		}
	
		b.WriteString(f.args[len(f.args) -1].String())
		b.WriteString(")")
	}

	b.WriteString(" ")

	// Directives
	for _, v := range f.directives {
		fmt.Fprintf(&b, "%s ", v.String())
	}

	// Fields
	if len(f.fields) != 0 {
		b.WriteString("{\n")

		for _, v := range f.fields {
			s, err := v.prettyString(spaces + inc, inc, visited)

			if err != nil {
				return "", err
			}

			fmt.Fprintf(&b, "%s\n", s)
		}

		fmt.Fprintf(&b, "%*c", spaces, ' ')
		b.WriteString("}")
	}

	return b.String(), nil
}

func (f *Field) String(spaces int) (string, error) {
	visited := map[*Field]struct{}{}

	return f.prettyString(spaces, spaces, visited)
}

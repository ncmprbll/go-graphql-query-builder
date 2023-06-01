package qb

import "fmt"

type Field struct {
	fieldName string

	args []*Arg
	fields    []*Field
	alias string

	depth uint8
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

func (f *Field) depthUpdate() {
	for _, field := range f.fields { 
		field.depth = f.depth + 1
		field.depthUpdate()
	}
}

func (f *Field) String() string {
	spaces := f.depth * PRETTY_PRINT_SPACES

	alias := ""

	if f.alias != "" {
		alias = f.alias + ": "
	}

	args := ""
	lenArgs := len(f.args)

	for i := 0; i < lenArgs - 1; i++ {
		args += fmt.Sprintf("%s, ", f.args[i].String())
	}

	if lenArgs > 0 {
		args = "(" + args + f.args[lenArgs - 1].String() + ")"
	}

	if len(f.fields) == 0 {
		return fmt.Sprintf("%*c%s%s%s", spaces, ' ', alias, f.fieldName, args)
	}

	fields := ""

	for _, v := range f.fields { 
		fields += fmt.Sprintf("%s\n", v.String())
	}

	return fmt.Sprintf("%*c%s%s%s {\n%s%*c}", spaces, ' ', alias, f.fieldName, args, fields, spaces, ' ')
}

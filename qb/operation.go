package qb

import "fmt"

type Operation struct {
	operationType int
	operationName string

	vars []*Var
	fields        []*Field
	fragments     []*Fragment
}

func NewOperation(operationType int, operationName string) *Operation {
	return &Operation{
		operationType: operationType,
		operationName: operationName,
		fields:        []*Field{},
	}
}

func NewQuery(queryName string) *Operation {
	return NewOperation(TYPE_QUERY, queryName)
}

func NewMutation(mutationName string) *Operation {
	return NewOperation(TYPE_MUTATION, mutationName)
}

func (o *Operation) Vars(vars ...*Var) *Operation {
	o.vars = append(o.vars, vars...) 

	return o
}

func (o *Operation) Fields(fields ...*Field) *Operation {
	for _, field := range fields {
		field.depth = 1
		field.depthUpdate()
		o.fields = append(o.fields, field)
	}

	return o
}

func (o *Operation) Fragments(fragments ...*Fragment) *Operation {
	o.fragments = append(o.fragments, fragments...)

	return o
}

func (o *Operation) String() string {
	if len(o.fields) == 0 {
		return o.operationName
	}

	fields := ""

	for _, field := range o.fields { 
		fields += fmt.Sprintf("%s\n", field.String())
	}

	operation := typeDescriptor[o.operationType] + " "
	operationName := o.operationName

	if o.operationType == TYPE_QUERY {
		operation = ""
	}

	vars := ""
	lenVars := len(o.vars)

	for i := 0; i < lenVars - 1; i++ {
		vars += fmt.Sprintf("%s, ", o.vars[i].String())
	}

	if lenVars > 0 {
		vars = "(" + vars + o.vars[lenVars - 1].String() + ")"
	}

	if operationName != "" {
		operationName += vars + " "
	}

	fragments := ""

	if len(o.fragments) != 0 {
		fragments += "\n\n"

		for _, fragment := range o.fragments { 
			fragments += fmt.Sprintf("%s\n\n", fragment.String())
		}

		fragments = fragments[:len(fragments) - 2]
	}

	return fmt.Sprintf("%s%s{\n%s}%s", operation, operationName, fields, fragments)
}

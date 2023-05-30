package qb

import "fmt"

type Operation struct {
	operationType int
	operationName string

	fields        []*Field
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

func (q *Operation) Fields(fields ...*Field) *Operation {
	for _, field := range fields {
		field.depth = 1
		field.depthUpdate()
		q.fields = append(q.fields, field)
	}

	return q
}

func (q *Operation) String() string {
	if len(q.fields) == 0 {
		return q.operationName + ""
	}

	fields := ""

	for _, v := range q.fields { 
		fields += fmt.Sprintf("%s\n", v.String())
	}

	operation := typeDescriptor[q.operationType]

	if q.operationName == "" {
		return fmt.Sprintf("%s {\n%s}", operation, fields)
	}

	return fmt.Sprintf("%s %s {\n%s}", operation, q.operationName, fields)
}

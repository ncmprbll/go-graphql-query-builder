package qb

const (
	TYPE_QUERY = iota
	TYPE_MUTATION

	TYPE_ARG_QUOTES
	TYPE_ARG_NO_QUOTES

	PRETTY_PRINT_SPACES = 2
)

var typeDescriptor = map[int]string{
	TYPE_QUERY:    "query",
	TYPE_MUTATION: "mutation",
}

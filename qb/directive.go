package qb

import (
	"fmt"
	"strings"
)

type Directive struct {
	name string

	args []*Arg
}

func NewDirective(directiveName string) *Directive {
	return &Directive{name: directiveName}
}

func (d *Directive) Args(args ...*Arg) *Directive {
	d.args = append(d.args, args...)

	return d
}

func (d *Directive) String() string {
	var b strings.Builder

	// Name
	b.WriteString(d.name)

	// Arguments
	if len(d.args) > 0 {
		b.WriteString("(")

		for i := 0; i < len(d.args) - 1; i++ {
			fmt.Fprintf(&b, "%s, ", d.args[i].String())
		}

		b.WriteString(d.args[len(d.args) - 1].String())
		b.WriteString(")")
	}

	return b.String()
}

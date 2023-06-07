package qb

import "fmt"

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
	args := ""

	for i := 0; i < len(d.args) - 1; i++ {
		args += fmt.Sprintf("%s, ", d.args[i].String())
	}

	if len(d.args) > 0 {
		args = "(" + args + d.args[len(d.args) - 1].String() + ")"
	}

	return fmt.Sprintf("%s%s", d.name, args)
}

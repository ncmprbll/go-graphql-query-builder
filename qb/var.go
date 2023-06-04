package qb

import "fmt"

type Var struct {
	name string
	typ  any
}

func NewVar(name, typ string) *Var {
	return &Var{name: name, typ: typ}
}

func UseVar(argName, arg string) *Arg {
	return NewArgNQ(argName, "$" + arg)
}

func (a *Var) ToArg(argName string) *Arg {
	return NewArgNQ(argName, "$" + a.name)
}

func (a *Var) String() string {
	return fmt.Sprintf("$%s: %s", a.name, a.typ)
}

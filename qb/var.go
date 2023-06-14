package qb

import (
	"fmt"
	"strings"
)

type Var struct {
	name, typ, df string
}

func NewVar(name, typ string) *Var {
	return &Var{name, typ, ""}
}

func (a *Var) Default(df string) *Var {
	a.df = df

	return a
}

func (a *Var) ToArg(argName string) *Arg {
	return NewArg(argName, a.name)
}

func (a *Var) String() string {
	var b strings.Builder

	// Name and type
	fmt.Fprintf(&b, "%s: %s", a.name, a.typ)

	// Default value
	if a.df != "" {
		b.WriteString(" = ")
		b.WriteString(a.df)
	}

	return b.String()
}

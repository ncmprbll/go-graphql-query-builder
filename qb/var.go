package qb

import "fmt"

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
	return NewArgNQ(argName, a.name)
}

func (a *Var) String() string {
	df := a.df

	if df != "" {
		df = " = " + df
	}

	return fmt.Sprintf("%s: %s%s", a.name, a.typ, df)
}

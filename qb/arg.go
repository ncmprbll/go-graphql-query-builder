package qb

import (
	"fmt"
	"reflect"
)

type Arg struct {
	wrapType int

	name  string
	value any
}

func NewArg(t int, argName string, value any) *Arg {
	return &Arg{t, argName, value}
}

func NewArgQ(argName string, value any) *Arg {
	return NewArg(TYPE_ARG_QUOTES, argName, value)
}

func NewArgNQ(argName string, value any) *Arg {
	return &Arg{TYPE_ARG_NO_QUOTES, argName, value}
}

func (a *Arg) String() string {
	value := ""

	v := reflect.ValueOf(a.value)

	// Rework
	if v.Kind() == reflect.Slice {
		valueLen := v.Len()

		for i := 0; i < valueLen - 1; i++ {
			if a.wrapType == TYPE_ARG_QUOTES {
				value += "\"" + fmt.Sprintf("%v", v.Index(i).Interface()) + "\", "
			} else if a.wrapType == TYPE_ARG_NO_QUOTES {
				value += fmt.Sprintf("%v", v.Index(i).Interface()) + ", "
			}
		}
	
		if valueLen > 0 {
			if a.wrapType == TYPE_ARG_QUOTES {
				value += "\"" + fmt.Sprintf("%v", v.Index(valueLen - 1).Interface()) + "\""
			} else if a.wrapType == TYPE_ARG_NO_QUOTES {
				value += fmt.Sprintf("%v", v.Index(valueLen - 1).Interface())
			}
		}

		value = "[" + value + "]"
	} else {
		if a.wrapType == TYPE_ARG_QUOTES {
			value = "\"" + fmt.Sprintf("%v", a.value) + "\""
		} else if a.wrapType == TYPE_ARG_NO_QUOTES {
			value = fmt.Sprintf("%v", a.value)
		}
	}

	return a.name + ": " + value
}

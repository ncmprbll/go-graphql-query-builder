package qb

import (
	"fmt"
	"reflect"
	"strings"
)

type Arg struct {
	name    string
	value   any
	wrap    bool
	spacing bool
}

func NewArg(name string, value any) *Arg {
	return &Arg{name, value, false, false}
}

func cloneEmptyNameArg(a *Arg, value any) *Arg {
	return &Arg{"", value, a.wrap, a.spacing}
}

func (a *Arg) Wrap() *Arg {
	a.wrap = true

	return a
}

func (a *Arg) Spacing() *Arg {
	a.spacing = true

	return a
}

func (a *Arg) String() string {
	var b strings.Builder
	v := reflect.ValueOf(a.value)

	// Name
	if a.name != "" {
		b.WriteString(a.name)
		b.WriteString(": ")
	}

	switch v.Kind() {
	case reflect.Slice:
		length := v.Len()

		if a.spacing {
			b.WriteString("[ ")
		} else {
			b.WriteByte('[')
		}

		for i := 0; i < length; i++ {
			if a.wrap {
				fmt.Fprintf(&b, "\"%v\"", cloneEmptyNameArg(a, v.Index(i).Interface()).String())
			} else {
				fmt.Fprintf(&b, "%v", cloneEmptyNameArg(a, v.Index(i).Interface()).String())
			}

			if i != length-1 {
				b.WriteString(", ")
			}
		}

		if a.spacing {
			b.WriteString(" ]")
		} else {
			b.WriteByte(']')
		}
	case reflect.Map:
		length := v.Len()

		if a.spacing {
			b.WriteString("{ ")
		} else {
			b.WriteByte('{')
		}

		for i, key := range v.MapKeys() {
			v := v.MapIndex(key)

			if a.wrap {
				fmt.Fprintf(&b, "%v: \"%v\"", key, cloneEmptyNameArg(a, v.Interface()).String())
			} else {
				fmt.Fprintf(&b, "%v: %v", key, cloneEmptyNameArg(a, v.Interface()).String())
			}

			if i != length-1 {
				b.WriteString(", ")
			}
		}

		if a.spacing {
			b.WriteString(" }")
		} else {
			b.WriteByte('}')
		}
	default:
		if a.wrap {
			fmt.Fprintf(&b, "\"%v\"", a.value)
		} else {
			fmt.Fprintf(&b, "%v", a.value)
		}
	}

	return b.String()
}

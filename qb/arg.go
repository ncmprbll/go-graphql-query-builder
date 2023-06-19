package qb

import (
	"fmt"
	"reflect"
	"strings"
)

type Arg struct {
	name  string
	value any
	wrap  bool
}

func NewArg(name string, value any) *Arg {
	return &Arg{name, value, false}
}

func (a *Arg) Wrap() *Arg {
	a.wrap = true

	return a
}

func (a *Arg) String() string {
	var b strings.Builder
	v := reflect.ValueOf(a.value)

	// Name
	b.WriteString(a.name)
	b.WriteString(": ")

	switch v.Kind() {
	case reflect.Slice:
		length := v.Len()

		b.WriteByte('[')

		for i := 0; i < length; i++ {
			if a.wrap {
				fmt.Fprintf(&b, "\"%v\"", v.Index(i).Interface())
			} else {
				fmt.Fprintf(&b, "%v", v.Index(i).Interface())
			}

			if i != length-1 {
				b.WriteString(", ")
			}
		}

		b.WriteByte(']')
	case reflect.Map:
		b.WriteByte('{')

		for _, key := range v.MapKeys() {
			v := v.MapIndex(key)

			if a.wrap {
				fmt.Fprintf(&b, "%v: \"%v\"", key, v)
			} else {
				fmt.Fprintf(&b, "%v: %v", key, v)
			}

			b.WriteString(", ")
		}

		return b.String()[:b.Len()-2] + "}"
	default:
		if a.wrap {
			fmt.Fprintf(&b, "\"%v\"", a.value)
		} else {
			fmt.Fprintf(&b, "%v", a.value)
		}
	}

	return b.String()
}

package ir

import (
	"strings"
	"unicode"
)

func capitalize(s string) string {
	var v []rune
	for i, c := range s {
		if i == 0 {
			c = unicode.ToUpper(c)
		}
		v = append(v, c)
	}
	return string(v)
}

func afterDot(v string) string {
	idx := strings.Index(v, ".")
	if idx > 0 {
		return v[idx+1:]
	}
	return v
}

func (t *Type) EncodeFn() string {
	if t.Is(KindArray) && t.Item.EncodeFn() != "" {
		return t.Item.EncodeFn() + "Array"
	}
	switch t.Primitive {
	case Int, Int64, Int32, String, Bool, Float32, Float64:
		return capitalize(t.Primitive.String())
	case UUID, Time:
		return afterDot(t.Primitive.String())
	default:
		return ""
	}
}

func (t Type) ToString() string {
	if t.EncodeFn() == "" {
		return ""
	}
	return t.EncodeFn() + "ToString"
}

func (t Type) FromString() string {
	if t.EncodeFn() == "" {
		return ""
	}
	return "To" + t.EncodeFn()
}

func (t *Type) IsInteger() bool {
	switch t.Primitive {
	case Int, Int8, Int16, Int32, Int64,
		Uint, Uint8, Uint16, Uint32, Uint64:
		return true
	default:
		return false
	}
}

func (t *Type) IsFloat() bool {
	switch t.Primitive {
	case Float32, Float64:
		return true
	default:
		return false
	}
}

func (t *Type) IsNumeric() bool { return t.IsInteger() || t.IsFloat() }
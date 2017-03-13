package compose

import (
	"bytes"
	"io"
)

// T is the template type.
type T struct{ t ti }

// S creates a template that renders the concatenation of the
// rendering of its arguments.
func S(args ...interface{}) T {
	a := make(tMany, 0, len(args))
	return T{&a}.S(args...)
}

// L creates a template that renders the concatenation of the
// rendering of its arguments, followed by a newline character.
func L(args ...interface{}) T {
	return S(args...).S('\n')
}

// L creates a template that renders the concatenation of the original
// template with the rendering of all the arguments, followed by a
// newline character.
func (t T) L(args ...interface{}) T {
	return t.S(args...).S('\n')
}

// S creates a template that renders the concatenation of the original
// template with the rendering of all the arguments.
func (t T) S(args ...interface{}) T {
	var a tMany
	if ar, ok := t.t.(*tMany); ok {
		a = *ar
	} else {
		a = make(tMany, 0, len(args)+1)
		a = append(a, t.t)
	}
	if len(a) == 0 && len(args) == 1 {
		return T{wrap(args[0])}
	}
	for _, arg := range args {
		a = append(a, wrap(arg))
	}
	t.t = &a
	return t
}

// Render renders the template to the given writer.
func (t T) Render(w io.Writer) {
	t.t.run(w)
}

// RenderString renders the template to a string.
func (t T) RenderString() string {
	var buf bytes.Buffer
	t.t.run(&buf)
	return buf.String()
}

// String returns an internal repreentation of the template to
// troubleshoot the template itself.
func (t T) String() string {
	var buf bytes.Buffer
	t.t.decompose(&buf)
	return buf.String()
}

// P is the type of a template placeholder.
type P string

// I creates a template which renders the concatenation of the
// original template with the placeholder defined by the key argument
// iterated over the remainder arguments.
//
// For example, `S(P("a"),"x").I("a", 3,2,1)` renders to `3x2x1x`.
func (t T) I(key string, args ...interface{}) T {
	iargs := make([]ti, len(args))
	for i, a := range args {
		iargs[i] = wrap(a)
	}

	return T{&tIter{t: t.t, key: key, args: iargs}}
}

// A creates a template which renders the original template with the
// placeholders given at even argument positions set to the values at
// the odd argument positions.
//
// For example, `S(P("b"),P("a")).A("a",12,"b",34")` renders to
// `3412`.
func (t T) A(v ...interface{}) T {
	if len(v)%2 != 0 {
		panic("must have even number of parameters")
	}
	args := Targs{}
	for i := 0; i < len(v); i = i + 2 {
		args[v[i].(string)] = T{wrap(v[i+1])}
	}
	return T{&tApply{tmpl: t.t, args: args}}
}

type Targs map[string]T

// Am is equivalent to A but provides the arguments as a map.
func (t T) Am(args Targs) T {
	return T{&tApply{tmpl: t.t, args: args}}
}

package compose

import (
	"bytes"
	"fmt"
	"io"
	"sort"
)

func (p P) decompose(buf *bytes.Buffer) {
	buf.WriteString("P(")
	buf.WriteString(string(p))
	buf.WriteByte(')')
}

func (p P) run(w io.Writer, env ...Targs) {
	key := string(p)
	for _, scope := range env {
		if t, ok := scope[key]; ok {
			t.t.run(w, env...)
			return
		}
	}
	panic(fmt.Sprintf("error: argument %s not provided", key))
}

type tIter struct {
	t    ti
	key  string
	args []ti
}

func (m *tIter) decompose(buf *bytes.Buffer) {
	m.t.decompose(buf)
	buf.WriteString(".Iter(")
	fmt.Fprintf(buf, "%q", m.key)
	for _, v := range m.args {
		buf.WriteString(", ")
		v.decompose(buf)
	}
	buf.WriteByte(')')
}

func (m *tIter) run(w io.Writer, env ...Targs) {
	largs := Targs{}
	env = append([]Targs{largs}, env...)
	for _, a := range m.args {
		largs[m.key] = T{a}
		m.t.run(w, env...)
	}
}

type tApply struct {
	tmpl ti
	args Targs
}

func (a *tApply) decompose(buf *bytes.Buffer) {
	a.tmpl.decompose(buf)
	buf.WriteString(".A(")
	keys := make([]string, 0, len(a.args))
	for k := range a.args {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(buf, "%q", k)
		buf.WriteString(", ")
		a.args[k].t.decompose(buf)
	}
	buf.WriteByte(')')
}

func (a *tApply) run(w io.Writer, env ...Targs) {
	a.tmpl.run(w, append([]Targs{a.args}, env...)...)
}

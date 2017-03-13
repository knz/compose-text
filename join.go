package compose

import (
	"bytes"
	"io"
)

type tJoin struct {
	delim ti
	a     tMany
}

func (o *tJoin) decompose(buf *bytes.Buffer) {
	buf.WriteString("J(")
	o.delim.decompose(buf)
	for i := range o.a {
		buf.WriteString(", ")
		o.a[i].decompose(buf)
	}
	buf.WriteByte(')')
}

func (o *tJoin) run(w io.Writer, env ...Targs) {
	for i, v := range o.a {
		if i > 0 {
			o.delim.run(w, env...)
		}
		v.run(w, env...)
	}
}

func J(delim interface{}, args ...interface{}) T {
	d := wrap(delim)
	t := S(args...)
	return T{&tJoin{delim: d, a: *t.t.(*tMany)}}
}

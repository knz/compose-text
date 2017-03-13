package compose

import (
	"bytes"
	"io"
)

type tMany []ti

func (a *tMany) decompose(buf *bytes.Buffer) {
	buf.WriteString("S(")
	for i, v := range *a {
		if i > 0 {
			buf.WriteString(", ")
		}
		v.decompose(buf)
	}
	buf.WriteByte(')')
}

func (a *tMany) run(w io.Writer, env ...Targs) {
	for _, o := range *a {
		o.run(w, env...)
	}
}

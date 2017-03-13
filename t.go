package compose

import (
	"bytes"
	"io"
)

func (t T) run(w io.Writer, env ...Targs) {
	t.t.run(w, env...)
}

func (t T) decompose(buf *bytes.Buffer) {
	buf.WriteString("T[")
	t.t.decompose(buf)
	buf.WriteByte(']')
}

type ti interface {
	run(w io.Writer, env ...Targs)
	decompose(buf *bytes.Buffer)
}

var _ ti = T{}
var _ ti = P("")
var _ ti = &tMany{}
var _ ti = &tJoin{}
var _ ti = &tOne{}
var _ ti = &tIter{}
var _ ti = &tApply{}

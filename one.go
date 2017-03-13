package compose

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type tOne struct {
	datum interface{}
}

func wrap(arg interface{}) ti {
	switch t := arg.(type) {
	case T:
		return t.t
	case ti:
		return t
	default:
		return &tOne{datum: arg}
	}
}

func (o *tOne) decompose(buf *bytes.Buffer) {
	switch v := o.datum.(type) {
	case T:
		buf.WriteString("T[")
		v.t.decompose(buf)
		buf.WriteByte(']')
	case ti:
		buf.WriteByte('(')
		v.decompose(buf)
		buf.WriteByte(')')
	case nil:
		buf.WriteString("<nil>")
	default:
		fmt.Fprintf(buf, "%T(%q)", o.datum, o.datum)
	}
}

func (o *tOne) run(w io.Writer, env ...Targs) {
	switch t := o.datum.(type) {
	case T:
		t.t.run(w, env...)
	case ti:
		t.run(w, env...)
	case string:
		w.Write([]byte(t))
	case *string:
		w.Write([]byte(*t))
	case fmt.Stringer:
		w.Write([]byte(t.String()))
	case bool:
		w.Write([]byte(strconv.FormatBool(t)))
	case int:
		w.Write([]byte(strconv.Itoa(t)))
	case int64:
		w.Write([]byte(strconv.FormatInt(t, 10)))
	case rune:
		w.Write([]byte(fmt.Sprintf("%c", t)))
	default:
		panic(fmt.Sprintf("don't know how to run %T(%q)", o.datum, o.datum))
	}
}

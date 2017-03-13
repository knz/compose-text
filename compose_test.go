package compose

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestSimple(t *testing.T) {
	testData := []struct {
		t        T
		expected string
	}{
		{S(""), ``},
		{S("hello"), `hello`},
		{S("hello", "world"), `helloworld`},
		{S("hello").S("world"), `helloworld`},
		{S(S("hello"), "world"), `helloworld`},
		{S(S("hello"), 123, '.'), `hello123.`},
		{S(S("hello"), true, '.'), `hellotrue.`},
		{S(S("hello"), int64(-1), '.'), `hello-1.`},
		{L("hello", "world"), "helloworld\n"},
	}

	for i, test := range testData {
		fmt.Fprintf(os.Stderr, "testing: %s\n", test.t)

		var buf bytes.Buffer
		test.t.run(&buf)
		if buf.String() != test.expected {
			t.Errorf("%d: expected %q, got %q", i, test.expected, buf.String())
		}
	}
}

func TestJoin(t *testing.T) {
	testData := []struct {
		t        T
		expected string
	}{
		{J(S("")), ``},
		{J(S(",")), ``},
		{J(S(""), 'a', "b", 123), `ab123`},
		{J(S(","), 'a', "b", true), `a,b,true`},
		{S("a", "b", J(S(","), "c", "d"), "e"), `abc,de`},
	}

	for i, test := range testData {
		fmt.Fprintf(os.Stderr, "testing: %s\n", test.t)

		var buf bytes.Buffer
		test.t.run(&buf)
		if buf.String() != test.expected {
			t.Errorf("%d: expected %q, got %q", i, test.expected, buf.String())
		}
	}
}

func TestFunctional(t *testing.T) {
	testData := []struct {
		t        T
		expected string
	}{
		{S("").A(), ``},
		{S(P("a")).A("a", 123), `123`},

		{S("hello, ", P("name")).
			A("name", "jane"),
			`hello, jane`},
		{S(L("hello, ", P("name")),
			L("bye, ", P("name"))).
			A("name", "jane"),
			"hello, jane\nbye, jane\n"},

		{J(S(", "), P("greeting"), P("name")).
			A("name", "jane").
			A("greeting", "hello"),
			`hello, jane`},

		{S(P("greeting"), P("name"), "!\n").
			A("greeting", "hello, ").
			I("name", "jane", "joe"),
			"hello, jane!\nhello, joe!\n"},

		{S("").Am(Targs{}), ``},
		{S(P("a")).Am(Targs{"a": S("hello")}), `hello`},
	}
	for i, test := range testData {
		fmt.Fprintf(os.Stderr, "testing: %s\n", test.t)

		var buf bytes.Buffer
		test.t.run(&buf)
		if buf.String() != test.expected {
			t.Errorf("%d: expected %q, got %q", i, test.expected, buf.String())
		}
	}
}

func TestReuse(t *testing.T) {
	v := ""
	tmpl := S("hello, ", &v)
	if tmpl.RenderString() != "hello, " {
		t.Error("nope")
	}
	v = "world"
	if tmpl.RenderString() != "hello, world" {
		t.Error("nope")
	}

	tmpl = S(P("greeting"), P("name"), "!\n").
		A("greeting", "hello, ")

	if tmpl.I("name", "jane", "joe").RenderString() !=
		"hello, jane!\nhello, joe!\n" {
		t.Error("nope")
	}
	if tmpl.I("name", "mark", "laura").RenderString() !=
		"hello, mark!\nhello, laura!\n" {
		t.Error("nope")
	}

}

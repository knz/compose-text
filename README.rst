==============
 compose-text
==============

Go text templating made easy
----------------------------

``compose-text`` is a library that provides a **text templating**
system easier to use than Go's standard ``template`` package.

It was inspired by the "StringTree" generator in `capnp-scala`__.

.. __: https://github.com/danhhz/capnp-scala/blob/master/codegen/src/main/scala/com/capnp/StringTree.scala

How to use
==========

1. ``go get github.com/knz/compose-text``
2. ``import (. "github.com/knz/compose-text")`` in your source.
3. use `S` etc. in your code.

For example:

.. code:: go

   // Template for _ x _ = _ (newline)
   var a, b, c int
   t := L(&a, " x ", &b, " = ", &c)

   // print the multiplication table of 7.
   a = 7
   for b = 1; b < 10; b++ {
     c = a * b
     t.Render(os.Stdout)
   }

This produces the following output::

   7 x 1 = 7
   7 x 2 = 14
   7 x 3 = 21
   ...
   7 x 9 = 63

The following code can also be used instead of using pointers to variables:

.. code:: go

   // Template for _ x _ = _ (newline)
   t := L(P("a"), " x ", P("b"), " = ", P("c"))

   // print the multiplication table of 7.
   for i := 1; i < 10; i++ {
     t.A("a", 7, "b", i, "c", i*7).Render(os.Stdout)
   }

Another example:

.. code:: go

   // Template for a greeting.
   g := L("hello, ", P("name"))

   // Greet multiple people.
   g.I("name", "jane", "mark", "isa").Render(os.Stdout)

Produces::

  hello, jane
  hello, mark
  hello, isa

Supported template operators
============================

``S`` (sequence)
   concatenation of zero or more templates.

``L`` (line)
   same as ``S`` followed by newline.

``P`` (placeholder)
   makes a "hole" template that can be filled by ``A`` or ``I``.

``A`` (apply)
   applies  a template to named arguments (creates a new template).

``I`` (iterate)
   "iterates" a template over a list of arguments (creates a new template
   equivalent to the concatenation of the results)

``J`` (join)
   joins the parts using a delimiter.

Methods on templates
====================

``Render(w io.Writer)``
  Executes the template to the specified writer.

``RenderString() string``
  Executes the template and return the resulting text as string.

``String() string``
  Produces a representation of the structure of the template
  (e.g. for troubleshooting)

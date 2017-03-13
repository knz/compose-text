==============
 compose-text
==============

Go text templating made easy
----------------------------

``compose-text`` is a library that provides a **text templating**
system easier to use than Go's standard ``template`` package.

How to use
==========

1. ``go get github.com/knz/compose-text``
2. ``import (c "github.com/knz/compose-text")`` in your source.
3. use `c.S` etc. in your code.

Example:

.. code:: go

   // Template for _ x _ = _ (newline)
   t := L(P("a"), " x ", P("b"), " = ", P("c"))

   // print the multiplication table of 7.
   for i := 1; i <= 10; i++ {
     t.A("a", 7, "b", i, "c", i*7).Render(os.Stdout)
   }

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

---
title: golang
date: 2020-03-25
tags:
layout: post
---

## Golang Tech Talk

### Performance

Go is pre-compiled, statically typed, and has first class support for
concurrency. For these reasons (and no doubt others), Go programs
*typically* run faster those written in dynamic, scripting languages like
Python, JavaScript, or Ruby.

However, it has garbage collection and modern, high-level language
abstractions (e.g. growable arrays or slices, native unicode support,
runtime-type checks), so Go programs almost certainly run slower than
properly optimized C, C++, or Rust code.

### Safety

Uninitialized variables get a default value.

There are runtime checks that you don't read past the end of an array.

Unused imports or declarations are a compile-time error.

### Hello World

```go
// Example 1
package main

import "fmt"

func main() {
  fmt.Println("Hello, world!")
}
```

* In Go, a "package" is the equivalent of a library or module in other
languages.
* The `import` keyword will give a file access to another package. In Go,
identifiers from another package are always referred to by a namespace, e.g.
`fmt.Println`.
* In Go, program execution starts and ends in package `main`, with the
function `main`. Every executable Go program must have exactly one of each.
* Go does not require semicolons at the end of statements, and `gofmt` will
remove them, except where syntactically necessary.

### Variables

In Go, there are several ways to declare a variable. The following are all
equivalent, and would be compiled to the same machine code:

```go
var x int = 42
var x = 42
x := 42
```

In Go, every variable must be statically typed; that is, it must be declared
with a type, and can only ever hold values of that type. However, if the type
can be inferred from the value a variable is being initialized with, its type
does not have to be explicitly specified. Unlike C, in Go types come after
variable names.

It is possible to declare a variable without initializing it with a value--
in this case, the compiler will initialize it with the variable type's default
value (e.g. for `int`, `0`, for a `String`, `""`).

### Functions

### Pointers

### Structs

Go does not have classes, and thus uses C-style `struct`s to define types:

```go
// Linked List
type LinkedNode struct {
  Value int
  Next *LinkedNode
}

func traverseList(node *LinkedNode) {
  for node.Next != nil {
    fmt.Println(node.Value)
    node = node.Next
  }
}
```

### Correct

No one ever lost their job over code that didn't compile. The most dangerous
code is the code that *looks* right, compiles, seems to "work", but behaves
unexpectedly.

### Links

[Documents at golang.org](https://golang.org/doc/)
[Why Golang is Successful by Rob Pike](https://www.youtube.com/watch?v=cQ7STILAS0M)

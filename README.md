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

### Hello World

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, world!")
}
```

### Links

[Documents at golang.org](https://golang.org/doc/)
[Why Golang is Successful by Rob Pike](https://www.youtube.com/watch?v=cQ7STILAS0M)

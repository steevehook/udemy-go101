# Section 1 - Introduction

## Go's Philosophy & Authors

### About the Authors

Go's authors are a big reason why the language is so successful. So we're
talking about:

- Rob Pike
- Ken Thompson
- Robert Griesemer

These are all gods in the computer science realm. They are authors & coauthors
of many things we use and know to be popular nowadays, things such as:

- UTF8
- UNIX
- B programming language
- Inferno
- Plan 9
- Go (Golang)
- Java HotSpot VM

Go's authors personally for me are a strong argument to choose this
amazing language, which brings me to the question: Why is this
language to be considered?

### About the Project

`"Go is an open source programming language
that makes it easy to build simple, reliable, and efficient software"`

This is what the headline of [Official Website](https://golang.org/) says.

From my experience it really hits well on all 3:

- Simplicity
- Reliability
- Efficiency

Go programming language is also a garbage collected language, meaning
you don't have to do memory management like in C or C++.

Besides the fact that is a statically typed language, which very much
looks like C with couple of improvements. What also makes Go different
and amazing in this manner are its:

- Simple & minimal modular language constructs
- Statically typed (almost feels dynamic)
- Concurrency model (CSP - Communicating Sequential Processes)
- Go Routines (not threads)
- Fast builds & compilation time
- Garbage Collection
- Runtime reflection
- Cross platform compilation without a VM

These are some pretty compelling reasons to want to learn and
master Go. And I'm here to teach you how to do it.

### Hello World

Here's a small Hello World program which proves how easy it is to
get up and running with Go programming language

**Hello World in C**
```cgo
#include<stdio.h>

int main() {
    printf("Hello World\n");
}
```

**Hello World in Go**
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

[Back](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>

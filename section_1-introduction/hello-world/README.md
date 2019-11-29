# Section 1 - Introduction

## Hello World

Here's the link to the **Hello World** [code](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction/hello-world/main.go)

### File components

Every file in Go has at least one or more of these components:

- Package declaration
- Import statements
- Constants
- Variables
- Types
- The `init` function
- The `main` function

Out of all these components, every file in Go only requires
the ***package* declaration** in order to successfully compile.

### Programs in Go

Every Go program has an entry point, and that is the `main` function.
The main function is only one and cannot be redeclared.

Every program in Go is also **UTF8 encoded** before compiled. That means
that you can use any UTF8 character in your programs for naming
identifiers i.e. for functions, variables and so on.

In regards to the `main` func, there is also the `main` package.
The `main` package has a special connotation in Go, and it means,
this is the entry point, where all things in a program start.

The `main` function declared anywhere else or outside the `package main`
will not have the same effect, it will act as a regular function called `main`

---
Note: Trying to execute a file or multiple files which don't belong to
the `main` package will result in an error.

Also trying to execute a file or multiple files which belong to the
`main` package but do not have the `func main` declaration in at least
one of the files will also result in an error.

### Order of evaluation

In every Go file things get evaluated in a very specific order.
Thus the way you write your code or reorder it does not really matter,
it will still evaluate and execute in the same order.

This is the order in which things are evaluate in every Go file:

1. Package declaration
2. Imports (recursively)
3. Constants
4. Variables
5. Types
6. `init`
7. `main`

[Back](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>

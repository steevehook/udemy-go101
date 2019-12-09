# Section 1 - Introduction

## Go environment variables

When it comes to the **Go tool**, there are a bunch of environment
variables which either **get exposed** by the Go tool or the developers are
allowed to **set** them in order to change the behavior
of the Go tool one way or another.

For displaying the entire list of Go environment variables,
make sure to run:

```bash
go env
```

However, the ones I think are the most important and relevant are the following:

### `$GOPATH`

Represents the **Go workspace**, which gets installed the moment you install the Go tool.
For more info make sure to check out the
[`$GOPATH`](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction/gopath) lecture

### `$GOROOT`

`$GOROOT` is usually referred to as the **version** of the Go tool (binary)
which is used to compile (build) the Go programs.

### `$GOBIN`

This is the location where all the **global binaries** live. If `$GOBIN`
is exported to `$PATH` variable, it makes it possible to
run any global generated Go binary from anywhere inside the terminal.

It can be overwritten to use another location.

### `$GOHOSTOS`

Displays the **current operating system type**.
Cannot be overwritten to trick the compiler in any way

### `$GOOS`

If set the compiler will compile the code for the **specified operating system**
and will generate a binary which can only be run on that operating system.

### `$GOHOSTARCH`

Displays the **current computer architecture type**.
Cannot be overwritten to trick the compiler in any way

### `$GOARCH`

If set the compiler will compile the code for the **specified computer architecture**
and will generate a binary which can only be run on that specific architecture.

### `GO111MODULE`

By default it equals to `empty string`, however if set to `on` it
enables Go modules features or Go module aware feature.

We'll talk more about this one later on in future sections

---

Note: When setting `$GOOS` and `$GOARCH` a variety of combinations
can be achieved, so in this way, we can say a Go program can be
**cross platform**, by generating multiple binaries from one single
host machine.

Here's a full list of supported values for
[`$GOOS` & `$GOARCH`](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63) in Go

For detailed info about each environment variable run:

```bash
go help environment
```

For a detailed list about supported operating systems & architectures run:

```bash
go tool dist list
```

[Back](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>

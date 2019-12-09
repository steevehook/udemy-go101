# Section 1 - Introduction

## `$GOPATH` & Go workspace

Right after installing the Go tool, there are some things you need to
be familiar with before compiling or running any program. One of them
is `$GOPATH`.

`$GOPATH` in Go is also known as the main workspace where all the
**source code**, **binaries** and **vendor libraries** live.

With newer versions of the language `$GOPATH` becomes a little
obsolete and needless to use, but it's still supported and probably still
going to be in use for some time, in fact most of the Go projects
still make use of it.

So, long story short, `$GOPATH` consists of 3 main directories:
`src`, `pkg` and `bin`

### `src` directory

The `src` directory is the directory we're mostly going to interact
with. This is the place where *all the source code lives* & older versions
of Go (which do not have Go modules enabled) will also store
*third party code* inside `$GOPATH/src`

The `src` directory pretty much looks like a GitHub url structure:

```
├── $GOPATH/
│   ├── src/
│   │   ├── github.com/
│   │   │   ├── organization/
│   │   │   │   ├── project1/
│   │   │   │   └── project2/
│   │   │   ├── user1/
│   │   │   │   ├── project1/
│   │   │   │   └── project2/
│   │   │   ├── pgk/
│   │   │   └── └── errors/
│   │   ├── gitlab.com/
│   │   │   ├── organization/
│   │   │   │   ├── project1/
│   │   │   │   └── project2/
│   │   │   ├── user1/
│   │   │   │   ├── project1/
│   │   │   │   └── project2/
│   │   │   ├── user2/
│   │   │   │   ├── project1/
│   │   │   └── └── project2/
│   │   ├── golang.org/
│   │   │   ├── x/
│   │   │   │   ├── crypto/
│   │   │   │   ├── lint/
│   └── └── └── └── sys/
└──
```

1. On the top level we have hosts: things like **github.com**, **gitlab.com**
or pretty much any host which has Go source code.

2. Then follows the **organization** or **user**, which is to indicate the belonging
of a certain project (repository)

3. Then follows the **project (repository)** itself, from which we import
the exposed packages.

### `pkg` directory

The `pkg` directory is where modern **(Go modules)** third party code
and also **Go archives** live.

Basically we'll interact less with this directory, because the Go tool
writes to it & manages it for us.

`pkg` is also a directory which can be used by other third party
package managers, for example **Go dep**

So, the `pkg` directory structure looks something like this:

```
├── $GOPATH/
│   ├── pkg/
│   │   ├── darwin_amd64/
│   │   │   ├── github.com/
│   │   │   │   ├── uber-go/
│   │   │   │   │   └── atomic.a
│   │   │   │   ├── user1/
│   │   │   │   │   ├── package1/
│   │   │   └── └── └── package2/
│   │   ├── mod/
│   │   │   ├── github.com/
│   │   │   │   ├── uber-go/
│   │   │   │   │   └── zap@1.8.0
│   │   │   │   │   └── zap@1.9.1
│   │   │   └── └── └── ...
│   │   ├── dep/
│   │   │   ├── sources/
│   │   │   │   ├── https---github.com-pkg-errors/
│   │   │   │   │   ├── errors.go
│   │   │   │   │   ├── errors_test.go
│   │   │   │   │   └── ...
│   │   │   │   ├── https---github.com-user1-project1/
│   │   │   │   │   ├── file1.go
│   │   │   │   │   ├── file1_test.go
│   │   ├── ├── ├── └── ...
│   │   │   │   ├── https---github.com-user2-project1/
│   │   │   │   │   ├── file1.go
│   │   │   │   │   ├── file1_test.go
│   │   │   │   │   ├── file2.go
│   │   │   │   │   ├── file2_test.go
│   │   └── └── └── └── ...
│   └──
└──
```

1. You will have a directory specific for your **architecture** (`amd64`, `i386`) and the operating
system (`windows`, `linux`, `darwin`). In my case `darwin_amd64`, which means this is the directory
which will store archives for **OS X amd64**.

If you're using Windows you should probably have something like `windows_amd64`
or `linux_amd64` in the case of Linux.

These **archive files** basically contain the **compiled package binary** code,
along with debug symbols and source information

It is these files you are referencing when you write `import foo/bar`.
It refers to `$GOROOT/pkg/$GOOS_$GOARCH/foo/bar.a` and not `$GOROOT/src/foo/bar/*.go`

2. Then you should also have a directory called `mod` if you have
installed any dependency using the new **Go modules** feature.

So basically the `mod` directory has the same structure as the `src`
directory except is uses **versions**. So a downloaded library
can have multiple versions under multiple directories as such:

```
...
├── mod/
│   ├── github.com/
│   │   ├── uber-go/
│   │   │   └── zap@1.8.0
│   │   │   └── zap@1.9.1
└── └── └── └── ...
...
```

### `bin` directory

After installing the Go tool, we have configured the environment
variables properly, which means: `$GOBIN`=`$GOPATH/bin`
and we exported it to `$PATH`.

Taking that into consideration, it means any binary that
is located inside `$GOPATH/bin` or `$GOBIN` will globally be available
to run from any place inside the terminal.

Also there is the special command `go install` which we'll cover
later on in this section. Which generates a binary from the current
project and places it inside `$GOBIN` or inside `$GOPATH/bin`.

So nothing fancy here. If you're coming from the *npm* or *composer* world,
you can consider this a **global package**, which is meant to be run
globally (from anywhere inside the terminal).

### Import paths

Import paths in Go may looks weird and a little bit verbose and long.
That is due to the fact they are import relative to `$GOPATH`.

These long paths have already become canonical in Go, so pretty much
any project out there, whether it's using Go modules or plain
`$GOPATH` style will have the same structure:

```go
import "github.com/user/project"
```

`HOST`/`USER`/`PROJECT`

For more info on `$GOPATH` & Go workspace check
[this](https://github.com/steevehook/udemy-go101/blob/master/go-workspace.md) out

[Back](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>

# Section 1 - Introduction

## Go basic commands

As opposed to other programming languages, by simply installing the
Go toolkit, you already have everything you need in order to:

- `Compile` (build) a program
- `Run` a program (compile & run)
- `Test` a program
- `Download` third party packages & libraries
- `Format` the code

and many others

Here are a list of Go basic commands which we'll heavily use
throughout the entire course:

#### `go run`

```bash
# compiles and runs the generated binary from main.go file
go run main.go

# compiles and runs the generated binary from a list of files
go run file1.go file2.go

# compiles and runs the generated binary from all go files in the current directory
go run *.go

# compiles and runs the generated binary from all go files making the specified package
go run github.com/user/repo/main-package

# NOTE: you can only run files which
# belong to the main package
# and have the main function declared

# for more info about run
go help run
```

#### `go build`

```bash
# compiles the main package and generates a binary named after
# the directory name, the main package is located in.
# for example if main.go is located inside $HOME/Desktop/program
# the binary name will be "program"
go build

# compiles all the files that belong to the main package
# and generates the binary with the specified name
go build -o binary-name
go build -o binary-name *.go
go build -o binary-name .

# compiles main.go and generates the binary with the specified name
go build -o binary-name main.go

# compiles a list of files and generates the binary with the specified name
go build -o binary-name f1.go f2.go

# compiles the specified package and generates the binary, naming it as the location directory
go build github.com/user/repo/pkg1

# compiles the specified package and generates the binary with the specified name
go build -o binary-name github.com/user/repo/pkg1

# compiles every main package in the current directory
# and store all the binaries in "bins" directory 
for pkg in $(go list -f '{{if eq .Name "main"}}{{.ImportPath}}{{end}}' ./...);\
go build -o ./bins/$(sed $'s/\\//\\\n/g' <<< $pkg | sed '$!d') $pkg

# for more info about build
go help build
```

#### `go get`

```bash
# downloads the latest version of the dependency and
# stores it inside $GOPATH/src OR inside $GOPATH/pkg/mod IF
# GO111MODULE=on variable is set (MODULE AWARE ACTIVATED)
# USED TO ADD a dependency
go get github.com/julienschmidt/httprouter

# downloads the dependency with 1.3 version and stores it inside $GOPATH/pkg/mod
# GO111MODULE=on variable needs to be set (MODULE AWARE ACTIVATED)
# USED TO ADD a dependency
go get github.com/julienschmidt/httprouter@v1.3

# downloads the dependency commit hash and stores it inside $GOPATH/pkg/mod
# GO111MODULE=on variable needs to be set (MODULE AWARE ACTIVATED)
# USED TO ADD a dependency
go get github.com/julienschmidt/httprouter@8e1132c

# recursively parses your code and downloads
# all imported third party dependencies
# USED TO FETCH all dependencies
go get ./...

# downloads or updates the dependency and stores it correspondingly
# either in $GOPATH/src or $GOPATH/pkg/mod
# USED TO UPDATE a dependency
go get -u github.com/julienschmidt/httprouter

# for more info about get
go help get
```

#### `go install`

```bash
# compiles and generates the binary inside $GOPATH/bin named by directory location
go install

# compiles recursively and generates binaries inside $GOPATH/bin named by directory locations
go install ./...

# compiles and generates "main" binary inside $GOPATH/bin
go install main.go hello.go

# compiles and generates "hello" binary inside $GOPATH/bin
go install hello.go main.go

# compiles and generates "pkg1" and "pkg2" binaries inside $GOPATH/bin
go install github.com/user/repo/pkg1 github.com/user/repo/pkg2

# compiles and generates the binary "shadow" inside $GOPATH/bin from REMOTE location
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow

# ONLY compiles the non-main package and DOES nothing
go install non-main.go
go install fmt

# for more info about install
go help install
```

#### `go test`

```bash
# compiles and tests the entire code inside the current directory
go test ./...

# compiles and tests the entire code inside the current directory (VERBOSE) 
go test -v ./...

# compiles and tests the entire code inside the current directory (NO CACHE)
go test -count=1 ./...

# compiles and tests code in file.go
# Note: the source code must be provided
go test file.go file_test.go

# compiles and tests code in pkg1 and pkg2
go test github.com/user/repo/pkg1 github.com/user/repo/pkg2

# compiles and tests code in pkg which matches the regular expression
go test github.com/user/repo/pkg -run TestRegex

# for more info about test
go help test
```

#### `go fmt`

```bash
# formats all files in the current directory
go fmt ./...

# for more info about fmt
go help fmt
```

#### More info

```bash
# for more info, feel free to run:
go help
```

[Back](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>
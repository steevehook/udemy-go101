# Section 1 - Introduction

## Binary size. Go vs C

As stated earlier in the philosophy of the language, I said
Go is a statically typed language with **Runtime reflection**

So when a piece of Go code is compiled, in result a binary is generated,
which is nothing else then 0s and 1s.

Now, as opposed to C, *in Go* a binary also contains **runtime information**,
the **Garbage Collector** and other useful things, which
are one of the reasons the simplest program in Go weights **more than 2MB**
as opposed to C which has roughly **20KB** for a hello world.

To prove you that, let's compile both programs and check out the result:

```bash
# compiles and generates a binary from the C code
gcc -o c-binary main.c

# compiles and generates a binary from the Go code
go build -o go-binary main.go

# prints the memory for each file inside the current directory
# If you're on Windows, run this with Git Bash
ls -sh c-binary go-binary
```

[Back](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>

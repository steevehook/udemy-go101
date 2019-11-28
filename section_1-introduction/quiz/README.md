# Section 1 Quiz

## Questions

**1)** Can I write Go source code using Chinese or Russian characters?

##### `Answer`: Yes

`Explanation`: Because Go source code is UTF8 encoded

**2)** Is Go a garbage collected language?

`Explanation`: Open up the language [docs](https://golang.org/doc/)
and convince yourself

##### `Answer`: Yes

**3)** Can I call the `init` function?

##### `Answer`: No

`Explanation`: The `init` function cannot be called.
It gets called automatically by the Go runtime.
Calling it will result in a reference error

**4)** What is the only thing a Go file requires?

##### `Answer`: **package** declaration

`Explanation`: Every file in Go must *belong* to a *package*

**5)** Will this piece of code compile?

```go
package main

import "fmt"
```

##### `Answer`: No

`Explanation`: Because the *import* is not used

**6)** Will this piece of code compile?

```go
package main

import "fmt"

func main() {
	fmt.Println(x)
}

const x = 10
```

##### `Answer`: Yes

`Explanation`: Because the order of *constants*, *variables*, *types*
or *functions* does not matter

[Back](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>

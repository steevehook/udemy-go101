# Section 2 - Language Basics (I)

## Overview

In this section we will cover the **Go language basics** and the bare bones
which make Go an amazing language to work with.

We will cover most of the **data types** in Go, some of the **built-in functions**
and **control structures**. We'll explore plenty of examples
to make it as clear as possible.

At the end of this section you'll have a small
[quiz](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/quiz)
which will consolidate the knowledge you gained during this section

You can skip this and the next section if you already have a decent
experience in this area. However if you're completely new to Go
or just want to *refresh your knowledge*, feel free to stick around 🤓

#### Topics:

- [Packages](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/packages)
- [Executable & non-executable packages](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/executable-vs-non-executable-packages)
- [Basic types](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/basic-types)
- [`const` & `var` declarations](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/const-var)
- [Blank identifier (`_`)](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/blank-identifier)
- [Reserved keywords & operators](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/reserved-keywords-and-operators)
- [Aggregate types overview](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/aggregate-types-overview)
- [Arrays](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/arrays)
- [Slices](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/slices)
- [`len` function](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/len-func)
- [`cap` function](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/cap-func)
- [`append` function](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/append-func)
- [`make` function](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/make-func)
- [Arrays/Slices addressing](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/arrays-slices-addressing)
- [Control structures overview](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/control-structures-overview)
- [`if`/`else` statements](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/if-else-statements)
- [`switch` statement](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/switch-statement)
- [`for` loops](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/for-loops)
- [Arrays/Slices looping](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/arrays-slices-looping)
- [Maps](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/maps)
- [Maps - addressing & looping](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/maps-addressing-looping)
- [Maps - comma, ok syntax](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/maps-comma-ok)
- [Nil maps](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/nil-maps)
- [Maps - delete function](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/maps-delete-func)
- [Map implementation](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/map-implementation)
- [Zero values](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/zero-values)
- [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1/quiz)

[Next](https://github.com/steevehook/udemy-go101/blob/master/section_3-language-basics-2) |
[Previous](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction) |
[Back](https://github.com/steevehook/udemy-go101)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>

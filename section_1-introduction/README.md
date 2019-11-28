# Section 1 - Introduction

## About the Author

My name is **Stefan Cirlig** aka **Steve Hook**. I am from the Republic of Moldova
and I'm super excited to share this course with all of you.

I am a self taught software engineer with over 3 years of professional
experience, and I've created `Go 101` to share all my knowledge
and best practices I've learnt while working with Go.

If you're reading this, it means you're at least interested in learning
Go & you don't necessarily have to enroll in the Udemy course because
this repository is very well documented and full of examples & demos.
So if you like reading & experimenting consider this entire course
hosted on GitHub entirely free.

However if you consider this course is worth rewarding, then you
can enroll and I really can't wait to see in the classroom.

This course is full of surprises, so stay tuned and enjoy it 🚀

## Prerequisites

This course is for anyone who want to learn Go from the very beginning
till the very advanced. So no requirements at all.

Having basic knowledge in the following is a plus, but not a requirement:

- Git
- C like language (anything)
- **Time & Passion**

## Course overview

This course is organized into **sections**, **small examples**,
**fully featured applications**, **quizzes** and **assignments**.

All the examples and applications are hosted on GitHub and I also made
sure they perfectly run on the following operating systems:

- `Windows`
- `OSX` (Mac)
- `Ubuntu`
- `Raspbian` (Raspberry Pi)

so without too much worry, got all of you covered.

In the first sections of this course we'll walk over through all the
**language basics**, so if you completely new to Go and don't have any previous
experience I got you covered.

Then we start build **smaller applications**. We're gonna build and test
a **CLI application** & we'll go through some **refactoring** and **improvement** cycles
to make the CLI app more maintainable and idiomatic.

Later on we dive into some heavy meat and we'll spend about 4 big sections
on **designing**, **developing** & **testing** a fully featured **REST API**, which uses
**MariaDB** and **Redis** and also has an **authentication** layer and stateless **application
session**. So if you have some Go experience, you might want to skip the
basics and jump straight into these sections.

At last we will explore more sections related to **Concurrency** in Go
**deploying applications**, **package management**, developing **cross platform** Go code
and even writing a **Real Time Chat** application.

If all of that sounds interesting, well I can't wait for you to find
out what I've prepared for you, because it's hell of a lot more interesting.

## Resources

Resources are everything when it comes to any programming language.
That why in this course you're never alone. We're all part of the same
challenge, and a gopher never leaves a gopher in need.

So besides the course which you can find on Udemy or play with it on
GitHub, there will be issues and you'll want to ask people who've already
been there, or ask me personally.

That's why besides Udemy's Q&A section I also created a designated
**Discord server** for all the participants os this course, whether you
enrolled on Udemy or are simply playing with it on GitHub.

[Udemy Go101 Discord](https://discord.gg/TxuJAs)

Besides that when you're asking a question, be **mindful of your question**.
Before asking a question in Q&A section or on Discord make sure
there is already someone who asked that question, chances are they did.
Because problems students are facing are usually the same.

When you ask a question make sure to be **straight to the point**, desirably
with **practical examples**. Speaking of examples, if you have longer pieces of code
try to wrap them in a [Go Playground](https://play.golang.org/) and **share the link** with the others.

And lastly, **be nice** to everyone in this course. There are people with
different levels of understanding and experience. So if you can help then do it,
if you can't help just be nice in the **community** and let others do it.

##### Other resources

Here's a list of resources which we will very frequently use, so
make sure to at least have a look:

- [Golang.org](https://golang.org/)
- [GoDoc](https://godoc.org/)
- [Go playground](https://play.golang.org/)
- [A Tour of Go](https://tour.golang.org/)
- [Go Tools](https://github.com/golang/tools)
- [Go Lint](https://github.com/golang/lint)
- [Go Imports](https://godoc.org/golang.org/x/tools/cmd/goimports)

## Installation

### Windows

```bash
# Head over to https://golang.org/dl/
# And download the MSI installer

# Follow the installation instructions

# Open a command prompt or type: Windows + R => cmd
# displays the version of Go runtime if installation was successful
go version

# Head over to https://git-scm.com/download/win
# Download Git Bash, which is a linux like terminal for Windows
```

### OSX

```bash
# install Homebrew
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

# Install the latest version of Go runtime
brew install go

# export necessary Go environment variables 
echo -e "\
export GOPATH=\$HOME/go\n
export GOBIN=\$GOPATH/bin\n
export PATH=\$PATH:/\$GOBIN\n" >> .bashrc

# displays the version of Go runtime if installation was successful
go version

# Head over to https://git-scm.com/download/mac
# Download the installer and follow the instructions
# Or try installing it via Homebrew
brew install git
```

### Ubuntu / Raspberry Pi (Linux)

```bash
# cd into HOME directory
cd

# download the Go distribution for ARM architecture
# If there's a newer version of Go make sure to navigate to

# UBUNTU
# https://storage.googleapis.com/golang and choose: linux-386/linux-amd64 version
# Depending if you're on 32/64 bit architecture
wget https://storage.googleapis.com/golang/go1.11.13.linux-amd64.tar.gz

# RASPBERRY PI
# https://storage.googleapis.com/golang and choose: linux-armv6l version
wget https://storage.googleapis.com/golang/go1.11.13.linux-armv6l.tar.gz

# export necessary Go environment variables 
echo -e "\
export GOROOT=/usr/local/go\n
export GOPATH=\$HOME/go\n
export GOBIN=\$GOPATH/bin\n
export PATH=\$PATH:\$GOROOT/bin:\$GOBIN\n" >> .bashrc

# refresh bash configuration
source .bashrc

# displays the version of Go runtime if installation was successful
go version

# Install Git both for Ubuntu/Raspberry Pi
sudo apt-get install git

# displays the version of Git if installation was successful
git version
```

### Text Editors / IDE(s)

There are many text editors and a few IDE(s) that support Go aka Golang
development, however the ones I recommend the most are:

- [VSCode - Free](https://code.visualstudio.com/download)
- [Goland - Paid (30 days Free Trial)](https://www.jetbrains.com/go/download)

Note: Goland IDE is a paid version, however you can get it for free 30 days
or try their [EAP](https://www.jetbrains.com/go/nextversion/) versions

For more info on Go distributions make sure to visit
the [Downloads](https://golang.org/dl/) page

#### Links

- [Go distribution downloads](https://golang.org/dl/)
- [Go runtime versions - XML](https://storage.googleapis.com/golang)
- [Git for Windows](https://git-scm.com/download/win)
- [Git for Mac](https://git-scm.com/download/mac)

#### [Quiz](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction/quiz)

[Back](https://github.com/steevehook/udemy-go101)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>

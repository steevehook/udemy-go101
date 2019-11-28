# udemy-go101

## Go Programming Language 101

## Course overview

- [Section 1 - Introduction](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction)
- [Section 2 - Language Basics (I)](https://github.com/steevehook/udemy-go101/blob/master/section_2-language-basics-1)
- [Section 3 - Language Basics (II)](https://github.com/steevehook/udemy-go101/blob/master/section_3-language-basics-2)
- [Section 4 - Note CLI app](https://github.com/steevehook/udemy-go101/blob/master/section_4-notes-cli-app)
- [Section 5 - Interfaces](https://github.com/steevehook/udemy-go101/blob/master/section_5-interfaces)
- [Section 6 - Exploring `net` & `net/http`](https://github.com/steevehook/udemy-go101/blob/master/section_6-exploring-net-http)
- [Section 7 - Notes CLI app improvements](https://github.com/steevehook/udemy-go101/blob/master/section_7-notes-cli-app-improvements)
- [Section 8 - Errors](https://github.com/steevehook/udemy-go101/blob/master/section_8-errors)
- [Section 9 - Refactor Notes CLI app](https://github.com/steevehook/udemy-go101/blob/master/section_9-refactor-notes-cli-app)
- [Section 10 - Testing](https://github.com/steevehook/udemy-go101/blob/master/section_10-testing)
- [Section 11 - Testing Notes CLI app](https://github.com/steevehook/udemy-go101/blob/master/section_11-testing-notes-cli-app)
- [Section 12 - Expenses REST API](https://github.com/steevehook/udemy-go101/blob/master/section_12-expenses-rest-api)
- [Section 13 - Test Expenses REST API](https://github.com/steevehook/udemy-go101/blob/master/section_13-test-expenses-rest-api)
- [Section 14 - Expenses REST API Auth](https://github.com/steevehook/udemy-go101/blob/master/section_14-expenses-rest-api-auth)
- [Section 15 - Expenses REST API Session](https://github.com/steevehook/udemy-go101/blob/master/section_15-expenses-rest-api-session)
- [Section 16 - Package management](https://github.com/steevehook/udemy-go101/blob/master/section_16-package-management)
- [Section 17 - Concurrency](https://github.com/steevehook/udemy-go101/blob/master/section_17-concurrency)
- [Section 18 - Real Time Chat app](https://github.com/steevehook/udemy-go101/blob/master/section_18-real-time-chat-app)
- [Section 19 - Platform Specific Software](https://github.com/steevehook/udemy-go101/blob/master/section_19-platform-specific-software)
- [Section 20 - Debugging](https://github.com/steevehook/udemy-go101/blob/master/section_20-debugging)
- [Section 21 - Deploying Go apps](https://github.com/steevehook/udemy-go101/blob/master/section_21-deploying-go-apps)
- [Section 22 - Go best practices](https://github.com/steevehook/udemy-go101/blob/master/section_22-go-best-practices)
- [Farewell & Resources](https://github.com/steevehook/udemy-go101/blob/master/farewell-and-resources)

For more info on the course check out the [Curriculum](https://github.com/steevehook/udemy-go101/blob/master/curriculum.md)

### Udemy course

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101-cover.jpg?sanitize=true" width="100%"/>

[Link to Udemy course HERE](https://www.google.com)

4-6 compelling images about the course HERE

### Installation

Before trying any of these examples make sure to have the `Go` binary installed on your platform.

OSX:

```bash
# Install Homebrew

/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

```bash
# Install Go

brew install go
```

Linux:

```bash
# Move into $HOME directory
cd ~

# Download the binary for your distribution
curl -O https://dl.google.com/go/go1.12.linux-amd64.tar.gz

# Verify that the downloaded binary is not corrupted. Check if the hash matches the one from downloads page
sha256sum go1.12.linux-amd64.tar.gz

# Extract the binary
tar xvf go1.12.linux-amd64.tar.gz

# Make root user the owner of Go workspace
sudo chown -R root:root ./go

# Move go directory to a standard location
sudo mv go /usr/local
```

For more details regarding `Go` installation check out

[Golang Installation](https://golang.org/doc/install)

For downloading `Go` binary for your platform checkout

[Golang Downloads](https://golang.org/dl/)

To check if you installed `Go` successfully type:

```bash
# Displays installed Go version
go version

# Displays all environment variables defined by Go
go env
```

Docker:

```bash
docker run -it golang:1.11
```

### Set $GOPATH variable
```bash
sudo nano ~/.profile

# Linux
export GOPATH=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:/$GOBIN

# OSX
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:/$GOBIN

# Save File

# Restart shell configuration
source ~/.profile
source ~/.zshrc
```

### Other Go tools

You can also install other `Go` tools which will help you have a more productive development process

- [Go Tools](https://github.com/golang/tools)
- [Go Lint](https://github.com/golang/lint)
- [Go Imports](https://godoc.org/golang.org/x/tools/cmd/goimports)

#### Go Tools

```bash
go get -u golang.org/x/tools/...
```

#### Go Lint

```bash
go get -u golang.org/x/lint/golint
```

#### Go Imports

```bash
go get -u golang.org/x/tools/cmd/goimports
```

#### Daily routine things

- Format your entire code recursively using
`go fmt ./...`

- Lint your entire code recursively using
`golint ./...` 

### For more info about Go check out

[Go Doc](https://golang.org/pkg)

[Go Packages](https://godoc.org/)

[Go Playground](https://play.golang.org/)

[Effective Go](https://golang.org/doc/effective_go.html)

[A Tour of Go](https://tour.golang.org/)


### For more info about Go commands & WORKSPACE check out

[Go commands explained](https://github.com/steevehook/udemy-go101/blob/master/go-commands.md)

[Go WORKSPACE explained](https://github.com/steevehook/udemy-go101/blob/master/go-workspace.md)

## FEEDBACK ⚗

[GopherTuts TypeForm](https://feedback.gophertuts.com)

## COMMUNITY 🙌

[Udemy Go101 Discord](https://discord.gg/TxuJAs)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>

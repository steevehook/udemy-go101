# Section 1 - Introduction

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

##### Windows ONLY

If you're on Windows make sure to change your default
terminal, so you can replicate the commands I'm running during
the entire course by using **Bash** instead of **cmd**

###### Goland

`Ctrl + Alt + S` => `shell` => `Tools -> Terminal` => `Shell path` => `bash.exe`

###### VSCode

`F1` => `Shell` => `Select Default Shell` => `Bash`

I also created some small useful tasks for VSCode to make your
developing experience as smooth as possible (Goland like)

Make sure to copy `.vscode.sample` inside your project and
rename it to `.vscode`

```bash
# install mockery dependency for working easier with tests
go get github.com/vektra/mockery/.../
```

OR: Go to ==> `Terminal` ==> `Configure Tasks` and paste
the JSON content from `.vscode.sample/tasks.json`

Type: `F6` to run any of the created tasks

---

Make sure to restart Goland or VSCode and check it the changes applied:

- **VSCode**: `Ctrl` + **\`** => `ls` => should work
- **Goland**: `F12` => `ls` => should work

### Other Go tools

You can also install other `Go` tools which will help you have a more productive development process

- [Go Tools](https://github.com/golang/tools)
- [Go Lint](https://github.com/golang/lint)
- [Go Imports](https://godoc.org/golang.org/x/tools/cmd/goimports)

##### Go Tools

```bash
go get -u golang.org/x/tools/...
```

##### Go Lint

```bash
go get -u golang.org/x/lint/golint
```

##### Go Imports

```bash
go get -u golang.org/x/tools/cmd/goimports
```

##### Daily routine

- Format your entire code recursively using
`go fmt ./...`

- Lint your entire code recursively using
`golint ./...` 

#### Links

- [Go distribution downloads](https://golang.org/dl/)
- [Go runtime versions - XML](https://storage.googleapis.com/golang)
- [Git for Windows](https://git-scm.com/download/win)
- [Git for Mac](https://git-scm.com/download/mac)

[Back](https://github.com/steevehook/udemy-go101/blob/master/section_1-introduction)

---

Happy hacking gophers 🚀🚀🚀

<img src="https://github.com/steevehook/udemy-go101/raw/master/udemy-go101.svg?sanitize=true" width="150px"/>

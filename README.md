# goprog

Learning "The Go Programming Language"

# go env

```bash
export GOPATH=${HOME}/golabs
export GOROOT=${HOME}/sdk/go1.17
export GOPROXY=https://goproxy.cn
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

export PS1='\[\e]0;\u@\h: \W\a\]${debian_chroot:+($debian_chroot)}\[\033[01;32m\]\u@\h\[\033[00m\]:\[\033[01;34m\]\W\[\033[00m\]\$ '

set -o ignoreeof

```

# goland settings

1. GOPATH: use GOPATH that's defined in system environment
2. New Project path: $HOME/src/github.com/timscm/goprog/

# final hierarchy

```
golabs/
|-- bin/
|   |__ goimports*
|   |__ hellworld*
|-- pkg/
|   |__ mod/
|   |   |__ cache/
|   |   |__ golang.org/
|   |__ sumdb/
|       |__ sum.golang.org/
|-- src/
    |__ github.com
        |__ timscm
            |__ goprog/
```

# goprog hierarchy

```
golang/
|-- ch1/
|   |-- helloworld/helloworld.go
|   |-- echo1/echo1.go
|-- ch2/
|-- .gitignore
|-- LICENSE
|-- README.md
```

# Running commands

```bash
go mod tidy
go mod vendor
go get rsc.io/quote

go get -u github.com/kardianos/govendor
govendor init
govendor fetch github.com/go-sql-driver/mysql
```
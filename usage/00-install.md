# 00 - Install

## Installation

```go
$ go install golang.org/dl/go1.21.6@latest
$ go1.21.6 download

$ go1.21.6 version
go version go1.21.6 linux/amd64
```

## Use a Go version manager: `G`

For bash:

```
curl -sSL https://git.io/g-install | sh -s -- bash
```

For zsh:

```
curl -sSL https://git.io/g-install | sh -s -- zsh
```

That will download the g script, put it inside `$GOPATH/bin/`, give it execution rights with chmod, and configure your default shell's initialization file, setting the GOPATH & GOROOT environment variables and adding $GOPATH/bin to the PATH.

Install a given version of Go:

```
$ g install 1.21.6
```

Install the latest version of Go:

```
$ g install latest
```

Check Go is correctly installed:

```
$ go version
```

Switch to another version of Go you have previously installed:

```
$ g

    1.18.4
  > 1.21.6
```



go-pr
=====
CLI written in Go to create a PR

Usage:
------

```
NAME:
   go-pr - Create a Pull Request for the current branch

USAGE:

    go-pr @target-branch # default target-user to your Github user
    go-pr target-user/target-repo@target-branch
    go-pr target-user/target-repo # default target-branch to master
    go-pr target-user # default target-branch to master and target-repo to the name of your Github repo


VERSION:
   0.0.1

AUTHOR(S):
   Bao Pham <gbaopham@gmail.com>

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```


Requirements:
-------------
Go, Git

Install:
--------

```bash
go get github.com/baopham/go-pr
```

Make sure you have `$GOPATH` set and `$GOPATH/bin` is in the `$PATH`, e.g.:

```bash
export GOPATH=$HOME/Projects/Go
export PATH=$PATH:/usr/local/opt/go/libexec/bin:$GOPATH/bin
```

License:
--------
MIT

Author:
-------
Bao Pham

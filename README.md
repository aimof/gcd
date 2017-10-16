# gcd
greater cd command for lazy programmer

## Overview

gcd is the greater cd command for lazy programmer.

Now, you no longer need to be confused by search suggestion.
In gcd, you are suggested only when you want.

When you use gcd, you move the directory in $GCDPATH or the Children of $GCDPATH.

## DEMO

```shell
aimof@host:~$ echo $GCDPATH 
/home/aimof/.vim:/home/aimof/go/src/github.com/aimof:/home/aimof
aimof@host:~$ gcd gcd            
aimof@host:~/go/src/github.com/aimof/gcd$ gcd .vim 
aimof@host:~/.vim$ 
```

### with ls

```shell
aimof@host:~$ echo $GCDPATH 
/home/aimof/.vim:/home/aimof/go/src/github.com/aimof:/home/aimof
aimof@host:~$ gcd gcd            
cmd  LICENSE  README.md
aimof@host:~/go/src/github.com/aimof/gcd$ gcd .vim 
autoload  ftplugin  plugged
aimof@host:~/.vim$ 
```

## Install and Settings

First, install 2 bin files named gcdhist and gcdpath.

```shell
go get github.com/aimof/gcd/cmd/gcdhist
go install github.com/aimof/gcd/cmd/gcdhist
go get github.com/aimof/gcd/cmd/gcdpath
go install github.com/aimof/gcd/cmd/gcdpath
```

Second, write below function in your shell setting file such as .bashrc, .zshrc.

```shell
gcd() {
	TARGET=`gcdpath $1`
	gcdhist add $TARGET
	cd $TARGET
	ls $TARGET # if you need to run ls when you move directory.
}
```

Third, set $GCDPATH and $GCDROOT

```
#example
export GCDPATH=$HOME/.vim:$GOPATH/src/github.com/aimof:$HOME
export GCDROOT=$HOME/.gcd
```

Now, you can use gcd command.

### Use history and Search

If you want to use history and search, you can setting these.

To use gcd history, write below function.
Off course, you must install fzf, peco or something.

```shell
gcds() {
	TARGET=$(gcdhist latest | fzf) # If you want, you can use peco or something like that.
	gcdhist add $TARGET
	cd $TARGET
	ls $TARGET # If you want.
}
```

In upper function, your suggestion order is latest.
If you want to be suggested order by frequency, write this.

```shell
gcds() {
	TARGET=$(gcdhist frequent | fzf) # If you want, you can use peco or something like that.
	gcdhist add $TARGET
	cd $TARGET
	ls $TARGET # If you want.
}
```

## the Priority of cd

1. If the argument is correct relative path from current directory, absolute path or ".", "..", "-", gcd behave same as cd.
2. Else, gcd decide path from $GCDPATH. The priority is below.
    1. The first directory in $GCDPATH.
    2. The children of the first directory in $GCDPATH.
    3. The second directory in $GCDPATH.
    4. The children of the second directory in $GCDPATH.
    5. The third directory in $GCDPATH.
    6. \.\.\.

## LICENCE

MIT

## Enjoy gcd.
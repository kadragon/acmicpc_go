#!/bin/sh

a=$1
b=1000
c=10000

d=`expr $a / $b`
d=`expr $d \* 1000`

if (( $d < $c ))
then
    d="0$d"
else
    d="$d"
fi

if (( $1 < $c ))
then
    e="0$1"
else
    e="$1"
fi

mkdir -p ./"$d"/"$e"

cp ./main.go ./"$d"/"$e"/main.go

# touch ./"$d"/"$e"/input.txt
## input content parsing
go run ./parser/main.go $1

## vscode --reuse-window
code -r ./"$d"/"$e"/input.txt
code -r ./"$d"/"$e"/main.go

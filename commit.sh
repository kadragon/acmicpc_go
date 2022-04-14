#!/bin/sh

if (( $1 < 10000 ))
then
    e="0$1"
else
    e="$1"
fi

git add .
git commit -m "$e solve"

#!/bin/env bash

PWD=$(pwd)
for d in $(find $PWD -maxdepth 1 -type d)
do
    [ -e "$d/.generated" ] && rm -r $d
done

cd _generator
go run main.go
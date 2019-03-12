#!/usr/bin/env bash

if [ -z "$1" ]
then
    echo "please input a folder name"
    exit 1
fi

if [ -d "$1" ]
then
    echo "folder $1 exist"
    exit 1
fi

mkdir "$1"
cp ./template.go ./$1/solution.go
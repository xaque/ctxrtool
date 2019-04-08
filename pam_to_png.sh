#!/usr/bin/env bash
for file in "$@"
do
    go run convertCTXR.go $file
    pamtopng $file.pam >${file%.ctxr}.png
    echo Wrote ${file%.ctxr}.png
    rm $file.pam
done

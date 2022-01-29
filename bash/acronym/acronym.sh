#!/usr/bin/env bash

sentence=${1//[^a-zA-Z\']/ }
for word in $sentence; do
    echo -n ${word:0:1} | tr '[:lower:]' '[:upper:]'
done
#!/usr/bin/env bash

source .env

if [[ "$1" == "" ]]; then
  go run main.go
else
  go run examples/"$1"
fi
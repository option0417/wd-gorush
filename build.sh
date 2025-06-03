#!/usr/bin/env bash

platform=$1
EXE_NAME=mpush

if [ -z "$platform" ]; then
  echo "build linux version as default."
  GOOS=linux GOARCH=amd64 go build -o $EXE_NAME
fi


if [ "$platform" == "linux" ]; then
  echo "build Linux version."
  GOOS=linux GOARCH=amd64 go build -o $EXE_NAME
fi

if [ "$platform" == "mac" ]; then
  echo "build Mac version."
  GOOS=darwin GOARCH=amd64 go build -o $EXE_NAME
fi

if [ "$platform" == "windows" ]; then
  echo "build Windows version."
  GOOS=windows GOARCH=amd64 go build -o $EXE_NAME
fi

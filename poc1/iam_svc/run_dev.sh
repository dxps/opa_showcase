#!/bin/sh

## --------------------------------------------------------------------
## This script restarts the main on code changes using the reflex tool.
## If you don't have it installed, do:
## - being outside of this (any any other nowadays Go Modules based) project directory, run:
##   go get github.com/cespare/reflex 
## - and have $HOME/go (or whether your GOPATH is defined) in your PATH
## --------------------------------------------------------------------

reflex -r '\.go' -s -- sh -c "go run ./cmd/main.go"

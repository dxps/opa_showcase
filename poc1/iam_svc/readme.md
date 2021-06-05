## The IAM Service

This is the Identity and Access Management (IAM) service, playing the authentication server role, thus being responsible with authenticating users and returing JWTs containing claims with relevant details.

### Run

Use the standard `go run ./cmd/main.go` approach. <br/>
For the flags options, run `go run ./cmd/main.go -h`.

A generally nicer experience is to have it automatically restarted on code changes you can use `./run_dev.sh` provided script.
For this, make sure you have `reflex` tool installed:

- by running `go get github.com/cespare/reflex` outside of this (any any other nowadays Go Modules based) project directory
- and have `$HOME/go/bin` in your `PATH`

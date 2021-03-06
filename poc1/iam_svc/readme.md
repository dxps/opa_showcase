## The IAM Service

This is the Identity and Access Management (IAM) service, playing the authentication server role, thus being responsible with authenticating users and returing JWTs containing claims with relevant details.

<br/>

### Prereqs

Have [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) tool installed for running the database migrations.

Run the database migrations using `./run_migrate.sh 3` before starting the app. This `3` is the latest version. Check out `ops/db_migrations/` directory for the latest version number that exists.

<br/>

### Run

Use the standard `go run ./cmd/main.go` approach. <br/>
For the flags options, run `go run ./cmd/main.go -h`.

A generally nicer experience is to have it automatically restarted on code changes you can use `./run_dev.sh` provided script.
For this, make sure you have `reflex` tool installed:

- by running `go get github.com/cespare/reflex` outside of this (any any other nowadays Go Modules based) project directory
- and have `$HOME/go/bin` in your `PATH`

<br/>

### Usage

- Register a subject using:<br/>
  `curl -i -d '{ "name":"John", "email":"john@doe.com", "password":"pass1234" }' localhost:3001/v1/subjects`<br/>>
  Example:

  ```shell
  $ curl -i -d '{ "name":"John", "email":"john@doe.com", "password":"pass1234" }' localhost:3001/v1/subjects
  HTTP/1.1 201 Created
  Content-Type: application/json
  Date: Mon, 07 Jun 2021 16:50:23 GMT
  Content-Length: 133

  {"id":"9f83a904-152a-4a5a-8c4e-6cf78dfbb399","created_at":"2021-06-07T16:50:24Z","name":"John","email":"john@doe.com","active":true}

  $
  ```

- Authenticate a subject using:<br/>
  `curl -i -d '{ "email":"john@doe.com", "password":"pass1234" }' localhost:3001/v1/authenticate`<br/>
  Example:

  ```shell
  $ curl -i -d '{ "email":"john@doe.com", "password":"pass1234" }' localhost:3001/v1/authenticate
  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Mon, 07 Jun 2021 09:01:49 GMT
  Content-Length: 339

  {"access_token":"eyJhbGciOiJFUzI1NiJ9.eyJpc3MiOiJpYW0uc2VydmljZSIsInN1YiI6IjkzZWZkYWFkLWFhNDctNGJkOS04ZTcxLTU1YjMwMzBmZTAyZCIsImF1ZCI6WyJhbnlvbmUiXSwiZXhwIjoxNjIzMDYwMTA5LjkyMzgzNzIsIm5iZiI6MTYyMzA1NjUwOS45MjM4MzcyLCJpYXQiOjE2MjMwNTY1MDkuOTIzODM3Mn0.jAc6lQ5B7loogHT6sacMzj6ksi7Kmd-XTnsLw7EXtFFESRNej2A96fmmVAhWUoelR9ss8YsOU_fsffDHUzpIBQ"}

  $
  ```

  The result of a successful authentication is a JWT token.

- Add an attribute to a subject using:<br/>
  `curl -i -d '{ "name": "jobRole", "value": "SomeRole" }' localhost:3001/v1/subjects/93efdaad-aa47-4bd9-8e71-55b3030fe02d/attributes`<br/>
  Example:

  ```shell
  $ curl -i -d '{ "name": "jobRole", "value": "SomeRole" }' localhost:3001/v1/subjects/93efdaad-aa47-4bd9-8e71-55b3030fe02d/attributes
  HTTP/1.1 201 Created
  Date: Mon, 07 Jun 2021 16:45:01 GMT
  Content-Length: 0

  $
  ```

- Get the attributes of a subject using:<br/>
  `curl localhost:3001/v1/subjects/93efdaad-aa47-4bd9-8e71-55b3030fe02d/attributes`<br/>
  Example:

  ```shell
  $ curl -i localhost:3001/v1/subjects/93efdaad-aa47-4bd9-8e71-55b3030fe02d/attributes
  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Mon, 07 Jun 2021 17:00:31 GMT
  Content-Length: 40

  [{"name":"jobRole","value":"SomeRole"}]

  $
  ```

# Go-Test-viewer

This app can remotely test your Golang code in github repository.

## Requirements

- [Golang 1.14](https://golang.org/dl/): backend.
- [Angular 9](https://angular.io/guide/setup-local): frontend framework with typescript.
- [ngrx](https://ngrx.io/guide/store/install) : for state management.
- [Angular Material UI](https://material.angular.io/guide/getting-started) : for faster and easier web development.
- The requirements to install or use the above.

## Installation

1. run backend server:

   1. go to the BE foler.
   2. run `go run main.go`

2. run frontend server:
   1. go to the FE folder.
   2. run `ng server --open`

## Demo to play with this app.

[Demo](https://drive.google.com/file/d/1oEcyIN4Hxpjz2AQQv0gRY2g3VnqtaXq3/view)

## TODO(just some thoughts):

1. add database to BE
2. Redesgin the route:
   - /info : clone repo, all the file for go file, return the file and file content with github api
   - /test: to test on the repo, save the cverage for each file in the database.
   - /delete: delete the cloned repo
     ...
3. better UI dataflow cleaner data model

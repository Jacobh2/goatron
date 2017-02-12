# Goatron - A RESTful TODO service!
_- Fun weekend to learn Golang!_

This is a very simple RESTful API service for creating/viewing/updating/deleting TODO notes!

 It is based on the guide by [Saddam H](https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3#.acdv9nml1)

## Run
`go run main/main.go`
This will start the server at `http://127.0.0.1:8080/api/v1/todos`

## Build
`go build main/main.go`

# Methods
### Create TODO
Type: POST
Path: `/`
Params:
- title
- completed (1 or 0 for true or false)
### Fetch all TODOs
Type: GET
Path: `/`
### Fetch single TODO
Type: GET
Path: `/id`
### Update TODO
Type: PUT
Path: `/id`
Params:
- title
- completed (1 or 0 for true or false)
### Delete TODO
Type: DELETE
Path: `/id`
### Delete all TODOs
Type: DELETE
Path: `/`
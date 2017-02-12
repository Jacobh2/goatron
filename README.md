# Goatron - A RESTful TODO service!
_- Fun weekend to learn Golang!_

This is a very simple RESTful API service for creating/viewing/updating/deleting TODO notes!

 It is based on the guide by [Saddam H](https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3#.acdv9nml1)

## Run
`go run main/main.go`
This will start the server at `http://127.0.0.1:8080/api/v1/todos`

## Build
`go build main/main.go`

## Methods
| Method                                            | Type          | Path  | Params
| :------------------------------------------------ | :------------ | :---- | :----------------
| <i class="icon-file"></i>Create a TODO            | POST          | `/`   | title, completed
| <i class="icon-folder-open"></i> Fetch all TODOs  | GET           | `/`   |
| <i class="icon-folder-open"></i> Fetch single TODO| POST          | `/id` |
| <i class="icon-pencil"></i> Update TODO           | PUT           | `/id` | title, completed
| <i class="icon-trash"></i> Delete TODO            | DELETE        | `/id` |
| <i class="icon-trash"></i> Delete all TODOs       | DELETE        | `/`   |

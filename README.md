# Go Fiber

![go version](https://img.shields.io/badge/Go-v1.19-blue)
![app version](https://img.shields.io/badge/version-v0.1.0-orange)

[Go Fiber](https://gofiber.io/) REST API web server starter boilerplate based on clean architecture.

Referenced from gofiber's own [boilerplate](https://github.com/gofiber/boilerplate) and ItsCosmas's [gofiber-boilerplate](https://github.com/ItsCosmas/gofiber-boilerplate).

## Requirements

- Go >= v1.19.x
- Postgres >= v14.x

### Dependencies

Dependecies are canonically described in the [go.mod](/go.mod) file

### Installing dependencies and packages

To install requirements, use command

```shell
go get
```

## Development

### Running Server

To run the server, use command

```shell
go run main
```

You can also use **hot-reloading** server created using cosmtrek's [air](https://github.com/cosmtrek/air)

For this you will need to install the air module (which is basically installing a package runable in the shell).

```shell
# installing module
go install github.com/cosmtrek/air@latest

# running server
air
```

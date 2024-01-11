# Value Service
Example implementation of Value service.

Main technologies/librariers:
* Go 1.21
* Gin
* Logrus
* Testify

# Development
Start from [installing](https://go.dev/doc/install) Golang.

To install dependencies & run application

```bash
make run-all
```

To visit application click [here](http://localhost:8080/index/1000)


### Configuration
> **NOTE:** By default `config/dev.yaml` file is read.

To use another file set environment variable `CONFIG_PATH`.

Relative path can be used from the root folder of the application.

Example:
```bash
export CONFIG_PATH=./config/example.yaml
```

To only install dependencies

```bash
make install-dev
```

To only start application

```bash
make run-app
```

Go to `http://localhost:8080/index/123` to see API response.

## Testing

> **NOTE:** Tests load config from `config/test.yaml` file.

To execute tests


```bash
make test
```

# Information

## Folder Structure

```bash
.
├── cmd - app entrypoint
├── config - app config
├── pkg
│   ├── domain - app business logic
│   └── http - gin application
└── tests - test files
    ├── pkg
    │   ├── domain - business logic tests
    │   └── http - gin application tests
    ├── setup.go - test setup
    └── stubs - stubs for tests
        └── value_source.go
```


## Functionality

- [x] Get Index by Value (via HTTP)


## Author Thoughts

I've consciously decided to:

* Not set GIN mode to release - project is just to showcase some skills, not to make it production-ready application
* Not use wiregen to inject dependencies - as it just adds complexity
* Not go through building binaries and running application with `go run` directly - same as first point


Also I have to admit that it was fun, small project for me.
I am aware that my Go skills are not comparable to my 
Python abilities, but I look at the language as a tool that
I am learning along the way, and I am more than open for feedback!

# Value Service
Example implementation of Value service.

Main technologies/librariers:
* Go 1.21
* Gin
* Logrus

# Development
Start from [installing](https://go.dev/doc/install) Golang.

Installing dependency

```bash
make install-dev
```

Setting dev configuration

```bash
cp config/example.yaml config/dev.yaml
```

To start application

```bash
# Start application
make run
```

Go to `http://localhost:8080/index/123` to see API response.

## Testing

> **NOTE:** Tests load config from `config/test.yaml` file.

To execute tests


```bash
make test
```

# Information

## Author Notes
I am consciously omiting things like:

* Setting up GIN release mode
* Using wiregen to inject dependencies
* Writing HTTP application tests

## Folder Structure

```bash
.
├── Makefile
├── README.md
├── api - HTTP application module
├── common - Application shared objects
├── cmd - Application entrypoint
├── config - Application configuration files
├── domain - Application business-related logic
├── task - Task-related information
|── tests - Application tests
├── go.mod
└── go.sum
```


## Functionality

- [x] Get Index by Value

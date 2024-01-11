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
├── LICENSE
├── Makefile
├── README.md
├── cmd - app entrypoint
├── config - app config
├── pkg - app packages
│   ├── domain - business-related logic - it shoud be used
│   └── http - HTTP interface package
├── task
│   ├── input.txt
│   └── task.md
└── tests
```


## Functionality

- [x] Get Index by Value

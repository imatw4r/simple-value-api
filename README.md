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

To start application

```bash
# Start application
make run
```

Go to `http://localhost:8080/index/123` to see API response.

## Testing

> **_NOTE:_**  Ensure that .env file is used to store configuration. I

To execute tests


```bash
make test
```


# Information

## Functionality

- [x] Get Index by Value


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

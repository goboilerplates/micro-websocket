# Boilerplate for WebSocket Microservice in Go
[![Build Status](https://travis-ci.org/goboilerplates/micro-websocket.svg?branch=master)](https://travis-ci.org/goboilerplates/micro-websocket)
[![codecov](https://codecov.io/gh/goboilerplates/micro-websocket/branch/master/graph/badge.svg)](https://codecov.io/gh/goboilerplates/micro-websocket)
[![Go Report Card](https://goreportcard.com/badge/github.com/goboilerplates/micro-websocket)](https://goreportcard.com/report/github.com/goboilerplates/micro-websocket)
[![GoDoc](https://godoc.org/github.com/goboilerplates/micro-websocket?status.svg)](https://godoc.org/github.com/goboilerplates/micro-websocket)
[![](https://images.microbadger.com/badges/image/goboilerplates/micro-websocket.svg)](https://microbadger.com/images/goboilerplates/micro-websocket)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/goboilerplates/micro-websocket/blob/master/LICENSE)

## Features
- WebSocket API
- Middlewares (cors, gzip and static)
- CI with Travis
- Docker Build

## Installation

Get the micro-websocket repository

```
go get github.com/goboilerplates/micro-websocket

cd echo $GOPATH/src/github.com/goboilerplates/micro-websocket
```

And install dependencies

```
go get -u github.com/golang/dep/cmd/dep

dep ensure
```

## Running the tests

Run all tests

```
go test ./...
```

Or run all tests with coverage

```
bash script/coverage.sh
```

## Build and Run

Run main.go
``` bash
go run main.go
# serve at localhost:9000
```

Build and run native binary

``` bash
bash script/Build.sh

./micro-websocket.out
```
Build native binary for multiple platforms (darwin, windows and linux)

```
bash script/BuildMulti.sh
```

## Environment variables

```bash
    # enable production mode, default is false
    env GBP_PROMODE=true
    
     # enable compression, default is true
    env GONI_COMPRESS=false

    # set read buffer size, default equals to 1024
    env GONI_RBSIZE=2048

    # set write buffer size, default equals to 1024
    env GONI_WBSIZE=2048
```
## Docker support 

Build docker image

```
bash script/Dockerbuild.sh
```

Run docker container

```
docker run -d --name micro-websocket -p 9000:9000 goboilerplates/micro-websocket
```
## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/goboilerplates/micro-websocket/tags). 

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details


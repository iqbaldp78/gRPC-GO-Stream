# Go gRPC Streaming

---
|Branch|Build Status|Coverage Report|
|:-------|:----------|:----|
|__Main__|[![build status](https://github.com/iqbaldp78/frankie-universal-sdk)](https://github.com/iqbaldp78/frankie-universal-sdk)|[![coverage report](https://github.com/iqbaldp78/frankie-universal-sdk)](https://github.com/iqbaldp78/frankie-universal-sdk)|


---

## Objective

This repo used to gRPC Server Stream

## Requirement

- [Go 1.15 or higher](https://golang.org/dl/)

## Installation

### Local environment

- Clone this repo

```bash
git clone https://github.com/iqbaldp78/frankie-universal-sdk.git
```

- Put outside `$GOPATH`
- Run this service

```bash
cd $REPO_DIR
go run main.go
```


- Run docker

```bash
docker build .

docker run -p 80:8080 <PUT_YOUR_DOCKER_IMAGE_HERE>
```

## Development

If you are __developer__ please see this readme first and information below may give help to you

### Project Structure

This is general project structure. If you add/change file the structure please change structure below

```ANSI
├── controllers                         //package for controllers
├── initialize                          //package for init function
├── model                               //package for model data
├── utils                               //package for utilities function
├── config.json                         //configuration file for this service
├── Dockerfile                          //dockerfile for build docker image
├── go.mod                              //gomod
├── go.sum                              //generated file from gomod
├── main.go                             //main function to run project
└── README.md                           //this file
```

### Environment Variables

This is changeable environment variables. If you add/change more variable on config file please change information below

|Variable|Description|Value|
|:-------|:----------|:----|
|SERVER_PORT|Server port listen|Default `8080`|


## Usage

This repo contain 1 endpoint :

```bash
METHOD POST :
localhost:8080/isgood
```
Please refer to apidocs.md for more documentation

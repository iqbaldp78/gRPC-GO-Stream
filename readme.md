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
git clone https://github.com/iqbaldp78/gRPC-GO-Stream.git
```

- Put outside `$GOPATH`

- Running the server
```bash
cd $REPO_DIR
go run main.go
```



## Development

If you are __developer__ please see this readme first and information below may give help to you


### Recompile Proto

If you want to change your proto file scheme please recompile your proto file with command below

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./modules/pb/pb.proto
```

### Project Structure

This is general project structure. If you add/change file the structure please change structure below

```ANSI
├── modules                             //package for controllers
├───── client                           //package for init client
├───── pb                               //package proto and generated file
├───── services                         //package for logic function
├── config.json                         //configuration file for this service
├── go.mod                              //gomod
├── go.sum                              //generated file from gomod
├── main.go                             //main function to run project
└── README.md                           //this file readme
```

### Environment Variables

This is changeable environment variables. If you add/change more variable on config file please change information below

|Variable|Description|Value|
|:-------|:----------|:----|
|SERVER_PORT|Server port listen|Default `10000`|


## Usage

This repo contain 1 server gRPC and 2 client example :

```bash
cd $REPO_DIR 

- running server 

go run main.go

- running client 
    - push 
       open another terminal
       cd $REPO_DIR/modules/client/push
        go run client_push.go
         Enter Number to push: {ENTER_NUMBER_YOU_WANT}
    - stream
       open another terminal
       cd $REPO_DIR/modules/client/stream 
       go run client_stream.go
        Enter Number of treshold: {ENTER_NUMBER_YOU_WANT}
```

### Set multiple client 
just open another terminal

Another Terminal:
```bash
cd $REPO_DIR/modules/client/push
```
go ahead running the client like example above

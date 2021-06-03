# Port Service

## Overview

Microservice architecture to persist port data on mongodb database.

The Client is a simple http service that parse json data and communicate over GRPC to second microservice that  
persist and return data from the database 

## Getting started

### docker-compose

Simple run ```docker-compose up``` command.

### Run locally

```make run_client && make mongo && make run_port_server```

### APIs

To parse and update new dato to database:

```curl localhost:8080/upload -X POST -F 'file=@ports.json'```

You can retrieve data from Port ID:

- Get one port:

```curl localhost:8080/id/{:id}```  

example: ```curl localhost:8080/id/CRCAL```




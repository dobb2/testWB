# Wildberries test L0

## Description
The service has a subscription to nats-streaming, 
getting the data from where it writes to the cache and postgres. 
On a get request, for example, 
such as http://localhost:8080/getOrder/id3 gives data data on id3

### running postgres and nats-streaming with docker-compose

```
docker-compose up --build
```

### run srcipt
Running a script that reads json data from a file and sends it to nats-streaming

```
go run cmd/pub-client/main.go
```

### run service

```
go run cmd/sub-client/main.go
```

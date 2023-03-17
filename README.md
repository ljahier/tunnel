# Tunneling service


## To run server
```
go run main.go server start --port 9000
```

## To run client
You must replace **127.0.0.1:9000** from **--server** using your server url and you must replace **8081** from **--port** by your port which running your web service
```
go run main.go client --port 8081 --server 127.0.0.1:9000
```
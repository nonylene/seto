# seto

A HTTP server to bridge utilities via Unix domain socket.

## Download (amd64)

```
$ cp config.json.example ${XDG_CONFIG_HOME}/seto/config.json
$ curl -L https://github.com/nonylene/seto/releases/latest/download/seto -o seto
$ chmod 755 seto
$ seto
```


## Development

### Seto (server side)

```
$ cp config.json.example config.json
$ go run cmd/seto/main.go -config config.json
```

To perform health check:

```
$ curl --unix-socket {Socket path} http://localhost/healthCheck
```

### Setoc (client side)

TBD

## Build

```
$ go build -o build/seto cmd/seto/main.go
```

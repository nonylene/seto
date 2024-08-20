# seto

A HTTP server to bridge utilities via Unix domain socket.

## Download (amd64)

### Seto

```
$ cp config.json.example ${XDG_CONFIG_HOME}/seto/config.json
$ curl -L https://github.com/nonylene/seto/releases/latest/download/seto -o seto
$ chmod 755 seto
$ ./seto
```

### Setoc

```
$ cp config.json.example ${XDG_CONFIG_HOME}/setoc/config.json
$ curl -L https://github.com/nonylene/seto/releases/latest/download/setoc -o setoc
$ chmod 755 setoc
$ ./setoc
```

### SSH

```
Host host
  RemoteForward {Full path for your socket file} {Remote sock file}
```

You may want to set `StreamLocalBindUnlink yes` on the remote server sshd config.

## Development

```
$ cp config.json.example config.json
$ go run cmd/seto/main.go -config config.json
$ go run cmd/setoc/main.go browser -config config.json
```

To perform health check:

```
$ curl --unix-socket {Socket path} http://localhost/healthCheck
```

## Build

```
$ go build -o build/seto cmd/seto/main.go
$ go build -o build/setoc cmd/setoc/main.go
```

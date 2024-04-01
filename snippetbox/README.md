format code before commit
```sh
go fmt ./...
```

server reload
```sh
watchexec -e=go,tmpl -rc -- "go run ./cmd/web -addr=:8080"
```

Restart server on file write
```bash
watchexec -e=.go,.html -rc -- "go run ./cmd/web -addr=':4000'"
```


# hello world http-server

логирует эндпоинты `localhost:8080/example/how/we/log` в файлик `server.log`
в формате:

```json
{"level":"INFO","time":"2019-11-30T13:39:20.219+0300","caller":"hello_world/server.go:25","message":"requested path: example/how/we/log"}
```

Запуск:
```bash
go run *.go
```

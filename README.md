uuid-gen
========
To pull the code:

```
go get github.com/abhi-bit/uuid-gen
```

Run the Server:

```
go run server/uuidServer.go
```

To grab UUID's:

```
$ curl -X GET "http://127.0.0.1:12345/_uuid?count=4"
{"uuid":["0dbacda040c9ece580410bf63c156ab2"]}
{"uuid":["a60faa2040a5ebb1804e3b7a98b90e9f"]}
{"uuid":["91085f0740e76ba180a5449231a0384e"]}
{"uuid":["12c31ccc400ebb618030a8554f67451d"]}
```

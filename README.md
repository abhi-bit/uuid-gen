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

Benchmark results:

```
go test -v -bench=".*" uuid.go uuid_test.go    
=== RUN TestUUIDv4
--- PASS: TestUUIDv4 (0.00 seconds)
        uuid_test.go:12: uuid[3d87de8f-4648-4f29-9425-319f4dbfa154]
=== RUN TestUUIDRand
--- PASS: TestUUIDRand (0.00 seconds)
        uuid_test.go:20: uuid[3b57686f-f1e9-011e-dd6e-182ca109ba42]
PASS
BenchmarkUUIDv4   500000              4143 ns/op
BenchmarkUUIDRand         200000             12628 ns/op
ok      command-line-arguments  61.632s
```

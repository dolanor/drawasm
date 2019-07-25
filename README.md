# Drawasm

Simple example from [Sendil Kumar N](https://dev.to/sendilkumarn/tiny-go-to-webassembly-5168)

## Run

### Go wasm

```shell
# We generate the wasm binary with standard Go and copy its wasm_exec.js counterpart
GOOS=js GOARCH=wasm go generate ./cmd/drawasm
go run ./cmd/server/server.go
```

### TinyGo wasm

```shell
# We generate the wasm binary with TinyGo and copy its wasm_exec.js counterpart
go generate ./cmd/drawasm
go run ./cmd/server/server.go
```

The TinyGo version crashes when you draw too many segments since it doesn't have GC yet.
But then, it's lighter!

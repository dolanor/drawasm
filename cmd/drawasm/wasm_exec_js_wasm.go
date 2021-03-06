package main

// Since the filename is explicitely *_js_wasm.go, we don't have to use // +build js,wasm build tag.
//go:generate echo Generating with standard Go draw.wasm and copy the correct version of wasm_exec.js
//go:generate mkdir -p ../../static/wasm
//go:generate go build -o ../../static/wasm/draw.wasm ../drawasm/
//go:generate cp wasm_exec_js_wasm.js ../../static/js/wasm_exec.js

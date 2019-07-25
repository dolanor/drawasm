// +build tinygo
package main

//go:generate mkdir -p ../../static/wasm
//go:generate tinygo build -o ../../static/wasm/draw.wasm ../drawasm/
//go:generate cp wasm_exec.tinygo.js ../../static/js/wasm_exec.js

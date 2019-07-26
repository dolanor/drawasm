// +build tinygo

package main

//go:generate echo Generating draw.wasm with TinyGo and copy the correct version of wasm_exec.js
//go:generate mkdir -p ../../static/wasm
//go:generate tinygo build -o ../../static/wasm/draw.wasm ../drawasm/
//go:generate cp wasm_exec_tinygo.js ../../static/js/wasm_exec.js

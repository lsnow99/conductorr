#!/bin/bash

GOOS=js GOARCH=wasm go build -o dist/csl.wasm ./cmd/csl
brotli -f dist/csl.wasm
cp $(go env GOROOT)/misc/wasm/wasm_exec.js web/src/util
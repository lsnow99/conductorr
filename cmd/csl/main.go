package main

import (
	"syscall/js"

	"github.com/lsnow99/conductorr/csl"
)

func Validate(this js.Value, args []js.Value) interface{} {
	_, err := csl.Parse(args[0].String())
	if err != nil {
		return err.Error()
	}
	return ""
}

func main() {
	js.Global().Set("Validate", js.FuncOf(Validate))
	select {}
}

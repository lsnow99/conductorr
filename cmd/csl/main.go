package main

import (
	"syscall/js"

	"github.com/lsnow99/conductorr/csl"
)

func Validate(this js.Value, args []js.Value) interface{} {
	callback := args[len(args)-1:][0]
	go func() {
		_, err := csl.Parse(args[0].String())
		if err != nil {
			callback.Invoke(false, err.Error())
			return
		}
		callback.Invoke(true, js.Null())
	}()
	return nil
}

func Run(this js.Value, args []js.Value) interface{} {
	callback := args[len(args)-1:][0]
	go func() {
		sexprs, err := csl.Parse(args[0].String())
		if err != nil {
			callback.Invoke(false, err.Error())
			return
		}
		env := make(map[string]interface{})
		result, trace := csl.Eval(sexprs, env)
		if trace.Err != nil {
			callback.Invoke(false, trace.Err.Error())
			return
		}
		if list, ok := result.(csl.List); ok {
			callback.Invoke(true, js.Null(), list.Elems)
			return
		}
		callback.Invoke(true, js.Null(), result)
	}()
	return nil
}

func main() {
	js.Global().Set("Validate", js.FuncOf(Validate))
	js.Global().Set("Run", js.FuncOf(Run))
	select {}
}

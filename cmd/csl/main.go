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

func buildRelease(release js.Value) csl.List {
	return csl.List{
		Elems: []interface{}{
			release.Get("title"),
			release.Get("indexer"),
			release.Get("download_type"),
			release.Get("content_type"),
			release.Get("rip_type"),
			release.Get("resolution"),
			release.Get("encoding"),
			release.Get("seeders"),
			release.Get("age"),
			release.Get("size"),
			release.Get("runtime"),
		},
	}
}

func Run(this js.Value, args []js.Value) interface{} {
	callback := args[len(args)-1:][0]
	env := make(map[string]interface{})
	a := args[1].Get("a")
	b := args[1].Get("b")
	env["a"] = buildRelease(a)
	if !b.IsUndefined() {
		env["b"] = buildRelease(b)
	}
	go func() {
		sexprs, err := csl.Parse(args[0].String())
		if err != nil {
			callback.Invoke(false, err.Error())
			return
		}
		result, trace := csl.Eval(sexprs, env)
		defer func() {
			if err := recover(); err != nil {
				if rErr, ok := err.(error); ok {
					errStr := rErr.Error()
					callback.Invoke(false, errStr)
				} else if str, ok := err.(string); ok {
					if str == "ValueOf: invalid value" {
						callback.Invoke(false, "ValueOf: invalid value - this usually means the return value of your script could not be converted to a valid javascript value")
					}
				} else {
					callback.Invoke(false, "Unexpected panic")
				}
			}
		}()
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

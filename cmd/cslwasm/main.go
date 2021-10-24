package main

import (
	"fmt"
	"syscall/js"

	"github.com/lsnow99/conductorr/csllib"
	_ "github.com/lsnow99/conductorr/internal/csl"
)

var DefaultEnv map[string]interface{} = make(map[string]interface{})

func Validate(this js.Value, args []js.Value) interface{} {
	csl := csllib.NewCSL(true)
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

func buildRelease(release js.Value) (csllib.List, error) {
	title := strOrNil(release.Get("title"))
	if title == nil {
		return csllib.List{}, fmt.Errorf("title is nil")
	}
	indexer := strOrNil(release.Get("indexer"))
	if indexer == nil {
		return csllib.List{}, fmt.Errorf("indexer is nil")
	}
	downloadType := strOrNil(release.Get("download_type"))
	if downloadType == nil {
		return csllib.List{}, fmt.Errorf("download_type is nil")
	}
	contentType := strOrNil(release.Get("content_type"))
	if contentType == nil {
		return csllib.List{}, fmt.Errorf("content_type is nil")
	}
	ripType := strOrNil(release.Get("rip_type"))
	if ripType == nil {
		return csllib.List{}, fmt.Errorf("rip_type is nil")
	}
	resolution := strOrNil(release.Get("resolution"))
	if resolution == nil {
		return csllib.List{}, fmt.Errorf("resolution is nil")
	}
	encoding := strOrNil(release.Get("encoding"))
	if encoding == nil {
		return csllib.List{}, fmt.Errorf("encoding is nil")
	}
	seeders := intOrNil(release.Get("seeders"))
	if seeders == nil {
		return csllib.List{}, fmt.Errorf("seeders is nil")
	}
	age := intOrNil(release.Get("age"))
	if age == nil {
		return csllib.List{}, fmt.Errorf("age is nil")
	}
	size := intOrNil(release.Get("size"))
	if size == nil {
		return csllib.List{}, fmt.Errorf("size is nil")
	}
	runtime := intOrNil(release.Get("runtime"))
	if runtime == nil {
		return csllib.List{}, fmt.Errorf("runtime is nil")
	}
	return csllib.List{
		Elems: []interface{}{
			title,
			indexer,
			contentType,
			downloadType,
			ripType,
			resolution,
			encoding,
			seeders,
			age,
			size,
			runtime,
		},
	}, nil
}

func strOrNil(val js.Value) interface{} {
	if val.IsNull() || val.IsUndefined() {
		return nil
	}
	return val.String()
}

func intOrNil(val js.Value) interface{} {
	if val.IsNull() || val.IsUndefined() {
		return nil
	}
	return int64(val.Float())
}

func Execute(this js.Value, args []js.Value) interface{} {
	csl := csllib.NewCSL(true)
	callback := args[len(args)-1:][0]
	go func() {
		sexprs, err := csl.Parse(args[0].String())
		if err != nil {
			callback.Invoke(false, err.Error())
			return
		}
		result, trace := csl.Eval(sexprs, DefaultEnv)
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
		if list, ok := result.(csllib.List); ok {
			callback.Invoke(true, js.Null(), list.Elems)
			return
		}
		callback.Invoke(true, js.Null(), result)
	}()
	return nil
}

func Run(this js.Value, args []js.Value) interface{} {
	callback := args[len(args)-1:][0]
	env := make(map[string]interface{})
	a := args[1].Get("a")
	b := args[1].Get("b")
	aR, err := buildRelease(a)
	if err != nil {
		callback.Invoke(false, err.Error())
		return nil
	}
	env["a"] = aR
	if !b.IsUndefined() {
		bR, err := buildRelease(b)
		if err != nil {
			callback.Invoke(false, err.Error())
			return nil
		}
		env["b"] = bR
	}
	go func() {
		csl := csllib.NewCSL(true)
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
		if list, ok := result.(csllib.List); ok {
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
	js.Global().Set("Execute", js.FuncOf(Execute))
	select {}
}

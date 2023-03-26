//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"syscall/js"

	"github.com/lsnow99/conductorr/internal/csl"
	"github.com/lsnow99/conductorr/pkg/csllib"
)

var DefaultEnv map[string]interface{} = make(map[string]interface{})
var CorsProxyServer string

func handleError(err error) error {
   if err != nil {
      _, filename, line, _ := runtime.Caller(1)
      return fmt.Errorf("%s:%d %v", filename, line, err)
   }
   return nil
}

var AppFetcher = func(is csllib.ImportableScript, importPath string, allowInsecureRequests bool) (string, error) {
	if _, ok := is.(csllib.FileScript); ok {
		return "", handleError(fmt.Errorf("cannot resolve import from file"))
	}

  fmt.Println("hello")

	u := url.URL{}
	u.Path = "/api/v1/fetchScript"
	q := u.Query()
	q.Set("importPath", importPath)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return "", handleError(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", handleError(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", handleError(err)
	}

	respString := string(data)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return "", handleError(fmt.Errorf("api server encountered error resolving import: %s", respString))
	}

	return respString, nil
}

var PlaygroundFetcher = func(is csllib.ImportableScript, importPath string, allowInsecureRequests bool) (string, error) {
	if _, ok := is.(csllib.FileScript); ok {
		return "", handleError(fmt.Errorf("cannot resolve import from file in playground"))
	} else if _, ok := is.(csllib.ProfileScript); ok {
		return "", handleError(fmt.Errorf("cannot resolve import from local profile in playground"))
	} else if gs, ok := is.(csllib.GitScript); ok {
		u := gs.GetURL()
		return attemptProxyFetch(u, allowInsecureRequests)
	} else if ws, ok := is.(csllib.WebScript); ok {
		u := ws.GetURL()
		return attemptProxyFetch(u, allowInsecureRequests)
	} else {
		return "", handleError(fmt.Errorf("unimplemented importable script scheme or type"))
	}
}

func attemptProxyFetch(u url.URL, allowInsecureRequests bool) (string, error) {
	proxyReqUrl, err := url.Parse(CorsProxyServer)
	if err != nil {
		return "", handleError(err)
	}
	q := proxyReqUrl.Query()
	q.Set("url", u.String())
	proxyReqUrl.RawQuery = q.Encode()

	proxyReq, err := http.NewRequest("GET", proxyReqUrl.String(), nil)
	if err != nil {
		return "", handleError(err)
	}

	resp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		return "", handleError(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		if u.Scheme == "https" && allowInsecureRequests {
			u.Scheme = "http"
			return attemptProxyFetch(u, allowInsecureRequests)
		} else {
			return "", handleError(fmt.Errorf("received non 2xx response code %d", resp.StatusCode))
		}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", handleError(err)
	}

	return string(data), nil
}

func Validate(this js.Value, args []js.Value) interface{} {
	csl := csl.NewCSL()
	callback := args[len(args)-1:][0]
	go func() {
		_, err := csl.Parse(args[0].String())
    err = handleError(err)
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
	downloadType := strOrNil(release.Get("downloadType"))
	if downloadType == nil {
		return csllib.List{}, fmt.Errorf("downloadType is nil")
	}
	contentType := strOrNil(release.Get("contentType"))
	if contentType == nil {
		return csllib.List{}, fmt.Errorf("contentType is nil")
	}
	ripType := strOrNil(release.Get("ripType"))
	if ripType == nil {
		return csllib.List{}, fmt.Errorf("ripType is nil")
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
		title,
		indexer,
		downloadType,
		contentType,
		ripType,
		resolution,
		encoding,
		seeders,
		age,
		size,
		runtime,
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
	script := args[0].String()
	callback := args[len(args)-1:][0]

	go func() {
		csl := csl.NewCSL()
		cslpm := csllib.NewCSLPackageManager(PlaygroundFetcher, true)
		if err := csl.PreprocessScript(script, "", cslpm); err != nil {
      err = handleError(err)
			callback.Invoke(false, err.Error())
			return
		}

		sexprs, err := csl.Parse(script)
		if err != nil {
      err = handleError(err)
			callback.Invoke(false, err.Error())
			return
		}

		result, trace := csl.Eval(sexprs, DefaultEnv)
		defer func() {
			if err := recover(); err != nil {
				if rErr, ok := err.(error); ok {
          rErr = handleError(rErr)
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
      trace.Err = handleError(trace.Err)
			callback.Invoke(false, trace.Err.Error())
			return
		}
		if list, ok := result.(csllib.List); ok {
			callback.Invoke(true, js.Null(), list)
			return
		}
		callback.Invoke(true, js.Null(), result)
	}()
	return nil
}

func Run(this js.Value, args []js.Value) interface{} {
	var aR, bR csllib.List
	var err error

	script := args[0].String()
	callback := args[len(args)-1:][0]
	a := args[1].Get("a")
	b := args[1].Get("b")

	aR, err = buildRelease(a)
	if err != nil {
		callback.Invoke(false, err.Error())
		return nil
	}
	if !b.IsUndefined() {
		bR, err = buildRelease(b)
		if err != nil {
			callback.Invoke(false, err.Error())
			return nil
		}
	}
	go func() {
		csl := csl.NewCSL()
		cslpm := csllib.NewCSLPackageManager(AppFetcher, true)

		result, err, trace := csl.ResolveDepsAndCall(cslpm, csllib.Script{
			Code: script,
		}, aR, bR)
		if err != nil {
			callback.Invoke(false, err.Error())
			return
		}
		if trace.Err != nil {
			callback.Invoke(false, trace.Err.Error())
			return
		}
		defer func() {
			if err := recover(); err != nil {
				if rErr, ok := err.(error); ok {
					errStr := rErr.Error()
					callback.Invoke(false, errStr)
				} else if str, ok := err.(string); ok {
					if str == "ValueOf: invalid value" {
						callback.Invoke(false, fmt.Sprintf("ValueOf: invalid value - this usually means the return value of your script could not be converted to a valid javascript value. Return value: %v", result))
					}
				} else {
					callback.Invoke(false, "Unexpected panic")
				}
			}
		}()
		callback.Invoke(true, js.Null(), result)
	}()
	return nil
}

func main() {
	fmt.Printf("Using %s as cors proxy server\n", CorsProxyServer)
	js.Global().Set("Validate", js.FuncOf(Validate))
	js.Global().Set("Run", js.FuncOf(Run))
	js.Global().Set("Execute", js.FuncOf(Execute))
	select {}
}

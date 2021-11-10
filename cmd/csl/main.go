package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lsnow99/conductorr/pkg/csllib"
)

var NoCache bool
var AllowInsecureRequests bool

func main() {
	flag.BoolVar(&NoCache, "nocache", false, "If set, the dependency cache will not be used")
	flag.BoolVar(&AllowInsecureRequests, "insecure", false, "If set, the dependency manager will allow plaintext http request fallbacks")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Printf("missing required argument action (run, or get)\n")
		os.Exit(1)
	} else if flag.NArg() < 2 {
		fmt.Printf("missing required argument script file\n")
		os.Exit(1)
	}

	csl := csllib.NewCSL(true)
	action := flag.Arg(0)
	importPath := flag.Arg(1)

	switch action {
	case "run":
		cslpm := csllib.NewCSLPackageManager(func(is csllib.ImportableScript, importPath string, allowInsecureRequests bool) (string, error) {
			if NoCache {
				return is.Fetch(allowInsecureRequests)
			}
			return csllib.DefaultFetcher(is, importPath, allowInsecureRequests)
		}, AllowInsecureRequests)
		script, err := cslpm.Resolve(importPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		result, err, trace := csl.ResolveDepsAndCall(cslpm, csllib.Script{
			Code: script,
			ImportPath: importPath,
		})
		if err != nil {
			fmt.Println("Error resolving dependencies: ", err)
		}
		if trace.Err != nil {
			fmt.Println("Error evaluating csl script:")
			fmt.Println(err)
			fmt.Println("Trace:")
			fmt.Println(trace.ExprTree)
			os.Exit(1)
		}
		fmt.Printf("%v\n", result)
	case "get":
		cslpm := csllib.NewCSLPackageManager(func(is csllib.ImportableScript, importPath string, allowInsecureRequests bool) (string, error) {
			script, err := is.Fetch(allowInsecureRequests)
			if err != nil {
				return "", err
			}

			cacheDir, err := csllib.GetDefaultCacheDir()
			if err != nil {
				return "", err
			}
			filename := csllib.GetDefaultCacheName(importPath)
			filename = filepath.Join(cacheDir, filename)
			err = os.WriteFile(filename, []byte(script), os.ModePerm)
			if err != nil {
				return "", err
			}

			return script, nil
		}, AllowInsecureRequests)
		script, err := cslpm.Resolve(importPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := csl.PreprocessScript(script, importPath, cslpm); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unrecognized action %s, available actions are [get, run]\n", action)
		os.Exit(1)
	}
}

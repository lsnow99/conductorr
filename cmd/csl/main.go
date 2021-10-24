package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lsnow99/conductorr/csllib"
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
	}

	switch flag.Arg(0) {
	case "run":
		if flag.NArg() < 2 {
			fmt.Printf("missing required argument script file\n")
			os.Exit(1)
		}
		csl := csllib.NewCSL(true)
		cslpm := csllib.NewCSLPackageManager(func(is csllib.ImportableScript, importPath string, allowInsecureRequests bool) (string, error) {
			if NoCache {
				return is.Fetch(allowInsecureRequests)
			}
			return csllib.DefaultFetcher(is, importPath, allowInsecureRequests)
		}, AllowInsecureRequests)
		data, err := os.ReadFile(flag.Arg(1))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		script := string(data)
		csl.PreprocessScript(script, cslpm)
		sexprs, err := csl.Parse(script)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		env := make(map[string]interface{})
		result, trace := csl.Eval(sexprs, env)
		if trace.Err != nil {
			fmt.Println("Error evaluating csl script:")
			fmt.Println(trace.Err)
			fmt.Println("Trace:")
			fmt.Println(trace.ExprTree)
			os.Exit(1)
		}
		fmt.Printf("%v\n", result)
	case "get":

	}
}

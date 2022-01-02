package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lsnow99/conductorr/internal/csl"
	"github.com/lsnow99/conductorr/pkg/csllib"
)

var NoCache bool
var AllowInsecureRequests bool

func main() {
	runCmd := flag.NewFlagSet("run", flag.ExitOnError)
	runCmd.BoolVar(&NoCache, "nocache", false, "If set, the dependency cache will not be used")
	runCmd.BoolVar(&AllowInsecureRequests, "insecure", false, "If set, the dependency manager will allow plaintext http request fallbacks")
	
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getCmd.BoolVar(&AllowInsecureRequests, "insecure", false, "If set, the dependency manager will allow plaintext http request fallbacks")
	flag.Parse()

	if len(os.Args) < 2 {
        fmt.Println("missing required argument action (run, or get)")
        os.Exit(1)
    }

	csl := csl.NewCSL()
	action := os.Args[1]

	switch action {
	case "run":
		if err := runCmd.Parse(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if len(runCmd.Args()) < 1 {
			fmt.Fprintln(os.Stderr, "no argument provided to run")
			os.Exit(1)
		}

		importPath := runCmd.Arg(0)

		runArgs, err := parseArgs(csl, runCmd.Args()[1:])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		cslpm := csllib.NewCSLPackageManager(func(is csllib.ImportableScript, importPath string, allowInsecureRequests bool) (string, error) {
			if NoCache {
				return is.Fetch(allowInsecureRequests)
			}
			return csllib.DefaultFetcher(is, importPath, allowInsecureRequests)
		}, AllowInsecureRequests)
		script, err := cslpm.Resolve(importPath)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		result, err, trace := csl.ResolveDepsAndCall(cslpm, csllib.Script{
			Code: script,
			ImportPath: importPath,
		}, runArgs...)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error resolving dependencies: ", err)
			os.Exit(1)
		}
		if trace.Err != nil {
			fmt.Fprintln(os.Stderr, "Error evaluating csl script:")
			fmt.Fprintln(os.Stderr, trace.Err)
			fmt.Fprintln(os.Stderr, "Trace:")
			fmt.Fprintln(os.Stderr, trace.ExprTree)
			os.Exit(1)
		}
		fmt.Printf("%v\n", result)
	case "get":
		if err := runCmd.Parse(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if len(runCmd.Args()) < 1 {
			fmt.Fprintln(os.Stderr, "no argument provided to get")
			os.Exit(1)
		}

		importPath := runCmd.Arg(0)

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

// parseArgs takes a csl instance and a slice of raw string arguments and evaluates them to
// usable CSL types.
func parseArgs(csl *csllib.CSL, rawArgs []string) ([]interface{}, error) {
	parsedArgs := make([]interface{}, 0, len(rawArgs))
	for _, arg := range rawArgs {
		sexprs, err := csl.Parse(arg)
		if err != nil {
			return nil, err
		}
		res, trace := csl.Invoke(sexprs)
		if trace.Err != nil {
			return nil, trace.Err
		}
		parsedArgs = append(parsedArgs, res)
	}
	return parsedArgs, nil
}

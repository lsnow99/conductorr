package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/blang/semver/v4"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Printf("expected 2 arguments (min version, testing version), got %d", len(args))
		os.Exit(1)
	}

	minV, err := semver.Make(args[0])
	if err != nil {
		fmt.Println(fmt.Errorf("error parsing version %s: %v", args[0], err))
		os.Exit(1)
	}

	testV, err := semver.Make(args[1])
	if err != nil {
		fmt.Println(fmt.Errorf("error parsing version %s: %v", args[1], err))
		os.Exit(1)
	}

	if !testV.GTE(minV) {
		fmt.Printf("%s < %s\n", testV, minV)
		os.Exit(1)
	}
	fmt.Printf("%s > %s\n", testV, minV)
}

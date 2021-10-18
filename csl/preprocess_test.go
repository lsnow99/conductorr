package csl

import (
	"fmt"
	"testing"
)

func TestPreprocessScript1(t *testing.T) {
	deps := PreprocessScript(`(define x 7)`)
	if len(deps.nodes) != 1 {
		t.Fatalf("got dependency graph with %d nodes, expected 1", len(deps.nodes))
	}
}

func TestPreprocessScript2(t *testing.T) {
	deps := PreprocessScript(`(import "github.com/lsnow99/myscripts:main.csl" main)`)
	fmt.Println(deps)
}

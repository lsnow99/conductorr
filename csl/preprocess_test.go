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

func TestDFS(t *testing.T) {
	deps := DepGraph{}
	deps.edges = make(map[Node][]NodeReference)
	deps.nodes = make([]*Node, 0)
	a := new(Node)
	a.script = "a"
	b := new(Node)
	b.script = "b"
	c := new(Node)
	c.script = "c"
	a2b := NodeReference{
		refNode: b,
	}
	b2c := NodeReference{
		refNode: c,
	}
	c2a := NodeReference{
		refNode: a,
	}
	deps.nodes = append(deps.nodes, a, b, c)
	deps.edges[*a] = append(deps.edges[*a], a2b)
	deps.edges[*b] = append(deps.edges[*b], b2c)
	deps.edges[*c] = append(deps.edges[*c], c2a)
	if !deps.dfs(a, b) {
		t.Fatalf("Expected to find a path from b to a but did not")
	}
}

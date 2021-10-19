package csl

import (
	"fmt"
	"testing"
)

func TestPreprocessScript1(t *testing.T) {
	deps := PreprocessScript(`(define x 7)`)
	if len(deps.Nodes) != 1 {
		t.Fatalf("got dependency graph with %d nodes, expected 1", len(deps.Nodes))
	}
}

func TestPreprocessScript2(t *testing.T) {
	deps := PreprocessScript(`(import "github.com/lsnow99/myscripts:main.csl" main)`)
	fmt.Println(deps)
}

func TestDFS(t *testing.T) {
	deps := DepGraph{}
	deps.Edges = make(map[*Node][]NodeReference)
	deps.Nodes = make([]*Node, 0)
	a := new(Node)
	b := new(Node)
	c := new(Node)
	a2b := NodeReference{
		RefNode: b,
	}
	b2c := NodeReference{
		RefNode: c,
	}
	c2a := NodeReference{
		RefNode: a,
	}
	deps.Nodes = append(deps.Nodes, a, b, c)
	deps.Edges[a] = append(deps.Edges[a], a2b)
	deps.Edges[b] = append(deps.Edges[b], b2c)
	deps.Edges[c] = append(deps.Edges[c], c2a)
	if !deps.DFS(a, b) {
		t.Fatalf("Expected to find a path from b to a but did not")
	}
}

